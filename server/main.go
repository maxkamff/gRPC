package main

import (
	"context"
	"log"
	"net"

	pb "grpc-todo/proto"
	"grpc-todo/server/postgres"

	"google.golang.org/grpc"
)

type StoreServer struct {
	pb.UnimplementedStoreServiceServer
}

func (s *StoreServer) CreateStore(ctx context.Context, in *pb.Store) (*pb.Store, error) {
	log.Printf("Recieved: %v", in.GetName())
	
	store, err := postgres.CreateStore(&pb.Store{
		Id: in.Id,
		Name: in.Name,
		Description: in.Description,
		Address: in.Address,
		IsOpen: in.IsOpen,
	})
	if err != nil{
		return nil, err   
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
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
