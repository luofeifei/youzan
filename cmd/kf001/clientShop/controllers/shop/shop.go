package shop

import (
	"base/client/pkg/comm"
	"base/client/pkg/middlewares"
	"base/client/pkg/validator"
	"base/model/client"
	"base/model/imp/serverShop"
	"base/pkg/app"
	"client/request"
	mytool "client/tools"
	"fmt"
	"github.com/gin-gonic/gin"
)

// @Tags co_shop_goods
// @Summary 商品编辑
// @description 商品编辑接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.ShopGoodsSave  true "data"
// @Success 200 {object} comm.Response
// @Router /goods/edit [put]
func EditGoods(ctx *gin.Context) {
	CoShopGoodsSave(ctx, "edit")
	return
}

// @Tags co_shop_goods
// @Summary 商品添加
// @description 商品添加接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.ShopGoods  true "data"
// @Success 200 {object} comm.Response
// @Router /goods/add [put]
func AddGoods(ctx *gin.Context) {
	//CoShopGoodsSave(ctx, "add")
	ShopGoodsSave(ctx, "add")
	return
}

// @Tags co_shop_goods
// @Summary 获取商品详情
// @description 根据商品id获取一件商品详情
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.Id true "商品ID"
// @Success 200 {object} comm.Response
// @Router /goods/detail [POST]
func GoodsDetail(ctx *gin.Context) {
	var form request.Id
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)

	// 先去redis缓存里面拿 商品的库存信息
	//var gds serverShop.GoodsDistributionStock
	////key:=strconv.FormatInt(form.ID+1,10)
	//key:=strconv.FormatInt(form.ID,10)
	//app.Println("key--------------------------")
	//key = "goods:7571707:10011"
	//app.Println(key)
	//
	//match:=fmt.Sprintf("goods:%d:*",form.ID)
	//count:=int64(999)
	////var duration time.Duration = (86400) * time.Second
	//goodsDistribution := redis.CacheGet(key, duration, func() string {
	//	//请求商品的库存信息
	//	var result string
	//	res, _ := client.ServerShopSys().GoodsDistributionDetail(serverShop.ReqId{Id: form.ID})
	//	app.Println("sql--------")
	//	if res.Id > 0 {
	//		result, _ = tools.JsonEncode(res)
	//	}
	//	return result
	//})


	//goodsDistribution:=redis.BatchHashGet(key,"sku_stock")
	//hmset goods:7571707:10011 coid 100000 stock_num 1199708  sold_num 7
	//goodsDistribution,_:=redis.BatchHashGetAll(key)
	//app.Println("goodsDistribution--------------------------------")
	//app.Println(goodsDistribution)
    //app.Unmarshal(goodsDistribution,&gds)
    //app.Println(len(goodsDistribution))
    //app.Println(match)
	//app.Println(count)
	//resAll:=redis.Client.HScan()
	//Scan(0,match,0)
	app.Println("resAll-----------------------------")
	//app.Println(resAll)
	//if len(goodsDistribution)>0 {
	//
	//	var	skuStocks []serverShop.SkuStock
	//	var	id int64
	//	var	coid int64
	//	for k, v := range goodsDistribution {
	//		item:=serverShop.SkuStock{}
	//		if strings.Contains(k,":stock_num") { // 库存
	//			itemSlice:=strings.Split(k,":")
	//			item.SkuId,_=strconv.ParseInt(itemSlice[2],10,64)
	//			item.StockNum,_=strconv.ParseInt(v,10,64)
	//
	//		}
	//
	//		skuStocks=append(skuStocks,item)
	//	}
	//
	//	if resId,ok:=goodsDistribution["id"];ok{
	//		id ,_= strconv.ParseInt(resId,10,64)
	//		gds.Id =id
	//	}
	//	if resCoid,ok:=goodsDistribution["coid"];ok{
	//		coid ,_= strconv.ParseInt(resCoid,10,64)
	//		gds.Id =coid
	//	}
	//	app.Unmarshal(skuStocks,&gds.SkuStock)
	//	app.Println("gds.SkuStock------------------------")
	//	app.Println(gds)
	//
	//}else{
	//	// 存入redis中
	//	monRes, _ := client.ServerShopSys().GoodsDistributionDetail(serverShop.ReqId{Id: form.ID})
	//	app.Println("sql----get--gds--")
	//	app.Unmarshal(monRes,&gds)
	//	app.Println(gds)
	//	app.Println(monRes)
	//
	//	if monRes.Id > 0 {
	//		//result, _ := tools.JsonEncode(res)
	//		// //hmset goods:7571707:10010 coid 100000 stock_num 1199707  sold_num 6
	//		var redisKeyValues map[string]map[string]interface{}
    //        if len(monRes.SkuStock)>0{
	//			for i:=0;i<len(monRes.SkuStock);i++{
	//				var redisfieldValues map[string]interface{}
	//				redisKey := fmt.Sprintf("goods:%d:%d",monRes.Id, monRes.SkuStock[i].SkuId)
	//				redisFieldCoidValue :=fmt.Sprintf("%d",monRes.Coid)
	//				redisValueStockNum:= fmt.Sprintf("%d",monRes.SkuStock[i].StockNum)
	//				redisValueSoldNum:= fmt.Sprintf("%d",monRes.SkuStock[i].SoldNum)
	//				redisfieldValues["coid"] = redisFieldCoidValue
	//				redisfieldValues["stock_num"] = redisValueStockNum
	//				redisfieldValues["sold_num"] = redisValueSoldNum
	//
	//				redisKeyValues[redisKey]=redisfieldValues
	//			}
	//
	//			if len(redisKeyValues)>0{
	//				for k,v := range redisKeyValues {
	//					settingStr:=redis.BatchHashSet(k,v)
	//					app.Println(settingStr)
	//				}
	//			}
	//		}
	//
	//		//var result map[string]interface{}
	//		//app.Unmarshal(result,&result)
	//		////result["sku_stock"]= app.Struct2Json(result["sku_stock"]) // hash表里面存不了切片类型sku_stock 需要转换成string
	//		//app.Println(result)
	//		//setString:=redis.BatchHashSet(key,result)
	//		//app.Println("setString---")
	//		//app.Println(setString)
	//		//if len(setString)>0{
	//		//
	//		//}
	//	}
	//
	//}



	res, err := client.ServerShopSys().GoodsDetail(serverShop.Uid{Uid: form.ID,Coid: coId})
	//app.Unmarshal(gds, &res.GoodsDisStock)
	comm.ApiResponse(ctx,&res,err,0,true)

}

