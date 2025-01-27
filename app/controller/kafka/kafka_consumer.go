//package controller
//
//import (
//	"faisal.com/bookProject/repository"
//	"github.com/Trendyol/kafka-cronsumer/pkg/kafka"
//	"github.com/Trendyol/kafka-konsumer/v2"
//	"log"
//)
//
//func StartConsumer() {
//	cfg := &kafka.ConsumerConfig{
//		Reader: kafka.ReaderConfig{
//			Brokers: []string{"localhost:9092"},
//			Topic:   "book-events",
//			GroupID: "book-group",
//		},
//		ConsumeFn: consumeFn,
//	}
//
//	consumer, err := kafka.NewConsumer(cfg)
//	if err != nil {
//		log.Fatalf("Failed to start Kafka consumer: %v", err)
//	}
//
//	defer consumer.Stop()
//	consumer.Consume()
//	log.Println("Kafka consumer started.")
//}
//
//func consumeFn(message *kafka.Message) error {
//	log.Printf("Received message: %s", message.Value)
//
//	bookID := string(message.Key)
//	updatedBookName := string(message.Value)
//
//	err := repository.UpdateBook(bookID, updatedBookName)
//	if err != nil {
//		log.Printf("Failed to update book: %v", err)
//		return err
//	}
//
//	log.Printf("Book with ID '%s' updated successfully", bookID)
//	return nil
//}
