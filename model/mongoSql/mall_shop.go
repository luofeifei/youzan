package mongoSql

// 完整的商品 一个主商品 对应不同规则的子商品
//type CoShopGoods struct {
//	ID         int64          `bson:"_id" json:"id"`                  // 商品 ID
//	Coid       int64          `bson:"coid" json:"coid"`               // 企业ID
//	GoodsAsset CoGoodsAsset   `bson:"goods_asset" json:"goods_asset"` // 无规则情况下的商品库存
//	GoodsRule  []CoGoodsRuleS `bson:"goods_rule" json:"goods_rule" `  // 商品规则 有规则情况下 商品库存失效
//	//Member Member     `bson:"goods_discount" json:"goods_discount"` // 会员折扣
//}
//type GoodsCount struct {
//	LevelOrBenefit int32 `bson:"level_or_benefit" json:"level_or_benefit"`  // 商品规则ID 不同等级的会员，享受不同的会员价优惠。当客户同时拥有等级和多张权益卡的时候，会选出会员折扣最低的卡或者等级，优先计算这张卡或者等级的会员价，如果这张卡或者等级没有自定义会员价，仍可以支持会员折扣。
//	DiscountType    int32 `bson:"discount_type" json:"discount_type"`       // 优惠方式
//	DiscountValue []float64 `bson:"goods_discount" json:"goods_discount"`   // 优惠值 打7.5折 减5.5元 指定100.5元 一个商品对应多个值
//	GoodsID     int64 `bson:"goods_id" json:"goods_id"`           // 商品ID
//}
type Member struct {
	GoodsId      int64          `json:"sku_id"`
	DiscountType []DiscountType `json:"discount_type"`
}
type MemberPrice struct {
	GoodsId      int64          `json:"goods_id"`
	SkuId        int64          `json:"sku_id"`
	DiscountType []DiscountType `json:"discount_type"`
}
type DiscountType struct {
	LevelOrBenefit int8      `binding:"oneof=1 2" json:"level_or_benefit"`  //1 按会员等级设置 2 按权益卡设置
	DiscountMethod int8      `binding:"oneof=1 2 3" json:"discount_method"` // 优惠方式：1打折 2减价 3指定价格
	DiscountValue  []float64 `json:"discount_value"`                        // 优惠值 打7.5折 减5.5元 指定100.5元 一个商品对应多个值
}

// 商品规则 颜色-尺码-性别 对应白色下的所有规则 白色-50码-男
//type CoGoodsRuleS struct {
//	Pic   string              `bson:"pic" json:"pic"`     // 规格图片 颜色为白色对应的图片
//	Type  int64               `bson:"type" json:"type"`   // 规则项目类型 1文本 2数字 3时间 4身份证
//	Name  string              `bson:"name" json:"name"`   // 规则名称 值
//	Alias []CoGoodsRuleSAlias `bson:"alias" json:"alias"` // 下级规则
//	Member     [] Member    `bson:"member" json:"member"`    // 会员价设置
//}
//type CoGoodsRuleSAlias struct {
//	Id         int64               `bson:"_id" json:"id"`                // 规则ID 前台或者生成短ID
//	Name       string              `bson:"name" json:"name"`             // 规则名称
//	Type       int64               `bson:"type" json:"type"`             // 规则类型  1文本 2数字 3时间 4身份证
//	StockID    int64               `bson:"stock_id" json:"stock_id"`     // 关联库存ID  关联特殊功能值 比如优惠券ID
//	Price      float64             `bson:"price" json:"price"`           // 商品价格
//	PriceDot   float64             `bson:"price_dot" json:"price_dot"`   // 划线价
//	PriceCost  float64             `bson:"price_cost" json:"price_cost"` // 成本价
//	GoodsNo    string              `bson:"goods_no" json:"goods_no"`     // 商品编码
//	GoodsAsset CoGoodsAsset        `bson:"goods_asset" json:"goods_asset"` // 商品库存
//	Member     [] Member    `bson:"member" json:"member"`    // 会员价设置
//	Alias      []CoGoodsRuleSAlias `bson:"alias" json:"alias"` // 下级规则
//}

