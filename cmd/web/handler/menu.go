package handler

import (
	"github.com/gin-gonic/gin"
	"online.shop.autmaple.com/cmd/web/service/menu"
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
