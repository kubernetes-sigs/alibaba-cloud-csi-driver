package oss

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	ossfpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager/oss"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/kata/directvolume"
	"k8s.io/klog/v2"
)

const (
	optDirectAssigned  = "direct"
	optAnnotations     = "annotations"
	optEncrypted       = "encrypted"
	optEncPasswd       = "encPasswd"
	optKmsKeyId        = "kmsKeyId"
	fsTypeSecureMount  = "secure_mount"
	volumeTypeOSS      = "alibaba-cloud-oss"
	sealedSecretPrefix = "sealed."
)

func (ns *nodeServer) publishDirectVolume(ctx context.Context, req *csi.NodePublishVolumeRequest, opt *ossfpm.Options) (*csi.NodePublishVolumeResponse, error) {
	logger := klog.FromContext(ctx).WithValues(
		"nodeServer", "OSSNodeServer",
		"volumeID", req.VolumeId,
		"target", req.TargetPath,
	)

	volumePath := req.TargetPath
	if isDirectVolumePath(volumePath) {
		logger.Info("NodePublishVolume: The mount info for DirectVolume already exist")
		return &csi.NodePublishVolumeResponse{}, nil
	}

	device := req.TargetPath
	volumeType := volumeTypeOSS
	fsType := fsTypeSecureMount
	var encrypted, encPasswd, kmsKeyId string
	annotationsObj := map[string]string{}

	if v := req.Secrets[optEncPasswd]; v != "" {
		encPasswd = v
	}
	for key, value := range req.VolumeContext {
		switch key {
		case optEncrypted:
			encrypted = strings.TrimSpace(value)
		case optEncPasswd:
			encPasswd = strings.TrimSpace(value)
		case optKmsKeyId:
			kmsKeyId = strings.TrimSpace(value)
		case optAnnotations:
			_ = json.Unmarshal([]byte(value), &annotationsObj)
		}
	}

	if annotationsObj["kata_fs_type"] != "" {
		fsType = annotationsObj["kata_fs_type"]
	}
	if annotationsObj["kata_volume_type"] != "" {
		volumeType = annotationsObj["kata_volume_type"]
	}
	if annotationsObj["kata_device"] != "" {
		device = annotationsObj["kata_device"]
	}
	annotations, _ := json.Marshal(annotationsObj)
	annotationsStr := string(annotations)
	if len(annotations) == 0 || annotationsStr == "null" || annotationsStr == "" {
		annotationsStr = "{}"
	}
	metadata := map[string]string{
		"bucket":      opt.Bucket,
		"url":         opt.URL,
		"otherOpts":   opt.OtherOpts,
		"path":        opt.Path,
		"encrypted":   encrypted,
		"kmsKeyId":    kmsKeyId,
		"annotations": annotationsStr,
		"volumeId":    req.GetVolumeId(),
		"readonly":    fmt.Sprintf("%v", req.Readonly),
		"targetPath":  req.GetTargetPath(),
	}
	if opt.AkSecret != "" && strings.HasPrefix(opt.AkSecret, sealedSecretPrefix) {
		// `AkID` and `AkSecret` are fixed for the protocol
		metadata[AkID] = opt.AkID
		metadata[AkSecret] = opt.AkSecret
	}
	if encPasswd != "" && strings.HasPrefix(encPasswd, sealedSecretPrefix) {
		metadata[optEncPasswd] = encPasswd
	}

	mountInfo := directvolume.MountInfo{
		VolumeType: volumeType,
		Device:     device,
		FsType:     fsType,
		Metadata:   metadata,
		Options:    strings.Fields(opt.OtherOpts),
	}

	logger.Info("NodePublishVolume:: Starting add mount info for DirectVolume")
	err := directvolume.AddMountInfo(volumePath, mountInfo)
	if err != nil {
		logger.Error(err, "NodePublishVolume:: Add mount info for DirectVolume failed")
		return nil, err
	}
	logger.Info("NodePublishVolume:: Add mount info for DirectVolume is successfully")

	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) unPublishDirectVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	logger := klog.FromContext(ctx).WithValues(
		"nodeServer", "OSSNodeServer",
		"volumeID", req.VolumeId,
		"target", req.TargetPath,
	)

	volumePath := req.TargetPath
	err := directvolume.Remove(volumePath)
	if err != nil {
		logger.Error(err, "NodeUnPublishVolume:: Remove mount info for DirectVolume failed")
		return nil, err
	}

	logger.Info("NodeUnPublishVolume:: Remove mount info for DirectVolume is successfully")
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func isDirectVolumePath(volumePath string) bool {
	_, err := directvolume.VolumeMountInfo(volumePath)
	return err == nil
}
