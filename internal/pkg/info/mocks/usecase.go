// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// GetCities mocks base method.
func (m *MockUseCase) GetCities() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCities")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCities indicates an expected call of GetCities.
func (mr *MockUseCaseMockRecorder) GetCities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCities", reflect.TypeOf((*MockUseCase)(nil).GetCities))
}

// GetCityId mocks base method.
func (m *MockUseCase) GetCityId(city string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCityId", city)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCityId indicates an expected call of GetCityId.
func (mr *MockUseCaseMockRecorder) GetCityId(city interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCityId", reflect.TypeOf((*MockUseCase)(nil).GetCityId), city)
}

// GetEducation mocks base method.
func (m *MockUseCase) GetEducation() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEducation")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEducation indicates an expected call of GetEducation.
func (mr *MockUseCaseMockRecorder) GetEducation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEducation", reflect.TypeOf((*MockUseCase)(nil).GetEducation))
}

// GetEducationId mocks base method.
func (m *MockUseCase) GetEducationId(education string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEducationId", education)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEducationId indicates an expected call of GetEducationId.
func (mr *MockUseCaseMockRecorder) GetEducationId(education interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEducationId", reflect.TypeOf((*MockUseCase)(nil).GetEducationId), education)
}

// GetHashtagId mocks base method.
func (m *MockUseCase) GetHashtagId(hashtagId []string) ([]uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHashtagId", hashtagId)
	ret0, _ := ret[0].([]uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHashtagId indicates an expected call of GetHashtagId.
func (mr *MockUseCaseMockRecorder) GetHashtagId(hashtagId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHashtagId", reflect.TypeOf((*MockUseCase)(nil).GetHashtagId), hashtagId)
}

// GetHashtags mocks base method.
func (m *MockUseCase) GetHashtags() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHashtags")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHashtags indicates an expected call of GetHashtags.
func (mr *MockUseCaseMockRecorder) GetHashtags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHashtags", reflect.TypeOf((*MockUseCase)(nil).GetHashtags))
}

// GetJobId mocks base method.
func (m *MockUseCase) GetJobId(job string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobId", job)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobId indicates an expected call of GetJobId.
func (mr *MockUseCaseMockRecorder) GetJobId(job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobId", reflect.TypeOf((*MockUseCase)(nil).GetJobId), job)
}

// GetJobs mocks base method.
func (m *MockUseCase) GetJobs() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobs indicates an expected call of GetJobs.
func (mr *MockUseCaseMockRecorder) GetJobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobs", reflect.TypeOf((*MockUseCase)(nil).GetJobs))
}

// GetReasons mocks base method.
func (m *MockUseCase) GetReasons() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReasons")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReasons indicates an expected call of GetReasons.
func (mr *MockUseCaseMockRecorder) GetReasons() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReasons", reflect.TypeOf((*MockUseCase)(nil).GetReasons))
}

// GetStatusId mocks base method.
func (m *MockUseCase) GetStatusId(status string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatusId", status)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatusId indicates an expected call of GetStatusId.
func (mr *MockUseCaseMockRecorder) GetStatusId(status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatusId", reflect.TypeOf((*MockUseCase)(nil).GetStatusId), status)
}

// GetStatuses mocks base method.
func (m *MockUseCase) GetStatuses() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatuses")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatuses indicates an expected call of GetStatuses.
func (mr *MockUseCaseMockRecorder) GetStatuses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatuses", reflect.TypeOf((*MockUseCase)(nil).GetStatuses))
}

// GetZodiacId mocks base method.
func (m *MockUseCase) GetZodiacId(zodiac string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetZodiacId", zodiac)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetZodiacId indicates an expected call of GetZodiacId.
func (mr *MockUseCaseMockRecorder) GetZodiacId(zodiac interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetZodiacId", reflect.TypeOf((*MockUseCase)(nil).GetZodiacId), zodiac)
}

// GetZodiacs mocks base method.
func (m *MockUseCase) GetZodiacs() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetZodiacs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetZodiacs indicates an expected call of GetZodiacs.
func (mr *MockUseCaseMockRecorder) GetZodiacs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetZodiacs", reflect.TypeOf((*MockUseCase)(nil).GetZodiacs))
}
