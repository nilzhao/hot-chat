// 鉴权
package middleware

import (
	"errors"
	"red-server/utils"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("X-Token")
		if token == "" {
			utils.ResFailed(ctx, errors.New("请先登录"), utils.CODE_UNAUTHORIZED)
			ctx.Abort()
			return
		}

		claims, err := utils.ParseToken(token)
		if claims == nil || err != nil {
			utils.ResFailed(ctx, errors.New("token已过期"), utils.CODE_UNAUTHORIZED)
			ctx.Abort()
			return
		}
		utils.SetUser(ctx, &claims.User)
		ctx.Next()
	}
}
