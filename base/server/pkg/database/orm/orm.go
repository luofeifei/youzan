package orm

import (
	"base/pkg/app"
	"base/pkg/config"
	"base/pkg/log"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync/atomic"
)

type (
	// Orm gorm 连接对象, 包含Master和Slaves, 由配置决定, Slaves 使用 atomic 包进行循环获取
	Orm struct {
		Master *gorm.DB
		Slaves []*gorm.DB
	}

	connInfo struct {
		Addr     string `yaml:"addr"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbname"`
		MaxIdle  int    `yaml:"max_idle"`
		MaxOpen  int    `yaml:"max_open"`
	}

	configs struct {
		Master connInfo   `yaml:"master"`
		Slaves []connInfo `yaml:"slave"`
	}
)

var (
	orm       = &Orm{}
	slavesLen int
	err       error
	cursor    int64
	conf      configs
)

func createConnectionURL(username, password, addr, dbName string) string {
	app.Logger().Debug(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, addr, dbName))
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, addr, dbName)
}

// Start 启动数据库
func Start() {
	err := config.Config().Bind(app.CfgName, "database", &conf, func() {
		conf = configs{}
		orm = &Orm{}
		if config.Config().Bind(app.CfgName, "database", &conf, nil) == nil {
			initStart()
		}
	})
	if  err != nil {
		app.Logger().Debug(fmt.Sprintf("mysql 配置错误 err: %s", err.Error()))
		return
	}
	initStart()
}

func initStart() {
	orm.Master, err = gorm.Open("mysql", createConnectionURL(conf.Master.Username, conf.Master.Password, conf.Master.Addr, conf.Master.DbName))
	if err != nil {
		app.Logger().Warn("database connect error, you can't use orm support: ")
		app.Logger().Warn(err)
		return
	}
	if app.Cfg.LogLevel == "DEBUG" {
		orm.Master.LogMode(true)
	}
	// 启用Logger，显示详细日志
	orm.Master.SetLogger(log.NewGormLogger())
	// 关闭复数表名，如果设置为true，`User`表的表名就会是`user`，而不是`users`
	orm.Master.SingularTable(true)
	orm.Master.DB().SetMaxIdleConns(conf.Master.MaxIdle)
	orm.Master.DB().SetMaxOpenConns(conf.Master.MaxOpen)
	for _, slave := range conf.Slaves {
		connect, err := gorm.Open("mysql", createConnectionURL(slave.Username, slave.Password, slave.Addr, slave.DbName))
		if err != nil {
			app.Logger().Warn("database connect error, you can't use orm support")
			app.Logger().Warn(err)
		}
		if app.Cfg.LogLevel == "DEBUG" {
			connect.LogMode(true)
		}
		connect.SetLogger(log.NewGormLogger())
		connect.SingularTable(true)
		orm.Slaves = append(orm.Slaves, connect)
	}
	slavesLen = len(orm.Slaves)
}

// Slave 获得一个从库连接对象, 使用 atomic.AddInt64 计算调用次数，然后按 Slave 连接个数和次数进行取模操作之后获取指定index的Slave
func Slave() *gorm.DB {
	rs := atomic.AddInt64(&cursor, 1)
	return orm.Slaves[rs%int64(slavesLen)]
}

// Master 获得主库连接
func Master() *gorm.DB {
	return orm.Master
}