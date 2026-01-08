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
	fpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager"
	ossfpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager/oss"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
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

const (
	// AkID and AkSecret are the key names we have always supported.
	// For compatibility with standard naming conventions, we also support customers configuring `accessKeyID` and `accessKeySecret`
	AkID     = "akId"
	AkSecret = "akSecret"
)

// parseCredentialsFromSecret retrieves credentials from the `Secret` field in request.
// It supports two credential formats:
//  1. Legacy format: `akId` and `akSecret` (case-sensitive, no trimming for key names)
//  2. Standard format: `AccessKeyId`, `AccessKeySecret`, `SecurityToken`, `Expiration` (case-insensitive)
//
// Priority and return logic:
//  1. Legacy keys (`akId` and `akSecret`):
//     - If both are set: returns AccessKey with legacy values, TokenSecret is empty (fixed credentials)
//     - If only one is set: returns empty AccessKey and empty TokenSecret (invalid, caller should handle)
//     - If both are empty: continues to check standard keys
//  2. Standard keys (searched in case-insensitive manner for SecretRef compatibility):
//     - If AccessKeyId or AccessKeySecret is missing: returns empty AccessKey and empty TokenSecret
//     - If SecurityToken is empty: returns AccessKey with standard values, TokenSecret is empty (fixed credentials)
//     - If SecurityToken is set: returns empty AccessKey, TokenSecret with all values (STS token credentials)
//
// Returns:
//   - AccessKey: Contains AkID and AkSecret for fixed credentials (either legacy or standard format)
//   - TokenSecret: Contains AccessKeyId, AccessKeySecret, SecurityToken, and Expiration for STS token credentials
//     Note: Only one of AccessKey or TokenSecret will be non-empty, never both
func parseCredentialsFromSecret(secrets map[string]string) (ossfpm.AccessKey, ossfpm.TokenSecret) {
	var nilAk = ossfpm.AccessKey{}
	var nilToken = ossfpm.TokenSecret{}
	// Maintain backward compatibility: legacy keys are case-sensitive and not trimmed
	akID := strings.TrimSpace(secrets[AkID])
	akSecret := strings.TrimSpace(secrets[AkSecret])
	if akID != "" && akSecret != "" {
		return ossfpm.AccessKey{
			AkID:     akID,
			AkSecret: akSecret,
		}, nilToken
	}
	// If either legacy key is set, return immediately without checking standard keys
	if akID != "" || akSecret != "" {
		return nilAk, nilToken
	}

	var token, expiration string
	// For SecretRef compatibility, search for standard keys in case-insensitive manner
	for k, v := range secrets {
		key := strings.TrimSpace(strings.ToLower(k))
		value := strings.TrimSpace(v)
		if value == "" {
			continue
		}
		// Can be either fixed credentials or rotating STS token
		switch key {
		case strings.ToLower(mounterutils.KeyAccessKeyId):
			akID = value
		case strings.ToLower(mounterutils.KeyAccessKeySecret):
			akSecret = value
		case strings.ToLower(mounterutils.KeySecurityToken):
			token = value
		case strings.ToLower(mounterutils.KeyExpiration):
			expiration = value
		}
	}
	// if akID or akSecret is empty, return nil
	if akID == "" || akSecret == "" {
		return nilAk, nilToken
	}
	// if token is empty, see as fixed credentials
	if token == "" {
		return ossfpm.AccessKey{
			AkID:     akID,
			AkSecret: akSecret,
		}, nilToken
	}
	// else, see as STS.Token
	return nilAk, ossfpm.TokenSecret{
		AccessKeyId:     akID,
		AccessKeySecret: akSecret,
		Expiration:      expiration, // optional
		SecurityToken:   token,
	}
}

