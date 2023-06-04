package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"online.shop.autmaple.com/cmd/web/services"
	"online.shop.autmaple.com/internal/dto"
	"online.shop.autmaple.com/internal/utils/handlerutil"
)

func InsertSku(c *gin.Context) {
	var skuForm dto.SkuForm
	err := c.ShouldBindJSON(&skuForm)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": http.StatusText(http.StatusUnprocessableEntity),
		})
		return
	}
	err = services.InsertSku(&skuForm)
	if err != nil {
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Sku Insert Successful",
	})
}
