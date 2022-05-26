package rpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/krls08/example-grpc-go/internal/infra/grpc/proto/chatpb_contract"
	"google.golang.org/grpc"
)

type server struct {
	chatpb_contract.UnimplementedChatServiceServer
}

func (s *server) SayHello(ctx context.Context, in *chatpb_contract.Message) (*chatpb_contract.Message, error) {
	log.Println("SayHello server in ->", in)
	return &chatpb_contract.Message{Body: "server response: " + in.Body}, nil
}

func Run() {
	fmt.Println("grpc server running...")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := server{}
	grpcServer := grpc.NewServer()

	chatpb_contract.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server gRPC server over port 9000: %v", err)
	}

}
