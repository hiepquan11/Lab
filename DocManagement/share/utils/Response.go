package utils

import (
	"github.com/gin-gonic/gin"
)

// SuccessResponse success response
func SuccessResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

// ErrorResponse failed response
func ErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    nil,
	})
}
