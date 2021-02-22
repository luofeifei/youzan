package mongoSql

// 公用规则  -  不是表
type CoLimitRule struct {
	Name  string                 `bson:"name" json:"name"`             // 定义规则名称
	Label string                 `bson:"label" json:"label"`           // 定义标签名
	Type  int8                   `bson:"type" json:"type"`             // 值类型 1数字 2字符串 3指定值 4布尔 5JSON 6富文本
	Limit []interface{}          `bson:"limit" json:"limit,omitempty"` // 参数限制 1最小最大限制 2 最长最短字符限制 3只能是其中之一
	Value interface{}            `bson:"value" json:"value"`           // 参数默认值
	Rule  map[string]interface{} `bson:"rule" json:"rule,omitempty"`   // 参数定义值 JSON数组
}

// co_type 企业类型表
type CoType struct {
	ID    int64  `bson:"_id" json:"id"`      // 地区编号
	Pid   int64  `bson:"pid" json:"pid"`     // 父级
	Name  string `bson:"name" json:"name"`   // 企业类型名称
	State int8   `bson:"state" json:"state"` // 状态
}

func (m *CoType) TableName() string {
	return "co_type"
}

// co_menu 企业ID后台菜单
type CoMenu struct {
	ID        int64    `bson:"_id" json:"id"`                // 菜单ID
	Pid       int64    `bson:"pid" json:"pid"`               // 父级
	Type      int32    `bson:"type" json:"type"`             // 类型 1菜单 2路由 3权限
	Title     string   `bson:"title" json:"title"`           // 规则名称
	Path      string   `bson:"path" json:"path"`             // 访问地址
	Component string   `bson:"component" json:"component"`   // 模板地址
	Icon      string   `bson:"icon" json:"icon"`             // 图标
	IsDisplay bool     `bson:"is_display" json:"is_display"` // 无权限是否显示
	Sort      int      `bson:"sort" json:"sort"`             // 权重
	State     int32    `bson:"state" json:"state"`           // 状态 1正常 2禁用 3公开 4 没有权限
	Children  []CoMenu `bson:"-" json:"children"`
}

func (m *CoMenu) TableName() string {
	return "co_menu"
}

// co_menu_api 区域
type CoMenuAPI struct {
	ID          int64  `bson:"_id" json:"id"`                  // 地区编号
	MenuID      int64  `bson:"menu_id" json:"menu_id"`         // 所属菜单ID
	Method      string `bson:"method" json:"method"`           // 访问方式
	Path        string `bson:"path" json:"path"`               // 访问路径
	Description string `bson:"description" json:"description"` // 接口简介
}

func (m *CoMenuAPI) TableName() string {
	return "co_menu_api"
}

// co_menu_group 菜单权限组 用户可控制权限
type CoMenuGroup struct {
	ID         int64   `bson:"_id" json:"id"`                  // 地区编号
	Pid        int64   `json:"pid" bson:"pid"`                 // 父菜单组权限ID -1顶级
	RoleID     int64   `json:"role_id" bson:"role_id"`         // 权限角色标签ID
	Name       string  `json:"name" bson:"name"`               // 组名称
	Note       string  `json:"note" bson:"note"`               // 权限提示信息
	Chain      int8    `json:"chain" bson:"chain"`             // 是否为连锁店模式
	IsSys      bool    `json:"is_sys" bson:"is_sys"`           // 是否为公用角色权限组
	Level      int     `json:"level" bson:"level"`             // 必须大于权限等级几？才能使用
	MenuList   []int64 `bson:"menu_list" json:"menu_list"`     // 拥有的菜单权限
	PluginList []int64 `bson:"plugin_list" json:"plugin_list"` // 拥有的插件权限
	State      int8    `json:"state" bson:"state"`             // 状态
}

func (m *CoMenuGroup) TableName() string {
	return "co_menu_group"
}

