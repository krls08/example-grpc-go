package main

import (
	"context"
	"fmt"
	"log"

	"github.com/krls08/example-grpc-go/internal/infra/grpc/proto/chatpb_contract"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("run client main")

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	c := chatpb_contract.NewChatServiceClient(conn)

	message := chatpb_contract.Message{
		Body: "hello from the client!",
	}
	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response From server: %s", response.Body)

}
