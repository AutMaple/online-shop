package handler

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"online.shop.autmaple.com/cmd/web/service/spu"
	"online.shop.autmaple.com/internal/models"
	"online.shop.autmaple.com/internal/response"
)

// QuerySpu will handle the `GET /spu/:id` request
func QuerySpu(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.InvalidParam(c, "id")
		return
	}
	spuDto, err := spu.QuerySpu(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.NotFound(c)
			return
		}
		response.ServerError(c)
		return
	}
	response.Ok(c, spuDto)
}

// PageQuerySpu will handle the `GET /spu?offset=10&size=20` request
func PageQuerySpu(c *gin.Context) {
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "1"))
	if err != nil {
		response.InvalidParam(c, "offset")
		return
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "20"))
	if err != nil {
		response.UnprocessableEntiy(c)
		return
	}
	spuDtoList, err := spu.PageQuerySpu(offset, size)
	if err != nil {
		response.ServerError(c)
		return
	}
	if len(spuDtoList) == 0 {
		response.NotFound(c)
		return
	}
	response.Ok(c, spuDtoList)
}

// InsertSpu will handle the `POST /spu` request
func InsertSpu(c *gin.Context) {
	var spuForm spu.Form
	if err := c.ShouldBind(&spuForm); err != nil {
		response.UnprocessableEntiy(c)
		return
	}
	if err := spu.InsertSpu(&spuForm); err != nil {
		response.ServerError(c)
		return
	}
	response.OkWithMessage(c, "Insert Successful")
}

// DeleteSpu will handle the `DELETE /spu/:id` request
func DeleteSpu(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.InvalidParam(c, "id")
		return
	}
	err = spu.DeleteSpu(id)
	if err != nil {
		if errors.Is(err, models.ErrRecordNotFound) {
			response.NotFound(c)
			return
		}
		response.ServerError(c)
		return
	}
	response.OkWithMessage(c, "Delete Successful")
}
