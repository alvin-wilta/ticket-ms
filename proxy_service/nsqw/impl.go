package nsqw

import (
	"log"

	"github.com/alvin-wilta/ticket-ms/proxy_service/graph/model"
)

func (h *HandlerNSQ) PublishCreateTicket(message model.NewTicket) error {
	log.Print("[NSQ] PublishCreateTicket request")
	createMessage := &TicketMessage{
		Title:       message.Title,
		Description: message.Description,
		Name:        message.Name,
	}
	byteCreateMessage := MarshalMessage(createMessage)
	err := h.producer.Publish("ticket"+"Create", byteCreateMessage)
	if err != nil {
		log.Panicf("[NSQ] Publish create ticket error: %v", err)
		return err
	}
	return nil
}

func (h *HandlerNSQ) PublishUpdateTicket(message model.UpdateTicket) error {
	log.Print("[NSQ] PublishUpdateTicket request")
	updateMessage := &UpdateMessage{
		Id:     int32(message.ID),
		Status: message.Status,
	}
	byteUpdateMessage := MarshalMessage(updateMessage)
	err := h.producer.Publish("ticket"+"Update", byteUpdateMessage)
	if err != nil {
		log.Panicf("[NSQ] Publish update ticket error: %v", err)
		return err
	}

	return nil
}

func (h *HandlerNSQ) PublishDeleteTicket(message model.DeleteTicket) error {
	log.Print("[NSQ] PublishDeleteTicket request")
	deleteMessage := &DeleteMessage{
		Id: int32(message.ID),
	}
	byteDeleteMessage := MarshalMessage(deleteMessage)
	err := h.producer.Publish("ticket"+"Delete", byteDeleteMessage)
	if err != nil {
		log.Panicf("[NSQ] Publish delete ticket error: %v", err)
		return err
	}
	return nil
}
