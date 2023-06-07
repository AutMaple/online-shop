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

// QueryUser will handle `GET /user/:id` request
func QueryUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		response.InvalidParam(c, "id")
		return
	}
	user, err := services.QueryUser(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.NotFound(c)
			return
		}
		response.ServerError(c)
		return
	}
	response.Ok(c, user)
}

// PageQueryUser will handle `get /user?offset=1&size=10` request
func PageQueryUser(c *gin.Context) {
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
	users, err := services.PageQueryUser(offset, size)
	if err != nil {
		response.ServerError(c)
		return
	}
	if len(users) == 0 {
		response.NotFound(c)
		return
	}
	response.Ok(c, users)
}

// InsertUser will handle `post /user` request
func InsertUser(c *gin.Context) {
	var user dto.UserForm
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.UnprocessableEntiy(c)
		return
	}
	err = services.InsertUser(&user)
	if err != nil {
		response.ServerError(c)
		return
	}
	response.Ok(c, "Insert Successful")
}

// UpdateUser will handle `put /user/:id` request
func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		response.InvalidParam(c, "id")
		return
	}
	var user dto.UserForm
	err = c.ShouldBindJSON(&user)
	if err != nil {
		response.UnprocessableEntiy(c)
		return
	}
	err = services.UpdateUser(id, &user)
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

// DeleteUser will handle `delete /user/:id` request
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		response.UnprocessableEntiy(c)
		return
	}
	err = services.DeleteUser(id)
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
