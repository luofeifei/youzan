package handler

import (
	"base/model/imp/serverShop"
	"base/model/modelSql/mall_shop"
	"base/model/mongoSql"
	"base/pkg/app"
	"base/server/pkg/database/orm"
	"base/tools"
	"time"

	//"base/tools"
	//"base/tools/encrypt"
	"context"
	"server/service"
)

type ShopSysImp struct {
}

//func (imp *ShopSysImp) CoShopGoodSave(ctx context.Context, req serverShop.CoShopGoods) (res serverShop.ResShopGoods, err error) {
//
//	var shopGoods mall_shop.ShopGoodsClear
//
//	var goodsRuleStock mongoSql.CoShopGoods
//
//	if err = app.UnmarshalJson(req.Info, &shopGoods); err != nil {
//		return
//	}
//	//app.Println("server---CoShopGoodSave-------shopGoods--------------------")
//	//app.Println(shopGoods)
//	//app.Println("server---CoShopGoodSave-------GoodsData--------------------")
//	//app.Println(req.GoodsData.Pic)
//
//	// 转换商品详情开始 start--- pb--->db
//	goodsData := mall_shop.ShopGoodsData{
//		ID:      req.Info.Id,
//		Video:   req.GoodsData.Video,
//		Share:   req.GoodsData.Share,
//		SubName: req.GoodsData.SubName,
//		Content: req.GoodsData.Content,
//	}
//
//	// 用这个转会报错 手动写一个  cannot unmarshal array into Go struct field ShopGoodsData.pic of type string
//	//if err = app.UnmarshalJson(req.GoodsData, &goodsData); err != nil {
//	//	return
//	//}
//
//	//转成["name","string","type","is_multi_row"]
//	if jsonEncode, err := tools.JsonEncode(req.GoodsData.Pic); err != nil {
//		return res, err
//
//	} else {
//		goodsData.Pic = jsonEncode
//	}
//	//转成["name","string","type","is_multi_row"]
//	//app.Println("server---CoShopGoodSave-------shopGoods--------string------------")
//	//app.Println(string(req.GoodsData.Word))
//	// base64 转成字符
//	goodsData.Word = string(req.GoodsData.Word)
//
//	//转换商品详情结束 end--- pb--->db
//	if err = app.UnmarshalJson(req, &goodsRuleStock); err != nil {
//		return
//	}
//	//goodsRuleStock.GoodsAsset.StockNum = req.GoodsAsset.StockNum
//
//
//	result, errInner := service.CoShopGoodSave("",  shopGoods, goodsData, goodsRuleStock)
//
//	// end  转换成新的存储方式
//
//	//result, errInner := service.CoShopGoodSave("", shopGoods, goodsData, goodsRuleStock)
//	if errInner != nil {
//		return res, errInner
//	}
//	err = app.Unmarshal(result, &res)
//	return res, err
//}


