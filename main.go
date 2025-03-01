package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// log.Println("Chat message stored successfully")
	//initiate godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// connect to rabbitmq
	rabbitMqConnect()
}
