package handlers

import (
	"github.com/Eklow-AI/Gotham/models"
	"github.com/gin-gonic/gin"
)

// CreateUser creates a new user instance in the database
func CreateUser() gin.HandlerFunc {
	var opts models.NewUserOptions
	return func(c *gin.Context) {
		token := c.Query("token")
		// check token has permission
		// patron := get patrong from token 
		if c.BindJSON(&opts) == nil {
			opts.Patron = opts.Patron
		}
	}
}
