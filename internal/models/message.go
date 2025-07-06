package models

type WSMessage struct {
	AuthorID   string `json:"author_id"`
	AuthorName string `json:"author_name"`
	Message    string `json:"message"`
}
