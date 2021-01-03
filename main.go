package main

import (
	"log"
	"time"

	"github.com/Eklow-AI/Gotham/handlers"
	"github.com/Eklow-AI/Gotham/middleware"
	"github.com/Eklow-AI/Gotham/models"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variable on dev environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Set up Sentry
	err = sentry.Init(sentry.ClientOptions{
		Dsn: "https://1e8d7ea2192a4f949bf5e878cfb2124e@o496200.ingest.sentry.io/5570172",
	})
	if err != nil {
		log.Fatal("sentry.Init:", err)
	}
	defer sentry.Flush(2 * time.Second)
	//Connect Postgres database
	models.ConnectDB()

	// Set up routing
	router := gin.Default()
	// Admin routes that control the Gotham API
	admin := router.Group("/admin", middleware.RequireAdmin())
	{
		admin.POST("/createOrg", handlers.CreateOrg())
	}

	// Private routes that only authorized Gotham projects can access
	private := router.Group("/private", middleware.CheckToken())
	{
		private.POST("/createUser", handlers.CreateUser())
		private.POST("/updateUtype", handlers.UpdateUserUtype())
	}
	router.Run()
}
