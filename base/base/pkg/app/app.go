package app

import (
	"errors"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/TarsCloud/TarsGo/tars/util/rogger"
	"strconv"
	"sync"
	"time"
)

var (
	start    int64
	Cfg      = tars.GetServerConfig()
	CfgName  = "settings.yml"
	services = &Services{services: make(map[string]interface{})}
)

// Services 服务汇总
type Services struct {
	lock     sync.Mutex
	services map[string]interface{}
}

// Logger 获取日志对象
func Logger() *rogger.Logger {
	return Get("logger").(*rogger.Logger)
}

func (service *Services) register(name string, se interface{}) {
	service.lock.Lock()
	defer service.lock.Unlock()
	service.services[name] = se
}

func (service *Services) get(name string) interface{} {
	if val, ok := service.services[name]; ok {
		return val
	}
	return nil
}

// Register 注册其他包的服务
func Register(name string, service interface{}) interface{} {
	services.register(name, service)
	return service
}

// Get 获取其他包的服务
func Get(name string) interface{} {
	return services.get(name)
}

func Err(message string) (err error) {
	return errors.New(message)
}

func IdInt(Id string, empty bool) (ID uint64, err error) {
	if Id != "" {
		ID, err := strconv.ParseUint(Id, 10, 64)
		if err == nil {
			return ID, err
		} else {
			return 0, errors.New("IdIsInt")
		}
	} else if empty == true {
		return 0, errors.New("IdIsInt")
	}
	return 0, nil
}

// start 调试一段程序消耗的 时间
func Println(req interface{}) {
	if start > 0 {
		tmp := time.Now().UnixNano() / 1e6
		fmt.Println(fmt.Sprintf("结束：%d 耗时：%d 毫秒", tmp, tmp-start))
		start = 0
	}
	if value, ok := req.(string); ok {
		if value == "start" {
			start = time.Now().UnixNano() / 1e6
			fmt.Println(fmt.Sprintf("开始：%d", start))
		} else {
			fmt.Println(value)
		}
	} else {
		fmt.Println(string(Struct2Json(req)))
	}
	return
}
