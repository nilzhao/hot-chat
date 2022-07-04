package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResCode int

const (
	CODE_SUCCESS      ResCode = 0
	CODE_BAD_REQUEST  ResCode = 400
	CODE_UNAUTHORIZED ResCode = 401
	CODE_FORBIDDEN    ResCode = 403
	CODE_NOT_FOUND    ResCode = 404
)

type Context struct {
	*gin.Context
}

type ResponseBody struct {
	Code ResCode     `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type HandlerFunc func(c *Context)

func (ctx *Context) NewResponse(code ResCode, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, ResponseBody{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func (ctx *Context) ResOk(data any) {
	ctx.NewResponse(CODE_SUCCESS, data, "ok")
}

func (ctx *Context) ResFailed(err error, codes ...ResCode) {
	code := CODE_BAD_REQUEST
	if len(codes) != 0 {
		code = codes[0]
	}
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	ctx.NewResponse(code, nil, msg)
}

func (ctx *Context) ResNotFound(msgs ...string) {
	msg := "您访问的资源不存在"
	if len(msgs) != 0 {
		msg = msgs[0]
	}
	ctx.NewResponse(CODE_NOT_FOUND, nil, msg)
}

func CreateHandlerFunc(handle HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			c,
		}
		handle(ctx)
	}
}
