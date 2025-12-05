package oss

import (
	"fmt"
	"os"
	"path"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	fpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager"
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

var (
	defaultOssfsImageTag        = "v1.88.4-80d165c-aliyun"
	defaultOssfsUpdatedImageTag = "v1.91.8.ack.3-b0e4403"
	defaultOssfs2ImageTag       = "v2.0.4.ack.1-5073ed2"
)

// keys for STS token
const (
	KeyAccessKeyId     = "AccessKeyId"
	KeyAccessKeySecret = "AccessKeySecret"
	KeyExpiration      = "Expiration"
	KeySecurityToken   = "SecurityToken"
)

func SetDefaultImage(fuseType string, m metadata.MetadataProvider, config *fpm.FuseContainerConfig) {
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

// CheckRRSAParams check parameters of RRSA
func CheckRRSAParams(o *Options) error {

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

// GetRRSAConfig get oidcProviderArn and roleArn
func GetRRSAConfig(o *Options, m metadata.MetadataProvider) (rrsaCfg *fpm.RrsaConfig, err error) {
	saName := fpm.FuseServiceAccountName
	if o.ServiceAccountName != "" {
		saName = o.ServiceAccountName
	}

	if o.OidcProviderArn != "" && o.RoleArn != "" {
		return &fpm.RrsaConfig{
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
	return &fpm.RrsaConfig{
		OidcProviderArn:    oidcProviderArn,
		RoleArn:            roleArn,
		ServiceAccountName: saName,
		AssumeRoleArn:      o.AssumeRoleArn,
	}, nil
}

// GetSTSEndpoint get STS endpoint
func GetSTSEndpoint(region string) string {

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

func GetPasswdSecretVolume(secretRef, fuseType string) (secret *corev1.SecretVolumeSource) {
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

func AppendRRSAAuthOptions(m metadata.MetadataProvider, options []string, volumeId, target string, authCfg *fpm.AuthConfig) ([]string, error) {
	if authCfg == nil {
		return options, nil
	}

	if authCfg.AuthType == "rrsa" {
		tokenFile, err := m.Get(metadata.RRSATokenFile)
		if err != nil {
			return nil, err
		}
		sessionName := mounterutils.GetRoleSessionName(volumeId, target, "ossfs")
		options = append(options, fmt.Sprintf("rrsa_oidc_provider_arn=%s", authCfg.RrsaConfig.OidcProviderArn))
		options = append(options, fmt.Sprintf("rrsa_role_arn=%s", authCfg.RrsaConfig.RoleArn))
		options = append(options, fmt.Sprintf("rrsa_role_session_name=%s", sessionName))
		options = append(options, fmt.Sprintf("rrsa_token_file=%s", tokenFile))
	}
	return options, nil
}
