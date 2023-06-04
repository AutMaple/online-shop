package main

import (
	"github.com/gin-gonic/gin"
	"online.shop.autmaple.com/cmd/web/handlers"
)

func RegisterRoutes(c *gin.Engine) {
	// c.GET("/spu/:id", handlers.SingleSpu)
	c.GET("/spu/:id", handlers.QuerySpu)
	c.POST("/spu", handlers.InsertSpu)
	// c.GET("/spu", handlers.PageSpu)
	c.GET("/spu", handlers.PageQuerySpu)
	// c.PUT("/spu/:id", handlers.UpdateSpuHandler)
	// c.DELETE("/spu/:id", handlers.DeleteSpuHandler)
	c.DELETE("/spu/:id", handlers.DeleteSpu)
}
