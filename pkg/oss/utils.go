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
	"strconv"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	cnfsv1beta1 "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/oss"
	mounter "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
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

// get Options for CreateVolume and PublishVolume
// volOptions: CreateVolumeRequest.GetParameters, PublishVolumeRequest.GetVolumeContext
// secrets: CreateVolumeRequest / PublishVolumeRequest.GetSecrets
// volCaps: CreateVolumeRequest.GetVolumeCapabilities, []{PublishVolumeRequest.GetVolumeCapability}
// readOnly: PublishVolumeRequest.GetReadonly
// region: for signature v4
// reqName: for subpath generating, CreateVolumeRequest.GetName
// onNode: run on nodeserver mode
func parseOptions(volOptions, secrets map[string]string, volCaps []*csi.VolumeCapability,
	readOnly bool, reqName string, onNode bool, m metadata.MetadataProvider) *oss.Options {

	if volOptions == nil {
		volOptions = map[string]string{}
	}
	if secrets == nil {
		secrets = map[string]string{}
	}

	opts := &oss.Options{
		UseSharedPath: true,
		Path:          "/",
		AccessKey: oss.AccessKey{
			AkID:     strings.TrimSpace(secrets[AkID]),
			AkSecret: strings.TrimSpace(secrets[AkSecret]),
		},
		TokenSecret: oss.TokenSecret{
			AccessKeyId:     strings.TrimSpace(secrets[KeyAccessKeyId]),
			AccessKeySecret: strings.TrimSpace(secrets[KeyAccessKeySecret]),
			Expiration:      strings.TrimSpace(secrets[KeyExpiration]),
			SecurityToken:   strings.TrimSpace(secrets[KeySecurityToken]),
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
				opts.DirectAssigned = res
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
			switch oss.SigVersion(value) {
			case oss.SigV1, oss.SigV4:
				opts.SigVersion = oss.SigVersion(value)
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
		// set fuseType by fsType
		if opts.FuseType == "" && c.GetMount() != nil && c.GetMount().FsType != "" {
			opts.FuseType = c.GetMount().FsType
		}
	}
	if readOnly {
		opts.ReadOnly = true
	}

	// default fuseType is ossfs
	if opts.FuseType == "" {
		opts.FuseType = OssFsType
	}

	url := opts.URL
	region := metadata.MustGet(m, metadata.RegionID)
	if region != "" && utils.GetNetworkType() == "vpc" {
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

	if !onNode || features.FunctionalMutableFeatureGate.Enabled(features.RundCSIProtocol3) {
		return opts
	}
	// get auth config from ENV / metadata service
	switch opts.AuthType {
	case "":
		// try to get ak/sk from env
		if opts.SecretRef == "" && opts.SecurityToken == "" && (opts.AkID == "" || opts.AkSecret == "") {
			ac := utils.GetEnvAK()
			opts.AkID = ac.AccessKeyID
			opts.AkSecret = ac.AccessKeySecret
		}

	case oss.AuthTypeSTS:
		// try to get default ECS worker role from metadata server
		if opts.RoleName == "" {
			workerRole, err := m.Get(metadata.RAMRoleName)
			if err != nil {
				klog.ErrorS(err, "get worker role name failed")
			} else {
				// Note: ossfs 1.0 supports 2 ram_role formats:
				// 1. utils.MetadataURL + utils.WorkerRoleResource + workerRole
				// 	Issues:
				//	* Incompatible with s3fs.
				// 	* Cannot support imdsv2 mode of ECS metadata server.
				// 2. workerRole
				//	Issues:
				// 	* Incompatible with older version (<=1.86) ossfs, which won't be used by containerized mounter.
				opts.RoleName = workerRole
			}
		}
	}

	return opts
}

func setCNFSOptions(ctx context.Context, cnfsGetter cnfsv1beta1.CNFSGetter, opts *oss.Options) error {
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

const (
	csiDefaultFsType = "csi.storage.k8s.io/fstype"
	fuseType         = "fuseType"
)

// setFsType assigns the value of fuseType to csiDefaultFsType.
func setFsType(vc map[string]string) {
	if vc == nil {
		return
	}
	_, ok1 := vc[csiDefaultFsType]
	_, ok2 := vc[fuseType]
	if ok2 && !ok1 {
		vc[csiDefaultFsType] = vc[fuseType]
	}
}

// Check oss options
func checkOssOptions(opt *oss.Options, fpm *oss.OSSFusePodManager) error {
	if fpm == nil {
		return WrapOssError(ParamError, "Unsupported fuseType %s", opt.FuseType)
	}
	// common
	if opt.URL == "" || opt.Bucket == "" {
		return WrapOssError(ParamError, "Url/Bucket empty")
	}

	err := validateEndpoint(opt.URL, opt.Bucket)
	if err != nil {
		return WrapOssError(UrlError, "url is invalid, %v", err)
	}

	// TODO: ossfs2 should matain compatibility with ossfs on encryption,
	// then remove this `switch`
	switch opt.FuseType {
	case OssFsType:
		if !strings.HasPrefix(opt.Path, "/") {
			return WrapOssError(PathError, "start with %s, should start with /", opt.Path)
		}

		if opt.Encrypted != "" && opt.Encrypted != oss.EncryptedTypeKms && opt.Encrypted != oss.EncryptedTypeAes256 {
			return WrapOssError(EncryptError, "invalid SSE encrypted type")
		}

	case OssFs2Type:
		if opt.Encrypted != "" {
			return WrapOssError(EncryptError, "ossfs2 does not support encryption")
		}
	}

	return nil
}

func makeAuthConfig(opt *oss.Options, fpm *oss.OSSFusePodManager, m metadata.MetadataProvider, onNode bool) (*mounter.AuthConfig, error) {
	err := fpm.PrecheckAuthConfig(opt, onNode)
	if err != nil {
		return nil, WrapOssError(AuthError, "%s: %v", opt.FuseType, err)
	}
	authCfg, err := fpm.MakeAuthConfig(opt, m)
	if err != nil {
		return nil, WrapOssError(AuthError, "%s: %v", opt.FuseType, err)
	}
	return authCfg, nil
}

func makeMountOptions(opt *oss.Options, fpm *oss.OSSFusePodManager, m metadata.MetadataProvider, volumeCapability *csi.VolumeCapability) (
	mountOptions []string, err error) {
	mountOptions, err = parseOtherOpts(opt.OtherOpts)
	if err != nil {
		return nil, err
	}

	// TODO: currently the common options for fuse like kernel_cache set on mountOptions will be ignored.
	if volumeCapability != nil && volumeCapability.GetMount() != nil && volumeCapability.GetMount().MountFlags != nil {
		if opt.FuseType == OssFs2Type {
			klog.Warningf("NodePublishVolume: ossfs2 does not support mountOptions, only support volumeAttributes")
		} else {
			mountOptions = append(mountOptions, volumeCapability.GetMount().MountFlags...)
		}
	}

	ops, err := fpm.MakeMountOptions(opt, m)
	if err != nil {
		return nil, err
	}
	mountOptions = append(mountOptions, ops...)
	return
}

func needRotateToken(opt *oss.Options, secrets map[string]string) (needRotate bool) {
	if len(secrets) == 0 {
		return false
	}
	// TODO: Remove this check if when ossfs support rotate fixed AKSK.
	ak := secrets[mounter.GetPasswdFileName(opt.FuseType)]
	if ak == "" {
		needRotate = true
	}
	return
}
