package handler

import (
	"github.com/gin-gonic/gin"
	"online.shop.autmaple.com/cmd/web/service/menu"
	"online.shop.autmaple.com/internal/configs/log"
	"online.shop.autmaple.com/internal/response"
)

func QueryMenu(c *gin.Context) {
	menus, err := menu.QueryMenu()
	if err != nil {
		response.ServerError(c)
		return
	}
	response.Ok(c, menus)
}

func InsertMenus(c *gin.Context) {
	var menus []*menu.Dto
	err := c.ShouldBindJSON(&menus)
	if err != nil {
		log.Error(err, err.Error())
		response.UnprocessableEntiy(c)
		return
	}
	err = menu.InsertMenus(menus)
	if err != nil {
		response.ServerError(c)
		return
	}
	response.OkWithMessage(c, "Insert Successful")
}
