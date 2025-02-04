package consumer

//kafka code
import (
	"context"
	"encoding/json"
	"log"

	"github.com/BlazeCode1/book-grpc/app/book/model/Book"
	"github.com/BlazeCode1/book-grpc/app/book/service"
	"github.com/segmentio/kafka-go"
)

func StartConsumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"broker:29092"},
		Topic:   "book-events",
		GroupID: "book-group",
	})
	defer reader.Close()

	log.Println("Kafka consumer started")

	// Test connection to broker
	conn, err := kafka.DialLeader(context.Background(), "tcp", "broker:29092", "book-events", 0)
	if err != nil {
		log.Printf("Failed to connect to Kafka broker: %v", err)
		return
	}
	conn.Close()
	log.Println("Successfully connected to Kafka broker")

	for {
		ctx := context.Background()
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
			log.Printf("Raw message value: %s", string(msg.Value))
			continue
		}

		log.Printf("Processing book update - ID: %s, BookName: %s", id, book.BookName)

		book.ID = id

		// Update the book in the database
		response, err := service.HandleUpdateBook(id, book)
		if err != nil {
			log.Printf("Error updating book: %v", err)
			continue
		}

		log.Printf("Response: %v", response)
	}
}
