package rpc

import (
	"log"
	"net"

	"hex-arch-go-grpc/internal/adapters/framework/in/grpc/pb"
	"hex-arch-go-grpc/internal/ports"

	"google.golang.org/grpc"
)

// Adapter is an rpc adapter the is compatible with
// the
type Adapter struct {
	api ports.APIPort
}

// NewAdapter creates a grpc entry adapter that is
// compatible with gRPCEntryPort
func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

// Run runs the grpc server
func (grpca Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	arithmeticServiceServer := &Adapter{}
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
