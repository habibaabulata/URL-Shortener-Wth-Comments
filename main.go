package main

import (
	// Importing configuration, controllers & database package
	"url-shortener/config"
	"url-shortener/controllers"
	"url-shortener/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig() // Load environment variables from .env file
	database.InitDB()   // Initialize the database connection

	// Create a new Gin router
	r := gin.Default()

	// Register routes for authentication and URL shortening
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/shorten", controllers.ShortenURL)
	r.GET("/:short_code", controllers.GetOriginalURL)

	r.Run(":8080")
}
