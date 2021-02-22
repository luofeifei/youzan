package jwt

import (
	"base/pkg/app"
	"base/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	TokenExpired     = app.Err("Token 已过期")
	TokenNotValidYet = app.Err("Token 未激活")
	TokenMalformed   = app.Err("这不是 Token")
	TokenInvalid     = app.Err("无法解析的 Token")
	jwfConf          jwfConfig
)

type jwfConfig struct {
	Secret  string        `yaml:"secret"`
	TimeOut time.Duration `yaml:"timeout"`
}

// 载荷，可以加一些自己需要的信息
type customClaims struct {
	Uid      int64  `json:"uid"`      // 用户UID
	Type     int32  `json:"type"`     // 用户登录类型 1 用户 2总后台
	Key      string `json:"key"`      // 签名密钥
	Platform string `json:"platform"` // 登录平台
	jwt.StandardClaims
}

func Start() {
	_ = config.Config().Bind(app.CfgName, "jwt", &jwfConf, func() {
		_ = config.Config().Bind(app.CfgName, "jwt", &jwfConf, nil)
	})
}

// CreateToken 生成一个token
// uid 用户ID
// Type 用户类型 1用户 2企业
func CreateToken(uid int64, Type int32, Key string, platform string) (string, error) {
	claims := &customClaims{}
	claims.Uid = uid
	claims.Type = Type
	claims.Key = Key
	claims.ExpiresAt = time.Now().Add(jwfConf.TimeOut * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwfConf.Secret))
}

// 解析Tokne
func ParseToken(tokenString string) (*customClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwfConf.Secret), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwfConf.Secret), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*customClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return CreateToken(claims.Uid, claims.Type, claims.Key, claims.Platform)
	}
	return "", TokenInvalid
}
