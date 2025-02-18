package main

import (
	kafka "github.com/BlazeCode1/book-grpc/app/book/consumer"
	"github.com/BlazeCode1/book-grpc/app/book/repository"
	"golang.org/x/net/context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BlazeCode1/book-grpc/app/book/controller"
	"github.com/BlazeCode1/book-grpc/app/book/service"
	"github.com/BlazeCode1/book-grpc/couchbase"
)

func main() {
	// Initialize Couchbase connection
	cb, err := couchbase.InitCouchbase("Administrator", "123123", "books_bucket")
	if err != nil {
		return
	}

	// Start Kafka consumer in a separate goroutine

	// Start gRPC client
	bookRepo := repository.NewBookRepository(cb)
	bookService := service.NewBookService(bookRepo)
	consumer := kafka.NewKafkaConsumer(bookService)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		if err := consumer.StartConsumer(ctx); err != nil {
			log.Fatalf("Kafka Consumer failed: %v", err)
		}
	}()
	go controller.StartGRPCServer(":50051", bookService)

	// Graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown

	log.Println("Shutting down the client...")
	cancel()
}

//
