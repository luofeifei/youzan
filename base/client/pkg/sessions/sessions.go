package sessions

import (
	"base/pkg/app"
	"base/pkg/config"
	"github.com/gin-contrib/sessions"
	redisSession "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"strconv"
)

type configs struct {
	Key          string `yaml:"key"`
	Name         string `yaml:"name"`
	Domain       string `yaml:"domain"`
	Addr         string `yaml:"addr"`
	Password     string `yaml:"password"`
	Db           int    `yaml:"db"`
	PoolSize     int    `yaml:"pool_size"`
	MinIdleConns int    `yaml:"min_idle_conns"`
}

var conf configs

// Inject 启动session服务, 在自定义的路由代码中调用, 传入 *gin.Engine 对象
func Inject(engine *gin.Engine) {
	err := config.Config().Bind(app.CfgName, "sessions", &conf, func() {
		if config.Config().Bind(app.CfgName, "sessions", &conf, nil) == nil {
			initStart(engine)
		}
	})
	if err == nil {
		initStart(engine)
	}
	return
}

func initStart(engine *gin.Engine) gin.IRoutes {
	store, err := redisSession.NewStoreWithDB(conf.PoolSize, "tcp", conf.Addr, conf.Password, strconv.Itoa(conf.Db), []byte(conf.Key))
	if err != nil {
		app.Logger().Error(err)
		return engine
	}
	store.Options(sessions.Options{MaxAge: 3600, Path: "/", Domain: conf.Domain, HttpOnly: true})
	return engine.Use(sessions.Sessions(conf.Name, store))
}

// Get 获取指定session
func Get(c *gin.Context, key string) string {
	sess := sessions.Default(c)
	val := sess.Get(key)
	if val != nil {
		return val.(string)
	}
	return ""
}

// Set 设置session
func Set(c *gin.Context, key, val string) {
	sess := sessions.Default(c)
	sess.Set(key, val)
	_ = sess.Save()
}

// Del 删除指定session
func Del(c *gin.Context, key string) {
	sess := sessions.Default(c)
	sess.Delete(key)
	_ = sess.Save()
}