// co_menu_group_limit 菜单权限组 规则限制
type CoMenuGroupLimit struct {
	ID          int64 `bson:"_id" json:"id"`                      // 地区编号
	LimitID     int64 `bson:"limit_id" json:"limit_id"`           // 限制 ID
	MenuGroupID int64 `bson:"menu_group_id" json:"menu_group_id"` // 菜单组 ID
	Level       int   `bson:"level" json:"level"`                 // 权限等级
	Value       int   `bson:"value" json:"value"`                 // 值
}

func (m *CoMenuGroupLimit) TableName() string {
	return "co_menu_group_limit"
}

// co_menu_group_role 菜单角色标签ID
type CoMenuGroupRole struct {
	ID     int64  `bson:"_id" json:"id"`        // 地区编号
	Name   string `json:"name" bson:"name"`     // 角色名
	Serial string `json:"serial" bson:"serial"` // 标签
	Tips   string `json:"tips" bson:"tips"`     // 提示信息
}

func (m *CoMenuGroupRole) TableName() string {
	return "co_menu_group_role"
}

// co_sys_vip 企业Vip
type CoSysVip struct {
	ID          int64                   `bson:"_id" json:"id,omitempty"`                      // VIP ID
	CreatedAt   int64                   `bson:"created_at" json:"created_at,omitempty"`       // 添加时间
	UpdatedAt   int64                   `bson:"updated_at" json:"updated_at,omitempty"`       // 更新时间
	Name        string                  `bson:"name" json:"name,omitempty"`                   // 商品名
	Title       string                  `bson:"title" json:"title,omitempty"`                 // 标题
	TypeID      []int64                 `bson:"type_id" json:"type_id,omitempty"`             // 所属企业分类数组
	Recommend   []int64                 `bson:"recommend" json:"recommend,omitempty"`         // 推荐VIP 产品列表
	PriceLabel  string                  `bson:"price_label" json:"price_label,omitempty"`     // 展示价格标签
	Cover       string                  `bson:"cover" form:"cover,omitempty"`                 // 封面图片
	Quota       int8                    `bson:"quota" json:"quota,omitempty"`                 // 限购 -1不限制 1终身 2每天 3每周 4每月 5每年
	QuotaNum    int8                    `bson:"quota_num" json:"quota_num,omitempty"`         // 购买限制单数 -1不限制
	Content     string                  `bson:"content" json:"content,omitempty"`             // 商品详情
	Rule        []CoSysVipRule          `bson:"rule" json:"rule,omitempty"`                   // 商品规则
	RuleVipList map[int8][]CoSysVipList `bson:"rule_vip_list" json:"rule_vip_list,omitempty"` // 商品规则 绑定的权限
	Sales       int32                   `bson:"sales" json:"sales,omitempty"`                 // 总销量
	Sort        int16                   `bson:"sort" json:"sort,omitempty"`                   // 排序
	State       int8                    `bson:"state" json:"state,omitempty"`                 // 状态 1 正常 2隐藏可通过推荐购买 3 下线
}

type CoSysVipRule struct {
	Name     string `bson:"name" json:"name"`           // 标题名
	Title    string `bson:"title" json:"title"`         // 副标题
	PriceDot int32  `bson:"price_dot" json:"price_dot"` // 原价 分
	Price    int32  `bson:"price" json:"price"`         // 现价 分
	Sales    int32  `bson:"sales" json:"sales"`         // 销量
}

type CoSysVipList struct {
	Label      string `bson:"label" json:"label"`             // VIP 权限标题
	Type       int8   `bson:"type" json:"type"`               // 类型 1 企业组权限时长 2 插件权限时长 3 企业限制规则
	Level      int8   `bson:"level" json:"level"`             // 类型 1 时为 会员等级 1~10
	LimitId    int64  `bson:"limit_id" json:"limit_id"`       // 根据类型决定 1 企业组ID 2 插件ID 3 限制规则ID
	LimitDays  int32  `bson:"limit_days" json:"limit_days"`   // 到期天数
	LimitValue int32  `bson:"limit_value" json:"limit_value"` // 限制值 1、2不填 3 限制规则的值
}

func (m *CoSysVip) TableName() string {
	return "co_sys_vip"
}

