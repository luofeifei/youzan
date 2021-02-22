package mongoSql

// sys_config 系统配置值
type SysConfig struct {
	ID          int64                  `bson:"_id" json:"id"`                    // 配置ID
	CreatedAt   int64                  `bson:"created_at" json:"created_at"`     // 添加时间
	UpdatedAt   int64                  `bson:"updated_at" json:"updated_at"`     // 更新时间
	Name        string                 `bson:"name" json:"name"`                 // 配置名称
	Alias       string                 `bson:"alias" json:"alias"`               // 配置别名
	LimitAdmin  []int64                `bson:"limit_admin" json:"limit_admin"`   // 限制可操作管理员
	ConfigRule  []CoLimitRule          `bson:"config_rule" json:"config_rule"`   // 用户配置对应名称
	ConfigValue map[string]interface{} `bson:"config_value" json:"config_value"` // 用户配置对应值
}

func (m *SysConfig) TableName() string {
	return "sys_config"
}

// UserRegion 区域
type UserRegion struct {
	ID       int32     `bson:"_id" json:"id"`              // 地区编号
	ParentID int32     `bson:"parent_id" json:"parent_id"` // 父地区
	Name     string    `bson:"name" json:"name"`           // 名称
	Loc      []float64 `bson:"loc" json:"loc"`             // 纬度 经度
}

func (m *UserRegion) TableName() string {
	return "user_region"
}

// 二维码登录扫码
type UserLoginQr struct {
	ID        int64  `bson:"_id" json:"_id"`               // 二维码扫描登录
	Code      string `bson:"code" json:"code"`             // 条码内容
	UId       int64  `bson:"uid" json:"uid"`               // 扫码用户
	State     int8   `bson:"state" json:"state"`           // 状态 1等待扫描 2扫描确认 3已确认
	CreatedAt int64  `bson:"created_at" json:"created_at"` // 创建时间
	UpdatedAt int64  `bson:"updated_at" json:"updated_at"` // 更新时间
}

func (m *UserLoginQr) TableName() string {
	return "user_login_qr"
}
