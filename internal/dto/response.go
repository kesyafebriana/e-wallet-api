package dto

import "github.com/gin-gonic/gin"

type ResponseMsg struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseSuccessJSON(c *gin.Context, statusCode int, message string, data interface{}) {
	c.AbortWithStatusJSON(statusCode, ResponseMsg{Message: message, Data: data})
}