package client

import (
	"base/model/imp/serverBase"
	"base/model/imp/serverPlugin"
	"base/pkg/app"
	"base/tools"
)

// 请求发号器微服务获取唯一ID
func GetID(spaceName string) (id int64, err error) {
	res, errs := ServerBaseSys().GetID(serverBase.SpaceName{SpaceName: spaceName})
	if errs != nil {
		return 0, errs
	}
	return res.Id, err
}

// API插件请求结构体
type ReqApiPlugin struct {
	Uid        int64 `json:"uid"`        // 调用插件的 用户ID
	Coid       int64 `json:"coid"`       // 调用插件的 企业ID
	FeatureId  int64 `json:"feature_id"` // 调用插件的 页面ID
	PluginList []ReqApiPluginList
}

type ReqApiPluginList struct {
	Alias       string                 `json:"alias"`        // 组件序 ID 前端生成短ID
	PluginId    int64                  `json:"plugin_id"`    // 企业插件 ID
	ModuleAlias string                 `json:"module_alias"` // 组件模型标签
	Value       map[string]interface{} `json:"value"`        // 插件配置值
}

// 请求获取API插件服务数据
func GetApiPluginData(request ReqApiPlugin) (res interface{}, err error) {
	result, errs := ServerPlugin().GetApiPluginData(serverPlugin.ReqApiPlugin{
		Coid:       request.Coid,
		FeatureId:  request.FeatureId,
		PluginList: app.Struct2Json(request.PluginList),
	})
	if errs != nil || result.Code != 1 {
		return 0, app.Err("数据请求失败，请重试！")
	}
	var v interface{}
	err = tools.JsonUnmarshal(result.Body, &v)
	return v, err
}
