package handler
//
//import (
//	"base/model/imp/serverUser"
//	"base/model/modelSql"
//	"base/pkg/app"
//	"base/pkg/des"
//	"context"
//	"github.com/TarsCloud/TarsGo/tars/util/current"
//	"server/service"
//)
//
//type UserSysImp struct {
//}
//
//var User modelSql.User
//var UserList modelSql.UserList
//var UserData modelSql.UserData
//var UserLogin modelSql.UserLogin
//
//// 用户登录
//// types 登录类型
//func (imp *UserSysImp) Login(ctx context.Context, req serverUser.Req_Pass) (res serverUser.UserData, err error) {
//	SelectDb, err := app.UnmarshalElect(req.UserList, &UserList)
//	if err == nil {
//		// 读取账户列表
//		result, err := service.GetUserList(SelectDb, UserList)
//		if err != nil {
//			return res, app.Err("未找到相关账户")
//		}
//		// 读取用户信息
//		User.ID = result.UId
//		userInfo, _ := service.GetUser([]string{"pass", "encrypt", "state"}, User)
//		// 用户名登录
//		if userInfo.State != 1 {
//			return res, app.Err("账户状态异常,请联系管理员！")
//		}
//		if UserList.Type == 1 || result.Type == 1 {
//			if des.Md5(req.Password+userInfo.Encrypt) != userInfo.Pass || req.Password == "" {
//				return res, app.Err("用户名或密码错误")
//			}
//		}
//		resultData, _ := service.GetUserDate([]string{"id", "realname", "nickname", "avatar"}, UserData)
//		_ = app.Unmarshal(resultData, &res)
//		return res, err
//	}
//	return res, nil
//}
//
//// 用户退出
//func (imp *UserSysImp) LoginOut(ctx context.Context, req serverUser.UserLogin) (res serverUser.ResultEmpty, err error) {
//	err = app.Unmarshal(req, &UserLogin)
//	if err == nil {
//		if err := service.LogoutUpData(UserLogin); err != nil {
//			// 退出成功
//		}
//	}
//	return res, nil
//}
//
//// 获取用户绑定信息
//func (imp *UserSysImp) ListData(ctx context.Context, req serverUser.UserList) (res serverUser.UserList, err error) {
//	SelectDb, err := app.UnmarshalElect(req, &UserList)
//	if err == nil {
//		result, err := service.GetUserList(SelectDb, UserList)
//		if err != nil {
//			return serverUser.UserList{}, app.Err("未找到相关账户")
//		}
//		_ = app.Unmarshal(result, &res)
//		return res, nil
//	}
//	return res, nil
//}
//
//// 设置用户登录记录
//func (imp *UserSysImp) LoginData(ctx context.Context, req serverUser.UserLogin) (res serverUser.ResultEmpty, err error) {
//	SelectDb, err := app.UnmarshalElect(req, &UserList)
//	if err == nil {
//		UserLogin.IP, _ = current.GetClientIPFromContext(ctx)
//		_, err := service.SaveUserLogin(SelectDb, UserLogin)
//		if err != nil {
//			return res, app.Err("未找到相关账户")
//		}
//	}
//	return res, err
//}
