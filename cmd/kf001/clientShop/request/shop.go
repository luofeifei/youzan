package request

type ShopGroup struct {
	ID   int64  `json:"id"`              // 分组id
	Name string `json:"name" example:""` // 商铺分组名称
}
type ReqShopGroupPage struct {
	PageInfo  PageInfo  `json:"page"`
	ShopGroup ShopGroup `json:"req"` // 企业ID
}
type ReqCouponPage struct {
	PageInfo      PageInfo      `json:"page"`
	ReqShopCoupon ReqShopCoupon `json:"req"` // 查询用的结构体
}
type ReqShopCoupon struct {
	Name  string `binding:"max=50" json:"name" form:"name"` // 优惠券名称 50个字符
	State int8   `binding:"oneof=-1 1 2" json:"state"`      // 状态-1 全部 1正常 2失效
	Type  int8   `binding:"oneof=-1 1 2 3 4" json:"type"`   // 优惠券类型-1全部 1满减券 2折扣券 3随机金额券 4商品兑换券
}

type ReqGoodsGroupPageCoid struct {
	Page     int32  `binding:"required" json:"page" form:"page"  example:"1"`          // 当前页
	PageSize int32  `binding:"required" json:"pageSize" form:"pageSize"  example:"10"` // 每页显示数
	OrderKey string `json:"orderKey" form:"orderKey"`                                  // 默认排序字段 -filed1,+field2,field3 (-Desc 降序)
}

type ReqGoodsPage struct {
	Page     int32  `binding:"required" json:"page" form:"page"`          // 当前页
	PageSize int32  `binding:"requ ired" json:"pageSize" form:"pageSize"` // 每页显示数
	OrderKey string `json:"orderKey" form:"orderKey"`                     // 默认排序字段 -filed1,+field2,field3 (-Desc 降序)
}

type SearchInfo struct {
	Page       PageInfo `json:"page"`
	GroupID    int64    `json:"group_id"`    // 所属分组ID
	Name       string   `json:"name"`        // 商品名称
	Type       int8     `json:"type"`        // 商品类型 所有 1 虚拟物品 2实物物品 3电子卡券 4付费优惠券
	StockType  int8     `json:"stock_type"`  // 库存扣减方式 1拍下减库存 2付款减库存
	PriceStart int32    `json:"price_start"` // 商品价格区间下线
	PriceEnd   int32    `json:"price_end"`   // 商品价格区间上线
	StartTime  int      `json:"start_time"`  // 上架开始时间
	EndTime    int      `json:"end_time"`    // 上架结束时间
	State      int8     `json:"state"`       // 状态 1待发布 2正常 3下架 4停用
}

//type ShopGoods struct {
//	Id        int64  `binding:"required" json:"id" form:"id"` // 商品ID
//	GroupID   int64  `json:"group_id" form:"group_id"`        // 所属分组ID
//	Name      string `json:"name" form:"name"`                // 商品名称
//	Title     string `json:"title" form:"title"`              // 商品标题
//	Type      int8   `json:"type" form:"type"`                // 商品类型 1 虚拟物品 2实物物品 3电子卡券 4付费优惠券
//	Cover     string `json:"cover" form:"cover"`              // 封面图片
//	StockID   int64  `json:"stock_id" form:"stock_id"`        // 关联库存ID
//	StockType int8   `json:"stock_type" form:"stock_type"`    // 库存扣减方式 1拍下减库存 2付款减库存
//	StockNum  int    `json:"stock_num" form:"stock_num"`      // 库存数量
//	IsVip     int8   `json:"is_vip" form:"is_vip"`            // 是否参加会员折扣
//	IsPoints  int8   `json:"is_points" form:"is_points"`      // 是否可使用积分购买
//	MaxPoints int    `json:"max_points" form:"max_points"`    // 最大可使用积分数
//	Price     int32  `json:"price" form:"price"`              // 商品价格 以分为单位
//	PriceDot  int32  `json:"price_dot" form:"price_dot"`      // 划线价 以分为单位
//	PriceCost int32  `json:"price_cost" form:"price_cost"`    // 成本价 以分为单位
//	GoodsNo   string `json:"goods_no" form:"goods_no"`        // 商品编码
//	SaleMin   int8   `json:"sale_min" form:"sale_min"`        // 每单最低购买数量
//	SaleMax   int8   `json:"sale_max" form:"sale_max"`        // 每单最多购买数量
//	Quota     int8   `json:"quota" form:"quota"`              // 限购 1终身 2每天 3每周 4每月 5每年
//	QuotaNum  int    `json:"quota_num" form:"quota_num"`      // 购买限制单数
//	StartTime int    `json:"start_time" form:"start_time"`    // 上架时间
//	EndTime   int    `json:"end_time" form:"end_time"`        // 下架时间
//	State     int8   `json:"state" form:"state"`              // 状态 1待发布 2正常 3下架 4停用
//}

