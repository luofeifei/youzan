package user

import (
	"base/client/pkg/comm"
	"base/client/pkg/middlewares"
	"base/model/client"
	"base/model/imp/serverUser"
	"github.com/gin-gonic/gin"
)

// @Tags system_user
// @Summary 用户退出
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Router /user/out [post]
func UserOut(ctx *gin.Context) {
	Uid, err := middlewares.GetUid(ctx)
	source := ctx.GetHeader("source")
	if err != nil {
		comm.TipResponse(ctx, err,0,nil)
		return
	}
	res, err := client.ServerUserSys().LoginOut(serverUser.UserLogin{Uid: Uid, Platform: source})
	comm.TipResponse(ctx, err,0,res)
}