//go:build !windows

package nas

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestControllerServer_ValidateVolumeCapabilitiesReadOnly(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)

	req := &csi.ValidateVolumeCapabilitiesRequest{
		VolumeCapabilities: []*csi.VolumeCapability{
			{
				AccessType: &csi.VolumeCapability_Mount{
					Mount: &csi.VolumeCapability_MountVolume{},
				},
				AccessMode: &csi.VolumeCapability_AccessMode{
					Mode: csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY,
				},
			},
		},
	}
	resp, err := cs.ValidateVolumeCapabilities(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.VolumeCapabilities, resp.Confirmed.VolumeCapabilities)
}

func TestControllerServer_ValidateVolumeCapabilitiesBlockVolume(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)

	req := &csi.ValidateVolumeCapabilitiesRequest{
		VolumeCapabilities: []*csi.VolumeCapability{
			{
				AccessType: &csi.VolumeCapability_Block{
					Block: &csi.VolumeCapability_BlockVolume{},
				},
				AccessMode: &csi.VolumeCapability_AccessMode{
					Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
				},
			},
		},
	}
	resp, err := cs.ValidateVolumeCapabilities(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, resp.Confirmed)
	assert.Contains(t, resp.Message, "block")
}

func TestControllerServer_ValidateVolumeCapabilitiesUnsupportedMode(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)

	req := &csi.ValidateVolumeCapabilitiesRequest{
		VolumeCapabilities: []*csi.VolumeCapability{
			{
				AccessMode: &csi.VolumeCapability_AccessMode{
					Mode: csi.VolumeCapability_AccessMode_Mode(99),
				},
			},
		},
	}
	resp, err := cs.ValidateVolumeCapabilities(context.Background(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, resp.Confirmed)
	assert.Contains(t, resp.Message, "unsupported access mode")
}

func TestControllerServer_ControllerGetCapabilities(t *testing.T) {
	cs := newMockControllerServer()
	assert.NotNil(t, cs)

	resp, err := cs.ControllerGetCapabilities(context.Background(), &csi.ControllerGetCapabilitiesRequest{})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

// TestControllerServerCreateVolumeAgenticFsKeepsForcedMountOptions is the end-to-end half of W-1.
// It drives the REAL generic wrapper in controllerserver.go with a StorageClass that sets
// parameters.options, and asserts the mandatory tls/ram survive the verbatim overwrite. The wrapper
// picks the agenticfs controller by volumeAs, calls its CreateVolume (which seeds options with
// defaultAgenticFsMountOptions), overwrites options with the StorageClass value, and THEN calls
// enforceAgenticFsMountOptions. Removing that hook - or moving it before the overwrite - makes the
// final VolumeContext lose tls and ram, so this test goes red.
func TestControllerServerCreateVolumeAgenticFsKeepsForcedMountOptions(t *testing.T) {
	ctrl := newAgenticfsCtrl(t, newFakeNasClientV2())
	cs := &controllerServer{
		ControllerFactory: &internal.ControllerFactory{
			Modes: map[string]internal.Controller{agenticFsVolumeAs: ctrl},
		},
		kubeClient: fake.NewSimpleClientset(),
		locks:      utils.NewVolumeLocks(),
	}

	// A StorageClass that sets options the agenticfs defaults do not include, exactly the case that
	// used to drop tls/ram.
	req := agenticfsCreateReq(testAgenticFsPVName, 20*GiB, map[string]string{"options": "vers=3,noresvport"})
	resp, err := cs.CreateVolume(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)

	opts := resp.Volume.VolumeContext["options"]
	require.NotEmpty(t, opts, "the final VolumeContext must carry options")
	for _, want := range []string{"tls", "ram", "vers=3", "noresvport"} {
		assert.Contains(t, opts, want,
			"W-1: the StorageClass options overwrite must not drop a mandatory agenticfs mount option")
	}
}
