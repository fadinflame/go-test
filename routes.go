package main

import (
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func (r routes) customerRoutes(rg *gin.RouterGroup) {
	var c = rg.Group("/customer")
	c.POST("/create", CreateCustomerController)
	c.POST("/get", GetCustomerController)
	c.POST("/deposit", DepositMoneyController)
}

func (r routes) frontendRoutes(rg *gin.RouterGroup) {
	rg.GET("", IndexController)
}

func NewRoutes() routes {
	var r = routes{
		router: gin.Default(),
	}
	r.router.LoadHTMLGlob("templates/*")
	var frontend = r.router.Group("/")
	var api = r.router.Group("/api/v1")
	r.customerRoutes(api)
	r.frontendRoutes(frontend)
	return r
}
