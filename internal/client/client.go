package client

import (
	"context"
	"fmt"

	"github.com/ASsssker/AnonTalk/internal/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WSClient struct {
	conn     *websocket.Conn
	UUID     string
	Username string
}

func NewWSClient(conn *websocket.Conn, username string) *WSClient {
	id := uuid.NewString()
	return &WSClient{
		conn:     conn,
		UUID:     id,
		Username: username,
	}
}

func (c *WSClient) MsgSubscribe(ctx context.Context, msgChan chan<- models.WSMessage) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			var wsMsg models.WSMessage
			if err := c.conn.ReadJSON(wsMsg); err != nil {
				return fmt.Errorf("failed tp read message: %w", err)
			}
			select {
			case <-ctx.Done():
				return nil
			case msgChan <- wsMsg:
			}
		}
	}
}

func (c *WSClient) Write(authorID string, msg string) error {
	wsMsg := models.WSMessage{
		AuthorID: authorID,
		Message:  msg,
	}

	if err := c.conn.WriteJSON(wsMsg); err != nil {
		return err
	}

	return nil
}

func (c *WSClient) GetID() string {
	return c.UUID
}

func (c *WSClient) Close(ctx context.Context) error {
	return c.conn.Close()
}
