package middlewares

import (
	"base/client/pkg/comm"
	"base/pkg/app"
	"github.com/gin-gonic/gin"
)

// 并发锁
func VerifyLock(c *gin.Context) {
	if comm.GinRedLock(c) == false {
		comm.NewResponse(app.Tips, comm.ResponseMsg{Icon: app.MsgWarn}, app.FailedToAcquireLock).End(c)
		c.Abort()
	} else {
		c.Next()
	}
}