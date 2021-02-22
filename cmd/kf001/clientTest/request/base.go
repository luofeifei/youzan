package request

type UserRegionLngLat struct {
	Lng float64 `binding:"required,max=9999" bson:"lng" json:"lng"`  // 纬度
	Lat float64 `binding:"required,max=9999" bson:"lat"  json:"lat"` // 经度
}

// 企业配置数据规则
type CoSysSetting struct {
	ID      int64            `json:"id"`                                               // 企业积分规则ID 企业ID
	MenuID  int64            `binding:"required" json:"menu_id"`                       // 所属菜单ID
	Name    string           `binding:"required,max=50" json:"name" maxLength:"50"`    // 配置名称
	AddInit bool             `json:"add_init" example:"true"`                          // 创建时是否初始化企业配置
	Rule    []CoSysLimitRule `binding:"required" json:"rule,omitempty"`                // 配置数据规则
	State   int8             `binding:"required,oneof=1 2 3" json:"state" example:"1"` // 状态 1正常 2下线 3开发中
}

type CoSysSettingPage struct {
	PageInfo
	MenuID int64  `json:"menu_id"`
	Name   string `json:"name"`
}

type CoSysLimitRule struct {
	Name  string        `binding:"required" json:"name" maxLength:"50"`           // 定义规则名称
	Label string        `binding:"required,alphanum" json:"label" maxLength:"20"` // 定义标签名
	Type  int8          `binding:"required,oneof=1 2 3 4 5 6" json:"type"`        // 值类型 1数字 2字符串 3指定值 4布尔 5JSON 6富文本
	Limit []interface{} `json:"limit"`                                            // 参数限制 1最小最大限制 2 最长最短字符限制 3只能是其中之一
	Value interface{}   `json:"value"`                                            // 参数默认值
	Rule  interface{}   `binding:"json" json:"rule"`                              // 参数定义值 JSON数组
}
