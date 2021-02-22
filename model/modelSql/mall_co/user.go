package mall_co

import "base/model/modelSql"

type CoUser struct {
	modelSql.Database
	Name     string `gorm:"column:name;type:varchar(100);not null" json:"name"`    // 企业名称
	Logo     string `gorm:"column:logo;type:varchar(100);not null" json:"logo"`    // 企业LOGO
	Chain    int8   `gorm:"column:chain;type:tinyint(2);not null" json:"chain"`    // 连锁点模式 1关闭 2主店 3子店
	Province int    `gorm:"column:province;type:int(10);not null" json:"province"` // 所在省
	City     int    `gorm:"column:city;type:int(10);not null" json:"city"`         // 所在市
	County   int    `gorm:"column:county;type:int(10);not null" json:"county"`     // 所在县
}

func (m *CoUser) TableName() string {
	return "co_user"
}

// 企业证书材料及申请
type CoUserCert struct {
	modelSql.Database
	TypeID    int64  `gorm:"index:type_id;column:type_id;type:bigint(17);not null" json:"type_id"`       // 企业类型ID
	TypeClass int64  `gorm:"index:type_id;column:type_class;type:bigint(17);not null" json:"type_class"` // 企业所属类目
	Type      int8   `gorm:"index:type;column:type;type:tinyint(2);not null" json:"type"`                // 类型 1企业 2个体户 3政府及事业单位 4其他组织 5个人
	UId       int64  `gorm:"column:uid;type:bigint(17);not null" json:"uid"`                             // 申请人UID
	Name      string `gorm:"column:name;type:varchar(100);not null" json:"name"`                         // 企业名称
	RegionID  int    `gorm:"column:region_id;type:int(10);not null" json:"region_id"`                    // 所在地区编号
	Chain     int8   `gorm:"column:chain;type:tinyint(2);not null" json:"chain"`                         // 是否为连锁店模式
	CardID    string `gorm:"column:card_id;type:varchar(50);not null" json:"card_id"`                    // 证件号 身份证、三证/五证合一营业执照
	CardPros  string `gorm:"column:card_pros;type:varchar(35);not null" json:"card_pros"`                // 证件图片1 正面
	CardCons  string `gorm:"column:card_cons;type:varchar(35);not null" json:"card_cons"`                // 证件图片2 反面
	Msg       string `gorm:"column:msg;type:varchar(50);not null" json:"msg"`                            // 管理员处理信息
	State     int8   `gorm:"index:state;column:state;type:tinyint(2);not null" json:"state"`             // 1 正常 2申请中 3申请失败
}

func (m *CoUserCert) TableName() string {
	return "co_user_cert"
}

// 连锁店关联 到期解除绑定关系
type CoUserChain struct {
	modelSql.Database
	Coid  int64 `gorm:"index:coid;column:coid;type:bigint(17);not null" json:"coid"`    // 主店企业ID
	Sonid int64 `gorm:"index:sonid;column:sonid;type:bigint(17);not null" json:"sonid"` // 绑定子店ID
	State int8  `gorm:"column:state;type:tinyint(2);not null" json:"state"`             // 1 正常 2邀请中 3暂停
}

func (m *CoUserChain) TableName() string {
	return "co_user_chain"
}

// 企业用户详细数据
type CoUserData struct {
	ID         int64  `gorm:"column:id;type:bigint(17);not null" json:"id" form:"id"`                            // 企业ID
	Address    string `gorm:"column:address;type:varchar(100);not null" json:"address" form:"address"`           // 企业所在地址
	CertID     int64  `gorm:"column:cert_id;type:bigint(17);not null" json:"cert_id" form:"cert_id"`             // 主体认证信息关联ID
	ChainMax   int8   `gorm:"column:chain_max;type:tinyint(4);not null" json:"chain_max" form:"chain_max"`       // 最大可关联连锁店数
	Phone      string `gorm:"column:phone;type:varchar(15);not null" json:"phone" form:"phone"`                  // 联系电话
	Qq         string `gorm:"column:qq;type:varchar(15);not null" json:"qq" form:"qq"`                           // 联系QQ
	WeChat     string `gorm:"column:we_chat;type:varchar(20);not null" json:"we_chat" form:"we_chat"`            // 联系微信
	Intro      string `gorm:"column:intro;type:varchar(100);not null" json:"intro" form:"intro"`                 // 企业简介
	CreatedUid int64  `gorm:"column:created_uid;type:bigint(17);not null" json:"created_uid" form:"created_uid"` // 创建用户 ID
}

func (m *CoUserData) TableName() string {
	return "co_user_data"
}

