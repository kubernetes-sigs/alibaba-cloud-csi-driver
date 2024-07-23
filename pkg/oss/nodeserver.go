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
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	mountutils "k8s.io/mount-utils"
)

type nodeServer struct {
	metadata metadata.MetadataProvider
	*csicommon.DefaultNodeServer
	nodeName        string
	clientset       kubernetes.Interface
	dynamicClient   dynamic.Interface
	sharedPathLock  *utils.VolumeLocks
	ossfsMounterFac *mounter.ContainerizedFuseMounterFactory
}

// Options contains options for target oss
type Options struct {
	directAssigned bool

	// oss options
	Bucket string `json:"bucket"`
	URL    string `json:"url"`
	Path   string `json:"path"`

	// authorization options
	// accesskey
	AkID      string `json:"akId"`
	AkSecret  string `json:"akSecret"`
	SecretRef string `json:"secretRef"`
	// RRSA
	RoleName           string `json:"roleName"`
	RoleArn            string `json:"roleArn"`
	OidcProviderArn    string `json:"oidcProviderArn"`
	ServiceAccountName string `json:"serviceAccountName"`
	// assume role
	AssumeRoleArn string `json:"assumeRoleArn"`
	ExternalId    string `json:"externalId"`
	// csi secret store
	SecretProviderClass string `json:"secretProviderClass"`

	// ossfs options
	OtherOpts  string `json:"otherOpts"`
	MetricsTop string `json:"metricsTop"`
	Encrypted  string `json:"encrypted"`
	KmsKeyId   string `json:"kmsKeyId"`

	// mount options
	UseSharedPath bool   `json:"useSharedPath"`
	AuthType      string `json:"authType"`
	FuseType      string `json:"fuseType"`
	ReadOnly      bool   `json:"readOnly"`
}

const (
	// OssfsCredentialFile is the path of oss ak credential file
	OssfsCredentialFile = "/host/etc/passwd-ossfs"
	// AkID is Ak ID
	AkID = "akId"
	// AkSecret is Ak Secret
	AkSecret = "akSecret"
	// OssFsType is the oss filesystem type
	OssFsType = "ossfs"
	// metricsPathPrefix
	metricsPathPrefix = "/host/var/run/ossfs/"
	// defaultMetricsTop
	defaultMetricsTop = "10"
	// fuseServieAccountName
	fuseServieAccountName = "csi-fuse-ossfs"
)

const (
	EncryptedTypeKms    = "kms"
	EncryptedTypeAes256 = "aes256"
)

func (ns *nodeServer) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return &csi.NodeGetCapabilitiesResponse{Capabilities: []*csi.NodeServiceCapability{
		{
			Type: &csi.NodeServiceCapability_Rpc{
				Rpc: &csi.NodeServiceCapability_RPC{
					Type: csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME,
				},
			},
		},
	}}, nil
}

