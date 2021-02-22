package mall_shop

import "base/model/modelSql"

// ShopGoods [...]

type ResShopGoods struct {
	GoodsId int64 `json:"goods_id" form:"goods_id"`
	RuleId  int64 `json:"rule_id" form:"rule_id"`
	DataId  int64 `json:"data_id" form:"data_id"`
}
type ShopGoodsClear struct {
	modelSql.Database
	Coid      int64   `gorm:"column:coid;type:bigint(17);not null" json:"coid" form:"coid"`                   // 所属企业
	GroupID   int64   `gorm:"column:group_id;type:bigint(17);not null" json:"group_id" form:"group_id"`       // 所属分组ID
	Name      string  `gorm:"column:name;type:varchar(50);not null" json:"name" form:"name"`                  // 商品名称
	Title     string  `gorm:"column:title;type:varchar(100);not null" json:"title" form:"title"`              // 商品标题
	Type      int8    `gorm:"column:type;type:tinyint(2);not null" json:"type" form:"type"`                   // 商品类型 1 虚拟物品 2实物物品 3电子卡券 4付费优惠券
	Cover     string  `gorm:"column:cover;type:varchar(35)" json:"cover" form:"cover"`                        // 封面图片
	StockID   int64   `gorm:"column:stock_id;type:bigint(17);not null" json:"stock_id" form:"stock_id"`       // 关联库存ID
	StockType int8    `gorm:"column:stock_type;type:tinyint(2);not null" json:"stock_type" form:"stock_type"` // 库存扣减方式 1拍下减库存 2付款减库存
	StockNum  int     `gorm:"column:stock_num;type:int(10);not null" json:"stock_num" form:"stock_num"`       // 库存数量
	IsVip     int8    `gorm:"column:is_vip;type:tinyint(2);not null" json:"is_vip" form:"is_vip"`             // 是否参加会员折扣
	IsPoints  int8    `gorm:"column:is_points;type:tinyint(2);not null" json:"is_points" form:"is_points"`    // 是否可使用积分购买
	MaxPoints int     `gorm:"column:max_points;type:int(10);not null" json:"max_points" form:"max_points"`    // 最大可使用积分数
	Price     float64 `gorm:"column:price;not null" json:"price" form:"price"`                                // 商品价格
	PriceDot  float64 `gorm:"column:price_dot;not null" json:"price_dot" form:"price_dot"`                    // 划线价
	PriceCost float64 `gorm:"column:price_cost;not null" json:"price_cost" form:"price_cost"`                 // 成本价
	GoodsNo   string  `gorm:"column:goods_no;type:varchar(20);not null" json:"goods_no" form:"goods_no"`      // 商品编码
	SaleMin   int8    `gorm:"column:sale_min;type:tinyint(4);not null" json:"sale_min" form:"sale_min"`       // 每单最低购买数量
	SaleMax   int8    `gorm:"column:sale_max;type:tinyint(4);not null" json:"sale_max" form:"sale_max"`       // 每单最多购买数量
	Quota     int8    `gorm:"column:quota;type:tinyint(2);not null" json:"quota" form:"quota"`                // 限购 1终身 2每天 3每周 4每月 5每年
	QuotaNum  int     `gorm:"column:quota_num;type:int(11);not null" json:"quota_num" form:"quota_num"`       // 购买限制单数
	StartTime int     `gorm:"column:start_time;type:int(10);not null" json:"start_time" form:"start_time"`    // 上架时间
	EndTime   int     `gorm:"column:end_time;type:int(10);not null" json:"end_time" form:"end_time"`          // 下架时间
	State     int8    `gorm:"column:state;type:tinyint(2);not null" json:"state" form:"state"`                // 状态 1待发布 2正常 3下架 4停用
}

func (m *ShopGoodsClear) TableName() string {
	return "shop_goods"
}

