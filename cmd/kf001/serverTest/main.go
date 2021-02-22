package main

import (
	"base/model/imp/serverUser"
	"base/model/mongoSql"
	"base/pkg/app"
	"base/pkg/log"
	"base/pkg/redis"
	"base/server/pkg/database/mongo"
	"base/server/pkg/database/orm"
	"base/server/pkg/queue"
	"github.com/TarsCloud/TarsGo/tars"
	"server/handler"
	"server/subscriber"
)

func start() {
	// 初始化服务
	log.Start()
	orm.Start()
	redis.Start()
	mongo.Start(mongoSql.CoIndex())
	if queue.Start() {
		go subscriber.RegisterSubscriber()
	}
}

func main() {
	start()
	// admin := tars.Admin{}
	// _, _ = admin.Notify("tars.pprof 8081")
	imp := new(handler.UserSysImp)
	apps := new(serverUser.UserSys)
	apps.AddServantWithContext(imp, app.Cfg.App+"."+app.Cfg.Server+".sysObj")
	tars.Run()
}
