package manager

import (
	"github.com/golang/mock/gomock"
)

// MockLoopDevice defines the mocker of loopdevice 
type MockLoopDevice struct {
	ctrl      *gomock.Controller
	recorder  *MockLoopDeviceMockRecorder
}

// MockLoopDeviceMockRecorder ... 
type MockLoopDeviceMockRecorder struct {
	mock *MockLoopDevice
}

func (m *MockLoopDevice) EXPECT() *MockLoopDeviceMockRecorder {
	return m.recorder
}

func NewMockLoopDevice(ctrl *gomock.Controller) *MockLoopDevice {
	mock := &MockLoopDevice{ctrl: ctrl}
	mock.recorder = &MockLoopDeviceMockRecorder{mock}
	return mock
}