package helloService

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaiHello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello  World!",
	})
}