type ShopCoupon struct {
	ID               int64          `json:"id" form:"id"`
	Type             int8           `binding:"required,oneof=1 2 3 4" json:"type"`                     // 优惠券类型 1满减券 2折扣券 3随机金额券 4商品兑换券
	Name             string         `binding:"required,max=50" json:"name" form:"name" maxLength:"50"` // 优惠券名称 50个字符
	Remark           string         `binding:"max=255" json:"remark" form:"remark" maxLength:"255"`    // 名称备注
	UseStockQty      int            `binding:"required" json:"use_stock_qty" form:"use_stock_qty"`     // 发放总量
	UsedQty          int            `json:"used_qty"`                                                  // 当前已使用量
	UseType          int8           `binding:"required,oneof=1 2 3" json:"use_type"`                   // 适用商品 1全部商品可用 2指定商品可用 3指定商品不可用
	UsingAmount      float64        `binding:"required" json:"using_amount"`                           // 使用门槛 -1无使用门槛 订单满多少元
	UseMode          int8           `binding:"required,oneof=1 2 3" json:"use_mode"`                   // 用券模式 1指定时间 2领券当日起 3领券次日起 若设置固定用券时间，编辑保存后对已领取未使用及后续领取的券均生效。若设置领券当日/次日n天内可用，编辑保存后仅对后续领取的券生效。
	UseDay           int8           `binding:"max=365" json:"use_day"`                                 // 用券模式 2、3时指定天数
	IsOverlay        int8           `binding:"oneof=1 2" json:"is_overlay"`                            // 优惠叠加 是否优惠券仅原价购买时可用 1是 2不是
	IsPublic         int8           `binding:"oneof=1 2" json:"is_public"`                             // 允许公开领取 1不允许 2允许
	IsSharable       int8           `binding:"oneof=1 2" json:"is_sharable"`                           // 是否可共享 1不可共享 2可共享
	IsHandSel        int8           `binding:"oneof=1 2" json:"is_hand_sel"`                           // 转赠设置 1不可 2允许转赠给好友
	IsExpireNotice   int8           `json:"is_expire_notice"`                                          // 过期提醒
	ExpireNoticeDays int8           `json:"expire_notice_days"`                                        // 过期前几天提醒
	LimitType        int8           `binding:"required,oneof=1 2" json:"limit_type"`                   // 领取人限制 1不限制 2限制
	LimitNum         int8           `binding:"required" json:"limit_num"`                              // 每人限领次数 -1不限次数 大于1领取次数
	StartTime        int            `json:"start_time"`                                                // 用券开始时间
	EndTime          int            `json:"end_time"`                                                  // 用券结束时间
	State            int8           `binding:"oneof=1 2" json:"state"`                                 // 状态 1正常 2失效
	CouponData       ShopCouponData `json:"coupon_data"`                                               // 优惠券规则详情（限定领取等级和使用商品）
}

type ShopCouponData struct {
	ID                 int64   `json:"id"`                                                               // 优惠券ID
	UseGoods           []int64 `binding:"max=255" json:"use_goods" maxLength:"255"`                      // 适用商品 商品ID数组
	UseRule            string  `binding:"required" json:"use_rule"`                                      // 使用规则数据 1减免?元 2打?折最多优惠?元 3随机?至?元
	LimitBenefitMember []int64 `binding:"max=255" json:"limit_benefit_member" maxLength:"255"`           // 领取人身份等级数组 客户身份权益卡 权益卡(权益卡1，权益卡2) [10001,10002,10003]
	Description        string  `binding:"max=255" json:"description" form:"description" maxLength:"255"` // 使用说明
	LimitLevelMember   []int64 `binding:"max=255" json:"limit_level_member" maxLength:"255"`             //领取人身份等级数组 客户身份会员等级 会员(黄金会员,白金会员) [10001,10002,10003]
}

