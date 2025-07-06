package room

import (
	"errors"
	"sync"
	"testing"

	mock "github.com/ASsssker/AnonTalk/mock/room"
	"github.com/ASsssker/AnonTalk/tests"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var clientsCount = 100

func TestRoomAddClients_GoodPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	room := NewRoom(tests.NewTestLogger(), "test-room")
	waitChan := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(clientsCount)

	for range clientsCount {
		client := mock.NewMockRoomClient(ctrl)
		client.EXPECT().
			GetID().
			Return(uuid.NewString()).
			AnyTimes()

		client.EXPECT().
			MsgSubscribe(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ any, _ any) error { <-waitChan; wg.Wait(); return nil }).
			AnyTimes()

		client.EXPECT().
			Close(gomock.Any()).
			Return(nil).
			AnyTimes()

		go room.AddClient(client)
	}

	for range clientsCount {
		waitChan <- struct{}{}
		defer wg.Done()
	}

	assert.Equal(t, clientsCount, room.ClientsCount())
}

func TestRoomDeleteClients_GoodPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	room := NewRoom(tests.NewTestLogger(), "test-room")
	wg := sync.WaitGroup{}
	wg.Add(clientsCount)

	for range clientsCount {
		client := mock.NewMockRoomClient(ctrl)
		client.EXPECT().
			GetID().
			Return(uuid.NewString()).
			AnyTimes()

		client.EXPECT().
			MsgSubscribe(gomock.Any(), gomock.Any()).
			Return(nil).
			AnyTimes()

		client.EXPECT().
			Close(gomock.Any()).
			Return(nil).
			DoAndReturn(func(_ any) error { defer wg.Done(); return nil }).
			AnyTimes()

		go room.AddClient(client)
	}

	wg.Wait()
	assert.Equal(t, 0, room.ClientsCount())
}

func TestRoomBroadcast_GoodPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	room := NewRoom(tests.NewTestLogger(), "test-room")
	wg := sync.WaitGroup{}
	wg.Add(clientsCount)

	for range clientsCount - 1 {
		client := mock.NewMockRoomClient(ctrl)
		client.EXPECT().
			GetID().
			Return(uuid.NewString()).
			AnyTimes()

		client.EXPECT().
			Write(gomock.Any(), gomock.Any()).
			Return(nil).
			AnyTimes()

		client.EXPECT().
			MsgSubscribe(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ any, _ any) error { defer wg.Wait(); return nil }).
			AnyTimes()

		client.EXPECT().
			Close(gomock.Any()).
			Return(nil).
			AnyTimes()

		go room.AddClient(client)
	}

	senderClient := mock.NewMockRoomClient(ctrl)
	senderClient.EXPECT().
		GetID().
		Return(uuid.NewString()).
		AnyTimes()

	senderClient.EXPECT().
		Write(gomock.Any(), gomock.Any()).
		Return(errors.New("the message was returned to the sender")).
		AnyTimes()

	senderClient.EXPECT().
		MsgSubscribe(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ any, _ any) error { defer wg.Wait(); return nil }).
		AnyTimes()

	senderClient.EXPECT().
		Close(gomock.Any()).
		Return(nil).
		AnyTimes()

	go room.AddClient(senderClient)
	err := room.Broadcast(senderClient.GetID(), "test-message")
	assert.NoError(t, err)

	for range clientsCount {
		wg.Done()
	}
}
