package mongo

import (
	"base/pkg/app"
	"base/pkg/config"
	baseMongo "base/pkg/mongo"
	"base/pkg/unique"
	"base/server/pkg/database"
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

var (
	client *mongo.Client
	conf   configs
)

type (
	// CollectionInfo 集合包含的连接信息和查询等操作信息
	CollectionInfo struct {
		Database *mongo.Database
		Table    *mongo.Collection
		filter   bson.M
		limit    int64
		skip     int64
		sort     bson.M
		fields   bson.M
	}
	configs struct {
		URL             string `yaml:"url"`
		Database        string `yaml:"database"`
		MaxConnIdleTime int    `yaml:"max_conn_idle_time"`
		MaxPoolSize     int    `yaml:"max_pool_size"`
		Username        string `yaml:"username"`
		Password        string `yaml:"password"`
	}
)

// Start 启动 mongo
func Start(index map[string]interface{}) {
	err := config.Config().Bind(app.CfgName, "mongo", &conf, func() {
		conf = configs{}
		if config.Config().Bind(app.CfgName, "mongo", &conf, nil) == nil {
			initStart()
		}
	})
	if err != nil {
		app.Logger().Debug(fmt.Sprintf("mongo 配置错误 err: %s", err.Error()))
		return
	}
	initStart()
	// 开始创建索引
	if len(index) > 0 {
		var num = 0
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for name, val := range index {
			var data []baseMongo.IndexData
			err = app.Unmarshal(val, &data)
			if err == nil {
				_, err := client.Database(conf.Database).Collection(name).Indexes().CreateMany(ctx, getIndexModel(data), options.CreateIndexes().SetMaxTime(10*time.Second))
				if err == nil {
					num++
				}
			}
		}
		app.Logger().Warn(fmt.Sprintf("mongo 创建索引: %d 表完成", num))
	}
}

func initStart() {
	var err error
	mongoOptions := options.Client()
	mongoOptions.SetMaxConnIdleTime(time.Duration(conf.MaxConnIdleTime) * time.Second)
	mongoOptions.SetMaxPoolSize(uint64(conf.MaxPoolSize))
	if conf.Username != "" && conf.Password != "" {
		mongoOptions.SetAuth(options.Credential{Username: conf.Username, Password: conf.Password})
	}
	client, err = mongo.NewClient(mongoOptions.ApplyURI(conf.URL))
	if err != nil {
		app.Logger().Warn("mongo connect error, you can't use orm support")
		app.Logger().Warn(err)
		return
	} else {
		app.Logger().Debug(fmt.Sprintf("mongo connect: %s %s", conf.URL, conf.Database))
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
}

// 转换为索引格式
func getIndexModel(data []baseMongo.IndexData) []mongo.IndexModel {
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

// Collection 得到一个mongo操作对象
func Collection(table database.Table) *CollectionInfo {
	db := client.Database(conf.Database)
	return &CollectionInfo{
		Database: db,
		Table:    db.Collection(table.TableName()),
		filter:   make(bson.M),
	}
}

// Database 获取数据库连接
func Database(name ...string) *CollectionInfo {
	var db *mongo.Database
	if len(name) == 1 {
		db = client.Database(name[0])
	}
	db = client.Database(conf.Database)
	collection := &CollectionInfo{
		Database: db,
		filter:   make(bson.M),
	}
	return collection
}

// 转换

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
func (collection *CollectionInfo) CreateMany(req []baseMongo.IndexData) (res []string, err error) {
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

// AggregateMany 聚合查询
func (collection *CollectionInfo) AggregateMany(documents interface{}, where []bson.M) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var pipeline interface{}
	if where == nil {
		pipeline = []bson.M{collection.filter}
	} else {
		pipeline = where
	}
	result, err := collection.Table.Aggregate(ctx, pipeline)
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
