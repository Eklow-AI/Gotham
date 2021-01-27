package main

import (
	"net/http"

	"github.com/Eklow-AI/Gotham/src/handlers"
	"github.com/Eklow-AI/Gotham/src/sdk"
	"github.com/gin-gonic/gin"
)

func main() {
	//Connect Postgres database
	//models.ConnectDB()
	//Setup SDK
	sdk.SetupSDK()

	// Set up routing
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "stonks"})
	})

	// Private routes that only authorized Gotham projects can access
	private := router.Group("/private")
	{
		private.POST("/score", handlers.GetScore())
	}
	router.Run()
}
