package database

import (
	"base/pkg/app"
	"base/pkg/unique"
	"base/tools"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type (
	// CollectionInfo 集合包含的连接信息和查询等操作信息
	CollectionInfo struct {
		client   *mongo.Client
		Database *mongo.Database
		Table    *mongo.Collection
		filter   bson.M
		limit    int64
		skip     int64
		sort     bson.M
		fields   bson.M
	}
	IndexKey struct {
		Name string `json:"name"` // 字段名
		Key  string `json:"key"`  // 1(ASC) -1 (DESC) 2D 2DSphere GeoHayStack Hashed Text
	}
	IndexData struct {
		Keys   []IndexKey `json:"keys"`
		Unique bool       `json:"unique"` // 是否唯一
		Exp    int32      `json:"exp"`    // 过期时间
	}
)

func MongoConnect(dbName, addr, username, password, maxConnIdleTime, maxPoolSize string) (db *CollectionInfo, err error) {
	mongoOptions := options.Client()
	if idle, err := tools.StringToInt(maxConnIdleTime); err == nil {
		mongoOptions.SetMaxConnIdleTime(time.Duration(idle) * time.Second)
	}
	if size, err := tools.StringToInt(maxPoolSize); err == nil {
		mongoOptions.SetMaxPoolSize(uint64(size))
	}
	if username != "" && password != "" {
		mongoOptions.SetAuth(options.Credential{Username: username, Password: password})
	}
	client, errs := mongo.NewClient(mongoOptions.ApplyURI("mongodb://" + addr))
	if errs != nil {
		app.Logger().Warn("mongo connect error, you can't use orm support")
		app.Logger().Warn(errs)
		return
	} else {
		app.Logger().Debug(fmt.Sprintf("mongo connect: %s %s", addr, dbName))
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		app.Logger().Error(err)
		return
	}
	if err = client.Ping(context.Background(), nil); err != nil {
		app.Logger().Debug(fmt.Sprintf("mongo Ping异常: %s", err.Error()))
		return
	}
	data := client.Database(dbName)
	return &CollectionInfo{
		client:   client,
		Database: data,
		filter:   make(bson.M),
	}, err
}

// 转换为索引格式
func getIndexModel(data []IndexData) []mongo.IndexModel {
	var index []mongo.IndexModel
	for _, val := range data {
		var opt = options.Index()
		opt.SetBackground(true)
		opt.SetUnique(val.Unique) //是否唯一
		if val.Exp > 0 {
			opt.SetExpireAfterSeconds(val.Exp) //设置过期时间
		}
		Keys := make(bson.M)
		for _, k := range val.Keys {
			if k.Key == "1" || k.Key == "-1" {
				Key, _ := strconv.Atoi(k.Key)
				Keys[k.Name] = Key
			} else {
				Keys[k.Name] = k.Key
			}
		}
		index = append(index, mongo.IndexModel{Keys: Keys, Options: opt})
	}
	return index
}

// 转换

// 获取 客户端
func (collection *CollectionInfo) GetClient() *mongo.Client {
	return collection.client
}

// Collection 得到一个mongo操作对象
func (collection *CollectionInfo) Collection(table Table) *CollectionInfo {
	collection.Table = collection.Database.Collection(table.TableName())
	return collection
}

// SetTable 设置集合名称
func (collection *CollectionInfo) SetTable(name string) *CollectionInfo {
	collection.Table = collection.Database.Collection(name)
	return collection
}

// Where 条件查询, bson.M{"field": "value"}
func (collection *CollectionInfo) Where(m bson.M) *CollectionInfo {
	collection.filter = m
	return collection
}

// Limit 限制条数
func (collection *CollectionInfo) Limit(n int64) *CollectionInfo {
	collection.limit = n
	return collection
}

// Skip 跳过条数
func (collection *CollectionInfo) Skip(n int64) *CollectionInfo {
	collection.skip = n
	return collection
}

// Sort 排序 bson.M{"created_at":-1}
func (collection *CollectionInfo) Sort(sorts bson.M) *CollectionInfo {
	collection.sort = sorts
	return collection
}

// Fields 指定查询字段  field1,field2
func (collection *CollectionInfo) Fields(fields string) *CollectionInfo {
	sorts := make(bson.M)
	if len(fields) > 0 {
		for _, v := range strings.Split(fields, ",") {
			sorts[strings.TrimSpace(v)] = 1
		}
		collection.fields = sorts
	}
	return collection
}

// CreateMany 批量创建索引
// mongo.Collection(&req).CreateMany([]mongo.IndexData{{Name: "uid", Key: "1", Unique: true}, {Name: "created_at", Key: "-1"}})
func (collection *CollectionInfo) CreateMany(req []IndexData) (res []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err = collection.Table.Indexes().CreateMany(ctx, getIndexModel(req), options.CreateIndexes().SetMaxTime(10*time.Second))
	return res, err
}

// InsertOne 写入单条数据
func (collection *CollectionInfo) InsertOne(document interface{}) (res *mongo.InsertOneResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err = collection.Table.InsertOne(ctx, BeforeCreate(document))
	return res, err
}

// InsertMany 写入多条数据
func (collection *CollectionInfo) InsertMany(documents interface{}) (res *mongo.InsertManyResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var data []interface{}
	data = BeforeCreate(documents).([]interface{})
	res, err = collection.Table.InsertMany(ctx, data)
	return res, err
}

// UpdateOrInsert 存在更新,不存在写入, documents 里边的文档需要有 _id 的存在
func (collection *CollectionInfo) UpdateOrInsert(documents interface{}) (res *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var upSert = true
	res, err = collection.Table.UpdateMany(ctx, documents, &options.UpdateOptions{Upsert: &upSert})
	return res, err
}

// UpdateOne 更新一条
func (collection *CollectionInfo) UpdateOne(document interface{}) (res *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err = collection.Table.UpdateOne(ctx, collection.filter, bson.M{"$set": BeforeUpdate(document)})
	return res, err
}

// UpdateMany 更新多条
func (collection *CollectionInfo) UpdateMany(document interface{}) (res *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err = collection.Table.UpdateMany(ctx, collection.filter, bson.M{"$set": BeforeUpdate(document)})
	return res, err
}

// FindOne 查询一条数据
func (collection *CollectionInfo) FindOne(document interface{}) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := collection.Table.FindOne(ctx, collection.filter, &options.FindOneOptions{
		Skip:       &collection.skip,
		Sort:       collection.sort,
		Projection: collection.fields,
	})
	err = result.Decode(document)
	return err
}

