package disk

import (
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGroupControllerGetCapabilities(t *testing.T) {
	cs := NewGroupControllerServer()

	resp, err := cs.GroupControllerGetCapabilities(t.Context(), &csi.GroupControllerGetCapabilitiesRequest{})

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotEmpty(t, resp.Capabilities)

	// Verify CREATE_DELETE_GET_VOLUME_GROUP_SNAPSHOT capability is present
	hasVolumeGroupSnapshotCap := false
	for _, cap := range resp.Capabilities {
		if rpc := cap.GetRpc(); rpc != nil {
			if rpc.Type == csi.GroupControllerServiceCapability_RPC_CREATE_DELETE_GET_VOLUME_GROUP_SNAPSHOT {
				hasVolumeGroupSnapshotCap = true
			}
		}
	}
	assert.True(t, hasVolumeGroupSnapshotCap, "should have CREATE_DELETE_GET_VOLUME_GROUP_SNAPSHOT capability")
}