type ShopCouponCard struct {
	ID        int64 `json:"id" form:"id"`                 // 兑换码
	Uid       int64 `json:"uid" form:"uid"`               // 所属用户ID
	StartTime int   `json:"start_time" form:"start_time"` // 开始时间
	EndTime   int   `json:"end_time" form:"end_time"`     // 到期时间
	State     int8  `json:"state" form:"state"`           // 状态 1待领取 2待使用 3使用中 3使用完成
}
type ShopStockCard struct {
	ID      int64  `json:"id" form:"id"`             // 库存列表 ID
	StockID int64  `json:"stock_id" form:"stock_id"` // 库存ID
	Card    string `json:"card" form:"card"`         // 卡号、券码
	Pass    string `json:"pass" form:"pass"`         // 密码
	State   int8   `json:"state" form:"state"`       // 状态 1有效 2已发送 3已使用
}

type ShopGoodsSave struct {
	ShopGoodsInfo      ShopGoodsInfo      `json:"info"`
	GoodsStockDisCount GoodsStockDisCount `json:"goods_rule,omitempty"` // 商品规则 有规则情况下 商品库存失效
	GoodsData          GoodsData          `json:"goods_data"`           // 商品详情数据
}
type GoodsAsset struct {
	GoodsID  int64 `bson:"goods_id" json:"goods_id"`   // 商品ID
	StockNum int64 `bson:"stock_num" json:"stock_num"` // 库存数量
	SoldNum  int64 `bson:"sold_num" json:"sold_num"`   // 销量
}
type ShopGoodsInfo struct {
	ID        int64  `json:"id"` // 商品 ID
	ShopId    int64  // 分销商品ID
	GroupID   int64  `binding:"required" json:"group_id"`                      // 所属分组ID
	Name      string `binding:"required,max=50" json:"name" maxLength:"50"`    // 商品名称 50个字符以内
	Title     string `binding:"required,max=100" json:"title" maxLength:"100"` // 商品标题 100个字符以内
	Type      int8   `binding:"required,oneof=1 2 3 4" json:"type"`            // 商品类型 1 虚拟物品 2实物物品 3电子卡券 4付费优惠券
	Cover     string `binding:"required" json:"cover"`                         // 封面图片 JSON数组对应图片ID 最多上传15张
	StockID   int64  `json:"stock_id"`                                         // 虚拟商品关联库存id
	StockType int8   `binding:"required,oneof=1 2"  json:"stock_type"`         // 库存扣减方式 1拍下减库存 2付款减库存
	IsVip     int8   `binding:"oneof=1 2" json:"is_vip"`                       // 是否参加会员折扣 1是 2不是
	IsPoints  int8   `binding:"oneof=1 2" json:"is_points"  `                  // 是否可使用积分购买 1是 2不是
	MaxPoints int    `json:"max_points"`                                       // 最大可使用积分数
	Price     int32  `binding:"required" json:"price" form:"price"`            // 商品价格 以分为单位
	PriceDot  int32  `json:"price_dot" form:"price_dot"`                       // 划线价 以分为单位
	PriceCost int32  `json:"price_cost" form:"price_cost"`                     // 成本价 以分为单位
	SaleMin   int8   `binding:"max=9999" json:"sale_min"`                      // 每单最低购买数量
	SaleMax   int8   `binding:"max=9999" json:"sale_max"`                      // 每单最多购买数量
	Quota     int8   `binding:"oneof=1 2 3 4 5" json:"quota"`                  // 限购 1终身 2每天 3每周 4每月 5每年
	QuotaNum  int    `json:"quota_num"`                                        // 购买限制单数
	//StartTime int    `validate:"datetime" json:"start_time"`                   // 上架时间
	//EndTime   int    `validate:"datetime" json:"end_time"`                     // 下架时间
	State   int8   `binding:"oneof=1 2 3 4" json:"state" `           // 状态 1待发布 2正常 3下架 4停用
	GoodsNo string `binding:"max=20" json:"goods_no" maxLength:"20"` // 商品编码 20个字符以内
}