func (imp *ShopSysImp) CoShopGoodSave(ctx context.Context, req serverShop.CoShopGoods) (res serverShop.ResShopGoods, err error) {

	app.Println(req.GoodsRule)
	// bytes -> map
	return
	var goodsRule map[int16]serverShop.GoodsRuleList

	//ttt:=string(req.GoodsRule)
	app.Unmarshal(req.GoodsRule,&goodsRule)

	//app.UnmarshalJson(req.GoodsRule,&goodsRule)
	app.Println("Unmarshal------")
	app.Println(goodsRule)
	return
	var shopGoods mall_shop.CoShopGoods
	var shopGoodsData mall_shop.CoShopGoodsData
	//var shopGoodsRule mall_shop.CoShopGoodsRule


	err = app.Unmarshal(req, &shopGoods)
	err = app.Unmarshal(req.AfterSale, &shopGoods)
	err = app.Unmarshal(req.PayLimit, &shopGoods)
	if err != nil {
		return
	}
	err = app.Unmarshal(req.GoodsData, &shopGoodsData)
	if err != nil {
		return
	}

	//var goodsStockDiscount mongoSql.GoodsStockDisCount

	//var distributionStock mongoSql.GoodsDistributionStock
	//if err = app.UnmarshalJson(req.Info, &shopGoods); err != nil {
	//	return
	//}
	//
	//// 转换商品详情开始 start--- pb--->db
	//goodsData := mall_shop.ShopGoodsData{
	//	ID:      req.Info.Id,
	//	Video:   req.GoodsData.Video,
	//	Share:   req.GoodsData.Share,
	//	SubName: req.GoodsData.SubName,
	//	Content: req.GoodsData.Content,
	//	SaleTime:req.GoodsData.SaleTime,
	//	IsShowRest: req.GoodsData.IsShowRest,
	//}
	//
	////转成["name","string","type","is_multi_row"]
	//if jsonPicEncode, err := tools.JsonEncode(req.GoodsData.Pic); err != nil {
	//	return res, err
	//
	//} else {
	//	goodsData.Pic = jsonPicEncode
	//}
	////转成["name","string","type","is_multi_row"]
	//if jsonWordEncode, err := tools.JsonEncode(req.GoodsData.Word); err != nil {
	//	return res, err
	//
	//} else {
	//	goodsData.Word = jsonWordEncode
	//}
	//if jsonElectronicEncode, err := tools.JsonEncode(req.GoodsData.ElectronicCoupon); err != nil {
	//	return res, err
	//
	//} else {
	//	goodsData.ElectronicCoupon = jsonElectronicEncode
	//}
	//
	////转换商品详情结束 end--- pb--->db
	//if err = app.UnmarshalJson(req.GoodsRule, &goodsStockDiscount); err != nil {
	//	return
	//}
	//
	////distributionStock
	//if req.GoodsRule.GoodsAsset.StockNum>0{
	//	// 主商品
	//	if goodsStockDiscount.ID>0{
	//		distributionStock.ID = goodsStockDiscount.ID
	//	}
	//	if req.GoodsRule.GoodsAsset.GoodsId>0{
	//		distributionStock.ID = req.GoodsRule.GoodsAsset.GoodsId
	//	}
	//
	//	distributionStock.Coid = goodsStockDiscount.Coid
	//	distributionStock.StockNum = req.GoodsRule.GoodsAsset.StockNum
	//	distributionStock.SoldNum = req.GoodsRule.GoodsAsset.SoldNum
	//	// sku商品
	//	var skus []mongoSql.SkuStock
	//	if len(goodsStockDiscount.GoodsRuleList)>0{
	//		for i:=0;i< len(goodsStockDiscount.GoodsRuleList);i++{
	//			skus =append(skus,mongoSql.SkuStock{SkuId:goodsStockDiscount.GoodsRuleList[i].SkuId,StockNum:req.GoodsRule.GoodsRuleList[i].GoodsStock.StockNum ,SoldNum: req.GoodsRule.GoodsRuleList[i].GoodsStock.SoldNum})
	//		}
	//	}
	//	distributionStock.SkuStock=skus
	//
	//}
	//
	//result, errInner := service.ShopGoodSave("", shopGoods, goodsData,goodsStockDiscount,distributionStock)
	//
	//// end  转换成新的存储方式
	//if errInner != nil {
	//	return res, errInner
	//}
	//err = app.Unmarshal(result, &res)
	return res, err
}


