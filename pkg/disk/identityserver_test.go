package disk

import (
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	featuregatetesting "k8s.io/component-base/featuregate/testing"
)

func TestGetPluginCapabilities(t *testing.T) {
	tests := []struct {
		name                       string
		enableVolumeGroupSnapshots bool
		expectedCapCount           int
		expectGroupControllerCap   bool
	}{
		{
			name:                       "volume group snapshots disabled",
			enableVolumeGroupSnapshots: false,
			expectedCapCount:           3, // CONTROLLER_SERVICE, VOLUME_ACCESSIBILITY_CONSTRAINTS, OFFLINE expansion
			expectGroupControllerCap:   false,
		},
		{
			name:                       "volume group snapshots enabled",
			enableVolumeGroupSnapshots: true,
			expectedCapCount:           4, // above + GROUP_CONTROLLER_SERVICE
			expectGroupControllerCap:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			featuregatetesting.SetFeatureGateDuringTest(t, features.FunctionalMutableFeatureGate, features.EnableVolumeGroupSnapshots, tt.enableVolumeGroupSnapshots)

			iden := NewIdentityServer()

			resp, err := iden.GetPluginCapabilities(t.Context(), &csi.GetPluginCapabilitiesRequest{})

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Len(t, resp.Capabilities, tt.expectedCapCount)

			// Verify expected capabilities
			hasControllerService := false
			hasVolumeAccessibility := false
			hasOfflineExpansion := false
			hasGroupController := false

			for _, cap := range resp.Capabilities {
				if svc := cap.GetService(); svc != nil {
					switch svc.Type {
					case csi.PluginCapability_Service_CONTROLLER_SERVICE:
						hasControllerService = true
					case csi.PluginCapability_Service_VOLUME_ACCESSIBILITY_CONSTRAINTS:
						hasVolumeAccessibility = true
					case csi.PluginCapability_Service_GROUP_CONTROLLER_SERVICE:
						hasGroupController = true
					}
				}
				if exp := cap.GetVolumeExpansion(); exp != nil {
					if exp.Type == csi.PluginCapability_VolumeExpansion_OFFLINE {
						hasOfflineExpansion = true
					}
				}
			}

			assert.True(t, hasControllerService, "should have CONTROLLER_SERVICE capability")
			assert.True(t, hasVolumeAccessibility, "should have VOLUME_ACCESSIBILITY_CONSTRAINTS capability")
			assert.True(t, hasOfflineExpansion, "should have OFFLINE volume expansion capability")
			assert.Equal(t, tt.expectGroupControllerCap, hasGroupController, "GROUP_CONTROLLER_SERVICE capability mismatch")
		})
	}
}
