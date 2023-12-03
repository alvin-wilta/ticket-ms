package nsq

import (
	"encoding/json"
	"log"

	"github.com/alvin-wilta/ticket-ms/ticket_service/config"
	"github.com/alvin-wilta/ticket-ms/ticket_service/db"
	"github.com/nsqio/go-nsq"
	"gorm.io/gorm"
)

type Handler struct {
	dbw *gorm.DB
}

func New(dbw *gorm.DB, cfg *config.Config) nsq.Handler {
	return &Handler{
		dbw: dbw,
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

	return nil
}
