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

// QueryStore will handle `GET /store/:id` request
func QueryStore(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		response.InvalidParam(c, "id")
		return
	}
	store, err := services.QueryStore(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.NotFound(c)
			return
		}
		response.ServerError(c)
		return
	}
	response.Ok(c, store)
}

// PageQueryStore will handle `GET /store?offset=1&size=10` request
func PageQueryStore(c *gin.Context) {
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "1"))
	if err != nil {
		response.InvalidParam(c, "offset")
		return
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil {
		response.InvalidParam(c, "size")
		return
	}
	stores, err := services.PageQueryStore(offset, size)
	if err != nil {
		response.ServerError(c)
		return
	}
	if len(stores) == 0 {
		response.NotFound(c)
		return
	}
	response.Ok(c, stores)
}

// InsertStore will handle `POST /store` request
func InsertStore(c *gin.Context) {
	var store dto.StoreForm
	err := c.ShouldBindJSON(&store)
	if err != nil {
		response.UnprocessableEntiy(c)
		return
	}
	err = services.InsertStore(&store)
	if err != nil {
		response.ServerError(c)
		return
	}
	response.OkWithMessage(c, "Insert Successful")
}

// UpdateStore will handle `PUT /store/:id` request
func UpdateStore(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		response.InvalidParam(c, "id")
		return
	}
	var store dto.StoreForm
	err = c.ShouldBindJSON(&store)
	if err != nil {
		response.UnprocessableEntiy(c)
		return
	}
	err = services.UpdateStore(id, &store)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.ServerError(c)
			return
		}
		response.ServerError(c)
		return
	}
	response.OkWithMessage(c, "Update Successful")
}

// DeleteStore will handle `DELETE /store/:id` request
func DeleteStore(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		response.InvalidParam(c, "id")
		return
	}
	err = services.DeleteStore(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.NotFound(c)
			return
		}
		response.ServerError(c)
		return
	}
	response.OkWithMessage(c, "Delte Successful")
}
