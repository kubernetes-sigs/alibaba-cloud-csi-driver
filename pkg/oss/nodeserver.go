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
	"path/filepath"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	cnfsv1beta1 "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

type nodeServer struct {
	metadata   metadata.MetadataProvider
	locks      *utils.VolumeLocks
	nodeName   string
	clientset  kubernetes.Interface
	cnfsGetter cnfsv1beta1.CNFSGetter
	rawMounter mountutils.Interface
	ossfs      mounter.FuseMounterType
	common.GenericNodeServer
	skipAttach bool
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
	valid, err := utils.ValidatePath(req.GetTargetPath())
	if !valid {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}
	return nil
}

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	klog.Infof("NodePublishVolume:: Starting Mount volume: %s mount with req: %+v", req.VolumeId, req)
	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)

	targetPath := req.GetTargetPath()
	if err := validateNodePublishVolumeRequest(req); err != nil {
		return nil, err
	}
	// check if already mounted
	notMnt, err := isNotMountPoint(ns.rawMounter, targetPath, true)
	if err != nil {
		return nil, err
	}
	if !notMnt {
		klog.Infof("NodePublishVolume: %s already mounted", targetPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	// parse options
	region, _ := ns.metadata.Get(metadata.RegionID)
	opts := parseOptions(req.GetVolumeContext(), req.GetSecrets(), []*csi.VolumeCapability{req.GetVolumeCapability()}, req.GetReadonly(), region, "")
	if err := setCNFSOptions(ctx, ns.cnfsGetter, opts); err != nil {
		return nil, err
	}
	// options validation
	if err := checkOssOptions(opts); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	switch opts.AuthType {
	case "":
		// try to get ak/sk from env
		if opts.SecretRef == "" && (opts.AkID == "" || opts.AkSecret == "") {
			ac := utils.GetEnvAK()
			opts.AkID = ac.AccessKeyID
			opts.AkSecret = ac.AccessKeySecret
		}
		if !features.FunctionalMutableFeatureGate.Enabled(features.RundCSIProtocol3) {
			if opts.SecretRef == "" && (opts.AkID == "" || opts.AkSecret == "") {
				return nil, status.Error(codes.InvalidArgument, "missing access key in node publish secret")
			}
		}
	case mounter.AuthTypeSTS:
		// try to get default ECS worker role from metadata server
		if opts.RoleName == "" {
			workerRole, _ := utils.GetMetaData(utils.WorkerRoleResource)
			if workerRole != "" {
				opts.RoleName = utils.MetadataURL + utils.WorkerRoleResource + workerRole
			} else {
				return nil, status.Error(codes.InvalidArgument, "missing roleName or ramRole in volume attributes")
			}
		}
	}

	// PVs with attribute direct=true will by mounted in rund guest os.
	// csi need only to prepare the direct-volume mountinfo in /run/kata-containers/shared/direct-volumes directory.
	if opts.directAssigned {
		return ns.publishDirectVolume(ctx, req, opts)
	}

	// make mount options and auth config
	mountSource := fmt.Sprintf("%s:%s", opts.Bucket, opts.Path)
	mountOptions, authCfg, err := opts.MakeMountOptionsAndAuthConfig(ns.metadata, req.VolumeCapability)
	if err != nil {
		return nil, err
	}
	if ns.ossfs != nil {
		mountOptions = ns.ossfs.AddDefaultMountOptions(mountOptions)
	}

	// rund 3.0 protocol
	if features.FunctionalMutableFeatureGate.Enabled(features.RundCSIProtocol3) {
		if ns.clientset != nil && utils.GetPodRunTime(req, ns.clientset) == utils.RundRunTimeTag {
			klog.Infof("NodePublishVolume: skip as %s enabled", features.RundCSIProtocol3)
			return &csi.NodePublishVolumeResponse{}, nil
		}
	}

	// get mount proxy socket path
	socketPath := req.PublishContext[mountProxySocket]
	if socketPath == "" {
		return nil, status.Errorf(codes.InvalidArgument, "%s not found in publishContext", mountProxySocket)
	}
	proxyMounter := mounter.NewProxyMounter(socketPath, ns.rawMounter)

	// When work as csi-agent, directly mount on the target path.
	if ns.skipAttach {
		utils.WriteMetricsInfo(metricsPathPrefix, req, opts.MetricsTop, OssFsType, "oss", opts.Bucket)
		err := proxyMounter.MountWithSecrets(mountSource, targetPath, opts.FuseType, mountOptions, authCfg.Secrets)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		klog.Infof("NodePublishVolume(csi-agent): successfully mounted %s on %s", mountSource, targetPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	// When work as csi nodeserver, mount on the attach path under /run/fuse.ossfs and then perform the bind mount.
	// check whether the attach path is mounted
	attachPath := mounter.GetOssfsAttachPath(req.VolumeId)
	notMnt, err = isNotMountPoint(ns.rawMounter, attachPath, false)
	if err != nil {
		return nil, err
	}
	if notMnt {
		utils.WriteSharedMetricsInfo(metricsPathPrefix, req, OssFsType, "oss", opts.Bucket, attachPath)
		err := mounter.NewProxyMounter(socketPath, ns.rawMounter).MountWithSecrets(
			mountSource, attachPath, opts.FuseType, mountOptions, authCfg.Secrets)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		klog.Infof("NodePublishVolume: successfully mounted volume %s on %s", req.VolumeId, attachPath)
	}

	// bind mount
	if err := ns.rawMounter.Mount(attachPath, targetPath, "", []string{"bind"}); err != nil {
		return nil, status.Errorf(codes.Internal, "bind mount failed: %v", err)
	}
	klog.Infof("NodePublishVolume: bind mounted %s to %s", attachPath, targetPath)

	return &csi.NodePublishVolumeResponse{}, nil
}

// Check oss options
func checkOssOptions(opt *Options) error {
	if opt.FuseType != OssFsType {
		return WrapOssError(ParamError, "only ossfs fuse type supported")
	}

	if opt.URL == "" || opt.Bucket == "" {
		return WrapOssError(ParamError, "Url/Bucket empty")
	}

	err := validateEndpoint(opt.URL, opt.Bucket)
	if err != nil {
		return WrapOssError(UrlError, "url is invalid, %v", err)
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
		if opt.SecretRef != "" {
			if opt.AkID != "" || opt.AkSecret != "" {
				return WrapOssError(AuthError, "AK and secretRef cannot be set at the same time")
			}
			if opt.SecretRef == mounter.OssfsCredentialSecretName {
				return WrapOssError(ParamError, "invalid SecretRef name")
			}
		}
	}

	if opt.Encrypted != "" && opt.Encrypted != EncryptedTypeKms && opt.Encrypted != EncryptedTypeAes256 {
		return WrapOssError(EncryptError, "invalid SSE encrypted type")
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
	klog.Infof("NodeUnpublishVolume: Starting Umount OSS: %s mount with req: %+v", req.TargetPath, req)
	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)
	targetPath := req.TargetPath
	err := validateNodeUnpublishVolumeRequest(req)
	if err != nil {
		return nil, err
	}
	if isDirectVolumePath(targetPath) {
		return ns.unPublishDirectVolume(ctx, req)
	}

	err = mountutils.CleanupMountPoint(targetPath, ns.rawMounter, true)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmount target %q: %v", targetPath, err)
	}
	klog.Infof("NodeUnpublishVolume: Umount OSS Successful: %s", targetPath)
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
	klog.Infof("NodeUnstageVolume: starting to unmount volume, volumeId: %s, target: %v", req.VolumeId, req.StagingTargetPath)
	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)

	attachPath := mounter.GetOssfsAttachPath(req.VolumeId)
	err := mountutils.CleanupMountPoint(attachPath, ns.rawMounter, false)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmount target %q: %v", attachPath, err)
	}

	// In the legacy mount process, NodePublishVolume creates ossfs pods in kube-system namespace to mount oss.
	// We still need to umount the mountpoint in case csi-plugin is upgraded from these versions.
	err = mountutils.CleanupMountPoint(req.StagingTargetPath, ns.rawMounter, false)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to unmount target %q: %v", req.StagingTargetPath, err)
	}

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

