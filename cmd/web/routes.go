package main

import (
	"github.com/gin-gonic/gin"
	"online.shop.autmaple.com/cmd/web/handlers"
)

func RegisterRoutes(c *gin.Engine) {
  c.GET("/spu", handlers.FetchCreateSpuPage)
  c.POST("/spu", handlers.InsertSpuHandler)
}
