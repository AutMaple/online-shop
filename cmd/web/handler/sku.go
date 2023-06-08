package handler

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"online.shop.autmaple.com/cmd/web/service/sku"
	"online.shop.autmaple.com/internal/response"
)

// InsertSku will handle the `POST /sku` request
func InsertSku(c *gin.Context) {
	var skuForm sku.Form
	err := c.ShouldBindJSON(&skuForm)
	if err != nil {
		response.UnprocessableEntiy(c)
		return
	}
	err = sku.InsertSku(&skuForm)
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
	skuDto, err := sku.QuerySku(id)
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
	skuDtos, err := sku.PageQuerySku(offset, size)
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
	var skuForm sku.Form
	if err := c.ShouldBindJSON(&skuForm); err != nil {
		response.UnprocessableEntiy(c)
		return
	}
	err = sku.UpdateSku(id, &skuForm)
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
	err = sku.DeleteSku(id)
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
