// Code generated by MockGen. DO NOT EDIT.
// Source: services/users/app/update_user.go
//
// Generated by this command:
//
//	mockgen -source=services/users/app/update_user.go -destination=services/users/mocks/update_user_interface_mock.go -package=mocks -write_generate_directive
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	app "github.com/italoservio/braz_ecommerce/services/users/app"
	gomock "go.uber.org/mock/gomock"
)

//go:generate mockgen -source=services/users/app/update_user.go -destination=services/users/mocks/update_user_interface_mock.go -package=mocks -write_generate_directive

// MockUpdateUserInterface is a mock of UpdateUserInterface interface.
type MockUpdateUserInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateUserInterfaceMockRecorder
}

// MockUpdateUserInterfaceMockRecorder is the mock recorder for MockUpdateUserInterface.
type MockUpdateUserInterfaceMockRecorder struct {
	mock *MockUpdateUserInterface
}

// NewMockUpdateUserInterface creates a new mock instance.
func NewMockUpdateUserInterface(ctrl *gomock.Controller) *MockUpdateUserInterface {
	mock := &MockUpdateUserInterface{ctrl: ctrl}
	mock.recorder = &MockUpdateUserInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateUserInterface) EXPECT() *MockUpdateUserInterfaceMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockUpdateUserInterface) Do(updateUser *app.UpdateUserInput, id string) (*app.UpdateUserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", updateUser, id)
	ret0, _ := ret[0].(*app.UpdateUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockUpdateUserInterfaceMockRecorder) Do(updateUser, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockUpdateUserInterface)(nil).Do), updateUser, id)
}
