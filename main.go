package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"bikrent/routes"
	"bikrent/models"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	loadEnv()

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/bikrentgoapp?parseTime=true", dbUser, dbPassword)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open a database connection: %v", err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	db.AutoMigrate(&models.User{}, &models.Bicycle{}, &models.Rental{}, &models.UserDetail{})

	e := echo.New()
	routes.SetupRoutes(e, db)

	e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORS())

	e.Start(":80")
}