type ShopGoods struct {
	modelSql.Database
	Coid      int64           `gorm:"column:coid;type:bigint(17);not null" json:"coid" form:"coid"`                   // 所属企业
	GroupID   int64           `gorm:"column:group_id;type:bigint(17);not null" json:"group_id" form:"group_id"`       // 所属分组ID
	Name      string          `gorm:"column:name;type:varchar(50);not null" json:"name" form:"name"`                  // 商品名称
	Title     string          `gorm:"column:title;type:varchar(100);not null" json:"title" form:"title"`              // 商品标题
	Type      int8            `gorm:"column:type;type:tinyint(2);not null" json:"type" form:"type"`                   // 商品类型 1 虚拟物品 2实物物品 3电子卡券 4付费优惠券
	Cover     string          `gorm:"column:cover;type:varchar(35)" json:"cover" form:"cover"`                        // 封面图片
	StockID   int64           `gorm:"column:stock_id;type:bigint(17);not null" json:"stock_id" form:"stock_id"`       // 关联库存ID
	StockType int8            `gorm:"column:stock_type;type:tinyint(2);not null" json:"stock_type" form:"stock_type"` // 库存扣减方式 1拍下减库存 2付款减库存
	StockNum  int             `gorm:"column:stock_num;type:int(10);not null" json:"stock_num" form:"stock_num"`       // 库存数量
	IsVip     int8            `gorm:"column:is_vip;type:tinyint(2);not null" json:"is_vip" form:"is_vip"`             // 是否参加会员折扣
	IsPoints  int8            `gorm:"column:is_points;type:tinyint(2);not null" json:"is_points" form:"is_points"`    // 是否可使用积分购买
	MaxPoints int             `gorm:"column:max_points;type:int(10);not null" json:"max_points" form:"max_points"`    // 最大可使用积分数
	Price     float64         `gorm:"column:price;not null" json:"price" form:"price"`                                // 商品价格
	PriceDot  float64         `gorm:"column:price_dot;not null" json:"price_dot" form:"price_dot"`                    // 划线价
	PriceCost float64         `gorm:"column:price_cost;not null" json:"price_cost" form:"price_cost"`                 // 成本价
	GoodsNo   string          `gorm:"column:goods_no;type:varchar(20);not null" json:"goods_no" form:"goods_no"`      // 商品编码
	SaleMin   int8            `gorm:"column:sale_min;type:tinyint(4);not null" json:"sale_min" form:"sale_min"`       // 每单最低购买数量
	SaleMax   int8            `gorm:"column:sale_max;type:tinyint(4);not null" json:"sale_max" form:"sale_max"`       // 每单最多购买数量
	Quota     int8            `gorm:"column:quota;type:tinyint(2);not null" json:"quota" form:"quota"`                // 限购 1终身 2每天 3每周 4每月 5每年
	QuotaNum  int             `gorm:"column:quota_num;type:int(11);not null" json:"quota_num" form:"quota_num"`       // 购买限制单数
	StartTime int             `gorm:"column:start_time;type:int(10);not null" json:"start_time" form:"start_time"`    // 上架时间
	EndTime   int             `gorm:"column:end_time;type:int(10);not null" json:"end_time" form:"end_time"`          // 下架时间
	State     int8            `gorm:"column:state;type:tinyint(2);not null" json:"state" form:"state"`                // 状态 1待发布 2正常 3下架 4停用
	Rule      []ShopGoodsRule `gorm:"ForeignKey:goods_id"`
	Data      ShopGoodsData   `json:"data" form:"data`
	Group     ShopGroup       `json:"group" form:"group`
	Stock     ShopStock       `json:"stock" form:"stock`
}
type SearchGoods struct {
	Coid       int64   `gorm:"column:coid;type:bigint(17);not null" json:"coid" form:"coid"`                   // 所属企业
	GroupID    int64   `gorm:"column:group_id;type:bigint(17);not null" json:"group_id" form:"group_id"`       // 所属分组ID
	Name       string  `gorm:"column:name;type:varchar(50);not null" json:"name" form:"name"`                  // 商品名称
	Type       int8    `gorm:"column:type;type:tinyint(2);not null" json:"type" form:"type"`                   // 商品类型 1 虚拟物品 2实物物品 3电子卡券 4付费优惠券
	StockType  int8    `gorm:"column:stock_type;type:tinyint(2);not null" json:"stock_type" form:"stock_type"` // 库存扣减方式 1拍下减库存 2付款减库存
	PriceStart float64 `gorm:"column:price;not null" json:"price_start" form:"price_start"`                    // 商品价格
	PriceEnd   float64 `gorm:"column:price;not null" json:"price_end" form:"price_end"`                        // 商品价格
	StartTime  int     `gorm:"column:start_time;type:int(10);not null" json:"start_time" form:"start_time"`    // 上架时间
	EndTime    int     `gorm:"column:end_time;type:int(10);not null" json:"end_time" form:"end_time"`          // 下架时间
	State      int8    `gorm:"column:state;type:tinyint(2);not null" json:"state" form:"state"`                // 状态 1待发布 2正常 3下架 4停用
}

func (m *SearchGoods) TableName() string {
	return "shop_goods"
}

func (m *ShopGoods) TableName() string {
	return "shop_goods"
}

// ShopGoodsData 存入
type ShopGoodsData struct {
	ID               int64  `gorm:"column:id;type:bigint(17);not null" json:"id"`                                 // 商品ID
	Pic              string `gorm:"column:pic;type:varchar(255);not null" json:"pic"`                             // 商品图  JSON数组对应图片ID
	Video            string `gorm:"column:video;type:varchar(35);not null" json:"video"`                          // 主图视频
	Share            string `gorm:"column:share;type:varchar(255);not null" json:"share"`                         // 分享描述
	Word             string `gorm:"column:word;type:varchar(255);not null" json:"word"`                           // 用户留言数组 {名称,模型,必填} JSON数组 如 电话,phone,1
	SubName          string `gorm:"column:sub_name;type:varchar(20);not null" json:"sub_name"`                    // 购买按钮名称
	Content          string `gorm:"column:content;not null" json:"content"`                                       // 商品详情
	SaleTime         int64  `gorm:"column:sale_time;not null" json:"sale_time"`                                   // 开售时间
	IsShowRest       int32  `gorm:"column:is_show_rest;not null" json:"is_show_rest"`                             // 商品详情、购物车是否显示剩余件数 1 显示 2不显示
	ElectronicCoupon string `gorm:"column:electronic_coupon;type:varchar(500);not null" json:"electronic_coupon"` // 电子卡券特有的属性 电子卡券有效期和退款
}

