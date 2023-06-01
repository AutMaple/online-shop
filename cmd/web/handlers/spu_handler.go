package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"online.shop.autmaple.com/cmd/web/services"
	"online.shop.autmaple.com/internal/configs/log"
	"online.shop.autmaple.com/internal/dto"
	"online.shop.autmaple.com/internal/models"
	"online.shop.autmaple.com/internal/utils/handlerutil"
)

// GET /spu/:id
func SingleSpu(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error(err, "")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": handlerutil.MsgInvalidId,
		})
		return
	}
	spu := &models.Spu{ID: id}
	err = spu.QueryById(nil)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("Spu %v not found", id),
			})
			return
		}
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, spu)
}

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
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, spuDto)
}

// GET /spu?offset=10&size=20
func PageSpu(c *gin.Context) {
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
	spu := &models.Spu{}
	spuList, err := spu.PageQuery(nil, offset, size)
	if err != nil {
		handlerutil.ServerError(c, err)
		return
	}
	if len(spuList) <= 0 {
		spuList = []*models.Spu{}
	}
	c.JSON(http.StatusOK, spuList)
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
func DeleteSpuHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error(err, "")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": handlerutil.MsgInvalidId,
		})
		return
	}
	var spu = &models.Spu{ID: id}
	err = spu.Delete(nil)
	if err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Successful",
	})
}

// PUT /spu/:id
func UpdateSpuHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error(err, "")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": handlerutil.MsgInvalidId,
		})
		return
	}
	spu := &models.Spu{ID: id}
	err = c.ShouldBindJSON(spu)
	if err != nil {
		log.Error(err, "")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": handlerutil.MsgInvalidParam,
		})
		return
	}
	err = spu.Update(nil)
	if err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Successful",
	})
}
