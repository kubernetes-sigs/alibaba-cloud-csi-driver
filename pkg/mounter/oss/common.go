package oss

import (
	"fmt"
	"os"
	"path"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
)

const (
	// OssFsType is the oss filesystem type
	OssFsType = "ossfs"
	// OssFs2Type is the ossfs2 filesystem type
	OssFs2Type = "ossfs2"
)

// keys for STS token
const (
	KeyAccessKeyId     = "AccessKeyId"
	KeyAccessKeySecret = "AccessKeySecret"
	KeyExpiration      = "Expiration"
	KeySecurityToken   = "SecurityToken"
)

func setDefaultImage(fuseType string, m metadata.MetadataProvider, config *mounterutils.FuseContainerConfig) {
	// set recovery image
	// TODO: remove this after recovery capacity reaches beta status.
	if features.FunctionalMutableFeatureGate.Enabled(features.EnableOssfsRecovery) {
		switch fuseType {
		case OssFsType:
			config.RecoveryImage = fmt.Sprintf("csi-%s:%s", fuseType, defaultOssfsRecoveryImageTag)
		case OssFs2Type:
			config.RecoveryImage = fmt.Sprintf("csi-%s:%s", fuseType, defaultOssfs2RecoveryImageTag)
		}
	}

	// deprecated
	if config.Image != "" {
		return
	}
	region, err := m.Get(metadata.RegionID)
	if err != nil {
		klog.Warningf("Failed to get region from metadata: %v", err)
	}
	prefix := utils.GetRepositoryPrefix(region)

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
	image := fmt.Sprintf("csi-%s:%s", fuseType, config.ImageTag)
	config.Image = path.Join(prefix, image)
	klog.Infof("Use ossfs image: %s", config.Image)
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
func getRRSAConfig(o *Options, m metadata.MetadataProvider) (rrsaCfg *mounterutils.RrsaConfig, err error) {
	saName := mounterutils.FuseServiceAccountName
	if o.ServiceAccountName != "" {
		saName = o.ServiceAccountName
	}

	if o.OidcProviderArn != "" && o.RoleArn != "" {
		return &mounterutils.RrsaConfig{
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
	provider := mounterutils.GetOIDCProvider(clusterId)
	oidcProviderArn, roleArn := mounterutils.GetArn(provider, accountId, o.RoleName)
	return &mounterutils.RrsaConfig{
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

// keep consistent with RAM response
var secretRefKeysToParse []string = []string{
	KeyAccessKeyId,
	KeyAccessKeySecret,
	KeyExpiration, // only used in ossfs1.0
	KeySecurityToken,
}

func getPasswdSecretVolume(secretRef, fuseType string) (secret *corev1.SecretVolumeSource) {
	if secretRef == "" {
		return nil
	}
	items := []corev1.KeyToPath{}
	for _, key := range secretRefKeysToParse {
		item := corev1.KeyToPath{
			Key:  key,
			Path: fmt.Sprintf("%s/%s", mounterutils.GetPasswdFileName(fuseType), key),
			Mode: tea.Int32(0600),
		}
		items = append(items, item)
	}
	secret = &corev1.SecretVolumeSource{
		SecretName: secretRef,
		Items:      items,
	}
	return
}
