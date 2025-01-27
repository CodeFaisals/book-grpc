package grpc

import (
	"faisal.com/bookProject/controller/grpc/handlers"
	pb "faisal.com/bookProject/server/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

// StartGRPCServer initializes and starts the gRPC server
func StartGRPCServer(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBookServiceServer(grpcServer, &handlers.BookService{})

	log.Printf("gRPC server started on %s", address)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
