package rpc

import (
	"log"
	"net"

	"hex/internal/adapters/framework/left/grpc/pb"
	"hex/internal/ports"

	"google.golang.org/grpc"
)

// Adapter implements the GRPCPort interface
type Adapter struct {
	api ports.APIPort
}

// NewAdapter creates a new Adapter
func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

// Run registers the ArithmeticServiceServer to a grpcServer and serves on
// the specified port
func (grpca Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}
}