// 企业用户地址库
type CoUserAddress struct {
	modelSql.Database
	Coid     int64               `gorm:"column:coid;type:bigint(17);not null" json:"coid" form:"coid"`            // 企业ID
	Name     string              `gorm:"column:name;type:varchar(30);not null" json:"name" form:"name"`           // 联系人
	Phone    string              `gorm:"column:phone;type:varchar(15);not null" json:"phone" form:"phone"`        // 联系方式
	Province int                 `gorm:"column:province;type:int(10);not null" json:"province" form:"province"`   // 所在省
	City     int                 `gorm:"column:city;type:int(10);not null" json:"city" form:"city"`               // 所在市
	County   int                 `gorm:"column:county;type:int(10);not null" json:"county" form:"county"`         // 所在县
	Address  string              `gorm:"column:address;type:varchar(100);not null" json:"address" form:"address"` // 详细地址
	Link     []CoUserAddressLink `gorm:"ForeignKey:address_id"`
}

func (m *CoUserAddress) TableName() string {
	return "co_user_address"
}

type CoUserAddressLink struct {
	AddressID int64 `gorm:"column:address_id;type:bigint(17);not null" json:"address_id" form:"address_id"`  // 企业用户地址库ID
	Type      int8  `gorm:"column:type;type:tinyint(2);not null" json:"type" form:"type"`                    // 地址类型 1退货地址 2收票地址 3发货地址
	Default   int8  `gorm:"default:1;column:default;type:tinyint(2);not null" json:"default" form:"default"` // 是否默认 2默认
}

func (m *CoUserAddressLink) TableName() string {
	return "co_user_address_link"
}

// 企业权益
type CoUserBenefit struct {
	modelSql.Database
	Coid            int64   `gorm:"column:coid;type:bigint(17);not null" json:"coid" form:"coid"`                                     // 企业ID
	Name            string  `gorm:"column:name;type:varchar(50);not null" json:"name" form:"name"`                                    // 权益名称
	ColorCode       string  `gorm:"column:color_code;type:varchar(10);not null" json:"color_code" form:"color_code"`                  // 背景设置-背景色
	CoverUrl        string  `gorm:"column:cover_url;type:varchar(50);not null" json:"cover_url" form:"cover_url"`                     // 背景设置-背景图
	ClaimType       int8    `gorm:"column:claim_type;type:tinyint(2);not null" json:"claim_type" form:"claim_type"`                   // 领取设置 1可直接领取 2满足条件领取 3需付费购买
	Price           float64 `gorm:"column:price;not null" json:"price" form:"price"`                                                  // 购买价格
	StockNum        int32   `gorm:"column:stock_num;type:mediumint(9);not null" json:"stock_num" form:"stock_num"`                    // 库存数量
	TermType        int8    `gorm:"column:term_type;type:tinyint(2);not null" json:"term_type" form:"term_type"`                      // 有效期 1永久有效 2领卡后指定天生效 3指定日期
	TermDays        int     `gorm:"column:term_days;type:int(5);not null" json:"term_days" form:"term_days"`                          // 指定生效天数
	TermBeginAt     int     `gorm:"column:term_begin_at;type:int(10);not null" json:"term_begin_at" form:"term_begin_at"`             // 指定生效时间
	TermEndAt       int     `gorm:"column:term_end_at;type:int(10);not null" json:"term_end_at" form:"term_end_at"`                   // 指定失效时间
	IsRepeatable    int8    `gorm:"column:is_repeatable;type:tinyint(2);not null" json:"is_repeatable" form:"is_repeatable"`          // 不限次数
	RepeatableLimit int8    `gorm:"column:repeatable_limit;type:tinyint(4);not null" json:"repeatable_limit" form:"repeatable_limit"` // 有效期内限领次数
	RequireMobile   int8    `gorm:"column:require_mobile;type:tinyint(1);not null" json:"require_mobile" form:"require_mobile"`       // 验证手机号
	RequireProfile  int8    `gorm:"column:require_profile;type:tinyint(1);not null" json:"require_profile" form:"require_profile"`    // 完善信息
	GoodsNo         string  `gorm:"column:goods_no;type:varchar(20);not null" json:"goods_no" form:"goods_no"`                        // 商品编码
	ServicePhone    string  `gorm:"column:service_phone;type:varchar(15);not null" json:"service_phone" form:"service_phone"`         // 客服电话
	Description     string  `gorm:"column:description;type:varchar(255);not null" json:"description" form:"description"`              // 使用须知
	State           int8    `gorm:"column:state;type:tinyint(1);not null" json:"state" form:"state"`                                  // 状态 1正常 2下架 3禁用
}

func (m *CoUserBenefit) TableName() string {
	return "co_user_benefit"
}

// 自定义权益
type CoUserBenefitDiy struct {
	modelSql.Database
	Coid        int64  `gorm:"column:coid;type:bigint(17);not null" json:"coid" form:"coid"`                       // 企业ID
	BenefitID   int64  `gorm:"column:benefit_id;type:bigint(17);not null" json:"benefit_id" form:"benefit_id"`     // 权益模板ID
	Type        int8   `gorm:"column:type;type:tinyint(2);not null" json:"type" form:"type"`                       // 权益类型 1权益 2礼包
	Mode        int8   `gorm:"column:mode;type:tinyint(2);not null" json:"mode" form:"mode"`                       // 服务模式 1系统核单 2商户线下核单
	Name        string `gorm:"column:name;type:varchar(50);not null" json:"name" form:"name"`                      // 展示名称
	Icon        string `gorm:"column:icon;type:varchar(35);not null" json:"icon" form:"icon"`                      // 权益图标
	Description string `gorm:"column:description;type:varchar(50);not null" json:"description" form:"description"` // 权益简介
	UseState    int8   `gorm:"column:use_state;type:tinyint(2);not null" json:"use_state" form:"use_state"`        // 使用状态
}