type publishRequest interface {
	GetVolumeCapability() *csi.VolumeCapability
	GetReadonly() bool
	GetVolumeContext() map[string]string
	GetSecrets() map[string]string
}

func (o *Options) MakeMountOptionsAndAuthConfig(m metadata.MetadataProvider, volumeCapability *csi.VolumeCapability) ([]string, *mounter.AuthConfig, error) {
	region, _ := m.Get(metadata.RegionID)

	mountOptions, err := parseOtherOpts(o.OtherOpts)
	if err != nil {
		return nil, nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if volumeCapability != nil && volumeCapability.GetMount() != nil {
		mountOptions = append(mountOptions, volumeCapability.GetMount().MountFlags...)
	}

	switch o.Encrypted {
	case EncryptedTypeAes256:
		mountOptions = append(mountOptions, "use_sse")
	case EncryptedTypeKms:
		if o.KmsKeyId == "" {
			mountOptions = append(mountOptions, "use_sse=kmsid")
		} else {
			mountOptions = append(mountOptions, fmt.Sprintf("use_sse=kmsid:%s", o.KmsKeyId))
		}
	}
	if o.ReadOnly {
		mountOptions = append(mountOptions, "ro")
	}
	// set use_metrics to enabled monitoring by default
	if features.FunctionalMutableFeatureGate.Enabled(features.UpdatedOssfsVersion) {
		mountOptions = append(mountOptions, "use_metrics")
	}
	if o.MetricsTop != "" {
		mountOptions = append(mountOptions, fmt.Sprintf("metrics_top=%s", o.MetricsTop))
	}

	switch o.SigVersion {
	case SigV1:
		mountOptions = append(mountOptions, "sigv1")
	case SigV4:
		if region == "" {
			return nil, nil, status.Errorf(codes.Internal, "SigV4 is not supported without region")
		}
		mountOptions = append(mountOptions, "sigv4")
		mountOptions = append(mountOptions, fmt.Sprintf("region=%s", region))
	}

	mountOptions = append(mountOptions, fmt.Sprintf("url=%s", o.URL))

	authCfg := &mounter.AuthConfig{AuthType: o.AuthType}
	switch o.AuthType {
	case mounter.AuthTypeRRSA:
		rrsaCfg, err := getRRSAConfig(o, m)
		if err != nil {
			return nil, nil, status.Errorf(codes.Internal, "Get RoleArn and OidcProviderArn for RRSA error: %v", err)
		}
		authCfg.RrsaConfig = rrsaCfg
		region, _ := m.Get(metadata.RegionID)
		mountOptions = append(mountOptions, fmt.Sprintf("rrsa_endpoint=%s", getSTSEndpoint(region)))
		if o.AssumeRoleArn != "" {
			mountOptions = append(mountOptions, fmt.Sprintf("assume_role_arn=%s", o.AssumeRoleArn))
			if o.ExternalId != "" {
				mountOptions = append(mountOptions, fmt.Sprintf("assume_role_external_id=%s", o.ExternalId))
			}
		}
	case mounter.AuthTypeCSS:
		authCfg.SecretProviderClassName = o.SecretProviderClass
		mountOptions = append(mountOptions, fmt.Sprintf("secret_store_dir=%s", "/etc/ossfs/secrets-store"))
	case mounter.AuthTypeSTS:
		authCfg.RoleName = o.RoleName
		if o.RoleName != "" {
			mountOptions = append(mountOptions, "ram_role="+o.RoleName)
		}
	default:
		if o.SecretRef != "" {
			authCfg.SecretRef = o.SecretRef
			mountOptions = append(mountOptions, fmt.Sprintf("passwd_file=%s", filepath.Join(mounter.PasswdMountDir, mounter.PasswdFilename)))
			mountOptions = append(mountOptions, "use_session_token")
		} else {
			authCfg.Secrets = map[string]string{
				mounter.OssfsPasswdFile: fmt.Sprintf("%s:%s:%s", o.Bucket, o.AkID, o.AkSecret),
			}
		}
	}

	return mountOptions, authCfg, nil
}