// ShopGoodsData 取出来
type ShopGoodsDataNew struct {
	ID      int64   `gorm:"column:id;type:bigint(17);not null" json:"id"`              // 商品ID
	Pic     string  `gorm:"column:pic;type:varchar(255);not null" json:"pic"`          // 商品图  JSON数组对应图片ID
	Video   string  `gorm:"column:video;type:varchar(35);not null" json:"video"`       // 主图视频
	Share   string  `gorm:"column:share;type:varchar(255);not null" json:"share"`      // 分享描述
	Word    []uint8 `gorm:"column:word;type:varchar(255);not null" json:"word"`        // 用户留言数组 {名称,模型,必填} JSON数组 如 电话,phone,1
	SubName string  `gorm:"column:sub_name;type:varchar(20);not null" json:"sub_name"` // 购买按钮名称
	Content string  `gorm:"column:content;not null" json:"content"`                    // 商品详情
}

func (m *ShopGoodsDataNew) TableName() string {
	return "shop_goods_data"
}
func (m *ShopGoodsData) TableName() string {
	return "shop_goods_data"
}

// ShopGoodsRule [...]
type ShopGoodsRule struct {
	modelSql.Database
	GoodsID   int64   `gorm:"column:goods_id;type:bigint(17);not null" json:"goods_id" form:"goods_id"`  // 商品ID
	StockID   int64   `gorm:"column:stock_id;type:bigint(17);not null" json:"stock_id" form:"stock_id"`  // 关联库存ID
	StockNum  int     `gorm:"column:stock_num;type:int(10);not null" json:"stock_num" form:"stock_num"`  // 库存数量
	ActStart  int     `gorm:"column:act_start;type:int(10);not null" json:"act_start" form:"act_start"`  // 规则开始时间 0总是生效
	ActEnd    int     `gorm:"column:act_end;type:int(10);not null" json:"act_end" form:"act_end"`        // 规则结束时间
	Price     float64 `gorm:"column:price;not null" json:"price" form:"price"`                           // 规则价格
	PriceCost float64 `gorm:"column:price_cost;not null" json:"price_cost" form:"price_cost"`            // 成本价
	GoodsNo   string  `gorm:"column:goods_no;type:varchar(20);not null" json:"goods_no" form:"goods_no"` // 商品编码
	Specs     string  `gorm:"column:specs" json:"specs" form:"specs"`                                    // 规则数据 JSON
	Sort      int     `gorm:"column:sort;type:int(11);not null" json:"sort" form:"sort"`                 // 排序
}

func (m *ShopGoodsRule) TableName() string {
	return "shop_goods_rule"
}

type ShopGroupDetail struct {
	ID        int64  `gorm:"primary_key;column:id;" json:"id"`
	CreatedAt int    `gorm:"column:created_at;type:int(10);not null" json:"created_at"` // 添加时间
	UpdatedAt int    `gorm:"column:updated_at;type:int(10);not null" json:"updated_at"` // 更新时间
	Coid      int64  `gorm:"column:coid;type:bigint(17);not null" json:"coid"`
	Name      string `gorm:"column:name;type:varchar(50);not null" json:"name"` // 商铺分组名称
}

// ShopGroup [...]
type ShopGroup struct {
	modelSql.Database
	Coid int64  `gorm:"column:coid;type:bigint(17);not null" json:"coid" form:"coid"`  // 所属 企业ID
	Name string `gorm:"column:name;type:varchar(50);not null" json:"name" form:"name"` // 商铺分组名称
}

func (m *ShopGroup) TableName() string {
	return "shop_group"
}

type ShopGroupVO struct {
	modelSql.Database
	Coid       int64  `gorm:"column:coid;type:bigint(17);not null" json:"coid" form:"coid"`  // 所属 企业ID
	Name       string `gorm:"column:name;type:varchar(50);not null" json:"name" form:"name"` // 商铺分组名称
	GoodsCount int32  `json:"goods_count" form:"goods_count"`                                // 不是表的字段 返回分组的商品数量
}

func (m *ShopGroupVO) TableName() string {
	return "shop_group"
}

