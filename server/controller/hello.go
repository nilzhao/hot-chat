package controller

import (
	helloService "red-server/service/hello"

	"github.com/gin-gonic/gin"
)

type HelloController struct{}

func NewHelloController() Controller {
	return &HelloController{}
}

func (c *HelloController) Hello(ctx *gin.Context) {
	helloService.SaiHello(ctx)
}

func (c *HelloController) Name() string {
	return "user"
}

func (c *HelloController) RegisterRoute(api *gin.RouterGroup) {
	api.GET("/", c.Hello)
}
