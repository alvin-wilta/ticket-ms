package main

type Ticket struct {
	ID          uint   `json:"id" gorm:"primaryKey,autoIncrement"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Name        string `json:"name"`
}
