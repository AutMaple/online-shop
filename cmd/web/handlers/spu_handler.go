package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"online.shop.autmaple.com/cmd/web/services"
	"online.shop.autmaple.com/internal/configs/log"
	"online.shop.autmaple.com/internal/dto"
)

func InsertSpuHandler(c *gin.Context) {
	var spuDto dto.SpuDto
	err := c.Bind(&spuDto)
	if err != nil {
		log.Error(err, "")
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": http.StatusText(http.StatusUnprocessableEntity),
		})
		return
	}
  services.InsertSpu(&spuDto)
	c.JSON(http.StatusOK, gin.H{
		"message": "Insert successFul",
	})
}
