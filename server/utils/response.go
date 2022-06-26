package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseCode int

const (
	SuccessCode      ResponseCode = 0
	BadRequestCode   ResponseCode = 400
	UnauthorizedCode ResponseCode = 401
	ForbiddenCode    ResponseCode = 403
	NotFoundCode     ResponseCode = 404
)

func ResOk(ctx *gin.Context, data any) {
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

func ResFail(ctx *gin.Context, message string, infos ...ResponseFailInfo) {
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
