package mongoSql

import (
	mongo "base/pkg/mongo"
)

// mall_sys 表索引
func SysIndex() map[string]interface{} {
	var req = make(map[string]interface{})
	req["sys_config"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "alias", Key: "1"}, {Name: "updated_at", Key: "1"}}}}
	req["user_region"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "parent_id", Key: "1"}}}}
	req["user_login_qr"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "code", Key: "text"}}, Unique: true, Exp: 3600}}
	return req
}

// mall_co 表索引
func CoIndex() map[string]interface{} {
	var req = make(map[string]interface{})
	req["co_menu"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "type", Key: "1"}}}, {Keys: []mongo.IndexKey{{Name: "state", Key: "1"}}}, {Keys: []mongo.IndexKey{{Name: "is_display", Key: "1"}}}}
	req["co_menu_api"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "path", Key: "1"}, {Name: "method", Key: "1"}}}}
	req["co_menu_group"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "pid", Key: "1"}, {Name: "role_id", Key: "1"}}}}
	req["co_sys_plugin"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "type", Key: "1"}}}, {Keys: []mongo.IndexKey{{Name: "model", Key: "1"}}}}
	req["co_sys_setting"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "menu_id", Key: "1"}}}, {Keys: []mongo.IndexKey{{Name: "state", Key: "1"}}}}
	req["co_sys_binding"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "type_id", Key: "1"}, {Name: "type", Key: "1"}}}}
	req["co_sys_binding_vip"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "type", Key: "1"}, {Name: "type_id", Key: "1"}, {Name: "limit_id", Key: "1"}}}}
	req["co_sys_module"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "plugin_id", Key: "1"}, {Name: "use_pos", Key: "1"}}}}
	req["co_user_loc"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "type", Key: "1"}, {Name: "of_id", Key: "1"}}}, {Keys: []mongo.IndexKey{{Name: "loc", Key: "2d"}}}}
	req["co_user_vip"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "coid", Key: "1"}, {Name: "due_date", Key: "1"}}}, {Keys: []mongo.IndexKey{{Name: "type", Key: "1"}, {Name: "limit_id", Key: "1"}}}}
	req["co_user_staff"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "coid", Key: "1"}}}, {Keys: []mongo.IndexKey{{Name: "uid", Key: "1"}}}, {Keys: []mongo.IndexKey{{Name: "group", Key: "1"}}}}
	req["co_user_points_rule"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "coid", Key: "1"}, {Name: "rule_type", Key: "1"}}}}
	req["co_user_setting"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "coid", Key: "1"}, {Name: "setting_id", Key: "1"}}}}
	req["co_user_shop_feature"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "coid", Key: "1"}}}, {Keys: []mongo.IndexKey{{Name: "page_type", Key: "1"}}}, {Keys: []mongo.IndexKey{{Name: "platform", Key: "1"}}}}
	req["co_user_shop_common"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "coid", Key: "1"}}}, {Keys: []mongo.IndexKey{{Name: "page_type", Key: "1"}}}, {Keys: []mongo.IndexKey{{Name: "platform", Key: "1"}}}}
	return req
}

// mall_sys 表索引
func FundIndex() map[string]interface{} {
	var req = make(map[string]interface{})
	req["fund_cash_flow"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "code", Key: "text"}}, Unique: true, Exp: 3600}}
	return req
}

// co_Shop_goods 文档索引
func ShopIndex() map[string]interface{} {
	var req = make(map[string]interface{})
	//req["co_shop_goods"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "coid", Key: "1"},{Name: "goods_rule.", Key: "1"}}}}
	req["co_shop_goods"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "coid", Key: "1"}}}}
	req["shop_goods_stock_discount"] = []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "coid", Key: "1"}}}}
	req["shop_goods_distribution_stock"] =  []mongo.IndexData{{Keys: []mongo.IndexKey{{Name: "coid", Key: "1"}}}}
	return req
}
