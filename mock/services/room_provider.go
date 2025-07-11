// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ASsssker/AnonTalk/internal/services (interfaces: RoomProvider)
//
// Generated by this command:
//
//	mockgen -package mock -destination /home/asker/code/AnonTalk/mock/services/room_provider.go . RoomProvider
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/ASsssker/AnonTalk/internal/models"
	room "github.com/ASsssker/AnonTalk/internal/room"
	gomock "go.uber.org/mock/gomock"
)

// MockRoomProvider is a mock of RoomProvider interface.
type MockRoomProvider struct {
	ctrl     *gomock.Controller
	recorder *MockRoomProviderMockRecorder
	isgomock struct{}
}

// MockRoomProviderMockRecorder is the mock recorder for MockRoomProvider.
type MockRoomProviderMockRecorder struct {
	mock *MockRoomProvider
}

// NewMockRoomProvider creates a new mock instance.
func NewMockRoomProvider(ctrl *gomock.Controller) *MockRoomProvider {
	mock := &MockRoomProvider{ctrl: ctrl}
	mock.recorder = &MockRoomProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoomProvider) EXPECT() *MockRoomProviderMockRecorder {
	return m.recorder
}

// DeleteRoom mocks base method.
func (m *MockRoomProvider) DeleteRoom(ctx context.Context, id string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRoom", ctx, id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRoom indicates an expected call of DeleteRoom.
func (mr *MockRoomProviderMockRecorder) DeleteRoom(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoom", reflect.TypeOf((*MockRoomProvider)(nil).DeleteRoom), ctx, id)
}

// GetRoom mocks base method.
func (m *MockRoomProvider) GetRoom(ctx context.Context, id string) (*room.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoom", ctx, id)
	ret0, _ := ret[0].(*room.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoom indicates an expected call of GetRoom.
func (mr *MockRoomProviderMockRecorder) GetRoom(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoom", reflect.TypeOf((*MockRoomProvider)(nil).GetRoom), ctx, id)
}

// GetRoomInfo mocks base method.
func (m *MockRoomProvider) GetRoomInfo(ctx context.Context, id string) (*models.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomInfo", ctx, id)
	ret0, _ := ret[0].(*models.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomInfo indicates an expected call of GetRoomInfo.
func (mr *MockRoomProviderMockRecorder) GetRoomInfo(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomInfo", reflect.TypeOf((*MockRoomProvider)(nil).GetRoomInfo), ctx, id)
}

// NewRoom mocks base method.
func (m *MockRoomProvider) NewRoom(ctx context.Context, name string) (*models.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRoom", ctx, name)
	ret0, _ := ret[0].(*models.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewRoom indicates an expected call of NewRoom.
func (mr *MockRoomProviderMockRecorder) NewRoom(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRoom", reflect.TypeOf((*MockRoomProvider)(nil).NewRoom), ctx, name)
}
