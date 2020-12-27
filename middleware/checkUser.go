package middleware

import (
	"github.com/gin-gonic/gin"
)

// CheckUser checks auth settings
func CheckUser() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
