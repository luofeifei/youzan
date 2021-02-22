package router

import (
	"client/controllers/base"
	"github.com/gin-gonic/gin"
)

// 公共开放接口
func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	router := Router.Group("base")
	{
		// 获取国内地区数据列表
		router.GET("region", base.GetRegion)
		router.POST("regionLngLat", base.GetRegionLngLat)
	}
	return router
}
