package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BlazeCode1/book-grpc/app/book/consumer"
	"github.com/BlazeCode1/book-grpc/app/book/controller"
	"github.com/BlazeCode1/book-grpc/couchbase"
)

func main() {
	var Client couchbase.Client
	var Consumer consumer.KafkaConsumer
	// Initialize Couchbase connection
	err := Client.InitCouchbase("Administrator", "123123", "books_bucket")
	if err != nil {
		return
	}

	// Start Kafka consumer in a separate goroutine

	go func() {
		err := Consumer.StartConsumer()
		if err != nil {

		}
	}()
	// Start gRPC client
	go controller.StartGRPCServer(":50051")

	// Graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown

	log.Println("Shutting down the client...")
}

//
