package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"log"

	"github.com/alvin-wilta/ticket-ms/proto"
	"github.com/alvin-wilta/ticket-ms/proxy_service/graph/model"
)

// CreateTicket is the resolver for the createTicket field.
func (r *mutationResolver) CreateTicket(ctx context.Context, input model.NewTicket) (*model.CreateTicketResponse, error) {
	log.Print("[GQL] Request CreateTicket")
	err := r.handlerNSQ.PublishCreateTicket(input)
	res := &model.CreateTicketResponse{
		Success: true,
	}
	if err != nil {
		res.Success = false
		log.Panicf("[GQL] CreateTicket err: %v", err)
		return res, err
	}
	return res, nil
}

// UpdateTicket is the resolver for the updateTicket field.
func (r *mutationResolver) UpdateTicket(ctx context.Context, input model.UpdateTicket) (*model.UpdateTicketResponse, error) {
	log.Print("[GQL] Request UpdateTicket")
	req := &proto.UpdateTicketRequest{
		Id:     int32(input.ID),
		Status: input.Status,
	}
	res := &model.UpdateTicketResponse{
		ID:      input.ID,
		Success: true,
	}
	grpcRes, err := r.Resolver.grpcClient.UpdateTicket(ctx, req)
	if err != nil {
		log.Panicf("[GQL] UpdateTicket err: %v", err)
	}
	if !grpcRes.Success {
		res.Success = false
	}
	return res, nil
}

// DeleteTicket is the resolver for the deleteTicket field.
func (r *mutationResolver) DeleteTicket(ctx context.Context, input model.DeleteTicket) (*model.DeleteTicketResponse, error) {
	log.Print("[GQL] Request DeleteTicket")
	req := &proto.DeleteTicketRequest{
		Id: int32(input.ID),
	}
	res := &model.DeleteTicketResponse{
		ID:      input.ID,
		Success: true,
	}
	grpcRes, err := r.Resolver.grpcClient.DeleteTicket(ctx, req)
	if err != nil {
		log.Panicf("[GQL] DeleteTicket err: %v", err)
	}
	if !grpcRes.Success {
		res.Success = false
	}
	return res, nil
}

// HealthCheck is the resolver for the healthCheck field.
func (r *queryResolver) HealthCheck(ctx context.Context) (string, error) {
	log.Print("[GQL] Request Healthcheck")
	res, err := r.Resolver.grpcClient.HealthCheck(ctx, &proto.Empty{})
	if err != nil {
		log.Panicf("[GQL] HealthCheck err: %v", err)
	}
	if !res.Success {
		return "failed", nil
	}
	return "success", nil
}

// Tickets is the resolver for the tickets field.
func (r *queryResolver) Tickets(ctx context.Context, input *model.TicketFilter) ([]*model.Ticket, error) {
	log.Print("[GQL] Request Tickets")
	var ticketList []*model.Ticket
	req := &proto.GetTicketListRequest{}
	if input != nil {
		if input.ID != nil {
			req.Id = int32(*input.ID)
		}
		if input.Status != nil {
			req.Status = *input.Status
		}
	}
	res, err := r.Resolver.grpcClient.GetTicketList(ctx, req)
	if err != nil {
		log.Panicf("[GQL] Tickets err: %v", err)
	}

	for _, ticket := range res.Tickets {
		// Map response
		newId := int(ticket.Id)
		newDescription := string(ticket.Description)
		newStatus := string(ticket.Status)

		ticketList = append(ticketList, &model.Ticket{
			ID:          &newId,
			Title:       ticket.Title,
			Description: &newDescription,
			Status:      &newStatus,
			Name:        ticket.Name,
		})
	}

	return ticketList, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
