package main

import (
	"log"

	pb "github.com/alvin-wilta/ticket-ms/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func initGRPC() *pb.TicketServiceClient {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	client := pb.NewTicketServiceClient(conn)
	return &client

}