func (imp *ShopSysImp) ShopGoodSave(ctx context.Context, req serverShop.ShopGoodsSave) (res serverShop.ResShopGoods, err error) {
	//app.Println("server---ShopGoodSave-------req--------------------")
	//app.Println(req)

	//app.Println("server---CoShopGoodSave-------shopGoods--------GoodsAsset------------")
	//app.Println(req.GoodsRule.GoodsAsset)

	var shopGoods mall_shop.ShopGoodsClear

	var goodsStockDiscount mongoSql.GoodsStockDisCount

    var distributionStock mongoSql.GoodsDistributionStock
	if err = app.UnmarshalJson(req.Info, &shopGoods); err != nil {
		return
	}
	//app.Println("server---ShopGoodSave-------shopGoods--------------------")
	//app.Println(shopGoods)
	//app.Println("server---CoShopGoodSave-------GoodsData--------------------")
	//app.Println(req.GoodsData.Pic)

	// 转换商品详情开始 start--- pb--->db
	goodsData := mall_shop.ShopGoodsData{
		ID:      req.Info.Id,
		Video:   req.GoodsData.Video,
		Share:   req.GoodsData.Share,
		SubName: req.GoodsData.SubName,
		Content: req.GoodsData.Content,
		SaleTime:req.GoodsData.SaleTime,
        IsShowRest: req.GoodsData.IsShowRest,
	}

	// 用这个转会报错 手动写一个  cannot unmarshal array into Go struct field ShopGoodsData.pic of type string
	//if err = app.UnmarshalJson(req.GoodsData, &goodsData); err != nil {
	//	return
	//}

	//转成["name","string","type","is_multi_row"]
	if jsonPicEncode, err := tools.JsonEncode(req.GoodsData.Pic); err != nil {
		return res, err

	} else {
		//app.Println("server---ShopGoodSave-------------goodsStockDiscount-----")
		//app.Println(req.GoodsData.Pic)
		//app.Println(jsonPicEncode)
		goodsData.Pic = jsonPicEncode
	}
	//转成["name","string","type","is_multi_row"]
	if jsonWordEncode, err := tools.JsonEncode(req.GoodsData.Word); err != nil {
		return res, err

	} else {
		goodsData.Word = jsonWordEncode
	}
	if jsonElectronicEncode, err := tools.JsonEncode(req.GoodsData.ElectronicCoupon); err != nil {
		return res, err

	} else {
		goodsData.ElectronicCoupon = jsonElectronicEncode
	}
	//app.Println("server---CoShopGoodSave-------shopGoods--------string------------")
	//app.Println(string(req.GoodsData.Word))
	// base64 转成字符
	//goodsData.Word = tools.JsonEncode(req.GoodsData.Word)

	//转换商品详情结束 end--- pb--->db
	if err = app.UnmarshalJson(req.GoodsRule, &goodsStockDiscount); err != nil {
		return
	}

	//distributionStock
	if req.GoodsRule.GoodsAsset.StockNum>0{
		// 主商品
		if goodsStockDiscount.ID>0{
			distributionStock.ID = goodsStockDiscount.ID
		}
		if req.GoodsRule.GoodsAsset.GoodsId>0{
			distributionStock.ID = req.GoodsRule.GoodsAsset.GoodsId
		}

		distributionStock.Coid = goodsStockDiscount.Coid
		distributionStock.StockNum = req.GoodsRule.GoodsAsset.StockNum
		distributionStock.SoldNum = req.GoodsRule.GoodsAsset.SoldNum
		// sku商品
		var skus []mongoSql.SkuStock
		if len(goodsStockDiscount.GoodsRuleList)>0{
			for i:=0;i< len(goodsStockDiscount.GoodsRuleList);i++{
				skus =append(skus,mongoSql.SkuStock{SkuId:goodsStockDiscount.GoodsRuleList[i].SkuId,StockNum:req.GoodsRule.GoodsRuleList[i].GoodsStock.StockNum ,SoldNum: req.GoodsRule.GoodsRuleList[i].GoodsStock.SoldNum})
			}
		}
		distributionStock.SkuStock=skus

	}

	//app.Println("server---CoShopGoodSave-------shopGoods--------string------------")
	//app.Println(goodsStockDiscount.GoodsAsset)
	//goodsRuleStock.GoodsAsset.StockNum = req.GoodsAsset.StockNum

	// start 转换成新的存储方式
	// var goodsStockDiscount mongoSql.GoodsStockDisCount
	// goodsStockDiscount := goodsRuleToStockDiscount(req)
	//app.Println("server---ShopGoodSave-------------shopGoods-----")
	//app.Println(shopGoods)
	//
	//app.Println("server---ShopGoodSave-------------goodsData-----")
	//app.Println(goodsData)

	//app.Println("server---ShopGoodSave-------------goodsStockDiscount-----")
	//app.Println(distributionStock)

	result, errInner := service.ShopGoodSave("", shopGoods, goodsData,goodsStockDiscount,distributionStock)

	// end  转换成新的存储方式

	if errInner != nil {
		return res, errInner
	}
	err = app.Unmarshal(result, &res)
	return res, err
}

func (imp *ShopSysImp) GoodsDistributionDetail(ctx context.Context, req serverShop.ReqId) (gds serverShop.GoodsDistributionStock,err error) {

	res, _:= service.GoodsDistributionDetail("", req.Id)
	err = app.Unmarshal(res, &gds)
	return gds,err
}
func (imp *ShopSysImp) GoodsDetail(ctx context.Context, req serverShop.Uid) (res serverShop.ResShopGoodsSave, err error) {

	//b,_:=service.GoodsDataDetailNew("", req.Id)
	//app.Println("GoodsDataDetailNew---------------")
	//app.Println(b)
	shopGoods, goodsData, goodsRuleStock, err := service.GoodsDetail("", req.Uid,req.Coid)

	if err != nil {
		return res, err
	}
	err = app.Unmarshal(goodsRuleStock, &res.GoodsRule)
	// 为啥要手动赋值
	err = app.Unmarshal(goodsData, &res.GoodsData)
	res.GoodsData.Word = goodsData.Word
	res.GoodsData.SubName = goodsData.SubName
	res.GoodsData.Content = goodsData.Content
    res.GoodsData.IsShowRest = goodsData.IsShowRest
	res.GoodsData.SaleTime = goodsData.SaleTime
	res.GoodsData.ElectronicCoupon = goodsData.ElectronicCoupon
	//app.Println(goodsData)
	//app.Println("handle-----GoodsDetail---Word")
	//app.Println(res.GoodsData)
	err = app.Unmarshal(shopGoods, &res.Info)

	return res, err
}

