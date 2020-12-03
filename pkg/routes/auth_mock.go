// Code generated by MockGen. DO NOT EDIT.
// Source: router_structs.go

// Package routes is a generated GoMock package.
package routes

import (
	gomock "github.com/golang/mock/gomock"
	repository "grinder/pkg/repository"
	reflect "reflect"
	time "time"
)

// MockISessionManager is a mock of ISessionManager interface
type MockISessionManager struct {
	ctrl     *gomock.Controller
	recorder *MockISessionManagerMockRecorder
}

// MockISessionManagerMockRecorder is the mock recorder for MockISessionManager
type MockISessionManagerMockRecorder struct {
	mock *MockISessionManager
}

// NewMockISessionManager creates a new mock instance
func NewMockISessionManager(ctrl *gomock.Controller) *MockISessionManager {
	mock := &MockISessionManager{ctrl: ctrl}
	mock.recorder = &MockISessionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockISessionManager) EXPECT() *MockISessionManagerMockRecorder {
	return m.recorder
}

// CreateSession mocks base method
func (m *MockISessionManager) CreateSession(userID, version int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", userID, version)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession
func (mr *MockISessionManagerMockRecorder) CreateSession(userID, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockISessionManager)(nil).CreateSession), userID, version)
}

// ValidateSession mocks base method
func (m *MockISessionManager) ValidateSession(arg0 string) (int64, int64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateSession", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int64)
	return ret0, ret1
}

// ValidateSession indicates an expected call of ValidateSession
func (mr *MockISessionManagerMockRecorder) ValidateSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateSession", reflect.TypeOf((*MockISessionManager)(nil).ValidateSession), arg0)
}

// GetTokenLiveTime mocks base method
func (m *MockISessionManager) GetTokenLiveTime() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenLiveTime")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// GetTokenLiveTime indicates an expected call of GetTokenLiveTime
func (mr *MockISessionManagerMockRecorder) GetTokenLiveTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenLiveTime", reflect.TypeOf((*MockISessionManager)(nil).GetTokenLiveTime))
}

// MockIUserRepo is a mock of IUserRepo interface
type MockIUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepoMockRecorder
}

// MockIUserRepoMockRecorder is the mock recorder for MockIUserRepo
type MockIUserRepoMockRecorder struct {
	mock *MockIUserRepo
}

// NewMockIUserRepo creates a new mock instance
func NewMockIUserRepo(ctrl *gomock.Controller) *MockIUserRepo {
	mock := &MockIUserRepo{ctrl: ctrl}
	mock.recorder = &MockIUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUserRepo) EXPECT() *MockIUserRepoMockRecorder {
	return m.recorder
}

// RegisterUser mocks base method
func (m *MockIUserRepo) RegisterUser(mail, password string) (int64, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", mail, password)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RegisterUser indicates an expected call of RegisterUser
func (mr *MockIUserRepoMockRecorder) RegisterUser(mail, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockIUserRepo)(nil).RegisterUser), mail, password)
}

// LoginUser mocks base method
func (m *MockIUserRepo) LoginUser(mail, password string) (int64, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", mail, password)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoginUser indicates an expected call of LoginUser
func (mr *MockIUserRepoMockRecorder) LoginUser(mail, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockIUserRepo)(nil).LoginUser), mail, password)
}

// CheckVersion mocks base method
func (m *MockIUserRepo) CheckVersion(userID, version int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckVersion", userID, version)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckVersion indicates an expected call of CheckVersion
func (mr *MockIUserRepoMockRecorder) CheckVersion(userID, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckVersion", reflect.TypeOf((*MockIUserRepo)(nil).CheckVersion), userID, version)
}

// GetUser mocks base method
func (m *MockIUserRepo) GetUser(userID, version int64) (*repository.User, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", userID, version)
	ret0, _ := ret[0].(*repository.User)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUser indicates an expected call of GetUser
func (mr *MockIUserRepoMockRecorder) GetUser(userID, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIUserRepo)(nil).GetUser), userID, version)
}

// UpdateUser mocks base method
func (m *MockIUserRepo) UpdateUser(u repository.User) (*repository.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", u)
	ret0, _ := ret[0].(*repository.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockIUserRepoMockRecorder) UpdateUser(u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockIUserRepo)(nil).UpdateUser), u)
}

// UpdateUserPassword mocks base method
func (m *MockIUserRepo) UpdateUserPassword(userID, userV int64, oldPass, newPass string) (*repository.User, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPassword", userID, userV, oldPass, newPass)
	ret0, _ := ret[0].(*repository.User)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateUserPassword indicates an expected call of UpdateUserPassword
func (mr *MockIUserRepoMockRecorder) UpdateUserPassword(userID, userV, oldPass, newPass interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPassword", reflect.TypeOf((*MockIUserRepo)(nil).UpdateUserPassword), userID, userV, oldPass, newPass)
}