package nsqw

type TicketMessage struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type UpdateMessage struct {
	Id     int32  `json:"id"`
	Status string `json:"status"`
}

type DeleteMessage struct {
	Id int32 `json:"id"`
}
