package main

import (
    "log"
    "github.com/joho/godotenv"
    "./routers"
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
