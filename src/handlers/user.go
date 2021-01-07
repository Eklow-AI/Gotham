package handlers

import (
	"net/http"

	"github.com/Eklow-AI/Gotham/src/models"
	"github.com/gin-gonic/gin"
)

// CreateUser creates a new user instance in the database
func CreateUser() gin.HandlerFunc {
	opts := models.NewUserOptions{}
	return func(c *gin.Context) {
		if err := c.BindJSON(&opts); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}

		if _, err := models.InsertNewUser(opts); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}
		opts = models.NewUserOptions{}
		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}

//UpdateUserUtype updates the utype of an existing user
func UpdateUserUtype() gin.HandlerFunc {
	opts := models.UpdateUserOptions{}
	return func(c *gin.Context) {
		if err := c.BindJSON(&opts); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}

		user := models.GetUser(opts.Email)
		if err := user.SetUtype(opts.Utype); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}
		opts = models.UpdateUserOptions{}
		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}
