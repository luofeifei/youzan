package config

import (
	"base/pkg/app"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// Configuration 应用配置
type Configuration struct {
	configs    map[string]*viper.Viper
	configsRun map[string][]configCb
}

type configCb func()

var (
	config = new(Configuration).init()
	uptime int64
)

// Config 得到config对象
func Config() *Configuration {
	return config
}

func (c *Configuration) init() *Configuration {
	c.configs = make(map[string]*viper.Viper)
	c.configsRun = make(map[string][]configCb)
	return c
}

// 获取 tars 远程配置
func (c *Configuration) GetRemoteConf(filename string) (conf *viper.Viper, err error) {
	nodeVal, ok := c.configs[filename]
	if !ok {
		// 读取本地配置
		bytes, err := ioutil.ReadFile("settings.yml")
		config := string(bytes)
		// 读取远程配置
		// remoteConf := tars.NewRConf(app.Cfg.App, app.Cfg.Server, app.Cfg.BasePath)
		// config, err := remoteConf.GetConfig(filename)
		if err != nil {
			return conf, err
		} else {
			if err := c.configSetup(filename, config); err != nil {
				return conf, nil
			}
			return c.configs[filename], nil
		}
	}
	return nodeVal, nil
}

func (c *Configuration) configSetup(filename string, config string) (err error) {
	c.configs[filename] = viper.New()
	c.configs[filename].SetConfigFile(app.Cfg.BasePath + "/" + filename)
	c.configs[filename].SetConfigType("yaml")
	if err = c.configs[filename].ReadConfig(strings.NewReader(os.ExpandEnv(config))); err != nil {
		return err
	}
	c.configs[filename].OnConfigChange(func(e fsnotify.Event) {
		if uptime+2 < time.Now().Unix() {
			uptime = time.Now().Unix()
			fmt.Println("监听到文件[" + e.Name + "]改变！")
			for _, val := range c.configsRun[filename] {
				val()
			}
		}
	})
	c.configs[filename].WatchConfig()
	return nil
}

func (c *Configuration) Bind(filename string, key string, obj interface{}, run configCb) error {
	objNode, err := c.GetRemoteConf(filename)
	if err == nil {
		objVal := objNode.Sub(key)
		if objVal != nil {
			if obj != nil {
				bs, err := yaml.Marshal(objVal.AllSettings())
				if err != nil {
					return err
				}
				return yaml.Unmarshal(bs, obj)
			}
			if run != nil {
				c.configsRun[filename] = append(c.configsRun[filename], run)
			}
			return nil
		}
	}
	errName := "读取配置[" + filename + " - " + key + "]失败"
	return app.Err(errName)
}
