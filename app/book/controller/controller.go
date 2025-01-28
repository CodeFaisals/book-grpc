package controller

import (
	"context"
	grpc2 "github.com/BlazeCode1/book-grpc/app/book/controller/grpc"
	"github.com/BlazeCode1/book-grpc/app/book/repository"
	"github.com/BlazeCode1/book-grpc/app/book/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	grpc2.UnimplementedBookServiceServer
}

func (s *server) AddBook(ctx context.Context, req *grpc2.BookRequest) (*grpc2.BookResponse, error) {
	bookInstance := repository.Book{
		BookName: req.BookName,
	}
	return service.HandleAddBook(bookInstance)
}

func (s *server) GetBooks(ctx context.Context, req *grpc2.EmptyRequest) (*grpc2.BookListResponse, error) {
	return service.HandleGetBooks()
}

func (s *server) DeleteBook(ctx context.Context, req *grpc2.BookDeletionRequest) (*grpc2.BookResponse, error) {
	return service.HandleDeleteBook(req.Id)
}

// StartGRPCServer initializes and starts the gRPC client
func StartGRPCServer(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
	}

	grpcServer := grpc.NewServer()
	grpc2.RegisterBookServiceServer(grpcServer, &server{})

	log.Printf("gRPC client started on %s", address)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC client: %v", err)
	}
}
