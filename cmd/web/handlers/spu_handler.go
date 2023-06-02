package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"online.shop.autmaple.com/cmd/web/services"
	"online.shop.autmaple.com/internal/configs/log"
	"online.shop.autmaple.com/internal/dto"
	"online.shop.autmaple.com/internal/models"
	"online.shop.autmaple.com/internal/utils/handlerutil"
)

// /spu/:id
func QuerySpu(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error(err, "")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": handlerutil.MsgInvalidId,
		})
		return
	}
	spuDto, err := services.QuerySpu(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": handlerutil.MsgRecordNotFound,
			})
			return
		}
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, spuDto)
}

// GET /spu?offset=10&size=20
func PageQuerySpu(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", "1")
	sizeStr := c.DefaultQuery("size", "20")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		log.Error(err, "")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid Param: size",
		})
		return
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		log.Error(err, "")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid Param: offset",
		})
		return
	}
	spuDtoList, err := services.PageQuerySpu(offset, size)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handlerutil.RecordNotFoundError(c)
			return
		}
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, spuDtoList)
}

// POST /spu
func InsertSpuHandler(c *gin.Context) {
	var spuForm dto.SpuForm
	err := c.Bind(&spuForm)
	if err != nil {
		log.Error(err, "")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": http.StatusText(http.StatusUnprocessableEntity),
		})
		return
	}
	services.InsertSpu(&spuForm)
	c.JSON(http.StatusOK, gin.H{
		"message": "Insert Successful",
	})
}

// DELETE /spu/:id
func DeleteSpu(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error(err, "")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": handlerutil.MsgInvalidId,
		})
		return
	}
	err = services.DeleteSpu(id)
	if err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			handlerutil.RecordNotFoundError(c)
			return
		}
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Successful",
	})
}