// co_sys_binding 企业绑定关系
type CoSysBinding struct {
	ID        int64   `bson:"_id" json:"id"`                // Id 用于绑定 各行业 拥有的所有功能列表
	Type      int32   `bson:"type" json:"type"`             // 类型 1菜单 2插件 3权限组
	TypeID    int64   `bson:"type_id" json:"type_id"`       // 所属企业分类ID
	List      []int64 `bson:"list" json:"list"`             // 绑定的ID 列表
	Count     int32   `bson:"count" json:"count"`           // 包含的ID 数量
	CreatedAt int64   `bson:"created_at" json:"created_at"` // 添加时间
	UpdatedAt int64   `bson:"updated_at" json:"updated_at"` // 更新时间
}

func (m *CoSysBinding) TableName() string {
	return "co_sys_binding"
}

// co_sys_binding_vip 企业绑定关系 - 无权限指定购买地址
type CoSysBindingVip struct {
	ID      int64   `bson:"_id" json:"id"`            // Id
	Type    int8    `bson:"type" json:"type"`         // 类型 1菜单 2插件
	TypeID  int64   `bson:"type_id" json:"type_id"`   // 所属企业分类ID
	LimitId int64   `bson:"limit_id" json:"limit_id"` // 列表中的ID
	VipId   []int64 `bson:"vip_id" json:"vip_id"`     // 绑定的VIP 购买地址
}

func (m *CoSysBindingVip) TableName() string {
	return "co_sys_binding_vip"
}

// co_sys_limit 企业限制规则
type CoSysLimit struct {
	ID        int64  `bson:"_id" json:"id"`                // 限制规则ID
	CreatedAt int64  `bson:"created_at" json:"created_at"` // 添加时间
	UpdatedAt int64  `bson:"updated_at" json:"updated_at"` // 更新时间
	Serial    string `bson:"serial" json:"serial"`         // 规则标签 开发用
	Name      string `bson:"name" json:"name"`             // 限制规则名称
	Interval  int    `bson:"interval" json:"interval"`     // 限制间隔分钟
	Unit      string `bson:"unit" json:"unit"`             // 单位
	Min       int    `bson:"min" json:"min"`               // 最小值
	Max       int    `bson:"max" json:"max"`               // 最大值
	Tips      string `bson:"tips" json:"tips"`             // 提示信息
	Note      string `json:"note" bson:"note"`             // 权限提示信息
}

func (m *CoSysLimit) TableName() string {
	return "co_sys_limit"
}

// co_sys_module 企业界面组件
type CoSysModule struct {
	ID         int64         `bson:"_id" json:"id,omitempty"`                  // 系统组件
	ClassID    int64         `bson:"class_id" json:"class_id,omitempty"`       // 所属组件分类
	PluginID   int64         `bson:"plugin_id" json:"plugin_id,omitempty"`     // 使用的插件ID -1 不使用插件 根据插件权限决定是否可用
	Name       string        `bson:"name" json:"name,omitempty"`               // 组件名称
	Icon       string        `bson:"icon" json:"icon,omitempty"`               // 组件图标
	Alias      string        `bson:"alias" json:"alias,omitempty"`             // 组件模型标签
	UseType    int8          `bson:"use_type" json:"use_type,omitempty"`       // 限制方式 1 不限制 2插件验权
	UsePos     []int8        `bson:"use_pos" json:"use_pos,omitempty"`         // 使用位置 -1 任意页面可调用 1微页面 2主页 3分类页 10导航条 11个人中心 12悬浮窗 13公共广告 ...
	LimitAdmin []int64       `bson:"limit_admin" json:"limit_admin"`           // 限制可操作管理员
	MaxNum     int8          `bson:"max_num" json:"max_num,omitempty"`         // 每页面最多可使用次数
	IsInit     bool          `bson:"is_init" json:"is_init,omitempty"`         // 是否初始化组件
	IsSort     bool          `bson:"is_sort" json:"is_sort,omitempty"`         // 是否支持排序
	FixSort    int8          `bson:"fix_sort" json:"fix_sort,omitempty"`       // 固定排序位置 序号 0~?
	IsDisplay  bool          `bson:"is_display" json:"is_display,omitempty"`   // 用户无权限是否显示
	IsSystem   bool          `bson:"is_system" json:"is_system,omitempty"`     // 是否为系统组件分类 如果为不显示
	IsDelete   bool          `bson:"is_delete" json:"is_delete,omitempty"`     // 是否支持删除
	IsConfig   bool          `bson:"is_config" json:"is_config,omitempty"`     // 是否支持功能设置
	State      int8          `bson:"state" json:"state,omitempty"`             // 状态 1正常 2下线 3开发中 4未开通权限
	ConfigRule []CoLimitRule `bson:"config_rule" json:"config_rule,omitempty"` // 参数设置规则
}

