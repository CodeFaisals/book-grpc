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
		book := Book.Book{}
		err = json.Unmarshal(msg.Value, &book)
		if err != nil {
			log.Printf("Error unmarshalling book: %v", err)
			continue
		}

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
