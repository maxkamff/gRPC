package main

import (
	"context"
	"log"
	"net"

	pb "grpc-todo/proto"

	"google.golang.org/grpc"
)

type StoreServer struct {
	pb.UnimplementedStoreServiceServer
}

func (s *StoreServer) CreateStore(ctx context.Context, in *pb.Store) (*pb.Store, error) {
	log.Printf("Recieved: %v", in.GetName())
	store := &pb.Store{
		Name:        in.GetName(),
		Description: in.GetDescription(),
		IsOpen:      false,
		Id:          1,
	}
	return store, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed connection: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterStoreServiceServer(s, &StoreServer{})

	log.Printf("server listening at %v", lis.Addr())
	if s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
