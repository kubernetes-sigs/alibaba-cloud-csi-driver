package fuse

import (
	"fmt"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/klog/v2"
)

const (
	// AkID is Ak ID
	AkID = "akId"
	// AkSecret is Ak Secret
	AkSecret = "akSecret"
	// OssFsType is the oss filesystem type
	OssFsType = "ossfs"
	// OssFs2Type is the ossfs2 filesystem type
	OssFs2Type = "ossfs2"
)

func setDefaultImage(fuseType string, m metadata.MetadataProvider, config *mounterutils.FuseContainerConfig) {
	if config.Image == "" {
		registry, _ := m.Get(metadata.RegistryURL)
		if registry == "" {
			region, err := m.Get(metadata.RegionID)
			if err == nil {
				registry = fmt.Sprintf("registry-%s-vpc.ack.aliyuncs.com", region)
			} else {
				klog.Warningf("DEFAULT_REGISTRY env not set, failed to get current region: %v, fallback to default registry: %s", err, defaultRegistry)
				registry = defaultRegistry
			}
		}
		if config.ImageTag == "" {
			switch fuseType {
			case OssFsType:
				if features.FunctionalMutableFeatureGate.Enabled(features.UpdatedOssfsVersion) {
					config.ImageTag = defaultOssfsUpdatedImageTag
				} else {
					config.ImageTag = defaultOssfsImageTag
				}
			case OssFs2Type:
				config.ImageTag = defaultOssfs2ImageTag
			default:
				klog.Warningf("Unknown fuse type: %s", fuseType)
				config.ImageTag = "latest"
			}

		}
		config.Image = fmt.Sprintf("%s/acs/csi-%s:%s", registry, fuseType, config.ImageTag)
		klog.Infof("Use ossfs image: %s", config.Image)
	}
}

// checkRRSAParams check parameters of RRSA
func checkRRSAParams(c *utils.AuthConfig) error {

	rc := c.RrsaConfig
	if rc != nil {
		if rc.RoleArn != "" && rc.OidcProviderArn != "" {
			return nil
		}
		if rc.RoleArn != "" || rc.OidcProviderArn != "" {
			return fmt.Errorf("use RRSA but one of the ARNs is empty, roleArn: %s, oidcProviderArn: %s", rc.RoleArn, rc.OidcProviderArn)
		}
	}
	if c.RoleName == "" {
		return fmt.Errorf("use RRSA but roleName is empty")
	}

	return nil
}
