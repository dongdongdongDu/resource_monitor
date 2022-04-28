package routers

import (
	"github.com/gin-gonic/gin"
	"resource_monitor/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/monitor", controller.UpdateInfoHandler)
	return r
}
