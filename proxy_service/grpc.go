package main

import (
	"fmt"
	"log"

	pb "github.com/alvin-wilta/ticket-ms/proto"
	"github.com/alvin-wilta/ticket-ms/proxy_service/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func initGRPC(cfg *config.Config) *pb.TicketServiceClient {
	addr := fmt.Sprintf("%s:%s", cfg.GrpcAddr, cfg.GrpcPort)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	client := pb.NewTicketServiceClient(conn)
	return &client

}
