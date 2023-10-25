package controller

import (
	"anruo-gin-demo/common"
	"github.com/gin-gonic/gin"
)

// 健康检查

type health struct{}

var Health health

func (*health) Health(ctx *gin.Context) {
	common.ReturnContext(ctx).Successful("success", "success")
}
