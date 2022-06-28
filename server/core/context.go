package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

type ResponseCode int
type HandlerFunc func(c *Context)

const (
	SuccessCode      ResponseCode = 0
	BadRequestCode   ResponseCode = 400
	UnauthorizedCode ResponseCode = 401
	ForbiddenCode    ResponseCode = 403
	NotFoundCode     ResponseCode = 404
)

func (ctx *Context) ResOk(data any) {
	ctx.JSON(http.StatusOK, map[string]any{
		"code":    SuccessCode,
		"data":    data,
		"message": "success",
		"errors":  nil,
	})
}

type ResponseFailInfo struct {
	Code   ResponseCode
	Errors []error
}

type ResponseBody struct {
	ResponseFailInfo
	Data    any
	Message string
}

func (ctx *Context) ResFail(message string, infos ...ResponseFailInfo) {
	info := ResponseFailInfo{}
	if len(infos) != 0 {
		info = infos[0]
	}
	code := info.Code
	if code == 0 {
		code = BadRequestCode
	}
	ctx.JSON(http.StatusOK, map[string]any{
		"code":    code,
		"data":    nil,
		"message": message,
		"errors":  info.Errors,
	})
}

func (ctx *Context) ResNotFound(messages ...string) {
	message := "您访问的资源不存在"
	if len(messages) != 0 {
		message = messages[0]
	}
	ctx.JSON(http.StatusOK, map[string]any{
		"code":    NotFoundCode,
		"data":    nil,
		"message": message,
		"errors":  nil,
	})
}

func CreateHandlerFunc(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			c,
		}
		h(ctx)
	}
}
