package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var elasticHost = os.Getenv("ELASTIC_HOST")
var elasticUsername = os.Getenv("ELASTIC_USERNAME")
var elasticPassword = os.Getenv("ELASTIC_PASSWORD")

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
