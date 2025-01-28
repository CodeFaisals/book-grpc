package consumer

import (
	"context"
	"github.com/BlazeCode1/book-grpc/app/book/repository"
	"log"

	"github.com/segmentio/kafka-go"
)

func StartConsumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "book-events",
		GroupID: "book-group",
	})
	defer reader.Close()

	log.Println("Kafka consumer started")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		id := string(msg.Key)
		bookName := string(msg.Value)
		// todo: put it in service not from database directly
		// Update the book in the database
		err = repository.UpdateBook(id, bookName)
		if err != nil {
			log.Printf("Error updating book: %v", err)
			continue
		}

		log.Printf("Successfully updated book %s with name %s", id, bookName)
	}
}
