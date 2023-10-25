package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 接口返回内容格式 msg为提示信息, data为数据

type BaseContext struct {
	ctx *gin.Context
}

// 返回格式

type ReturnMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 成功返回

func ReturnContext(ctx *gin.Context) *BaseContext {
	return &BaseContext{ctx: ctx}
}
func (BaseContext *BaseContext) Successful(msg string, data interface{}) {
	resp := &ReturnMsg{
		Code: 20000,
		Msg:  msg,
		Data: data,
	}
	BaseContext.ctx.JSON(http.StatusOK, resp)
}

// 失败返回

func (BaseContext *BaseContext) Failed(msg string, data interface{}) {
	resp := &ReturnMsg{
		Code: 50000,
		Msg:  msg,
		Data: data,
	}
	BaseContext.ctx.JSON(http.StatusOK, resp)
}
