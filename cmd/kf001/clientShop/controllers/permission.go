package controllers

import (
	"base/client/pkg/comm"
	"base/client/pkg/middlewares"
	"base/model/client"
	"base/model/imp/serverCo"
	"base/pkg/app"
	"github.com/gin-gonic/gin"
)

// CheckPermission 验证用户访问企业权限
func CheckPermission(ctx *gin.Context) {
	coId, uid, err := middlewares.GetCoIdUid(ctx)

	if err == nil {
		if HasPermission(coId, uid, ctx) {
			ctx.Set("coid", coId)
			ctx.Next()
			return
		}
	}
	comm.NewResponse(app.Tips, comm.ResponseMsg{Icon: app.MsgError}, app.PermissionDeniedMessage).End(ctx)
	ctx.Abort()
	return
}

func HasPermission(coId, uid int64, ctx *gin.Context) bool {
	ok, _ := client.ServerCoUser().PermissionsByRequest(serverCo.Permissions{Coid: coId, Uid: uid, Path: ctx.Request.URL.Path, Method: ctx.Request.Method})
	return ok.Code == 1
}
