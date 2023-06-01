package handlerutil

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"online.shop.autmaple.com/internal/configs/log"
)

const (
	MsgInvalidId      = "Invalid ID"
	MsgInvalidParam   = "Invalid Params"
	MsgRecordNotFound = "Record Not Found"
)

func ServerError(c *gin.Context, err error) {
	log.Error(err, "")
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": http.StatusText(http.StatusInternalServerError),
	})
}

func ClientError(c *gin.Context, status int, err error, msg string) {
	log.Error(err, "")
	c.JSON(status, gin.H{
		"message": msg,
	})
}
