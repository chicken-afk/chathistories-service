package main

import (
	"context"
	"time"

	"github.com/olivere/elastic/v7"
)

type ChatMessage struct {
	RoomID    string    `json:"room_id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Message   string    `json:"message"`
}

func connectElasticSearch(url, username, password string) (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetBasicAuth(username, password),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func storeChatMessage(client *elastic.Client, roomID string, message ChatMessage) error {
	message.RoomID = roomID
	ctx := context.Background()
	_, err := client.Index().
		Index("chat_histories").
		BodyJson(message).
		Do(ctx)
	if err != nil {
		return err
	}
	return nil
}
