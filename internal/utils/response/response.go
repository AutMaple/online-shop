package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func OkWithMessage(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}

func ServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": http.StatusText(http.StatusInternalServerError),
	})
}

func InvalidParam(c *gin.Context, paramName string) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"message": fmt.Sprintf("Ivalid Param: %s", paramName),
	})
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": http.StatusText(http.StatusNotFound),
	})
}

func UnprocessableEntiy(c *gin.Context) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"message": http.StatusText(http.StatusUnprocessableEntity),
	})
}
