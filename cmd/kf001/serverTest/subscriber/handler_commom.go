package subscriber

import (
	"base/pkg/app"
)

// 检查手机号是否已在系统注册 如果存在发送邀请 如果不存在 发送注册邀请 短信
func UserStaffData(body interface{}) (err error) {
	var msg struct {
		Type int8  `json:"type"` // 1添加 2删除
		Id   int64 `json:"id"`   // 企业员工 ID
	}
	err = app.Unmarshal(body, &msg)
	if err == nil {
		// 检查企业用户是否已注册如果没有发送短信

	}
	app.Println(msg)

	return nil
}
