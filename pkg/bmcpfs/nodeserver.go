//go:build !windows

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

package bmcpfs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
)

type nodeServer struct {
	common.GenericNodeServer
	mounter mount.Interface
	locks   *utils.VolumeLocks

	detachDelay      time.Duration
	unstageStartTime sync.Map // map[volumeID]time.Time
	// credsRoot is where per-volume EFC credential files live; a field so
	// tests can redirect it away from the /run/cnfs hostPath.
	credsRoot string
}

const (
	defaultAlinasMountProxySocket = "/run/cnfs/alinas-mounter.sock"
	metricsPathPrefix             = "/run/cnfs/efc/"

	// EFC mount options managed by this driver. The credential file options
	// are always generated from the driver-managed credentials directory and
	// stripped from user-supplied mount options.
	optionKeyAKFile  = "g_unas_AKFile"
	optionKeySTSFile = "g_unas_STSFile"

	// Access point mount option naming: the test-phase EFC client consumes
	// g_unas_Accesspoint, GA clients accept accesspoint. Switched via the
	// BMCPFS_AP_OPTION_STYLE env ("legacy" default, "ga").
	apOptionKeyLegacy = "g_unas_Accesspoint"
	apOptionKeyGA     = "accesspoint"
	apOptionStyleEnv  = "BMCPFS_AP_OPTION_STYLE"
	apOptionStyleGA   = "ga"
)

func newNodeServer(meta metadata.MetadataProvider) (*nodeServer, error) {
	var nodeID string
	data, err := os.ReadFile(metadata.LingjunConfigFile)
	if err != nil {
		if os.IsNotExist(err) {
			// if not lingjun instance
			isVsc, err := meta.Get(metadata.IsVscEnable)
			if err == nil && isVsc == "true" {
				instance, err := meta.Get(metadata.InstanceID)
				if err != nil {
					return nil, fmt.Errorf("get instance ID failed: %w", err)
				}
				nodeID = VSCNodeIDPrefix + instance
			} else {
				klog.InfoS("Not a lingjun instance", "isVsc", isVsc, "err", err, "nodeID", nodeID)
				nodeID = CommonNodeIDPrefix + os.Getenv(metadata.KUBE_NODE_NAME_ENV)
			}
		} else {
			return nil, fmt.Errorf("read lingjun_config file: %w", err)
		}
	} else {
		var lingjunConfig struct {
			NodeId string `json:"NodeId"`
		}
		if err := json.Unmarshal(data, &lingjunConfig); err != nil {
			return nil, fmt.Errorf("parse lingjun_config: %w", err)
		}
		if lingjunConfig.NodeId == "" {
			return nil, errors.New("unexpected lingjun_config: NodeId is empty")
		}
		nodeID = LingjunNodeIDPrefix + lingjunConfig.NodeId
	}
	klog.Infof("bmcpfsplugin nodeId: %s", nodeID)
	mounter := mounter.NewProxyMounter(defaultAlinasMountProxySocket, mount.NewWithoutSystemd(""))

	var detachDelay time.Duration
	if delay := os.Getenv("BMCPFS_DETACH_DELAY"); delay != "" {
		if delay, err := time.ParseDuration(delay); err == nil {
			detachDelay = delay
		} else {
			return nil, fmt.Errorf("parse BMCPFS_DETACH_DELAY: %w", err)
		}
	}
	return &nodeServer{
		GenericNodeServer: common.GenericNodeServer{NodeID: nodeID},
		locks:             utils.NewVolumeLocks(),
		mounter:           mounter,
		detachDelay:       detachDelay,
		credsRoot:         defaultCredentialsRoot,
	}, nil
}

func (ns *nodeServer) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: []*csi.NodeServiceCapability{
			{
				Type: &csi.NodeServiceCapability_Rpc{
					Rpc: &csi.NodeServiceCapability_RPC{
						Type: csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME,
					},
				},
			},
		},
	}, nil
}

func (ns *nodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	ns.unstageStartTime.Delete(req.VolumeId)   // in case NodeUnstageVolume never finished and next Pod comes
	return &csi.NodeStageVolumeResponse{}, nil // no-op
}

