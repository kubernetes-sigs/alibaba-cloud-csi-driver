package internal

import (
	"context"
	"errors"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"testing"
)

type MockController struct{}

func (m MockController) VolumeAs() string {
	return "FileSystem"
}

func (m MockController) CreateVolume(context.Context, *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			CapacityBytes: 1024,
			VolumeId:      "id",
		},
	}, nil
}

func (m MockController) DeleteVolume(context.Context, *csi.DeleteVolumeRequest, *corev1.PersistentVolume) (*csi.DeleteVolumeResponse, error) {
	return &csi.DeleteVolumeResponse{}, nil
}

func (m MockController) ControllerExpandVolume(context.Context, *csi.ControllerExpandVolumeRequest, *corev1.PersistentVolume) (*csi.ControllerExpandVolumeResponse, error) {
	return &csi.ControllerExpandVolumeResponse{
		CapacityBytes: 1024,
	}, nil
}

func newMockController(*ControllerConfig) (Controller, error) {
	return MockController{}, nil
}

func newErrorController(*ControllerConfig) (Controller, error) {
	return nil, errors.New("")
}

func TestRegisterControllerMode(t *testing.T) {
	clearControllerInitFuncs()
	RegisterControllerMode(newMockController)
	assert.Len(t, controllerInitFuncs, 1)
}

func clearControllerInitFuncs() {
	controllerInitFuncs = nil
}

func TestNewControllerFactory(t *testing.T) {
	type args struct {
		defaultVolumeAs  string
		controller       ControllerInitFunc
		controllerConfig ControllerConfig
	}

	tests := []struct {
		args     args
		expected *ControllerFactory
		wantErr  bool
	}{
		{
			args: args{
				defaultVolumeAs:  "FileSystem",
				controller:       newMockController,
				controllerConfig: ControllerConfig{},
			},
			expected: &ControllerFactory{
				modes: map[string]Controller{
					"":           MockController{},
					"FileSystem": MockController{},
				},
			},
			wantErr: false,
		},
		{
			args: args{
				defaultVolumeAs:  "",
				controller:       newErrorController,
				controllerConfig: ControllerConfig{},
			},
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		clearControllerInitFuncs()
		RegisterControllerMode(tt.args.controller)
		actual, err := NewControllerFactory(&tt.args.controllerConfig, tt.args.defaultVolumeAs)
		assert.Equal(t, tt.expected, actual)
		if tt.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestVolumeAs(t *testing.T) {
	clearControllerInitFuncs()
	RegisterControllerMode(newMockController)
	factory, _ := NewControllerFactory(&ControllerConfig{}, "FileSystem")

	tests := []struct {
		args     string
		expected Controller
		wantErr  bool
	}{
		{
			args:     "FileSystem",
			expected: MockController{},
			wantErr:  false,
		},
		{
			args:     "InvalidVolumeAs",
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		actual, err := factory.VolumeAs(tt.args)
		assert.Equal(t, actual, tt.expected)
		if tt.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
