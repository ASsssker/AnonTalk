package room

import (
	"errors"
	"testing"

	mock "github.com/ASsssker/AnonTalk/mock/room"
	"github.com/ASsssker/AnonTalk/tests"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var clientsCount = 10

func TestNewRoom_GoodPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	clients := make([]RoomClient, 0, clientsCount)
	for range clientsCount {
		client := mock.NewMockRoomClient(ctrl)
		client.EXPECT().
			GetID().
			Return(uuid.NewString()).
			AnyTimes()

		clients = append(clients, client)
	}

	room := NewRoom(tests.NewTestLogger(), "test-room", clients...)
	assert.Equal(t, clientsCount, len(room.clients))
}

func TestRoomAddClients_GoodPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	clients := make([]RoomClient, 0, clientsCount)
	for range clientsCount {
		client := mock.NewMockRoomClient(ctrl)
		client.EXPECT().
			GetID().
			Return(uuid.NewString()).
			AnyTimes()

		clients = append(clients, client)
	}

	room := NewRoom(tests.NewTestLogger(), "test-room")
	room.AddClients(clients...)
	assert.Equal(t, clientsCount, len(room.clients))
}

func TestRoomDeleteClients_GoodPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	clients := make([]RoomClient, 0, clientsCount)
	for range clientsCount {
		client := mock.NewMockRoomClient(ctrl)
		client.EXPECT().
			GetID().
			Return(uuid.NewString()).
			AnyTimes()

		clients = append(clients, client)
	}

	room := NewRoom(tests.NewTestLogger(), "test-room", clients...)
	clientsID := make([]string, 0, clientsCount)
	for _, client := range clients {
		clientsID = append(clientsID, client.GetID())
	}

	room.DeleteClients(clientsID...)
	assert.Equal(t, 0, len(room.clients))
}

func TestRoomBroadcast_GoodPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	clients := make([]RoomClient, 0, clientsCount)
	for range clientsCount - 1 {
		client := mock.NewMockRoomClient(ctrl)
		client.EXPECT().
			GetID().
			Return(uuid.NewString()).
			AnyTimes()

		client.EXPECT().
			Write(gomock.Any()).
			Return(nil).
			AnyTimes()

		clients = append(clients, client)
	}

	senderClient := mock.NewMockRoomClient(ctrl)
	senderClient.EXPECT().
		GetID().
		Return(uuid.NewString()).
		AnyTimes()

	senderClient.EXPECT().
		Write(gomock.Any()).
		Return(errors.New("the message was returned to the sender")).
		AnyTimes()

	clients = append(clients, senderClient)

	room := NewRoom(tests.NewTestLogger(), "test-room", clients...)
	err := room.Broadcast(senderClient.GetID(), "test-message")
	assert.NoError(t, err)
}
