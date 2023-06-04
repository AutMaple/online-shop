package main

import (
	"github.com/gin-gonic/gin"
	"online.shop.autmaple.com/cmd/web/handlers"
)

func RegisterRoutes(c *gin.Engine) {
	// spu
	c.GET("/spu/:id", handlers.QuerySpu)
	c.POST("/spu", handlers.InsertSpu)
	c.GET("/spu", handlers.PageQuerySpu)
	c.DELETE("/spu/:id", handlers.DeleteSpu)

	// sku
	c.POST("/sku", handlers.InsertSku)
}
