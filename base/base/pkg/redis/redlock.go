package redis

import (
	"encoding/base64"
	"errors"
	"github.com/go-redis/redis"
	"math/rand"
	"sync"
	"time"
)

const (
	DefaultRetry        = 5                      //默认重试
	DefaultInterval     = 100 * time.Millisecond //默认间隔
	DefaultExpire       = 5 * time.Second        //默认过期
	DefaultExpireUnsame = 600 * time.Second      //默认过期时间
)

var ErrFailed = errors.New("failed to acquire lock")

type RedisLock struct {
	Expiry        time.Duration
	Retry         int
	RetryInterval time.Duration
	lockKey       string
	lockValue     string
	sameLock      bool
	mutex         sync.Mutex
}

func NewRedisLock(lk string, sameLock bool) *RedisLock {
	if sameLock {
		return &RedisLock{lockKey: lk, Expiry: DefaultExpire, Retry: DefaultRetry, RetryInterval: DefaultInterval}
	} else {
		return &RedisLock{lockKey: lk, Expiry: DefaultExpireUnsame, Retry: 1, RetryInterval: 0}
	}
}

func NewRedisLockWithParam(lk string, expire int, retry int, interval int) *RedisLock {
	return &RedisLock{lockKey: lk, Expiry: time.Duration(expire) * time.Second, Retry: retry, RetryInterval: time.Duration(interval) * time.Millisecond}
}

func (rl *RedisLock) Lock(rds *redis.Client) error {
	if rl.sameLock {
		rl.mutex.Lock()
		err := rl.lock(rds)
		if err != nil {
			rl.mutex.Unlock()
		}
		return err
	} else {
		err := rl.lock(rds)
		return err
	}
}

func (rl *RedisLock) lock(rds *redis.Client) error {
	if rl.sameLock {
		b := make([]byte, 16)
		_, err := rand.Read(b)
		if err != nil {
			return err
		}
		rl.lockValue = base64.StdEncoding.EncodeToString(b)
	} else {
		rl.lockValue = ""
	}
	for i := 0; i < rl.Retry; i++ {
		ok, err := rds.SetNX(rl.lockKey, rl.lockValue, rl.Expiry).Result()
		if err != nil {
			return err
		}
		if ok {
			return nil
		}
		if i < rl.Retry-1 {
			time.Sleep(rl.RetryInterval)
		}
	}
	return ErrFailed
}

const delScript = `
if redis.call("get", KEYS[1]) == ARGV[1] then
	return redis.call("del", KEYS[1])
else
	return 0
end`

func (rl *RedisLock) UnLock(rds *redis.Client) error {
	if rl.sameLock {
		_, err := rds.Eval(delScript, []string{rl.lockKey}, rl.lockValue).Result()
		rl.mutex.Unlock()
		return err
	} else {
		_, err := rds.Del(rl.lockKey).Result()
		return err
	}
}

// 简单锁 name锁名称 expire 有效期 retry 重试次数
func RedLock(name string, expire int, retry int) (lock *RedisLock, res bool) {
	lock = NewRedisLockWithParam("redLock:"+name, expire, retry, 100)
	if err := lock.Lock(Client); err != nil {
		return nil, false
	}
	return lock, true
}