// 用户订单
type ShopOrder struct {
	modelSql.Database
	GoodsID   int64   `gorm:"column:goods_id;type:bigint(17);not null" json:"goods_id" form:"goods_id"`       // 商品ID
	Coid      int64   `gorm:"column:coid;type:bigint(17);not null" json:"coid" form:"coid"`                   // 商家ID
	Uid       int64   `gorm:"column:uid;type:bigint(17);not null" json:"uid" form:"uid"`                      // 购买用户ID
	MailingID int64   `gorm:"column:mailing_id;type:bigint(17);not null" json:"mailing_id" form:"mailing_id"` // 用户选择的邮寄地址
	Type      int8    `gorm:"column:type;type:tinyint(2)" json:"type" form:"type"`                            // 商品类型 1 虚拟物品 2实物物品 3电子卡券 4付费优惠券
	Amount    float64 `gorm:"column:amount;not null" json:"amount" form:"amount"`                             // 总金额
	Price     float64 `gorm:"column:price;not null" json:"price" form:"price"`                                // 需支付价格
	Discount  float64 `gorm:"column:discount;not null" json:"discount" form:"discount"`                       // 优惠价格
	Freight   float64 `gorm:"column:freight;not null" json:"freight" form:"freight"`                          // 运费
	State     int8    `gorm:"column:state;type:tinyint(2);not null" json:"state" form:"state"`                // 状态 1待付款 2待发货 3已发货 4已完成 5已关闭 6售后中
}

func (m *ShopOrder) TableName() string {
	return "shop_order"
}

// ShopOrderData [...]
type ShopOrderData struct {
	ID          int64  `gorm:"column:id;type:bigint(20);not null" json:"id" form:"id"`                           // 用户订单ID
	PackageType int    `gorm:"column:package_type;type:int(5);not null" json:"package_type" form:"package_type"` // 快递类型
	PackageNo   string `gorm:"column:package_no;type:varchar(30);not null" json:"package_no" form:"package_no"`  // 订单号
	Word        string `gorm:"column:word;type:varchar(100);not null" json:"word" form:"word"`                   // 用户留言数组 JSON数组 对应商品 用户留言数组
	Notes       string `gorm:"column:notes;type:varchar(100);not null" json:"notes" form:"notes"`                // 用户备注
	PayTime     int    `gorm:"column:pay_time;type:int(10);not null" json:"pay_time" form:"pay_time"`            // 支付时间
}

func (m *ShopOrderData) TableName() string {
	return "shop_order_data"
}

// ShopStock [...]
type ShopStock struct {
	modelSql.Database
	Coid    int64  `gorm:"column:coid;type:bigint(17);not null" json:"coid" form:"coid"`         // 所属企业 ID
	Type    int8   `gorm:"column:type;type:tinyint(2);not null" json:"type" form:"type"`         // 类型 1自有 2会员权益 3优惠券
	Aid     int64  `gorm:"column:aid;type:bigint(17);not null" json:"aid" form:"aid"`            // 2会员权益ID
	Name    string `gorm:"column:name;type:varchar(30);not null" json:"name" form:"name"`        // 库存名称
	Notes   string `gorm:"column:notes;type:varchar(100);not null" json:"notes" form:"notes"`    // 库存备注
	AutoNum int    `gorm:"column:auto_num;type:int(4);not null" json:"auto_num" form:"auto_num"` // 库存不足 自动增加数量
}

func (m *ShopStock) TableName() string {
	return "shop_stock"
}

// ShopStockList [...]
type ShopStockCard struct {
	ID      int64  `gorm:"column:id;type:bigint(17);not null" json:"id" form:"id"`                   // 库存列表 ID
	StockID int64  `gorm:"column:stock_id;type:bigint(17);not null" json:"stock_id" form:"stock_id"` // 库存ID
	Card    string `gorm:"column:card;type:varchar(30);not null" json:"card" form:"card"`            // 卡号、券码
	Pass    string `gorm:"column:pass;type:varchar(30);not null" json:"pass" form:"pass"`            // 密码
	State   int8   `gorm:"column:state;type:tinyint(2);not null" json:"state" form:"state"`          // 状态 1有效 2已发送 3已使用
}

func (m *ShopStockCard) TableName() string {
	return "shop_stock_card"
}

