package handlers

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"online.shop.autmaple.com/cmd/web/services"
	"online.shop.autmaple.com/internal/dto"
	"online.shop.autmaple.com/internal/utils/response"
)

// InsertSku will handle the `POST /sku` request
func InsertSku(c *gin.Context) {
	var skuForm dto.SkuForm
	err := c.ShouldBindJSON(&skuForm)
	if err != nil {
		response.UnprocessableEntiy(c)
		return
	}
	err = services.InsertSku(&skuForm)
	if err != nil {
		response.ServerError(c)
		return
	}
	response.OkWithMessage(c, "Insert Successful")
}

// QuerySku will handle the `GET /sku/:id` request
func QuerySku(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		response.InvalidParam(c, "id")
		return
	}
	skuDto, err := services.QuerySku(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.NotFound(c)
			return
		}
		response.ServerError(c)
		return
	}
	response.Ok(c, skuDto)
}

// PageQuerySku will handle the `GET /sku?offset=1&szie=10` request
func PageQuerySku(c *gin.Context) {
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "1"))
	if err != nil || offset <= 0 {
		response.InvalidParam(c, "offset")
		return
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil || size <= 0 {
		response.InvalidParam(c, "size")
		return
	}
	skuDtos, err := services.PageQuerySku(offset, size)
	if err != nil {
		response.ServerError(c)
		return
	}
	if len(skuDtos) == 0 {
		response.NotFound(c)
		return
	}
	response.Ok(c, skuDtos)
}

// UpdateSku will handle the `PUT /sku/:id` request
func UpdateSku(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		response.InvalidParam(c, "id")
		return
	}
	var skuForm dto.SkuForm
	if err := c.ShouldBindJSON(&skuForm); err != nil {
		response.UnprocessableEntiy(c)
		return
	}
	err = services.UpdateSku(id, &skuForm)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.NotFound(c)
			return
		}
		response.ServerError(c)
		return
	}
	response.OkWithMessage(c, "Update Successful")
}

// DeleteSku will handle the `DELETE /sku/:id` request
func DeleteSku(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		response.InvalidParam(c, "id")
		return
	}
	err = services.DeleteSku(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.NotFound(c)
			return
		}
		response.ServerError(c)
		return
	}
	response.OkWithMessage(c, "Delete Successful")
}