// 商品资产库 用于查询更新库存
//type CoGoodsAsset struct {
//	GoodsID     int64 `bson:"goods_id" json:"goods_id"`           // 商品ID
//	GoodsRuleID int64 `bson:"goods_rule_id" json:"goods_rule_id"` // 商品规则ID
//	StockNum    int64 `bson:"stock_num" json:"stock_num" binding:"required"`       // 库存数量
//	Sales       int64 `bson:"sales" json:"sales"`                 // 销量
//}
//func (m *CoShopGoods) TableName() string {
//	return "co_shop_goods"
//}

// start 新的存储方式
type GoodsStockDisCount struct {
	ID            int64       `bson:"_id" json:"id"`                                    // 商品 ID
	Coid          int64       `bson:"coid" json:"coid"`                                 // 企业ID
	GoodsRuleList []GoodsRule `bson:"goods_rule_list" json:"goods_rule_list,omitempty"` // 库存和打折信息
	//GoodsAsset GoodsAsset  `bson:"goods_asset" json:"goods_asset"`    // 商品总库存
}
type GoodsDistributionStock struct {
	ID       int64      `bson:"_id" json:"id"`              // 商品ID
	Coid     int64      `bson:"coid" json:"coid"`           // 企业ID
	StockNum int64      `bson:"stock_num" json:"stock_num"` // 总库存数量
	SoldNum  int64      `bson:"sold_num" json:"sold_num"`   // 总销量
	SkuStock []SkuStock `bson:"sku_stock" json:"sku_stock"` // sku库存
}

type SkuStock struct {
	SkuId    int64 `bson:"sku_id" json:"sku_id"`       // sku ID
	StockNum int64 `bson:"stock_num" json:"stock_num"` // 库存数量
	SoldNum  int64 `bson:"sold_num" json:"sold_num"`   // 销量
}

func (m *GoodsDistributionStock) TableName() string {
	return "shop_goods_distribution_stock"
}

type GoodsAsset struct {
	GoodsID  int64 `bson:"goods_id" json:"goods_id"`   // 商品ID
	StockNum int64 `bson:"stock_num" json:"stock_num"` // 库存数量
	SoldNum  int64 `bson:"sold_num" json:"sold_num"`   // 销量
}

func (m *GoodsStockDisCount) TableName() string {
	return "shop_goods_stock_discount"
}

type GoodsRule struct {
	SkuId             int64           `bson:"sku_id" json:"sku_id"`                                     // sku ID
	GoodsStock        GoodsStock      `bson:"goods_stock" json:"goods_stock,omitempty"`                 // 库存信息
	GoodsDiscountList []GoodsDiscount `bson:"goods_discount_list" json:"goods_discount_list,omitempty"` // 打折信息
	GoodsRuleInfo     []GoodsRuleInfo `bson:"goods_rule_info" json:"goods_rule_info,omitempty"`         // 打折信息
}
type GoodsStock struct {
	//ID         int64          `bson:"_id" json:"id"`                  // ID
	StockID   int64   `bson:"stock_id" json:"stock_id"`     // 关联库存ID  关联特殊功能值 比如优惠券ID
	Price     float64 `bson:"price" json:"price"`           // 商品价格
	PriceDot  float64 `bson:"price_dot" json:"price_dot"`   // 划线价
	PriceCost float64 `bson:"price_cost" json:"price_cost"` // 成本价
	GoodsNo   string  `bson:"goods_no" json:"goods_no"`     // 商品编码
	//StockNum    int64 `bson:"stock_num" json:"stock_num"`  // 库存数量
	//SoldNum       int64 `bson:"sold_num" json:"sold_num"`  // 销量
}
type GoodsDiscount struct {
	ID            int64 `bson:"_id" json:"id"`                        // GOODS_ID
	SkuId         int64 `bson:"sku_id" json:"sku_id"`                 // sku_id
	BindType      int32 `bson:"bind_type" json:"bind_type"`           // 1 会员卡 2 权益卡
	DiscountType  int32 `bson:"discount_type" json:"discount_type"`   // 优惠方式 1打折 2减价 3指定价格
	DiscountValue int32 `bson:"discount_value" json:"discount_value"` // 打折
}
type GoodsRuleInfo struct {

	Name  string `bson:"name" json:"name"`   // 规则名称
	Type  int64  `bson:"type" json:"type"`   // 规则类型  1文本 2数字 3时间 4身份证
	Value string `bson:"value" json:"value"` // 规则值
}

// end 新的存储方式
