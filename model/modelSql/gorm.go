package modelSql

import (
	"base/model/client"
	"base/pkg/unique"
	"github.com/jinzhu/gorm"
	"time"
)

type Database struct {
	ID        int64 `gorm:"primary_key;column:id;" json:"id"`
	CreatedAt int   `gorm:"column:created_at;type:int(10);not null" json:"created_at"` // 添加时间
	UpdatedAt int   `gorm:"column:updated_at;type:int(10);not null" json:"updated_at"` // 更新时间
}

func (m *Database) BeforeCreate(scope *gorm.Scope) error {
	return BeforeAdd(scope, m.ID, false)
}

func (m *Database) BeforeUpdate(scope *gorm.Scope) error {
	return BeforeUp(scope)
}

// c false 使用ID发号器 true sonyflake 唯一值
func BeforeAdd(scope *gorm.Scope, id int64, c bool) error {
	if id == 0 {
		if c == false {
			// 链接到 grpc 获取ID值
			resId, err := client.GetID(scope.TableName())
			if err != nil {
				return err
			}
			_ = scope.SetColumn("id", resId)
		} else {
			// 生成 唯一值
			_ = scope.SetColumn("id", unique.ID())
		}
	}
	t := time.Now().Unix()
	_ = scope.SetColumn("created_at", t)
	_ = scope.SetColumn("updated_at", t)
	return nil
}

func BeforeUp(scope *gorm.Scope) error {
	t := time.Now().Unix()
	scope.Set("gorm:update_column", true)
	_ = scope.SetColumn("updated_at", t)
	return nil
}
