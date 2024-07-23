package config

import (
	"log"
	"os"

	"github.com/joho/godotenv" // Importing the godotenv package to load environment variables
)

// LoadConfig loads the environment variables from the .env file
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}
}

// GetDSN constructs the Data Source Name (DSN) for connecting to the MySQL database
func GetDSN() string {
	return os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
}
