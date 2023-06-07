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

// QueryUser will handle `GET /user/:id` request
func QueryUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": handlerutil.MsgInvalidId,
		})
		return
	}
	user, err := services.QueryUser(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			handlerutil.RecordNotFoundError(c, err)
			return
		}
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

// PageQueryUser will handle `get /user?offset=1&size=10` request
func PageQueryUser(c *gin.Context) {
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "1"))
	if err != nil || offset <= 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Ivalid param: offset",
		})
		return
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil || size <= 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Ivalid param: size",
		})
		return
	}
	users, err := services.PageQueryUser(offset, size)
	if err != nil {
		handlerutil.ServerError(c, err)
		return
	}
	if len(users) == 0 {
		handlerutil.RecordNotFoundError(c, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

// InsertUser will handle `post /user` request
func InsertUser(c *gin.Context) {
	var user dto.UserForm
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": http.StatusText(http.StatusUnprocessableEntity),
		})
		return
	}
	err = services.InsertUser(&user)
	if err != nil {
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Insert Successful",
	})
}

// UpdateUser will handle `put /user/:id` request
func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": handlerutil.MsgInvalidId,
		})
		return
	}
	var user dto.UserForm
	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": http.StatusText(http.StatusUnprocessableEntity),
		})
		return
	}
	err = services.UpdateUser(id, &user)
	if err != nil {
		handlerutil.ServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Successful",
	})
}

// DeleteUser will handle `delete /user/:id` request
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": handlerutil.MsgInvalidId,
		})
		return
	}
	err = services.DeleteUser(id)
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
