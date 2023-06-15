package main

import (
	"github.com/gin-gonic/gin"
	"online.shop.autmaple.com/cmd/web/handler"
)

func RegisterRoutes(c *gin.Engine) {
	// spu
	c.GET("/spu/:id", handler.QuerySpu)
	c.POST("/spu", handler.InsertSpu)
	c.GET("/spu", handler.PageQuerySpu)
	c.DELETE("/spu/:id", handler.DeleteSpu)

	// sku
	c.GET("/sku/:id", handler.QuerySku)
	c.POST("/sku", handler.InsertSku)
	c.GET("/sku", handler.PageQuerySku)
	c.PUT("/sku/:id", handler.UpdateSku)
	c.DELETE("/sku/:id", handler.DeleteSku)

	// brand
	c.GET("/brand/:id", handler.QueryBrand)
	c.POST("/brand", handler.InsertBrand)
	c.GET("/brand", handler.PageQueryBrand)
	c.PUT("/brand/:id", handler.UpdateBrand)
	c.DELETE("/brand/:id", handler.DeleteBrand)

	// category
	c.GET("/category/:id", handler.QueryCategory)
	c.POST("/category", handler.InsertCategory)
	c.GET("/category", handler.PageQueryCategory)
	c.PUT("/category/:id", handler.UpdateCategory)
	c.DELETE("/category/:id", handler.DeleteCategory)

	// store
	c.GET("/store/:id", handler.QueryStore)
	c.POST("/store", handler.InsertStore)
	c.GET("/store", handler.PageQueryStore)
	c.PUT("/store/:id", handler.UpdateStore)
	c.DELETE("/store/:id", handler.DeleteStore)

	// user
	c.GET("/user/:id", handler.QueryUser)
	c.POST("/user", handler.InsertUser)
	c.GET("/user", handler.PageQueryUser)
	c.PUT("/user/:id", handler.UpdateUser)
	c.DELETE("/user/:id", handler.DeleteUser)

	// menu
	c.GET("/menu", handler.QueryMenu)
  c.POST("/menu/batch", handler.InsertMenus)
}
