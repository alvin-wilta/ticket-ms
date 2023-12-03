package nsq

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/alvin-wilta/ticket-ms/ticket_service/config"
	"github.com/alvin-wilta/ticket-ms/ticket_service/db"
	"github.com/nsqio/go-nsq"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Handler struct {
	dbw *gorm.DB
	rdb *redis.Client
}

func New(dbw *gorm.DB, rdb *redis.Client, cfg *config.Config) nsq.Handler {
	return &Handler{
		dbw: dbw,
		rdb: rdb,
	}
}

func NewHandler(handler func(m *nsq.Message) error) nsq.HandlerFunc {
	return func(message *nsq.Message) error {
		return handler(message)
	}
}

func (h *Handler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}
	log.Print("[NSQ] CreateTicket")
	ticket := &db.Ticket{}
	err := json.Unmarshal(m.Body, ticket)
	if err != nil {
		log.Panicf("[NSQ] Unmarshal 'create' message failed: %v", err)
	}

	result := h.dbw.Create(&ticket)
	err = result.Error
	if err != nil {
		log.Panic(result.Error)
		return err
	}

	// Cache result

	ticketId := string(int32(ticket.ID))
	ctx := context.Background()
	h.rdb.Set(ctx, ticketId, m.Body, 5*time.Minute)

	return nil
}
