package request

type Id struct {
	ID int64 `binding:"required" json:"id" form:"id"` // ID 本系统为17位数字
}

type PageInfo struct {
	Page     int32  `binding:"required" json:"page" form:"page"`         // 当前页
	PageSize int32  `binding:"required" json:"pageSize" form:"pageSize"` // 每页显示数
	OrderKey string `json:"orderKey" form:"orderKey"`                    // 默认排序字段 -filed1,+field2,field3 (-Desc 降序)
}
