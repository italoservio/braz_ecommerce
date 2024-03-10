// Code generated by MockGen. DO NOT EDIT.
// Source: services/users/app/update_user_by_id.go
//
// Generated by this command:
//
//	mockgen -source=services/users/app/update_user_by_id.go -destination=services/users/mocks/update_user_by_id_interface_mock.go -package=mocks -write_generate_directive
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	app "github.com/italoservio/braz_ecommerce/services/users/app"
	gomock "go.uber.org/mock/gomock"
)

//go:generate mockgen -source=services/users/app/update_user_by_id.go -destination=services/users/mocks/update_user_by_id_interface_mock.go -package=mocks -write_generate_directive

// MockUpdateUserByIdInterface is a mock of UpdateUserByIdInterface interface.
type MockUpdateUserByIdInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateUserByIdInterfaceMockRecorder
}

// MockUpdateUserByIdInterfaceMockRecorder is the mock recorder for MockUpdateUserByIdInterface.
type MockUpdateUserByIdInterfaceMockRecorder struct {
	mock *MockUpdateUserByIdInterface
}

// NewMockUpdateUserByIdInterface creates a new mock instance.
func NewMockUpdateUserByIdInterface(ctrl *gomock.Controller) *MockUpdateUserByIdInterface {
	mock := &MockUpdateUserByIdInterface{ctrl: ctrl}
	mock.recorder = &MockUpdateUserByIdInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateUserByIdInterface) EXPECT() *MockUpdateUserByIdInterfaceMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockUpdateUserByIdInterface) Do(ctx context.Context, id string, input *app.UpdateUserByIdInput) (*app.UpdateUserByIdOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", ctx, id, input)
	ret0, _ := ret[0].(*app.UpdateUserByIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockUpdateUserByIdInterfaceMockRecorder) Do(ctx, id, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockUpdateUserByIdInterface)(nil).Do), ctx, id, input)
}
