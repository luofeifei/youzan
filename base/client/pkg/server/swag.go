package server

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	swagHandler gin.HandlerFunc
)


func init() {
	swagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler)
}
