package routers

import (
	"anruo-gin-demo/controller"
	"github.com/gin-gonic/gin"
)

// 监控检测接口

func InitHealthRouter(r *gin.RouterGroup) {

	r.GET("/health", controller.Health.Health)
	r.POST("/health", controller.Health.Health)
}
