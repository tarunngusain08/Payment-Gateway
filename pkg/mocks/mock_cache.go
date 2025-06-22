package mocks

import (
	"context"

	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCacheStore is a mock of CacheStore interface.
type MockCacheStore struct {
	ctrl     *gomock.Controller
	recorder *MockCacheStoreMockRecorder
}

// MockCacheStoreMockRecorder is the mock recorder for MockCacheStore.
type MockCacheStoreMockRecorder struct {
	mock *MockCacheStore
}

// NewMockCacheStore creates a new mock instance.
func NewMockCacheStore(ctrl *gomock.Controller) *MockCacheStore {
	mock := &MockCacheStore{ctrl: ctrl}
	mock.recorder = &MockCacheStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheStore) EXPECT() *MockCacheStoreMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockCacheStore) Get(ctx context.Context, key string) (interface{}, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCacheStoreMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCacheStore)(nil).Get), ctx, key)
}

// Set mocks base method.
func (m *MockCacheStore) Set(ctx context.Context, key string, value interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCacheStoreMockRecorder) Set(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCacheStore)(nil).Set), ctx, key, value)
}
