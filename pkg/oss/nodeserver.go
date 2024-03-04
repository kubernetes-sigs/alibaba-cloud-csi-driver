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
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
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

	Bucket        string `json:"bucket"`
	URL           string `json:"url"`
	OtherOpts     string `json:"otherOpts"`
	AkID          string `json:"akId"`
	AkSecret      string `json:"akSecret"`
	Path          string `json:"path"`
	UseSharedPath bool   `json:"useSharedPath"`
	AuthType      string `json:"authType"`
	FuseType      string `json:"fuseType"`
	MetricsTop    string `json:"metricsTop"`
	ReadOnly      bool   `json:"readOnly"`
	Encrypted     string `json:"encrypted"`
	KmsKeyId      string `json:"kmsKeyId"`
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
	// JindoFsType tag
	JindoFsType = "jindofs"
	// metricsPathPrefix
	metricsPathPrefix = "/host/var/run/ossfs/"
	// defaultMetricsTop
	defaultMetricsTop = "10"
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
	opt := &Options{}
	opt.UseSharedPath = true
	opt.FuseType = OssFsType
	for key, value := range req.VolumeContext {
		key = strings.ToLower(key)
		if key == "bucket" {
			opt.Bucket = strings.TrimSpace(value)
		} else if key == "url" {
			opt.URL = strings.TrimSpace(value)
		} else if key == "otheropts" {
			opt.OtherOpts = strings.TrimSpace(value)
		} else if key == "akid" {
			opt.AkID = strings.TrimSpace(value)
		} else if key == "aksecret" {
			opt.AkSecret = strings.TrimSpace(value)
		} else if key == "path" {
			if v := strings.TrimSpace(value); v == "" {
				opt.Path = "/"
			} else {
				opt.Path = v
			}
		} else if key == "usesharedpath" {
			if useSharedPath, err := strconv.ParseBool(value); err == nil {
				opt.UseSharedPath = useSharedPath
			} else {
				log.Warnf("invalid useSharedPath: %q", value)
			}
		} else if key == "authtype" {
			opt.AuthType = strings.ToLower(strings.TrimSpace(value))
		} else if key == "fusetype" {
			opt.FuseType = strings.ToLower(strings.TrimSpace(value))
		} else if key == "metricstop" {
			opt.MetricsTop = strings.ToLower(strings.TrimSpace(value))
		} else if key == "containernetworkfilesystem" {
			cnfsName = value
		} else if key == optDirectAssigned {
			opt.directAssigned, _ = strconv.ParseBool(strings.TrimSpace(value))
		} else if key == "encrypted" {
			opt.Encrypted = strings.ToLower(strings.TrimSpace(value))
		} else if key == "kmskeyid" {
			opt.KmsKeyId = value
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

	// Default oss path
	if opt.Path == "" {
		opt.Path = "/"
	}

	// support set ak by secret
	if opt.AkID == "" || opt.AkSecret == "" {
		if value, ok := req.Secrets[AkID]; ok {
			opt.AkID = value
		}
		if value, ok := req.Secrets[AkSecret]; ok {
			opt.AkSecret = value
		}
	}

	// check parameters
	if err := checkOssOptions(opt); err != nil {
		log.Errorf("Check oss input error: %s", err.Error())
		return nil, errors.New("Check oss input error: " + err.Error())
	}

	argStr := fmt.Sprintf("Bucket: %s, url: %s, , OtherOpts: %s, Path: %s, UseSharedPath: %s, authType: %s, encrypted: %s, kmsid: %s",
		opt.Bucket, opt.URL, opt.OtherOpts, opt.Path, strconv.FormatBool(opt.UseSharedPath), opt.AuthType, opt.Encrypted, opt.KmsKeyId)
	log.Infof("NodePublishVolume:: Starting Oss Mount: %s", argStr)

	if opt.directAssigned {
		return ns.publishDirectVolume(ctx, req, opt)
	}
	var ossMounter mountutils.Interface
	switch opt.FuseType {
	case JindoFsType:
		ossMounter = mounter.NewConnectorMounter(mountutils.New(""), "/etc/jindofs-tool/jindo-fuse")
	case OssFsType, "":
		// When useSharedPath options is set to false,
		// mount operations need to be atomic to ensure that no fuse pods are left behind in case of failure.
		// Because kubelet will not call NodeUnpublishVolume when NodePublishVolume never succeeded.
		authCfg := &mounter.AuthConfig{AuthType: opt.AuthType}
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

	switch opt.AuthType {
	case mounter.AuthTypeSTS:
		if opt.FuseType == OssFsType {
			mountOptions = append(mountOptions, GetRAMRoleOption())
		} else if opt.FuseType == JindoFsType {
			mountOptions = append(mountOptions, "fs.oss.provider.endpoint=ECS_ROLE")
		}
	default:
		if opt.FuseType == OssFsType {
			// ossfs fuse pod will mount the secret to access credentials
			err := mounter.SetupOssfsCredentialSecret(ctx, ns.clientset, ns.nodeName, req.VolumeId, opt.Bucket, opt.AkID, opt.AkSecret)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to setup ossfs credential secret: %v", err)
			}
		} else if opt.FuseType == JindoFsType {
			mountOptions = append(mountOptions, fmt.Sprintf("fs.oss.accessKeyId=%s,fs.oss.accessKeySecret=%s", opt.AkID, opt.AkSecret))
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

	regionID, _ := utils.GetRegionID()
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
		return errors.New("Oss Parametes error: Url/Bucket empty ")
	}

	if !strings.HasPrefix(opt.Path, "/") {
		return errors.New("Oss path error: start with " + opt.Path + ", should start with / ")
	}

	if opt.FuseType == JindoFsType {
		return nil
	}

	switch opt.AuthType {
	case mounter.AuthTypeSTS:
	default:
		// if not input ak from user, use the default ak value
		if opt.AkID == "" || opt.AkSecret == "" {
			ac := utils.GetEnvAK()
			opt.AkID = ac.AccessKeyID
			opt.AkSecret = ac.AccessKeySecret
		}
		if opt.AkID == "" || opt.AkSecret == "" {
			return errors.New("Oss Parametes error: AK and authType are both empty or invalid ")
		}
	}

	if opt.Encrypted != "" && opt.Encrypted != EncryptedTypeKms && opt.Encrypted != EncryptedTypeAes256 {
		return errors.New("Oss Encrypted error: invalid SSE encryted type ")
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