// 公用添加 修改商品操作
func CoShopGoodsSave(ctx *gin.Context, way string) {
	var form request.ShopGoodsSave
	//var form request.ShopGoodsSave
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	app.Println("CoShopGoodsSave----------")
	app.Println(form)

	return
	// 验证 规格 1文本 2数字 3时间 4身份证
	if  len(form.GoodsStockDisCount.GoodsRuleList)>0{
		for i:=0;i<len(form.GoodsStockDisCount.GoodsRuleList);i++ {

			//app.Println(form.GoodsStockDisCount.GoodsRuleList[i])

			if len(form.GoodsStockDisCount.GoodsRuleList[i].GoodsRuleInfo)>0{
				for j:=0;j<len(form.GoodsStockDisCount.GoodsRuleList[i].GoodsRuleInfo);j++ {

					//app.Println(form.GoodsStockDisCount.GoodsRuleList[i].GoodsRuleInfo[j].Type)
					//app.Println(form.GoodsStockDisCount.GoodsRuleList[i].GoodsRuleInfo[j].Value)
					if form.GoodsStockDisCount.GoodsRuleList[i].GoodsRuleInfo[j].Type==1{ // 文本


					} else if form.GoodsStockDisCount.GoodsRuleList[i].GoodsRuleInfo[j].Type==2{ // 数字
						if !mytool.IsDigital(form.GoodsStockDisCount.GoodsRuleList[i].GoodsRuleInfo[j].Value){
							//app.Println("数字规格校验有误")
							err:= app.Err("数字规格校验有误")
							comm.ApiResponse(ctx,nil,err,0,true)
							break

						}

					}else if form.GoodsStockDisCount.GoodsRuleList[i].GoodsRuleInfo[j].Type==3{ // 时间
						if !mytool.IsRuleTime(form.GoodsStockDisCount.GoodsRuleList[i].GoodsRuleInfo[j].Value){
							//app.Println("时间规格校验有误")
							err:= app.Err("时间规格校验有误")
							comm.ApiResponse(ctx,nil,err,0,true)
							break

						}

					}else if form.GoodsStockDisCount.GoodsRuleList[i].GoodsRuleInfo[j].Type==4{ // 身份证
						if !mytool.IsIdCard(form.GoodsStockDisCount.GoodsRuleList[i].GoodsRuleInfo[j].Value){
							//app.Println("身份证规格校验有误")
							err:= app.Err("身份证规格校验有误")
							comm.ApiResponse(ctx,nil,err,0,true)
							break

						}

					}
				}
			}
		}
	}

	if way == "edit" {
		if err := validator.Get().Var(form.ShopGoodsInfo.ID, "required"); err != nil {
			comm.ApiResponse(ctx,nil,app.Err("ID值异常"),0,true)
		}
	}
	var Form serverShop.ShopGoodsSave
	if err := app.UnmarshalJson(form, &Form); err != nil {
		comm.ApiResponse(ctx,nil,err,0,true)
	}
	//app.Println("Form----")
	//app.Println(Form)
	//return
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	Form.Info.Coid = coId
	Form.GoodsRule.Coid = coId
	//res, err := client.ServerShopSys().CoShopGoodSave(Form)
	//comm.ApiResponse(ctx,&res,err,0,true)


}



