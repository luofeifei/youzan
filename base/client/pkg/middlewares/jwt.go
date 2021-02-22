package middlewares

import (
	"base/client/pkg/comm"
	"base/client/pkg/jwt"
	"base/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	// JwtHeaderKey jwt token 在HTTP请求中的header名
	JwtHeaderKey = "Authorization"
)

// VerifyAuth 验证用户有效性的中间件
func VerifyAuth(c *gin.Context) {
	token := c.GetHeader(JwtHeaderKey)
	if token != "" {
		claims, err := jwt.ParseToken(token)
		if err == nil {
			c.Set("uid", claims.Uid)
			c.Set("uType", claims.Type)
			c.Header(JwtHeaderKey, token)
			c.Next()
			return
		}
	}
	comm.NewResponse(app.AuthFail, nil, app.AuthFailMessage).End(c, http.StatusUnauthorized)
	c.Abort()
}

// 获取当前JWT 用户ID
func GetUid(c *gin.Context) (UID int64, err error) {
	jwfUid, _ := c.Get("uid")
	return jwfUid.(int64), nil
}

// 获取当前JWT 用户类型
func GetType(c *gin.Context) (Type int32, err error) {
	AdminUser, exists := c.Get("uType")
	if !exists {
		return 0, app.Err("未获取到用户类型")
	}
	return AdminUser.(int32), nil
}

// 获取当前企业ID及用户ID
func GetCoIdUid(c *gin.Context) (Coid, Uid int64, err error) {
	jwfCoId, errs := strconv.ParseInt(c.GetHeader("coid"), 10, 64)
	if errs != nil {
		return 0, 0, app.Err("企业ID读取错误")
	}
	jwfUid, ts1 := c.Get("uid")
	if !ts1 {
		return 0, 0, app.Err("读取用户ID错误")
	}
	return jwfCoId, jwfUid.(int64), nil
}

func GetFastCoIdUid(c *gin.Context) (Coid, Uid int64) {
	jwfCoId, _ := c.Get("coid")
	jwfUid, _ := c.Get("uid")
	return jwfCoId.(int64), jwfUid.(int64)
}
