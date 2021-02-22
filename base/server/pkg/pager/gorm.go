package pager

import (
	"base/server/pkg/database/orm"
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
	"strings"
)

// Gorm 查询
type Gorm struct {
	conn  *gorm.DB
	query string
	args  []interface{}
	limit int64
	skip  int64
	index string
	sorts string
	elect string
}

// NewGormDriver gorm 分页实现
func NewGormDriver() *Gorm {
	return &Gorm{
		conn: orm.Slave(),
	}
}

// Where 构建查询条件
func (orm *Gorm) Where(kv Where) {
	for k, v := range kv {
		if strings.Index(k, "?") == -1 {
			orm.query = orm.query + fmt.Sprintf("AND `%s` = ? ", k)
		} else {
			orm.query = orm.query + fmt.Sprintf("AND %s ", k)
		}
		orm.args = append(orm.args, v)
	}
}

// Index table 表名
func (orm *Gorm) Select(elect string) {
	orm.elect = elect
}

// Limit 每页数量
func (orm *Gorm) Limit(limit int64) {
	orm.limit = limit
}

// Skip 跳过数量
func (orm *Gorm) Skip(skip int64) {
	orm.skip = skip
}

// Index table 表名
func (orm *Gorm) Index(index string) {
	orm.index = index
}

// Sort 排序
func (orm *Gorm) Sort(kv map[string]Sort) {
	for k, v := range kv {
		if v == Asc {
			orm.sorts += fmt.Sprintf("AND %s ASC ", k)
		} else {
			orm.sorts += fmt.Sprintf("AND %s DESC ", k)
		}
	}
	orm.sorts = strings.TrimPrefix(orm.sorts, "AND")
}

// Find 从数据库查询数据
func (orm *Gorm) Find(data interface{}) error {
	db := orm.conn.Table(orm.index)
	if len(orm.elect) > 0 {
		db = db.Select(orm.elect)
	}
	orm.sorts = strings.TrimSpace(orm.sorts)
	orm.query = strings.TrimSpace(strings.TrimPrefix(orm.query, "AND"))
	return db.Where(orm.query, orm.args...).Limit(orm.limit).Offset(orm.skip).Order(orm.sorts).Find(data).Error
}

// Count 计算指定查询条件的总数量
func (orm *Gorm) Count() int64 {
	var count int64
	orm.conn.Table(orm.index).Where(orm.query, orm.args...).Count(&count)
	return count
}

// SetTyp sql 对数据查询的类型不敏感
func (orm *Gorm) SetTyp(typ reflect.Type) {
	return
}