// 公用添加 修改商品操作
func ShopGoodsSave(ctx *gin.Context, way string) {
	var form request.ShopGoods
	//var form request.ShopGoodsSave
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	//app.Println("CoShopGoodsSave----------")
	//app.Println(form)
    // 验证 商品价格库存规则 1->价格必填 库存必填 2->规则信息 1文本 2数字 3时间 4身份证
    if len(form.GoodsRule)>0{
    	for skuId,item:= range form.GoodsRule{
    		if !(item.GoodsPriceName.Price>0){

				err:= app.Err(fmt.Sprintf("Id为%d的商品价格输入有误",skuId))
				comm.ApiResponse(ctx,nil,err,0,true)
				break
			}
			if !(item.GoodsPriceStock.StockNum>0){
				err:= app.Err(fmt.Sprintf("Id为%d的商品库存数量输入有误",skuId))
				comm.ApiResponse(ctx,nil,err,0,true)
				break
			}
			if(skuId!=0){ // 排除没有skuId的商品 (只有一个总商品)
				if len(item.GoodsRuleListInfo)>0{
					for _,v:=range item.GoodsRuleListInfo{

						switch v.Type {
						case 1:

						case 2:// 数字
							if !mytool.IsDigital(v.Value){
								//app.Println("数字规格校验有误")
								err:= app.Err("规则信息,数字规格校验有误")
								comm.ApiResponse(ctx,nil,err,0,true)
								break
							}
						case 3: // 时间
							if !mytool.IsRuleTime(v.Value){
								  //app.Println("时间规格校验有误")
								 err:= app.Err("规则信息,时间规格校验有误")
								 comm.ApiResponse(ctx,nil,err,0,true)
								 break
							}

						case 4: // 身份证
							 if !mytool.IsIdCard(v.Value){
								 //app.Println("身份证规格校验有误")
								 err:= app.Err("规则信息,身份证规格校验有误")
								 comm.ApiResponse(ctx,nil,err,0,true)
								 break

							 }
						default:
						}
					}
				 }
			}
		}
	}
	//1 虚拟物品(限购和起售) 2实物物品(限购和起售) 3电子卡券（限购和起售） 4付费优惠券（限购和起售） 5 酒店（限购） 6蛋糕烘培（限购和起售）
	//商品购买限制规则 只允许特定用户购买 如果指定了那些用户可购买 bind_type 只能是(1 会员卡 2 权益卡 3用户标签) 并且bind_id数组不能为空
	if len(form.PayLimit.LimitSpecific)>0 {
		for _,it :=range form.PayLimit.LimitSpecific{
			if it.BindType>0{
				if !(len(it.BindId)>0){
					err:= app.Err("请选择特定用户")
					comm.ApiResponse(ctx,nil,err,0,true)
					break
				}
			}else{
				err:= app.Err("请选择允许特定用户购买的会员/会员卡/用户标签")
				comm.ApiResponse(ctx,nil,err,0,true)
				break
			}
		}

	}

	switch form.Type {
	case 1: // 虚拟物品
		//app.Println("虚拟物品-----")
	case 2: // 实物物品
		//app.Println("实物物品-----")
	    //app.Println(len(form.GoodsData.SpecialRules.GoodsReal.DeliveryType))
		if !mytool.IsNil(form.GoodsData.SpecialRules.GoodsReal)  {
			if !(len(form.GoodsData.SpecialRules.GoodsReal.DeliveryType)>0){ // 配送方式必填
				// 如果给了就不合法
				err:= app.Err("配送方式必填")
				comm.ApiResponse(ctx,nil,err,0,true)
				break
			}
		}
	case 3: // 电子卡券
		// validity_mode 卡券生效模式 1 立即生效 2 次日生效 3多少小时后生效 为3必须指定小时（validity_day）
		// validity_type 卡券生效类型 ( -1长期有效 1指定天数 2指定时间段) 为2 validity_start_time和validity_end_time 必须有值
		if !mytool.IsNil(form.GoodsData.SpecialRules.GoodsCard)  {

			if form.GoodsData.SpecialRules.GoodsCard.ValidityType==-1{ // 长期有效 不需要再给指定天数和指定时间段的值了
				// 如果给了就不合法
				if form.GoodsData.SpecialRules.GoodsCard.ValidityDay>0{
					err:= app.Err("您选择的是长期有效,不需要再传指定天数了")
					comm.ApiResponse(ctx,nil,err,0,true)
				}
				if form.GoodsData.SpecialRules.GoodsCard.ValidityStartTime>0{
					err:= app.Err("您选择的是长期有效,不需要再传有效期开始时间了")
					comm.ApiResponse(ctx,nil,err,0,true)
				}

			} else  if form.GoodsData.SpecialRules.GoodsCard.ValidityType==1{ // 指定天数

				if !(form.GoodsData.SpecialRules.GoodsCard.ValidityDay>0){
					err:= app.Err("您选择的是指定天数有效,请输入合法的天数")
					comm.ApiResponse(ctx,nil,err,0,true)
				}

			}else  if form.GoodsData.SpecialRules.GoodsCard.ValidityType==2{ // 指定时间段
				if !(form.GoodsData.SpecialRules.GoodsCard.ValidityStartTime>0) || !(form.GoodsData.SpecialRules.GoodsCard.ValidityEndTime>0){
					err:= app.Err("您选择的是指定时间段有效,请输入合法的时间段")
					comm.ApiResponse(ctx,nil,err,0,true)
				}
			}
	    }
	case 4: // 付费优惠券
		//fmt.Printf("3")
	case 5: // 酒店
		//fmt.Printf("4, 5, 6")
	case 6: // 蛋糕烘培
		//app.Println("GoodsCake---len")
		//app.Println(len(form.GoodsData.SpecialRules.GoodsCake.StockUp))
		//app.Println(len(form.GoodsRule))
		//app.Println(form.GoodsData.SpecialRules.GoodsCake)
		if !mytool.IsNil(form.GoodsData.SpecialRules.GoodsCake)  {
            // -1 不需要预留备货时间 其他
			if form.GoodsData.SpecialRules.GoodsCake.StockUpTime==-1{ // -1 启用不同规格单独设置备货时间 校验
				if len(form.GoodsData.SpecialRules.GoodsCake.StockUp)>0 && len(form.GoodsData.SpecialRules.GoodsCake.StockUp)==len(form.GoodsRule){
					// 验证备货时间 map[SKUID]时间 单位分钟
					for _,v:=range form.GoodsData.SpecialRules.GoodsCake.StockUp{
						if !(v>0){
							err:= app.Err("您输入的备货时间有误,请重新输入")
							comm.ApiResponse(ctx,nil,err,0,true)
							break
						}
					}
				}else {
					err:= app.Err("您输入的不同规格备货时间有误,请重新输入")
					comm.ApiResponse(ctx,nil,err,0,true)
					break
				}
			}else { // 不启用不同规格单独设置备货时间 不同规格只用一个时间
				if !(form.GoodsData.SpecialRules.GoodsCake.StockUpTime>0){
					err:= app.Err("您输入的主备货时间有误,请重新输入")
					comm.ApiResponse(ctx,nil,err,0,true)
					break
				}
			}
		}
	default:
		//fmt.Printf("Default")
	}


	var Form serverShop.CoShopGoods
	if err := app.UnmarshalJson(form, &Form); err != nil {
		comm.ApiResponse(ctx,nil,err,0,true)
	}
	//Form.GoodsRule = app.Struct2Json(form.GoodsRule)
	app.Println("Form----")
	app.Println(Form.GoodsRule)
	//return



	//coId, _ := middlewares.GetFastCoIdUid(ctx)
	coId := int64(100000)
	Form.Coid = coId
	//res, err := client.ServerShopSys().CoShopGoodSave(Form)
	//comm.ApiResponse(ctx,&res,err,0,true)


}