// 电子卡券特有的属性 有效期和退款
type ElectronicCoupon struct {
	ValidityStartTime   int  `validate:"datetime" json:"validity_start_time"`   // 有效期开始时间
	ValidityEndTime     int  `validate:"datetime" json:"validity_end_time"`     // 有效期结束时间
	IsHolidaysAvailable int8 `binding:"oneof=1 2" json:"is_holidays_available"` // 节假日是否可用  1是 2不是
	IsSupportRefund     int8 `binding:"oneof=1 2" json:"is_support_Refund" `    // 是否支持买家申请退款  1是 2不是
	IsLongTimeSupport   int8 `binding:"oneof=1 2" json:"is_long_time_support" ` // 是否 未核销卡券无论是否过期均支持退款  1是 2不是
	RefundOverdueTime   int  `"json:"refund_overdue_time"`                      // 退款截至时间 未核销卡券在过期前可退款
}

// start 新的存储方式
type GoodsStockDisCount struct {
	ID int64 `json:"id"` // 商品ID
	//Coid          int64          `json:"coid"` // 企业ID
	GoodsAsset    GoodsAsset     `json:"goods_asset"`
	GoodsRuleList []ReqGoodsRule `json:"goods_rule_list,omitempty"` // 库存和打折信息

}
type ReqGoodsRule struct {
	//ID         int64          `json:"id"`                  // ID
	SkuId             int64           `json:"sku_id"`                        // sku ID
	GoodsStock        GoodsStock      `json:"goods_stock,omitempty"`         // 库存信息
	GoodsDiscountList []GoodsDiscount `json:"goods_discount_list,omitempty"` // 打折信息
	GoodsRuleInfo     []GoodsRuleInfo `json:"goods_rule_info,omitempty"`     // 规则信息
}
type GoodsStock struct {
	// 商品ID
	// 属于那个企业的ID
	// 分销ID

	//ID         int64          `json:"id"`                  // ID
	StockID   int64  `json:"stock_id"`   // 关联库存ID  关联特殊功能值 比如优惠券ID
	Price     int32  `json:"price"`      // 商品价格 以分为单位
	PriceDot  int32  `json:"price_dot"`  // 划线价 以分为单位
	PriceCost int32  `json:"price_cost"` // 成本价 以分为单位
	GoodsNo   string `json:"goods_no"`   // 商品编码
	StockNum  int64  `json:"stock_num"`  // 库存数量
	SoldNum   int64  `json:"sold_num"`   // 销量
}

//type GoodsDiscount struct {
//	//ID         int64          `json:"id"`                  // ID
//	BindType      int64 `json:"bind_type"`      // 1 会员卡 2 权益卡
//	DiscountType  int64 `json:"discount_type"`  // 优惠方式 1打折 2减价 3指定价格
//	DiscountValue int64 `json:"discount_value"` // 打折 打折7.5折传75 加给以分为单位
//}
type GoodsRuleInfo struct {
	//Id         int64               `json:"id"`                // 规则ID 前台或者生成短ID
	Name  string `json:"name"`  // 规则名称
	Type  int64  `json:"type"`  // 规则类型  1文本 2数字 3时间 4身份证
	Value string `json:"value"` // 规则值
}

// end 新的存储方式

type GoodsData struct {
	Pic              []string         `binding:"required,max=15,dive,max=255" json:"pic"` //商品图 JSON数组对应图片ID 最多上传15张
	Video            string           `json:"video"`                                      // 主图视频  35个字符长度限制
	Share            string           `binding:"max=36" json:"share"`                     // 分享描述 微信分享给好友时会显示，建议36个字以内
	Word             []Word           `binding:"max=10" json:"word"`                      // 用户留言数组 {名称,模型,必填} JSON数组 如 电话,phone,1 购买商品时让买家输入留言，最多可设置 10 条留言
	SubName          string           `binding:"max=20" json:"sub_name"`                  // 购买按钮名称 20个字以内
	Content          string           `json:"content"`                                    // 商品详情
	SaleTime         int              `validate:"datetime" json:"sale_time"`              // 开售时间
	IsShowRest       int8             `binding:"required,oneof=1 2" json:"is_show_rest"`  // 商品详情、购物车是否显示剩余件数 1 显示 2不显示
	ElectronicCoupon ElectronicCoupon `json:"electronic_coupon,omitempty"`                // 电子卡券特有的属性 电子卡券有效期和退款
}

