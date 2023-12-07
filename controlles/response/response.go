package response

import (
	"github.com/gin-gonic/gin"
	"test/domain"
)

func ReturnSuccess(c *gin.Context, code int, message interface{}, data interface{}) {
	json := domain.Response{Code: code, Data: data, Message: message}
	c.JSON(200, json)
}

func ReturnError(c *gin.Context, code int, message interface{}) {
	json := domain.Response{Code: code, Message: message}
	c.JSON(400, json)
}
