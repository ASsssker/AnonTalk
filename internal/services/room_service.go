package services

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/ASsssker/AnonTalk/internal/client"
	"github.com/ASsssker/AnonTalk/internal/models"
	"github.com/ASsssker/AnonTalk/internal/room"
	"github.com/ASsssker/AnonTalk/internal/storage"
	"github.com/gorilla/websocket"
)

var (
	ErrRoomNotFound = errors.New("room not found")
	ErrUnexpected   = errors.New("unexpected error")
)

//go:generate mockgen -package mock -destination $MOCK_FOLDER/services/room_provider.go . RoomProvider
type RoomProvider interface {
	GetRoom(ctx context.Context, id string) (*room.Room, error)
	GetRoomInfo(ctx context.Context, id string) (*models.Room, error)
	NewRoom(ctx context.Context, name string) (*models.Room, error)
	DeleteRoom(ctx context.Context, id string) (string, error)
}

type RoomService struct {
	log          *slog.Logger
	roomProvider RoomProvider
}

func NewRoomService(log *slog.Logger, roomProvider RoomProvider) *RoomService {
	return &RoomService{
		log:          log.With(slog.String("service", "room_service")),
		roomProvider: roomProvider,
	}
}

func (r *RoomService) AddUserToRoom(ctx context.Context, roomID string, clientName string, clientConn *websocket.Conn) error {
	log := r.log.With(slog.String("op", "add_user_to_room"))
	log.Debug("start operation")

	room, err := r.roomProvider.GetRoom(ctx, roomID)
	if err != nil {
		if errors.Is(err, storage.ErrRoomNotFound) {
			return fmt.Errorf("failed to get room: %w: %w", ErrRoomNotFound, err)
		}

		return fmt.Errorf("failed to get room: %w: %w", ErrUnexpected, err)
	}

	client := client.NewWSClient(clientConn, clientName)

	if err := room.AddClient(client); err != nil {
		return fmt.Errorf("failed to add client: %w", err)
	}

	log.Debug("operation completed")

	return nil
}

func (r *RoomService) CreateNewRoom(ctx context.Context, roomName string) (*models.Room, error) {
	log := r.log.With(slog.String("op", "create_new_room"))
	log.Debug("start operation")

	room, err := r.roomProvider.NewRoom(ctx, roomName)
	if err != nil {
		return nil, fmt.Errorf("failed to create new room: %w: %w", ErrUnexpected, err)
	}

	log.Debug("operation completed")

	return room, nil
}
func (r *RoomService) GetRoom(ctx context.Context, id string) (*models.Room, error) {
	log := r.log.With(slog.String("op", "get_room"))
	log.Debug("start operation")

	roomInfo, err := r.roomProvider.GetRoomInfo(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrRoomNotFound) {
			return nil, fmt.Errorf("failed to get room: %w: %w", ErrRoomNotFound, err)
		}

		return nil, fmt.Errorf("failed to get room info: %w: %w", ErrUnexpected, err)
	}

	log.Debug("operation completed")

	return roomInfo, nil
}
