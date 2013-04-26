// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/mozilla-services/heka/pipeline (interfaces: PluginRunner)

package pipeline

import (
	gomock "code.google.com/p/gomock/gomock"
)

// Mock of PluginRunner interface
type MockPluginRunner struct {
	ctrl     *gomock.Controller
	recorder *_MockPluginRunnerRecorder
}

// Recorder for MockPluginRunner (not exported)
type _MockPluginRunnerRecorder struct {
	mock *MockPluginRunner
}

func NewMockPluginRunner(ctrl *gomock.Controller) *MockPluginRunner {
	mock := &MockPluginRunner{ctrl: ctrl}
	mock.recorder = &_MockPluginRunnerRecorder{mock}
	return mock
}

func (_m *MockPluginRunner) EXPECT() *_MockPluginRunnerRecorder {
	return _m.recorder
}

func (_m *MockPluginRunner) LogError(_param0 error) {
	_m.ctrl.Call(_m, "LogError", _param0)
}

func (_mr *_MockPluginRunnerRecorder) LogError(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LogError", arg0)
}

func (_m *MockPluginRunner) LogMessage(_param0 string) {
	_m.ctrl.Call(_m, "LogMessage", _param0)
}

func (_mr *_MockPluginRunnerRecorder) LogMessage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LogMessage", arg0)
}

func (_m *MockPluginRunner) Name() string {
	ret := _m.ctrl.Call(_m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockPluginRunnerRecorder) Name() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Name")
}

func (_m *MockPluginRunner) Plugin() Plugin {
	ret := _m.ctrl.Call(_m, "Plugin")
	ret0, _ := ret[0].(Plugin)
	return ret0
}

func (_mr *_MockPluginRunnerRecorder) Plugin() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Plugin")
}

func (_m *MockPluginRunner) SetName(_param0 string) {
	_m.ctrl.Call(_m, "SetName", _param0)
}

func (_mr *_MockPluginRunnerRecorder) SetName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetName", arg0)
}