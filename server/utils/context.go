package utils

import (
	"hot-chat/model"

	"github.com/gin-gonic/gin"
)

const USER_CONTEXT_KEY = "user"

func SetCurrentUser(ctx *gin.Context, user *model.User) {
	if ctx == nil || user == nil {
		return
	}
	ctx.Set(USER_CONTEXT_KEY, user)
}

func GetCurrentUser(ctx *gin.Context) *model.User {
	val, ok := ctx.Get(USER_CONTEXT_KEY)
	if !ok {
		return nil
	}
	user, ok := val.(*model.User)
	if !ok {
		return nil
	}

	return user
}