func (ns *nodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	if ns.detachDelay == 0 {
		if err := ns.cleanupCredentials(req.VolumeId); err != nil {
			return nil, status.Errorf(codes.Internal, "cleanup credentials: %v", err)
		}
		return &csi.NodeUnstageVolumeResponse{}, nil
	}
	actual, _ := ns.unstageStartTime.LoadOrStore(req.VolumeId, time.Now())
	startTime := actual.(time.Time)
	select {
	case <-time.After(time.Until(startTime.Add(ns.detachDelay))):
	case <-time.After(1 * time.Second):
		// Return fast, so the returning pod don't need to wait for the long timeout.
		// But still wait 1s to avoid busy loop.
		return nil, status.Error(codes.DeadlineExceeded, "delaying detach for possible reuse")
	case <-ctx.Done():
		return nil, status.Errorf(codes.DeadlineExceeded, "delaying detach for possible reuse: %v", ctx.Err())
	}
	ns.unstageStartTime.Delete(req.VolumeId)
	// All publishes for this volume on the node are gone; drop its EFC
	// credential files. Only done on the success path so a returning pod
	// within the detach delay keeps its credentials.
	if err := ns.cleanupCredentials(req.VolumeId); err != nil {
		return nil, status.Errorf(codes.Internal, "cleanup credentials: %v", err)
	}
	return &csi.NodeUnstageVolumeResponse{}, nil
}

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (
	*csi.NodePublishVolumeResponse, error) {
	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)

	notMounted, err := ns.mounter.IsLikelyNotMountPoint(req.TargetPath)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(req.TargetPath, os.ModePerm); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			notMounted = true
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	if !notMounted {
		// Republish (CSIDriver requiresRepublish): the mount point must not be
		// touched, only volume contents may change. Refresh the STS credential
		// file from the latest secrets so external rotation propagates.
		if err := ns.refreshCredentialsOnRepublish(req); err != nil {
			return nil, err
		}
		klog.InfoS("NodePublishVolume: target path is already mounted", "targetPath", req.TargetPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	var (
		mountOptions = req.VolumeCapability.GetMount().GetMountFlags()
		networkType  = req.PublishContext[_networkType]
		source       string
	)
	switch networkType {
	case networkTypeVPC:
		source = req.PublishContext[_vpcMountTarget]
		mountOptions = append(mountOptions, "net=tcp")
	case networkTypeVSC:
		mountOptions = append(mountOptions, "_netdev,net=vsc")
		source = req.PublishContext[_vscMountTarget]
	default:
		return nil, status.Errorf(codes.InvalidArgument, "unknown CPFS mountTarget networkType: %q", networkType)
	}
	if source == "" {
		return nil, status.Error(codes.InvalidArgument, "mountTarget is empty")
	}
	if path := req.VolumeContext[_path]; path != "" {
		source = fmt.Sprintf("%s:%s", source, path)
	}
	klog.InfoS("Mounting mount target", "targetPath", req.TargetPath, "source", source)

	// Credential file options are driver-managed; never accept user-supplied
	// paths for them.
	mountOptions, stripped := stripMountOptionKeys(mountOptions, optionKeyAKFile, optionKeySTSFile)
	if len(stripped) > 0 {
		klog.InfoS("NodePublishVolume: stripped driver-managed credential options from mount options", "volumeId", req.VolumeId, "keys", stripped)
	}

	accessPointID := req.VolumeContext[_accessPointID]
	mode, err := detectAuthMode(req.Secrets)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid secret for RAM auth: %v", err)
	}
	if mode != authModeNone && accessPointID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "RAM auth requires %q in volume attributes", _accessPointID)
	}
	if accessPointID != "" {
		mountOptions = append(mountOptions, accessPointMountOption(accessPointID))
	}
	switch mode {
	case authModeAK, authModeSTS:
		credsDir, err := credentialsDirForVolume(ns.credsRoot, req.VolumeId)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		if mode == authModeAK {
			akPath, err := writeAKFile(credsDir, req.Secrets)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "write AK credentials: %v", err)
			}
			mountOptions = append(mountOptions, optionKeyAKFile+"="+akPath)
		} else {
			stsPath, _, err := writeSTSFile(credsDir, req.Secrets)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "write STS credentials: %v", err)
			}
			mountOptions = append(mountOptions, optionKeySTSFile+"="+stsPath)
		}
		klog.V(2).InfoS("NodePublishVolume: prepared RAM auth credentials", "volumeId", req.VolumeId, "authMode", mode.String(), "accessPoint", accessPointID)
	}

	// Default g_lease_Enable=false unless user explicitly specified it
	// it turns on write caching for both data and metadata (backend support required), reducing read/write latency for small files. The risk is that it may increase the possibility of abnormal data loss in extreme cases.
	if !hasMountOption(mountOptions, "g_lease_Enable") {
		mountOptions = append(mountOptions, "g_lease_Enable=false")
	}

	mountOptions = append(mountOptions, "efc,protocol=efc,fstype=cpfs")
	err = ns.mounter.Mount(source, req.TargetPath, "alinas", mountOptions)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	utils.WriteMetricsInfo(metricsPathPrefix, req, "10", "efc", "cpfs", req.VolumeId)

	klog.InfoS("NodePublishVolume: succeeded to mount", "volumeId", req.VolumeId, "targetPath", req.TargetPath)
	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (
	*csi.NodeUnpublishVolumeResponse, error) {
	if !ns.locks.TryAcquire(req.VolumeId) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for %s", req.VolumeId)
	}
	defer ns.locks.Release(req.VolumeId)
	err := mount.CleanupMountPoint(req.TargetPath, ns.mounter, false)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cleanup mount point: %v", err)
	}
	klog.InfoS("NodeUnpublishVolume: succeeded to umount", "targetPath", req.TargetPath)
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

