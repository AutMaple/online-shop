package handler

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"online.shop.autmaple.com/cmd/web/service/category"
	"online.shop.autmaple.com/internal/response"
)

// QueryCategory will handle `GET /category/:id` request
func QueryCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		response.InvalidParam(c, "id")
		return
	}
	categoryDto, err := category.QueryCategory(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.NotFound(c)
			return
		}
		response.ServerError(c)
		return
	}
	response.Ok(c, categoryDto)
}

// PageQueryCategory will handle `GET /category?offset=1&size=10`
func PageQueryCategory(c *gin.Context) {
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
	categorys, err := category.PageQueryCategory(offset, size)
	if err != nil {
		response.ServerError(c)
		return
	}
	if len(categorys) == 0 {
		response.NotFound(c)
		return
	}
	response.Ok(c, categorys)
}

// InsertCategory will hanle `POST /category` request
func InsertCategory(c *gin.Context) {
	var categoryForm category.Form
	err := c.ShouldBindJSON(&categoryForm)
	if err != nil {
		response.UnprocessableEntiy(c)
		return
	}
	err = category.InsertCategory(&categoryForm)
	if err != nil {
		response.ServerError(c)
		return
	}
	response.OkWithMessage(c, "Insert Successful")
}

// UpdateCategory will handle `PUT /category/:id` request
func UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.InvalidParam(c, "id")
		return
	}
	var categoryForm category.Form
	err = c.ShouldBindJSON(&categoryForm)
	if err != nil {
		response.UnprocessableEntiy(c)
		return
	}
	err = category.UpdateCategory(id, &categoryForm)
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

// DeleteCategory will handle `DELETE /category/:id` request
func DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.InvalidParam(c, "id")
		return
	}
	err = category.DeleteCategory(id)
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
