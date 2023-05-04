package main

import (
	"context"
	"errors"
	"log"
	"net"
	pb "server-reflection/proto"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
	productMap map[string]*pb.Product
}

func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*wrappers.StringValue, error) {
	out, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	in.Id = out.String()

	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	return &wrappers.StringValue{Value: in.Id}, nil
}

func (s *server) GetProduct(ctx context.Context, in *wrappers.StringValue) (*pb.Product, error) {
	value, ok := s.productMap[in.Value]
	if ok {
		return value, nil
	}
	return nil, errors.New("Product does not exist for the ID: " + in.Value)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})
	// Register reflection service on gRPC server
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v", err)
	}
}
