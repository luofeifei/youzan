// Package pager 分页工具
//  目前可使用具体实现
//  var driver pager.Driver
//  driver = NewMongoDriver()
//  driver = NewGormDriver()
//  pager.New(driver).SetIndex(c.entity.TableName()).Find()

package pager

import (
	"base/server/pkg/database/orm"
	"strings"
)

// Where 查询条件
type Where map[string]interface{}

// Sort 排序
type Sort int

const (
	// Desc 降序
	Desc Sort = -1
	// Asc 升序
	Asc = 1
)

// Driver 查询驱动, Pagination会在真正查询的时候传入参数
type Driver interface {
	// 查询条件
	Where(kv Where)
	// 每页数量
	Limit(limit int64)
	// 需要查询的字段
	Select(elect string)
	// 跳过行
	Skip(skip int64)
	// Index 索引名称,可以是表名,或者比如 es 的 index 名称等,标识具体资源集合的东西
	Index(index string)
	// 排序
	Sort(kv map[string]Sort)
	// 查询具体操作
	Find(data interface{}) error
	// 计算总数
	Count() int64
}

// Pagination 分页工具
type Pagination struct {
	defaultWhere Where
	page         int64
	pageSize     int64
	sorts        string
	elect        string
	index        string
	driver       Driver
}

// New 新的分页工具实例
func New(driver Driver, pageInfo orm.IndexPage) *Pagination {
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 10
	}
	return &Pagination{driver: driver, page: pageInfo.Page, pageSize: pageInfo.PageSize, sorts: pageInfo.OrderKey}
}

// Select 设置需要查询的字段
func (pagination *Pagination) Fields(elect string) *Pagination {
	pagination.elect = elect
	return pagination
}

// SetIndex 设置集合或表名，有的驱动可能需要这个, 比如 elastic search
func (pagination *Pagination) SetIndex(index string) *Pagination {
	pagination.index = index
	return pagination
}

// Where 传入默认查询参数, 该传入参数将会在这个实例中一直被使用
func (pagination *Pagination) Where(kv Where) *Pagination {
	pagination.defaultWhere = kv
	return pagination
}

// Sort 排序 sorts=-filed1,+field2,field3
func (pagination *Pagination) Sort(order string) *Pagination {
	pagination.sorts = order
	return pagination
}

// Find 查询数据
//  structure 不需要传类似 []struct 这样的类型, 直接传入 struct 就行, 非指针
func (pagination *Pagination) Find(structure interface{}) (total int64, err error) {
	pagination.driver.Limit(pagination.pageSize)
	pagination.driver.Sort(ParseSorts(pagination.sorts))
	pagination.driver.Index(pagination.index)
	pagination.driver.Skip(pagination.pageSize * (pagination.page - 1))
	pagination.driver.Select(pagination.elect)
	pagination.driver.Where(pagination.defaultWhere)
	err = pagination.driver.Find(structure)
	if err == nil {
		total = pagination.driver.Count()
	}
	return total, err
}

// ParseSorts 解析排序字段, 传入规则为  sorts=-filed1,+field2,field3
//  "-"号标识降序
//  "+"或者无符号标识升序
func ParseSorts(query string) map[string]Sort {
	var sortMap = make(map[string]Sort)
	if query == "" {
		return sortMap
	}
	sorts := strings.Split(query, ",")
	for _, sort := range sorts {
		if strings.HasPrefix(sort, "-") {
			sortMap[strings.TrimPrefix(sort, "-")] = Desc
		} else {
			sortMap[strings.TrimPrefix(sort, "+")] = Asc
		}
	}
	return sortMap
}
