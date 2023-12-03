package main

import (
	"context"
	"log"

	pb "github.com/alvin-wilta/ticket-ms/proto"
	"github.com/alvin-wilta/ticket-ms/ticket_service/config"
	"github.com/alvin-wilta/ticket-ms/ticket_service/db"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedTicketServiceServer
	db  *gorm.DB
	rdb *redis.Client
}

func InitGRPC(db *gorm.DB, cfg *config.Config) *grpc.Server {
	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterTicketServiceServer(grpcServer, &Server{db: db})
	log.Print("[GRPC] GRPC Server running")
	return grpcServer
}

func (s *Server) HealthCheck(ctx context.Context, req *pb.Empty) (*pb.HealthCheckResponse, error) {
	log.Print("[RPC] Health check succeeded")
	return &pb.HealthCheckResponse{
		Success: true,
	}, nil
}

func (s *Server) GetTicketList(ctx context.Context, req *pb.GetTicketListRequest) (*pb.GetTicketListResponse, error) {
	log.Print("[RPC] GetTicketList")
	var err error
	var tickets []db.Ticket
	var res pb.GetTicketListResponse
	query := s.db

	if req.Id != 0 {
		query = query.Where("id = ?", req.Id)
	}
	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}

	result := query.Find(&tickets)
	if result.Error != nil {
		err = result.Error
		log.Panic(err)
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

	return &res, err
}

func (s *Server) CreateTicket(ctx context.Context, req *pb.CreateTicketRequest) (*pb.CreateTicketResponse, error) {
	log.Print("[RPC] CreateTicket")
	var err error
	res := pb.CreateTicketResponse{}

	ticket := db.Ticket{
		Title:       req.Title,
		Description: req.Description,
		Status:      "New",
		Name:        req.Name,
	}
	result := s.db.Create(&ticket)
	if result.Error != nil {
		err = result.Error
		log.Panic(err)
	}

	res.Id = int32(ticket.ID)
	res.Success = true

	return &res, err
}

func (s *Server) UpdateTicket(ctx context.Context, req *pb.UpdateTicketRequest) (*pb.UpdateTicketResponse, error) {
	log.Printf("[RPC] UpdateTicket with id %v", req.Id)
	var err error
	res := pb.UpdateTicketResponse{
		Id:      req.Id,
		Success: true,
	}

	result := s.db.Model(&db.Ticket{}).Where("id = ?", req.Id).Update("status", req.Status)
	if result.Error != nil {
		err = result.Error
		log.Panic(err)
	}
	if result.RowsAffected == 0 {
		res.Success = false
	}
	return &res, err
}

func (s *Server) DeleteTicket(ctx context.Context, req *pb.DeleteTicketRequest) (*pb.DeleteTicketResponse, error) {
	log.Printf("[RPC] DeleteTicket with id %v", req.Id)
	var err error
	res := pb.DeleteTicketResponse{}
	res.Id = req.Id
	res.Success = true

	result := s.db.Delete(&db.Ticket{}, uint(req.Id))
	if result.Error != nil {
		err = result.Error
		log.Panic(err)
	}
	if result.RowsAffected == 0 {
		res.Success = false
	}

	return &res, err
}
