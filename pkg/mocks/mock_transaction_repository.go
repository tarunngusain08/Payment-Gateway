// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/transaction_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	constants "Payment-Gateway/internal/constants"
	models "Payment-Gateway/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTransactionRepository is a mock of TransactionRepository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// CreateTransaction mocks base method.
func (m *MockTransactionRepository) CreateTransaction(tx *models.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockTransactionRepositoryMockRecorder) CreateTransaction(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockTransactionRepository)(nil).CreateTransaction), tx)
}

// GetTransactionByID mocks base method.
func (m *MockTransactionRepository) GetTransactionByID(id string) (*models.Transaction, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByID", id)
	ret0, _ := ret[0].(*models.Transaction)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetTransactionByID indicates an expected call of GetTransactionByID.
func (mr *MockTransactionRepositoryMockRecorder) GetTransactionByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByID", reflect.TypeOf((*MockTransactionRepository)(nil).GetTransactionByID), id)
}

// UpdateTransactionStatus mocks base method.
func (m *MockTransactionRepository) UpdateTransactionStatus(id string, status constants.TransactionStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransactionStatus", id, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTransactionStatus indicates an expected call of UpdateTransactionStatus.
func (mr *MockTransactionRepositoryMockRecorder) UpdateTransactionStatus(id, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransactionStatus", reflect.TypeOf((*MockTransactionRepository)(nil).UpdateTransactionStatus), id, status)
}
