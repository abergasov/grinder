// Code generated by MockGen. DO NOT EDIT.
// Source: router_structs.go

// Package routes is a generated GoMock package.
package routes

import (
	gin "github.com/gin-gonic/gin"
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

// AuthMiddleware mocks base method
func (m *MockISessionManager) AuthMiddleware(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AuthMiddleware", arg0)
}

// AuthMiddleware indicates an expected call of AuthMiddleware
func (mr *MockISessionManagerMockRecorder) AuthMiddleware(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthMiddleware", reflect.TypeOf((*MockISessionManager)(nil).AuthMiddleware), arg0)
}

// GetUserAndVersion mocks base method
func (m *MockISessionManager) GetUserAndVersion(c *gin.Context) (int64, int64, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAndVersion", c)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// GetUserAndVersion indicates an expected call of GetUserAndVersion
func (mr *MockISessionManagerMockRecorder) GetUserAndVersion(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAndVersion", reflect.TypeOf((*MockISessionManager)(nil).GetUserAndVersion), c)
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
func (m *MockIUserRepo) UpdateUser(u *repository.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", u)
	ret0, _ := ret[0].(error)
	return ret0
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

// MockIPersonsRepo is a mock of IPersonsRepo interface
type MockIPersonsRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIPersonsRepoMockRecorder
}

// MockIPersonsRepoMockRecorder is the mock recorder for MockIPersonsRepo
type MockIPersonsRepoMockRecorder struct {
	mock *MockIPersonsRepo
}

// NewMockIPersonsRepo creates a new mock instance
func NewMockIPersonsRepo(ctrl *gomock.Controller) *MockIPersonsRepo {
	mock := &MockIPersonsRepo{ctrl: ctrl}
	mock.recorder = &MockIPersonsRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPersonsRepo) EXPECT() *MockIPersonsRepoMockRecorder {
	return m.recorder
}

// LoadPersons mocks base method
func (m *MockIPersonsRepo) LoadPersons(offset int64) ([]repository.Person, []repository.PersonRight, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadPersons", offset)
	ret0, _ := ret[0].([]repository.Person)
	ret1, _ := ret[1].([]repository.PersonRight)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoadPersons indicates an expected call of LoadPersons
func (mr *MockIPersonsRepoMockRecorder) LoadPersons(offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadPersons", reflect.TypeOf((*MockIPersonsRepo)(nil).LoadPersons), offset)
}

// GetRightsMap mocks base method
func (m *MockIPersonsRepo) GetRightsMap() map[int64]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRightsMap")
	ret0, _ := ret[0].(map[int64]string)
	return ret0
}

// GetRightsMap indicates an expected call of GetRightsMap
func (mr *MockIPersonsRepoMockRecorder) GetRightsMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRightsMap", reflect.TypeOf((*MockIPersonsRepo)(nil).GetRightsMap))
}

// UpdateUser mocks base method
func (m *MockIPersonsRepo) UpdateUser(userID int64, firstName, lastName, email string, active bool) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", userID, firstName, lastName, email, active)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockIPersonsRepoMockRecorder) UpdateUser(userID, firstName, lastName, email, active interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockIPersonsRepo)(nil).UpdateUser), userID, firstName, lastName, email, active)
}

// MockIRightsChecker is a mock of IRightsChecker interface
type MockIRightsChecker struct {
	ctrl     *gomock.Controller
	recorder *MockIRightsCheckerMockRecorder
}

// MockIRightsCheckerMockRecorder is the mock recorder for MockIRightsChecker
type MockIRightsCheckerMockRecorder struct {
	mock *MockIRightsChecker
}

// NewMockIRightsChecker creates a new mock instance
func NewMockIRightsChecker(ctrl *gomock.Controller) *MockIRightsChecker {
	mock := &MockIRightsChecker{ctrl: ctrl}
	mock.recorder = &MockIRightsCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIRightsChecker) EXPECT() *MockIRightsCheckerMockRecorder {
	return m.recorder
}

// CheckRight mocks base method
func (m *MockIRightsChecker) CheckRight(rights []int64, ver func(*gin.Context) (int64, int64, bool)) gin.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRight", rights, ver)
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// CheckRight indicates an expected call of CheckRight
func (mr *MockIRightsCheckerMockRecorder) CheckRight(rights, ver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRight", reflect.TypeOf((*MockIRightsChecker)(nil).CheckRight), rights, ver)
}