func (imp *ShopSysImp) GoodsFrontDetail(ctx context.Context, req serverShop.ReqId) (res serverShop.ResShopGoodsSave, err error) {

	//b,_:=service.GoodsDataDetailNew("", req.Id)
	//app.Println("GoodsDataDetailNew---------------")
	//app.Println(b)
	shopGoods, goodsData, goodsRuleStock, err := service.GoodsFrontDetail("", req.Id)

	if err != nil {
		return res, err
	}
	err = app.Unmarshal(goodsRuleStock, &res.GoodsRule)
	// 为啥要手动赋值
	err = app.Unmarshal(goodsData, &res.GoodsData)
	res.GoodsData.Word = goodsData.Word
	res.GoodsData.SubName = goodsData.SubName
	res.GoodsData.Content = goodsData.Content
	//app.Println(goodsData)
	//app.Println("handle-----GoodsDetail---Word")
	//app.Println(res.GoodsData)
	err = app.Unmarshal(shopGoods, &res.Info)
	return res, err
}
func (imp *ShopSysImp) GoodsDelete(ctx context.Context, req serverShop.ShopGoodsPure) (res serverShop.ResultCount, err error) {
	var info mall_shop.ShopGoodsClear
	err = app.Unmarshal(req, &info)
	if err == nil {
		result, err := service.GoodsDelete(info)
		//app.Println(result)
		if err != nil {
			return res, err
		}
		if result > 0 { // 有删除数据
			res.Count = result
			return res, err
		}

		return res, err
	}
	return res, err
}

func (imp *ShopSysImp) GoodsBatchDelete(ctx context.Context, req serverShop.IdArray) (res serverShop.ResultCount, err error) {
	//err = app.Unmarshal(req, &info)
	var ids []int64
	err = app.Unmarshal(req, &ids)
	result, err := service.GoodsBatchDelete(ids)
	//app.Println(result)
	if err != nil {
		return res, err
	}
	if result > 0 { // 有删除数据
		res.Count = result
		return res, err
	}

	return res, err
}