func (m *CoSysModule) TableName() string {
	return "co_sys_module"
}

// co_sys_module_class
type CoSysModuleClass struct {
	ID    int64  `bson:"_id" json:"id"`      // 系统组件
	Name  string `bson:"name" json:"name"`   // 组件类型名称
	State int8   `bson:"state" json:"state"` // 状态 1正常 2停用
}

func (m *CoSysModuleClass) TableName() string {
	return "co_sys_module_class"
}

// co_sys_template 企业模板市场数据
type CoSysTemplate struct {
	ID         int64  `bson:"_id" json:"id"`                  // 页面ID
	Coid       int64  `bson:"coid" json:"coid"`               // 企业ID
	Name       string `bson:"name" json:"name"`               // 标题
	Designer   string `bson:"designer" json:"designer"`       // 作者
	ExampleUrl string `bson:"example_url" json:"example_url"` // 模板访问地址
	PayType    int8   `bson:"pay_type" json:"pay_type"`       // 支付类型 1免费 2收费
	TplType    int8   `bson:"tpl_type" json:"tpl_type"`       // 模板类型 1官网 2用户
	Components string `bson:"components" json:"components"`   // 模板数据
	State      int8   `bson:"state" json:"state"`             // 状态 1正常 2停用
}

func (m *CoSysTemplate) TableName() string {
	return "co_sys_template"
}

// co_sys_plugin 企业插件
type CoSysPlugin struct {
	ID         int64   `bson:"_id" json:"id"`                  // 插件编号
	Group      string  `bson:"group" json:"group"`             // 插件组名
	Type       int8    `bson:"type" json:"type"`               // 插件类型 1 空插件用于插件权限检测 2 so插件 3 gRpc中转
	Model      int8    `bson:"model" json:"model"`             // 插件模式 1 接口服务 2 工具过滤 3 队列处理 （一般页面为读取接口数据）
	UseType    int8    `bson:"use_type" json:"use_type"`       // 使用类型 -1 空 1 领卡时 2 购买前 3 购买后 （2才有改方法）
	LimitAdmin []int64 `bson:"limit_admin" json:"limit_admin"` // 限制可操作管理员
	Name       string  `bson:"name" json:"name"`               // 插件名称
	Note       string  `json:"note" bson:"note"`               // 权限提示信息
	Alias      string  `bson:"alias" json:"alias"`             // 插件别名
	Ver        string  `bson:"ver" json:"ver"`                 // 插件版本
	Topic      string  `bson:"topic" json:"topic"`             // 订阅消费者话题 话题:channel组
	Path       string  `bson:"path" json:"path"`               // 插件路径 SO文件上传后的路径
	State      int8    `bson:"state" json:"state"`             // 状态 1正常 2异常
}

func (m *CoSysPlugin) TableName() string {
	return "co_sys_plugin"
}

