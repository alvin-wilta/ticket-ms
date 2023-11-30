package graph

import (
	pb "github.com/alvin-wilta/ticket-ms/proto"
)

type Resolver struct {
	grpcClient pb.TicketServiceClient
}

func InitResolver(grpcClient pb.TicketServiceClient) *Resolver {
	return &Resolver{
		grpcClient: grpcClient,
	}
}
