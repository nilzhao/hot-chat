package controller

import (
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUserController() *User {
	return &User{}
}

func (c *User) RegisterRoute(api *gin.RouterGroup) {
}
