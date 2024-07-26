package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/container-storage-interface/spec/lib/go/csi"
	corev1 "k8s.io/api/core/v1"
)

const MockVolumeAs = "Mock"

type MockController struct{}
type MockErrorController struct{}

func (m MockController) VolumeAs() string {
	return MockVolumeAs
}

func (m MockController) CreateVolume(context.Context, *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	fmt.Println("MockController.CreateVolume")
	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			CapacityBytes: 1024,
			VolumeId:      "id",
			VolumeContext: map[string]string{},
		},
	}, nil
}

func (m MockController) DeleteVolume(context.Context, *csi.DeleteVolumeRequest, *corev1.PersistentVolume) (*csi.DeleteVolumeResponse, error) {
	fmt.Println("MockController.DeleteVolume")
	return &csi.DeleteVolumeResponse{}, nil
}

func (m MockController) ControllerExpandVolume(context.Context, *csi.ControllerExpandVolumeRequest, *corev1.PersistentVolume) (*csi.ControllerExpandVolumeResponse, error) {
	fmt.Println("MockController.ControllerExpandVolume")
	return &csi.ControllerExpandVolumeResponse{
		CapacityBytes: 1024,
	}, nil
}

func (m MockErrorController) VolumeAs() string {
	return ""
}

func (m MockErrorController) CreateVolume(context.Context, *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	fmt.Println("MockErrorController.CreateVolume")
	return nil, errors.New("")
}

func (m MockErrorController) DeleteVolume(context.Context, *csi.DeleteVolumeRequest, *corev1.PersistentVolume) (*csi.DeleteVolumeResponse, error) {
	fmt.Println("MockErrorController.DeleteVolume")
	return nil, errors.New("")
}

func (m MockErrorController) ControllerExpandVolume(context.Context, *csi.ControllerExpandVolumeRequest, *corev1.PersistentVolume) (*csi.ControllerExpandVolumeResponse, error) {
	fmt.Println("MockErrorController.ControllerExpandVolume")
	return nil, errors.New("")
}
