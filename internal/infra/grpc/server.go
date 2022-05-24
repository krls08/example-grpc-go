package rpc

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func Run() {
	fmt.Println("grpc server running...")
	lis, err := net.Listen("tpc", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}
	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server gRPC server over port 9000: %v", err)
	}

}
