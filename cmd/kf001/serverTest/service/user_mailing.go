package service

import (
	"base/model/modelSql/mall_user"
	"base/server/pkg/database/orm"
)

// 获取用户邮寄信息
func UserMailingData(elect string, req mall_user.UserMailing) (res mall_user.UserMailing, err error) {
	_, err = orm.First(elect, &req, &res)
	return
}

// 获取用户邮寄信息列表
func UserMailingList(elect string, m mall_user.UserMailing) (infos []mall_user.UserMailing, err error) {
	err = orm.Find(elect, m, &infos)
	return
}