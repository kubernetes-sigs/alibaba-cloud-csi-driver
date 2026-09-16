//go:build !windows

package customfuse

import (
	"context"
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateVolumePassesParametersThrough(t *testing.T) {
	cs := &controllerServer{}
	resp, err := cs.CreateVolume(context.Background(), &csi.CreateVolumeRequest{
		Name: "pvc-abc",
		Parameters: map[string]string{
			"source":           "redis://host:6379/1",
			"url":              "oss-cn-hangzhou-internal.aliyuncs.com",
			"entrypointConfig": "my-config",
		},
	})
	require.NoError(t, err)
	assert.Equal(t, "pvc-abc", resp.Volume.VolumeId)
	assert.Equal(t, "redis://host:6379/1", resp.Volume.VolumeContext["source"])
	assert.Equal(t, "my-config", resp.Volume.VolumeContext["entrypointConfig"])
}

// The requested size reaching the entrypoint as $capacity is the whole reason this
// RPC exists, since nothing else about the volume is provisioned.
func TestCreateVolumeTurnsRequestedSizeIntoCapacity(t *testing.T) {
	cs := &controllerServer{}
	resp, err := cs.CreateVolume(context.Background(), &csi.CreateVolumeRequest{
		Name:          "pvc-abc",
		CapacityRange: &csi.CapacityRange{RequiredBytes: 100 * 1024 * 1024 * 1024},
	})
	require.NoError(t, err)
	assert.Equal(t, "100Gi", resp.Volume.VolumeContext["capacity"])
	assert.EqualValues(t, 100*1024*1024*1024, resp.Volume.CapacityBytes)
}

func TestCreateVolumeWithoutCapacityRange(t *testing.T) {
	cs := &controllerServer{}
	resp, err := cs.CreateVolume(context.Background(), &csi.CreateVolumeRequest{Name: "pvc-abc"})
	require.NoError(t, err)
	_, set := resp.Volume.VolumeContext["capacity"]
	assert.False(t, set, "no requested size must leave capacity unset rather than pass 0 through")
}

// A capacity stated in the StorageClass is more specific than one derived from the
// PVC's size, so it wins.
func TestCreateVolumeKeepsExplicitCapacityParameter(t *testing.T) {
	cs := &controllerServer{}
	resp, err := cs.CreateVolume(context.Background(), &csi.CreateVolumeRequest{
		Name:          "pvc-abc",
		Parameters:    map[string]string{"capacity": "7Gi"},
		CapacityRange: &csi.CapacityRange{RequiredBytes: 100 * 1024 * 1024 * 1024},
	})
	require.NoError(t, err)
	assert.Equal(t, "7Gi", resp.Volume.VolumeContext["capacity"])
}

// Accepting Delete would promise a deprovisioning that never happens, leaving the
// customer's storage behind while Kubernetes reports the volume gone.
func TestCreateVolumeRejectsNonRetainReclaimPolicy(t *testing.T) {
	cs := &controllerServer{}
	key := common.CsiAlibabaCloudPrefix + "/reclaimPolicy"

	_, err := cs.CreateVolume(context.Background(), &csi.CreateVolumeRequest{
		Name:       "pvc-abc",
		Parameters: map[string]string{key: "Delete"},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Retain")

	_, err = cs.CreateVolume(context.Background(), &csi.CreateVolumeRequest{
		Name:       "pvc-abc",
		Parameters: map[string]string{key: "Retain"},
	})
	assert.NoError(t, err)
}

func TestDeleteVolumeSucceeds(t *testing.T) {
	cs := &controllerServer{}
	_, err := cs.DeleteVolume(context.Background(), &csi.DeleteVolumeRequest{VolumeId: "pvc-abc"})
	assert.NoError(t, err)
}

// The capability is advertised, so external-provisioner calls both RPCs; they have
// to be implemented rather than inherited as Unimplemented.
func TestControllerAdvertisesImplementedCapabilities(t *testing.T) {
	cs := &controllerServer{}
	resp, err := cs.ControllerGetCapabilities(context.Background(), nil)
	require.NoError(t, err)

	advertised := map[csi.ControllerServiceCapability_RPC_Type]bool{}
	for _, c := range resp.Capabilities {
		advertised[c.GetRpc().GetType()] = true
	}
	assert.True(t, advertised[csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME])
	assert.True(t, advertised[csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME])
}
