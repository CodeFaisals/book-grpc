package controller

import (
	"context"

	pb "github.com/BlazeCode1/book-grpc/app/controller/grpc"
	"github.com/BlazeCode1/book-grpc/app/repository"
	"github.com/BlazeCode1/book-grpc/app/service/book"

	//"github.com/BlazeCode1/book-grpc/app/service/book" //service/handler
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBookServiceServer
}

// Implement the required methods using your existing handlers
func (s *server) AddBook(ctx context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {
	bookInstance := repository.Book{
		BookName: req.BookName,
	}
	return book.HandleAddBook(bookInstance)
}

func (s *server) GetBooks(ctx context.Context, req *pb.EmptyRequest) (*pb.BookListResponse, error) {
	return book.HandleGetBooks()
}

func (s *server) DeleteBook(ctx context.Context, req *pb.BookDeletionRequest) (*pb.BookResponse, error) {
	return book.HandleDeleteBook(req.Id)
}

// StartGRPCServer initializes and starts the gRPC server
func StartGRPCServer(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBookServiceServer(grpcServer, &server{})

	log.Printf("gRPC server started on %s", address)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
