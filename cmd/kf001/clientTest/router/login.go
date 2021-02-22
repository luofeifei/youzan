package router

import (
	"client/controllers"
	"github.com/gin-gonic/gin"
)

func InitLoginRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	Router.POST("login", controllers.Login)       // 用户登录
	Router.GET("captcha", controllers.CaptchaImg) // 获取验证码
	return Router
}
