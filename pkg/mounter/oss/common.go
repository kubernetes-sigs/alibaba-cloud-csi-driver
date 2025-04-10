package oss

import (
	"fmt"
	"os"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/klog/v2"
)

const (
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
			if region != "" {
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
func checkRRSAParams(o *Options) error {

	if o.RoleArn != "" && o.OidcProviderArn != "" {
		return nil
	}
	if o.RoleArn != "" || o.OidcProviderArn != "" {
		return fmt.Errorf("use RRSA but one of the ARNs is empty, roleArn: %s, oidcProviderArn: %s", o.RoleArn, o.OidcProviderArn)
	}

	if o.RoleName == "" {
		return fmt.Errorf("use RRSA but roleName is empty")
	}

	return nil
}

// getRRSAConfig get oidcProviderArn and roleArn
func getRRSAConfig(o *Options, m metadata.MetadataProvider) (rrsaCfg *utils.RrsaConfig, err error) {
	saName := utils.FuseServiceAccountName
	if o.ServiceAccountName != "" {
		saName = o.ServiceAccountName
	}

	if o.OidcProviderArn != "" && o.RoleArn != "" {
		return &utils.RrsaConfig{
			OidcProviderArn:    o.OidcProviderArn,
			RoleArn:            o.RoleArn,
			ServiceAccountName: saName,
			AssumeRoleArn:      o.AssumeRoleArn,
		}, nil
	}

	accountId, err := m.Get(metadata.AccountID)
	if err != nil {
		return nil, fmt.Errorf("Get accountId error: %v", err)
	}
	clusterId, err := m.Get(metadata.ClusterID)
	if err != nil {
		return nil, fmt.Errorf("Get clusterId error: %v", err)
	}
	provider := utils.GetOIDCProvider(clusterId)
	oidcProviderArn, roleArn := utils.GetArn(provider, accountId, o.RoleName)
	return &utils.RrsaConfig{
		OidcProviderArn:    oidcProviderArn,
		RoleArn:            roleArn,
		ServiceAccountName: saName,
		AssumeRoleArn:      o.AssumeRoleArn,
	}, nil
}

// getSTSEndpoint get STS endpoint
func getSTSEndpoint(region string) string {

	// for PrivateCloud
	if os.Getenv("STS_ENDPOINT") != "" {
		return os.Getenv("STS_ENDPOINT")
	}

	if region == "" {
		return "https://sts.aliyuncs.com"
	}
	return fmt.Sprintf("https://sts-vpc.%s.aliyuncs.com", region)
}
