package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoUser string
	MongoPassword string
	MongoURI string
	ServerPort string
}

func LoadConfig() *Config {
    rootDir, err := filepath.Abs("../../")
    if err != nil {
        log.Fatal("Failed to determine root directory:", err)
    }

    // Load the .env file from the root directory
    err = godotenv.Load(filepath.Join(rootDir, ".env"))
    if err != nil {
        log.Println("No .env file found, using system environment variables")
    }
	mongoUser := os.Getenv("MONGO_USER")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	serverPort := os.Getenv("SERVER_PORT")

	if mongoUser == "" || mongoPassword == "" {
		log.Fatal("MongoDB credentials are not set in environment variables")
	}

	mongoURI := "mongodb://" + mongoUser + ":" + mongoPassword + "@localhost:27017/?authSource=admin"

	return &Config{
		MongoUser: mongoUser,
		MongoPassword: mongoPassword,
		MongoURI: mongoURI,
		ServerPort: serverPort,
	}
}