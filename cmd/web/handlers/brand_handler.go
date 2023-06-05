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

// QueryBrand will handle `GET /brand/:id` request
func QueryBrand(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": handlerutil.MsgInvalidId,
		})
		return
	}
	brandDto, err := services.QueryBrand(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handlerutil.RecordNotFoundError(c, err)
			return
		}
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, brandDto)
}

// PageQueryBrand will handle `GET /brand?offset=1&size=10` request
func PageQueryBrand(c *gin.Context) {
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil || offset == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid Param: offset",
		})
		return
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil || size == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid Param: size",
		})
		return
	}

	brandDtos, err := services.PageQueryBrand(offset, size)
	if err != nil {
		handlerutil.ServerError(c, err)
		return
	}
	if len(brandDtos) == 0 {
		handlerutil.RecordNotFoundError(c, err)
		return
	}
	c.JSON(http.StatusOK, brandDtos)
}

// InsertBrand will handle `PoST /brand` request
func InsertBrand(c *gin.Context) {
	var brandForm dto.BrandForm
	err := c.ShouldBindJSON(&brandForm)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": http.StatusText(http.StatusUnprocessableEntity),
		})
		return
	}
	err = services.InsertBrand(&brandForm)
	if err != nil {
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Insert Successful",
	})
}

// UpdateBrand will handle `PUT /brand/:id` request
func UpdateBrand(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": handlerutil.MsgInvalidId,
		})
		return
	}
	var brandForm dto.BrandForm
	err = c.ShouldBindJSON(&brandForm)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": http.StatusText(http.StatusUnprocessableEntity),
		})
		return
	}
	err = services.UpdateBrand(id, &brandForm)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handlerutil.RecordNotFoundError(c, err)
			return
		}
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Successful",
	})
}

// DeleteBrand will handle `DELETE /brand/:id` request
func DeleteBrand(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": handlerutil.MsgInvalidId,
		})
		return
	}
	err = services.DeleteBrand(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handlerutil.RecordNotFoundError(c, err)
			return
		}
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Successful",
	})
}