// get Options for CreateVolume and PublishVolume
// volOptions: CreateVolumeRequest.GetParameters, PublishVolumeRequest.GetVolumeContext
// secrets: CreateVolumeRequest / PublishVolumeRequest.GetSecrets
// volCaps: CreateVolumeRequest.GetVolumeCapabilities, []{PublishVolumeRequest.GetVolumeCapability}
// readOnly: PublishVolumeRequest.GetReadonly
// region: for signature v4
// reqName: for subpath generating, CreateVolumeRequest.GetName
// onNode: run on nodeserver mode
func parseOptions(volOptions, secrets map[string]string, volCaps []*csi.VolumeCapability,
	readOnly bool, reqName string, onNode bool, m metadata.MetadataProvider) *ossfpm.Options {

	if volOptions == nil {
		volOptions = map[string]string{}
	}
	if secrets == nil {
		secrets = map[string]string{}
	}

	// credientials
	accessKey, tokenSecret := parseCredentialsFromSecret(secrets)
	opts := &ossfpm.Options{
		UseSharedPath: true,
		Path:          "/",
		AccessKey:     accessKey,
		TokenSecret:   tokenSecret,
	}

	var volumeAsSubpath bool
	var runtimeClassValue, directAssignedValue string
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
			directAssignedValue = value
		case "encrypted":
			opts.Encrypted = strings.ToLower(value)
		case "kmskeyid":
			opts.KmsKeyId = value
		case "assumerolearn":
			opts.AssumeRoleArn = value
		case "externalid":
			opts.ExternalId = value
		case "sigversion":
			switch ossfpm.SigVersion(value) {
			case ossfpm.SigV1, ossfpm.SigV4:
				opts.SigVersion = ossfpm.SigVersion(value)
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
		case "dnspolicy":
			switch strings.ToLower(value) {
			case strings.ToLower(string(corev1.DNSClusterFirst)),
				strings.ToLower(string(corev1.DNSClusterFirstWithHostNet)),
				strings.ToLower(string(corev1.DNSDefault)):
				opts.DnsPolicy = corev1.DNSPolicy(value)
			default:
				klog.Warning(WrapOssError(ParamError, "the value(%q) of %q is invalid, only support %v, %v, %v",
					v, k, corev1.DNSClusterFirst, corev1.DNSClusterFirstWithHostNet, corev1.DNSDefault).Error())
			}
		case "runtimeclass":
			runtimeClassValue = value
		// deprecated:
		case strings.ToLower(AkID):
			opts.AkID = value
		case strings.ToLower(AkSecret):
			opts.AkSecret = value
		}
	}

	opts.DirectAssigned = parseDirectAssigned(runtimeClassValue, directAssignedValue)

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
		opts.FuseType = mounterutils.OssFsType
	}

	url := opts.URL
	region, _ := m.Get(metadata.RegionID)
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
		if opts.SecretRef == "" &&
			(opts.AkID == "" && opts.AkSecret == "") &&
			(opts.AccessKeyId == "" && opts.AccessKeySecret == "" && opts.SecurityToken == "") {
			ac := utils.GetEnvAK()
			opts.AkID = ac.AccessKeyID
			opts.AkSecret = ac.AccessKeySecret
		}

	case ossfpm.AuthTypeSTS:
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

func setCNFSOptions(ctx context.Context, cnfsGetter cnfsv1beta1.CNFSGetter, opts *ossfpm.Options) error {
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
func checkOssOptions(opt *ossfpm.Options, fpm *ossfpm.OSSFusePodManager) error {
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
	case mounterutils.OssFsType:
		if !strings.HasPrefix(opt.Path, "/") {
			return WrapOssError(PathError, "start with %s, should start with /", opt.Path)
		}

		if opt.Encrypted != "" && opt.Encrypted != ossfpm.EncryptedTypeKms && opt.Encrypted != ossfpm.EncryptedTypeAes256 {
			return WrapOssError(EncryptError, "invalid SSE encrypted type")
		}

	case mounterutils.OssFs2Type:
		if opt.Encrypted != "" {
			return WrapOssError(EncryptError, "ossfs2 does not support encryption")
		}
	}

	return nil
}

