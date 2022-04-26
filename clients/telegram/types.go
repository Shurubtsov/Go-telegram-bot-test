package telegram

type UpdatesResponse struct {
	Status bool     `json: "ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID      int    `json:"update_id"`
	Message string `json:"message"`
}
