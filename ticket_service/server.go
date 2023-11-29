package main

import (
	"context"
	"log"

	pb "github.com/alvin-wilta/ticket-ms/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedTicketServiceServer
}

func InitGRPC() *grpc.Server {
	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterTicketServiceServer(grpcServer, &Server{})
	return grpcServer
}

func (s *Server) HealthCheck(ctx context.Context, req *pb.Empty) (*pb.HealthCheckResponse, error) {
	log.Print("[RPC] Health check succeeded")
	return &pb.HealthCheckResponse{
		Success: true,
	}, nil
}

func (s *Server) GetTicketList(ctx context.Context, req *pb.Empty) (*pb.GetTicketListResponse, error) {
	log.Print("[RPC] GetTicketList")
	var tickets []Ticket
	var ticketListResponse pb.GetTicketListResponse
	db.Find(&tickets)

	for _, ticket := range tickets {
		ticketListResponse.Tickets = append(ticketListResponse.Tickets, &pb.Ticket{
			Id:          int32(ticket.ID),
			Title:       ticket.Title,
			Description: ticket.Description,
			Status:      ticket.Status,
			Name:        ticket.Name,
		})
	}

	return &ticketListResponse, nil
}