type Word struct {
	Name       string `json:"name"`                             // 用户留言字段名称 phone
	Type       int8   `binding:"oneof=1 2 3 4" json:"type"`     // 用户留言字段对应的值  1文本 2数字 3时间 4身份证
	IsMultiRow int8   `binding:"oneof=1 2" json:"is_multi_row"` // 是否多行 1 是多行 2 不是多行
	IsRequired int8   `binding:"oneof=1 2" json:"is_required"`  // 是否必填 1 是必填 2 不是必填
	IsSingle   int8   `binding:"oneof=1 2" json:"is_single"`    // 1 是 (只需要用户填1次) 2否 (根据购买数量填写,买N张门票需要填写N个身份证)
}

type ShopStock struct {
	Type    int8   `json:"type" form:"type"`         // 类型 1自有 2会员权益 3优惠券
	Aid     int64  `json:"aid" form:"aid"`           // 2会员权益ID 3优惠券ID
	Name    string `json:"name" form:"name"`         // 库存名称
	Notes   string `json:"notes" form:"notes"`       // 库存备注
	AutoNum int    `json:"auto_num" form:"auto_num"` // 库存不足 自动增加数量
}

//type ShopGoodsData struct {
//	ID      int64  `json:"id" form:"id"`             // 商品ID
//	Pic     string `json:"pic" form:"pic"`           // 商品图  JSON数组对应图片ID
//	Video   string `json:"video" form:"video"`       // 主图视频
//	Share   string `json:"share" form:"share"`       // 分享描述
//	Word    string `json:"word" form:"word"`         // 用户留言数组 {名称,模型,必填} JSON数组 如 电话,phone,1
//	SubName string `json:"sub_name" form:"sub_name"` // 购买按钮名称
//	Content string `json:"content" form:"content"`   // 商品详情
//}

type CoShopGoodsOperate struct {
	Type        int32       `binding:"required" json:"type"`         // 操作方式  1删除  2正常(上架) 3下架 4停用  5分组 6会员价 7商品名称文字替换  8上架时间 9每人限购 10会员折扣 11限定用户(限定用户有购买的权限) 12 配送方式 13商品模板
	GoodsList   []int64     `binding:"required" json:"goods_list"`   // 要操作的商品ID数组
	FeatureList FeatureList `binding:"required" json:"feature_list"` //要修改的的商品属性值
}
type FeatureList struct {
	GroupId    int64           `json:"group_id"`   // 分组id
	StartTime  int64           `json:"start_time"` // 上架时间
	Quota      int8            `json:"quota"`      // 每人限购类型 1终身 2每天 3每周 4每月 5每年
	QuotaNum   int32           `json:"quota_num"`  // 每人购买限制单数
	IsVip      int8            `json:"is_vip"`     //是否参加会员折扣 1参加 2 不参加
	Member     []MemberPrice   `json:"member"`     // 主商品设置会员价 可以设置成按会员和按权益卡两种 所有是个数组
	NameChange GoodsNameChange `json:"goods_name"` // 商品名称修改
}
type MemberPrice struct {
	GoodsId      int64          `json:"goods_id"`
	SkuId        int64          `json:"sku_id"`
	DiscountType []DiscountType `json:"discount_type"`
}
type DiscountType struct {
	LevelOrBenefit int8    `binding:"oneof=1 2" json:"level_or_benefit"`  //1 按会员等级设置 2 按权益卡设置
	DiscountMethod int8    `binding:"oneof=1 2 3" json:"discount_method"` // 优惠方式：1打折 2减价 3指定价格
	DiscountValue  []int32 `json:"discount_value"`                        // 优惠值 打7.5折 减5.5元 指定100.5元 一个商品对应多个值 打折对应的值乘以10 其他乘以100(以分为单位)
}

