package mall_system

import "base/model/modelSql"

// SystemAPI [...]
type SystemAPI struct {
	modelSql.Database
	APIGroup    string `gorm:"index:api_group;column:api_group;type:varchar(50);not null" json:"api_group"` // API组名
	Method      string `gorm:"index:path;column:method;type:char(6);not null" json:"method"`                // 访问方式
	Path        string `gorm:"index:path;column:path;type:varchar(100);not null" json:"path"`               // 访问路径
	Description string `gorm:"column:description;type:varchar(50);not null" json:"description"`             // 接口简介
}

func (m *SystemAPI) TableName() string {
	return "system_api"
}

// SystemGroup [...]
type SystemGroup struct {
	modelSql.Database
	Pid   int64  `gorm:"index:pid;column:pid;type:bigint(17);not null" json:"pid"`       // 父ID
	Name  string `gorm:"column:name;type:varchar(50);not null" json:"name"`              // 管理组名称
	API   string `gorm:"column:api;type:text" json:"api"`                                // API权限
	Menus string `gorm:"column:menus;type:text" json:"menus"`                            // 表单权限
	State int8   `gorm:"index:state;column:state;type:tinyint(2);not null" json:"state"` // 状态
}

func (m *SystemGroup) TableName() string {
	return "system_group"
}

// SystemGroupsLink [...]
type SystemGroupsLink struct {
	ID      int64 `gorm:"index:uid;column:id;type:bigint(17);not null" json:"id"`         // 会员ID
	GroupID int64 `gorm:"index:group_id;column:group_id;type:bigint(17)" json:"group_id"` // 级别ID
}

func (m *SystemGroupsLink) TableName() string {
	return "system_groups_link"
}

// 管理员日志
type SystemLog struct {
	modelSql.Database
	Type int8   `gorm:"index:type;column:type;type:tinyint(4);not null" json:"type"` // 类型
	Msg  string `gorm:"column:msg;type:varchar(255);not null" json:"msg"`            // 消息
	Lid  int64  `gorm:"column:lid;type:bigint(17);not null" json:"lid"`              // 关系ID
	Data string `gorm:"column:data;type:text" json:"data"`                           // 改变数据
	IP   string `gorm:"column:ip;type:char(12);not null" json:"ip"`                  // 客户端IP
}

func (m *SystemLog) TableName() string {
	return "system_log"
}

// 管理员总后台菜单
type SystemMenu struct {
	modelSql.Database
	Pid       int64  `gorm:"column:pid;type:bigint(17);not null" json:"pid"`                // 父级
	Type      int8   `gorm:"index:type;column:type;type:tinyint(2)" json:"type"`            // 类型 0 菜单 1路由
	Title     string `gorm:"column:title;type:varchar(20)" json:"title"`                    // 规则名称
	Path      string `gorm:"index:path;column:path;type:varchar(255);not null" json:"path"` // 访问地址
	Component string `gorm:"column:component;type:varchar(255);not null" json:"component"`  // 模板地址
	Keepalive int8   `gorm:"column:keepalive;type:tinyint(2)" json:"keepalive"`
	Icon      string `gorm:"column:icon;type:varchar(50)" json:"icon"`                       // 图标
	Sort      int    `gorm:"index:weigh;column:sort;type:int(5)" json:"sort"`                // 权重
	State     int8   `gorm:"index:state;column:state;type:tinyint(2);not null" json:"state"` // 状态
}

func (m *SystemMenu) TableName() string {
	return "system_menu"
}

// SystemUpCommon [...]
type SystemUpCommon struct {
	modelSql.Database
	Type     int8  `gorm:"column:type;type:tinyint(2);not null" json:"type"`           // 公用类型
	FileType int8  `gorm:"column:file_type;type:tinyint(2);not null" json:"file_type"` // 附件类型
	FileID   int64 `gorm:"column:file_id;type:bigint(17);not null" json:"file_id"`     // 附件ID
}

func (m *SystemUpCommon) TableName() string {
	return "system_up_common"
}

// 图片上传附件
type SystemUpPic struct {
	modelSql.Database
	Type   int8  `gorm:"unique_index:type;column:type;type:tinyint(2);not null" json:"type"`   // 用户类型 1企业 2用户 3管理员
	Aid    int64 `gorm:"unique_index:type;column:aid;type:bigint(17);not null" json:"aid"`     // 所属用户ID
	FileID int64 `gorm:"index:file_id;column:file_id;type:bigint(17);not null" json:"file_id"` // 关联图片附件ID
}

func (m *SystemUpPic) TableName() string {
	return "system_up_pic"
}

// SystemUpPicFile [...]
type SystemUpPicFile struct {
	modelSql.Database
	Hash  string `gorm:"unique;column:hash;type:char(40);not null" json:"hash"`          // 文件 hash 值
	Path  string `gorm:"column:path;type:varchar(50);not null" json:"path"`              // 目录
	Name  string `gorm:"column:name;type:varchar(100);not null" json:"name"`             // 文件名
	Tag   string `gorm:"column:tag;type:varchar(10);not null" json:"tag"`                // 后缀
	Size  int    `gorm:"column:size;type:int(11);not null" json:"size"`                  // 文件大小
	State int8   `gorm:"index:state;column:state;type:tinyint(2);not null" json:"state"` // 状态
}

