package response

import (
	"base/model/modelSql/mall_co"
)

type CoUserBenefit struct {
	mall_co.CoUserBenefit
	MeetBag []struct {
		MeetType  int8  `json:"meet_type,omitempty"`  // 满足领取条件类型 1累计支付成功 2累计消费金额 3累计总积分为
		MeetValue int32 `json:"meet_value,omitempty"` // 满足条件值
	} `json:"meet_bag,omitempty"`                                           // 满足条件值
	BenefitBag []CoUserBenefitBenefitBag `json:"use_benefit_bag,omitempty"` // 已使用权益
}

type CoUserBenefitBenefitBag struct {
	ID          int64                              `json:"id"`
	BenefitID   int64                              `json:"benefit_id"`             // 权益模板ID
	Mode        int8                               `json:"mode"`                   // 服务模式 1系统核单 2商户线下核单
	Name        string                             `json:"name"`                   // 展示名称
	Icon        string                             `json:"icon"`                   // 权益图标
	Description string                             `json:"description"`            // 权益简介
	UseState    int8                               `json:"use_state"`              // 使用状态 1未使用 2使用中
	BenefitInfo CoUserBenefitBenefitBagBenefitInfo `json:"benefit_info,omitempty"` // 权益列表及设置
}

type CoUserBenefitBenefitBagBenefitInfo struct {
	ClassID   int64 `json:"class_id,omitempty"`  // 所属权益分类
	PluginID  int64 `json:"plugin_id,omitempty"` // 使用的插件ID 根据插件权限决定是否可用
	Mode      int8  `json:"mode,omitempty"`      // 服务模式 1系统核单 2商户线下核单
	Parameter map[string]struct {
		Name  string      `json:"name"`
		Value int64       `json:"value"`
		Rule  interface{} `json:"rule"`
	} `json:"parameter,omitempty"` // 参数定义值 JSON数组
}