// co_sys_benefit 企业权益模板
type CoSysBenefit struct {
	ID          int64         `bson:"_id" json:"id,omitempty"`                  // 企业权益模板 企业ID
	ClassID     int64         `bson:"class_id" json:"class_id,omitempty"`       // 所属权益分类
	PluginID    int64         `bson:"plugin_id" json:"plugin_id,omitempty"`     // 使用的插件ID 根据插件权限决定是否可用
	Name        string        `bson:"name" json:"name,omitempty"`               // 权益模板名称
	Type        int8          `bson:"type" json:"type,omitempty"`               // 权益类型 1权益 2礼包
	Mode        int8          `bson:"mode" json:"mode,omitempty"`               // 服务模式 1系统核单 2商户线下核单
	Icon        string        `bson:"icon" json:"icon,omitempty"`               // 默认权益图标
	Description string        `bson:"description" json:"description,omitempty"` // 权益简介
	State       int8          `bson:"state" json:"state,omitempty"`             // 状态 1已生效 2预上线 3已下线
	Parameter   []CoLimitRule `bson:"parameter" json:"parameter,omitempty"`     // 参数定义值
}

func (m *CoSysBenefit) TableName() string {
	return "co_sys_benefit"
}

// co_sys_benefit_class 企业权益分类
type CoSysBenefitClass struct {
	ID       int64  `bson:"_id" json:"id"`                // 企业权益模板 企业ID
	Name     string `bson:"name" json:"name"`             // 权益分类名称
	UseCount int    `bson:"use_count" json:"use_count"`   // 被使用数
	State    int8   `bson:"state" json:"state,omitempty"` // 状态 1正常 2停用
}

func (m *CoSysBenefitClass) TableName() string {
	return "co_sys_benefit_class"
}

// 企业配置数据规则
type CoSysSetting struct {
	ID        int64         `bson:"_id" json:"id"`                          // 企业配置ID
	CreatedAt int64         `bson:"created_at" json:"created_at,omitempty"` // 添加时间
	UpdatedAt int64         `bson:"updated_at" json:"updated_at,omitempty"` // 更新时间
	MenuID    int64         `bson:"menu_id" json:"menu_id,omitempty"`       // 所属菜单ID
	Name      string        `bson:"name" json:"name,omitempty"`             // 配置名称
	AddInit   bool          `bson:"add_init" json:"add_init"`               // 创建时否初始化企业配置
	Rule      []CoLimitRule `bson:"rule" json:"rule,omitempty"`             // 配置数据规则
	State     int8          `bson:"state" json:"state,omitempty"`           // 状态 1正常 2下线 3开发中
}

func (m *CoSysSetting) TableName() string {
	return "co_sys_setting"
}

// co_user 企业用户 信息
type CoUser struct {
	ID         int64   `bson:"_id" json:"id"`                  // 所属企业ID
	TypeIds    []int64 `bson:"type_ids" json:"type_ids"`       // 所属企业分类ID
	CreatedUid int64   `bson:"created_uid" json:"created_uid"` // 企业最高管理员
	CertType   int8    `bson:"cert_type" json:"cert_type"`     // 主体认证状态 1未认证 2已认证
	Domain     string  `bson:"domain" json:"domain"`           // 绑定域名 未确定是否开放
	Style      string  `bson:"style" json:"style"`             // 全站配色风格
	State      int8    `bson:"state" json:"state"`             // 状态 1 正常 2已到期 3禁用
}

func (m *CoUser) TableName() string {
	return "co_user"
}

// co_user_menu 企业用户 菜单信息 缓存
type CoUserAuthority struct {
	ID         int64           `bson:"_id" json:"id"`                    // 所属企业ID
	MenuGroup  []int64         `bson:"menu_group" json:"menu_group"`     // 拥有的菜单组ID
	MenuIds    []int64         `bson:"menu_ids" json:"menu_ids"`         // 拥有权限的菜单ID
	AllMenuIds []int64         `bson:"all_menu_ids" json:"all_menu_ids"` // 拥有的所有菜单ID
	PluginIds  []int64         `bson:"plugin_ids" json:"plugin_ids"`     // 拥有的所有插件ID
	LimitIds   map[int64]int32 `bson:"limit_ids" json:"limit_ids"`       // 拥有的企业限制规则
	IsChange   int8            `bson:"is_change" json:"is_change"`       // 权限变更 1 是
}

func (m *CoUserAuthority) TableName() string {
	return "co_user_authority"
}