type ReqShopCoupon struct {
	SC ShopCoupon
	SD ShopCouponData
}
type ShopCouponRelatedVO struct {
	modelSql.Database
	Coid             int64          `gorm:"column:coid;type:bigint(17);not null" json:"coid" form:"coid"`                                           // 所属企业ID
	Type             int8           `gorm:"column:type;type:tinyint(1);not null" json:"type" form:"type"`                                           // 优惠券类型 1满减券 2折扣券 3随机金额券 4商品兑换券
	Name             string         `gorm:"column:name;type:varchar(50);not null" json:"name" form:"name"`                                          // 优惠券名称
	Remark           string         `gorm:"column:remark;type:varchar(50);not null" json:"remark" form:"remark"`                                    // 名称备注
	UseStockQty      int            `gorm:"column:use_stock_qty;type:int(10);not null" json:"use_stock_qty" form:"use_stock_qty"`                   // 发放总量
	UsedQty          int            `gorm:"column:used_qty;type:int(10);not null" json:"used_qty" form:"used_qty"`                                  // 当前已使用量
	UseType          int8           `gorm:"column:use_type;type:tinyint(2);not null" json:"use_type" form:"use_type"`                               // 适用商品 1全部商品可用 2指定商品可用 3指定商品不可用
	UsingAmount      float64        `gorm:"column:using_amount;not null" json:"using_amount" form:"using_amount"`                                   // 使用门槛 -1无使用门槛 订单满多少元
	UseMode          int8           `gorm:"column:use_mode;type:tinyint(2);not null" json:"use_mode" form:"use_mode"`                               // 用券模式 1指定时间 2领券当日起 3领券次日起
	UseDay           int8           `gorm:"column:use_day;type:tinyint(4);not null" json:"use_day" form:"use_day"`                                  // 用券模式 2、3时指定天数
	IsOverlay        int8           `gorm:"column:is_overlay;type:tinyint(2);not null" json:"is_overlay" form:"is_overlay"`                         // 优惠叠加 优惠券仅原价购买时可用
	IsPublic         int8           `gorm:"column:is_public;type:tinyint(2);not null" json:"is_public" form:"is_public"`                            // 允许公开领取 1不允许 2允许
	IsSharable       int8           `gorm:"column:is_sharable;type:tinyint(2);not null" json:"is_sharable" form:"is_sharable"`                      // 是否可共享 1不可共享 2可共享
	IsHandSel        int8           `gorm:"column:is_hand_sel;type:tinyint(2);not null" json:"is_hand_sel" form:"is_hand_sel"`                      // 转赠设置 1不可 2允许转赠给好友
	IsExpireNotice   int8           `gorm:"column:is_expire_notice;type:tinyint(4);not null" json:"is_expire_notice" form:"is_expire_notice"`       // 过期提醒
	ExpireNoticeDays int8           `gorm:"column:expire_notice_days;type:tinyint(4);not null" json:"expire_notice_days" form:"expire_notice_days"` // 过期前几天提醒
	LimitType        int8           `gorm:"column:limit_type;type:tinyint(2);not null" json:"limit_type" form:"limit_type"`                         // 领取人限制
	LimitNum         int8           `gorm:"column:limit_num;type:tinyint(4);not null" json:"limit_num" form:"limit_num"`                            // 每人限领次数 -1不限次数 大于1领取次数
	StartTime        int            `gorm:"column:start_time;type:int(10);not null" json:"start_time" form:"start_time"`                            // 用券开始时间
	EndTime          int            `gorm:"column:end_time;type:int(10);not null" json:"end_time" form:"end_time"`                                  // 用券结束时间
	State            int8           `gorm:"column:state;type:tinyint(2);not null" json:"state" form:"state"`                                        // 状态 1正常 2失效
	ShopCouponData   ShopCouponData `gorm:"ForeignKey:id" json:"coupon_data"`
}

func (m *ShopCouponRelatedVO) TableName() string {
	return "shop_coupon"
}

type ShopCoupon struct {
	modelSql.Database
	Coid             int64          `gorm:"column:coid;type:bigint(17);not null" json:"coid" form:"coid"`                                           // 所属企业ID
	Type             int8           `gorm:"column:type;type:tinyint(1);not null" json:"type" form:"type"`                                           // 优惠券类型 1满减券 2折扣券 3随机金额券 4商品兑换券
	Name             string         `gorm:"column:name;type:varchar(50);not null" json:"name" form:"name"`                                          // 优惠券名称
	Remark           string         `gorm:"column:remark;type:varchar(50);not null" json:"remark" form:"remark"`                                    // 名称备注
	UseStockQty      int            `gorm:"column:use_stock_qty;type:int(10);not null" json:"use_stock_qty" form:"use_stock_qty"`                   // 发放总量
	UsedQty          int            `gorm:"column:used_qty;type:int(10);not null" json:"used_qty" form:"used_qty"`                                  // 当前已使用量
	UseType          int8           `gorm:"column:use_type;type:tinyint(2);not null" json:"use_type" form:"use_type"`                               // 适用商品 1全部商品可用 2指定商品可用 3指定商品不可用
	UsingAmount      float64        `gorm:"column:using_amount;not null" json:"using_amount" form:"using_amount"`                                   // 使用门槛 -1无使用门槛 订单满多少元
	UseMode          int8           `gorm:"column:use_mode;type:tinyint(2);not null" json:"use_mode" form:"use_mode"`                               // 用券模式 1指定时间 2领券当日起 3领券次日起
	UseDay           int8           `gorm:"column:use_day;type:tinyint(4);not null" json:"use_day" form:"use_day"`                                  // 用券模式 2、3时指定天数
	IsOverlay        int8           `gorm:"column:is_overlay;type:tinyint(2);not null" json:"is_overlay" form:"is_overlay"`                         // 优惠叠加 优惠券仅原价购买时可用
	IsPublic         int8           `gorm:"column:is_public;type:tinyint(2);not null" json:"is_public" form:"is_public"`                            // 允许公开领取 1不允许 2允许
	IsSharable       int8           `gorm:"column:is_sharable;type:tinyint(2);not null" json:"is_sharable" form:"is_sharable"`                      // 是否可共享 1不可共享 2可共享
	IsHandSel        int8           `gorm:"column:is_hand_sel;type:tinyint(2);not null" json:"is_hand_sel" form:"is_hand_sel"`                      // 转赠设置 1不可 2允许转赠给好友
	IsExpireNotice   int8           `gorm:"column:is_expire_notice;type:tinyint(4);not null" json:"is_expire_notice" form:"is_expire_notice"`       // 过期提醒
	ExpireNoticeDays int8           `gorm:"column:expire_notice_days;type:tinyint(4);not null" json:"expire_notice_days" form:"expire_notice_days"` // 过期前几天提醒
	LimitType        int8           `gorm:"column:limit_type;type:tinyint(2);not null" json:"limit_type" form:"limit_type"`                         // 领取人限制
	LimitNum         int8           `gorm:"column:limit_num;type:tinyint(4);not null" json:"limit_num" form:"limit_num"`                            // 每人限领次数 -1不限次数 大于1领取次数
	StartTime        int            `gorm:"column:start_time;type:int(10);not null" json:"start_time" form:"start_time"`                            // 用券开始时间
	EndTime          int            `gorm:"column:end_time;type:int(10);not null" json:"end_time" form:"end_time"`                                  // 用券结束时间
	State            int8           `gorm:"column:state;type:tinyint(2);not null" json:"state" form:"state"`                                        // 状态 1正常 2失效
	CouponData       ShopCouponData `json:"-" form:"-"`
}

