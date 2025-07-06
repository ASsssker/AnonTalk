package services

import (
	"context"
	"testing"

	"github.com/ASsssker/AnonTalk/internal/models"
	mock "github.com/ASsssker/AnonTalk/mock/services"
	"github.com/ASsssker/AnonTalk/tests"
	"github.com/google/uuid"
	"github.com/jackc/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetRoomInfo_GoodPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	roomProvider := mock.NewMockRoomProvider(ctrl)
	room := &models.Room{
		UUID: uuid.NewString(),
		Name: fake.FirstName(),
	}

	roomProvider.EXPECT().
		GetRoomInfo(gomock.Any(), gomock.Eq(room.UUID)).
		Return(room, nil)

	roomService := NewRoomService(tests.NewTestLogger(), roomProvider)

	actualRoom, err := roomService.GetRoom(context.TODO(), room.UUID)
	require.NoError(t, err)
	assert.Equal(t, room.UUID, actualRoom.UUID)
	assert.Equal(t, room.Name, actualRoom.Name)
}

func TestCreateNewRoom_GoodPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	roomProvider := mock.NewMockRoomProvider(ctrl)
	room := &models.Room{
		UUID: uuid.NewString(),
		Name: fake.FirstName(),
	}

	roomProvider.EXPECT().
		NewRoom(gomock.Any(), gomock.Eq(room.Name)).
		Return(room, nil)

	roomService := NewRoomService(tests.NewTestLogger(), roomProvider)

	actualRoom, err := roomService.CreateNewRoom(context.TODO(), room.Name)
	require.NoError(t, err)
	assert.Equal(t, room.UUID, actualRoom.UUID)
	assert.Equal(t, room.Name, actualRoom.Name)
}
