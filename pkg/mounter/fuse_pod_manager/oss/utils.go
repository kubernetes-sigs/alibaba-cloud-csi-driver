package oss

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	fpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
)

var (
	defaultOssfsImageTag        = "v1.88.4-80d165c-aliyun"
	defaultOssfsUpdatedImageTag = "v1.91.10.ack.5-f5e325e"
	defaultOssfs2ImageTag       = "v2.0.7.ack.1-2d5bf24"
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
		case mounterutils.OssFsType:
			if features.FunctionalMutableFeatureGate.Enabled(features.UpdatedOssfsVersion) {
				config.ImageTag = defaultOssfsUpdatedImageTag
			} else {
				config.ImageTag = defaultOssfsImageTag
			}
		case mounterutils.OssFs2Type:
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

// parseRegionFromURL extracts region from known OSS endpoint patterns:
//   - oss-{region}[-internal].aliyuncs.com
//   - {region}[-internal].oss-data-acc.aliyuncs.com
//
// Same patterns as setNetworkType in pkg/oss/utils.go (duplicated due to circular dependency).
func parseRegionFromURL(rawURL string) string {
	var protocol string
	if strings.HasPrefix(rawURL, "https://") {
		protocol = "https://"
	} else if strings.HasPrefix(rawURL, "http://") {
		protocol = "http://"
	}
	endpoint := strings.TrimSuffix(strings.TrimPrefix(rawURL, protocol), "/")

	// {region}[-internal].oss-data-acc.aliyuncs.com — check first, more specific suffix
	if after, ok := strings.CutSuffix(endpoint, ".oss-data-acc.aliyuncs.com"); ok {
		return strings.TrimSuffix(after, "-internal")
	}

	// oss-{region}[-internal].aliyuncs.com
	if after, ok := strings.CutPrefix(endpoint, "oss-"); ok {
		if region, ok := strings.CutSuffix(after, ".aliyuncs.com"); ok {
			return strings.TrimSuffix(region, "-internal")
		}
	}

	return ""
}

// UseV4 returns true if the effective signature version is v4.
// Region being set implies v4 unless sigVersion is explicitly v1.
func UseV4(o *Options) bool {
	if o.SigVersion == SigV1 {
		return false
	}
	return o.SigVersion == SigV4 || o.Region != ""
}

// ResolveRegion resolves the OSS bucket region with the following priority:
//  1. User-specified o.Region
//  2. Parsed from o.URL
//  3. Node metadata (cluster region)
//
// When sigVersion=v4 is configured without an explicit region, a warning is logged.
// This is intentionally not an error for backward compatibility: existing users whose
// cluster and bucket are in the same region will continue to work without changes.
// However, cross-region access may fail silently in this case.
func ResolveRegion(o *Options, m metadata.MetadataProvider) string {
	if o.Region != "" {
		return o.Region
	}

	var region, source string

	if r := parseRegionFromURL(o.URL); r != "" {
		region = r
		source = "URL"
	} else if r, _ := m.Get(metadata.RegionID); r != "" {
		region = r
		source = "node metadata"
	}

	if o.SigVersion == SigV4 {
		if region != "" {
			klog.Warningf("sigVersion=v4 is configured without explicit region, using %q from %s; cross-region access may fail", region, source)
		} else {
			klog.Warning("sigVersion=v4 is configured but region cannot be determined; set region in volume parameters")
		}
	}

	return region
}

func GetAgentIdentityTokenFilePath(sandboxId string) string {
	return fmt.Sprintf("/var/opt/sandbox/agent-token/%s.token", sandboxId)
}

func GetAgentIdentityEndpoint() string {
	if ep := os.Getenv("AGENT_IDENTITY_ENDPOINT"); ep != "" {
		return ep
	}
	return "https://credential-provider.ack-agent-identity.svc:8443/"
}

// keep consistent with RAM response
var secretRefKeysToParse []string = []string{
	mounterutils.KeyAccessKeyId,
	mounterutils.KeyAccessKeySecret,
	mounterutils.KeyExpiration, // only used in ossfs1.0
	mounterutils.KeySecurityToken,
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
