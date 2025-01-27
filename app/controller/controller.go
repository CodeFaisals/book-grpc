package controller

import (
	pb "github.com/BlazeCode1/book-grpc/app/controller/grpc"
	//"github.com/BlazeCode1/book-grpc/app/service/book" //service/handler
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedBookServiceServer
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
