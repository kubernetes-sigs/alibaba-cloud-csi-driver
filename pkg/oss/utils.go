/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package oss

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	cnfsv1beta1 "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

// SigVersion is the version of signature for ossfs
type SigVersion string

const (
	// default version for ossfs
	SigV1 SigVersion = "v1"
	// need set region
	SigV4 SigVersion = "v4"
)

// VolumeAs determines the mounting target path in OSS
type VolumeAsType string

const (
	// VolumeAsDirect means mount on the `path` directly
	VolumeAsDirect VolumeAsType = "direct"
	// VolumeAsSharePath is same with `direct`, keep compatible with CNFS-NAS
	VolumeAsSharePath VolumeAsType = "sharepath"
	// VolumeAsSubpath means mount on the subpath of `path` which is automatically generated
	VolumeAsSubpath VolumeAsType = "subpath"
)

type AccessKey struct {
	AkID     string `json:"akId"`
	AkSecret string `json:"akSecret"`
}
type TokenSecret struct {
	AccessKeyId     string `json:"AccessKeyId"`
	AccessKeySecret string `json:"AccessKeySecret"`
	Expiration      string `json:"Expiration"`
	SecurityToken   string `json:"SecurityToken"`
}

// Options contains options for target oss
type Options struct {
	directAssigned bool
	CNFSName       string

	// oss options
	Bucket string `json:"bucket"`
	URL    string `json:"url"`
	Path   string `json:"path"`

	// authorization options
	// accesskey
	AccessKey   `json:",inline"`
	TokenSecret `json:",inline"`
	SecretRef   string `json:"secretRef"`

	// RRSA
	RoleName           string `json:"roleName"` // also for STS
	RoleArn            string `json:"roleArn"`
	OidcProviderArn    string `json:"oidcProviderArn"`
	ServiceAccountName string `json:"serviceAccountName"`
	// assume role
	AssumeRoleArn string `json:"assumeRoleArn"`
	ExternalId    string `json:"externalId"`
	// csi secret store
	SecretProviderClass string `json:"secretProviderClass"`

	// ossfs options
	OtherOpts  string     `json:"otherOpts"`
	MetricsTop string     `json:"metricsTop"`
	Encrypted  string     `json:"encrypted"`
	KmsKeyId   string     `json:"kmsKeyId"`
	SigVersion SigVersion `json:"sigVersion"`

	// mount options
	UseSharedPath bool   `json:"useSharedPath"`
	AuthType      string `json:"authType"`
	FuseType      string `json:"fuseType"`
	ReadOnly      bool   `json:"readOnly"`
}

