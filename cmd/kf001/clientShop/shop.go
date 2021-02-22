package main

type ShopGoods struct {
	ID         int64                   `json:"id"`                                               // 商品 ID
	Type       int8                    `binding:"required,oneof=1 2 3 4" json:"type"`            // 商品类型 1 虚拟物品 2实物物品 3电子卡券 4付费优惠券
	Name       string                  `binding:"required,max=50" json:"name" maxLength:"50"`    // 商品名称 50个字符以内
	Title      string                  `binding:"required,max=100" json:"title" maxLength:"100"` // 商品标题 100个字符以内
	GroupID    int64                   `binding:"required" json:"group_id"`                      // 所属分组ID
	SnapshotId int64                   `json:"group_id"`                                         // 快照ID 保存前做售出检查 如果数据基本信息变化 做快照备份
	RetailId   int64                   `binding:"required" json:"group_id"`                      // 分销商品ID
	Cover      string                  `binding:"required" json:"cover"`                         // 封面图片 JSON数组对应图片ID 最多上传15张
	StockID    int64                   `json:"stock_id"`                                         // 虚拟商品关联库存id
	StockType  int8                    `binding:"required,oneof=1 2"  json:"stock_type"`         // 库存扣减方式 1拍下减库存 2付款减库存
	IsVip      int8                    `binding:"oneof=1 2" json:"is_vip"`                       // 是否参加会员折扣 1是 2不是
	IsPoints   int8                    `binding:"oneof=1 2" json:"is_points"  `                  // 是否可使用积分购买 1是 2不是
	MaxPoints  int64                   `json:"max_points"`                                       // 最大可使用积分数
	PriceDot   int32                   `json:"price_dot" form:"price_dot"`                       // 划线价 以分为单位 不存在读 价格规则内最低值
	IsShowRest int8                    `binding:"required,oneof=1 2" json:"is_show_rest"`        // 商品详情、购物车是否显示剩余件数 1 显示 2不显示
	SaleTime   int64                   `validate:"datetime" json:"sale_time"`                    // 开售时间  -1 放入仓库 改变状态
	AfterSale  ShopGoodsAfterSale      `json:"after_sale,omitempty"`                             // 售后服务规则
	PayLimit   ShopGoodsPayLimit       `json:"pay_limit,omitempty"`                              // 购买限制规则
	GoodsData  ShopGoodsData           `json:"goods_data"`                                       // 商品详情数据
	GoodsRule  map[int16]GoodsRuleList `json:"goods_rule,omitempty"`                             // 商品价格库存规则 有规则情况下 商品库存失效  0 为总库存信息 其他自定义下 int16 为前端生成规则短ID
	State      int8                    `binding:"oneof=1 2 3 4" json:"state" `                   // 状态 1待发布 2正常 3下架 4 已售罄 5仓库中
}

// 售后服务规则 mysql
type ShopGoodsAfterSale struct {
	// 实物
	BuyerExchange      bool  `json:"buyer_exchange,omitempty"`       // 支持买家申请换货
	BuyerSevenDays     bool  `json:"buyer_exchange,omitempty"`       // 7天无理由退货
	BuyerSpeedExchange bool  `json:"buyer_speed_exchange,omitempty"` // 极速退款
	BuyerRefund        bool  `json:"buyer_refund,omitempty"`         // 支持买家申请退款
	BuyerRefundExpire  int16 `json:"buyer_refund_expire,omitempty"`  // 未核销卡券过期多久前可退款 -1 未核销均支持退款 单位小时
}

// 商品购买限制规则 mysql
type ShopGoodsPayLimit struct {
	LimitPayMin      int8                        `binding:"max=9999" json:"limit_pay_min"`          // 每单最低购买数量
	LimitPayQuota    int8                        `binding:"oneof=1 2 3 4 5" json:"limit_pay_quota"` // 限购 1每单 2终身 3每天 4每周 5每月 6每年
	LimitPayQuotaNum int16                       `json:"limit_pay_quota_num"`                       // 购买限制单数
	LimitSpecific    []ShopGoodsPayLimitSpecific `json:"limit_specific"`                            // 指定用户可购买 根据 商品类型确定是否有改选项
}

// 指定限制规则
type ShopGoodsPayLimitSpecific struct {
	BindType int8    `json:"bind_type"` // 1 会员卡 2 权益卡 3用户标签
	BindId   []int64 `json:"bind_id"`   // 绑定的相关ID列表
}

type GoodsRuleList struct {
	GoodsPriceName    GoodsPriceName      `json:"goods_price_name,omitempty"`  // 库存基本信息
	GoodsPriceStock   GoodsPriceStock     `json:"goods_price_stock,omitempty"` // 库存数量信息
	GoodsDiscount     []GoodsDiscount     `json:"goods_discount,omitempty"`    // 打折信息
	GoodsRuleListInfo []GoodsRuleListInfo `json:"goods_info,omitempty"`        // 规则信息
}

// 商品价格规则 名称
type GoodsPriceName struct {
	Price     int32  `json:"price"`                   // 商品价格 以分为单位
	PriceDot  int32  `json:"price_dot"`               // 划线价 以分为单位
	PriceCost int32  `json:"price_cost"`              // 成本价 以分为单位
	GoodsNo   string `json:"goods_no" maxLength:"20"` // 商品编码 20个字符以内
}