func (m *ShopCoupon) TableName() string {
	return "shop_coupon"
}

type ShopCouponData struct {
	ID                 int64  `gorm:"column:id;type:bigint(17);not null" json:"id" form:"id"`                                                         // 优惠券ID
	UseGoods           string `gorm:"column:use_goods;type:varchar(255);not null" json:"use_goods" form:"use_goods"`                                  // 适用商品 商品ID数组
	UseRule            string `gorm:"column:use_rule;type:varchar(255);not null" json:"use_rule" form:"use_rule"`                                     // 使用规则数据 1减免?元 2打?折最多优惠?元 3随机?至?元
	LimitBenefitMember string `gorm:"column:limit_benefit_member;type:varchar(255);not null" json:"limit_benefit_member" form:"limit_benefit_member"` // 领取人身份 客户身份权益卡
	LimitLevelMember   string `gorm:"column:limit_level_member;type:varchar(255);not null" json:"limit_level_member" form:"limit_level_member"`       // 领取人身份 客户身份会员
	Description        string `gorm:"column:description;type:varchar(255);not null" json:"description" form:"description"`                            // 使用说明
}

func (m *ShopCouponData) TableName() string {
	return "shop_coupon_data"
}

type ShopCouponCard struct {
	ID        int64 `gorm:"column:id;type:bigint(17);not null" json:"id" form:"id"`                      // 兑换码
	Coid      int64 `gorm:"column:coid;type:bigint(17);not null" json:"coid" form:"coid"`                // 所属企业ID
	Uid       int64 `gorm:"column:uid;type:bigint(17);not null" json:"uid" form:"uid"`                   // 所属用户ID
	StartTime int   `gorm:"column:start_time;type:int(10);not null" json:"start_time" form:"start_time"` // 开始时间
	EndTime   int   `gorm:"column:end_time;type:int(10);not null" json:"end_time" form:"end_time"`       // 到期时间
	State     int8  `gorm:"column:state;type:tinyint(2);not null" json:"state" form:"state"`             // 状态 1待领取 2待使用 3使用中 3使用完成
}

func (m *ShopCouponCard) TableName() string {
	return "shop_coupon_card"
}

type ShopGoodsAsset struct {
	modelSql.Database
	GoodsID     int64 `gorm:"column:goods_id;type:bigint(17);not null" json:"goods_id" form:"goods_id"`                // 商品ID
	GoodsRuleID int64 `gorm:"column:goods_rule_id;type:bigint(17);not null" json:"goods_rule_id" form:"goods_rule_id"` // 商品规则ID
	StockNum    int   `gorm:"column:stock_num;type:int(10);not null" json:"stock_num" form:"stock_num"`                // 库存数量
	Sales       int   `gorm:"column:sales;type:int(10);not null" json:"sales" form:"sales"`                            // 销量
}

func (m *ShopGoodsAsset) TableName() string {
	return "shop_goods_asset"
}

