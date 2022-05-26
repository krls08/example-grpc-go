package main

import (
	"fmt"

	rpc "github.com/krls08/example-grpc-go/internal/infra/grpc"
)

func main() {
	rpc.Run()

	fmt.Println("-------main end-------")
}
