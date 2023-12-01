package response

import (
	"github.com/gin-gonic/gin"
)

type JsonStruct struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func ReturnSuccess(c *gin.Context, code int, message interface{}, data interface{}) {
	json := JsonStruct{Code: code, Data: data, Message: message}
	c.JSON(200, json)
}

func ReturnError(c *gin.Context, code int, message interface{}) {
	json := JsonStruct{Code: code, Message: message}
	c.JSON(400, json)
}
