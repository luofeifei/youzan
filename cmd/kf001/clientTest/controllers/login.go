package controllers

import (
	"base/client/pkg/captcha"
	"base/client/pkg/comm"
	"base/client/pkg/jwt"
	"base/client/pkg/validator"
	"base/model/client"
	"base/model/imp/serverUser"
	"base/pkg/app"
	"client/request"
	"client/tools"
	"github.com/gin-gonic/gin"
)

// @Tags base
// @Summary 用户登录
// @description 用户登录接口
// @accept json
// @Produce json
// @Param platform header string true "接口访问来源"
// @Param data body request.Login true "date"
// @Success 200 {object} comm.Response
// @Router /login [post]
func Login(ctx *gin.Context) {
	var form request.Login
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	platform := ctx.GetHeader("platform")
	if tools.IsPlatform(platform) == false {
		comm.TipResponse(ctx, app.Err("来源未知"), 0, nil)
		return
	}
	if captcha.Verify(form.CodeId, form.Code) == false {
		//comm.TipResponse(ctx, app.Err("验证码错误"), 0, nil)
		//return
	}
	if form.Type == 1 || form.Type == 2 {
		UserList := serverUser.UserList{User: form.User, Type: form.Type}
		res, err := client.ServerUserSys().Login(serverUser.ReqPass{Password: form.Pass, UserList: &UserList})
		if err == nil {
			token, err := jwt.CreateToken(res.Id, 1, "", platform)
			if err == nil {
				// 发送用户登录记录
				_, err = client.ServerUserSys().LoginData(serverUser.UserLogin{Uid: res.Id, Platform: platform, Token: token})
				if err != nil {
					comm.TipResponse(ctx, err, 0, nil)
					return
				}
				comm.NewResponse(app.Success, gin.H{"token": token, "user": res}).End(ctx)
			}
		}
		if err != nil {
			comm.TipResponse(ctx, err, 0, nil)
			return
		}
	} else {
		comm.TipResponse(ctx, app.Err("不支持的登录方式"), 0, nil)
	}
	return
}

// @Tags base
// @Summary 生成验证码
// @Produce json
// @Success 200 {object} comm.Response
// @Router /captcha [get]
func CaptchaImg(ctx *gin.Context) {
	cpat := captcha.New("")
	comm.NewResponse(app.Success, gin.H{"content": cpat.ToBase64EncodeString(), "code_id": cpat.CaptchaID}).End(ctx)
}
