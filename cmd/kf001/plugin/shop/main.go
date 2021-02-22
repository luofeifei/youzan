package main

import (
	"base/pkg/app"
	"base/plugin/common"
	"fmt"
	"plugin/shop/service"
	"reflect"
	"strconv"
)

/*
编译
go build --buildmode=plugin  -o ../../../base/plugin/shop_1.0.1.so shop/main.go
go build --buildmode=plugin -ldflags "-w -s"  -o ../../../base/plugin/shop_1.0.1.so shop/main.go

*/

var (
	engine     common.IEngine
	pluginInfo common.PluginInfo
	//shopFuncs  = map[string]interface{}{
	//	"shopGroupList": ShopGroupList,
	//	"shopList":      ShopList,
	//}
	//pageInfo IndexPage
)

// 分页参数返回
//type IndexPage struct {
//	Total    int64  `json:"total"`    //总数
//	Page     int64  `json:"page"`     //页数
//	PageSize int64  `json:"pageSize"` //每页显示数
//	OrderKey string `json:"orderKey"` // 默认排序字段 -filed1,+field2,field3 (-Desc 降序)
//}

// 插件注册入口
func Initialize(e common.IEngine, info common.PluginInfo) error {
	engine = e
	pluginInfo = info
	return nil
}

// 调试 编译时删除
func init() {

	//db, _ := database.MysqlConnect("mall_shop", "192.168.2.22:3306", "user", "123456", "100", "100")
	//common.SetOrm("mall_shop","123456",db)
	//app.Println(common.OrmClient)
	//return
	//pageInfo = IndexPage{Total: 0,Page:1,PageSize:20}
	//apiPlugin:=common.PluginBaseInfo{
	//	Uid:123,
	//	Coid:123,
	//	FeatureId:123456,
	//	PluginInfo:common.ApiPluginInfo{
	//
	//		Alias:"shop",
	//		PluginId:1493510,
	//		ModuleId:469713,
	//		ModuleAlias:"SHOP",
	//		Value: map[string]interface{}{
	//			"funcName":"shopGroupList",
	//			"pageInfo":map[string]interface{}{
	//				"PageSize":10,
	//				"Page":1,
	//				"OrderKey":"",
	//			},
	//		},
	//
	//	},
	//}

	//apiPlugin := common.PluginBaseInfo{
	//	Uid:       1,
	//	Coid:      100000,
	//	FeatureId: 123456,
	//	PluginInfo: common.ApiPluginInfo{
	//		Alias:       "shop",
	//		PluginId:    1493510,
	//		ModuleAlias: "SHOP",
	//		Value: map[string]interface{}{
	//			"list":[]int64{7287369,7396617,7418427},
	//		},
	//	},
	//}
	//_, _ = PluginEntry(apiPlugin)

}

func PluginEntry(m common.PluginBaseInfo) (res interface{}, err error) {

	if m.PluginInfo.Alias == "shop" {
		if m.Coid<1{
			err=app.Err("请输入合法的企业coid")
			return
		}
		if len(m.PluginInfo.Value) > 0 {
			if listValue, ok := m.PluginInfo.Value["list"]; ok {
				if  reflect.TypeOf(listValue).Kind()== reflect.Slice{
					 var goodsIds []int64
					 app.Unmarshal(listValue,&goodsIds)
					 //app.Println("goodsIds-----------")
					 //app.Println(goodsIds)
					if len(goodsIds)>0{
						//检测ORM数据库是否连接
						//if common.HashOrm("shop_goods","123456"){

							frontGoodsList,errInner:=service.ShopList(engine,m.Coid,goodsIds)
							if len(frontGoodsList.List)>0{
								for i:=0;i<len(frontGoodsList.List);i++ {
									//frontGoodsList.List[i].Price
									frontGoodsList.List[i].Price,_ = strconv.ParseFloat(fmt.Sprintf("%.2f",  frontGoodsList.List[i].Price/100), 32)
								}
							}
							//app.Println("errInner----")
							//app.Println(errInner)
							//app.Println("frontGoodsList------------------")
							app.Println(frontGoodsList)
						    app.Println( app.Struct2Json(frontGoodsList))
					    	return app.Struct2Json(frontGoodsList), err
							if errInner!=nil{
								//app.Println("frontGoodsList------------------")
								//app.Println(frontGoodsList)

							}
							//app.Println("frontGoodsList------------------")
							//app.Println(app.Struct2Json(frontGoodsList))
						//}
					}
				}

			}
		}

	}
	return
}
// 数据处理入口
//func PluginEntry(m common.PluginBaseInfo) (res interface{}, err error) {
//	// 根据调用组件 区别读取数据
//	// m 插件配置值
//
//	if m.PluginInfo.Alias == "shop" { // 商城插件的别名
//		// 根据传过来的参数 决定去哪里拿数据 value里面传的值 list [21456]
//		if len(m.PluginInfo.Value) > 0 {
//
//			//if pager, ok := m.PluginInfo.Value["pageInfo"]; ok {
//			//	app.Unmarshal(pager, &pageInfo)
//			//	app.Println(pageInfo)
//			//}
//			if funcName, ok := m.PluginInfo.Value["list"]; ok {
//
//				if result, err := CallShopFuncs(shopFuncs, funcName.(string), m.Coid, pageInfo); err == nil {
//
//					if result[0].Kind() == reflect.Struct {
//						//返回结构体指针(进行强制转换)
//						rrr := result[0].Interface()
//						if reflect.TypeOf(rrr).Kind().String() == "struct" {
//							//app.Println(app.Struct2Json(rrr))
//							return app.Struct2Json(rrr), err
//						}
//
//					}
//
//				}
//			}
//
//		}
//
//	}
//
//	return
//}
//
//func CallShopFuncs(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
//	f := reflect.ValueOf(m[name])
//	if len(params) != f.Type().NumIn() {
//		err = app.Err("参数个数有误！")
//		return
//	}
//
//	in := make([]reflect.Value, len(params))
//	for k, param := range params {
//		in[k] = reflect.ValueOf(param)
//	}
//	result = f.Call(in)
//
//	return
//}