type GoodsNameChange struct {
	NameList []string `json:"name_list"` // 商品名称完整的数组 跟GoodsList对应
	OldName  string   `json:"old_name"`  // 商品名称里的文字
	NewName  string   `json:"new_name"`  // 商品名称文字替换为新文字 为空表示删除商品名称里的oldName文字
}

type FrontSearchInfo struct {
	PageInfo   PageInfo        `binding:"required" json:"page_info"`
	SearchInfo SearchInfoFront `binding:"required" json:"search_info"`
}
type SearchInfoFront struct {
	GroupID    int64  `json:"group_id"`    // 所属分组ID
	Name       string `json:"name"`        // 商品名称
	Type       int8   `json:"type"`        // 商品类型 1 虚拟物品 2实物物品 3电子卡券 4付费优惠券
	StockType  int8   `json:"stock_type"`  // 库存扣减方式 1拍下减库存 2付款减库存
	PriceStart int32  `json:"price_start"` // 商品价格 以分为单位
	PriceEnd   int32  `json:"price_end"`   // 商品价格 以分为单位
	StartTime  int    `json:"start_time"`  // 上架时间
	EndTime    int    `json:"end_time"`    // 下架时间
	State      int8   `json:"state"`       // 状态 1待发布 2正常 3下架 4停用
}

// start--------新的前端页面提交规则
type ShopGoods struct {
	ID         int64                   `json:"id"`                                               // 商品 ID
	Type       int8                    `binding:"required,oneof=1 2 3 4 5 6" json:"type"`            // 商品类型 1 虚拟物品 2实物物品 3电子卡券 4付费优惠券
	Name       string                  `binding:"required,max=100" json:"name" maxLength:"50"`   // 商品名称 100个字符以内
	Title      string                  `binding:"required,max=100" json:"title" maxLength:"100"` // 商品卖点 100个字符以内
	GroupID    int64                   `binding:"required" json:"group_id"`                      // 所属分组ID
	SnapshotId int64                   `json:"snapshot_id"`                                      // 快照ID 保存前做售出检查 如果数据基本信息变化 做快照备份
	RetailId   int64                   `json:"retail_id"`                                        // 分销商品ID
	Cover      string                  `binding:"required" json:"cover"`                         // 封面图片 JSON数组对应图片ID 最多上传15张
	StockID    int64                   `json:"stock_id"`                                         // 虚拟商品关联库存id
	StockType  int8                    `binding:"required,oneof=1 2" json:"stock_type"`          // 库存扣减方式 1拍下减库存 2付款减库存
	IsVip      int8                    `binding:"oneof=1 2" json:"is_vip"`                       // 是否参加会员折扣 1是 2不是
	IsPoints   int8                    `binding:"oneof=1 2" json:"is_points"`                    // 是否可使用积分购买 1是 2不是
	MaxPoints  int64                   `json:"max_points"`                                       // 最大可使用积分数
	PriceDot   int32                   `json:"price_dot" form:"price_dot"`                       // 划线价 以分为单位 不存在读 价格规则内最低值
	IsShowRest int8                    `binding:"required,oneof=1 2" json:"is_show_rest"`        // 商品详情、购物车是否显示剩余件数 1 显示 2不显示
	SaleTime   int64                   `validate:"datetime" json:"sale_time"`                    // 开售时间  -1 放入仓库 改变状态
	AfterSale  ShopGoodsAfterSale      `json:"after_sale,omitempty"`                             // 售后服务规则
	PayLimit   ShopGoodsPayLimit       `json:"pay_limit"`                                        // 购买限制规则
	GoodsData  ShopGoodsData           `json:"goods_data"`                                       // 商品详情数据
	GoodsRule  map[int16]GoodsRuleList `json:"goods_rule,omitempty"`                             // 商品价格库存规则 有规则情况下 商品库存失效  0 为总库存信息 其他自定义下 int16 为前端生成规则短ID
	State      int8                    `binding:"oneof=1 2 3 4" json:"state"`                   // 状态 1待发布 2正常 3下架 4 已售罄 5仓库中

}