// hasMountOption checks whether any option in the slice has the given key prefix (e.g. "g_lease_Enable").
func hasMountOption(options []string, key string) bool {
	prefix := key + "="
	return slices.ContainsFunc(options, func(opt string) bool {
		return opt == key || strings.HasPrefix(opt, prefix)
	})
}

// accessPointMountOption renders the EFC access point option for the given
// AP ID, in the legacy (test phase) or GA naming depending on
// BMCPFS_AP_OPTION_STYLE.
func accessPointMountOption(accessPointID string) string {
	if os.Getenv(apOptionStyleEnv) == apOptionStyleGA {
		return apOptionKeyGA + "=" + accessPointID
	}
	return apOptionKeyLegacy + "=" + accessPointID
}

// stripMountOptionKeys removes options whose key matches any of the given
// keys, handling comma-joined compound entries. It returns the filtered
// options and the list of stripped keys.
func stripMountOptionKeys(options []string, keys ...string) (kept []string, stripped []string) {
	banned := make(map[string]struct{}, len(keys))
	for _, k := range keys {
		banned[k] = struct{}{}
	}
	for _, opt := range options {
		parts := strings.Split(opt, ",")
		keptParts := parts[:0]
		for _, part := range parts {
			key, _, _ := strings.Cut(part, "=")
			if _, ok := banned[key]; ok {
				stripped = append(stripped, key)
				continue
			}
			keptParts = append(keptParts, part)
		}
		if len(keptParts) > 0 {
			kept = append(kept, strings.Join(keptParts, ","))
		}
	}
	return kept, stripped
}

// refreshCredentialsOnRepublish handles the already-mounted NodePublishVolume
// path. The auth mode was fixed at the first publish and is recorded by which
// credential file exists; only STS-mode volumes are refreshed, so external
// Secret rotation propagates without touching the mount. Shape changes of the
// secret after mount are warned and ignored (a remount is required).
func (ns *nodeServer) refreshCredentialsOnRepublish(req *csi.NodePublishVolumeRequest) error {
	credsDir, err := credentialsDirForVolume(ns.credsRoot, req.VolumeId)
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	stsExists := fileExists(filepath.Join(credsDir, stsFileName))
	akExists := fileExists(filepath.Join(credsDir, akFileName))

	mode, err := detectAuthMode(req.Secrets)
	if err != nil {
		if stsExists {
			// The rotation source is broken; surface it so kubelet retries and
			// operators see the event before the mounted credential expires.
			return status.Errorf(codes.InvalidArgument, "invalid secret on republish: %v", err)
		}
		klog.ErrorS(err, "NodePublishVolume: ignoring invalid secret on republish of a non-STS volume", "volumeId", req.VolumeId)
		return nil
	}

	switch {
	case stsExists:
		if mode != authModeSTS {
			klog.InfoS("NodePublishVolume: secret shape changed after mount, ignoring; recreate the pod to apply", "volumeId", req.VolumeId, "mountedMode", authModeSTS.String(), "secretMode", mode.String())
			return nil
		}
		_, changed, err := writeSTSFile(credsDir, req.Secrets)
		if err != nil {
			return status.Errorf(codes.Internal, "refresh STS credentials: %v", err)
		}
		if changed {
			klog.V(2).InfoS("NodePublishVolume: refreshed STS credentials", "volumeId", req.VolumeId)
		}
	case akExists:
		if mode != authModeAK {
			klog.InfoS("NodePublishVolume: secret shape changed after mount, ignoring; recreate the pod to apply", "volumeId", req.VolumeId, "mountedMode", authModeAK.String(), "secretMode", mode.String())
		}
		// AK does not support hot update.
	default:
		if mode != authModeNone {
			klog.InfoS("NodePublishVolume: secret appeared after an unauthenticated mount, ignoring; recreate the pod to apply", "volumeId", req.VolumeId, "secretMode", mode.String())
		}
	}
	return nil
}

// cleanupCredentials removes the per-volume EFC credential directory.
func (ns *nodeServer) cleanupCredentials(volumeID string) error {
	credsDir, err := credentialsDirForVolume(ns.credsRoot, volumeID)
	if err != nil {
		// Publish would never have created a directory for such an ID.
		klog.ErrorS(err, "NodeUnstageVolume: skip credentials cleanup", "volumeId", volumeID)
		return nil
	}
	return os.RemoveAll(credsDir)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