func makeAuthConfig(opt *ossfpm.Options, fpm *ossfpm.OSSFusePodManager, m metadata.MetadataProvider, onNode bool) (*fpm.AuthConfig, error) {
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
func makePodTemplateConfig(opt *ossfpm.Options) *fpm.PodTemplateConfig {
	return &fpm.PodTemplateConfig{
		DnsPolicy: opt.DnsPolicy,
	}
}

func makeMountOptions(opt *ossfpm.Options, fpm *ossfpm.OSSFusePodManager, m metadata.MetadataProvider, volumeCapability *csi.VolumeCapability) (
	mountOptions []string, err error) {
	mountOptions, err = parseOtherOpts(opt.OtherOpts)
	if err != nil {
		return nil, err
	}

	// TODO: currently the common options for fuse like kernel_cache set on mountOptions will be ignored.
	if volumeCapability != nil && volumeCapability.GetMount() != nil && volumeCapability.GetMount().MountFlags != nil {
		if opt.FuseType == mounterutils.OssFs2Type {
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

// parseDirectAssigned parses the DirectAssigned value from volOptions with the following priority:
// 1. If both optDirectAssigned and "runtimeclass" are not set, use DEFAULT_RUNTIME_CLASS from env
// 2. If "runtimeclass" is set, use getDirectAssignedValue to get DirectAssigned (can override 1)
// 3. If optDirectAssigned is set, directly override 1 and 2; but if error, fallback to 1 or 2
func parseDirectAssigned(runtimeClassValue, directAssignedValue string) bool {
	// Priority 1 & 2: default from env, or from runtimeclass if set
	result := getDirectAssignedValue(runtimeClassValue)

	// Priority 3: if optDirectAssigned is set, try to parse and override (if error, keep previous value)
	if directAssignedValue != "" {
		if res, err := strconv.ParseBool(directAssignedValue); err == nil {
			result = res
		} else {
			klog.Warning(WrapOssError(ParamError, "the value(%q) of %q is invalid, fallback to previous value", directAssignedValue, optDirectAssigned).Error())
		}
	}

	return result
}

// RuntimeType represents the container runtime type
type RuntimeType string

const (
	RuntimeTypeCOCO    RuntimeType = "coco"
	RuntimeTypeRunD    RuntimeType = "rund"
	RuntimeTypeRunC    RuntimeType = "runc"
	RuntimeTypeMicroVM RuntimeType = "eci"
	RuntimeTypeUnknown RuntimeType = "unknown"
)

// runtimeTypeResult represents the result of runtime type determination
type runtimeTypeResult struct {
	runtimeType RuntimeType
	err         error
}

// runtimeTypeLookupTable is a lookup table for all 8 possible combinations of
// (directAssigned, hasSocketPath, skipAttach).
// Index calculation: index = (directAssigned ? 4 : 0) + (hasSocketPath ? 2 : 0) + (skipAttach ? 1 : 0)å
var runtimeTypeLookupTable = [8]runtimeTypeResult{
	// Index 0: directAssigned=false, hasSocketPath=false, skipAttach=false -> error (should not occur)
	{
		runtimeType: RuntimeTypeUnknown,
		err:         fmt.Errorf("socket path is empty in non csi-agent mode (should not occur)"),
	},
	// Index 1: directAssigned=false, hasSocketPath=false, skipAttach=true -> MicroVM
	{
		runtimeType: RuntimeTypeMicroVM,
		err:         nil,
	},
	// Index 2: directAssigned=false, hasSocketPath=true, skipAttach=false -> RunC
	{
		runtimeType: RuntimeTypeRunC,
		err:         nil,
	},
	// Index 3: directAssigned=false, hasSocketPath=true, skipAttach=true -> RunD (pure rund cluster, no need to specify directAssigned)
	{
		runtimeType: RuntimeTypeRunD,
		err:         nil,
	},
	// Index 4: directAssigned=true, hasSocketPath=false, skipAttach=false -> COCO
	{
		runtimeType: RuntimeTypeCOCO,
		err:         nil,
	},
	// Index 5: directAssigned=true, hasSocketPath=false, skipAttach=true -> MicroVM (directAssigned not effective)
	{
		runtimeType: RuntimeTypeMicroVM,
		err:         nil,
	},
	// Index 6: directAssigned=true, hasSocketPath=true, skipAttach=false -> error (should not occur, controller ensures empty socketPath for COCO)
	{
		runtimeType: RuntimeTypeUnknown,
		err:         fmt.Errorf("rund cannot be used in non csi-agent mode (should not occur)"),
	},
	// Index 7: directAssigned=true, hasSocketPath=true, skipAttach=true -> RunD
	{
		runtimeType: RuntimeTypeRunD,
		err:         nil,
	},
}

// DetermineRuntimeType determines the container runtime type based on directAssigned, socketPath, and skipAttach.
//
// This function uses a table-driven approach with a 3-bit index calculated from the three boolean parameters.
// The lookup table contains all 8 possible combinations, providing O(1) lookup performance.
//
// Returns:
//   - RuntimeType: the determined runtime type
//   - error: if the combination is invalid (should not occur scenarios)
func DetermineRuntimeType(directAssigned bool, socketPath string, skipAttach bool) (RuntimeType, error) {
	hasSocketPath := socketPath != ""

	// Calculate index: (directAssigned ? 4 : 0) + (hasSocketPath ? 2 : 0) + (skipAttach ? 1 : 0)
	index := 0
	if directAssigned {
		index += 4
	}
	if hasSocketPath {
		index += 2
	}
	if skipAttach {
		index += 1
	}

	// Lookup result from table
	result := runtimeTypeLookupTable[index]
	return result.runtimeType, result.err
}

// getDirectAssignedValue returns the default value for DirectAssigned option based on the runtime class.
// The function reads DEFAULT_RUNTIME_CLASS environment variable and determines the appropriate default value:
// - For "rund" runtime, returns true (direct assignment enabled by default)
// - For "runc" runtime or empty value, returns false (direct assignment disabled by default)
// - For any other value, logs a warning and returns false
func getDirectAssignedValue(runtimeClass string) bool {
	// Get runtime class from environment variable
	if runtimeClass == "" {
		runtimeClass = os.Getenv("DEFAULT_RUNTIME_CLASS")
	}

	// Validate runtime class, only allow "runc", "rund" or empty (treated as "runc")
	// Note: Do not consider the confidential container scenario (coco),
	//  because in this scenario, PV.attributes must explicitly specify whether directAssigned
	switch strings.ToLower(runtimeClass) {
	case strings.ToLower(utils.RundRunTimeTag):
		return true
	case strings.ToLower(utils.RuncRunTimeTag), "":
		return false
	default:
		// Invalid runtime class, see as runc and return error
		klog.Warningf("invalid runtimeClass value: %q, only %s and %s are allowed", runtimeClass, utils.RuncRunTimeTag, utils.RundRunTimeTag)
		return false
	}
}

func needRotateToken(fuseType string, secrets map[string]string) (needRotate bool) {
	if len(secrets) == 0 {
		return false
	}
	// TODO: Remove this check if when ossfs support rotate fixed AKSK.
	ak := secrets[mounterutils.GetPasswdFileName(fuseType)]
	token := secrets[mounterutils.KeySecurityToken]
	if ak == "" && token != "" {
		needRotate = true
	}
	return
}
