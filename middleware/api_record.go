package middleware

import (
	"github.com/gin-gonic/gin"
)

func SaveApiRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
