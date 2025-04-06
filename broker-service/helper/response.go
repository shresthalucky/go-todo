package helper

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Test  string      `json:"test,omitempty"`
}

func SuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	res := Response{Data: data}
	c.JSON(statusCode, res)
}

func ErrorResponse(c *gin.Context, statusCode int, err error) {
	res := Response{Error: err.Error(), Data: nil}
	c.JSON(statusCode, res)
}
