package common

// 插件列表
type PluginInfo struct {
	Alias   string `json:"alias"`    // 插件名称
	Ver     string `json:"ver"`      // 插件版本
	Type    int8   `json:"type"`     // 插件类型 1 空插件用于插件权限检测 2 so插件 3 gRpc中转
	Model   int8   `json:"model"`    // 插件模式 1 接口服务 2 工具过滤 3 队列处理 （一般页面为读取接口数据）
	UseType int8   `json:"use_type"` // 使用类型 -1 空 1 领卡时 2 购买前 3 购买后 （2才有改方法）
	Path    string `json:"path"`     // 插件下载路径
	Topic   string `json:"topic"`    // 订阅消费者话题 话题:channel组
}

// 存储插件信息
type PluginBaseInfo struct {
	Uid        int64 `json:"uid"`        // 调用插件的 用户ID
	Coid       int64 `json:"coid"`       // 调用插件的 企业ID
	FeatureId  int64 `json:"feature_id"` // 调用插件的 页面ID
	PluginInfo ApiPluginInfo
}

type ApiPluginInfo struct {
	Alias       string                 `json:"alias"`        // 组件序 ID 前端生成短ID
	PluginId    int64                  `json:"plugin_id"`    // 企业插件 ID
	ModuleAlias string                 `json:"module_alias"` // 组件模型标签
	Value       map[string]interface{} `json:"value"`        // 插件配置值
}
