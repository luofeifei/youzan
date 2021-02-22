package handler

import (
	"base/model/imp/serverUser"
	"base/model/modelSql/mall_user"
	"base/pkg/app"
	"base/tools"
	"base/tools/encrypt"
	"context"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"server/service"
)

type UserSysImp struct {
}

/**
TODO: 用户登录
QinWong 2020.09.14
*/
// types 登录类型
func (imp *UserSysImp) Login(ctx context.Context, req serverUser.ReqPass) (res serverUser.UserData, err error) {
	var User mall_user.User
	var UserList mall_user.UserList
	var UserData mall_user.UserData
	if err := app.Unmarshal(req.UserList, &UserList); err != nil {
		return res, err
	}
	// 读取账户列表
	result, err := service.GetUserList(app.GetContext(ctx, "elect"), UserList)
	if err != nil {
		return res, app.Err("未找到相关账户")
	}
	// 读取用户信息
	User.ID = result.UId
	userInfo, _ := service.GetUser("pass, encrypt, state", User)
	// 用户名登录
	if userInfo.State != 1 {
		return res, app.Err("账户状态异常,请联系管理员！")
	}
	if UserList.Type == 1 || result.Type == 1 {
		if encrypt.Md5(req.Password+userInfo.Encrypt) != userInfo.Pass || req.Password == "" {
			return res, app.Err("用户名或密码错误")
		}
	}
	resultData, _ := service.GetUserDate("id, real_name, nick_name, avatar", UserData)
	err = app.Unmarshal(resultData, &res)
	return res, err
}

// 用户退出
func (imp *UserSysImp) LoginOut(ctx context.Context, req serverUser.UserLogin) (res serverUser.ResultEmpty, err error) {
	var info mall_user.UserLogin
	err = app.Unmarshal(req, &info)
	if err != nil {
		return res, err
	}
	return res, service.LogoutUpData(info)
}

// 设置用户登录记录
func (imp *UserSysImp) LoginData(ctx context.Context, req serverUser.UserLogin) (res serverUser.ResultEmpty, err error) {
	var info mall_user.UserLogin
	if err := app.Unmarshal(req, &info); err != nil {
		return res, err
	}
	ip, _ := current.GetClientIPFromContext(ctx)
	info.IP = tools.Ip2long(ip)
	_, err = service.SaveUserLogin("", info)
	if err != nil {
		return res, app.Err("未找到相关账户")
	}
	return res, err
}

/**
TODO: 用户信息
QinWong 2020.08.21
*/
// 获取用户详细信息
func (imp *UserSysImp) UserDataData(ctx context.Context, req serverUser.UserData) (res serverUser.UserData, err error) {
	var info mall_user.UserData
	err = app.Unmarshal(req, &info)
	if err == nil {
		result, err := service.GetUserDate(app.GetContext(ctx, "elect"), info)
		if err != nil {
			return res, err
		}
		err = app.Unmarshal(result, &res)
		return res, err
	}
	return res, err
}

// 获取用户绑定信息
func (imp *UserSysImp) UserListData(ctx context.Context, req serverUser.UserList) (res serverUser.UserList, err error) {
	var info mall_user.UserList
	err = app.Unmarshal(req, &info)
	if err == nil {
		result, err := service.GetUserList(app.GetContext(ctx, "elect"), info)
		if err != nil {
			return res, app.Err("未找到相关账户")
		}
		err = app.Unmarshal(result, &res)
		return res, err
	}
	return res, err
}

// 获取用户登录设备列表
func (imp *UserSysImp) UserLoginList(ctx context.Context, req serverUser.UserLogin) (res serverUser.ResUserLogin, err error) {
	var info mall_user.UserLogin
	err = app.Unmarshal(req, &info)
	if err == nil {
		result, err := service.UserLoginList(app.GetContext(ctx, "elect"), info)
		if err != nil {
			return res, err
		}
		res.Count = int32(len(result))
		err = app.Unmarshal(result, &res.List)
		return res, err
	}
	return res, err
}

/**
TODO: 用户邮寄信息
QinWong 2020.09.09
*/
// 获取用户邮寄信息
func (imp *UserSysImp) UserMailingData(ctx context.Context, req serverUser.UserMailing) (res serverUser.UserMailing, err error) {
	var info mall_user.UserMailing
	err = app.Unmarshal(req, &info)
	if err == nil {
		result, err := service.UserMailingData(app.GetContext(ctx, "elect"), info)
		if err != nil {
			return res, app.Err("未找到相关账户")
		}
		err = app.Unmarshal(result, &res)
		return res, err
	}
	return res, err
}

// 获取用户邮寄信息列表
func (imp *UserSysImp) UserMailingList(ctx context.Context, req serverUser.UserMailing) (res serverUser.ResUserMailing, err error) {
	var info mall_user.UserMailing
	err = app.Unmarshal(req, &info)
	if err == nil {
		result, err := service.UserMailingList(app.GetContext(ctx, "elect"), info)
		if err != nil {
			return res, err
		}
		res.Count = int32(len(result))
		err = app.Unmarshal(result, &res.List)
		return res, err
	}
	return res, err
}

/**
TODO: 地区查询
QinWong 2020.08.29
*/
// 查询地区所以下级数据
func (imp *UserSysImp) UserRegionList(ctx context.Context, req serverUser.UserRegion) (res serverUser.ResUserRegion, err error) {
	result, err := service.GetUserRegionList(app.GetContext(ctx, "elect"), req.ParentId)
	if err != nil {
		return res, err
	}
	res.Count = int32(len(result))
	err = app.Unmarshal(result, &res.List)
	return res, err
}

// 查询地区数据
func (imp *UserSysImp) UserRegionData(ctx context.Context, req serverUser.UserRegion) (res serverUser.UserRegion, err error) {
	result, err := service.GetUserRegion(app.GetContext(ctx, "elect"), req.Id)
	if err != nil {
		return res, err
	}
	_ = app.Unmarshal(result, &res)
	return res, nil
}

// 查询地区名字
func (imp *UserSysImp) UserRegionName(ctx context.Context, req serverUser.ResUserRegionName) (res serverUser.Result, err error) {
	result, err := service.UserRegionName("name", req.Province, req.City, req.County)
	if err != nil {
		return res, err
	}
	res.Code = 1
	res.Msg = result
	return res, nil
}

// 根据经纬度返回地区信息
func (imp *UserSysImp) UserRegionLngLat(ctx context.Context, req serverUser.UserRegion) (res serverUser.ResUserRegion, err error) {
	result, err := service.UserRegionLngLat("name", req.Loc[0], req.Loc[1])
	if err != nil {
		return res, err
	}
	res.Count = int32(len(result))
	err = app.Unmarshal(result, &res.List)
	return res, err
}
