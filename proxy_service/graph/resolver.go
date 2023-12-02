package graph

import (
	pb "github.com/alvin-wilta/ticket-ms/proto"
	"github.com/alvin-wilta/ticket-ms/proxy_service/nsqw"
)

type Resolver struct {
	grpcClient pb.TicketServiceClient
	handlerNSQ nsqw.HandlerNSQ
}

func InitResolver(grpcClient pb.TicketServiceClient, handlerNSQ nsqw.HandlerNSQ) *Resolver {
	return &Resolver{
		grpcClient: grpcClient,
		handlerNSQ: handlerNSQ,
	}
}
