package main

import (
	"log"
	"time"

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

	models.ConnectDB()

	router := gin.Default()
	router.Run()
}
