package client

import (
	"base/model/imp/serverAdmin"
	"base/tools"
	"time"
)

// sys_config 系统配置值
type sysConfig struct {
	UpdatedAt   int64                  `json:"updated_at"`   // 更新时间
	ConfigValue map[string]interface{} `json:"config_value"` // 用户配置对应值
}

var (
	conf map[string]sysConfig
)

func InitSysConfig(run func(), alias ...string) {
	tmp := make(map[string]int64, len(alias))
	if len(conf) == 0 {
		// 第一次进入
		conf = make(map[string]sysConfig)
		for _, val := range alias {
			tmp[val] = 0
			conf[val] = sysConfig{}
		}
		go GetSysConfigWatch(run)
	} else {
		// 重入将更新时间带上
		for key, val := range conf {
			tmp[key] = val.UpdatedAt
		}
	}
	change := false
	res, err := ServerAdminSys().SysConfigValue(serverAdmin.ResSysConfigValue{Alias: tmp})
	if err == nil {
		tmpConf := map[string]sysConfig{}
		err = tools.JsonUnmarshal(res.Body, &tmpConf)
		if err == nil {
			for key, val := range tmpConf {
				change = true
				conf[key] = val
			}
		}
	}
	if change && run != nil {
		run()
	}
	return
}

// 每隔10秒拉取一次配置
func GetSysConfigWatch(run func()) {
	go func() {
		timer2 := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-timer2.C:
				InitSysConfig(run)
			}
		}
	}()
}

// 读取系统配置下的所有的配置
func GetSysConfig(alias string) (map[string]interface{}, error) {
	nodeVal, ok := conf[alias]
	if ok {
		return nodeVal.ConfigValue, nil
	}
	return nil, nil
}

// 读取系统配置下的指定值
func GetSysConfigVal(alias string, value string) (interface{}, error) {
	nodeVal, ok := conf[alias]
	if ok {
		return nodeVal.ConfigValue[value], nil
	}
	return "", nil
}
