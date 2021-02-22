package client

import (
	"base/model/imp/serverCo"
	"base/tools"
)

type CoConfig struct {
	UpdatedAt   int64                  `json:"updated_at"`   // 更新时间
	ConfigValue map[string]interface{} `json:"config_value"` // 用户配置对应值
}

// 读取系统配置下的所有的配置
func GetCoConfig(coid int64, settingId ...int64) (res map[int64]CoConfig, err error) {
	// 做企业配置缓存处理 --- 预留功能
	result, errs := ServerCoUser().UserSettingValue(serverCo.ResCoUserSettingValue{Coid: coid, SettingId: settingId})
	if errs != nil {
		return res, errs
	}
	err = tools.JsonUnmarshal(result.Body, &res)
	return res, err
}

// 读取系统配置下的指定值
func GetCoConfigVal(coid int64, settingId int64, value string) (res interface{}, err error) {
	nodeVal, errs := GetCoConfig(coid, settingId)
	if errs != nil {
		return nil, errs
	}
	return nodeVal[settingId].ConfigValue[value], nil
}