func (imp *ShopSysImp) GoodsBatchOperate(ctx context.Context, req serverShop.ReqCoShopGoodsOperate) (res serverShop.Result, err error) {
	// 验证type 的数据是否合法

	var info mall_shop.CoShopGoodsBatchSettingVO
	err = app.Unmarshal(req, &info)
	if err != nil {
		return
	}

	if len(req.GoodsList) > 0 && req.Type > 0 {
		// 1删除  2正常(上架) 3下架 4停用  5分组 6会员价 7商品名称文字替换  8上架时间 9每人限购 10会员折扣 11限定用户(限定用户有购买的权限) 12 配送方式 13商品模板
		var count int64

		switch req.Type {
		case 1: //
			count, err = service.GoodsBatchDelete(req.GoodsList)
		case 2, 3, 4:
			count, err = service.GoodsBatchState(req.GoodsList, req.Type,time.Now().Unix())
		case 5:
			count, err = service.GoodsChageGroup(req.GoodsList, req.FeatureList.GroupId)
		case 6: // 会员价
			//app.Println("会员价---------------------------------")
			// {"level_or_benefit":1,"discount_type":1,"goods_count":[{"goods_id":3701691,"count_value":[10,20,30]},{"goods_id":3786006,"count_value":[10,20,30]}]}
			//app.Println(req.FeatureList.Member)
			skuNum:=len(req.FeatureList.Member)
			//app.Println(skuNum)
			var gsdc []mongoSql.GoodsDiscount
			for i:=0;i<skuNum;i++{ // 商品循环
				if len(req.FeatureList.Member[i].DiscountType)>0{
					for j:=0;j<len(req.FeatureList.Member[i].DiscountType);j++{
						if len(req.FeatureList.Member[i].DiscountType[j].DiscountValue)>0{
							for k:=0;k<len(req.FeatureList.Member[i].DiscountType[j].DiscountValue);k++{
								gsdc = append(gsdc,mongoSql.GoodsDiscount{
									ID:req.FeatureList.Member[i].GoodsId,
									SkuId: req.FeatureList.Member[i].SkuId,
									BindType:req.FeatureList.Member[i].DiscountType[j].LevelOrBenefit,
									DiscountType: req.FeatureList.Member[i].DiscountType[j].DiscountMethod,
                                    DiscountValue: req.FeatureList.Member[i].DiscountType[j].DiscountValue[k],
								})
							}
						}

					}
				}

				goodsRule := mongoSql.GoodsRule{SkuId:req.FeatureList.Member[i].SkuId,GoodsDiscountList: gsdc}

				count,err=service.GoodsSettingSku(req.Coid,req.FeatureList.Member[i].GoodsId,req.FeatureList.Member[i].SkuId,goodsRule)

			}




			//return
			//goodsNum:=len(req.FeatureList.Member)
			//if len(req.FeatureList.Member) > 0 {
			//	for i:=0;i<goodsNum;i++ {
			//
			//		// 修改SKU商品 查出来 赋值 再更新
			//
			//		goodsStockRule, _ := service.GoodsStockRule("", req.FeatureList.Member[i].GoodsId)
			//		//app.Println(goodsStockRule)
			//		//app.Println(goodsStockRule.GoodsRule)
			//		//app.Println("----------------------------------------------------")
			//		//app.Println(len(goodsStockRule.GoodsRule))                               // 2
			//		//app.Println(goodsStockRule.GoodsRule[0].Name)                            // 白色"
			//		//app.Println(len(goodsStockRule.GoodsRule[0].Alias))                      // 3
			//		//app.Println(goodsStockRule.GoodsRule[0].Alias[0].Name)                   // 39码
			//		//app.Println(len(goodsStockRule.GoodsRule[0].Alias[0].Alias))             // 2
			//		//app.Println(goodsStockRule.GoodsRule[0].Alias[0].Alias[0].Name)          // 男款
			//		//app.Println(len(goodsStockRule.GoodsRule[0].Alias[0].Alias[0].Alias))    // 1
			//		//app.Println(goodsStockRule.GoodsRule[0].Alias[0].Alias[0].Alias[0].Name) // ""
			//
			//		// 最多只有三级
			//		// 找到底 白色-39码-男款 会员价 赋值
			//		if len(goodsStockRule.GoodsRule) > 0 { // 第一级
			//			for i := 0; i < len(goodsStockRule.GoodsRule); i++ {
			//				if len(goodsStockRule.GoodsRule[i].Alias) > 0 { // 第二级
			//					for j := 0; j < len(goodsStockRule.GoodsRule[i].Alias); j++ {
			//						if len(goodsStockRule.GoodsRule[i].Alias[j].Alias) > 0 { // 第三级
			//							for k := 0; k < len(goodsStockRule.GoodsRule[i].Alias[j].Alias); k++ {
			//
			//								//app.Println(goodsStockRule.GoodsRule[i].Name)
			//								//app.Println(goodsStockRule.GoodsRule[i].Alias[j].Name)
			//								//app.Println(goodsStockRule.GoodsRule[i].Alias[j].Alias[k].Name)
			//
			//								app.Unmarshal(req.FeatureList.Member, &goodsStockRule.GoodsRule[i].Alias[j].Alias[k].Member)
			//
			//								//app.Println(req.FeatureList.Member)
			//								//app.Println(goodsStockRule.GoodsRule[i].Alias[j].Alias[k].Member)
			//							}
			//
			//						} else { //只有两级
			//
			//							app.Unmarshal(req.FeatureList.Member, &goodsStockRule.GoodsRule[i].Alias[j].Member)
			//						}
			//					}
			//				} else { //  只有一级 如蓝色和白色 要设置会员价
			//					app.Unmarshal(req.FeatureList.Member, &goodsStockRule.GoodsRule[i].Member)
			//				}
			//			}
			//
			//		}
			//		// SKU商品设置完成 再设置主商品的会员价
			//
			//		//goodsStockRule[i].ID=req.FeatureList.Member[i].GoodsId
			//		//err = app.Unmarshal(req.FeatureList.Member[i], &goodsStockRule.Member)
			//
			//		//app.Println(goodsStockRule)
			//		count, err = service.GoodsSettingDiscount(goodsStockRule)
			//
			//		//
			//		if err!=nil || count==0{
			//			return
			//		}
			//	}
			//
			//
			//
			//
			//}

		case 7, 8, 9, 10, 11, 12, 13: //7商品名称文字替换 8上架时间 9每人限购 10会员折扣 11限定用户(限定用户有购买的权限) 12 配送方式 13商品模板

			//app.Println("default--------7,8,9,10,11,12,13-------------")
			//app.Println(req.Type)
			count = service.GoodsBatchSetting(info)
		default:
			app.Println("default---------------------")
			return
		}
		if err == nil && count > 0 { // 有删除或修改数据
			res = serverShop.Result{
				Msg:    "操作成功",
				Code:   1,
				Header: nil,
				Body:   []byte(""),
			}
		} else {
			res = serverShop.Result{
				Msg:    "操作失败",
				Code:   2,
				Header: nil,
				Body:   []byte(""),
			}
		}
		return
	}
	return
}

func (imp *ShopSysImp) GoodsListFrontSearchPage(ctx context.Context, req serverShop.ReqFrontSearchInfo) (res serverShop.ResFrontGoodsList, err error) {
	var info mall_shop.SearchGoods
	var ormPage orm.IndexPage
	app.Unmarshal(req.PageInfo, &ormPage)
	app.Unmarshal(req.SearchInfo, &info)
	if err != nil {
		return res, err
	} else {

		res, err = service.GoodsListFrontSearchPage("*", info, ormPage)

		if err != nil {
			return res, err
		}
		res.Count = int32(len(res.List))
		return res, err
	}
	return res, err
}

