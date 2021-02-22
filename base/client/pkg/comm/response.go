package comm

import (
	"base/client/pkg/validator"
	"base/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response HTTP返回数据结构体, 可使用这个, 也可以自定义
type Response struct {
	Code    int         `json:"code"`    // 状态码,这个状态码是与前端和APP约定的状态码,非HTTP状态码
	Data    interface{} `json:"data"`    // 返回数据
	Message string      `json:"message"` // 自定义返回的消息内容
}

// Response 提示输出信息结构体
type ResponseMsg struct {
	Type int8        `json:"type"` // 提示类型 1 消息提示 2 通知 3 弹框
	Name string      `json:"name"` // 标题
	Icon string      `json:"icon"` // 主体、图标 success warning info error
	Time int8        `json:"time"` // 自动关闭时间
	Data interface{} `json:"data"` // 附带数据
}

// Response 模板输出内容
type ResponseTmlMsg struct {
	Name string `json:"name"` // 标题
	Icon string `json:"icon"` // 主体、图标 success warning info error
	Time int8   `json:"time"` // 自动关闭时间
	Jump string `json:"jump"` // 跳转地址
	Err  string `json:"err"`  // 错误提示
}

// End 在调用了这个方法之后,还是需要 return 的
func (rsp *Response) End(c *gin.Context, httpStatus ...int) {
	status := http.StatusOK
	if len(httpStatus) > 0 {
		status = httpStatus[0]
	}
	// 如果存在并发锁 输出时解锁
	GinUnLock(c)
	c.JSON(status, rsp)
}

// NewResponse 接口返回统一使用这个
//  code 服务端与客户端和web端约定的自定义状态码
//  data 具体的返回数据
//  message 可不传,自定义消息
func NewResponse(code int, data interface{}, message ...string) *Response {
	msg := ""
	if len(message) > 0 {
		msg = message[0]
	}
	return &Response{Code: code, Data: data, Message: msg}
}

// 输出表单验证失败提示
func ValidatorResponse(ctx *gin.Context, err *validator.ValidErrors) {
	NewResponse(app.Tips, ResponseMsg{Type: 1, Icon: app.MsgWarn, Data: err.ErrorsInfo}, "Validator").End(ctx)
}

// 信息输出
// tips 提示类型 0 仅输出 1 消息提示 2 通知 3 弹框
// errTips 如果无错误是否忽略弹出提示
func ApiResponse(ctx *gin.Context, data interface{}, err error, tips int8, ignore ...bool) {
	if tips == 0 {
		if err != nil {
			NewResponse(app.Fail, data, err.Error()).End(ctx)
		} else {
			NewResponse(app.Success, data, app.SUCCESS).End(ctx)
		}
	} else {
		if err != nil {
			NewResponse(app.Tips, ResponseMsg{Type: tips, Icon: app.MsgWarn, Data: data}, err.Error()).End(ctx)
		} else {
			if len(ignore) > 0 {
				NewResponse(app.Success, data, app.SUCCESS).End(ctx)
			} else {
				NewResponse(app.Tips, ResponseMsg{Type: tips, Icon: app.MsgSuccess, Data: data}, app.SUCCESS).End(ctx)
			}
		}
	}
}

func TplResponse(ctx *gin.Context, name string, err error) {
	ctx.HTML(http.StatusOK, name, ResponseTmlMsg{
		Err: err.Error(),
	})
}