// get Options for CreateVolume and PublishVolume
// volOptions: CreateVolumeRequest.GetParameters, PublishVolumeRequest.GetVolumeContext
// secrets: CreateVolumeRequest / PublishVolumeRequest.GetSecrets
// volCaps: CreateVolumeRequest.GetVolumeCapabilities, []{PublishVolumeRequest.GetVolumeCapability}
// readOnly: PublishVolumeRequest.GetReadonly
// region: for signature v4
// reqName: for subpath generating, CreateVolumeRequest.GetName
func parseOptions(volOptions, secrets map[string]string, volCaps []*csi.VolumeCapability,
	readOnly bool, region, reqName string) *Options {

	if volOptions == nil {
		volOptions = map[string]string{}
	}
	if secrets == nil {
		secrets = map[string]string{}
	}

	opts := &Options{
		UseSharedPath: true,
		FuseType:      OssFsType,
		Path:          "/",
		AccessKey: AccessKey{
			AkID:     strings.TrimSpace(secrets[AkID]),
			AkSecret: strings.TrimSpace(secrets[AkSecret]),
		},
		TokenSecret: TokenSecret{
			AccessKeyId:     strings.TrimSpace(secrets[mounter.KeyAccessKeyId]),
			AccessKeySecret: strings.TrimSpace(secrets[mounter.KeyAccessKeySecret]),
			Expiration:      strings.TrimSpace(secrets[mounter.KeyExpiration]),
			SecurityToken:   strings.TrimSpace(secrets[mounter.KeySecurityToken]),
		},
		MetricsTop: defaultMetricsTop,
	}

	var volumeAsSubpath bool
	for k, v := range volOptions {
		key := strings.TrimSpace(strings.ToLower(k))
		value := strings.TrimSpace(v)
		if value == "" {
			continue
		}
		switch key {
		case "bucket":
			opts.Bucket = value
		case "url":
			opts.URL = value
		case "otheropts":
			opts.OtherOpts = value
		case "akid":
			opts.AkID = value
		case "aksecret":
			opts.AkSecret = value
		case "secretref":
			opts.SecretRef = value
		case "path":
			opts.Path = value
		case "usesharedpath":
			if res, err := strconv.ParseBool(value); err == nil {
				opts.UseSharedPath = res
			} else {
				klog.Warning(WrapOssError(ParamError, "the value(%q) of %q is invalid", v, k).Error())
			}
		case "authtype":
			opts.AuthType = strings.ToLower(value)
		case "rolename", "ramrole":
			opts.RoleName = value
		case "rolearn":
			opts.RoleArn = value
		case "oidcproviderarn":
			opts.OidcProviderArn = value
		case "serviceaccountname":
			opts.ServiceAccountName = value
		case "secretproviderclass":
			opts.SecretProviderClass = value
		case "fusetype":
			opts.FuseType = strings.ToLower(value)
		case "metricstop":
			opts.MetricsTop = strings.ToLower(value)
		case "containernetworkfilesystem":
			opts.CNFSName = value
		case optDirectAssigned:
			if res, err := strconv.ParseBool(value); err == nil {
				opts.directAssigned = res
			} else {
				klog.Warning(WrapOssError(ParamError, "the value(%q) of %q is invalid", v, k).Error())
			}
		case "encrypted":
			opts.Encrypted = strings.ToLower(value)
		case "kmskeyid":
			opts.KmsKeyId = value
		case "assumerolearn":
			opts.AssumeRoleArn = value
		case "externalid":
			opts.ExternalId = value
		case "sigversion":
			switch SigVersion(value) {
			case SigV1, SigV4:
				opts.SigVersion = SigVersion(value)
			default:
				klog.Warning(WrapOssError(ParamError, "the value(%q) of %q is invalid, only support v1 and v4", v, k).Error())
			}
		case "volumeas":
			switch VolumeAsType(value) {
			case VolumeAsDirect, VolumeAsSharePath:
				// do nothing
			case VolumeAsSubpath:
				volumeAsSubpath = true
			default:
				klog.Warning(WrapOssError(ParamError, "the value(%q) of %q is invalid, only support direct and subpath", v, k).Error())
			}
		}
	}
	for _, c := range volCaps {
		switch c.AccessMode.GetMode() {
		case csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY, csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY:
			opts.ReadOnly = true
		}
	}
	if readOnly {
		opts.ReadOnly = true
	}

	url := opts.URL
	if region != "" {
		url, _ = setNetworkType(url, region)
	}
	url, _ = setTransmissionProtocol(url)
	if url != opts.URL {
		klog.Infof("Changed oss URL from %s to %s", opts.URL, url)
		opts.URL = url
	}

	// only work for CreateVolume
	if volumeAsSubpath {
		opts.Path = path.Join(opts.Path, reqName)
	}

	return opts
}

func setCNFSOptions(ctx context.Context, cnfsGetter cnfsv1beta1.CNFSGetter, opts *Options) error {
	if opts.CNFSName == "" {
		return nil
	}
	cnfs, err := cnfsGetter.GetCNFS(ctx, opts.CNFSName)
	if err != nil {
		return err
	}
	if cnfs.Status.FsAttributes.EndPoint == nil {
		return fmt.Errorf("missing endpoint in status of CNFS %s", opts.CNFSName)
	}
	opts.Bucket = cnfs.Status.FsAttributes.BucketName
	opts.URL = cnfs.Status.FsAttributes.EndPoint.Internal
	return nil
}

