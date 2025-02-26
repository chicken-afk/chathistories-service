package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"encoding/json"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func rabbitMqConnect() {
	var rabbitHost = os.Getenv("RABITMQ_HOST")
	fmt.Println("Connecting to RabbitMQ")
	fmt.Println(rabbitHost)
	conn, err := amqp.Dial(rabbitHost)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"queue_chat_histories",               // name
		true,                                 // durable
		false,                                // delete when unused
		false,                                // exclusive
		false,                                // no-wait
		amqp.Table{"x-queue-type": "quorum"}, // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var message ChatMessage
			err := json.Unmarshal(d.Body, &message)
			if err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}
			loc, err := time.LoadLocation("Asia/Bangkok")
			if err != nil {
				log.Printf("Failed to load location: %v", err)
				continue
			}
			//Get date now
			message.CreatedAt = time.Now()
			message.CreatedAt = message.CreatedAt.In(loc)

			// Store message to elastic
			fmt.Println("Storing message to elastic")
			fmt.Println(elasticHost)
			fmt.Println(elasticUsername)
			fmt.Println(elasticPassword)
			client, err := connectElasticSearch(elasticHost, elasticUsername, elasticPassword)
			if err != nil {
				log.Fatalf("Failed to connect to Elasticsearch: %v", err)
			}

			err = storeChatMessage(client, message.RoomID, message)
			if err != nil {
				log.Fatalf("Failed to store chat message: %v", err)
			}
		}
	}()

	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
