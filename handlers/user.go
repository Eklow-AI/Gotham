package handlers

import (
	"net/http"

	"github.com/Eklow-AI/Gotham/models"
	"github.com/gin-gonic/gin"
)

// CreateUser creates a new user instance in the database
func CreateUser() gin.HandlerFunc {
	opts := models.NewUserOptions{}
	return func(c *gin.Context) {
		err := c.BindJSON(&opts)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}

		err = models.InsertNewUser(opts)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}
		opts = models.NewUserOptions{}
		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}
