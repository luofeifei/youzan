package request

type Login struct {
	User   string `binding:"required,max=12" form:"user"`  // 用户账户
	Pass   string `binding:"required,max=128" form:"pass"` // 账户密码或短信验证码
	CodeId string `form:"code_id"`                         // 验证码id
	Code   string `form:"code"`                            // 验证码
	Type   int32  `binding:"required,max=1" form:"type"`   // 登录类型  0系统ID 1用户名 2手机号 3微信 4支付宝 5字节跳动 只支持 1与2登录
}
