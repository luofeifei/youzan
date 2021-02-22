package middlewares

import (
	"base/client/pkg/comm"
	"base/client/pkg/sessions"
	"base/pkg/app"
	"base/pkg/config"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type configs struct {
	Domain string `yaml:"domain"`
}

var (
	conf configs
	once sync.Once
)

var key = "csrf_token"

func generateToken() string {
	length := 12
	rand.Seed(time.Now().UnixNano())
	var token = make([]byte, length, length)
	for i := 0; i < length; i++ {
		token[i] = byte(rand.Intn(127))
	}

	return base64.URLEncoding.EncodeToString(token)
}

// CsrfToken 对外中间件
func CsrfToken(ctx *gin.Context) {
	once.Do(func() {
		_ = config.Config().Bind(app.CfgName,"captcha", &conf, func() {
			_ = config.Config().Bind(app.CfgName, "sessions", &conf, nil)
		})
	})
	switch ctx.Request.Method {
	case http.MethodGet:
		sessions.Del(ctx, key)
		token := generateToken()
		sessions.Set(ctx, key, token)
		ctx.SetCookie(key, token, 3600, "/", conf.Domain, false, false)
	default:
		token, err := ctx.Cookie(key)
		if err != nil || token == "" {
			comm.NewResponse(app.Fail, nil, "CsrfTokenError").End(ctx)
			ctx.Abort()
			return
		}
		if token != sessions.Get(ctx, key) {
			comm.NewResponse(app.Fail, nil, "CsrfTokenError").End(ctx)
			ctx.Abort()
			return
		}
		sessions.Del(ctx, key)
		ctx.Next()
	}
}