// 售后服务规则 mysql
type ShopGoodsAfterSale struct {
	// 实物
	BuyerExchange      bool  `json:"buyer_exchange,omitempty"`       // 支持买家申请换货true false不支持
	BuyerSevenDays     bool  `json:"buyer_seven_days,omitempty"`     // 7天无理由退货
	BuyerSpeedExchange bool  `json:"buyer_speed_exchange,omitempty"` // 极速退款
	BuyerRefund        bool  `json:"buyer_refund,omitempty"`         // 支持买家申请退款true false不支持
	BuyerRefundExpire  int16 `json:"buyer_refund_expire,omitempty"`  // 未核销卡券过期多久前可退款 -1 未核销均支持退款 单位小时
}

// 商品购买限制规则 mysql
type ShopGoodsPayLimit struct {
	LimitPayMin      int8                        `binding:"max=9999" json:"limit_pay_min"`             // 每单最低购买数量
	LimitPayQuota    int8                        `binding:"oneof=-1 1 2 3 4 5" json:"limit_pay_quota"` // 限购 -1不限制 1每单 2终身 3每天 4每周 5每月 6每年
	LimitPayQuotaNum int16                       `json:"limit_pay_quota_num,omitempty"`                // 购买限制单数
	LimitSpecific    []ShopGoodsPayLimitSpecific `json:"limit_specific,omitempty"`                     // 指定用户可购买 根据 商品类型确定是否有改选项
}

// 指定限制规则
type ShopGoodsPayLimitSpecific struct {
	BindType int8    `binding:"oneof=1 2 3" json:"bind_type,omitempty"` // 1 会员卡 2 权益卡 3用户标签
	BindId   []int64 `json:"bind_id"`   // 绑定的相关ID列表
}

type GoodsRuleList struct {
	GoodsPriceName    GoodsPriceName      `json:"goods_price_name,omitempty"`  // 库存基本信息
	GoodsPriceStock   GoodsPriceStock     `json:"goods_price_stock,omitempty"` // 库存数量信息
	//GoodsDiscount     []GoodsDiscount     `json:"goods_discount,omitempty"`    // 打折信息
	GoodsRuleListInfo []GoodsRuleListInfo `json:"goods_info,omitempty"`        // 规则信息
}

// 商品价格规则 名称
type GoodsPriceName struct {
	Price     int32  `binding:"required" json:"price"` // 商品价格 以分为单位
	PriceDot  int32  `json:"price_dot"`                // 划线价 以分为单位
	PriceCost int32  `json:"price_cost"`               // 成本价 以分为单位
	GoodsNo   string `json:"goods_no" maxLength:"20"`  // 商品编码 20个字符以内
}

// 商品价格规则 储存至 mongodb
type GoodsPriceStock struct {
	StockID  int64 `json:"stock_id"`                     // 关联库存ID 关联特殊功能值 比如优惠券ID
	StockNum int64 `binding:"required" json:"stock_num"` // 库存数量
	SoldNum  int64 `json:"sold_num"`                     // 销量
}

// 会员价设置
type GoodsDiscount struct {
	BindType      int64 `binding:"required,oneof="1 2" json:"bind_type"` // 1 会员卡 2 权益卡
	DiscountType  int64 `json:"discount_type"`                // 优惠方式 1打折 2减价 3指定价格
	DiscountValue int64 `json:"discount_value"`               // 打折 打折7.5折传75 价格以分为单位
}

// API 接口 实现
type GoodsRuleListInfo struct {
	Name  string `binding:"required" json:"name"`  // 规则名称
	Type  int64  `binding:"required,oneof=1 2 3 4" json:"type"`  // 规则类型  1文本 2数字 3时间 4身份证
	Value string `binding:"required" json:"value"` // 规则值
}