// 企业员工、管理员
type CoUserStaff struct {
	ID       int64              `bson:"_id" json:"id"`              // ID
	Coid     int64              `bson:"coid" json:"coid"`           // 所属企业ID
	Uid      int64              `bson:"uid" json:"uid"`             // 所属用户
	StaffNo  string             `bson:"staff_no" json:"staff_no"`   // 员工编号
	RealName string             `bson:"real_name" json:"real_name"` // 员工姓名
	Phone    string             `bson:"phone" json:"phone"`         // 员工联系电话
	Sort     int32              `bson:"sort" json:"sort"`           // 显示排序
	IsChange int8               `bson:"is_change" json:"is_change"` // 权限变更 1 是
	MenuList []int64            `bson:"menu_list" json:"menu_list"` // 拥有的菜单权限 缓存
	Group    []CoUserStaffGroup `bson:"group" json:"group"`         // 角色组列表
	State    int8               `bson:"state" json:"state"`         // 状态 1 正常 2邀请中 3停用
}

// 企业员工 拥有的角色组
type CoUserStaffGroup struct {
	RoleType int8  `bson:"role_type" json:"role_type"` // 角色组类型 1系统 2用户自定义
	RoleID   int64 `bson:"role_id" json:"role_id"`     // 角色组关联ID
}

func (m *CoUserStaff) TableName() string {
	return "co_user_staff"
}

// 企业用户配置
type CoUserSetting struct {
	ID        int64                  `bson:"_id" json:"id"`                // 企业积分规则ID 企业ID
	CreatedAt int64                  `bson:"created_at" json:"created_at"` // 添加时间
	UpdatedAt int64                  `bson:"updated_at" json:"updated_at"` // 更新时间
	Coid      int64                  `bson:"coid" json:"coid"`             // 企业ID
	SettingId int64                  `bson:"setting_id" json:"setting_id"` // 企业配置参数ID
	Value     map[string]interface{} `bson:"value" json:"value"`           // 用户配置存储值
}

func (m *CoUserSetting) TableName() string {
	return "co_user_setting"
}

// 企业限制记录
type CoUserLimit struct {
	ID       int64 `bson:"_id" json:"id"`              // 企业限制 ID
	Coid     int64 `bson:"coid" json:"coid"`           // 企业ID
	LimitID  int64 `bson:"limit_id" json:"limit_id"`   // 企业限制规则ID
	ActStart int   `bson:"act_start" json:"act_start"` // 限制开始时间
	ActEnd   int   `bson:"act_end" json:"act_end"`     // 限制清理时间 定时器 清理
	Value    int   `bson:"value" json:"value"`         // 当前值 大于 0 进行清理
}

func (m *CoUserLimit) TableName() string {
	return "co_user_limit"
}

// 企业权益 绑定数据
type CoUserBenefitBag struct {
	ID         int64                  `bson:"_id" json:"id"`                  // 企业权益ID
	Coid       int64                  `bson:"coid" json:"coid"`               // 企业ID
	PluginID   map[int64]int64        `bson:"plugin_id" json:"plugin_id"`     // 使用的插件ID 与关联的用户的权益ID
	MeetBag    []CoUserBenefitMeetBag `bson:"meet_bag" json:"meet_bag"`       // 满足领取条件值数据
	BenefitBag []CoUserBenefitBagList `bson:"benefit_bag" json:"benefit_bag"` // 权益列表及设置
}

type CoUserBenefitMeetBag struct {
	MeetType  int8  `bson:"meet_type" json:"meet_type"`   // 满足领取条件类型 1累计支付成功 2累计消费金额 3累计总积分为
	MeetValue int32 `bson:"meet_value" json:"meet_value"` // 满足条件值
}

type CoUserBenefitBagList struct {
	ID    int64                  `bson:"_id" json:"id"` // 用户自定义模板ID
	Value map[string]interface{} `bson:"value"  json:"value"`
}

func (m *CoUserBenefitBag) TableName() string {
	return "co_user_benefit_bag"
}

