package service
//
//import (
//	"base/model/modelSql"
//	"base/pkg/app"
//	"base/server/pkg/database/orm"
//)
//
//// 根据用户名获取信息
//func GetUser(elect string, req modelSql.User) (res modelSql.User, err error) {
//	_, err = orm.First(elect, &req, &res)
//	return
//}
//
//// 获取用户详细资料
//func GetUserDate(elect string, req modelSql.UserData) (res modelSql.UserData, err error) {
//	_, err = orm.First(elect, &req, &res)
//	return
//}
//
//// 获取用户绑定信息
//func GetUserList(elect string, req modelSql.UserList) (res modelSql.UserList, err error) {
//	_, err = orm.First(elect, &req, &res)
//	return
//}
//
//// 添加用户登录记录
//func SaveUserLogin(elect string, req modelSql.UserLogin) (res modelSql.UserLogin, err error) {
//	notFound, _ := orm.First(elect, &modelSql.UserLogin{UId: req.UId, Source: req.Source}, &res)
//	if notFound == true {
//		err = orm.Create(&req)
//	} else {
//		err = orm.Updates(&res, &req)
//	}
//	return
//}
//
//func LogoutUpData(req modelSql.UserLogin) (err error) {
//	count, _ := orm.DeleteByWhere(&req, req)
//	if count == 0 {
//		return app.Err("删除失败")
//	}
//	return
//}
