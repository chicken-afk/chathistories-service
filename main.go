package main

var elasticHost = "http://localhost:9200"
var elasticUsername = "elastic"
var elasticPassword = "yourpassword"

func main() {
	// client, err := connectElasticSearch(elasticHost, elasticUsername, elasticPassword)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to Elasticsearch: %v", err)
	// }

	// message := ChatMessage{
	// 	Email:     "user2@example.com",
	// 	CreatedAt: time.Now(),
	// 	Message:   "Hallo jugaa!",
	// }

	// err = storeChatMessage(client, "room1", message)
	// if err != nil {
	// 	log.Fatalf("Failed to store chat message: %v", err)
	// }

	// log.Println("Chat message stored successfully")
	rabbitMqConnect()
}
