package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"online.shop.autmaple.com/cmd/web/services"
	"online.shop.autmaple.com/internal/dto"
	"online.shop.autmaple.com/internal/utils/handlerutil"
)

// InsertSku will handle the `POST /sku` request
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

// QuerySku will handle the `GET /sku/:id` request
func QuerySku(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": handlerutil.MsgInvalidId,
		})
		return
	}
	skuDto, err := services.QuerySku(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handlerutil.RecordNotFoundError(c, err)
			return
		}
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, skuDto)
}

// PageQuerySku will handle the `GET /sku?offset=1&szie=10` request
func PageQuerySku(c *gin.Context) {

}

// UpdateSku will handle the `PUT /sku/:id` request
func UpdateSku(c *gin.Context) {

}

// DeleteSku will handle the `DELETE /sku/:id` request
func DeleteSku(c *gin.Context) {

}
