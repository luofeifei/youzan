package mongoSql

// 支付流水号 永久记录 （可保留一段时间）
type PayCashFlow struct {
	ID        int64 `json:"id"`         // 支付流水订单
	InType    int8  `json:"in_type"`    // 来源类型 1用户订单 2企业订单
	OrderNo   int64 `json:"order_no"`   // 来源订单
	PayAmount int64 `json:"pay_amount"` // 支付金额 单位分
	PayAt     int64 `json:"pay_at"`     // 支付完成时间
	CreatedAt int64 `json:"created_at"` // 添加时间
	UpdatedAt int64 `json:"updated_at"` // 更新时间
}

// 企业对账单 （每日0点后更新）
type CoFundBill struct {
	Coid          int64 `json:"coid"`           // 企业ID
	BillDate      int64 `json:"bill_date"`      // 账单日期
	AccountType   int8  `json:"account_type"`   // 账户类型 1店铺余额 2储存资金 3保证金 4广告投放金 5标记资金
	IncomeAmount  int64 `json:"income_amount"`  // 收入金额
	IncomeCount   int64 `json:"income_count"`   // 收入笔数
	OutcomeAmount int64 `json:"outcome_amount"` // 支出金额
	OutcomeCount  int64 `json:"outcome_count"`  // 支出笔数
}

// 企业对账单明细
type CoFundRecord struct {
	Coid        int64 `json:"coid"`         // 企业ID
	OrderNo     int64 `json:"order_no"`     // 业务单号
	PayFlowNo   int64 `json:"pay_flow_no"`  // 支付流水单号
	AccountType int64 `json:"account_type"` // 账户类型 1店铺余额 2储存资金 3保证金 4广告投放金 5标记资金
	PayType     int8  `json:"pay_type"`     // 支付方式 1微信支付 2支付宝 3花呗支付 4银行卡 5QQ支付 6余额支付 7线下标记 8微信支付-自有 9支付宝支付-自有
	BizType     int8  `json:"trade_type"`   // 账单类型 入账 - 1充值 2订单入账 3提现退款 4转账退款 5交易手续费退回 6退款 7分销入账 8广告分润 9平台奖励 / 出账 - 20提现 21转账 22余额支付 23交易手续费 24云服务费 25平台服务费 26佣金支出 27仲裁 28赔付追偿 29平台扣款
	TradeType   int64 `json:"trade_type"`   // 商品类型 1商城 2企业VIP 3模板市场
	TradeId     int64 `json:"trade_id"`     // 商品关联 ID
	TradeDesc   int64 `json:"trade_desc"`   // 商品信息
	Balance     int64 `json:"balance"`      // 当前企业余额
	ChangeMoney int   `json:"change_money"` // 改变金额
	State       int8  `json:"state"`        // 状态 1进行中 2失败 3成功 （成功订单才会入账）
	CreatedAt   int   `json:"created_at"`   // 添加时间
	UpdatedAt   int   `json:"updated_at"`   // 更新时间
}

// 企业提现记录


// 用户对账单明细
type UserFundRecord struct {
	Uid         int64 `json:"uid"`          // 用户ID
	Coid        int64 `json:"coid"`         // 目标企业ID
	OrderNo     int64 `json:"order_no"`     // 业务单号
	PayFlowNo   int64 `json:"pay_flow_no"`  // 支付流水单号
	PayType     int64 `json:"trade_type"`   // 账单类型 入账 - 1充值（充入平台） 2订单退款 3提现退款 4分销入账 / 出账 - 20提现 21余额支付 22购买
	TradeType   int64 `json:"trade_type"`   // 商品类型 1商城
	TradeId     int64 `json:"trade_id"`     // 商品关联 ID
	TradeDesc   int64 `json:"trade_desc"`   // 商品信息
	ChangeMoney int   `json:"change_money"` // 改变金额
	State       int8  `json:"state"`        // 状态 1进行中 2失败 3成功
	CreatedAt   int   `json:"created_at"`   // 添加时间
	UpdatedAt   int   `json:"updated_at"`   // 更新时间
}