type CoShopGoodsBatchSettingVO struct {
	Coid        int64         `json:"coid" form:"coid"`                 // 企业ID
	Type        int32         `json:"type" form:"type"`                 // 修改类型  1删除  2正常(上架) 3下架 4停用  5分组 6会员价 7商品名称替换  8上架时间 9每人限购 10会员折扣 11限定用户(限定用户有购买的权限) 12 配送方式 13商品模板
	GoodsList   []int64       `json:"goods_list" form:"goods_list"`     // 要操作的商品ID数组
	FeatureList FeatureListVO `json:"feature_list" form:"feature_list"` //要修改的的商品属性值
}

type FeatureListVO struct {
	GroupID   int64             `json:"group_id"`
	StartTime int64             `json:"start_time"` // 上架时间
	Quota     int32             `json:"quota"`      // 每人限购类型 1终身 2每天 3每周 4每月 5每年
	QuotaNum  int32             `json:"quota_num"`  // 每人购买限制单数
	IsVip     int32             `json:"is_vip"`     //是否参加会员折扣 1参加 2 不参加
	Member    []MemberPriceVO   `json:"member"`     // 会员价设置
	GoodsName GoodsNameChangeVO `json:"goods_name"`
}
type MemberPriceVO struct {
	GoodsId      int64          `json:"goods_id"`
	SkuId        int64          `json:"sku_id"`
	DiscountType []DiscountType `json:"discount_type"`
}
type DiscountType struct {
	LevelOrBenefit int32     `json:"level_or_benefit"` //1 按会员等级设置 2 按权益卡设置
	DiscountMethod int32     `json:"discount_method"`  // 优惠方式：1打折 2减价 3指定价格
	CountValue     []float64 `json:"count_value"`      // 优惠值 打7.5折 减5.5元 指定100.5元 一个商品对应多个值
}
type GoodsCountVO struct {
	GoodsId    int64     `json:"goods_id"`
	CountValue []float64 `json:"count_value"` // 优惠值 打7.5折 减5.5元 指定100.5元 一个商品对应多个值
}
type GoodsNameChangeVO struct {
	NameList []string `json:"name_list"` // 商品名称完整的数组 跟GoodsList对应
	OldName  string   `json:"old_name"`  // 商品名称里的文字
	NewName  string   `json:"new_name"`  // 商品名称文字替换为新文字 为空表示删除商品名称里的oldName文字
}

type CoShopGoods struct {
	modelSql.Database
	Coid               int64  `gorm:"column:coid;type:bigint(17);not null" json:"coid" form:"coid"`                   // 所属企业
	GroupID            int64  `gorm:"column:group_id;type:bigint(17);not null" json:"group_id" form:"group_id"`       // 所属分组ID
	Name               string `gorm:"column:name;type:varchar(50);not null" json:"name" form:"name"`                  // 商品名称
	Title              string `gorm:"column:title;type:varchar(100);not null" json:"title" form:"title"`              // 商品标题
	Type               int8   `gorm:"column:type;type:tinyint(2);not null" json:"type" form:"type"`                   // 商品类型 1 虚拟物品 2实物物品 3电子卡券 4付费优惠券
	Cover              string `gorm:"column:cover;type:varchar(35)" json:"cover" form:"cover"`                        // 封面图片
	StockID            int64  `gorm:"column:stock_id;type:bigint(17);not null" json:"stock_id" form:"stock_id"`       // 关联库存ID
	StockType          int8   `gorm:"column:stock_type;type:tinyint(2);not null" json:"stock_type" form:"stock_type"` // 库存扣减方式 1拍下减库存 2付款减库存
	StockNum           int32  `gorm:"column:stock_num;type:int(10);not null" json:"stock_num" form:"stock_num"`       // 库存数量
	IsVip              int8   `gorm:"column:is_vip;type:tinyint(2);not null" json:"is_vip" form:"is_vip"`             // 是否参加会员折扣
	IsPoints           int8   `gorm:"column:is_points;type:tinyint(2);not null" json:"is_points" form:"is_points"`    // 是否可使用积分购买
	MaxPoints          int32  `gorm:"column:max_points;type:int(10);not null" json:"max_points" form:"max_points"`    // 最大可使用积分数
	Price              int32  `gorm:"column:price;not null" json:"price" form:"price"`                                // 商品价格
	PriceDot           int32  `gorm:"column:price_dot;not null" json:"price_dot" form:"price_dot"`                    // 划线价
	PriceCost          int32  `gorm:"column:price_cost;not null" json:"price_cost" form:"price_cost"`                 // 成本价
	GoodsNo            string `gorm:"column:goods_no;type:varchar(20);not null" json:"goods_no" form:"goods_no"`      // 商品编码
	LimitPayMin        int8   `gorm:"column:limit_pay_min;type:tinyint(4);not null" json:"limit_pay_min"`             // 每单最低购买数量
	SaleMax            int8   `gorm:"column:sale_max;type:tinyint(4);not null" json:"sale_max" form:"sale_max"`       // 每单最多购买数量
	Quota              int8   `gorm:"column:quota;type:tinyint(2);not null" json:"quota" form:"quota"`                // 限购 1终身 2每天 3每周 4每月 5每年
	LimitPayQuotaNum   int32  `gorm:"column:limit_pay_quota_num;type:int(11);not null" json:"limit_pay_quota_num"`    // 购买限制单数
	StartTime          int32  `gorm:"column:start_time;type:int(10);not null" json:"start_time" form:"start_time"`    // 上架时间
	EndTime            int32  `gorm:"column:end_time;type:int(10);not null" json:"end_time" form:"end_time"`          // 下架时间
	State              int8   `gorm:"column:state;type:tinyint(2);not null" json:"state" form:"state"`                // 状态 1待发布 2正常 3下架 4停用
	BuyerExchange      int8   `gorm:"column:buyer_exchange;type:tinyint(2);not null" json:"buyer_exchange"`
	BuyerSevenDays     int8   `gorm:"column:buyer_seven_days;type:tinyint(2);not null" json:"buyer_seven_days"`
	BuyerSpeedExchange int8   `gorm:"column:buyer_speed_exchange;type:tinyint(2);not null" json:"buyer_speed_exchange"`
	BuyerRefund        int8   `gorm:"column:buyer_refund;type:tinyint(2);not null" json:"buyer_refund"`
	BuyerRefundExpire  int32  `gorm:"column:buyer_refund_expire;type:tinyint(2);not null" json:"buyer_refund_expire"`
	LimitSpecific      string `gorm:"column:limit_specific;type:tinyint(2);not null" json:"limit_specific"`
	IsShowRest         int8   `gorm:"column:is_show_rest;type:tinyint(2);not null" json:"is_show_rest"`
	SaleTime           int32  `gorm:"column:sale_time;type:tinyint(2);not null" json:"sale_time"`
	RetailId           int64  `gorm:"column:retail_id;type:tinyint(2);not null" json:"retail_id"`
	SnapshotId         int64  `gorm:"column:snapshot_id;type:tinyint(2);not null" json:"snapshot_id"`
}