// 企业积分配置
type CoUserPoints struct {
	ID            int64  `bson:"_id" json:"coid"`                      // 企业积分规则ID 企业ID
	Name          string `bson:"name" json:"name"`                     // 积分自定义名称
	LimitPerDay   int32  `bson:"limit_per_day" json:"limit_per_day"`   // 积分获取上限
	ProtectedTime int16  `bson:"protected_time" json:"protected_time"` // 积分保护期 -1不启用 单位天
	GeneralRatio  int32  `bson:"general_ratio" json:"general_ratio"`   // 积分兑换比例 ?积分=1元
	Deduct        struct {
		DeductAmountLimitVal int32 `bson:"deduct_amount_limit" json:"deduct_amount_limit"` // 订单金额门槛 -1 不限制 订单最低为?元可抵现 单位分
		MinOrderAmount       int32 `bson:"min_order_amount" json:"min_order_amount"`       // 抵现金额上限 -1 不限制 每笔订单最多抵扣?元 单位分
	} `bson:"deduct" json:"deduct"`                                       // 积分抵现规则
	ExpireDetail ExpireDetail `bson:"expire_detail" json:"expire_detail"` // 积分到期规则
}

type ExpireDetail struct {
	ExpireType int8   `bson:"expire_type" json:"expire_type"` // 到期类型 1永久有效 2从获得开始 3每笔积分有效期
	TimePeriod string `bson:"time_period" json:"time_period"` // 2 指定失效日期 0101 1月1日
	TimeDay    int16  `bson:"time_day" json:"time_day"`       // 3 有效期天 一年365天
}

func (m *CoUserPoints) TableName() string {
	return "co_user_points"
}

// 企业积分规则
type CoUserPointsRule struct {
	ID          int64 `bson:"_id" json:"id"`                    // 企业积分规则ID 企业ID
	CreatedAt   int64 `bson:"created_at" json:"created_at"`     // 添加时间
	UpdatedAt   int64 `bson:"updated_at" json:"updated_at"`     // 更新时间
	Coid        int64 `bson:"coid"  json:"coid"`                // 企业ID
	Points      int32 `bson:"points" json:"points"`             // 奖励分值
	RuleType    int8  `bson:"rule_type" json:"rule_type"`       // 规则类型 1 关注我的微信 2每成功交易 3每购买金额
	PointsLimit int32 `bson:"points_limit" json:"points_limit"` // 规则限制 2 每成功交易(笔) 3每购买金额(元)
	SendMessage bool  `bson:"send_message" json:"send_message"` // 通知设置 1不通知 2通知
	BonusPoints int64 `bson:"bonus_points" json:"bonus_points"` // 已奖励积分数
}

func (m *CoUserPointsRule) TableName() string {
	return "co_user_points_rule"
}

// co_user_page 企业用户 页面数据
type CoUserFeature struct {
	ID         int64                     `bson:"_id" json:"id"`                // 页面ID
	CreatedAt  int64                     `bson:"created_at" json:"created_at"` // 添加时间
	UpdatedAt  int64                     `bson:"updated_at" json:"updated_at"` // 更新时间
	Coid       int64                     `bson:"coid" json:"coid"`             // 企业ID
	Name       string                    `bson:"name" json:"name"`             // 标题
	Platform   int8                      `bson:"platform" json:"platform"`     // 平台标识ID -1通用 1 H5 2 WEB 3 APP 4 微信公众号 5 微信小程序 6 支付宝 7 字节跳动
	Source     int8                      `bson:"source" json:"source"`         // 来源 1 H5 2 WEB 3 APP 4 微信公众号 5 微信小程序 6 支付宝 7 字节跳动
	IsDelete   bool                      `bson:"is_delete" json:"is_delete"`   // 是否可删除
	PageType   int8                      `bson:"page_type" json:"page_type"`   // 页面类型 1微页面 2主页 3分类页
	State      int8                      `bson:"state" json:"state"`           // 状态 1正常 2草稿箱
	ApiPlugin  []CoUserFeaturePluginIDs  `bson:"api_plugin" json:"api_plugin"` // 使用的接口插件数据列表
	Components []CoUserFeatureComponents `bson:"components" json:"components"` // 页面使用的组件数据 JSON
}

