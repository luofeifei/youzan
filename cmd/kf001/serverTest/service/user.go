package service

import (
	"base/model/modelSql/mall_user"
	"base/pkg/app"
	"base/server/pkg/database/orm"
)

// 根据用户名获取信息
func GetUser(elect string, req mall_user.User) (res mall_user.User, err error) {
	_, err = orm.First(elect, &req, &res)
	return
}

// 获取用户详细资料
func GetUserDate(elect string, req mall_user.UserData) (res mall_user.UserData, err error) {
	_, err = orm.First(elect, &req, &res)
	return
}

// 获取用户绑定信息
func GetUserList(elect string, req mall_user.UserList) (res mall_user.UserList, err error) {
	_, err = orm.First(elect, &req, &res)
	return
}

// 获取用户登录设备列表
func UserLoginList(elect string, m mall_user.UserLogin) (infos []mall_user.UserLogin, err error) {
	err = orm.Find(elect, m, &infos)
	return
}

// 添加用户登录记录
func SaveUserLogin(elect string, req mall_user.UserLogin) (res mall_user.UserLogin, err error) {
	notFound, _ := orm.First(elect, &mall_user.UserLogin{UId: req.UId, Platform: req.Platform}, &res)
	if notFound == true {
		err = orm.Create(&req)
	} else {
		err = orm.Updates(&res, &req)
	}
	return
}

func LogoutUpData(req mall_user.UserLogin) (err error) {
	count, _ := orm.DeleteByWhere(&req, req)
	if count == 0 {
		return app.Err("删除失败")
	}
	return
}