func (imp *ShopSysImp) GoodsListPage(ctx context.Context, req serverShop.ReqShopGoodsPage) (res serverShop.ResShopGoodsList, err error) {
	var info mall_shop.ShopGoods
	var ormPage orm.IndexPage

	err = app.Unmarshal(req.Req, &info)
	if err != nil {
		return res, err
	}
	err = app.Unmarshal(req.Page, &ormPage)
	if err != nil {
		return res, err
	} else {

		result, _, err := service.GoodsListPage("*", info, ormPage)
		if err != nil {
			return res, err
		}
		res.Count = int32(len(result))
		err = app.Unmarshal(result, &res.List)
		return res, err
	}
	return res, err
}

func (imp *ShopSysImp) GoodsListSearchPage(ctx context.Context, req serverShop.ReqSearchGoods) (res serverShop.ResFrontGoodsList, err error) {
	var info mall_shop.SearchGoods
	var ormPage orm.IndexPage

	err = app.Unmarshal(req.SearchInfo, &info)
	if err != nil {
		return res, err
	}

	err = app.Unmarshal(req.Page, &ormPage)
	if err != nil {
		return res, err
	} else {

		res, err := service.GoodsListSearchPage("*", info, ormPage)
		res.Count = int32(len(res.List))
		if err != nil {
			return res, err
		}
		return res, err
	}
	return
}

/**
TODO: 修改/添加商品分组
Roger 2020.11.13
*/
func (imp *ShopSysImp) ShopGroupSave(ctx context.Context, req serverShop.ShopGroup) (res serverShop.ShopGroup, err error) {

	var info mall_shop.ShopGroup
	err = app.Unmarshal(req, &info)

	if err != nil {
		return res, err
	}

	result, err := service.SaveGroupData("", info)
	if err != nil {
		return res, err
	}
	err = app.Unmarshal(result, &res)
	return res, err

}
func (imp *ShopSysImp) GroupListByCoid(ctx context.Context, req serverShop.ReqGoodsGroupCoid) (res serverShop.ResGoodsGroup, err error) {
	var info mall_shop.ShopGroup
	err = app.Unmarshal(req, &info)
	if err == nil {
		result, err := service.GroupListByCoid(app.GetContext(ctx, "elect"), info)
		if err != nil {
			return res, err
		}
		res.Count = int32(len(result))
		err = app.Unmarshal(result, &res.List)
		return res, err
	}
	return res, err
}
func (imp *ShopSysImp) GroupListByCoidPage(ctx context.Context, req serverShop.ReqGoodsGroupObj) (res serverShop.ResGoodsGroupList, err error) {
	var info mall_shop.ShopGroupVO
	var ormPage orm.IndexPage
	err = app.Unmarshal(req.Req, &info)
	if err != nil {
		return res, err
	}
	err = app.Unmarshal(req.Page, &ormPage)
	if err != nil {
		return res, err
	} else {
		result, err := service.GroupListByCoidPage("*", info, ormPage)
		if err != nil {
			return res, err
		}
		res.Count = int32(len(result))
		err = app.Unmarshal(result, &res.List)
		return res, err
	}
	return res, err
}

func (imp *ShopSysImp) GroupDelete(ctx context.Context, req serverShop.ShopGroup) (res serverShop.ResultCount, err error) {
	var info mall_shop.ShopGroup
	err = app.Unmarshal(req, &info)
	if err == nil {
		result, err := service.GroupDelete(info)
		if err != nil {
			return res, err
		}
		if result > 0 { // 有删除数据
			res.Count = result

			app.Println(res)
			return res, err
		}
	}
	return res, err
}

/**
  TODO: 商品库存操作
  Roger 2020.11.13
*/
func (imp *ShopSysImp) StockAdd(ctx context.Context, req serverShop.ShopStock) (res serverShop.ShopStock, err error) {
	var info mall_shop.ShopStock
	err = app.Unmarshal(req, &info)
	if err != nil {
		return res, err
	}
	result, err := service.GoodsStockSave("", info)
	if err != nil {
		return res, app.Err("添加失败")
	}
	err = app.Unmarshal(result, &res)
	return res, err
}

func (imp *ShopSysImp) StockEdit(ctx context.Context, req serverShop.ShopStock) (res serverShop.ShopStock, err error) {
	var info mall_shop.ShopStock
	err = app.Unmarshal(req, &info)
	if err != nil {
		return res, err
	}
	result, err := service.GoodsStockSave("", info)
	if err != nil {
		return res, app.Err("添加失败")
	}
	err = app.Unmarshal(result, &res)
	return res, err
}

