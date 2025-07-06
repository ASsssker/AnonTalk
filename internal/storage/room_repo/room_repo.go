package roomrepo

import (
	"context"
	"log/slog"
	"sync"

	"github.com/ASsssker/AnonTalk/internal/models"
	"github.com/ASsssker/AnonTalk/internal/room"
	"github.com/ASsssker/AnonTalk/internal/storage"
)

type RoomRepo struct {
	log   *slog.Logger
	rooms map[string]*room.Room
	mu    *sync.RWMutex
}

func NewRoomRepo(log *slog.Logger) *RoomRepo {
	return &RoomRepo{
		log:   log,
		rooms: make(map[string]*room.Room),
		mu:    &sync.RWMutex{},
	}
}

func (r *RoomRepo) GetRoom(_ context.Context, id string) (*room.Room, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	room, exists := r.rooms[id]
	if !exists {
		return nil, storage.ErrRoomNotFound
	}

	return room, nil
}

func (r *RoomRepo) GetRoomInfo(_ context.Context, id string) (*models.Room, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	room, exists := r.rooms[id]
	if !exists {
		return nil, storage.ErrRoomNotFound
	}

	return &models.Room{UUID: room.ID, Name: room.Name}, nil
}

func (r *RoomRepo) NewRoom(_ context.Context, name string) (*models.Room, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	newRoom := room.NewRoom(r.log, name)
	r.rooms[newRoom.ID] = newRoom

	return &models.Room{UUID: newRoom.ID, Name: newRoom.Name}, nil
}

func (r *RoomRepo) DeleteRoom(_ context.Context, id string) (string, error) {
	return  "", nil
}
