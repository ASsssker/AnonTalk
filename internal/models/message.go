package models

type WSMessage struct {
	AuthorID string `json:"author_id"`
	Message  string `json:"message"`
}