//func ShopGroupList(coid int64, pageInfo IndexPage) (list serverShop.ResGoodsGroupList, err error) {
//	// 统计分组有多少商品 第二种方法
//	//Table()
//	db, err := database.MysqlConnect("mall_shop", "192.168.2.22:3306", "user", "123456", "100", "100")
//	//关闭数据库连接
//	defer db.Close()
//	if err != nil {
//		return
//	}
//	if db.HasTable("shop_group") {
//		Db := db.Table("shop_group").
//			Select("id,coid,name,created_at")
//		if coid > 0 {
//			Db = Db.Where("coid = ?", coid)
//		}
//		//分页
//		Db = Db.Limit(pageInfo.PageSize).Offset((pageInfo.Page - 1) * pageInfo.PageSize).Order(pageInfo.OrderKey)
//
//		err = Db.Find(&list.List).Error
//
//		if len(list.List) > 0 {
//			for i, val := range list.List { // 查找对应的商品数量
//				//app.Println("查找对应的商品数量")
//				var out int32
//				db.Model(&mall_shop.ShopGoods{}).Where(mall_shop.ShopGoods{GroupID: val.Id}).Select("count(*)").Limit(1).Count(&out)
//				list.List[i].GoodsCount = out
//			}
//			list.Count = (int32)(len(list.List))
//			return list, err
//		}
//
//	}
//	return
//}
//func ShopList(coid int64, pageInfo IndexPage) (list serverShop.ResFrontGoodsList, err error) {
//	// 统计分组有多少商品 第二种方法
//	db, err := database.MysqlConnect("mall_shop", "192.168.2.22:3306", "user", "123456", "100", "100")
//	//关闭数据库连接
//	defer db.Close()
//	if err != nil {
//		return
//	}
//	if db.HasTable("shop_goods") {
//
//		Db := db.Table("shop_goods item").
//			Select(
//				"item.id,\n" +
//					"item.name,\n" +
//					"item.group_id,\n" +
//					"item.title,\n" +
//					"item.price,\n" +
//					"item.stock_num,\n" +
//					"item.cover,\n" +
//					"item.coid,\n" +
//					"item_data.video,\n" +
//					"item_data.share,\n" +
//					"item_data.sub_name,\n" +
//					"item_data.pic").
//			Joins("LEFT JOIN shop_goods_data item_data ON	item.id = item_data.id")
//		if coid > 0 {
//			Db = Db.Where("item.coid = ?", coid)
//		}
//		//分页
//		Db = Db.Limit(pageInfo.PageSize).Offset((pageInfo.Page - 1) * pageInfo.PageSize).Order(pageInfo.OrderKey)
//		err = Db.Find(&list.List).Error
//		list.Count = (int32)(len(list.List))
//
//		return list, err
//	}
//	return
//}
//
//func main(){
//
//}