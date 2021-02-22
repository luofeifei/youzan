package database

import (
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

type (
	RedisClient struct {
		Client *redis.Client
	}
)

func RedisConnect(dbName, addr, password, dbId, poolSize, minIdleConn string) (db *RedisClient, err error) {
	dbTmp, _ := strconv.Atoi(dbId)
	poolSizeTmp, _ := strconv.Atoi(poolSize)
	minIdleConnTmp, _ := strconv.Atoi(minIdleConn)
	dbs := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           dbTmp,
		PoolSize:     poolSizeTmp,
		MinIdleConns: minIdleConnTmp,
	})
	return &RedisClient{
		Client: dbs,
	}, err
}

// 获取 REDIS 客户端
func (collection *RedisClient) GetClient() *redis.Client {
	return collection.Client
}

// CacheGet 获取指定key的值,如果值不存在,就执行f方法将返回值存入redis
func (collection *RedisClient) CacheGet(key string, expiration time.Duration, f func() string) string {
	var val string
	result, _ := collection.Client.Get(key).Result()
	if len(result) == 0 {
		collection.Client.Set(key, f(), expiration)
		return val
	}
	return result
}

func (collection *RedisClient) CacheDel(key string) error {
	_, err := collection.Client.Del(key).Result()
	return err
}
