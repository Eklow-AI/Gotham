package handlers

import (
	"net/http"

	"github.com/Eklow-AI/Gotham/src/sdk"
	"github.com/gin-gonic/gin"
)

type baseRequest struct {
	Cage string `json:"cage" binding:"required"`
	Cid  string `json:"cid" binding:"required"`
}

// GetScore handler gets a single compatability score for a vendor
func GetScore() gin.HandlerFunc {
	params := baseRequest{}
	return func(c *gin.Context) {
		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}
		score := sdk.GetScore(params.Cage, params.Cid)
		params = baseRequest{}
		c.JSON(http.StatusOK, gin.H{"success": true, "score": score})
	}
}
