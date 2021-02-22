package database

import (
	"base/pkg/app"
	"base/tools"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func createConnectionURL(username, password, addr, dbName string) string {
	//app.Logger().Debug(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, addr, dbName))
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, addr, dbName)
}

func MysqlConnect(dbName, addr, username, password, maxIdle, maxOpen string) (orm *gorm.DB, err error) {
	orm, err = gorm.Open("mysql", createConnectionURL(username, password, addr, dbName))
	if err != nil {
		app.Println("database connect error, you can't use orm support: ")
		//app.Logger().Warn("database connect error, you can't use orm support: ")
		//app.Logger().Warn(err)
		return orm, err
	}
	//if app.Cfg.LogLevel == "DEBUG" {
	//	orm.LogMode(true)
	//}
	orm.LogMode(true)
	// 启用Logger，显示详细日志
	//orm.SetLogger(log.NewGormLogger())
	// 关闭复数表名，如果设置为true，`User`表的表名就会是`user`，而不是`users`
	orm.SingularTable(true)
	if idle, err := tools.StringToInt(maxIdle); err == nil {
		orm.DB().SetMaxIdleConns(idle)
	}
	if open, err := tools.StringToInt(maxOpen); err == nil {
		orm.DB().SetMaxOpenConns(open)
	}
	return orm, nil
}
