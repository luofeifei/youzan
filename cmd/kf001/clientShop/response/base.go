package response

import "base/model/imp/serverShop"

type PageResult struct {
	List    [] *serverShop.ShopGroup  `json:"groupList"`
	Total int32        `json:"total"`
	Page     int32       `json:"page"`
	PageSize int32    `json:"pageSize"`
}

type GoodsPageResult struct {
	List    [] *serverShop.ShopGoods
	Total int32        `json:"total"`
	Page     int32       `json:"page"`
	PageSize int32    `json:"pageSize"`
}