package utils

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

type ResponseBody struct {
	Code ResCode     `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(ctx *gin.Context, code ResCode, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, ResponseBody{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func ResOk(ctx *gin.Context, data interface{}) {
	NewResponse(ctx, CODE_SUCCESS, data, "ok")
}

func ResFailed(ctx *gin.Context, err error, codes ...ResCode) {
	code := CODE_BAD_REQUEST
	if len(codes) != 0 {
		code = codes[0]
	}
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	NewResponse(ctx, code, nil, msg)
}

func ResNotFound(ctx *gin.Context, msgs ...string) {
	msg := "您访问的资源不存在"
	if len(msgs) != 0 {
		msg = msgs[0]
	}
	NewResponse(ctx, CODE_NOT_FOUND, nil, msg)
}