func (imp *ShopSysImp) StockDelete(ctx context.Context, req serverShop.ShopStock) (res serverShop.Result, err error) {
	var info mall_shop.ShopStock
	err = app.Unmarshal(req, &info)
	if err == nil {
		result, err := service.StockDelete(info)
		if err != nil {
			return res, err
		}
		if result > 0 { // 有删除数据

			body := []byte("删除成功提醒")
			res = serverShop.Result{
				Msg:  "",
				Code: app.Tips,
				Body: body,
			}
			return res, err
		}
		return res, err
	}
	return res, err
}

/**
  TODO: 商品库存卡操作
  Roger 2020.11.13
*/
func (imp *ShopSysImp) StockCardAdd(ctx context.Context, req serverShop.ShopStockCard) (res serverShop.ShopStockCard, err error) {
	var info mall_shop.ShopStockCard
	err = app.Unmarshal(req, &info)

	if err != nil {
		return res, err
	}
	result, err := service.StockCardSave("", info)
	if err != nil {
		return res, app.Err("添加失败")
	}
	err = app.Unmarshal(result, &res)
	return res, err
}

func (imp *ShopSysImp) StockCardEdit(ctx context.Context, req serverShop.ShopStockCard) (res serverShop.ShopStockCard, err error) {
	var info mall_shop.ShopStockCard
	err = app.Unmarshal(req, &info)
	if err != nil {
		return res, err
	}
	result, err := service.StockCardSave("", info)
	if err != nil {
		return res, app.Err("添加失败")
	}
	err = app.Unmarshal(result, &res)
	return res, err
}

func (imp *ShopSysImp) StockCardDelete(ctx context.Context, req serverShop.ShopStockCard) (res serverShop.ResultCount, err error) {
	var info mall_shop.ShopStockCard
	err = app.Unmarshal(req, &info)
	if err == nil {
		result, err := service.StockCardDelete(info)
		if err != nil {
			return res, err
		}
		if result > 0 { // 有删除数据
			res.Count = result
			return res, err
		}

		return res, err
	}
	return res, err
}

/**
  TODO: 优惠券
  Roger 2020.11.13
*/
func (imp *ShopSysImp) CoCouponSave(ctx context.Context, req serverShop.ShopCoupon) (res serverShop.ShopCoupon, err error) {

	var info mall_shop.ShopCoupon
	err = app.Unmarshal(req, &info)

	//app.Println("CLIENT-====Form=========----")
	limitLevelMember, _ := tools.JsonEncode(req.CouponData.LimitLevelMember)
	limitBenefitMember, _ := tools.JsonEncode(req.CouponData.LimitBenefitMember)
	useGoods, _ := tools.JsonEncode(req.CouponData.UseGoods)
	info.CouponData = mall_shop.ShopCouponData{
		ID:                 req.CouponData.Id,
		LimitLevelMember:   limitLevelMember,
		LimitBenefitMember: limitBenefitMember,
		UseGoods:           useGoods,
		UseRule:            req.CouponData.UseRule,
		Description:        req.CouponData.Description,
	}

	if err != nil {
		return res, err
	}
	result, err := service.CoCouponSave("", info)
	if err != nil {
		return res, err
	}
	err = app.Unmarshal(result, &res)
	return res, err

}

func (imp *ShopSysImp) CouponDetail(ctx context.Context, req serverShop.ReqId) (res serverShop.ResShopCoupon, err error) {
	var info mall_shop.ShopCoupon
	err = app.Unmarshal(req, &info)
	if err != nil {
		return res, err
	}
	result, err := service.CouponDetail("", info)

	//app.Println("CouponDetail--result---SD-----------")
	//app.Println(result.SD)
	if err != nil {
		return res, err
	}
	err = app.Unmarshal(result.SC, &res)
	err = app.Unmarshal(result.SD, &res.CouponData)

	return res, err
}

func (imp *ShopSysImp) CouponDelete(ctx context.Context, req serverShop.ShopCoupon) (res serverShop.Result, err error) {
	var info mall_shop.ShopCoupon
	err = app.Unmarshal(req, &info)
	if err == nil {
		result, err := service.CouponDelete(info)
		if err != nil {
			return res, err
		}
		if result > 0 { // 有删除数据

			body := []byte("删除成功提醒")
			res = serverShop.Result{
				Msg:  "删除成功提醒",
				Code: app.Tips,
				Body: body,
			}
			return res, err
		}
		return res, err
	}
	return res, err
}

func (imp *ShopSysImp) CouponListPage(ctx context.Context, req serverShop.ReqCouponPage) (res serverShop.ResShopCouponList, err error) {
	var info mall_shop.ShopCouponRelatedVO
	var ormPage orm.IndexPage
	err = app.Unmarshal(req.Req, &info)
	if err != nil {
		return res, err
	}
	err = app.Unmarshal(req.Page, &ormPage)
	if err != nil {
		return res, err
	} else {
		result, _, err := service.CouponListPage("*", info, ormPage)
		if err != nil {
			return res, err
		}
		res.Count = int32(len(result))

		err = app.Unmarshal(result, &res.List)
		return res, err
	}
	return res, err
}

