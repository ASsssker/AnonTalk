package room

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"

	"github.com/ASsssker/AnonTalk/internal/models"
	"github.com/google/uuid"
)

var (
	ErrClientNotFound = errors.New("room client not found")
	ErrMsgWrite       = errors.New("room client write message error")
)

//go:generate mockgen -package mock -destination $MOCK_FOLDER/room/client.go . RoomClient
type RoomClient interface {
	GetID() string
	Write(authorID string, msg models.WSMessage) error
	Close(ctx context.Context) error
	MsgSubscribe(ctx context.Context, msgChan chan<- models.WSMessage) error
}

type Room struct {
	log     *slog.Logger
	ID      string
	Name    string
	clients map[string]RoomClient
	msgChan chan models.WSMessage
	mu      *sync.RWMutex
	ctx     context.Context
	cancel  context.CancelFunc
}

func NewRoom(log *slog.Logger, name string) *Room {
	id := uuid.NewString()
	log = log.With(slog.String("roomID", id))
	ctx, cancel := context.WithCancel(context.Background())

	return &Room{
		log:     log,
		ID:      id,
		Name:    name,
		msgChan: make(chan models.WSMessage),
		clients: make(map[string]RoomClient),
		mu:      &sync.RWMutex{},
		ctx:     ctx,
		cancel:  cancel,
	}
}

func (r *Room) Run() {
	for {
		select {
		case msg := <-r.msgChan:
			if err := r.Broadcast(msg.AuthorID, msg); err != nil {
				r.log.Error("failed to broadcast message:" + err.Error())
			}
		case <-r.ctx.Done():
			return
		}
	}
}

func (r *Room) AddClient(client RoomClient) error {
	r.mu.Lock()

	r.clients[client.GetID()] = client
	r.log.Debug("add client", slog.String("clientID", client.GetID()))

	r.mu.Unlock()

	defer r.DeleteClients(client.GetID())
	if err := client.MsgSubscribe(r.ctx, r.msgChan); err != nil {
		return fmt.Errorf("failed to msg subscribe: %w", err)
	}

	return nil
}

func (r *Room) DeleteClients(clientID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.clients[clientID]; !exists {
		r.log.Debug("delete client not found", slog.String("clientID", clientID))
		return fmt.Errorf("failed to delete client: %w", ErrClientNotFound)
	}

	r.clients[clientID].Close(context.TODO())
	delete(r.clients, clientID)

	return nil
}

func (r *Room) Broadcast(withoutClientID string, msg models.WSMessage) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var err error

	for clientID, client := range r.clients {
		if clientID == withoutClientID {
			continue
		}

		if err := client.Write(withoutClientID, msg); err != nil {
			err = errors.Join(err, fmt.Errorf("failed to write msg for client %s: %w", clientID, ErrMsgWrite))
			r.log.Error("failed to write message for client", slog.String("clientID", clientID))
		}
	}

	return err
}

func (r *Room) ClientsCount() int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(r.clients)
}

func (r *Room) Close(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	defer r.ctx.Done()

	var err error
	for _, client := range r.clients {
		if err := client.Close(ctx); err != nil {
			err = errors.Join(err, fmt.Errorf("failed to close room client"))
		}
	}

	return err
}
