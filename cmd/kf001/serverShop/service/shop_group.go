package service

import (
	"base/model/imp/serverShop"
	"base/model/modelSql/mall_shop"
	"base/model/mongoSql"
	"base/pkg/app"
	"base/server/pkg/database/mongo"
	"base/server/pkg/database/orm"
	"base/server/pkg/pager"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

/*
商品操作
*/
//func CoShopGoodSave(elect string, sg mall_shop.ShopGoodsClear, sd mall_shop.ShopGoodsData, msgr mongoSql.CoShopGoods) (res mall_shop.ResShopGoods, err error) {
//	if sg.ID > 0 {
//		//修改
//		//开启事务修改
//		tx := orm.Master().Begin()
//
//		defer func() {
//			if r := recover(); r != nil {
//				app.Println("Rollback----测试---------------")
//				tx.Rollback()
//			}
//		}()
//
//		// 先查一遍
//		var sgOld mall_shop.ShopGoodsClear
//		notFound, _ := orm.FirstByID("", &sgOld, sg.ID)
//		if notFound == true {
//			err = app.Err("未找到商品记录")
//		} else {
//
//			err = tx.Model(&mall_shop.ShopGoodsClear{}).Omit("id,coid").Where("id=?", sg.ID).Updates(&sg).Error
//			if err != nil {
//				tx.Rollback()
//				return
//			} else {
//				app.Println("商品信息更新成功---------------------")
//			}
//			// 再查一遍商品详情
//			var sdOld mall_shop.ShopGoodsData
//			sd.ID = sg.ID
//			notFound, _ := orm.FirstByID("", &sdOld, sd.ID)
//			if notFound == true {
//				err = app.Err("未找到商品详情记录")
//			} else {
//
//				err = tx.Model(&mall_shop.ShopGoodsData{}).Omit("id").Where("id=?", sd.ID).Updates(&sd).Error
//				if err != nil {
//					tx.Rollback()
//					return
//				} else {
//					app.Println("商品详情更新成功---------------------")
//
//				}
//			}
//			// 再查一遍商品规则
//			msgr.ID = sg.ID
//			//app.Println("sleep----start-----------")
//			//time.Sleep(20 * time.Second)
//			//app.Println("sleep----end-------------mongo----")
//			if count, _ := mongo.Collection(&msgr).Where(bson.M{"_id": msgr.ID}).Count(); count > 0 {
//				_, errInner := mongo.Collection(&mongoSql.CoShopGoods{}).UpdateOne(&msgr)
//				if errInner == nil {
//					//panic(app.Err("123456"))
//					//app.Println("mongo商品规则更新成功------")
//					//app.Println(result) {"MatchedCount":1,"ModifiedCount":1,"UpsertedCount":0,"UpsertedID":null}
//				} else {
//					//app.Println("mongo创建商品规则失败---回滚--测试---------------")
//					// mongoDB更新失败 不能回滚这次的 删除本次的更新 再插入之前保存的记录
//					tx.Rollback()
//					return
//				}
//			} else {
//				//app.Println("mongo创建商品规则---不存在不存在！！！！！！！")
//			}
//
//			err = tx.Commit().Error
//			return res, err
//		}
//
//	} else {
//		//app.Println("CoShopGoodSave---------------CREATE-------------")
//		// 添加
//		// 事务处理 都成功才返回正确
//		tx := orm.Master().Begin()
//		defer func() {
//			if r := recover(); r != nil {
//				//app.Println("Rollback----测试---------------")
//				tx.Rollback()
//			}
//		}()
//
//		if err = tx.Error; err != nil {
//
//			return
//		}
//		// 第一步创建商品 获取商品ID
//		if err = tx.Create(&sg).Error; err != nil {
//			//app.Println("创建商品失败---回滚--测试----")
//			tx.Rollback()
//			return
//		}
//
//		if sg.ID > 0 { // 创建商品成功
//			//app.Println("SG-----创建商品成功-----------------------------")
//			//app.Println(sg)
//			//res.GoodsId = sg.ID // 商品id
//			sd.ID = sg.ID
//			if err = tx.Create(&sd).Error; err != nil {
//				//app.Println("创建商品详情失败---回滚---测试-----")
//				tx.Rollback()
//				return
//			} else {
//
//				//panic(app.Err("回滚测试---------------"))
//				//app.Println("sd-----创建商品详情成功-----------------------------")
//				//app.Println(sd)
//				res.DataId = sd.ID
//				msgr.ID = sg.ID
//				//app.Println("sleep----start-----------")
//				//time.Sleep(20 * time.Second)
//				//app.Println("sleep----end-------------mongo----")
//				if count, _ := mongo.Collection(&msgr).Where(bson.M{"_id": msgr.ID}).Count(); count == 0 {
//					result, errInner := mongo.Collection(&mongoSql.CoShopGoods{}).InsertOne(&msgr)
//					if errInner == nil {
//						//app.Println("sd-----mongo创建商品规则成功-----------------------------")
//						//app.Println(msgr)
//						res.GoodsId = result.InsertedID.(int64)
//						res.RuleId = res.GoodsId
//					} else {
//						//app.Println("mongo创建商品规则失败---回滚--测试---------------")
//						tx.Rollback()
//						return
//					}
//				} else {
//					//app.Println("mongo创建商品规则---存在存在！！！！！！！")
//				}
//			}
//		}
//
//		err = tx.Commit().Error
//		return res, err
//
//	}
//	return
//}


func ShopGoodSave(elect string, sg mall_shop.ShopGoodsClear, sd mall_shop.ShopGoodsData,stockDisCount mongoSql.GoodsStockDisCount,distributionStock mongoSql.GoodsDistributionStock) (res mall_shop.ResShopGoods, err error) {
	if sg.ID > 0 {
		//修改
		//开启事务修改
		tx := orm.Master().Begin()

		defer func() {
			if r := recover(); r != nil {
				//app.Println("Rollback----测试---------------")
				tx.Rollback()
			}
		}()

		// 先查一遍
		var sgOld mall_shop.ShopGoodsClear
		notFound, _ := orm.FirstByID("", &sgOld, sg.ID)

		if notFound == true {
			err = app.Err("未找到商品记录")
		} else if sgOld.Coid !=sg.Coid{

				err = app.Err("企业非法操作")
				return res,err

		}else {
			err = tx.Model(&mall_shop.ShopGoodsClear{}).Omit("id,coid").Where("id=?", sg.ID).Updates(&sg).Error
			if err != nil {
				tx.Rollback()
				return
			} else {
				res.GoodsId=sg.ID
				//app.Println("商品信息更新成功---------------------")
			}
			// 再查一遍商品详情
			var sdOld mall_shop.ShopGoodsData
			sd.ID = sg.ID
			notFound, _ := orm.FirstByID("", &sdOld, sd.ID)
			if notFound == true {
				err = app.Err("未找到商品详情记录")
			} else {

				err = tx.Model(&mall_shop.ShopGoodsData{}).Omit("id").Where("id=?", sd.ID).Updates(&sd).Error
				if err != nil {
					tx.Rollback()
					return
				} else {
					//app.Println("商品详情更新成功---------------------")
					res.DataId=sd.ID

				}
			}
			// 再查一遍商品规则
			stockDisCount.ID = sg.ID
			//app.Println("sleep----start-----------")
			//time.Sleep(20 * time.Second)
			//app.Println("sleep----end-------------mongo----")
			if count, _ := mongo.Collection(&stockDisCount).Where(bson.M{"_id": stockDisCount.ID}).Count(); count > 0 {
				_, errInner := mongo.Collection(&stockDisCount).Where(bson.M{"_id": stockDisCount.ID}).UpdateOne(&stockDisCount)
				if errInner == nil {
					//panic(app.Err("123456"))
					//app.Println("mongo商品规则更新成功------")
					//app.Println(result) {"MatchedCount":1,"ModifiedCount":1,"UpsertedCount":0,"UpsertedID":null}
					// 更新库存 查一遍判断库存是否有变化 库存有变化才更新
					res.RuleId=stockDisCount.ID
					var oldDistributionStock mongoSql.GoodsDistributionStock

					mongo.Collection(&distributionStock).Where(bson.M{"_id": distributionStock.ID}).FindOne(&oldDistributionStock)
					//app.Println(err)
					//app.Println("old_distributionStock--StockNum---")
					//app.Println(oldDistributionStock.StockNum)
					//app.Println(distributionStock.StockNum)

					if oldDistributionStock.ID>0{
						//app.Println("old_distributionStock--StockNum----1---------------------")

						if oldDistributionStock.StockNum!=distributionStock.StockNum{
							resUp, _ := mongo.Collection(&distributionStock).UpdateOne(&distributionStock)
							app.Println(resUp)
							if resUp.ModifiedCount>0{
								app.Println("mongo商品库存更新成功------")
							}else{
								tx.Rollback()
								return
							}
						}

					}else{

						//app.Println("old_distributionStock--StockNum----2---------------------")

						tx.Rollback()
						return
					}

				} else {
					//app.Println("mongo创建商品规则失败---回滚--测试---------------")
					// mongoDB更新失败 不能回滚这次的 删除本次的更新 再插入之前保存的记录
					tx.Rollback()
					return
				}
			} else {
				//app.Println("mongo创建商品规则---不存在不存在！！！！！！！")
			}

			err = tx.Commit().Error

			app.Println(res)
			return res, err
		}





	} else {
		//app.Println("shopGoodSave---------------CREATE-------------")
		// 添加
		// 事务处理 都成功才返回正确
		tx := orm.Master().Begin()
		defer func() {
			if r := recover(); r != nil {
				//app.Println("Rollback----测试---------------")
				tx.Rollback()
			}
		}()

		if err = tx.Error; err != nil {

			return
		}
		// 第一步创建商品 获取商品ID
		if err = tx.Create(&sg).Error; err != nil {
			//app.Println("创建商品失败---回滚--测试----")
			tx.Rollback()
			return
		}

		if sg.ID > 0 { // 创建商品成功
			//app.Println("SG-----创建商品成功-----------------------------")
			//app.Println(sg)
			//res.GoodsId = sg.ID // 商品id
			sd.ID = sg.ID
			if err = tx.Create(&sd).Error; err != nil {
				//app.Println("创建商品详情失败---回滚---测试-----")
				tx.Rollback()
				return
			} else {

				//panic(app.Err("回滚测试---------------"))
				//app.Println("sd-----创建商品详情成功-----------------------------")
				//app.Println(sd)
				res.DataId = sd.ID
				stockDisCount.ID = sg.ID
				//stockDisCount.GoodsID = sg.ID
				//var stockDisCount mongoSql.GoodsStockDisCount
				//if(len(stockDisCount.GoodsRuleList))==1{ // 只有主商品 sku_id = 商品id
				//	stockDisCount.GoodsRuleList[0].ID = sg.ID
				//	stockDisCount.GoodsRuleList[0].SkuId = sg.ID
				//}

				//app.Println("sleep----start-----------")
				//time.Sleep(20 * time.Second)
				//app.Println("sleep----end-------------mongo----")


				if count, _ := mongo.Collection(&stockDisCount).Where(bson.M{"_id": stockDisCount.ID}).Count(); count == 0 {
					result, errInner := mongo.Collection(&mongoSql.GoodsStockDisCount{}).InsertOne(&stockDisCount)
					if errInner == nil {
						//app.Println("sd-----mongo创建商品规则成功-----------------------------")
						//app.Println(msgr)
						res.GoodsId = result.InsertedID.(int64)
						res.RuleId = res.GoodsId
						// 再创建商品库存文档 成功后缓存库存文档
						distributionStock.ID = res.GoodsId
						_, dsErr := mongo.Collection(&distributionStock).InsertOne(&distributionStock)
						if dsErr!=nil{
							tx.Rollback()
							return
						}
					} else {
						//app.Println("mongo创建商品规则失败---回滚--测试---------------")
						tx.Rollback()
						return
					}
				} else {
					//app.Println("mongo创建商品规则---存在存在！！！！！！！")
				}
			}
		}

		err = tx.Commit().Error
		return res, err

	}
	return
}
func GoodsSettingSku( Coid int64 ,goodsId int64,skuId int64,gsdc mongoSql.GoodsRule)(count int64, err error) {

	filter := bson.M{"_id": goodsId,"coid":Coid, "goods_rule_list.sku_id": skuId}
	update := bson.M{"goods_rule_list.$.goods_discount_list": gsdc.GoodsDiscountList}
	//arrayFilter := bson.M{"item.sku_id": skuId}
	res,errInner:=mongo.Collection(&mongoSql.GoodsStockDisCount{}).Where(filter).UpdateMany(update)
	app.Println(res)
	return res.ModifiedCount ,errInner
	//mongo.Collection(&mongoSql.GoodsStockDisCount{}).Where(filter).UpdateOne(&gsdc)


	//res :=  mongo.Collection(&mongoSql.GoodsStockDisCount{}).Table.FindOneAndUpdate(context.Background(),
	//	filter,
	//	bson.M{"$set": update},
	//	options.FindOneAndUpdate().SetArrayFilters(
	//		options.ArrayFilters{
	//			Filters: []interface{}{
	//				arrayFilter,
	//			},
	//		},
	//	))
	//if res.Err() != nil {
	//
	//}
	//err = mongo.Collection(&mongoSql.GoodsStockDisCount{}).Where(bson.M{"_id": goodsId,}).FindOne(&gsdcddd)
    // app.Println("gsdcddd---")
	//app.Println(gsdcddd)
	//res, errInner:= mongo.Collection(&mongoSql.GoodsStockDisCount{}).Where(bson.M{"_id": goodsId,"coid":Coid,"goods_rule_list.sku_id":skuId}).UpdateOne(gsdc.GoodsDiscountList)
		//Update(gsdc.GoodsDiscountList)
	// 修改主商品会员价格成功后 再修改sku商品的价格 蓝色-50码 商品的会员价格
	//app.Println(res)
	//res :=  mongo.Collection().FindOneAndUpdate(context.Background(),
	//	filter,
	//	bson.M{"$set": update},
	//	options.FindOneAndUpdate().SetArrayFilters(
	//		options.ArrayFilters{
	//			Filters: []interface{}{
	//				arrayFilter,
	//			},
	//		},
	//	))


}
//func GoodsStockRule(elect string, goodsId int64) (msgr mongoSql.CoShopGoods, err error) {
//	err = mongo.Collection(&mongoSql.CoShopGoods{}).Where(bson.M{"_id": goodsId}).FindOne(&msgr)
//	return
//}
//func GoodsDetail(elect string, goodsId int64) (sg mall_shop.ShopGoodsClear, sd mall_shop.ShopGoodsDataNew, msgr mongoSql.CoShopGoods, err error) {
//
//	notFound, _ := orm.FirstByID("", &sg, goodsId)
//	if notFound == true {
//		err = app.Err("未找到商品记录")
//		return
//	}
//	//一对一 关联查询 或者
//	//orm.Slave().Model(&sg).Related(&sd, "id").Find(&sg)
//	//
//	//app.Println("inner--join---Related--sg-")
//	//app.Println(sg)
//	//app.Println("inner--join---Related---sd")
//	//app.Println(sd)
//	//var sdd mall_shop.ShopGoodsDataNew
//	//err = orm.Find("", &mall_shop.ShopGoodsDataNew{ID: goodsId}, &sdd)
//	//app.Println("inner--join---Related---sdd")
//	//app.Println(sdd)
//	//
//	//notFound, _ = orm.FirstByID("", &sd, goodsId)
//
//	notFound, _ = orm.FirstByID("", &sd, goodsId)
//	if notFound == true {
//		err = app.Err("未找到商品详情记录")
//		return
//	}
//	//app.Println("inner--join---Related---sd")
//	//app.Println(sd)
//	// 商品如果有下级的规则 没有就为空
//	err = mongo.Collection(&mongoSql.CoShopGoods{}).Where(bson.M{"_id": goodsId}).FindOne(&msgr)
//	//app.Println("msgr---------------------------------------------------------------------")
//	//app.Println(msgr)
//
//	if err != nil {
//		// err = app.Err("未找到商品规则记录")
//		return
//	}
//	return sg, sd, msgr, err
//}

func GoodsDistributionDetail(elect string, goodsId int64) (gds mongoSql.GoodsDistributionStock,err error) {
	err=mongo.Collection(&mongoSql.GoodsDistributionStock{}).Where(bson.M{"_id": goodsId}).FindOne(&gds)
    return
}

func GoodsDetail(elect string, goodsId int64,coid int64) (sg mall_shop.ShopGoodsClear, sd mall_shop.ShopGoodsData, msgr mongoSql.GoodsStockDisCount, err error) {

	notFound, _ := orm.FirstByID("", &sg, goodsId)
	if notFound == true {
		err = app.Err("未找到商品记录")
		return
	}
	//一对一 关联查询 或者
	//orm.Slave().Model(&sg).Related(&sd, "id").Find(&sg)
	//
	//app.Println("inner--join---Related--sg-")
	//app.Println(sg)
	//app.Println("inner--join---Related---sd")
	//app.Println(sd)
	//var sdd mall_shop.ShopGoodsDataNew
	//err = orm.Find("", &mall_shop.ShopGoodsDataNew{ID: goodsId}, &sdd)

	//
	//notFound, _ = orm.FirstByID("", &sd, goodsId)

	notFound, _ = orm.FirstByID("", &sd, goodsId)
	if notFound == true {
		err = app.Err("未找到商品详情记录")
		return
	}

	// 商品如果有下级的规则 没有就为空
	err = mongo.Collection(&mongoSql.GoodsStockDisCount{}).Where(bson.M{"_id": goodsId}).FindOne(&msgr)

	//app.Println("inner--join---Related---SG")
	//app.Println(sg)

	//app.Println("inner--join---Related---sd")
	//app.Println(sd)

	//app.Println("msgr---------------------------------------------------------------------")
	//app.Println(msgr)

	if err != nil {
		// err = app.Err("未找到商品规则记录")
		return
	}
	return sg, sd, msgr, err
}

func GoodsFrontDetail(elect string, goodsId int64) (sg mall_shop.ShopGoodsClear, sd mall_shop.ShopGoodsData, msgr mongoSql.GoodsStockDisCount, err error) {

	notFound, _ := orm.FirstByID("", &sg, goodsId)
	if notFound == true {
		err = app.Err("未找到商品记录")
		return
	}
	//一对一 关联查询 或者
	//orm.Slave().Model(&sg).Related(&sd, "id").Find(&sg)
	//
	//app.Println("inner--join---Related--sg-")
	//app.Println(sg)
	//app.Println("inner--join---Related---sd")
	//app.Println(sd)
	//var sdd mall_shop.ShopGoodsDataNew
	//err = orm.Find("", &mall_shop.ShopGoodsDataNew{ID: goodsId}, &sdd)

	//
	//notFound, _ = orm.FirstByID("", &sd, goodsId)

	notFound, _ = orm.FirstByID("", &sd, goodsId)
	if notFound == true {
		err = app.Err("未找到商品详情记录")
		return
	}


	// 商品如果有下级的规则 没有就为空
	err = mongo.Collection(&mongoSql.GoodsStockDisCount{}).Where(bson.M{"_id": goodsId}).FindOne(&msgr)

	//app.Println("inner--join---Related---SG")
	//app.Println(sg)

	//app.Println("inner--join---Related---sd")
	//app.Println(sd)

	//app.Println("msgr---------------------------------------------------------------------")
	//app.Println(msgr)

	if err != nil {
		// err = app.Err("未找到商品规则记录")
		return
	}
	return sg, sd, msgr, err
}

func GoodsDelete(m mall_shop.ShopGoodsClear) (count int64, err error) {
	// 不需要事务 直接硬删除 1--> 删除对应的规则 2--> 删除商品详情 3--> 删除商品
	_, err = mongo.Collection(&mongoSql.GoodsStockDisCount{}).Where(bson.M{"_id": m.ID}).Delete()
	if err != nil {
		return
	}
	_, err = orm.DeleteByID(mall_shop.ShopGoodsDataNew{}, m.ID)
	if err != nil {
		return
	}
	count, err = orm.DeleteByID(mall_shop.ShopGoodsClear{}, m.ID)
	if err != nil {
		return
	}
	return count, err
}
func GoodsBatchDelete(ids []int64) (count int64, err error) {
	// 不需要事务 直接硬删除 1--> 删除对应的规则 2--> 删除商品详情 3--> 删除商品
	// bson.M{"foo": "bar", "hello": "world", "pi": 3.14159}.

	//app.Println("GoodsBatchDelete------")
	if ids != nil && len(ids) > 0 {

		count, err = orm.DeleteByIDS(&mall_shop.ShopGoodsClear{}, ids)

		if err != nil || count == 0 {
			return count, err
		} else {
			//app.Println("GoodsBatchDelete---删除商品成功---")
			count, err = orm.DeleteByIDS(&mall_shop.ShopGoodsDataNew{}, ids)
			if err != nil || count == 0 {
				return count, err
			} else {
				//app.Println("GoodsBatchDelete---删除商品详情成功---")

				for _, v := range ids {
					_, err = mongo.Collection(&mongoSql.GoodsStockDisCount{}).Where(bson.M{"_id": v}).Delete()
					if err != nil {
						return count, err
					}

				}

			}
		}

		//return
		//for _,v:=range ids{
		//	_,err=mongo.Collection(&mongoSql.CoShopGoods{}).Where(bson.M{"_id": v}).Delete()
		//	if err!=nil{
		//		return
		//	}
		//	_, err = orm.DeleteByID(&mall_shop.ShopGoodsDataNew{},v)
		//	if err!=nil{
		//		return
		//	}
		//	_, err = orm.DeleteByID(&mall_shop.ShopGoodsClear{},v)
		//	if err!=nil{
		//		return
		//	}
		//}
	}
	return

}
func GoodsBatchState(ids []int64, state int32,timestamp int64) (count int64, err error) {
	// 修改商品状态state
	if len(ids) > 0 {
		if state==2{ // 上架
			count = orm.Master().Debug().Model(&mall_shop.ShopGoodsClear{}).Where("id in (?)", ids).Updates(map[string]interface{}{"state": state, "start_time": timestamp}).RowsAffected
		}else if state==3{ // 下架
			count = orm.Master().Model(&mall_shop.ShopGoodsClear{}).Where("id in (?)", ids).Updates(map[string]interface{}{"state": state, "end_time": timestamp}).RowsAffected
		}else { // 停用
			count = orm.Master().Model(&mall_shop.ShopGoodsClear{}).Where("id in (?)", ids).Update(map[string]interface{}{"state": state, "end_time": timestamp}).RowsAffected
		}
	}

	return
}
func GoodsChageGroup(ids []int64, groupId int64) (count int64, err error) {
	// 修改商品分组
	if ids != nil && len(ids) > 0 {
		count = orm.Master().Model(&mall_shop.ShopGoodsClear{}).Select("group_id").Where("id in (?)", ids).Update("group_id", groupId).RowsAffected
	}
	return
}
//func GoodsBatchSettingDiscount(msgr []mongoSql.GoodsStockDisCount) (count int64, err error) {
//
//	count = 0
//	for _, v := range msgr {
//		count = count + 1
//		_, err= mongo.Collection(&mongoSql.GoodsStockDisCount{}).Where(bson.M{"_id": v.ID}).UpdateOne(v)
//		// 修改主商品会员价格成功后 再修改sku商品的价格 蓝色-50码 商品的会员价格
//
//
//	}
//
//	return count,err
//}
//func GoodsSettingDiscount(msgr mongoSql.CoShopGoods) (count int64, err error) {
//
//
//	res, errInner:= mongo.Collection(&mongoSql.CoShopGoods{}).Where(bson.M{"_id": msgr.ID}).UpdateOne(msgr)
//		// 修改主商品会员价格成功后 再修改sku商品的价格 蓝色-50码 商品的会员价格
//
//	return res.ModifiedCount ,errInner
//}
func GoodsBatchSetting(vo mall_shop.CoShopGoodsBatchSettingVO) (count int64) {
	// 7商品名称文字替换  8上架时间 9每人限购 10会员折扣 11限定用户(限定用户有购买的权限) 12 配送方式 13商品模板

	switch vo.Type {

	case 7: //
		//app.Println("name--------------")
		//app.Println(vo)
		if len(vo.GoodsList) > 0 && len(vo.FeatureList.GoodsName.NameList) > 0 && len(vo.FeatureList.GoodsName.OldName) > 0 {
			count = 0
			for i, val := range vo.FeatureList.GoodsName.NameList {
				newName := strings.Replace(val, vo.FeatureList.GoodsName.OldName, vo.FeatureList.GoodsName.NewName, -1)
				count = count + orm.Master().Model(&mall_shop.ShopGoodsClear{}).Select("name").Where("id = ?", vo.GoodsList[i]).Update("name", newName).RowsAffected

			}
		}

	case 8:
		if len(vo.GoodsList) > 0 && vo.FeatureList.StartTime > 0 {
			count = orm.Master().Model(&mall_shop.ShopGoodsClear{}).Select("start_time").Where("id in (?)", vo.GoodsList).Update("start_time", vo.FeatureList.StartTime).RowsAffected
		}
	case 9:
		if len(vo.GoodsList) > 0 && vo.FeatureList.QuotaNum > 0 && vo.FeatureList.Quota > 0 {
			quota := map[string]int32{"quota_num": vo.FeatureList.QuotaNum, "quota": vo.FeatureList.Quota}
			//app.Println(quota)
			//test :=map[string]interface{}{"quota": 1, "quota_num": 10}
			//app.Println(test)
			//ttttt:=&mall_shop.ShopGoodsClear{Quota:int8(vo.FeatureList.Quota),QuotaNum: int(vo.FeatureList.QuotaNum) }
			//app.Println("count--------------------------")
			//			//Select("quota,quota_num")
			//app.Println(count)
			count = orm.Master().Model(&mall_shop.ShopGoodsClear{}).Where("id in (?)", vo.GoodsList).Updates(quota).RowsAffected
		}
	case 10:
		if len(vo.GoodsList) > 0 && vo.FeatureList.IsVip > 0 {
			count = orm.Master().Model(&mall_shop.ShopGoodsClear{}).Select("is_vip").Where("id in (?)", vo.GoodsList).Update("is_vip", vo.FeatureList.IsVip).RowsAffected
		}
	case 11, 12, 13:
		//if len(vo.GoodsList) > 0 && len(vo.FeatureList.Member[0].GoodsCount) > 0 {
		//	app.Println("会员价设置--------------------------------")
		//	//app.Println( vo.FeatureList.Member)
		//}

	default:
		app.Println("GoodsBatchSetting --------default---------------------")
		return
	}

	return
}
func GoodsListFrontSearchPage(elect string, m mall_shop.SearchGoods, pageInfo orm.IndexPage) (list serverShop.ResFrontGoodsList, err error) {

	//var where = make(pager.Where)

	Db := orm.Slave().Table("shop_goods item").
		Select(
			"item.id,\n" +
				"item.name,\n" +
				"item.group_id,\n" +
				"item.title,\n" +
				"item.price,\n" +
				"item.coid,\n" +
				"item.price_dot,\n" +
				"item_data.video,\n" +
				"item_data.share,\n" +
				"item_data.sub_name,\n" +
				"item_data.pic").
		Joins("LEFT JOIN shop_goods_data item_data ON	item.id = item_data.id")
	if m.Coid > 0 {
		Db = Db.Where("coid = ?", m.Coid)
	}
	if m.GroupID > 0 {
		Db = Db.Where("group_id = ?", m.GroupID)
	}
	if m.Type > 0 {
		Db = Db.Where("type = ?", m.Type)
	}
	if m.StockType > 0 {
		Db = Db.Where("stock_type = ?", m.StockType)
	}
	if m.State > 0 {
		Db = Db.Where("state = ?", m.State)
	}
	if m.PriceStart >= 0 && m.PriceEnd > 0 {
		Db = Db.Where("price BETWEEN ? AND ?", m.PriceStart, m.PriceEnd)
	}
	if m.StartTime > 0 && m.EndTime > 0 {
		Db = Db.Where("start_time BETWEEN ? AND ?", m.StartTime, m.EndTime)
	}
	if len(m.Name) > 0 {
		Db = Db.Where("name LIKE (?)", "%"+m.Name+"%")
	}
	//分页
	Db = Db.Limit(pageInfo.PageSize).Offset((pageInfo.Page - 1) * pageInfo.PageSize).Order(pageInfo.OrderKey)

	err = Db.Find(&list.List).Error
	// mongo里面查询商品的库存和销量
	if len(list.List)>0{
		for index, v := range list.List {
			var gsd mongoSql.GoodsDistributionStock
			err = mongo.Collection(&mongoSql.GoodsDistributionStock{}).Fields("stock_num,sold_num").Where(bson.M{"_id": v.Id}).FindOne(&gsd)
			app.Println("err----")
			app.Println(err)
			if err == nil {
				list.List[index].StockNum =gsd.StockNum
				list.List[index].SoldNum =gsd.SoldNum
			}

		}
	}
	return list, err

}

func GoodsListPage(elect string, m mall_shop.ShopGoods, pageInfo orm.IndexPage) (list []mall_shop.ShopGoods, total int64, err error) {

	var where = make(pager.Where)
	if m.Coid > 0 {
		where["coid"] = m.Coid
	}
	total, err = pager.New(pager.NewGormDriver(), pageInfo).Fields(elect).SetIndex(m.TableName()).Where(where).Find(&list)
	for i, val := range list {
		_ = orm.Find("*", &mall_shop.ShopGoodsRule{GoodsID: val.ID}, &list[i].Rule)
		_, _ = orm.FirstByID("", &list[i].Data, val.ID)
		_, _ = orm.FirstByID("", &list[i].Group, val.GroupID)
		_, _ = orm.FirstByID("", &list[i].Stock, val.StockID)
	}
	return list, total, err

}

func GoodsListSearchPage(elect string, m mall_shop.SearchGoods, pageInfo orm.IndexPage) (list serverShop.ResFrontGoodsList, err error) {

	//where := make(pager.Where)
	//_, err = pager.New(pager.NewGormDriver(), pageInfo).Fields(elect).SetIndex(m.TableName()).Where(where).Find(&list.List)
	//
	////for i, val := range list.List {
	////	_ = orm.Find("*", &mall_shop.ShopGoodsRule{GoodsID: val.ID}, &list[i].Rule)
	////	_, _ = orm.FirstByID("", &list[i].Data, val.ID)
	////	_, _ = orm.FirstByID("", &list[i].Group, val.GroupID)
	////	_, _ = orm.FirstByID("", &list[i].Stock, val.StockID)
	////}
	//return list, err
	//return

	Db := orm.Slave().Table("shop_goods item").
		Select(
			"item.id,\n" +
				"item.name,\n" +
				"item.group_id,\n" +
				"item.title,\n" +
				"item.price,\n" +
				"item.stock_num,\n" +
				"item.price_dot,\n" +
				"item.cover,\n" +
				"item.coid,\n" +
				"item_data.video,\n" +
				"item_data.share,\n" +
				"item_data.sub_name,\n" +
				"item_data.pic").
		Joins("LEFT JOIN shop_goods_data item_data ON	item.id = item_data.id")
	if m.Coid > 0 {
		Db = Db.Where("item.coid = ?", m.Coid)
	}
	if m.GroupID > 0 {
		Db = Db.Where("item.group_id = ?", m.GroupID)
	}
	if m.Type > 0 {
		Db = Db.Where("item.type = ?", m.Type)
	}
	if m.StockType > 0 {
		Db = Db.Where("item.stock_type = ?", m.StockType)
	}
	if m.State > 0 {
		Db = Db.Where("item.state = ?", m.State)
	}
	if m.PriceStart >= 0 && m.PriceEnd > 0 {
		Db = Db.Where("item.price BETWEEN ? AND ?", m.PriceStart, m.PriceEnd)
	}
	if m.StartTime > 0 && m.EndTime > 0 {
		Db = Db.Where("item.start_time BETWEEN ? AND ?", m.StartTime, m.EndTime)
	}
	if len(m.Name) > 0 {
		Db = Db.Where("item.name LIKE (?)", "%"+m.Name+"%")
	}
	//分页
	Db = Db.Limit(pageInfo.PageSize).Offset((pageInfo.Page - 1) * pageInfo.PageSize).Order(pageInfo.OrderKey)

	err = Db.Find(&list.List).Error
	if len(list.List)>0{
		for index, v := range list.List {
			var gsd mongoSql.GoodsDistributionStock
			err = mongo.Collection(&mongoSql.GoodsDistributionStock{}).Fields("stock_num,sold_num").Where(bson.M{"_id": v.Id}).FindOne(&gsd)
			if err == nil {
				list.List[index].StockNum =gsd.StockNum
				list.List[index].SoldNum =gsd.SoldNum
			}
		}
	}
	return list, err

}

// 添加/修改商品分组
func SaveGroupData(elect string, m mall_shop.ShopGroup) (res mall_shop.ShopGroup, err error) {

	if m.ID > 0 { // 修改分组
		//notFound, err := orm.FirstByID("", &res, m.ID)
		notFound, err := orm.FirstByID("", &res,m.ID)
		//app.Println("SaveGroupData----------")
		//app.Println(notFound)
		//app.Println(m)
		//app.Println(res)
		if notFound == true {
			err = app.Err("未找到记录")
			return res,err
		} else if res.Coid !=m.Coid{
			err = app.Err("企业非法操作")
			return res,err

		} else {
			err = orm.Updates(&res, m)
			if err != nil {
				err = app.Err("更新商品分组失败")
			}
			return res, err
		}
	} else { // 创建分组

		// 是否创建过

		err := orm.Slave().Where("name = ?", m.Name).Find(&res).Error

		if res.ID>0 {
			err = app.Err("商品分组已经存在")
			return res,err
		}else{
			err = orm.Create(&m)
			if err != nil {
				err = app.Err("新增商品分组失败")
			}
			return m, err

		}


	}
	return

}

func GroupListByCoidPage(elect string, m mall_shop.ShopGroupVO, pageInfo orm.IndexPage) (list []mall_shop.ShopGroupVO, err error) {
	// 统计分组有多少商品 第一种方法 纯sql
	//Db := orm.Slave().Table("shop_group item").
	//	Select(
	//		"item.coid,\n"+
	//			"item.name,\n"+
	//			"item.id,\n"+
	//			"count(*) as goods_count").
	//	Joins("LEFT JOIN shop_goods item_data ON	item_data.group_id = item.id ").Group("item.id")
	//if m.Coid > 0 {
	//	Db = Db.Where("item.coid = ?", m.Coid)
	//}
	//
	//if len(m.Name) > 0 {
	//	Db = Db.Where("item.name LIKE (?)", "%"+m.Name+"%")
	//}
	////分页
	//Db = Db.Limit(pageInfo.PageSize).Offset((pageInfo.Page - 1) * pageInfo.PageSize).Order(pageInfo.OrderKey)
	//
	//err = Db.Find(&list).Error

	// 统计分组有多少商品 第二种方法
	var where = make(pager.Where)
	if m.Coid > 0 {
		where["coid"] = m.Coid
	}
	if len(m.Name) > 0 {
		where["name"] = m.Name
	}
	_, err = pager.New(pager.NewGormDriver(), pageInfo).Fields(elect).SetIndex(m.TableName()).Where(where).Find(&list)
	if len(list) > 0 {
		for i, val := range list { // 查找对应的商品数量
            //app.Println("查找对应的商品数量")
			//app.Println(val)


			var out int32
			orm.Slave().Model(&mall_shop.ShopGoods{}).Where(mall_shop.ShopGoods{GroupID: val.ID}).Select("count(*)").Limit(1).Count(&out)

			list[i].GoodsCount = out
		}
	}
	return list, err

}

func GroupListByCoid(elect string, m mall_shop.ShopGroup) (infos []mall_shop.ShopGroup, err error) {

	err = orm.Find(elect, m, &infos)

	return
}

func GroupDelete(m mall_shop.ShopGroup) (count int64, err error) {

	count, err = orm.DeleteByModel(m)

	return
}

// 添加/修改库存
func GoodsStockSave(elect string, m mall_shop.ShopStock) (res mall_shop.ShopStock, err error) {
	if m.ID > 0 {
		notFound, err := orm.FirstByID("", &res, m.ID)

		if notFound == true {
			err = app.Err("未找到记录")

		} else {
			err = orm.Updates(&res, m)
			if err != nil {
				err = app.Err("更新库存失败")
			}

			return res, err
		}
	} else {
		err = orm.Create(&m)
		if err != nil {
			err = app.Err("新增库存失败")
		}
		return m, err
	}
	return res, err
}
func StockDelete(m mall_shop.ShopStock) (count int64, err error) {

	count, err = orm.DeleteByModel(m)

	return
}

// 添加/修改库存卡
func StockCardSave(elect string, m mall_shop.ShopStockCard) (res mall_shop.ShopStockCard, err error) {
	if m.ID > 0 {
		notFound, err := orm.FirstByID("", &res, m.ID)

		if notFound == true {
			err = app.Err("未找到记录")

		} else {
			err = orm.Updates(&res, m)
			if err != nil {
				err = app.Err("更新库存失败")
			}

			return res, err
		}
	} else {
		err = orm.Create(&m)
		if err != nil {
			err = app.Err("新增库存失败")
		}
		return m, err
	}
	return res, err
}
func StockCardDelete(m mall_shop.ShopStockCard) (count int64, err error) {

	count, err = orm.DeleteByModel(m)

	return
}

func CouponListPage(elect string, m mall_shop.ShopCouponRelatedVO, pageInfo orm.IndexPage) (list []mall_shop.ShopCouponRelatedVO, total int64, err error) {

	var tb mall_shop.ShopCoupon
	tx := orm.Slave().Model(tb.TableName()).Preload("ShopCouponData")
	if m.Coid > 0 {
		tx = tx.Where("coid = ?", m.Coid)
	}
	if m.Name != "" {
		tx = tx.Where("name like ?", "%"+m.Name+"%")
	}
	if pageInfo.PageSize < 1 {
		pageInfo.PageSize = 10
	}
	if pageInfo.Page < 1 {
		pageInfo.Page = 1
	}

	var totalCount uint
	errCount := tx.Count(&totalCount).Error
	if err != nil {
		return list, 0, errCount
	}
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	err = tx.Limit(pageInfo.PageSize).Offset(offset).Find(&list).Error
	if err != nil {
		return list, 0, err
	}
	return list, total, err

}

func CoCouponSave(elect string, m mall_shop.ShopCoupon) (res mall_shop.ShopCoupon, err error) {

	if m.ID > 0 {
		//修改
		//开启事务修改
		tx := orm.Master().Begin()

		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		var sgOld mall_shop.ShopCoupon
		notFound, _ := orm.FirstByID("", &sgOld, m.ID)
		if notFound == true {
			err = app.Err("未找到优惠券记录")
		} else {

			err = tx.Model(&mall_shop.ShopCoupon{}).Omit("id,coid").Where("id=?", m.ID).Updates(&m).Error
			if err != nil {
				tx.Rollback()
				return
			} else {
				//app.Println("优惠券信息更新成功---------更新优惠券详情----有修改则更新-------")
				var scOld mall_shop.ShopCouponData
				notFound, _ := orm.FirstByID("", &scOld, m.ID)
				if notFound == true {
					err = app.Err("未找到优惠券详情记录")
				} else {

					if scOld != m.CouponData {
						err = tx.Model(&mall_shop.ShopCouponData{}).Omit("id").Where("id=?", m.ID).Updates(&m.CouponData).Error
						if err != nil {
							tx.Rollback()
							return
						}
					}

				}

			}

		}
		err = tx.Commit().Error
		return res, err

	} else {

		// 事务处理 都成功才返回正确
		tx := orm.Master().Begin()
		defer func() {
			if r := recover(); r != nil {
				//app.Println("Rollback----测试---------------")
				tx.Rollback()
			}
		}()
		if err = tx.Error; err != nil {
			return
		}
		// 第一步创建优惠券
		if err = tx.Create(&m).Error; err != nil {
			//app.Println("创建优惠券失败---回滚--测试----")
			tx.Rollback()
			return
		}

		if m.ID > 0 { // 创建优惠券成功
			//第二步 创建优惠券详情
			m.CouponData.ID = m.ID
			if err = tx.Create(&m.CouponData).Error; err != nil {
				//app.Println("创建优惠券详情失败---回滚---测试-----")
				tx.Rollback()
				return
			}

		}

		err = tx.Commit().Error
		return res, err

	}
	return
}
func CouponDelete(m mall_shop.ShopCoupon) (count int64, err error) {

	count, err = orm.DeleteByModel(m)
	return
}

// 添加/修改优惠券兑换卡
func CouponCardSave(elect string, m mall_shop.ShopCouponCard) (res mall_shop.ShopCouponCard, err error) {

	if m.ID > 0 {
		notFound, _ := orm.FirstByID("id", &res, m.ID)
		if notFound == true {
			//err = orm.Create(&m)
			err = app.Err("新增库存失败")

		} else {
			err = orm.Updates(&res, m)
			if err != nil {
				err = app.Err("更新库存失败")
			}
			return res, err
		}
	} else {
		err = orm.Create(&m)
		if err != nil {
			err = app.Err("新增库存失败")
		}
		return m, err
	}
	return
}

func CouponCardDelete(m mall_shop.ShopCouponCard) (count int64, err error) {

	count, err = orm.DeleteByModel(m)

	return
}

func ShopGroupDetail(elect string, m mall_shop.ShopGroup) (res mall_shop.ShopGroup, err error) {

	notFound, _ := orm.FirstByID("", &res, m.ID)

	if notFound == true {
		err = app.Err("未找到记录")
	}
	return res, err
}
func StockDetail(elect string, m mall_shop.ShopStock) (res mall_shop.ShopStock, err error) {

	notFound, _ := orm.FirstByID("", &res, m.ID)

	if notFound == true {
		err = app.Err("未找到记录")
	}
	return res, err
}
func StockCardDetail(elect string, m mall_shop.ShopStockCard) (res mall_shop.ShopStockCard, err error) {

	notFound, _ := orm.FirstByID("", &res, m.ID)

	if notFound == true {
		err = app.Err("未找到记录")
	}
	return res, err
}

func CouponDetail(elect string, m mall_shop.ShopCoupon) (res mall_shop.ReqShopCoupon, err error) {

	notFound, _ := orm.FirstByID("", &res.SC, m.ID)

	if notFound == true {
		err = app.Err("未找到记录")
	} else {
		notFound, _ := orm.FirstByID("", &res.SD, m.ID)
		if notFound == true {
			err = app.Err("未找到记录")
		}

	}

	return res, err
}

func CouponCardDetail(elect string, m mall_shop.ShopCouponCard) (res mall_shop.ShopCouponCard, err error) {

	notFound, _ := orm.FirstByID("", &res, m.ID)

	if notFound == true {
		err = app.Err("未找到记录")
	}
	return res, err
}
