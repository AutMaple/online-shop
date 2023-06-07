package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"online.shop.autmaple.com/cmd/web/services"
	"online.shop.autmaple.com/internal/dto"
	"online.shop.autmaple.com/internal/utils/response"
)

// QueryBrand will handle `GET /brand/:id` request
func QueryBrand(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.InvalidParam(c, "id")
		return
	}
	brandDto, err := services.QueryBrand(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.NotFound(c)
			return
		}
		response.ServerError(c)
		return
	}
	response.Ok(c, brandDto)
}

// PageQueryBrand will handle `GET /brand?offset=1&size=10` request
func PageQueryBrand(c *gin.Context) {
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil || offset == 0 {
		response.InvalidParam(c, "offset")
		return
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil || size == 0 {
		response.InvalidParam(c, "size")
		return
	}

	brandDtos, err := services.PageQueryBrand(offset, size)
	if err != nil {
		response.ServerError(c)
		return
	}
	if len(brandDtos) == 0 {
		response.NotFound(c)
		return
	}
	c.JSON(http.StatusOK, brandDtos)
}

// InsertBrand will handle `PoST /brand` request
func InsertBrand(c *gin.Context) {
	var brandForm dto.BrandForm
	err := c.ShouldBindJSON(&brandForm)
	if err != nil {
		response.UnprocessableEntiy(c)
		return
	}
	err = services.InsertBrand(&brandForm)
	if err != nil {
		response.ServerError(c)
		return
	}
	response.OkWithMessage(c, "Insert Successful")
}

// UpdateBrand will handle `PUT /brand/:id` request
func UpdateBrand(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.InvalidParam(c, "id")
		return
	}
	var brandForm dto.BrandForm
	err = c.ShouldBindJSON(&brandForm)
	if err != nil {
		response.UnprocessableEntiy(c)
		return
	}
	err = services.UpdateBrand(id, &brandForm)
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

// DeleteBrand will handle `DELETE /brand/:id` request
func DeleteBrand(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.InvalidParam(c, "id")
		return
	}
	err = services.DeleteBrand(id)
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