// 商品详情规则储存到 mysql
type ShopGoodsData struct {
	Video        string                    `json:"video"`                                      // 主图视频  35个字符长度限制
	Pic          []string                  `binding:"required,max=15,dive,max=255" json:"pic"` //商品图 JSON数组对应图片ID 最多上传15张
	Share        string                    `binding:"max=36" json:"share"`                     // 分享描述 微信分享给好友时会显示，建议36个字以内
	SubName      string                    `binding:"max=20" json:"sub_name"`                  // 购买按钮名称 20个字以内
	SpecialRules ShopGoodsDataSpecialRules `json:"special_rules,omitempty"`                    // 特殊规则、自定义商品规则 根据商品类型确定
	Word         []ShopGoodsDataWords      `json:"word"`                                       // 用户留言数组 {名称,模型,必填} JSON数组 如 电话,phone,1 购买商品时让买家输入留言，最多可设置 10 条留言
	Content      string                    `json:"content"`                                    // 使用说明
}

// 商品留言信息规则 mysql
type ShopGoodsDataWords struct {
	Name       string `json:"name"`                             // 用户留言字段名称 phone
	Type       int8   `binding:"oneof=1 2 3 4" json:"type"`     // 用户留言字段对应的值  1文本 2数字 3时间 4身份证 5 手机号
	IsMultiRow int8   `binding:"oneof=1 2" json:"is_multi_row"` // 是否多行 1 是多行 2 不是多行
	IsRequired int8   `binding:"oneof=1 2" json:"is_required"`  // 是否必填 1 是必填 2 不是必填
	IsSingle   int8   `binding:"oneof=1 2" json:"is_single"`    // 1 是 (只需要用户填1次) 2否 (根据购买数量填写,买N张门票需要填写N个身份证)
}

// 自定义规则
type ShopGoodsDataSpecialRules struct {
	GoodsReal *ShopGoodsDataSpecialRulesGoodsReal `json:"goods_real,omitempty"` // 实物商品自定义规则 omitempty有值就需要校验
	GoodsCake *ShopGoodsDataSpecialRulesGoodsCake `json:"goods_cake,omitempty"` // 蛋糕烘焙 omitempty有值就需要校验
	GoodsCard *ShopGoodsDataSpecialRulesGoodsCard `json:"goods_card,omitempty"` // 电子卡券规则 omitempty有值就需要校验
}

// 自定义规则 实物商品
type ShopGoodsDataSpecialRulesGoodsReal struct {
	DeliveryType []int8 `binding:"required" json:"delivery_type"`         // 配送方式 1 快递发货 2 同城配送 3 到店自提
	FreightPrice int32  `binding:"required" json:"freight_price,omitempty"` // 运费价格 统一价格 -1 为调用运费模板
	FreightTpl   int32  `json:"freight_tpl"`                                 // 运费模板
}

// 自定义规则 电子卡券规则
type ShopGoodsDataSpecialRulesGoodsCard struct {
	ValidityMode        int8   `json:"validity_mode"`       // 卡券生效模式 1 立即生效 2 次日生效 3多少小时后生效
	ValidityType        int8   `binding:"oneof=-1 1 2" json:"validity_type"`       // 卡券生效类型 ( -1长期有效 1指定天数 2指定时间段)
	ValidityDay         int8   `json:"validity_day"`                              // 卡券生效时长 (生效类型 如：5 指5天内有效)
	ValidityStartTime   int64  `validate:"datetime" json:"validity_start_time"`   // 具体限制卡券 有效期开始时间
	ValidityEndTime     int64  `validate:"datetime" json:"validity_end_time"`     // 具体限制卡券 有效期结束时间
	IsHolidaysAvailable int8   `binding:"oneof=1 2" json:"is_holidays_available"` // 节假日是否可用 1是 2不是
	UseNotes            string `json:"use_notes"`                                 // 使用说明
}

// 自定义规则 蛋糕烘焙
type ShopGoodsDataSpecialRulesGoodsCake struct {
	DeliveryType []int8 `binding:"required" json:"delivery_type"`    // 配送方式 1 同城配送 2 到店自提
	AttributeIds []int64         `json:"attribute_ids"` // 商品附加属性值 关联到 商品属性库ID
	StockUpTime  int64           `json:"stock_up_time"` // 统一备货时间 单位分钟 -1 启用不同规格单独设置备货时间
	StockUp      map[int16]int64 `json:"attribute"`     // 备货时间 map[SKUID]时间 单位分钟
}

// end--------新的前端页面提交规则
