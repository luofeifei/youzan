package router
import (
	"client/controllers/shop"
	"github.com/gin-gonic/gin"
)

// 商品管理路由
func InitFrontShopRouter(Router *gin.RouterGroup) (R gin.IRoutes) {


	// 根据关键字搜索商品
	Router.POST("goods/pageFrontSearch", shop.GoodsListFrontSearchPage)
	// 根据商品Id查询商品详细信息
	Router.POST("goods/detailFront", shop.GoodsFrontDetail)
	// 刷新购物车里的商品
	Router.POST("goods/refresh", shop.QueryGoodsByIds)
    // 模拟购买
	Router.POST("goods/buy",shop.FakerBuy)

	return Router
}