func (m *SystemUpPicFile) TableName() string {
	return "system_up_pic_file"
}

// 文件上传附件
type SystemUpFile struct {
	modelSql.Database
	Type   int8  `gorm:"unique_index:type;column:type;type:tinyint(2);not null" json:"type"`   // 用户类型 1企业 2用户 3管理员
	Aid    int64 `gorm:"unique_index:type;column:aid;type:bigint(17);not null" json:"aid"`     // 所属用户ID
	FileID int64 `gorm:"index:file_id;column:file_id;type:bigint(17);not null" json:"file_id"` // 关联图片附件ID
}

func (m *SystemUpFile) TableName() string {
	return "system_up_file"
}

// SystemUpPicFile [...]
type SystemUpFileFile struct {
	modelSql.Database
	Hash  string `gorm:"unique;column:hash;type:char(40);not null" json:"hash"`          // 文件 hash 值
	Path  string `gorm:"column:path;type:varchar(50);not null" json:"path"`              // 目录
	Name  string `gorm:"column:name;type:varchar(100);not null" json:"name"`             // 文件名
	Tag   string `gorm:"column:tag;type:varchar(10);not null" json:"tag"`                // 后缀
	Size  int    `gorm:"column:size;type:int(11);not null" json:"size"`                  // 文件大小
	State int8   `gorm:"index:state;column:state;type:tinyint(2);not null" json:"state"` // 状态
}

func (m *SystemUpFileFile) TableName() string {
	return "system_up_file_file"
}

// 视频上传附件
type SystemUpVideo struct {
	modelSql.Database
	ID     int64 `gorm:"primary_key;column:id;type:bigint(17);not null" json:"id"`             // 视频上传附件
	Type   int8  `gorm:"index:type;column:type;type:tinyint(2);not null" json:"type"`          // 用户类型 1企业 2用户 3管理员
	Aid    int64 `gorm:"index:type;column:aid;type:bigint(17);not null" json:"aid"`            // 所属用户ID
	FileID int64 `gorm:"index:file_id;column:file_id;type:bigint(17);not null" json:"file_id"` // 关联图片附件ID
}

func (m *SystemUpVideo) TableName() string {
	return "system_up_video"
}

// SystemUpVideoFile [...]
type SystemUpVideoFile struct {
	modelSql.Database
	Hash     string `gorm:"unique;column:hash;type:varchar(40);not null" json:"hash"`       // 文件 hash 值
	Path     string `gorm:"column:path;type:varchar(50);not null" json:"path"`              // 目录
	Name     string `gorm:"column:name;type:varchar(100);not null" json:"name"`             // 文件名
	Tag      string `gorm:"column:tag;type:varchar(10);not null" json:"tag"`                // 后缀
	Size     int    `gorm:"column:size;type:int(11);not null" json:"size"`                  // 文件大小
	Duration int    `gorm:"column:duration;type:int(11);not null" json:"duration"`          // 影片时长
	State    int8   `gorm:"index:state;column:state;type:tinyint(2);not null" json:"state"` // 状态
}

func (m *SystemUpVideoFile) TableName() string {
	return "system_up_video_file"
}

// SystemUser [...]
type SystemUser struct {
	modelSql.Database
	Username      string `gorm:"index:user;column:username;type:varchar(20);not null" json:"username"` // 帐号
	Password      string `gorm:"column:password;type:char(128);not null" json:"password"`              // 密码
	Encrypt       string `gorm:"column:encrypt;type:char(6);not null" json:"encrypt"`                  // 加密密匙
	NickName      string `gorm:"column:nick_name;type:varchar(64);not null" json:"nick_name"`          // 昵称姓名
	WeiXin        string `gorm:"column:wei_xin;type:varchar(20);not null" json:"wei_xin"`              // 绑定微信号
	Mobile        string `gorm:"column:mobile;type:varchar(12);not null" json:"mobile"`                // 联系电话
	Avatar        string `gorm:"column:avatar;type:varchar(100);not null" json:"avatar"`               // 用户图像
	Online        int8   `gorm:"column:online;type:tinyint(2);not null" json:"online"`                 // 在线状态
	LoginFailure  int8   `gorm:"column:login_failure;type:tinyint(2);not null" json:"login_failure"`   // 密码错误次数
	LastLoginIP   string `gorm:"column:last_login_ip;type:char(12);not null" json:"last_login_ip"`     // 最后登录IP
	LastLoginTime int    `gorm:"column:last_login_time;type:int(10);not null" json:"last_login_time"`  // 最后登录时间
	LoginNum      int16  `gorm:"column:login_num;type:smallint(6);not null" json:"login_num"`          // 登陆次数
	State         int8   `gorm:"index:user;column:state;type:tinyint(2);not null" json:"state"`        // 账号状态
}

func (m *SystemUser) TableName() string {
	return "system_user"
}
