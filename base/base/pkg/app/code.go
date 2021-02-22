package app

// 返回自定义状态码
const (
	Success  = 0   // 响应正确
	Tips     = 1   // 弹窗提示
	Fail     = 2   // 错误
	NotFound = 404 // 404页面不存在
	AuthFail = 401 // 登录验证失败
)

// 自定义的一些错误消息的返回
const (
	AuthFailMessage         = "认证失败"
	PermissionDeniedMessage = "无操作权限"
	FailedToAcquireLock     = "请求限制尝试过快"
	SUCCESS                 = "操作成功"
	MsgInfo                 = "info"
	MsgSuccess              = "success"
	MsgWarn                 = "warning"
	MsgError                = "error"
)
