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
	c.GET("/sku/:id", handlers.QuerySku)
	c.POST("/sku", handlers.InsertSku)
	c.GET("/sku", handlers.PageQuerySku)
	c.PUT("/sku/:id", handlers.UpdateSku)
	c.DELETE("/sku/:id", handlers.DeleteSku)

	// brand
	c.GET("/brand/:id", handlers.QueryBrand)
	c.POST("/brand", handlers.InsertBrand)
	c.GET("/brand", handlers.PageQueryBrand)
	c.PUT("/brand/:id", handlers.UpdateBrand)
	c.DELETE("/brand/:id", handlers.DeleteBrand)

	// category
	c.GET("/category/:id", handlers.QueryCategory)
	c.POST("/category", handlers.InsertCategory)
	c.GET("/category", handlers.PageQueryCategory)
	c.PUT("/category/:id", handlers.UpdateCategory)
	c.DELETE("/category/:id", handlers.DeleteCategory)

	// store
	c.GET("/store/:id", handlers.QueryStore)
	c.POST("/store", handlers.InsertStore)
	c.GET("/store", handlers.PageQueryStore)
	c.PUT("/store/:id", handlers.UpdateStore)
	c.DELETE("/store/:id", handlers.DeleteStore)
}
