package main

import (
	"context"
	"log"

	pb "github.com/alvin-wilta/ticket-ms/proto"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedTicketServiceServer
	db *gorm.DB
}

func InitGRPC(db *gorm.DB) *grpc.Server {
	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterTicketServiceServer(grpcServer, &Server{db: db})
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
	var res pb.GetTicketListResponse

	result := s.db.Find(&tickets)
	if result.Error != nil {
		log.Panic(result.Error)
	}

	for _, ticket := range tickets {
		res.Tickets = append(res.Tickets, &pb.Ticket{
			Id:          int32(ticket.ID),
			Title:       ticket.Title,
			Description: ticket.Description,
			Status:      ticket.Status,
			Name:        ticket.Name,
		})
	}

	return &res, nil
}

func (s *Server) CreateTicket(ctx context.Context, req *pb.CreateTicketRequest) (*pb.CreateTicketResponse, error) {
	log.Print("[RPC] GetTicketList")
	res := pb.CreateTicketResponse{}

	ticket := Ticket{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Name:        req.Name,
	}
	result := s.db.Create(&ticket)
	if result.Error != nil {
		log.Panic(result.Error)
	}

	res.Id = int32(ticket.ID)
	res.Success = true

	return &res, nil
}

func (s *Server) DeleteTicket(ctx context.Context, req *pb.DeleteTicketRequest) (*pb.DeleteTicketResponse, error) {
	log.Printf("[RPC] DeleteTicket with id %v", req.Id)
	res := pb.DeleteTicketResponse{}
	res.Id = req.Id
	res.Success = true

	result := s.db.Delete(&Ticket{
		ID: uint(req.Id),
	})
	if result.Error != nil {
		log.Panic(result.Error)
	}
	if result.RowsAffected == 0 {
		res.Success = false
	}

	return &res, nil
}
