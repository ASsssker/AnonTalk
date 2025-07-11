// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ASsssker/AnonTalk/internal/rest/v1 (interfaces: RoomService)
//
// Generated by this command:
//
//	mockgen -package mock -destination /home/asker/code/AnonTalk/mock/rest/v1/handler/services.go . RoomService
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/ASsssker/AnonTalk/internal/models"
	websocket "github.com/gorilla/websocket"
	gomock "go.uber.org/mock/gomock"
)

// MockRoomService is a mock of RoomService interface.
type MockRoomService struct {
	ctrl     *gomock.Controller
	recorder *MockRoomServiceMockRecorder
	isgomock struct{}
}

// MockRoomServiceMockRecorder is the mock recorder for MockRoomService.
type MockRoomServiceMockRecorder struct {
	mock *MockRoomService
}

// NewMockRoomService creates a new mock instance.
func NewMockRoomService(ctrl *gomock.Controller) *MockRoomService {
	mock := &MockRoomService{ctrl: ctrl}
	mock.recorder = &MockRoomServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoomService) EXPECT() *MockRoomServiceMockRecorder {
	return m.recorder
}

// AddUserToRoom mocks base method.
func (m *MockRoomService) AddUserToRoom(ctx context.Context, roomID, clientName string, clientConn *websocket.Conn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserToRoom", ctx, roomID, clientName, clientConn)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserToRoom indicates an expected call of AddUserToRoom.
func (mr *MockRoomServiceMockRecorder) AddUserToRoom(ctx, roomID, clientName, clientConn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserToRoom", reflect.TypeOf((*MockRoomService)(nil).AddUserToRoom), ctx, roomID, clientName, clientConn)
}

// CreateNewRoom mocks base method.
func (m *MockRoomService) CreateNewRoom(ctx context.Context, roomName string) (*models.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewRoom", ctx, roomName)
	ret0, _ := ret[0].(*models.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewRoom indicates an expected call of CreateNewRoom.
func (mr *MockRoomServiceMockRecorder) CreateNewRoom(ctx, roomName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewRoom", reflect.TypeOf((*MockRoomService)(nil).CreateNewRoom), ctx, roomName)
}

// GetRoom mocks base method.
func (m *MockRoomService) GetRoom(ctx context.Context, id string) (*models.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoom", ctx, id)
	ret0, _ := ret[0].(*models.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoom indicates an expected call of GetRoom.
func (mr *MockRoomServiceMockRecorder) GetRoom(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoom", reflect.TypeOf((*MockRoomService)(nil).GetRoom), ctx, id)
}