// @Tags co_shop_goods
// @Summary 商品批量操作
// @description 商品批量操作接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.CoShopGoodsOperate true "data"
// @Success 200 {object} comm.Response
// @Router /goods/operate [put]
func GoodsBatchOperate(ctx *gin.Context) {
	var form request.CoShopGoodsOperate
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	//校验type和对应的值是否合法
	var err error
	if len(form.GoodsList) > 0 && form.Type > 0 {
		// 1删除  2正常(上架) 3下架 4停用  5分组 6会员价 7商品名称替换  8上架时间 9每人限购 10会员折扣 11限定用户(限定用户有购买的权限) 12 配送方式 13商品模板
		switch form.Type {

		case 1,2, 3, 4:
			if !(len(form.GoodsList) > 0) {
				err = app.Err("商品id有误")
			}
		case 5:
			if !(form.FeatureList.GroupId > 0) {
				err = app.Err("分组id有误")
			}
		case 6: // 会员价
			//app.Println(form.FeatureList.Member)
			if !(len(form.FeatureList.Member) > 0) {
				err = app.Err("会员价设置有误")
			}
		case 7: // 商品名称替换
			if !(len(form.FeatureList.NameChange.NameList) > 0) || !(len(form.FeatureList.NameChange.OldName) > 0) {
				err = app.Err("商品名称文字输入有误")
			}
		case 8: // 上架时间
			if !(form.FeatureList.StartTime > 0) {
				err = app.Err("上架时间设置有误")
			}

		case 9: // 每人限购
			if !(form.FeatureList.Quota > 0) || !(form.FeatureList.QuotaNum > 0) {
				err = app.Err("用户限购填写有误")
			}
		case 10: // 会员折扣
			if !(form.FeatureList.IsVip > 0) {
				err = app.Err("会员是否参与折扣输入有误")
			}
		case 11: // 限定用户(限定用户有购买的权限)
			err = app.Err("限定用户待开发")
		case 12: // 配送方式
			err = app.Err("配送方式待开发")
		case 13: // 13商品模板
			err = app.Err("商品模板待开发")
		default:
			err = app.Err("未知类型")
		}

		if err != nil {
			comm.ApiResponse(ctx,nil,err,0,true)
		}

	} else {
		comm.ApiResponse(ctx,nil,app.Err("商品列表或操作类型有误"),0,true)
	}

	var Form serverShop.ReqCoShopGoodsOperate

	coId, _ := middlewares.GetFastCoIdUid(ctx)
	Form.Coid = coId

	if err := app.UnmarshalJson(form, &Form); err != nil {
		comm.ApiResponse(ctx,nil,err,0,true)
		return
	}
	//app.Println(Form)
	res, err := client.ServerShopSys().GoodsBatchOperate(Form)
	comm.ApiResponse(ctx,&res,err,0,true)
}

