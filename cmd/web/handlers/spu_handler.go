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

// QuerySpu will handle the `GET /spu/:id` request
func QuerySpu(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
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
			handlerutil.RecordNotFoundError(c, err)
			return
		}
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, spuDto)
}

// PageQuerySpu will handle the `GET /spu?offset=10&size=20` request
func PageQuerySpu(c *gin.Context) {
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "1"))
	if err != nil {
		log.Error(err, "")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid Param: offset",
		})
		return
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "20"))
	if err != nil {
		log.Error(err, "")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid Param: size",
		})
		return
	}
	spuDtoList, err := services.PageQuerySpu(offset, size)
	if err != nil {
		handlerutil.ServerError(c, err)
		return
	}
	if len(spuDtoList) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": handlerutil.MsgRecordNotFound,
		})
		return
	}
	c.JSON(http.StatusOK, spuDtoList)
}

// InsertSpu will handle the `POST /spu` request
func InsertSpu(c *gin.Context) {
	var spuForm dto.SpuForm
	if err := c.ShouldBind(&spuForm); err != nil {
		log.Error(err, "")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": http.StatusText(http.StatusUnprocessableEntity),
		})
		return
	}
	if err := services.InsertSpu(&spuForm); err != nil {
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Insert Successful",
	})
}

// DeleteSpu will handle the `DELETE /spu/:id` request
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
