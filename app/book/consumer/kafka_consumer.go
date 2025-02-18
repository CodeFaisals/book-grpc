package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/BlazeCode1/book-grpc/app/book/model/Book"
	"github.com/BlazeCode1/book-grpc/app/book/service"
	"github.com/segmentio/kafka-go"
)

type KafkaConsumer interface {
	StartConsumer(ctx context.Context) error
}

type kafkaConsumer struct {
	service service.BookService
}

// NewKafkaConsumer New constructor for KafkaConsumer
func NewKafkaConsumer(svc service.BookService) KafkaConsumer {
	return &kafkaConsumer{service: svc}
}

// StartConsumer Updated to accept a context for proper shutdown
func (s *kafkaConsumer) StartConsumer(ctx context.Context) error {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"broker:29092"},
		Topic:   "book-events",
		GroupID: "book-group",
	})
	defer func(reader *kafka.Reader) {
		err := reader.Close()
		if err != nil {
			log.Println(err)
		}
	}(reader)

	log.Println("Kafka consumer started and listening for messages...")

	for {
		select {
		case <-ctx.Done(): //  Graceful shutdown handling
			log.Println("Kafka consumer shutting down...")
			return nil
		default:
			msg, err := reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("Error reading message: %v", err)
				continue
			}

			log.Printf("Received message: Topic=%s, Partition=%d, Offset=%d, Key=%s, Value=%s",
				msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))

			id := string(msg.Key)
			book := Book.Book{}
			err = json.Unmarshal(msg.Value, &book)
			if err != nil {
				log.Printf("Error unmarshalling book: %v", err)
				continue
			}

			log.Printf("Processing book update - ID: %s, BookName: %s, Author: %s", id, book.BookName, book.Author)

			book.ID = id

			// Update book in the database
			response, err := s.service.HandleUpdateBook(id, book)
			if err != nil {
				log.Printf("Error updating book: %v", err)
				continue
			}

			log.Printf("Update successful: %v", response)
		}
	}
}