func validateNodePublishVolumeRequest(req *csi.NodePublishVolumeRequest) error {
	valid, err := utils.CheckRequest(req.GetVolumeContext(), req.GetTargetPath())
	if !valid {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}
	return nil
}

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	// logout oss paras
	log.Infof("NodePublishVolume:: Starting Mount volume: %s mount with req: %+v", req.VolumeId, req)
	mountPath := req.GetTargetPath()
	if err := validateNodePublishVolumeRequest(req); err != nil {
		return nil, err
	}
	var cnfsName string
	opt := &Options{
		UseSharedPath: true,
		FuseType:      OssFsType,
		Path:          "/",
		AkID:          req.Secrets[AkID],
		AkSecret:      req.Secrets[AkSecret],
	}
	for k, v := range req.VolumeContext {
		key := strings.TrimSpace(strings.ToLower(k))
		value := strings.TrimSpace(v)
		if value == "" {
			continue
		}
		switch key {
		case "bucket":
			opt.Bucket = value
		case "url":
			opt.URL = value
		case "otheropts":
			opt.OtherOpts = value
		case "akid":
			opt.AkID = value
		case "aksecret":
			opt.AkSecret = value
		case "secretref":
			opt.SecretRef = value
		case "path":
			opt.Path = value
		case "usesharedpath":
			if res, err := strconv.ParseBool(value); err == nil {
				opt.UseSharedPath = res
			} else {
				log.Warnf("Oss parameters error: the value(%q) of %q is invalid", v, k)
			}
		case "authtype":
			opt.AuthType = strings.ToLower(value)
		case "rolename":
			opt.RoleName = value
		case "rolearn":
			opt.RoleArn = value
		case "oidcproviderarn":
			opt.OidcProviderArn = value
		case "serviceaccountname":
			opt.ServiceAccountName = value
		case "secretproviderclass":
			opt.SecretProviderClass = value
		case "fusetype":
			opt.FuseType = strings.ToLower(value)
		case "metricstop":
			opt.MetricsTop = strings.ToLower(value)
		case "containernetworkfilesystem":
			cnfsName = value
		case optDirectAssigned:
			if res, err := strconv.ParseBool(value); err == nil {
				opt.directAssigned = res
			} else {
				log.Warnf("Oss parameters error: the value(%q) of %q is invalid", v, k)
			}
		case "encrypted":
			opt.Encrypted = strings.ToLower(value)
		case "kmskeyid":
			opt.KmsKeyId = value
		case "assumerolearn":
			opt.AssumeRoleArn = value
		case "externalid":
			opt.ExternalId = value
		default:
			log.Warnf("Oss parameters error: the key(%q) is unknown", k)
		}
	}

	if req.Readonly {
		opt.ReadOnly = true
	} else {
		switch req.VolumeCapability.AccessMode.GetMode() {
		case csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER, csi.VolumeCapability_AccessMode_MULTI_NODE_SINGLE_WRITER, csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER:
			opt.ReadOnly = false
		default:
			opt.ReadOnly = true
		}
	}

	if len(opt.Bucket) == 0 {
		cnfs, err := v1beta1.GetCnfsObject(ns.dynamicClient, cnfsName)
		if err != nil {
			return nil, err
		}
		if cnfs.Status.FsAttributes.EndPoint == nil {
			return nil, errors.New("Cnfs " + cnfsName + " is not ready, endpoint is empty.")
		}
		opt.Bucket = cnfs.Status.FsAttributes.BucketName
		opt.URL = cnfs.Status.FsAttributes.EndPoint.Internal
	}

	// check parameters
	if err := checkOssOptions(opt); err != nil {
		log.Errorf("Check oss input error: %s", err.Error())
		return nil, errors.New("Check oss input error: " + err.Error())
	}

	argStr := fmt.Sprintf("Bucket: %s, url: %s, OtherOpts: %s, Path: %s, UseSharedPath: %s, authType: %s, encrypted: %s, kmsid: %s",
		opt.Bucket, opt.URL, opt.OtherOpts, opt.Path, strconv.FormatBool(opt.UseSharedPath), opt.AuthType, opt.Encrypted, opt.KmsKeyId)
	log.Infof("NodePublishVolume:: Starting Oss Mount: %s", argStr)

	if opt.directAssigned {
		return ns.publishDirectVolume(ctx, req, opt)
	}
	var ossMounter mountutils.Interface
	switch opt.FuseType {
	case OssFsType, "":
		// When useSharedPath options is set to false,
		// mount operations need to be atomic to ensure that no fuse pods are left behind in case of failure.
		// Because kubelet will not call NodeUnpublishVolume when NodePublishVolume never succeeded.

		var rrsaCfg *mounter.RrsaConfig
		var err error
		if opt.AuthType == mounter.AuthTypeRRSA {
			rrsaCfg, err = getRRSAConfig(opt, ns.metadata)
			if err != nil {
				return nil, errors.New("Get RoleArn and OidcProviderArn for RRSA error: " + err.Error())
			}
		}
		authCfg := &mounter.AuthConfig{AuthType: opt.AuthType, RrsaConfig: rrsaCfg, SecretProviderClassName: opt.SecretProviderClass, SecretRef: opt.SecretRef}
		ossMounter = ns.ossfsMounterFac.NewFuseMounter(ctx, req.VolumeId, authCfg, !opt.UseSharedPath)
	default:
		return nil, status.Errorf(codes.InvalidArgument, "unknown fuseType: %q", opt.FuseType)
	}

	notMnt, err := mountutils.IsNotMountPoint(ossMounter, mountPath)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(mountPath, os.ModePerm); err != nil {
				log.Errorf("NodePublishVolume: mkdir %s: %v", mountPath, err)
				return nil, status.Error(codes.Internal, err.Error())
			}
			notMnt = true
		} else {
			log.Errorf("NodePublishVolume: check mountpoint %s: %v", mountPath, err)
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	if !notMnt {
		log.Infof("NodePublishVolume: %s already mounted", mountPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	var mountOptions []string
	if req.VolumeCapability != nil && req.VolumeCapability.GetMount() != nil {
		mountOptions = req.VolumeCapability.GetMount().MountFlags
	}

	regionID, _ := ns.metadata.Get(metadata.RegionID)
	switch opt.AuthType {
	case mounter.AuthTypeSTS:
		mountOptions = append(mountOptions, GetRAMRoleOption())
	case mounter.AuthTypeRRSA:
		if regionID == "" {
			mountOptions = append(mountOptions, "rrsa_endpoint=https://sts.aliyuncs.com")
		} else {
			mountOptions = append(mountOptions, fmt.Sprintf("rrsa_endpoint=https://sts-vpc.%s.aliyuncs.com", regionID))
		}
		if opt.AssumeRoleArn != "" {
			mountOptions = append(mountOptions, fmt.Sprintf("assume_role_arn=%s", opt.AssumeRoleArn))
			if opt.ExternalId != "" {
				mountOptions = append(mountOptions, fmt.Sprintf("assume_role_external_id=%s", opt.ExternalId))
			}
		}
	case mounter.AuthTypeCSS:
	default:
		// ossfs fuse pod will mount the secret to access credentials
		if opt.SecretRef == "" {
			err := mounter.SetupOssfsCredentialSecret(ctx, ns.clientset, ns.nodeName, req.VolumeId, opt.Bucket, opt.AkID, opt.AkSecret)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to setup ossfs credential secret: %v", err)
			}
		}
	}

	switch opt.Encrypted {
	case EncryptedTypeAes256:
		if opt.FuseType == OssFsType {
			mountOptions = append(mountOptions, "use_sse")
		}
	case EncryptedTypeKms:
		if opt.FuseType == OssFsType {
			if opt.KmsKeyId == "" {
				mountOptions = append(mountOptions, "use_sse=kmsid")
			} else {
				mountOptions = append(mountOptions, fmt.Sprintf("use_sse=kmsid:%s", opt.KmsKeyId))
			}
		}
	}

	if opt.ReadOnly {
		mountOptions = append(mountOptions, "ro")
	}

	// set use_metrics to enabled monitoring by default
	if features.FunctionalMutableFeatureGate.Enabled(features.UpdatedOssfsVersion) {
		mountOptions = append(mountOptions, "use_metrics")
	}

	if regionID == "" {
		log.Warnf("NodePublishVolume:: failed to get region id from both env and metadata, use original URL: %s", opt.URL)
	} else if url, done := setNetworkType(opt.URL, regionID); done {
		log.Warnf("NodePublishVolume:: setNetworkType: modified URL from %s to %s", opt.URL, url)
		opt.URL = url
	}
	if url, done := setTransmissionProtocol(opt.URL); done {
		log.Warnf("NodePublishVolume:: setTransmissionProtocol: modified URL from %s to %s", opt.URL, url)
		opt.URL = url
	}

	if opt.UseSharedPath {
		sharedPath := GetGlobalMountPath(req.GetVolumeId())
		notMnt, err := ossMounter.IsLikelyNotMountPoint(sharedPath)
		if err != nil {
			if os.IsNotExist(err) {
				if err := os.MkdirAll(sharedPath, os.ModePerm); err != nil {
					log.Errorf("NodePublishVolume: mkdir %s: %v", sharedPath, err)
					return nil, status.Error(codes.Internal, err.Error())
				}
				notMnt = true
			} else if mountutils.IsCorruptedMnt(err) {
				log.Warnf("Umount corrupted mountpoint %s", sharedPath)
				err := mountutils.New("").Unmount(sharedPath)
				if err != nil {
					return nil, status.Errorf(codes.Internal, "umount corrupted mountpoint %s: %v", sharedPath, err)
				}
				notMnt = true
			} else {
				log.Errorf("NodePublishVolume: check mountpoint %s: %v", sharedPath, err)
				return nil, status.Error(codes.Internal, err.Error())
			}
		}
		if notMnt {
			// serialize node publish operations on the same volume when using sharedpath
			if lock := ns.sharedPathLock.TryAcquire(req.VolumeId); !lock {
				log.Errorf("NodePublishVolume: aborted because failed to acquire volume %s lock", req.VolumeId)
				return nil, status.Errorf(codes.Aborted, "NodePublishVolume operation on shared path of volume %s already exists", req.VolumeId)
			}
			defer ns.sharedPathLock.Release(req.VolumeId)
			utils.WriteSharedMetricsInfo(metricsPathPrefix, req, OssFsType, "oss", opt.Bucket, sharedPath)
			if opt.MetricsTop != "" {
				mountOptions = append(mountOptions, fmt.Sprintf("metrics_top=%s", opt.MetricsTop))
			}
			if err := doMount(ossMounter, sharedPath, *opt, mountOptions); err != nil {
				log.Errorf("NodePublishVolume: failed to mount")
				return nil, err
			}
		} else {
			log.Infof("NodePublishVolume: sharedpath %s already mounted", sharedPath)
		}
		log.Infof("NodePublishVolume:: Start mount operation from source [%s] to dest [%s]", sharedPath, mountPath)
		if err := ossMounter.Mount(sharedPath, mountPath, "", []string{"bind"}); err != nil {
			log.Errorf("Ossfs mount error: %v", err.Error())
			return nil, errors.New("Create oss volume fail: " + err.Error())
		}
	} else {
		if opt.FuseType == OssFsType {
			metricsTop := defaultMetricsTop
			if opt.MetricsTop != "" {
				metricsTop = opt.MetricsTop
			}
			utils.WriteMetricsInfo(metricsPathPrefix, req, metricsTop, OssFsType, "oss", opt.Bucket)
		}
		if err := doMount(ossMounter, mountPath, *opt, mountOptions); err != nil {
			return nil, err
		}
	}

	log.Infof("NodePublishVolume: mounted oss successfully, volume %s, targetPath: %s", req.VolumeId, mountPath)
	return &csi.NodePublishVolumeResponse{}, nil
}

// Check oss options
func checkOssOptions(opt *Options) error {
	if opt.URL == "" || opt.Bucket == "" {
		return WrapOssError(ParamError, "Oss parameters error: Url/Bucket empty")
	}

	if !strings.HasPrefix(opt.Path, "/") {
		return WrapOssError(PathError, "start with "+opt.Path+", should start with /")
	}

	switch opt.AuthType {
	case mounter.AuthTypeSTS:
	case mounter.AuthTypeRRSA:
		if err := checkRRSAParams(opt); err != nil {
			return WrapOssError(AuthError, err.Error())
		}
	case mounter.AuthTypeCSS:
		if opt.SecretProviderClass == "" {
			return WrapOssError(AuthError, "use CsiSecretStore but secretProviderClass is empty")
		}
	default:
		// if not input ak from user, use the default ak value
		if opt.AkID == "" || opt.AkSecret == "" {
			ac := utils.GetEnvAK()
			opt.AkID = ac.AccessKeyID
			opt.AkSecret = ac.AccessKeySecret
		}
		if opt.SecretRef != "" {
			if opt.AkID != "" || opt.AkSecret != "" {
				return WrapOssError(AuthError, "AK and secretRef cannot be set at the same time")
			}
			if opt.SecretRef == mounter.OssfsCredentialSecretName {
				return WrapOssError(ParamError, "invalid SecretRef")
			}
		} else if opt.AkID == "" || opt.AkSecret == "" {
			return WrapOssError(AuthError, "AK and authType are both empty or invalid")
		}
	}

	if opt.Encrypted != "" && opt.Encrypted != EncryptedTypeKms && opt.Encrypted != EncryptedTypeAes256 {
		return WrapOssError(EncryptError, "invalid SSE encryted type")
	}

	if opt.AssumeRoleArn != "" && opt.AuthType != mounter.AuthTypeRRSA {
		return WrapOssError(AuthError, "only support access OSS through STS AssumeRole when authType is RRSA")
	}

	return nil
}

func validateNodeUnpublishVolumeRequest(req *csi.NodeUnpublishVolumeRequest) error {
	valid, err := utils.ValidatePath(req.GetTargetPath())
	if !valid {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}
	return nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	log.Infof("NodeUnpublishVolume: Starting Umount OSS: %s mount with req: %+v", req.TargetPath, req)
	mountPoint := req.TargetPath
	err := validateNodeUnpublishVolumeRequest(req)
	if err != nil {
		return nil, err
	}
	if isDirectVolumePath(mountPoint) {
		return ns.unPublishDirectVolume(ctx, req)
	}
	err = ns.cleanupMountPoint(ctx, req.VolumeId, req.TargetPath)
	if err != nil {
		log.Errorf("NodeUnpublishVolume: failed to unmount %q: %v", mountPoint, err)
		return nil, status.Errorf(codes.Internal, "failed to unmount target %q: %v", mountPoint, err)
	}
	log.Infof("NodeUnpublishVolume: Umount OSS Successful: %s", mountPoint)
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeStageVolume(
	ctx context.Context,
	req *csi.NodeStageVolumeRequest) (
	*csi.NodeStageVolumeResponse, error) {
	return &csi.NodeStageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnstageVolume(
	ctx context.Context,
	req *csi.NodeUnstageVolumeRequest) (
	*csi.NodeUnstageVolumeResponse, error) {
	log.Infof("NodeUnstageVolume: starting to unmount volume, volumeId: %s, target: %v", req.VolumeId, req.StagingTargetPath)
	// unmount for sharedPath
	mountpoint := GetGlobalMountPath(req.VolumeId)
	err := ns.cleanupMountPoint(ctx, req.VolumeId, mountpoint)
	if err != nil {
		log.Errorf("NodeUnstageVolume: failed to unmount %q: %v", mountpoint, err)
		return nil, status.Errorf(codes.Internal, "failed to unmount target %q: %v", mountpoint, err)
	}
	log.Infof("NodeUnstageVolume: umount OSS Successful: %s", mountpoint)
	// if secretRef set, return nil directly as default secret or related key not found here
	err = mounter.CleanupOssfsCredentialSecret(ctx, ns.clientset, ns.nodeName, req.VolumeId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to cleanup ossfs credential secret: %v", err)
	}
	return &csi.NodeUnstageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (
	*csi.NodeExpandVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (ns *nodeServer) cleanupMountPoint(ctx context.Context, volumeId string, mountpoint string) error {
	m := mountutils.New("")
	var err error
	forceUnmounter, ok := m.(mountutils.MounterForceUnmounter)
	if ok {
		err = mountutils.CleanupMountWithForce(mountpoint, forceUnmounter, true, time.Minute)
	} else {
		err = mountutils.CleanupMountPoint(mountpoint, m, true)
	}
	if err != nil {
		return err
	}
	return ns.ossfsMounterFac.NewFuseMounter(ctx, volumeId, nil, false).Unmount(mountpoint)
}