// checkRRSAParams check parameters of RRSA
func checkRRSAParams(opt *Options) error {
	if opt.RoleArn != "" && opt.OidcProviderArn != "" {
		return nil
	}
	if opt.RoleArn != "" || opt.OidcProviderArn != "" {
		return fmt.Errorf("use RRSA but one of the ARNs is empty, roleArn: %s, oidcProviderArn: %s", opt.RoleArn, opt.OidcProviderArn)
	}
	if opt.RoleName == "" {
		return fmt.Errorf("use RRSA but roleName is empty")
	}
	return nil
}

func checkDefaultAuthOptions(opt *Options) error {
	if opt == nil {
		return nil
	}
	// Using fixed AKSK for authentication first if set.
	if opt.AkID != "" && opt.AkSecret != "" {
		return nil
	}

	// If AKSK if not set, check if a Secret maintained externally is configured,
	//   which stores Token authentication information.
	// For runc scenarios, set the SecretRef parameter.
	runc := opt.SecretRef != ""
	// For rund or eci scenarios, configure Token in nodePublishSecretRef or nodeStageSecretRef.
	rund := opt.AccessKeyId != "" && opt.AccessKeySecret != "" && opt.Expiration != "" && opt.SecurityToken != ""
	if runc && rund {
		return fmt.Errorf("Token and secretRef cannot be set at the same time")
	}
	if !features.FunctionalMutableFeatureGate.Enabled(features.RundCSIProtocol3) {
		if !runc && !rund {
			return fmt.Errorf("missing authentication information")
		}
	}
	if opt.SecretRef == mounter.OssfsCredentialSecretName {
		return fmt.Errorf("invalid secretRef name")
	}

	return nil
}

// getRRSAConfig get oidcProviderArn and roleArn
func getRRSAConfig(opt *Options, m metadata.MetadataProvider) (rrsaCfg *mounter.RrsaConfig, err error) {
	saName := fuseServiceAccountName
	if opt.ServiceAccountName != "" {
		saName = opt.ServiceAccountName
	}

	if opt.OidcProviderArn != "" && opt.RoleArn != "" {
		return &mounter.RrsaConfig{ServiceAccountName: saName, OidcProviderArn: opt.OidcProviderArn, RoleArn: opt.RoleArn}, nil
	}

	accountId, err := m.Get(metadata.AccountID)
	if err != nil {
		return nil, fmt.Errorf("Get accountId error: %v", err)
	}
	clusterId, err := m.Get(metadata.ClusterID)
	if err != nil {
		return nil, fmt.Errorf("Get clusterId error: %v", err)
	}
	provider := mounter.GetOIDCProvider(clusterId)
	oidcProviderArn, roleArn := mounter.GetArn(provider, accountId, opt.RoleName)
	return &mounter.RrsaConfig{ServiceAccountName: saName, OidcProviderArn: oidcProviderArn, RoleArn: roleArn}, nil
}

