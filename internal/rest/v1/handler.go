package v1

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ASsssker/AnonTalk/internal/models"
	bp "github.com/ASsssker/AnonTalk/internal/rest/v1/boilerplate"
	"github.com/labstack/echo/v4"
)

type RoomService interface {
	CreateNewRoom(ctx context.Context, roomName string) (*models.Room, error)
	GetRoom(ctx context.Context, id string) (*models.Room, error)
}

type Handler struct {
	log         *slog.Logger
	roomService RoomService
}

func NewHandler(log *slog.Logger, roomService RoomService) *Handler {
	return &Handler{
		log:         log,
		roomService: roomService,
	}
}

func RegisterHandler(server *echo.Echo, handler bp.ServerInterface) {
	bp.RegisterHandlers(server, handler)
}

var _ bp.ServerInterface = Handler{}

func (h Handler) Healthcheck(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h Handler) GetApi(c echo.Context) error {
	swagger, err := bp.GetSwagger()
	if err != nil {
		return fmt.Errorf("failed to get swagger: %w", err)
	}

	spec, err := swagger.MarshalJSON()
	if err != nil {
		return fmt.Errorf("failed to marshall swagger spec: %w", err)
	}

	specStr := string(spec)
	response := bp.API{Api: &specStr}

	return c.JSON(http.StatusOK, response)
}

func (h Handler) CreateNewRoom(c echo.Context) error {
	ctx := c.Request().Context()

	var createRoomData bp.CreateNewRoomJSONRequestBody
	if err := c.Bind(&createRoomData); err != nil {
		return fmt.Errorf("failed to parse create new room data: %w", err)
	}

	room, err := h.roomService.CreateNewRoom(ctx, createRoomData.Name)
	if err != nil {
		return fmt.Errorf("failed to create new room: %w", err)
	}

	response := bp.RoomInfo{Id: room.UUID, Name: room.Name}

	return c.JSON(http.StatusCreated, response)
}

func (h Handler) GetRoomInfo(c echo.Context, id string) error {
	ctx := c.Request().Context()

	room, err := h.roomService.GetRoom(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get room: %w", err)
	}

	response := bp.RoomInfo{Id: room.UUID, Name: room.Name}

	return c.JSON(http.StatusOK, response)
}

func (h Handler) ConnectRoom(ctx echo.Context, id string, params bp.ConnectRoomParams) error {

	// TODO: Websocket connect

	return nil
}
