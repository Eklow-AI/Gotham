package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB instance
var DB *gorm.DB

// ConnectDB establishes a postgres database connection by reference
func ConnectDB() {
	connURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", os.Getenv("dbHost"),
		os.Getenv("dbPort"), os.Getenv("dbUser"), os.Getenv("dbName"), os.Getenv("dbPassword"), os.Getenv("sslmode"))
	database, err := gorm.Open("postgres", connURI)

	if err != nil {
		log.Fatal("Problem connecting to database:", err)
	}
	//Migrate models
	database.AutoMigrate(&User{})
	database.AutoMigrate(&Org{})
	DB = database
}
