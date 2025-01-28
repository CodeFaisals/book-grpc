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
	// Initialize Couchbase connection
	couchbase.InitCouchbase("admin", "1q2w3e4r5t", "books_bucket")

	// Start Kafka consumer in a separate goroutine
	go consumer.StartConsumer()

	// Start gRPC client
	go controller.StartGRPCServer(":50051")

	// Graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown

	log.Println("Shutting down the client...")
}

//
