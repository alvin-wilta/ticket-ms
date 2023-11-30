// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateTicketResponse struct {
	ID      int  `json:"id"`
	Success bool `json:"success"`
}

type DeleteTicket struct {
	ID int `json:"id"`
}

type DeleteTicketResponse struct {
	ID      int  `json:"id"`
	Success bool `json:"success"`
}

type NewTicket struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type Ticket struct {
	ID          *int    `json:"id,omitempty"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	Status      *string `json:"Status,omitempty"`
	Name        string  `json:"Name"`
}

type UpdateTicket struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type UpdateTicketResponse struct {
	ID      int  `json:"id"`
	Success bool `json:"success"`
}