func (m *CoUserBenefitDiy) TableName() string {
	return "co_user_benefit_diy"
}

// 用户自定义角色
type CoUserGroupRole struct {
	modelSql.Database
	Coid   int64  `gorm:"column:coid;type:bigint(19);not null" json:"coid" form:"coid"`  // 企业ID
	Name   string `gorm:"column:name;type:varchar(10);not null" json:"name" form:"name"` // 角色名称
	Tips   string `gorm:"column:tips;type:varchar(50)" json:"tips" form:"tips"`          // 角色描述
	MenuID string `gorm:"column:menu_id;not null" json:"-" form:"menu_id"`               // 拥有的菜单JSON数据
}

func (m *CoUserGroupRole) TableName() string {
	return "co_user_group_role"
}

// 企业用户门店管理
type CoUserStore struct {
	modelSql.Database
	Coid         int64  `gorm:"column:coid;type:bigint(17);not null" json:"coid" form:"coid"`                            // 企业ID
	Name         string `gorm:"column:name;type:varchar(20);not null" json:"name" form:"name"`                           // 门店名称
	Phone        string `gorm:"column:phone;type:varchar(15);not null" json:"phone" form:"phone"`                        // 联系电话
	Pic          string `gorm:"column:pic;type:varchar(255);not null" json:"-" form:"pic"`                               // 门店照片 JSON数组对应图片ID
	Province     int    `gorm:"column:province;type:int(10);not null" json:"province" form:"province"`                   // 所在省
	City         int    `gorm:"column:city;type:int(10);not null" json:"city" form:"city"`                               // 所在市
	County       int    `gorm:"column:county;type:int(10);not null" json:"county" form:"county"`                         // 所在县
	Address      string `gorm:"column:address;type:varchar(100);not null" json:"address" form:"address"`                 // 详细地址
	Description  string `gorm:"column:description;type:varchar(200);not null" json:"description" form:"description"`     // 商家推荐介绍
	BusinessType int8   `gorm:"column:business_type;type:tinyint(2);not null" json:"business_type" form:"business_type"` // 营业时间 1全天 2每天重复 3每周重复
	BusinessTime string `gorm:"column:business_time;type:varchar(100);not null" json:"-" form:"business_time"`           // 营业不为全天 储存的时间规则JSON
}

func (m *CoUserStore) TableName() string {
	return "co_user_store"
}

// 企业绑定用户
type CoUserCustomer struct {
	modelSql.Database
	Coid      int64 `gorm:"index:eduid;column:coid;type:bigint(17);not null" json:"coid"` // 企业ID
	UId       int64 `gorm:"index:uid;column:uid;type:bigint(17);not null" json:"uid"`     // 关联用户ID
	Sort      int   `gorm:"index:uid;column:sort;type:int(11);not null" json:"sort"`      // 显示排序
	Points    int   `gorm:"column:points;type:int(10);not null" json:"points"`            // 用户积分
	PointsAll int   `gorm:"column:points_all;type:int(10);not null" json:"points_all"`    // 累计发放积分
}

func (m *CoUserCustomer) TableName() string {
	return "co_user_customer"
}

// 企业用户订单
type CoUserOrder struct {
	modelSql.Database
	Coid    int64   `gorm:"index:coid;column:coid;type:bigint(17);not null" json:"coid"`         // 企业ID
	BizType int8    `gorm:"index:biz_type;column:type;type:tinyint(2);not null" json:"biz_type"` // 类型
	Aid     int64   `gorm:"index:type;column:aid;type:bigint(17);not null" json:"aid"`           // 关联商品ID
	Price   float64 `gorm:"column:price;type:decimal(10,2);not null" json:"price"`               // 需支付价格
	State   int8    `gorm:"index:state;column:state;type:tinyint(2);not null" json:"state"`      // 状态 1待支付 2取消 3已支付 4已完成
	PayTime int     `gorm:"column:pay_time;type:int(10);not null" json:"pay_time"`               // 支付时间
}

func (m *CoUserOrder) TableName() string {
	return "co_user_order"
}

// 企业资金
type CoUserFund struct {
	Coid      int64 `gorm:"index:coid;column:coid;type:bigint(17);not null" json:"coid"` // 企业ID
	CreatedAt int   `gorm:"column:created_at;type:int(10);not null" json:"created_at"`   // 添加时间
	UpdatedAt int   `gorm:"column:updated_at;type:int(10);not null" json:"updated_at"`   // 更新时间
	Balance   int   `gorm:"column:balance;type:int(10);not null" json:"balance"`         // 企业余额
	Frozen    int   `gorm:"column:frozen;type:int(10);not null" json:"frozen"`           // 冻结资金
}
