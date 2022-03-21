package main

import (
	"context"
	"vending/proto/example"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	grpcServer := grpc.NewServer()

	example.RunServer(ctx, grpcServer, "50011", nil)
}
