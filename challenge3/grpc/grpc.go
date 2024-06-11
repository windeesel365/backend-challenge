package grpc

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGRPCServer() (*grpc.Server, net.Listener, error) {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "9090" //default gRPC port
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer) //register gRPC services

	return grpcServer, lis, nil
}
