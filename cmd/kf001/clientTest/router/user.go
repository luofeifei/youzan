package router

import (
	"client/controllers/user"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	SysRouter := Router.Group("user")
	{
		SysRouter.POST("out", user.UserOut) //用户退出
	}
	return SysRouter
}
