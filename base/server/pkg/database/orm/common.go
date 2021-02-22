package orm

import (
	"github.com/jinzhu/gorm"
)

// 分页条件
type PageWhere struct {
	Where string
	Value []string
}

// 分页参数返回
type IndexPage struct {
	Total    int64  `json:"total"`    //总数
	Page     int64  `json:"page"`     //页数
	PageSize int64  `json:"pageSize"` //每页显示数
	OrderKey string `json:"orderKey"` // 默认排序字段 -filed1,+field2,field3 (-Desc 降序)
}

// Create
func Create(value interface{}) error {
	return Master().Create(value).Error
}

// Save
func Save(value interface{}) error {
	return Master().Save(value).Error
}

// Updates
func Updates(where interface{}, value interface{}) error {
	return Master().Model(where).Omit("id").Updates(value).Error
}

// Delete
func DeleteByModel(model interface{}) (count int64, err error) {
	db := Master().Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// Delete
func DeleteByWhere(model, where interface{}) (count int64, err error) {
	db := Master().Where(where).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// Delete
func DeleteByID(model interface{}, id int64) (count int64, err error) {
	db := Master().Where("id=?", id).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// Delete
func DeleteByIDS(model interface{}, ids []int64) (count int64, err error) {
	db := Master().Where("id in (?)", ids).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// First
func FirstByID(elect string, out interface{}, id int64) (notFound bool, err error) {
	if len(elect) > 0 {
		err = Slave().Select(elect).First(out, id).Error
	} else {
		err = Slave().First(out, id).Error
	}
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

func Count(model, where interface{}, elect string) (out interface{}) {
	if elect != "" {
		var info []string
		err := Slave().Model(model).Where(where).Limit(1).Pluck(elect, &info).Error
		if err != nil {
			return ""
		}
		return info[0]
	} else {
		err = Slave().Model(model).Where(where).Select("count(*)").Limit(1).Count(&out).Error
	}
	return ""
}

// First
func First(elect string, where interface{}, out interface{}) (notFound bool, err error) {
	if len(elect) > 0 {
		err = Slave().Select(elect).Where(where).First(out).Error
	} else {
		err = Slave().Where(where).First(out).Error
	}
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

// Find
func Find(elect string, where interface{}, out interface{}, orders ...string) error {
	db := Slave().Where(where)
	if len(elect) > 0 {
		db = db.Select(elect)
	}
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Find(out).Error
}

// Scan
func Scan(elect string, model, where interface{}, out interface{}) (notFound bool, err error) {
	db := Slave().Model(model)
	if len(elect) > 0 {
		db = db.Select(elect)
	}
	err = db.Where(where).Scan(out).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

// ScanList
func ScanList(elect string, model, where interface{}, out interface{}, orders ...string) error {
	db := Slave().Model(model).Where(where)
	if len(elect) > 0 {
		db = db.Select(elect)
	}
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Scan(out).Error
}

// PluckList
func PluckList(elect string, model, where interface{}, fieldName string, out interface{}) error {
	db := Slave().Model(model).Where(where)
	if len(elect) > 0 {
		db = db.Select(elect)
	}
	return db.Pluck(fieldName, out).Error
}
