package routers

import (
	"anruo-gin-demo/common"
	"anruo-gin-demo/controller"
	"anruo-gin-demo/middle"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func BaseRouters() *gin.Engine {
	r := gin.New()
	// 自定义日志格式
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("[%s | method: %s | path: %s | host: %s | proto: %s | code: %d | %s | %s ]\n",
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.ClientIP,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
		)
	}))
	// 健康检查
	InitHealthRouter(r.Group(""))
	// 不需要做鉴权的接口
	PublicGroup := r.Group("/api/v1")
	{
		PublicGroup.GET("/health", controller.Health.Health)
	}
	// 需要做鉴权的接口
	PrivateGroup := r.Group("/api/v1")
	// 鉴权
	PrivateGroup.Use(gin.Recovery()).Use(middle.Cors())
	{

	}
	r.NoRoute(func(ctx *gin.Context) {
		common.ReturnContext(ctx).Failed("fail", "该接口未开放")
		return
	})
	return r
}
