package main

import (
	"github.com/gin-gonic/gin"
	"online.shop.autmaple.com/cmd/web/handlers"
)

func RegisterRoutes(c *gin.Engine) {
	c.GET("/spu/:id", handlers.SingleSpu)
	c.POST("/spu", handlers.InsertSpuHandler)
	c.GET("/spu", handlers.PageSpu)
	c.PUT("/spu/:id", handlers.UpdateSpuHandler)
	c.DELETE("/spu/:id", handlers.DeleteSpuHandler)
}
