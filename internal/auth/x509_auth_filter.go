package auth

import (
	"github.com/gin-gonic/gin"
)

func X5009AuthFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
