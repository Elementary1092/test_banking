// Code generated by MockGen. DO NOT EDIT.
// Source: /home/quark/alif/internal/domain/customer/command/dao.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/Elementary1092/test_banking/internal/domain/customer/command/model"
	gomock "github.com/golang/mock/gomock"
)

// MockWriteDAO is a mock of WriteDAO interface.
type MockWriteDAO struct {
	ctrl     *gomock.Controller
	recorder *MockWriteDAOMockRecorder
}

// MockWriteDAOMockRecorder is the mock recorder for MockWriteDAO.
type MockWriteDAOMockRecorder struct {
	mock *MockWriteDAO
}

// NewMockWriteDAO creates a new mock instance.
func NewMockWriteDAO(ctrl *gomock.Controller) *MockWriteDAO {
	mock := &MockWriteDAO{ctrl: ctrl}
	mock.recorder = &MockWriteDAOMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriteDAO) EXPECT() *MockWriteDAOMockRecorder {
	return m.recorder
}

// CreateCustomer mocks base method.
func (m *MockWriteDAO) CreateCustomer(ctx context.Context, customer *model.WriteModel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCustomer", ctx, customer)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCustomer indicates an expected call of CreateCustomer.
func (mr *MockWriteDAOMockRecorder) CreateCustomer(ctx, customer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCustomer", reflect.TypeOf((*MockWriteDAO)(nil).CreateCustomer), ctx, customer)
}
