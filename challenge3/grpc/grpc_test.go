package grpc

import (
	"os"
	"testing"
)

func TestStartGRPCServer(t *testing.T) {
	os.Setenv("GRPC_PORT", "50051")

	grpcServer, lis, err := StartGRPCServer()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if grpcServer == nil {
		t.Fatalf("Expected grpcServer to be non-nil")
	}

	if lis == nil {
		t.Fatalf("Expected lis to be non-nil")
	}

	// clean up
	grpcServer.GracefulStop()
	lis.Close()
}
