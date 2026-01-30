package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nekowawolf/MSC/internal/app"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	app.Run()
}