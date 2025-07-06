package wschat

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	v1 "github.com/ASsssker/AnonTalk/internal/rest/v1"
	"github.com/ASsssker/AnonTalk/internal/services"
	roomrepo "github.com/ASsssker/AnonTalk/internal/storage/room_repo"
	"github.com/labstack/echo/v4"
)

type App struct {
	log    *slog.Logger
	server *http.Server
}

func NewApp(log *slog.Logger) *App {
	log.Info("start service")
	roomProvider := roomrepo.NewRoomRepo(log)
	roomService := services.NewRoomService(log, roomProvider)

	server := echo.New()
	server.Server.Addr = ":8000"
	handler := v1.NewHandler(log, roomService)

	v1.RegisterHandler(server, handler)

	return &App{
		log:    log,
		server: server.Server,
	}
}

func (a *App) Run() error {
	a.log.Info("service server started")
	if err := a.server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	a.log.Info("service stopping")
	return a.server.Shutdown(ctx)
}