func (m *CoShopGoods) TableName() string {
	return "shop_goods"
}

type CoShopGoodsData struct {
	ID           int64  `gorm:"column:id;type:bigint(17);not null" json:"id"`                         // 商品ID
	Pic          string `gorm:"column:pic;type:varchar(255);not null" json:"pic"`                     // 商品图  JSON数组对应图片ID
	Video        string `gorm:"column:video;type:varchar(35);not null" json:"video"`                  // 主图视频
	Share        string `gorm:"column:share;type:varchar(255);not null" json:"share"`                 // 分享描述
	Word         string `gorm:"column:word;type:varchar(255);not null" json:"word"`                   // 用户留言数组 {名称,模型,必填} JSON数组 如 电话,phone,1
	SubName      string `gorm:"column:sub_name;type:varchar(20);not null" json:"sub_name"`            // 购买按钮名称
	Content      string `gorm:"column:content;not null" json:"content"`                               // 商品详情
	SpecialRules string `gorm:"column:special_rules;type:varchar(255);not null" json:"special_rules"` // 电子卡券特有的属性 电子卡券有效期和退款
}

func (m *CoShopGoodsData) TableName() string {
	return "shop_goods_data"
}

type CoShopGoodsRule struct {
	modelSql.Database
	GoodsID       int64  `gorm:"column:goods_id;type:bigint(17);not null" json:"goods_id"`  // 商品ID
	SkuID         int64  `gorm:"column:sku_id;type:bigint(17);not null" json:"sku_id"`      // 商品ID
	StockID       int64  `gorm:"column:stock_id;type:bigint(17);not null" json:"stock_id"`  // 关联库存ID
	StockNum      int32  `gorm:"column:stock_num;type:int(10);not null" json:"stock_num"`   // 库存数量
	SoldNum       int32  `gorm:"column:sold_num;type:int(10);not null" json:"sold_num"`     // 库存数量
	ActStart      int32  `gorm:"column:act_start;type:int(10);not null" json:"act_start"`   // 规则开始时间 0总是生效
	ActEnd        int32  `gorm:"column:act_end;type:int(10);not null" json:"act_end"`       // 规则结束时间
	Price         int32  `gorm:"column:price;not null" json:"price"`                        // 规则价格
	PriceCost     int32  `gorm:"column:price_cost;not null" json:"price_cost"`              // 成本价
	PriceDot      int32  `gorm:"column:price_dot;not null" json:"price_dot"`                // 成本价
	GoodsNo       string `gorm:"column:goods_no;type:varchar(20);not null" json:"goods_no"` // 商品编码
	Specs         string `gorm:"column:specs" json:"specs"`                                 // 备用 JSON
	GoodsDiscount string `gorm:"column:goods_discount" json:"goods_discount"`               // 打折信息 JSON
	GoodsInfo     string `gorm:"column:goods_info" json:"goods_info"`                       // 规则数据 JSON
	Sort          int    `gorm:"column:sort;type:int(11);not null" json:"sort"`             // 排序
}
func (m *CoShopGoodsRule) TableName() string {
	return "shop_goods_rule"
}
