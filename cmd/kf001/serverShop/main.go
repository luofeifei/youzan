package main

import (
	"base/model/imp/serverShop"
	"base/pkg/app"
	"base/pkg/log"
	"base/server/pkg/database/orm"
	"github.com/TarsCloud/TarsGo/tars"
	"server/handler"
)

func start() {
	// 初始化服务
	log.Start()
	orm.Start()
	//redis.Start()
	//mongo.Start(mongoSql.ShopIndex())
}

func main() {
	start()
	imp := new(handler.ShopSysImp)
	apps := new(serverShop.ShopSys)
	apps.AddServantWithContext(imp, app.Cfg.App+"."+app.Cfg.Server+".sysObj")
	tars.Run()
}
