package room

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"

	"github.com/google/uuid"
)

var (
	ErrClientNotFound = errors.New("room client not found")
	ErrMsgWriteErr    = errors.New("room client write message error")
)

//go:generate mockgen -package mock -destination $MOCK_FOLDER/room/client.go . RoomClient
type RoomClient interface {
	GetID() string
	Write(msg string) error
	Close(ctx context.Context) error
}

type Room struct {
	log     *slog.Logger
	ID      string
	Name    string
	clients map[string]RoomClient
	mu      *sync.RWMutex
}

func NewRoom(log *slog.Logger, name string, clients ...RoomClient) *Room {
	roomClients := make(map[string]RoomClient)
	if len(clients) > 0 {
		for _, client := range clients {
			roomClients[client.GetID()] = client
		}
	}

	id := uuid.NewString()
	log = log.With(slog.String("roomID", id))

	return &Room{
		log:     log,
		ID:      id,
		Name:    name,
		clients: roomClients,
		mu:      &sync.RWMutex{},
	}
}

func (r *Room) AddClients(clients ...RoomClient) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, client := range clients {
		r.clients[client.GetID()] = client
		r.log.Debug("add client", slog.String("clientID", client.GetID()))
	}
}

func (r *Room) DeleteClients(clientsID ...string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, clientID := range clientsID {
		if _, exists := r.clients[clientID]; !exists {
			r.log.Debug("delete client not found", slog.String("clientID", clientID))
			continue
		}
		delete(r.clients, clientID)
	}
}

func (r *Room) Broadcast(withoutClientID string, msg string) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var err error

	for clientID, client := range r.clients {
		if clientID == withoutClientID {
			continue
		}

		if err := client.Write(msg); err != nil {
			err = errors.Join(err, fmt.Errorf("failed to write msg for client %s: %w", clientID, ErrMsgWriteErr))
			r.log.Error("failed to write message for client", slog.String("clientID", clientID))
		}
	}

	return err
}

func (r *Room) Close(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	var err error
	for _, client := range r.clients {
		if err := client.Close(ctx); err != nil {
			err = errors.Join(err, fmt.Errorf("failed to close room client"))
		}
	}

	return err
}
