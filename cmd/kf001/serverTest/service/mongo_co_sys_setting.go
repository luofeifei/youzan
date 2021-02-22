package service

import (
	"base/model/client"
	"base/model/mongoSql"
	"base/pkg/app"
	"base/server/pkg/commom"
	"base/server/pkg/database/mongo"
	"base/server/pkg/database/orm"
	"base/server/pkg/pager"
	"go.mongodb.org/mongo-driver/bson"
)

// 读取列表 配置数据规则 列表
func SettingPage(elect string, m mongoSql.CoSysSetting, pageInfo orm.IndexPage) (list []mongoSql.CoSysSetting, total int64, err error) {
	var where = make(pager.Where)
	if m.MenuID > 0 {
		where["menu_id"] = m.MenuID
	}
	if m.Name != "" {
		where["name"] = m.Name
	}
	total, err = pager.New(pager.NewMongoDriver(), pageInfo).Fields(elect).SetIndex(m.TableName()).Where(where).Find(&list)
	return list, total, err
}

// 获取配置数据规则数据
func SettingData(elect string, m mongoSql.CoSysSetting) (res mongoSql.CoSysSetting, err error) {
	err = mongo.Collection(&m).Fields(elect).Where(commom.ToBson(m)).FindOne(&res)
	return
}

// 修改添加菜单 配置数据规则
func SettingSave(m mongoSql.CoSysSetting) (res mongoSql.CoSysSetting, err error) {
	if m.ID > 0 {
		if count, _ := mongo.Collection(&res).Where(bson.M{"_id": m.ID}).Count(); count == 0 {
			return m, app.Err("未找到记录")
		} else {
			_, err = mongo.Collection(&m).Where(bson.M{"_id": m.ID}).UpdateOne(&m)
			return m, err
		}
	} else {
		if count, _ := mongo.Collection(&res).Where(bson.M{"name": m.Name}).Count(); count == 0 {
			m.ID, err = client.GetID(m.TableName())
			if err == nil {
				_, err = mongo.Collection(&m).InsertOne(&m)
				return m, err
			}
		} else {
			return m, app.Err("存在重复记录")
		}
	}
	return res, err
}

// 删除配置数据规则
func SettingDelete(m mongoSql.CoSysSetting) (err error) {
	result, err := mongo.Collection(&m).Where(commom.ToBson(m)).Delete()
	if result == 0 {
		return app.Err("删除失败")
	}
	return
}
