package router


import (
	"client/controllers/shop"
	"github.com/gin-gonic/gin"
)

// 商品管理路由
func InitShopRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
		// 添加商品
		Router.PUT("goods/add", shop.AddGoods)
		// 修改商品
		Router.PUT("goods/edit", shop.EditGoods)
		// 筛选商品
		Router.POST("goods/pageSearch", shop.GoodsListSearchPage)
		// 根据商品id获取一件商品
		Router.POST("goods/detail", shop.GoodsDetail)
		// 批量操作商品
		Router.PUT("goods/operate", shop.GoodsBatchOperate)
		// 添加企业分组
		Router.PUT("group/add", shop.AddGroup)
		// 修改企业分组
		Router.PUT("group/edit", shop.EditGroup)
		// 商品分组列表
		Router.POST("group/page", shop.GroupListByCoidPage)
		// 商品删除
		Router.DELETE("group/delete", shop.GroupDelete)
		// 获取商品分组详情
		Router.POST("group/detail", shop.ShopGroupDetail)
		// 优惠券兑换卡新增
		Router.PUT("coupon/card/add", shop.CouponCardAdd)
		// 优惠券兑换卡修改
		Router.PUT("coupon/card/edit", shop.CouponCardEdit)
		// 优惠券兑换卡删除
		Router.DELETE("coupon/card/delete", shop.CouponCardDelete)
		// 查询优惠券兑换卡
		Router.POST("coupon/card/detail", shop.CouponCardDetail)

		// 优惠券添加测试
		Router.PUT("coupon/add", shop.CouponAdd)
		// 优惠券修改测试
		Router.PUT("coupon/edit", shop.CouponEdit)
		// 优惠券详情
		Router.PUT("coupon/detail", shop.CouponDetail)
		// 优惠券分页列表
		Router.POST("coupon/page", shop.CouponListPage)

		// 商品库存新增
		Router.PUT("stock/add", shop.StockAdd)
		// 商品库存修改 加减数量
		Router.PUT("stock/edit", shop.StockEdit)
		// 商品库存删除
		Router.DELETE("stock/delete", shop.StockDelete)
		// 获取库存卡详情
		Router.POST("stock/detail", shop.StockDetail)
		// 添加库存卡
		Router.PUT("stock/card/add", shop.StockCardAdd)
		// 修改库存卡
		Router.PUT("stock/card/edit", shop.StockCardEdit)
		// 删除库存卡
		Router.DELETE("stock/card/delete", shop.StockCardDelete)
		// 查询库存卡
		Router.POST("stock/card/detail", shop.StockCardDetail)

	return Router
}