// getDefaultAuthConfig gets accesskey with or without token
func getDefaultAuthConfig(opt *Options) (authCfg *mounter.AuthConfig, mountOptions []string) {
	authCfg = &mounter.AuthConfig{}
	// fixed AKSK
	if opt.AkID != "" && opt.AkSecret != "" {
		authCfg.Secrets = map[string]string{
			mounter.OssfsPasswdFile: fmt.Sprintf("%s:%s:%s", opt.Bucket, opt.AkID, opt.AkSecret),
		}
		return
	}

	// secretRef for RunC
	if opt.SecretRef != "" {
		authCfg.SecretRef = opt.SecretRef
		mountOptions = append(mountOptions, fmt.Sprintf("passwd_file=%s", filepath.Join(mounter.PasswdMountDir, mounter.PasswdFilename)))
		mountOptions = append(mountOptions, "use_session_token")
		return
	}

	// token secret for RunD
	authCfg.Secrets = map[string]string{
		filepath.Join(mounter.OssfsPasswdFile, mounter.KeyAccessKeyId):     opt.AccessKeyId,
		filepath.Join(mounter.OssfsPasswdFile, mounter.KeyAccessKeySecret): opt.AccessKeySecret,
		filepath.Join(mounter.OssfsPasswdFile, mounter.KeySecurityToken):   opt.SecurityToken,
		filepath.Join(mounter.OssfsPasswdFile, mounter.KeyExpiration):      opt.Expiration,
	}
	mountOptions = append(mountOptions, "use_session_token")
	return
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

// parseOtherOpts extracts mount options from parameters.otherOpts string.
// example: "-o max_stat_cache_size=0 -o allow_other -ogid=1000,uid=1000" => {"max_stat_cache_size=0", "allow_other", "uid=1000", "gid=1000"}
func parseOtherOpts(otherOpts string) (mountOptions []string, err error) {
	elements := strings.Fields(otherOpts)
	accepting := false
	for _, ele := range elements {
		if accepting {
			eles := strings.Split(ele, ",")
			mountOptions = append(mountOptions, eles...)
			accepting = false
		} else {
			if ele == "-o" {
				accepting = true
			} else if strings.HasPrefix(ele, "-o") {
				eles := strings.Split(strings.TrimPrefix(ele, "-o"), ",")
				mountOptions = append(mountOptions, eles...)
			} else {
				// missing -o
				return nil, status.Errorf(codes.InvalidArgument, "invalid otherOpts: %q", otherOpts)
			}
		}
	}
	return mountOptions, nil
}

func isNotMountPoint(mounter mountutils.Interface, target string) (notMnt bool, err error) {
	notMnt, err = mounter.IsLikelyNotMountPoint(target)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(target, os.ModePerm); err != nil {
				return false, status.Errorf(codes.Internal, "mkdir: %v", err)
			}
			return true, nil
		} else if mountutils.IsCorruptedMnt(err) {
			klog.Warningf("Umount corrupted mountpoint %s", target)
			err := mounter.Unmount(target)
			if err != nil {
				return false, status.Errorf(codes.Internal, "umount corrupted mountpoint %s: %v", target, err)
			}
			return true, nil
		}
		return false, status.Errorf(codes.Internal, "check mountpoint: %v", err)
	}
	return notMnt, nil
}

func setNetworkType(originURL, regionID string) (URL string, modified bool) {
	URL = originURL
	if utils.IsPrivateCloud() {
		return
	}
	var protocol string
	if strings.HasPrefix(originURL, "https://") {
		protocol = "https://"
	} else if strings.HasPrefix(originURL, "http://") {
		protocol = "http://"
	}
	endpoint := strings.TrimPrefix(originURL, protocol)

	switch endpoint {
	case fmt.Sprintf("oss-%s.aliyuncs.com", regionID):
		endpoint = fmt.Sprintf("oss-%s-internal.aliyuncs.com", regionID)
	case fmt.Sprintf("%s.oss-data-acc.aliyuncs.com", regionID):
		endpoint = fmt.Sprintf("%s-internal.oss-data-acc.aliyuncs.com", regionID)
	default:
		return
	}
	URL = protocol + endpoint
	modified = true
	return
}

func setTransmissionProtocol(originURL string) (URL string, modified bool) {
	URL = originURL
	// for PrivateCloud, use the default protocol (http) of ossfs if not set.
	if utils.IsPrivateCloud() || strings.HasPrefix(URL, "http://") || strings.HasPrefix(URL, "https://") {
		return
	}
	if strings.HasSuffix(strings.TrimRight(URL, "/"), "-internal.aliyuncs.com") {
		return "http://" + URL, true
	}
	return "https://" + URL, true
}

// validateEndpoint will try to validate endpoint：
//
//	ban the bucket domain name like: bucket.oss-cn-beijing.aliyuncs.com
func validateEndpoint(originURL, bucket string) error {
	if utils.IsPrivateCloud() {
		return nil
	}
	if len(strings.Split(originURL, ".")) < 4 {
		return nil
	}
	woProtocol := strings.TrimPrefix(strings.TrimPrefix(originURL, "http://"), "https://")
	if strings.HasPrefix(woProtocol, bucket+".") {
		return fmt.Errorf("%s is a bucket domain name, please use endpoint instead", originURL)
	}
	return nil
}
