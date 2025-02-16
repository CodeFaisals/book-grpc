package controller

import (
	"context"
	"log"
	"net"

	grpc2 "github.com/BlazeCode1/book-grpc/app/book/controller/grpc"
	b "github.com/BlazeCode1/book-grpc/app/book/model/Book"
	"github.com/BlazeCode1/book-grpc/app/book/service"

	"google.golang.org/grpc"
)

type server struct {
	grpc2.UnimplementedBookServiceServer
	service service.BookService
}

func (s *server) AddBook(ctx context.Context, req *grpc2.BookRequest) (*grpc2.BookResponse, error) {
	bookInstance := b.Book{
		BookName: req.BookName,
		Author:   req.Author,
	}
	return s.service.HandleAddBook(bookInstance)
}

func (s *server) GetBooks(ctx context.Context, req *grpc2.EmptyRequest) (*grpc2.BookListResponse, error) {
	return s.service.HandleGetBooks()
}

func (s *server) DeleteBook(ctx context.Context, req *grpc2.BookDeletionRequest) (*grpc2.BookResponse, error) {
	return s.service.HandleDeleteBook(req.Id)
}

// StartGRPCServer initializes and starts the gRPC client
func StartGRPCServer(address string, bookService service.BookService) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
	}
	bookServer := &server{service: bookService}
	grpcServer := grpc.NewServer()
	grpc2.RegisterBookServiceServer(grpcServer, bookServer)

	log.Printf("gRPC client started on %s", address)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC client: %v", err)
	}
}
