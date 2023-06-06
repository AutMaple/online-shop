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

// QueryStore will handle `GET /store/:id` request
func QueryStore(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": handlerutil.MsgInvalidId,
		})
		return
	}
	store, err := services.QueryStore(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handlerutil.RecordNotFoundError(c, err)
			return
		}
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, store)
}

// PageQueryStore will handle `GET /store?offset=1&size=10` request
func PageQueryStore(c *gin.Context) {
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "1"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid Param: offset",
		})
		return
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "1"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid Param: size",
		})
		return
	}
	stores, err := services.PageQueryStore(offset, size)
	if err != nil {
		handlerutil.ServerError(c, err)
		return
	}
	if len(stores) == 0 {
		handlerutil.RecordNotFoundError(c, nil)
		return
	}
	c.JSON(http.StatusOK, stores)
}

// InsertStore will handle `POST /store` request
func InsertStore(c *gin.Context) {
	var store dto.StoreForm
	err := c.ShouldBindJSON(&store)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": http.StatusText(http.StatusUnprocessableEntity),
		})
		return
	}
	err = services.InsertStore(&store)
	if err != nil {
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Insert Successful",
	})
}

// UpdateStore will handle `PUT /store/:id` request
func UpdateStore(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": handlerutil.MsgInvalidId,
		})
		return
	}
	var store dto.StoreForm
	err = c.ShouldBindJSON(&store)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": http.StatusText(http.StatusUnprocessableEntity),
		})
		return
	}
	err = services.UpdateStore(id, &store)
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

// DeleteStore will handle `DELETE /store/:id` request
func DeleteStore(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": handlerutil.MsgInvalidId,
		})
		return
	}
	err = services.DeleteStore(id)
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
