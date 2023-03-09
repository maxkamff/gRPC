package main

import (
	"context"
	"fmt"
	pb "grpc-todo/proto"
	"time"

	"google.golang.org/grpc"
)

type Store struct {
	id           int64
	Name         string
	Descrioption string
	Address      []string
	IsOpen       bool
}

const (
	serverAddress = "localhost:8000"
)

func main() {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := pb.NewStoreServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	store, err := c.CreateStore(ctx, &pb.Store{
		Name:        "New gRPC store",
		Description: "gRPC test description",
		IsOpen:      true,
		Address: []string{
			"new address 1",
			"new address 2",
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(store)
}
