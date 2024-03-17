package ens

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	csicommon "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/agent/csi-common"
	log "github.com/sirupsen/logrus"
)

type identityServer struct {
	*csicommon.DefaultIdentityServer
}

// NewIdentityServer create identity server
func NewIdentityServer(d *csicommon.CSIDriver) csi.IdentityServer {
	return &identityServer{
		csicommon.NewDefaultIdentityServer(d),
	}
}

// GetPluginCapabilities returns available capabilities of the plugin
func (iden *identityServer) GetPluginCapabilities(ctx context.Context, req *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	log.Infof("Identity:GetPluginCapabilities is called")
	resp := &csi.GetPluginCapabilitiesResponse{
		Capabilities: []*csi.PluginCapability{
			{
				Type: &csi.PluginCapability_Service_{
					Service: &csi.PluginCapability_Service{
						Type: csi.PluginCapability_Service_CONTROLLER_SERVICE,
					},
				},
			},
			{
				Type: &csi.PluginCapability_Service_{
					Service: &csi.PluginCapability_Service{
						Type: csi.PluginCapability_Service_VOLUME_ACCESSIBILITY_CONSTRAINTS,
					},
				},
			},
			{
				Type: &csi.PluginCapability_VolumeExpansion_{
					VolumeExpansion: &csi.PluginCapability_VolumeExpansion{
						Type: csi.PluginCapability_VolumeExpansion_OFFLINE,
					},
				},
			},
		},
	}
	return resp, nil
}