type CoUserFeaturePluginIDs struct {
	I           int                    `bson:"i" json:"i"`                       // 界面组件ID
	Alias       string                 `bson:"alias" json:"alias"`               // 组件序 ID 前端生成短ID
	PluginID    int64                  `bson:"plugin_id" json:"plugin_id"`       // 使用的插件ID
	ModuleAlias string                 `bson:"module_alias" json:"module_alias"` // 组件模型标签
	Value       map[string]interface{} `bson:"value" json:"value"`               // 组件配置值
}

type CoUserFeatureComponents struct {
	Alias       string                 `bson:"alias" json:"alias"`               // 组件序 ID 前端生成短ID
	ModuleId    int64                  `bson:"module_id" json:"module_id"`       // 企业界面组件 ID
	ModuleAlias string                 `bson:"module_alias" json:"module_alias"` // 组件模型标签
	Value       map[string]interface{} `bson:"value" json:"value"`               // 组件配置值
}

func (m *CoUserFeature) TableName() string {
	return "co_user_shop_feature"
}

// 企业用户 店铺全局配置
type CoUserShopCommon struct {
	ID         int64                     `bson:"_id" json:"id"`                          // 页面ID
	CreatedAt  int64                     `bson:"created_at" json:"created_at"`           // 添加时间
	UpdatedAt  int64                     `bson:"updated_at" json:"updated_at"`           // 更新时间
	Coid       int64                     `bson:"coid" json:"coid,omitempty"`             // 企业ID
	PageType   int8                      `bson:"page_type" json:"page_type"`             // 全局配置类型 10导航条 11个人中心 12悬浮窗 13公共广告
	Platform   int8                      `bson:"platform" json:"platform"`               // 使用平台标识ID -1通用 1 H5 2 WEB 3 APP 4 微信公众号 5 微信小程序 6 支付宝 7 字节跳动
	ApiPlugin  []CoUserFeaturePluginIDs  `bson:"api_plugin" json:"api_plugin,omitempty"` // 使用的接口插件数据列表
	Components []CoUserFeatureComponents `bson:"components" json:"components"`           // 页面使用的配给或组件数据
}

func (m *CoUserShopCommon) TableName() string {
	return "co_user_shop_common"
}

// 企业用户关联权限
type CoUserVip struct {
	ID         int64 `bson:"_id" json:"id"`                          // 用户拥有的权限 Id
	CreatedAt  int64 `bson:"created_at" json:"created_at,omitempty"` // 添加时间
	UpdatedAt  int64 `bson:"updated_at" json:"updated_at,omitempty"` // 更新时间
	Coid       int64 `bson:"coid" json:"coid"`                       // 所属企业ID
	Type       int32 `bson:"type" json:"type"`                       // 类型 1 企业组权限时长 2 插件权限时长 3 企业限制规则时长
	PayType    int8  `bson:"pay_type" json:"pay_type"`               // 购买方式 1系统赠送 2购买 3积分兑换
	LimitId    int64 `bson:"limit_id" json:"limit_id"`               // 根据类型决定 1 企业组ID 2 插件ID 3 限制规则ID
	LimitValue int32 `bson:"limit_value" json:"limit_value"`         // 限制值 1 会员等级 1~10 2不填 3 限制规则的值
	DueDate    int64 `bson:"due_date" json:"due_date"`               // 到期时间
}

func (m *CoUserVip) TableName() string {
	return "co_user_vip"
}

// 企业经纬度储存
type CoUserLoc struct {
	ID   int64     `bson:"_id" json:"id"`      // Id
	Type int32     `bson:"type" json:"type"`   // 类型 1 企业位置 2 门店位置
	OfId int64     `bson:"of_id" json:"of_id"` // 相关 ID
	Loc  []float64 `bson:"loc" json:"loc"`     // 纬度 经度
}

func (m *CoUserLoc) TableName() string {
	return "co_user_loc"
}