// FindMany 查询多条数据
func (collection *CollectionInfo) FindMany(documents interface{}) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.Table.Find(ctx, collection.filter, &options.FindOptions{
		Skip:       &collection.skip,
		Limit:      &collection.limit,
		Sort:       collection.sort,
		Projection: collection.fields,
	})
	if err != nil {
		return err
	}
	defer result.Close(ctx)
	val := reflect.ValueOf(documents)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Slice {
		return app.Err("result argument must be a slice address")
	}
	slice := reflect.MakeSlice(val.Elem().Type(), 0, 0)
	itemTyp := val.Elem().Type().Elem()
	for result.Next(ctx) {
		item := reflect.New(itemTyp)
		err := result.Decode(item.Interface())
		if err != nil {
			app.Logger().Error(err)
			break
		}
		slice = reflect.Append(slice, reflect.Indirect(item))
	}
	val.Elem().Set(slice)
	return nil
}

// Delete 删除数据,并返回删除成功的数量
func (collection *CollectionInfo) Delete() (count int64, err error) {
	if collection.filter == nil || len(collection.filter) == 0 {
		return 0, app.Err("you can't delete all documents, it's very dangerous")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, errs := collection.Table.DeleteMany(ctx, collection.filter)
	if errs != nil {
		return 0, errs
	}
	return result.DeletedCount, errs
}

// Count 根据指定条件获取总条数
func (collection *CollectionInfo) Count() (count int64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	count, err = collection.Table.CountDocuments(ctx, collection.filter)
	return count, err
}

// BeforeCreate 创建数据前置操作
func BeforeCreate(document interface{}) interface{} {
	val := reflect.ValueOf(document)
	typ := reflect.TypeOf(document)
	switch typ.Kind() {
	case reflect.Ptr:
		return BeforeCreate(val.Elem().Interface())
	case reflect.Array, reflect.Slice:
		var sliceData = make([]interface{}, val.Len(), val.Cap())
		for i := 0; i < val.Len(); i++ {
			sliceData[i] = BeforeCreate(val.Index(i).Interface()).(bson.M)
		}
		return sliceData
	case reflect.Struct:
		var data = make(bson.M)
		for i := 0; i < typ.NumField(); i++ {
			tag := typ.Field(i).Tag.Get("bson")
			if tag != "-" {
				data[tag] = val.Field(i).Interface()
			}
		}
		if val.FieldByName("ID").Type() == reflect.TypeOf(primitive.ObjectID{}) {
			data["_id"] = primitive.NewObjectID()
		}
		if val.FieldByName("ID").Kind() == reflect.String && val.FieldByName("ID").Interface() == "" {
			data["_id"] = unique.ID()
		}
		if IsIntn(val.FieldByName("ID").Kind()) && val.FieldByName("ID").Interface().(int64) == 0 {
			data["_id"] = unique.ID()
		}
		now := time.Now().Unix()
		data["created_at"] = now
		data["updated_at"] = now
		return data
	default:
		if val.Type() == reflect.TypeOf(bson.M{}) {
			if !val.MapIndex(reflect.ValueOf("_id")).IsValid() {
				val.SetMapIndex(reflect.ValueOf("_id"), reflect.ValueOf(primitive.NewObjectID()))
			}
			val.SetMapIndex(reflect.ValueOf("created_at"), reflect.ValueOf(time.Now().Unix()))
			val.SetMapIndex(reflect.ValueOf("updated_at"), reflect.ValueOf(time.Now().Unix()))
		}
		return val.Interface()
	}
}

// BeforeUpdate 更新数据前置操作
func BeforeUpdate(document interface{}) interface{} {
	val := reflect.ValueOf(document)
	typ := reflect.TypeOf(document)
	switch typ.Kind() {
	case reflect.Ptr:
		return BeforeUpdate(val.Elem().Interface())
	case reflect.Array, reflect.Slice:
		var sliceData = make([]interface{}, val.Len(), val.Cap())
		for i := 0; i < val.Len(); i++ {
			sliceData[i] = BeforeUpdate(val.Index(i).Interface()).(bson.M)
		}
		return sliceData
	case reflect.Struct:
		var data = make(bson.M)
		for i := 0; i < typ.NumField(); i++ {
			if !isZero(val.Field(i)) {
				tag := strings.Split(typ.Field(i).Tag.Get("bson"), ",")[0]
				if tag != "_id" && tag != "-" {
					data[tag] = val.Field(i).Interface()
				}
			}
		}
		data["updated_at"] = time.Now().Unix()
		return data
	default:
		if val.Type() == reflect.TypeOf(bson.M{}) {
			val.SetMapIndex(reflect.ValueOf("updated_at"), reflect.ValueOf(time.Now().Unix()))
		}
		return val.Interface()
	}
}

// IsIntn 是否为整数
func IsIntn(p reflect.Kind) bool {
	return p == reflect.Int || p == reflect.Int32 || p == reflect.Int64 || p == reflect.Uint64 || p == reflect.Uint32
}

func isZero(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		// return !value.Bool()
		return false
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}
