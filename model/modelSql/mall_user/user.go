package mall_user

import "base/model/modelSql"

// User [...]
type User struct {
	modelSql.Database
	Pass    string `gorm:"column:pass;type:char(32);not null" json:"pass"`         // 密码
	Encrypt string `gorm:"column:encrypt;type:char(6);not null" json:"encrypt"`    // 签名密钥
	IsReal  int8   `gorm:"column:is_real;type:tinyint(2);not null" json:"is_real"` // 实名状态 1已实名 2未实名
	State   int8   `gorm:"column:state;type:tinyint(2);not null" json:"state"`     // 状态
}

func (m *User) TableName() string {
	return "user"
}

// UserAssets [...]
type UserAssets struct {
	ID    int64   `gorm:"primary_key;column:id;type:bigint(17);not null" json:"id"` // 用户资金表
	Money float64 `gorm:"column:money;type:decimal(10,2);not null" json:"money"`    // 账户余额
}

func (m *UserAssets) TableName() string {
	return "user_assets"
}

// UserData [...]
type UserData struct {
	ID       int64  `gorm:"primary_key;column:id;type:bigint(17);not null" json:"id"`
	Avatar   string `gorm:"column:avatar;type:varchar(35);not null" json:"avatar"`       // 用户图像
	Sex      bool   `gorm:"column:sex;type:tinyint(1);not null" json:"sex"`              // 性别
	Source   string `gorm:"column:source;type:char(5);not null" json:"source"`           // 注册来源 1 H5 2 WEB 3 APP 4 微信公众号 5 微信小程序 6 支付宝 7 字节跳动
	NickName string `gorm:"column:nick_name;type:varchar(30);not null" json:"nick_name"` // 昵称
	RealName string `gorm:"column:real_name;type:varchar(30);not null" json:"real_name"` // 真实姓名
	IDCard   string `gorm:"column:id_card;type:varchar(30);not null" json:"id_card"`     // 证件号码
	RealType int8   `gorm:"column:real_type;type:tinyint(2);not null" json:"real_type"`  // 实名认证来源
}

func (m *UserData) TableName() string {
	return "user_data"
}

// 用户账户列表
type UserList struct {
	modelSql.Database
	UId  int64  `gorm:"index:uid;column:uid;type:bigint(17);not null" json:"uid"`     // 所属用户 ID
	Type int8   `gorm:"index:user;column:type;type:tinyint(2);not null" json:"type"`  // 用户类型 0系统ID 1用户名 2手机号 3微信 4支付宝 5字节跳动
	User string `gorm:"index:user;column:user;type:varchar(36);not null" json:"user"` // 用户名
}

func (m *UserList) TableName() string {
	return "user_list"
}

// UserLogin [...]
type UserLogin struct {
	modelSql.Database
	UId      int64  `gorm:"index:uid;column:uid;type:bigint(17);not null" json:"uid"`        // 用户ID
	Device   string `gorm:"column:device;type:varchar(50);not null" json:"device"`           // 设备标识
	Platform string `gorm:"index:uid;column:platform;type:char(5);not null" json:"platform"` // 登录来源
	IP       string `gorm:"column:ip;type:char(12);not null" json:"ip"`                      // IP
	Token    string `gorm:"column:token;type:text" json:"token"`                             // 记录 JWF
}

func (m *UserLogin) TableName() string {
	return "user_login"
}

// 用户邮寄地址
type UserMailing struct {
	modelSql.Database
	Default  int8   `gorm:"column:default;type:tinyint(2);not null" json:"default"`   // 是否为默认
	Province int    `gorm:"column:province;type:int(10);not null" json:"province"`    // 所在省
	City     int    `gorm:"column:city;type:int(10);not null" json:"city"`            // 所在市
	County   int    `gorm:"column:county;type:int(10);not null" json:"county"`        // 所在县
	Address  string `gorm:"column:address;type:varchar(200);not null" json:"address"` // 详细地址
	Name     string `gorm:"column:name;type:varchar(10);not null" json:"name"`        // 收货人姓名
	Mobile   string `gorm:"column:mobile;type:varchar(15);not null" json:"mobile"`    // 联系电话
}

func (m *UserMailing) TableName() string {
	return "user_mailing"
}
