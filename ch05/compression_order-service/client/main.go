package main

import (
	pb "compression_order-service/order-service-gen"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:9090"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to Dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewOrderManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
	defer cancel()

	// RPC: Add Order
	order1 := pb.Order{Id: "120", Items: []string{"iPhone XS", "Mac Book Pro"}, Destination: "HPU:Lxc", Price: -123.00}
	res, _ := client.AddOrder(ctx, &order1)

	log.Printf("AddOrder Response: %v", res.Value)

}
