// Automatically generated by MockGen. DO NOT EDIT!
// Source: sqllib.go

package sqllib_test

import (
	sql "database/sql"
	gomock "github.com/golang/mock/gomock"
)

// Mock of DB interface
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *_MockDBRecorder
}

// Recorder for MockDB (not exported)
type _MockDBRecorder struct {
	mock *MockDB
}

func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &_MockDBRecorder{mock}
	return mock
}

func (_m *MockDB) EXPECT() *_MockDBRecorder {
	return _m.recorder
}

func (_m *MockDB) Prepare(_param0 string) (*sql.Stmt, error) {
	ret := _m.ctrl.Call(_m, "Prepare", _param0)
	ret0, _ := ret[0].(*sql.Stmt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDBRecorder) Prepare(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Prepare", arg0)
}
