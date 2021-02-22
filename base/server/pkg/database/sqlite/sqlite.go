package sqlite

import (
	"base/pkg/app"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type (
	Database struct {
		ID        string `gorm:"primary_key;column:id;" json:"id" form:"id"`
		CreatedAt int    `gorm:"column:created_at;index:created_at" json:"created_at"`
		UpdatedAt int    `gorm:"column:updated_at;index:updated_at" json:"updated_at"`
	}
)

var (
	sqLite *gorm.DB
	err    error
)

func Start(dbName string) {
	sqLite, err = gorm.Open("sqlite3", dbName)
	if err != nil {
		app.Logger().Warn("database connect error, you can't use sqllite support: ")
		app.Logger().Warn(err)
	}
	sqLite.LogMode(false)
}

func Get() *gorm.DB {
	return sqLite
}