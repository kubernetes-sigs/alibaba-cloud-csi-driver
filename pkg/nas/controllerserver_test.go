package nas

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	mytesting "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/testing"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/fake"
)

const (
	mockPVName = "MockPV"

	mockInvalidPVName = "MockInvalidPV"

	mockPVJson = `
{
	"apiVersion": "v1",
	"kind": "PersistentVolume",
	"metadata": {
		"name": "` + mockPVName + `"
	},
	"spec": {
		"csi": {
			"volumeAttributes": {
				"volumeAs": "` + internal.MockVolumeAs + `"
			}
		}
	}
}`

	mockInvalidVolumeAsPVJson = `
{
	"apiVersion": "v1",
	"kind": "PersistentVolume",
	"metadata": {
		"name": "` + mockInvalidPVName + `"
	},
	"spec": {
		"csi": {
			"volumeAttributes": {
				"volumeAs": "InvalidVolumeAs"
			}
		}
	}
}`
)

func TestValidateCreateVolumeRequestValid(t *testing.T) {
	type args struct {
		name       string
		parameters map[string]string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid volume request",
			args: args{
				name: "ValidVolume",
				parameters: map[string]string{
					"key": "value",
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid volume request",
			args: args{
				name: "InvalidVolume",
				parameters: map[string]string{
					"key": "(invalid-value)",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &csi.CreateVolumeRequest{
				Name:       tt.args.name,
				Parameters: tt.args.parameters,
			}
			err := validateCreateVolumeRequest(req)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestControllerServer_CreateVolume(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)

	req := &csi.CreateVolumeRequest{
		Name: "TestVolume",
		Parameters: map[string]string{
			"volumeAs": internal.MockVolumeAs,
			"mode":     "mode",
			"modeType": "",
			"options":  "options",
		},
	}
	resp, err := cs.CreateVolume(context.Background(), req)
	assert.NotNil(t, resp)
	assert.NoError(t, err)
}

func newMockControllerServer() *controllerServer {
	internal.RegisterControllerMode(newMockController)
	var pv corev1.PersistentVolume
	_ = json.Unmarshal([]byte(mockPVJson), &pv)
	return &controllerServer{
		ControllerFactory: &internal.ControllerFactory{
			Modes: map[string]internal.Controller{
				internal.MockVolumeAs: internal.MockController{},
			},
		},
		kubeClient: fake.NewSimpleClientset(&pv),
		locks:      utils.NewVolumeLocks(),
	}
}

func newMockController(*internal.ControllerConfig) (internal.Controller, error) {
	return internal.MockController{}, nil
}

func TestControllerServer_CreateVolumeError(t *testing.T) {
	cs := newMockErrorControllerServer()
	assert.NotNil(t, cs)

	req := &csi.CreateVolumeRequest{
		Name: "TestVolume",
		Parameters: map[string]string{
			"volumeAs": internal.MockVolumeAs,
		},
	}
	resp, err := cs.CreateVolume(context.Background(), req)
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func newMockErrorControllerServer() *controllerServer {
	internal.RegisterControllerMode(newMockErrorController)
	var pv, invalidVolumeAsPV corev1.PersistentVolume
	_ = json.Unmarshal([]byte(mockPVJson), &pv)
	_ = json.Unmarshal([]byte(mockInvalidVolumeAsPVJson), &invalidVolumeAsPV)
	return &controllerServer{
		ControllerFactory: &internal.ControllerFactory{
			Modes: map[string]internal.Controller{
				internal.MockVolumeAs: internal.MockErrorController{},
			},
		},
		kubeClient: fake.NewSimpleClientset(&pv, &invalidVolumeAsPV),
		locks:      utils.NewVolumeLocks(),
	}
}

func newMockErrorController(*internal.ControllerConfig) (internal.Controller, error) {
	return internal.MockErrorController{}, nil
}

func TestControllerServer_CreateVolumeInvalidRequest(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)

	req := &csi.CreateVolumeRequest{
		Name: "TestVolume",
		Parameters: map[string]string{
			"key": "(invalid value)",
		},
	}
	resp, err := cs.CreateVolume(context.Background(), req)
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestControllerServer_CreateVolumeInvalidVolumeAs(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)

	req := &csi.CreateVolumeRequest{
		Name: "TestVolume",
		Parameters: map[string]string{
			"volumeAs": "InvalidVolumeAs",
		},
	}
	resp, err := cs.CreateVolume(context.Background(), req)
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestControllerServer_DeleteVolume(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)

	resp, err := cs.DeleteVolume(context.Background(), &csi.DeleteVolumeRequest{VolumeId: mockPVName})
	assert.NotNil(t, resp)
	assert.NoError(t, err)
}

func TestControllerServer_DeleteVolumeInvalidVolumeAs(t *testing.T) {
	cs := newMockErrorControllerServer()
	assert.NotNil(t, cs)

	resp, err := cs.DeleteVolume(context.Background(), &csi.DeleteVolumeRequest{VolumeId: mockInvalidPVName})
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestControllerServer_DeleteVolumePVGetError(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)

	resp, err := cs.DeleteVolume(context.Background(), &csi.DeleteVolumeRequest{VolumeId: ""})
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestControllerServer_DeleteVolumeError(t *testing.T) {
	cs := newMockErrorControllerServer()
	assert.NotNil(t, cs)

	_, err := cs.DeleteVolume(context.Background(), &csi.DeleteVolumeRequest{VolumeId: mockPVName})
	assert.Error(t, err)
}

func TestControllerServer_ControllerExpandVolume(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)

	resp, err := cs.ControllerExpandVolume(context.Background(), &csi.ControllerExpandVolumeRequest{VolumeId: mockPVName})
	assert.NotNil(t, resp)
	assert.NoError(t, err)
}

func TestControllerServer_ControllerExpandVolumeInvalidVolumeAs(t *testing.T) {
	cs := newMockErrorControllerServer()
	assert.NotNil(t, cs)

	resp, err := cs.ControllerExpandVolume(context.Background(), &csi.ControllerExpandVolumeRequest{VolumeId: mockInvalidPVName})
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestControllerServer_ControllerExpandVolumePVGetError(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)

	resp, err := cs.ControllerExpandVolume(context.Background(), &csi.ControllerExpandVolumeRequest{VolumeId: ""})
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func TestControllerServer_ControllerExpandVolumeError(t *testing.T) {
	cs := newMockErrorControllerServer()
	assert.NotNil(t, cs)

	_, err := cs.ControllerExpandVolume(context.Background(), &csi.ControllerExpandVolumeRequest{VolumeId: mockPVName})
	assert.Error(t, err)
}

func TestControllerServer_ValidateVolumeCapabilitiesConfirmed(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)

	req := &csi.ValidateVolumeCapabilitiesRequest{
		VolumeCapabilities: []*csi.VolumeCapability{
			{
				AccessMode: &csi.VolumeCapability_AccessMode{
					Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
				},
			},
		},
	}
	resp, err := cs.ValidateVolumeCapabilities(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.VolumeCapabilities, resp.Confirmed.VolumeCapabilities)
}

func TestControllerServer_ValidateVolumeCapabilitiesNotConfirmed(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)

	req := &csi.ValidateVolumeCapabilitiesRequest{
		VolumeCapabilities: []*csi.VolumeCapability{
			{
				AccessMode: &csi.VolumeCapability_AccessMode{
					Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY,
				},
			},
		},
	}
	resp, err := cs.ValidateVolumeCapabilities(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	mytesting.AssertProtoEqual(t, &csi.ValidateVolumeCapabilitiesResponse{}, resp)
}

func TestControllerServer_ControllerGetCapabilities(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)

	resp, err := cs.ControllerGetCapabilities(context.Background(), &csi.ControllerGetCapabilitiesRequest{})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestControllerServer_ControllerPublishVolume(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)
	resp, err := cs.ControllerPublishVolume(context.Background(), &csi.ControllerPublishVolumeRequest{})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.Unimplemented, ""), err)
}

func TestControllerServer_ControllerUnpublishVolume(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)
	resp, err := cs.ControllerUnpublishVolume(context.Background(), &csi.ControllerUnpublishVolumeRequest{})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.Unimplemented, ""), err)
}

func TestControllerServer_ListVolumes(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)
	resp, err := cs.ListVolumes(context.Background(), &csi.ListVolumesRequest{})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.Unimplemented, ""), err)
}

func TestControllerServer_GetCapacity(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)
	resp, err := cs.GetCapacity(context.Background(), &csi.GetCapacityRequest{})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.Unimplemented, ""), err)
}

func TestControllerServer_CreateSnapshot(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)
	resp, err := cs.CreateSnapshot(context.Background(), &csi.CreateSnapshotRequest{})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.Unimplemented, ""), err)
}

func TestControllerServer_DeleteSnapshot(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)
	resp, err := cs.DeleteSnapshot(context.Background(), &csi.DeleteSnapshotRequest{})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.Unimplemented, ""), err)
}

func TestControllerServer_ListSnapshots(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)
	resp, err := cs.ListSnapshots(context.Background(), &csi.ListSnapshotsRequest{})
	assert.Nil(t, resp)
	assert.Equal(t, status.Error(codes.Unimplemented, ""), err)
}