// 商品价格规则 储存至 mongodb
type GoodsPriceStock struct {
	StockID  int64 `json:"stock_id"`  // 关联库存ID 关联特殊功能值 比如优惠券ID
	StockNum int64 `json:"stock_num"` // 库存数量
	SoldNum  int64 `json:"sold_num"`  // 销量
}

type GoodsDiscount struct {
	BindType      int64 `json:"bind_type"`      // 1 会员卡 2 权益卡
	DiscountType  int64 `json:"discount_type"`  // 优惠方式 1打折 2减价 3指定价格
	DiscountValue int64 `json:"discount_value"` // 打折 打折7.5折传75 加给以分为单位
}

// API 接口 实现
type GoodsRuleListInfo struct {
	Name  string `json:"name"`  // 规则名称
	Type  int64  `json:"type"`  // 规则类型  1文本 2数字 3时间 4身份证
	Value string `json:"value"` // 规则值
}

// 商品详情规则储存到 mysql
type ShopGoodsData struct {
	Video        string                    `json:"video"`                                      // 主图视频  35个字符长度限制
	Pic          []string                  `binding:"required,max=15,dive,max=255" json:"pic"` //商品图 JSON数组对应图片ID 最多上传15张
	Share        string                    `binding:"max=36" json:"share"`                     // 分享描述 微信分享给好友时会显示，建议36个字以内
	SubName      string                    `binding:"max=20" json:"sub_name"`                  // 购买按钮名称 20个字以内
	SpecialRules ShopGoodsDataSpecialRules `json:"special_rules,omitempty"`                    // 特殊规则、自定义商品规则 根据商品类型确定
	Word         []ShopGoodsDataWords      `json:"word"`                                       // 用户留言数组 {名称,模型,必填} JSON数组 如 电话,phone,1 购买商品时让买家输入留言，最多可设置 10 条留言
	Content      string                    `json:"content"`                                    // 商品详情
}

// 商品留言信息规则 mysql
type ShopGoodsDataWords struct {
	Name       string `json:"name"`                             // 用户留言字段名称 phone
	Type       int8   `binding:"oneof=1 2 3 4" json:"type"`     // 用户留言字段对应的值  1文本 2数字 3时间 4身份证
	IsMultiRow int8   `binding:"oneof=1 2" json:"is_multi_row"` // 是否多行 1 是多行 2 不是多行
	IsRequired int8   `binding:"oneof=1 2" json:"is_required"`  // 是否必填 1 是必填 2 不是必填
	IsSingle   int8   `binding:"oneof=1 2" json:"is_single"`    // 1 是 (只需要用户填1次) 2否 (根据购买数量填写,买N张门票需要填写N个身份证)
}

// 自定义规则
type ShopGoodsDataSpecialRules struct {
	GoodsReal ShopGoodsDataSpecialRulesGoodsReal `json:"goods_real,omitempty"` // 实物商品自定义规则
	GoodsCake ShopGoodsDataSpecialRulesGoodsCake `json:"goods_cake,omitempty"` // 蛋糕烘焙
	GoodsCard ShopGoodsDataSpecialRulesGoodsCard `json:"goods_card,omitempty"` // 电子卡券规则
}

// 自定义规则 实物商品
type ShopGoodsDataSpecialRulesGoodsReal struct {
	DeliveryType []int8 `json:"delivery_type"`                     // 配送方式 1 快递发货 2 同城配送 3 到店自提
	FreightPrice int32  `binding:"oneof=1 2" json:"freight_price"` // 运费价格 统一价格 -1 为调用运费模板
	FreightTpl   int32  `json:"freight_tpl"`                       // 运费模板
}

// 自定义规则 电子卡券规则
type ShopGoodsDataSpecialRulesGoodsCard struct {
	ValidityMode        int8   `json:"validity_mode"`                             // 卡券生效模式 1 立即生效 2 次日生效 3多少小时后生效
	ValidityType        int8   `json:"validity_type"`                             // 卡券生效类型 (1立即生效、 -1长期有效 1指定天数 2指定时间段)
	ValidityDay         int8   `json:"validity_day"`                              // 卡券生效时长 (生效类型 如：5 指5天内有效)
	ValidityStartTime   int64  `validate:"datetime" json:"validity_start_time"`   // 具体限制卡券 有效期开始时间
	ValidityEndTime     int64  `validate:"datetime" json:"validity_end_time"`     // 具体限制卡券 有效期结束时间
	IsHolidaysAvailable int8   `binding:"oneof=1 2" json:"is_holidays_available"` // 节假日是否可用 1是 2不是
	UseNotes            string `json:"use_notes"`                                 // 使用说明
}

// 自定义规则 蛋糕烘焙
type ShopGoodsDataSpecialRulesGoodsCake struct {
	AttributeIds []int64         `json:"attribute_ids"` // 商品附加属性值 关联到 商品属性库ID
	StockUpTime  int64           `json:"stock_up_time"` // 统一备货时间 单位分钟 -1 启用不同规格单独设置备货时间
	StockUp      map[int16]int64 `json:"attribute"`     // 备货时间 map[SKUID]时间 单位分钟
}
