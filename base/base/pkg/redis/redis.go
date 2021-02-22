package redis

import (
	"base/pkg/app"
	"base/pkg/config"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type configs struct {
	Addr         string `yaml:"addr"`
	Password     string `yaml:"password"`
	Db           int    `yaml:"db"`
	PoolSize     int    `yaml:"pool_size"`
	MinIdleConns int    `yaml:"min_idle_conns"`
}

var (
	// Client redis连接资源
	Client *redis.Client
	conf   configs

)

// Start 启动redis
func Start() {
	err := config.Config().Bind(app.CfgName, "redis", &conf, func() {
		if config.Config().Bind(app.CfgName, "redis", &conf, nil) == nil {
			initStart()
		}
	})
	if  err != nil {
		app.Logger().Debug(fmt.Sprintf("redis 配置错误 err: %s", err.Error()))
		return
	}
	initStart()
}

func initStart() {
	Client = redis.NewClient(&redis.Options{
		Addr:         conf.Addr,
		Password:     conf.Password,
		DB:           conf.Db,
		PoolSize:     conf.PoolSize,
		MinIdleConns: conf.MinIdleConns,
	})

}

// CacheGet 获取指定key的值,如果值不存在,就执行f方法将返回值存入redis
func CacheGet(key string, expiration time.Duration, f func() string) string {
	var val string
	result, _ := Client.Get(key).Result()
	if len(result) == 0 {

		app.Println("CacheGet----")
		app.Println(key)
		Client.Set(key, f(), expiration)
		return val
	}
	return result
}

func CacheDel(key string) error {
	_, err := Client.Del(key).Result()
	return err
}
// 2020-12-16 roger add-----start-----------
// 普通队列推入
func Push(key string, value string) (error) {
	err := Client.LPush(key, value).Err()
	if err != nil {
		return err
	}
	return nil
}
// 普通队列弹出
func Pop(key string, value string) (string, error) {
	val, err := Client.RPop(key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
// 有序队列增加
func Zadd(key string, value interface{}, score int64) error {
	res, err := Client.ZAdd(key, redis.Z{float64(score), value}).Result()
	if res == 0 {
		return app.Err("加入有序队列失败")
	}
	if err != nil {
		return err
	}
	return nil
}
// 从有序队列中删除
func Zrem(key string, member interface{}) error {
	err := Client.ZRem(key, member).Err()

	if err != nil {
		return err
	}

	return nil
}
// 从有序队列取出排名最靠前的
func GetTop(key string) (string, int64, error) {
	val, err := Client.ZRangeWithScores(key, 0, 0).Result();
	if err != nil {
		return "", 0, err
	}
	if len(val) == 0 {
		return "", 0, nil
	}
	score := int64(val[0].Score)
	value := val[0].Member.(string)
	return value, score, nil
}
// 向key的hash中添加元素field的值
func HashSet(key, field string, data interface{})(error) {
	err := Client.HSet(key, field, data)
	if err != nil {
		return app.Err("Redis HSet Error")
	}
	return  nil
}

// 批量向key的hash添加对应元素field的值
func BatchHashSet(key string, fields map[string]interface{}) string {
	val, err := Client.HMSet(key, fields).Result()
	if err != nil {
		return err.Error()
	}
	return val
}

// 通过key获取hash的元素值
func HashGet(key, field string) string {
	result := ""
	val, err := Client.HGet(key, field).Result()
	if err == redis.Nil {
		return result
	}else if err != nil {
		return result
	}
	return val
}


func BatchHashGetAll(key string)  (res map[string]string,err error) {

	res, err = Client.HGetAll(key).Result()
	return
}
// 批量获取key的hash中对应多元素值
//func BatchHashGet(key string, fields ...string) map[string]string {
func BatchHashGet(key string, fields ...string) string {
	//resMap := make(map[string]string)
	//var resSli []interface{}

	//for _, field := range fields {
	//	var result string
	//	val, err := Client.HGet(key, fmt.Sprintf("%s", field)).Result()
	//	if err == redis.Nil {
	//		resMap[field] = result
	//	}else if err != nil {
	//		resMap[field] = result
	//	}
	//	if val != "" {
	//		resMap[field] = val
	//	}else {
	//		resMap[field] = result
	//	}
	//}

	//resMap, _ = Client.HGetAll(key).Result()

    resString,_:=Client.Get(key).Result()
	return resString
	//return resMap
}

// 获取自增唯一ID
func Incr(key string) int {
	val, err := Client.Incr(key).Result()
	if err != nil {
		return 0
	}
	return int(val)
}
// 添加集合数据
func SetAdd(key, val string){
	Client.SAdd(key, val)
}
// 从集合中获取数据
func SetGet(key string)[]string{
	val, err := Client.SMembers(key).Result()
	if err != nil{
		return []string{err.Error()}
	}
	return val
}
// 2020-12-16 roger add-----end-----------