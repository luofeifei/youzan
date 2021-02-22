package service

import (
	"base/plugin/common"
	"base/plugin/pkg/database"
)
type ResFrontSearchInfo struct {
	 Id int64  `json:"id"`
	 GroupId int64 `json:"group_id"`
	 Coid int64 `json:"coid"`
	 Name string `json:"name"`
	 Title string `json:"title"`
	 Price float64 `json:"price"`
	 Video string `json:"video"`
	 Pic string `json:"pic"`
	 StockNum int64 `json:"stock_num"`
	 Cover string `json:"cover"`
	 PriceDot float64 `json:"price_dot"`
}
type ResFrontGoodsList struct{
	List [] ResFrontSearchInfo  `json:"list"`
	Count int32 `json:"count"`
}

// 模块传批量的商品id过来 返回数据
func ShopList(engine common.IEngine,coid int64, goodsIds []int64) (list ResFrontGoodsList, err error) {
	db, err := database.MysqlConnect("mall_shop", "192.168.2.22:3306", "user", "123456", "100", "100")
	//关闭数据库连接
	defer db.Close()
	if err != nil {
		return
	}
	//app.Println("engine---------")
	//app.Println(engine)
	//db:=engine.Orm("mall_shop")
	//app.Println("db------------------------")
	//app.Println(db)
	if db.HasTable("shop_goods"){
		Db := db.Table("shop_goods item").
			Select(
				"item.id,\n" +
					"item.name,\n" +
					"item.group_id,\n" +
					"item.title,\n" +
					"item.price,\n" +
					"item.price_dot,\n" +
					"item.stock_num,\n" +
					"item.cover,\n" +
					"item.coid,\n" +
					"item_data.video,\n" +
					"item_data.share,\n" +
					"item_data.sub_name,\n" +
					"item_data.pic").
			Joins("LEFT JOIN shop_goods_data item_data ON	item.id = item_data.id")
		if coid > 0 {
			Db = Db.Where("item.coid = ?", coid)
		}
		if len(goodsIds)>0{
			Db = Db.Where("item.id  IN (?)", goodsIds)
		}
		//分页
		//Db = Db.Limit(pageInfo.PageSize).Offset((pageInfo.Page - 1) * pageInfo.PageSize).Order(pageInfo.OrderKey)

		err = Db.Find(&list.List).Error
		list.Count = (int32)(len(list.List))

		return list, err
	}
	return
}
