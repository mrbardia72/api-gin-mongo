package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mrbardia72/api-gin-mongo/routers"
)

func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {
	routers.RouterApp()
}
