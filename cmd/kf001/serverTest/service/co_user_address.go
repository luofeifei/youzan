package service

import (
	"base/model/modelSql/mall_co"
	"base/pkg/app"
	"base/server/pkg/database/orm"
	"base/server/pkg/pager"
)

func UserAddressPage(elect string, m mall_co.CoUserAddress, pageInfo orm.IndexPage) (list []mall_co.CoUserAddress, total int64, err error) {
	var where = make(pager.Where)
	if m.Coid > 0 {
		where["coid"] = m.Coid
	}
	total, err = pager.New(pager.NewGormDriver(), pageInfo).Fields(elect).SetIndex(m.TableName()).Where(where).Find(&list)
	for i, val := range list {
		_ = orm.Find("type,`default`", &mall_co.CoUserAddressLink{AddressID: val.ID}, &list[i].Link)
	}
	return list, total, err
}

func UserAddressData(elect string, m mall_co.CoUserAddress) (res mall_co.CoUserAddress, err error) {
	_, err = orm.First(elect, &m, &res)
	if err == nil {
		_ = orm.Find("type,`default`", &mall_co.CoUserAddressLink{AddressID: m.ID}, &res.Link)
	}
	return
}

// 设置默认地址
func UserAddressDefaultUp(id, coid int64, m []mall_co.CoUserAddressLink) (err error) {
	for _, val := range m {
		if val.Default == 2 {
			var ids []int64
			_ = orm.Slave().Table("co_user_address a").Joins("INNER JOIN co_user_address_link b ON a.id = b.address_id").Where("a.id != ? AND a.coid = ? AND b.type = ?", id, coid, val.Type).Pluck("b.address_id", &ids)
			if len(ids) > 0 {
				orm.Master().Table("co_user_address_link").Where("address_id IN (?)", ids).Update("default", 1)
			}
		}
	}
	return
}

func UserAddressSave(m mall_co.CoUserAddress) (res mall_co.CoUserAddress, err error) {
	if m.ID > 0 {
		notFound, _ := orm.FirstByID("id", &res, m.ID)
		if notFound == true {
			err = app.Err("未找到记录")
		} else {
			_, _ = orm.DeleteByWhere(&mall_co.CoUserAddressLink{}, &mall_co.CoUserAddressLink{AddressID: m.ID})
			err = orm.Updates(&res, m)
			if err == nil {
				err = UserAddressDefaultUp(m.ID, m.Coid, m.Link)
			}
			return m, err
		}
	} else {
		err = orm.Create(&m)
		if err == nil {
			err = UserAddressDefaultUp(m.ID, m.Coid, m.Link)
		}
		return m, err
	}
	return res, err
}

func UserAddressDelete(m mall_co.CoUserAddress) (err error) {
	count, _ := orm.DeleteByWhere(&mall_co.CoUserAddress{}, m)
	if count == 0 {
		return app.Err("用户地址删除失败")
	} else {
		_, _ = orm.DeleteByWhere(&mall_co.CoUserAddressLink{}, &mall_co.CoUserAddressLink{AddressID: m.ID})
	}
	return
}
