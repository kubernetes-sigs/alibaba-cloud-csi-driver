package internal

import (
	"context"
	"fmt"

	"github.com/container-storage-interface/spec/lib/go/csi"
	corev1 "k8s.io/api/core/v1"
)

type Controller interface {
	VolumeAs() string
	CreateVolume(context.Context, *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error)
	DeleteVolume(context.Context, *csi.DeleteVolumeRequest, *corev1.PersistentVolume) (*csi.DeleteVolumeResponse, error)
	ControllerExpandVolume(context.Context, *csi.ControllerExpandVolumeRequest, *corev1.PersistentVolume) (*csi.ControllerExpandVolumeResponse, error)
}

type ControllerInitFunc func(*ControllerConfig) (Controller, error)

var controllerInitFuncs []ControllerInitFunc

func RegisterControllerMode(f ControllerInitFunc) {
	controllerInitFuncs = append(controllerInitFuncs, f)
}

type ControllerFactory struct {
	modes map[string]Controller
}

func NewControllerFactory(config *ControllerConfig, defaultVolumeAs string) (*ControllerFactory, error) {
	modes := map[string]Controller{}
	for _, f := range controllerInitFuncs {
		mode, err := f(config)
		if err != nil {
			return nil, err
		}
		volumeAs := mode.VolumeAs()
		modes[volumeAs] = mode
		if volumeAs == defaultVolumeAs {
			modes[""] = mode
		}
	}
	return &ControllerFactory{modes: modes}, nil
}

func (fac *ControllerFactory) VolumeAs(what string) (Controller, error) {
	c, ok := fac.modes[what]
	if !ok {
		return nil, fmt.Errorf("invalid volumeAs: %q", what)
	}
	return c, nil
}
