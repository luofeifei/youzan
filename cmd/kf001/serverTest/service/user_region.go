package service

import (
	"base/model/mongoSql"
	"base/server/pkg/commom"
	"base/server/pkg/database/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

// 查询地区所以下级数据
func GetUserRegionList(elect string, ParentId int32) (infos []mongoSql.UserRegion, err error) {
	err = mongo.Collection(&mongoSql.UserRegion{}).Fields(elect).Where(bson.M{"parent_id": ParentId}).FindMany(&infos)
	return
}

// 查询地区数据
func GetUserRegion(elect string, Id int32) (res mongoSql.UserRegion, err error) {
	// 列子
	var m mongoSql.UserRegion
	m.ID = Id
	err = mongo.Collection(&mongoSql.UserRegion{}).Fields(elect).Where(commom.ToBson(m)).FindOne(&res)
	return
}

// 获取地区名称拼接
func UserRegionName(elect string, province int32, city int32, county int32) (res string, err error) {
	var result []string
	var info mongoSql.UserRegion
	if province > 0 {
		if err = mongo.Collection(&info).Where(bson.M{"_id": province}).Fields(elect).FindOne(&info); err == nil {
			result = append(result, info.Name)
		}
	}
	if city > 0 {
		if err = mongo.Collection(&info).Where(bson.M{"_id": city}).Fields(elect).FindOne(&info); err == nil {
			result = append(result, info.Name)
		}
	}
	if county > 0 {
		if err = mongo.Collection(&info).Where(bson.M{"_id": county}).Fields(elect).FindOne(&info); err == nil {
			result = append(result, info.Name)
		}
	}
	return strings.Join(result, ", "), nil
}

// 根据经纬度返回地区信息
func UserRegionLngLat(elect string, Lat, Lng float64) (res []mongoSql.UserRegion, err error) {
	var info mongoSql.UserRegion
	err = mongo.Collection(&mongoSql.UserRegion{}).Where(bson.M{"loc": bson.M{"$near": []float64{Lng, Lat}}}).FindOne(&info)
	if err == nil {
		res = append(res, info)
		for {
			info, err = GetUserRegion("", info.ParentID)
			if err == nil {
				res = append(res, info)
			} else {
				break
			}
		}
	}
	return res, nil
}
