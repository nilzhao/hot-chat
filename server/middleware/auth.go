// 鉴权
package middleware

import (
	"errors"
	"red-server/utils"

	"github.com/gin-gonic/gin"
)

func getToken(ctx *gin.Context) string {
	const tokenKey = "X-Token"
	token := ctx.Request.Header.Get(tokenKey)
	if token == "" {
		token = ctx.Query(tokenKey)
	}
	return token
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := getToken(ctx)
		if token == "" {
			utils.ResFailed(ctx, errors.New("请先登录"), utils.CODE_UNAUTHORIZED)
			ctx.Abort()
			return
		}

		claims, err := utils.ParseToken(token)
		if claims == nil || err != nil {
			utils.ResFailed(ctx, errors.New("无效的 token"), utils.CODE_UNAUTHORIZED)
			ctx.Abort()
			return
		}
		utils.SetCurrentUser(ctx, &claims.User)
		ctx.Next()
	}
}