/**
  TODO: 优惠券兑换卡
  Roger 2020.11.13
*/
func (imp *ShopSysImp) CouponCardAdd(ctx context.Context, req serverShop.ShopCouponCard) (res serverShop.ShopCouponCard, err error) {
	var info mall_shop.ShopCouponCard
	err = app.Unmarshal(req, &info)
	if err != nil {
		return res, err
	}
	result, err := service.CouponCardSave("", info)
	if err != nil {
		return res, app.Err("添加失败")
	}
	err = app.Unmarshal(result, &res)
	return res, err
}

func (imp *ShopSysImp) CouponCardEdit(ctx context.Context, req serverShop.ShopCouponCard) (res serverShop.ShopCouponCard, err error) {
	var info mall_shop.ShopCouponCard
	err = app.Unmarshal(req, &info)
	if err != nil {
		return res, err
	}
	result, err := service.CouponCardSave("", info)
	if err != nil {
		return res, app.Err("添加失败")
	}
	err = app.Unmarshal(result, &res)
	return res, err
}

func (imp *ShopSysImp) CouponCardDelete(ctx context.Context, req serverShop.ShopCouponCard) (res serverShop.Result, err error) {
	var info mall_shop.ShopCouponCard
	err = app.Unmarshal(req, &info)
	if err == nil {
		result, err := service.CouponCardDelete(info)
		if err != nil {
			return res, err
		}
		if result > 0 { // 有删除数据

			body := []byte("删除成功提醒")
			res = serverShop.Result{
				Msg:  "删除成功提醒",
				Code: app.Tips,
				Body: body,
			}
			return res, err
		}
		return res, err
	}
	return res, err
}

func (imp *ShopSysImp) ShopGroupDetail(ctx context.Context, req serverShop.ReqId) (res serverShop.ShopGroup, err error) {
	var info mall_shop.ShopGroup
	err = app.Unmarshal(req, &info)
	if err != nil {
		return res, err
	}
	result, err := service.ShopGroupDetail("", info)
	if err != nil {
		return res, err
	}
	err = app.Unmarshal(result, &res)
	return res, err
}
func (imp *ShopSysImp) StockDetail(ctx context.Context, req serverShop.ReqId) (res serverShop.ShopStock, err error) {
	var info mall_shop.ShopStock
	err = app.Unmarshal(req, &info)
	if err != nil {
		return res, err
	}
	result, err := service.StockDetail("", info)
	if err != nil {
		return res, err
	}
	err = app.Unmarshal(result, &res)
	return res, err
}
func (imp *ShopSysImp) StockCardDetail(ctx context.Context, req serverShop.ReqId) (res serverShop.ShopStockCard, err error) {
	var info mall_shop.ShopStockCard
	err = app.Unmarshal(req, &info)
	if err != nil {
		return res, err
	}
	result, err := service.StockCardDetail("", info)
	if err != nil {
		return res, err
	}
	err = app.Unmarshal(result, &res)
	return res, err
}

func (imp *ShopSysImp) CouponCardDetail(ctx context.Context, req serverShop.ReqId) (res serverShop.ShopCouponCard, err error) {
	var info mall_shop.ShopCouponCard
	err = app.Unmarshal(req, &info)
	if err != nil {
		return res, err
	}
	result, err := service.CouponCardDetail("", info)
	if err != nil {
		return res, err
	}

	err = app.Unmarshal(result, &res)
	return res, err
}
func (imp *ShopSysImp) FakerBuy (ctx context.Context, req serverShop.Faker) (res serverShop.Faker, err error) {
	//remoteSpike:=server.RemoteSpikeKeys{
	//	TotalInventoryKey:"ticket_total_nums",
	//}
	//remoteSpike.RemoteDeductionStock()


	return
}
//func (imp *ShopSysImp ) GoodsBatchSetting(ctx context.Context, req serverShop.CoShopGoodsBatchSetting)(res serverShop.Result, err error)  {
//	// 验证type 的数据是否合法
//
//	if len(req.GoodsList)>0 && req.Type>0{
//		// 1商品名称  2上架时间 3每人限购 4限定用户(限定用户有购买的权限) 5会员折扣 6 配送方式 7商品模板
//		// handle验证处理 不到数据库模型里面做
//		var info mall_shop.CoShopGoodsBatchSettingVO
//		err = app.Unmarshal(req, &info)
//		if err!=nil{
//			return
//		}
//		err = service.GoodsBatchSetting(info)
//
//		return res, err
//	}
//	return
//}
