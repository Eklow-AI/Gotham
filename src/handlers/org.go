package handlers

import (
	"net/http"

	"github.com/Eklow-AI/Gotham/src/models"
	"github.com/gin-gonic/gin"
)

// CreateOrg creates a new user instance in the database
func CreateOrg() gin.HandlerFunc {
	opts := models.NewOrgOptions{}
	return func(c *gin.Context) {
		err := c.BindJSON(&opts)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}

		org, err := models.InsertNewOrg(opts)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}
		opts = models.NewOrgOptions{}
		c.JSON(http.StatusOK, gin.H{"success": true, "token": org.Token})
	}
}
