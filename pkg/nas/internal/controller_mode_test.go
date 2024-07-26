package internal

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
				defaultVolumeAs:  MockVolumeAs,
				controller:       newMockController,
				controllerConfig: ControllerConfig{},
			},
			expected: &ControllerFactory{
				Modes: map[string]Controller{
					"":           MockController{},
					MockVolumeAs: MockController{},
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
	factory, _ := NewControllerFactory(&ControllerConfig{}, MockVolumeAs)

	tests := []struct {
		name     string
		args     string
		expected Controller
		wantErr  bool
	}{
		{
			name:     "Valid volume as",
			args:     MockVolumeAs,
			expected: MockController{},
			wantErr:  false,
		},
		{
			name:     "Invalid volume as",
			args:     "InvalidVolumeAs",
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := factory.VolumeAs(tt.args)
			assert.Equal(t, actual, tt.expected)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
