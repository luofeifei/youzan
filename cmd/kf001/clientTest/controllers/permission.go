package controllers

import (
	"base/client/pkg/comm"
	"base/client/pkg/middlewares"
	"base/pkg/app"
	"github.com/gin-gonic/gin"
)

// CheckPermission 验证用户权限
func CheckPermission(context *gin.Context) {
	Uid, err := middlewares.GetUid(context)
	if err == nil {
		if HasPermission(Uid, context) {
			context.Next()
			return
		}
	}
	comm.NewResponse(app.Tips, comm.ResponseMsg{Icon: app.MsgError}, app.PermissionDeniedMessage).End(context)
	context.Abort()
	return
}

func HasPermission(ID int64, ctx *gin.Context) bool {
	/*
	ok, _ := client.ServerAdminSys().PermissionsByRequest(ID, ctx.Request.URL.Path, ctx.Request.Method)
	return ok
	 */
	return true
}
