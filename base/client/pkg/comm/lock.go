package comm

import (
	"base/pkg/redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GIN 并发处理检查
func GinRedLock(c *gin.Context) bool {
	if c.Request.Method == http.MethodPut || c.Request.Method == http.MethodDelete {
		Uid, exists := c.Get("uid")
		name := "redLock:"
		if !exists {
			name += c.ClientIP() + ":" + c.Request.URL.Path
		} else {
			name += strconv.Itoa(int(Uid.(int64)))
		}
		fmt.Println(name)
		lock := redis.NewRedisLock(name, false)
		if err := lock.Lock(redis.Client); err != nil {
			return false
		}
		c.Set("lock", lock)
	}
	return true
}

// GIN 解除并发锁
func GinUnLock(c *gin.Context) bool {
	if value, exists := c.Get("lock"); exists {
		lock, ok := value.(*redis.RedisLock)
		if ok {
			_ = lock.UnLock(redis.Client)
		}
	}
	return true
}