// @Tags co_shop_goods
// @Summary 筛选商品分页
// @description 筛选商品分页接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.SearchInfo  true "data"
// @Success 200 {object} comm.Response
// @Router /goods/pageSearch [post]
func GoodsListSearchPage(ctx *gin.Context) {
	var form request.SearchInfo
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	var Form serverShop.ReqSearchGoods
	if err := app.UnmarshalJson(form, &Form.SearchInfo); err != nil {
		comm.ApiResponse(ctx,nil,err,0,true)
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	Form.SearchInfo.Coid = coId
	if err := app.UnmarshalJson(form.Page, &Form.Page); err != nil {
		comm.ApiResponse(ctx,nil,err,0,true)
	}
	res, err := client.ServerShopSys().GoodsListSearchPage(Form)
	comm.ApiResponse(ctx,res,err,0,true)
}

// @Tags shop_goods_front
// @Summary 根据关键字搜索商品
// @description 根据关键字搜索商品分页接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param data body request.FrontSearchInfo  true "data"
// @Success 200 {object} comm.Response
// @Router /goods/pageFrontSearch [post]
func GoodsListFrontSearchPage(ctx *gin.Context) {
	var form request.FrontSearchInfo
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	var Form serverShop.ReqFrontSearchInfo
	if err := app.UnmarshalJson(form, &Form); err != nil {
		comm.ApiResponse(ctx,nil,err,0,true)
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	Form.SearchInfo.Coid = coId
	res, err := client.ServerShopSys().GoodsListFrontSearchPage(Form)
	comm.ApiResponse(ctx,&res,err,0,true)
}

// @Tags shop_goods_front
// @Summary 刷新购物车里的商品
// @description 刷新购物车里的商品接口
// @accept json
// @Produce json
// @Param data body request.FrontSearchInfo  true "data"
// @Success 200 {object} comm.Response
// @Router /goods/refresh [post]
func QueryGoodsByIds(ctx *gin.Context) {
	var form request.FrontSearchInfo
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	var Form serverShop.ReqFrontSearchInfo
	if err := app.UnmarshalJson(form, &Form); err != nil {
		comm.ApiResponse(ctx,nil,err,0,true)
		//comm.TipResponse(ctx, err, 0, nil)
		return
	}
	res, err := client.ServerShopSys().GoodsListFrontSearchPage(Form)
	comm.ApiResponse(ctx,&res,err,0,true)

	return

}

// @Tags shop_goods_front
// @Summary 获取商品详情
// @description 根据商品id获取一件商品详情
// @accept json
// @Produce json
// @Param data body request.Id true "商品ID"
// @Success 200 {object} comm.Response
// @Router /goods/detailFront [POST]
func GoodsFrontDetail(ctx *gin.Context) {
	var form request.Id
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	res, err := client.ServerShopSys().GoodsFrontDetail(serverShop.ReqId{Id: form.ID})
	// pb返回数据要进行格式转换 BASE64的值
	comm.ApiResponse(ctx,&res,err,0,true)
}

// @Tags co_shop_group
// @Summary 添加商品分组
// @description 添加商品分组接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.ShopGroup  true "data"
// @Success 200 {object} comm.Response
// @Router /group/add [put]
func AddGroup(ctx *gin.Context) {
	CoShopGroupSave(ctx, "add")
	return
}

// @Tags co_shop_group
// @Summary 修改商品分组
// @description 修改商品分组接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.ShopGroup  true "data"
// @Success 200 {object} comm.Response
// @Router /group/edit [put]
func EditGroup(ctx *gin.Context) {
	CoShopGroupSave(ctx, "edit")
	return
}

// 公用添加 修改分组操作
func CoShopGroupSave(ctx *gin.Context, way string) {
	var form request.ShopGroup
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	if way == "edit" {
		if err := validator.Get().Var(form.ID, "required"); err != nil {
			comm.ApiResponse(ctx,nil,app.Err("ID值异常"),0,true)
			return
		}
	}
	var Form serverShop.ShopGroup
	if err := app.UnmarshalJson(form, &Form); err != nil {
		comm.ApiResponse(ctx,nil,err,0,true)
		return
	}

	coId, _ := middlewares.GetFastCoIdUid(ctx)
	Form.Coid = coId

	res, err := client.ServerShopSys().ShopGroupSave(Form)
	comm.ApiResponse(ctx,&res,err,0,true)

}

// @Tags co_shop_group
// @Summary 商品分组分页列表
// @description 商品Id获取商品分组列表分页接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.ReqShopGroupPage  true "data"
// @Success 200 {object} comm.Response
// @Router /group/page [post]
func GroupListByCoidPage(ctx *gin.Context) {
	var form request.ReqShopGroupPage
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	var Form serverShop.ReqGoodsGroupObj

	if err := app.UnmarshalJson(form, &Form); err != nil {

		comm.ApiResponse(ctx,nil,err,0,true)
	}

	coId, _ := middlewares.GetFastCoIdUid(ctx)

	Form.Req.Coid = coId

	res, err := client.ServerShopSys().GroupListByCoidPage(Form)
	comm.ApiResponse(ctx,&res,err,0,true)

}

// @Tags co_shop_group
// @Summary 商品分组删除
// @description 商品分组删除接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.Id true "商品分组ID"
// @Success 200 {object} comm.Response
// @Router /group/delete [delete]
func GroupDelete(ctx *gin.Context) {

	var form request.Id
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	//idStr := ctx.GetHeader("coid")
	//coId := tools.StringToInt64(idStr)
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	res, err := client.ServerShopSys().GroupDelete(serverShop.ShopGroup{Id: form.ID, Coid: coId})
	comm.ApiResponse(ctx,res.Count,err,0,true)
}

// @Tags co_shop_group
// @Summary 获取商品分组详情
// @description 获取商品分组详情
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.Id true "ID"
// @Success 200 {object} comm.Response
// @Router /group/detail [POST]
func ShopGroupDetail(ctx *gin.Context) {
	var form request.Id
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	res, err := client.ServerShopSys().ShopGroupDetail(serverShop.ReqId{Id: form.ID})
	comm.ApiResponse(ctx,&res,err,0,true)
}

// @Tags co_shop_stock
// @Summary 商品库存新增
// @description 商品库存新增接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.ShopStock  true "data"
// @Success 200 {object} comm.Response
// @Router /stock/add [put]
func StockAdd(ctx *gin.Context) {
	var form request.ShopStock
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}

	var Form serverShop.ShopStock
	if err := app.UnmarshalJson(form, &Form); err != nil {
		comm.ApiResponse(ctx,nil,err,0,true)
		return
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	Form.Coid = coId

	res, err := client.ServerShopSys().StockAdd(Form)
	comm.ApiResponse(ctx,&res,err,0,true)

}

// @Tags co_shop_stock
// @Summary 商品库存修改
// @description 商品库存修改接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.ShopStock  true "data"
// @Success 200 {object} comm.Response
// @Router /stock/edit [put]
func StockEdit(ctx *gin.Context) {
	var form request.ShopStock
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	var Form serverShop.ShopStock
	if err := app.UnmarshalJson(form, &Form); err != nil {
		comm.ApiResponse(ctx,nil,err,0,true)
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	Form.Coid = coId
	res, err := client.ServerShopSys().StockEdit(Form)
	comm.ApiResponse(ctx,&res,err,0,true)

}

// @Tags co_shop_stock
// @Summary 商品库存删除
// @description 商品库存删除接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Success 200 {object} comm.Response
// @Router /stock/delete [delete]
func StockDelete(ctx *gin.Context) {
	var form request.Id
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	res, err := client.ServerShopSys().StockDelete(serverShop.ShopStock{Id: form.ID})
	comm.ApiResponse(ctx,res,err,0,true)
}

// @Tags co_shop_coupon
// @Summary 优惠券添加测试
// @description 优惠券新增接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.ShopCoupon  true "data"
// @Success 200 {object} comm.Response
// @Router /coupon/add [put]
func CouponAdd(ctx *gin.Context) {
	CoCouponSave(ctx, "add")
	return
}

// @Tags co_shop_coupon
// @Summary 优惠券修改测试
// @description 优惠券修改组接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.ShopCoupon  true "data"
// @Success 200 {object} comm.Response
// @Router /coupon/edit [put]
func CouponEdit(ctx *gin.Context) {
	CoCouponSave(ctx, "edit")
	return
}

// 公用添加 修改优惠券操作
func CoCouponSave(ctx *gin.Context, way string) {
	var form request.ShopCoupon
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	// 验证参数
	var err error
	if form.UseType != 1 { // 适用商品 1全部商品可用 2指定商品可用 3指定商品不可用
		if len(form.CouponData.UseGoods) == 0 {
			err = app.Err("请输入适用指定商品")
		}
	}
	if form.LimitType == 2 { //领取人限制
		if len(form.CouponData.LimitBenefitMember) == 0 && len(form.CouponData.LimitLevelMember) == 0 {
			err = app.Err("请输入合适限制等级")
		}
	}
	if way == "edit" {
		if err := validator.Get().Var(form.ID, "required"); err != nil {
			comm.ApiResponse(ctx,nil, app.Err("ID值异常"),0,true)
		}
	}
	var Form serverShop.ShopCoupon
	if err := app.UnmarshalJson(form, &Form); err != nil {
		comm.ApiResponse(ctx,nil,err,0,true)
		return
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	Form.Coid = coId
	res, err := client.ServerShopSys().CoCouponSave(Form)
	comm.ApiResponse(ctx,&res,err,0,true)
}

// @Tags co_shop_coupon
// @Summary 优惠券删除
// @description 优惠券删除接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.Id true "优惠劵ID"
// @Success 200 {object} comm.Response
// @Router /coupon/delete [delete]
func CouponDelete(ctx *gin.Context) {
	var form request.Id
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	res, err := client.ServerShopSys().CouponDelete(serverShop.ShopCoupon{Id: form.ID})
	comm.ApiResponse(ctx,res,err,0,true)
}

// @Tags co_shop_coupon
// @Summary 优惠券分页列表
// @description 优惠券分页接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.ReqCouponPage  true "data"
// @Success 200 {object} comm.Response
// @Router /coupon/page [post]
func CouponListPage(ctx *gin.Context) {
	var form request.ReqCouponPage
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	var Form serverShop.ReqCouponPage
	if err := app.UnmarshalJson(form, &Form); err != nil {
		comm.ApiResponse(ctx,nil,err,0,true)
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	Form.Req.Coid = coId
	res, err := client.ServerShopSys().CouponListPage(Form)
	comm.ApiResponse(ctx,&res,err,0,true)
}

// @Tags co_shop_couponCard
// @Summary 优惠券兑换卡新增
// @description 优惠券兑换卡接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.ShopCouponCard  true "data"
// @Success 200 {object} comm.Response
// @Router /coupon/card/add [put]
func CouponCardAdd(ctx *gin.Context) {
	var form request.ShopCouponCard
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	var Form serverShop.ShopCouponCard
	if err := app.UnmarshalJson(form, &Form); err != nil {
		comm.ApiResponse(ctx,nil,err,0,true)
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	Form.Coid = coId
	res, err := client.ServerShopSys().CouponCardAdd(Form)
	comm.ApiResponse(ctx,&res,err,0,true)

}

// @Tags co_shop_couponCard
// @Summary 优惠券兑换卡修改
// @description 优惠券兑换卡修改接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.ShopCouponCard  true "data"
// @Success 200 {object} comm.Response
// @Router /coupon/card/edit [put]
func CouponCardEdit(ctx *gin.Context) {
	var form request.ShopCouponCard
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	var Form serverShop.ShopCouponCard
	if err := app.UnmarshalJson(form, &Form); err != nil {
		comm.ApiResponse(ctx,nil,err,0,true)
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	Form.Coid = coId
	res, err := client.ServerShopSys().CouponCardEdit(Form)
	comm.ApiResponse(ctx,&res,err,0,true)
}

// @Tags co_shop_couponCard
// @Summary 优惠券兑换卡删除
// @description 优惠券兑换卡删除接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param coid header string true "企业ID"
// @Param data body request.Id true "商品ID"
// @Success 200 {object} comm.Response
// @Router /coupon/card/delete [delete]
func CouponCardDelete(ctx *gin.Context) {
	var form request.Id
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	coId, _ := middlewares.GetFastCoIdUid(ctx)
	res, err := client.ServerShopSys().CouponCardDelete(serverShop.ShopCouponCard{Id: form.ID, Coid: coId})
	comm.ApiResponse(ctx,res,err,0,true)
}

// @Tags co_shop_stock
// @Summary 获取库存卡详情
// @description 获取库存卡详情
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param data body request.Id true "ID"
// @Success 200 {object} comm.Response
// @Router /stock/detail [POST]
func StockDetail(ctx *gin.Context) {
	var form request.Id
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	res, err := client.ServerShopSys().StockDetail(serverShop.ReqId{Id: form.ID})
	comm.ApiResponse(ctx,&res,err,0,true)
}

// @Tags co_shop_stockCard
// @Summary 查询库存卡
// @description 查询库存卡
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param data body request.Id true "ID"
// @Success 200 {object} comm.Response
// @Router /stock/card/detail [POST]
func StockCardDetail(ctx *gin.Context) {
	var form request.Id
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	res, err := client.ServerShopSys().StockCardDetail(serverShop.ReqId{Id: form.ID})
	comm.ApiResponse(ctx,&res,err,0,true)
}

// @Tags co_shop_coupon
// @Summary 查询优惠券
// @description 查询优惠券
// @accept json
// @Produce json
// @Param data body request.Id true "ID"
// @Success 200 {object} comm.Response
// @Router /coupon/detail [POST]
func CouponDetail(ctx *gin.Context) {
	var form request.Id
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	res, err := client.ServerShopSys().CouponDetail(serverShop.ReqId{Id: form.ID})
	comm.ApiResponse(ctx,&res,err,0,true)
}

// @Tags co_shop_couponCard
// @Summary 查询优惠券兑换卡
// @description 查询优惠券兑换卡
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param data body request.Id true "ID"
// @Success 200 {object} comm.Response
// @Router /coupon/card/detail [POST]
func CouponCardDetail(ctx *gin.Context) {
	var form request.Id
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	res, err := client.ServerShopSys().CouponCardDetail(serverShop.ReqId{Id: form.ID})
	comm.ApiResponse(ctx,&res,err,0,true)

}

// @Tags co_shop_stockCard
// @Summary 添加库存卡
// @description 添加库存卡
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param data body request.ShopStockCard  true "data"
// @Success 200 {object} comm.Response
// @Router /stock/card/add [put]
func StockCardAdd(ctx *gin.Context) {
	var form request.ShopStockCard
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	var Form serverShop.ShopStockCard
	if err := app.UnmarshalJson(form, &Form); err != nil {
		comm.ApiResponse(ctx,nil,err,0,true)
	}
	//idStr :=ctx.GetHeader("id")
	//Form.Id = tools.StringToInt64(idStr)
	res, err := client.ServerShopSys().StockCardAdd(Form)
	comm.ApiResponse(ctx,&res,err,0,true)
}

// @Tags co_shop_stockCard
// @Summary 修改库存卡
// @description 修改库存卡
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param data body request.ShopStockCard  true "data"
// @Success 200 {object} comm.Response
// @Router /stock/card/edit [put]
func StockCardEdit(ctx *gin.Context) {
	var form request.ShopStockCard
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	var Form serverShop.ShopStockCard
	if err := app.UnmarshalJson(form, &Form); err != nil {
		comm.ApiResponse(ctx,nil,err,0,true)
	}
	//idStr :=ctx.GetHeader("id")
	//Form.Id = tools.StringToInt64(idStr)
	res, err := client.ServerShopSys().StockCardEdit(Form)
	comm.ApiResponse(ctx,&res,err,0,true)
}

// @Tags co_shop_stockCard
// @Summary 删除库存卡
// @description 删除库存卡接口
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param data body request.Id true "ID"
// @Success 200 {object} comm.Response
// @Router /card/delete [delete]
func StockCardDelete(ctx *gin.Context) {
	var form request.Id
	if err := validator.Bind(ctx, &form); !err.IsValid() {
		comm.ValidatorResponse(ctx, err)
		return
	}
	res, err := client.ServerShopSys().StockCardDelete(serverShop.ShopStockCard{Id: form.ID})
	comm.ApiResponse(ctx,res,err,0,true)
}

func FakerBuy(ctx *gin.Context){

	_, err := client.ServerShopSys().FakerBuy(serverShop.Faker{Id: 1})
	app.Println(err)
}
//Router.POST("goods/refresh", shop.QueryGoodsByIds)
