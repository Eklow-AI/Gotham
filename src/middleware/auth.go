package middleware

import (
	"net/http"

	"github.com/Eklow-AI/Gotham/src/models"
	"github.com/gin-gonic/gin"
)

// CheckToken validates the given token is valid
func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if !models.GetOrg(token).IsValid {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "401 - Invalid token"})
			c.AbortWithStatus(401)
		}
	}
}

// RequireAdmin requires API key to be admin level
// this also checks that the API key isValid
// this way we lower the number of database reads to only one
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		org := models.GetOrg(token)
		if !org.IsValid {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "401 - Invalid token"})
			c.AbortWithStatus(401)
			return
		}
		if org.Clearance < 3 {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "403 - Endpoint requires admin priveleges"})
			c.AbortWithStatus(403)
		}
	}
}
