package main

import (
	"net/http"

	"github.com/Eklow-AI/Gotham/src/handlers"
	"github.com/Eklow-AI/Gotham/src/middleware"
	"github.com/Eklow-AI/Gotham/src/models"
	"github.com/Eklow-AI/Gotham/src/sdk"
	"github.com/gin-gonic/gin"
)

func main() {
	//Connect Postgres database
	models.ConnectDB()
	//Setup SDK
	sdk.SetupSDK()

	// Set up routing
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "stonks"})
	})
	// Admin routes that control the Gotham API
	admin := router.Group("/admin", middleware.RequireAdmin())
	{
		admin.POST("/createOrg", handlers.CreateOrg())
	}

	// Private routes that only authorized Gotham projects can access
	private := router.Group("/private", middleware.CheckToken())
	{
		private.POST("getScore", handlers.GetScore())
	}
	router.Run()
}
