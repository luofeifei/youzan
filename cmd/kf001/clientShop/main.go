package main

import (
	"base/client/pkg/jwt"
	"base/client/pkg/middlewares"
	"base/client/pkg/validator"
	"base/pkg/app"
	"base/pkg/log"
	"base/pkg/redis"
	_ "client/docs"
	"client/router"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func start() {
	// 初始化服务
	log.Start()
	jwt.Start()
	redis.Start()
	// 将 gin 的验证器替换为 v10 版本
	binding.Validator = new(validator.Validator)
}

// @title Go 商城接口
// @version 1.0.0
// @description 提供商品添加、修改等所有相关操作
// @contact.name Roger
// @license.name Apache 2.0
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http https
// @host 192.168.31.114:9006
// @BasePath /api/Shop
func main() {
	start()
	engine := gin.New()
	// 注册公用的中间件
	engine.Use(middlewares.CORS)
	// swag 文档访问地址
	if gin.Mode() != gin.ReleaseMode {
		engine.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	ApiGroup := engine.Group("/api/Shop")
	// 前端页面API
	//router.InitFrontShopRouter(ApiGroup)
	//ApiGroup.Use(middlewares.VerifyAuth).Use(middlewares.VerifyLock)
	//ApiGroup.Use(middlewares.VerifyAuth)
	//ApiGroup.Use(controllers.CheckPermission)
	router.InitShopRouter(ApiGroup)
	mux := &tars.TarsHttpMux{}
	mux.HandleFunc("/", engine.ServeHTTP)
	tars.AddHttpServant(mux, app.Cfg.App+"."+app.Cfg.Server+".shopObj")
	tars.Run()
}
