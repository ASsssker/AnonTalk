package v1

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ASsssker/AnonTalk/internal/models"
	bp "github.com/ASsssker/AnonTalk/internal/rest/v1/boilerplate"
	mock "github.com/ASsssker/AnonTalk/mock/rest/v1/handler"
	"github.com/ASsssker/AnonTalk/tests"
	"github.com/google/uuid"
	"github.com/jackc/fake"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestHealthcheck_GoodPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	roomService := mock.NewMockRoomService(ctrl)
	handler := NewHandler(tests.NewTestLogger(), roomService)
	server := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/healthchek", nil)
	rec := httptest.NewRecorder()
	c := server.NewContext(req, rec)

	err := handler.Healthcheck(c)
	require.NoError(t, err)
	require.Equal(t, rec.Result().StatusCode, http.StatusOK)
}

func TestGetApi_GoodPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	roomService := mock.NewMockRoomService(ctrl)
	handler := NewHandler(tests.NewTestLogger(), roomService)
	server := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/swagger", nil)
	rec := httptest.NewRecorder()
	c := server.NewContext(req, rec)

	err := handler.GetApi(c)
	require.NoError(t, err)
	require.Equal(t, rec.Result().StatusCode, http.StatusOK)

	var api bp.API
	err = json.Unmarshal(rec.Body.Bytes(), &api)
	require.NoError(t, err)

	require.NotNil(t, api.Api)
	require.NotEmpty(t, *api.Api)
}

func TestCreateNewRoom(t *testing.T) {
	cases := []*struct {
		name         string
		req          *http.Request
		rec          *httptest.ResponseRecorder
		c            echo.Context
		expectedName string
	}{
		{
			name:         "empty name",
			req:          newRequest(http.MethodGet, "/room", strings.NewReader(`{"name": ""}`), map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}),
			expectedName: "",
		},
		{
			name:         "non empty name",
			req:          newRequest(http.MethodGet, "/room", strings.NewReader(`{"name": "room_name"}`), map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}),
			expectedName: "room_name",
		},
		{
			name:         "empty request body",
			req:          newRequest(http.MethodGet, "/room", nil, map[string]string{echo.HeaderContentType: echo.MIMEApplicationJSON}),
			expectedName: "",
		},
	}

	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	roomService := mock.NewMockRoomService(ctrl)
	server := echo.New()

	for _, tt := range cases {
		tt.rec = httptest.NewRecorder()
		tt.c = server.NewContext(tt.req, tt.rec)

		roomService.EXPECT().
			CreateNewRoom(gomock.Any(), gomock.Eq(tt.expectedName)).
			Return(&models.Room{UUID: uuid.NewString(), Name: tt.expectedName}, nil).
			AnyTimes()

	}

	handler := NewHandler(tests.NewTestLogger(), roomService)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := handler.CreateNewRoom(tt.c)
			require.NoError(t, err)
			assert.Equal(t, tt.rec.Result().StatusCode, http.StatusCreated)

			var room bp.RoomInfo
			err = json.Unmarshal(tt.rec.Body.Bytes(), &room)
			require.NoError(t, err)

			assert.Equal(t, room.Name, tt.expectedName)
		})
	}

}

func TestGetRoomInfo_GoodPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	room := &models.Room{
		UUID: uuid.NewString(),
		Name: fake.Company(),
	}
	roomService := mock.NewMockRoomService(ctrl)
	roomService.EXPECT().
		GetRoom(gomock.Any(), gomock.Eq(room.UUID)).
		Return(room, nil).
		AnyTimes()

	handler := NewHandler(tests.NewTestLogger(), roomService)

	server := echo.New()
	req := newRequest(http.MethodGet, fmt.Sprintf("/room/%s", room.UUID), nil, nil)
	rec := httptest.NewRecorder()
	c := server.NewContext(req, rec)

	err := handler.GetRoomInfo(c, room.UUID)
	require.NoError(t, err)

	var roomResp bp.RoomInfo
	err = json.Unmarshal(rec.Body.Bytes(), &roomResp)
	require.NoError(t, err)

	assert.Equal(t, room.UUID, roomResp.Id)
	assert.Equal(t, room.Name, roomResp.Name)
}

func newRequest(method string, target string, body io.Reader, headers map[string]string) *http.Request {
	req := httptest.NewRequest(method, target, body)
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req
}
