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
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	efloclient "github.com/alibabacloud-go/eflo-controller-20221215/v3/client"
	nasclient "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	alicred_old "github.com/aliyun/credentials-go/credentials"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/bmcpfs/internal"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/credentials"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
)

type controllerServer struct {
	common.GenericControllerServer
	vscManager     *internal.PrimaryVscManagerWithCache
	attachDetacher internal.CPFSAttachDetacher
	filesetManager internal.CPFSFileSetManager
	locks          *utils.VolumeLocks
	nasClient      *nasclient.Client
	skipDetach     bool
}

func newControllerServer(region string) (*controllerServer, error) {
	efloClient, err := newEfloClient(region)
	if err != nil {
		return nil, err
	}

	skipDetach := false
	if skipDetachVal, _ := strconv.ParseBool(os.Getenv("SKIP_BMCPFS_DETACH")); skipDetachVal {
		skipDetach = skipDetachVal
	}

	nasClient, err := cloud.NewNasClientV2(region)
	if err != nil {
		return nil, err
	}

	return &controllerServer{
		vscManager:     internal.NewPrimaryVscManagerWithCache(efloClient),
		attachDetacher: internal.NewCPFSAttachDetacher(nasClient),
		filesetManager: internal.NewCPFSFileSetManager(nasClient),
		locks:          utils.NewVolumeLocks(),
		nasClient:      nasClient,
		skipDetach:     skipDetach,
	}, nil
}

// CreateVolume ...
func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	logger := klog.FromContext(ctx)

	if !cs.locks.TryAcquire(req.Name) {
		return nil, status.Errorf(codes.Aborted, "There is already an operation for volume %s", req.Name)
	}
	defer cs.locks.Release(req.Name)

	// Validate parameters
	if err := validateFileSetParameters(req); err != nil {
		logger.Error(err, "failed to validate parameters", "volumeName", req.Name)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid parameters from input: %v, with error: %v", req.Name, err)
	}

	// Extract parameters
	params := req.GetParameters()
	bmcpfsID := params["bmcpfsId"]
	if bmcpfsID == "" {
		return nil, status.Error(codes.InvalidArgument, "bmcpfsId parameter is required")
	}

	// Construct volume path
	volumeName := req.Name
	var fullPath string
	if rootPath, ok := params["path"]; ok && rootPath != "" {
		fullPath = filepath.Join(rootPath, volumeName)
	} else {
		fullPath = filepath.Join("/", volumeName)
	}
	// filepath.Join will skip the tailing slash we need to add it manually
	fullPath = fullPath + "/"
	klog.InfoS("CreateVolume: Constructing volume path", "fullpath", fullPath)

	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())

	// Prepare volume context
	volumeContext := req.GetParameters()
	if volumeContext == nil {
		volumeContext = make(map[string]string)
	}
	volumeContext[_path] = fullPath
	volumeContext = updateVolumeContext(volumeContext)

	// Check if fileset already exists as nas doesn't provide idempotent create fileset API
	fileset, err := cs.filesetManager.DescribeFilesetByFilePath(ctx, bmcpfsID, fullPath)
	if err == nil && fileset != nil {
		tmpVol := createVolumeResponse(*fileset.FsetId, bmcpfsID, volSizeBytes, volumeContext)
		return &csi.CreateVolumeResponse{Volume: tmpVol}, nil
	}

	// If the error indicates the fileset was not found, continue to create it
	// Otherwise, return the error
	if err != nil && !errors.Is(err, internal.ErrFilesetNotFound) {
		klog.ErrorS(err, "CreateVolume: failed to describe fileset", "filesystemID", bmcpfsID, "filesetPath", fullPath)
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Create fileset
	fileSetID, err := cs.filesetManager.CreateFileSet(ctx, bmcpfsID, req.Name, fullPath, 1000000, volSizeBytes, false)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = wait.PollUntilContextCancel(ctx, time.Second*2, true, func(ctx context.Context) (bool, error) {
		var err error
		resp, err := cs.filesetManager.DescribeFilesetByFsetID(ctx, bmcpfsID, fileSetID)
		if err == nil {
			if resp.Status != nil && *resp.Status == "CREATED" {
				klog.InfoS("CreateVolume: fileset created successfully", "filesetID", fileSetID, "filesystemID", bmcpfsID)
				return true, nil
			}
			klog.InfoS("CreateVolume: fileset status incorrected", "filesetID", fileSetID, "filesystemID", bmcpfsID, "status", *resp.Status)
			return false, nil
		}
		return false, err
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "CreateVolume: failed to describe fileset %s: %s", fileSetID, err.Error())
	}

	klog.Infof("CreateVolume: Successfully created FileSet %s: id[%s], filesystem[%s], path[%s]", req.GetName(), fileSetID, bmcpfsID, fullPath)

	tmpVol := createVolumeResponse(fileSetID, bmcpfsID, volSizeBytes, volumeContext)

	return &csi.CreateVolumeResponse{Volume: tmpVol}, nil
}

func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(2).Info("starting")

	// Parse volume ID to extract filesystem ID and fileset ID
	fsID, fileSetID, err := parseVolumeID(req.VolumeId)
	if err != nil {
		klog.Errorf("DeleteVolume: failed to parse volume ID %s: %v", req.VolumeId, err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	klog.Infof("DeleteVolume: deleting fileset %s from filesystem %s", fileSetID, fsID)

	// Delete the fileset
	err = cs.filesetManager.DeleteFileSet(ctx, fsID, fileSetID)
	if err != nil {
		klog.Errorf("DeleteVolume: failed to delete fileset %s from filesystem %s: %v", fileSetID, fsID, err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	klog.Infof("DeleteVolume: successfully deleted fileset %s from filesystem %s", fileSetID, fsID)
	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: common.ControllerRPCCapabilities(
			csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
			csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		),
	}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	if !strings.HasPrefix(req.NodeId, LingjunNodeIDPrefix) {
		if req.VolumeContext[_vpcMountTarget] == "" {
			return nil, status.Errorf(codes.InvalidArgument, "missing %q config in volume context", _vpcMountTarget)
		}
		// TODO: try to use existing vpc mount target
		klog.InfoS("ControllerPublishVolume: use VPC MountTarget", "nodeId", req.NodeId)
		return &csi.ControllerPublishVolumeResponse{
			PublishContext: map[string]string{
				_networkType:    networkTypeVPC,
				_vpcMountTarget: req.VolumeContext[_vpcMountTarget],
			},
		}, nil
	}

	cpfsID, _ := parseVolumeHandle(req.VolumeId)

	mt := req.VolumeContext[_vscMountTarget]
	if mt == "" {
		var err error
		mt, err = getMountTarget(cs.nasClient, cpfsID, networkTypeVSC)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get VscMountTarget: %v", err)
		}
	}

	// Get Primary vsc of Lingjun node
	lingjunInstanceID := strings.TrimPrefix(req.NodeId, LingjunNodeIDPrefix)
	if LingjunNodeIDPrefix == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid node id")
	}
	vscID, err := cs.vscManager.EnsurePrimaryVsc(ctx, lingjunInstanceID, false)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	klog.Info("Use VSC MountTarget for lingjun node", "nodeId", req.NodeId, "vscId", vscID)

	// Attach CPFS to VSC
	err = cs.attachDetacher.Attach(ctx, cpfsID, vscID)
	if err != nil {
		if autoSwitch, _ := strconv.ParseBool(req.VolumeContext[_mpAutoSwitch]); autoSwitch && internal.IsAttachNotSupportedError(err) {
			if req.VolumeContext[_vpcMountTarget] == "" {
				return nil, status.Errorf(codes.InvalidArgument, "missing %q config in volume context as vsc mountpoint not supported", _vpcMountTarget)
			}
			klog.InfoS("ControllerPublishVolume: vsc mountpoint not supported, switch to vpc mountpoint", "nodeId", req.NodeId, "mt", mt)
			return &csi.ControllerPublishVolumeResponse{
				PublishContext: map[string]string{
					_networkType:    networkTypeVPC,
					_vpcMountTarget: req.VolumeContext[_vpcMountTarget],
				},
			}, nil
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// TODO: if the cached vscid is already deleted, try to recreate a new primary vsc for lingjun node

	klog.InfoS("ControllerPublishVolume: attached cpfs to vsc", "vscMountTarget", mt, "vscId", vscID, "node", req.NodeId)
	return &csi.ControllerPublishVolumeResponse{
		PublishContext: map[string]string{
			_networkType:    networkTypeVSC,
			_vscID:          vscID,
			_vscMountTarget: mt,
		},
	}, nil
}

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	if !strings.HasPrefix(req.NodeId, LingjunNodeIDPrefix) || cs.skipDetach {
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}
	// Create Primary vsc for Lingjun node
	lingjunInstanceID := strings.TrimPrefix(req.NodeId, LingjunNodeIDPrefix)
	if LingjunNodeIDPrefix == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid node id")
	}
	vsc, err := cs.vscManager.GetPrimaryVscOf(lingjunInstanceID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get vsc error: %v", err)
	}
	if vsc == nil {
		klog.InfoS("ControllerUnpublishVolume: skip detaching cpfs from vsc as vsc not found", "node", req.NodeId)
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}

	// If `req.VolumeId` is a combination of `cpfsID` and `fsetID`, Detach will trigger an error.
	err = cs.attachDetacher.Detach(ctx, req.VolumeId, vsc.VscID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	klog.InfoS("ControllerUnpublishVolume: detached cpfs from vsc", "node", req.NodeId, "filesystem", req.VolumeId)
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

// KubernetesAlicloudIdentity is the user agent string for Eflo client
var KubernetesAlicloudIdentity = fmt.Sprintf("Kubernetes.Alicloud/CsiProvision.Bmcpfs-%s", version.VERSION)

const efloConnTimeout = 10

func newEfloClient(region string) (*efloclient.Client, error) {
	// lingjun region could be different from ack region
	// use another environment variable EFLO_CONTROLLER_REGION for special lingjun pre regions
	if r := os.Getenv("EFLO_CONTROLLER_REGION"); r != "" {
		region = r
	}
	config := new(openapi.Config).
		SetUserAgent(KubernetesAlicloudIdentity).
		SetRegionId(region).
		SetConnectTimeout(efloConnTimeout).
		SetGlobalParameters(&openapi.GlobalParameters{
			Queries: map[string]*string{
				"RegionId": &region,
			},
		})
	// set credential
	provider, err := credentials.NewProvider()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch credential: %w", err)
	}
	config = config.SetCredential(alicred_old.FromCredentialsProvider(provider.GetProviderName(), provider))
	// set endpoint
	ep := os.Getenv("EFLO_CONTROLLER_ENDPOINT")
	if ep != "" {
		config = config.SetEndpoint(ep)
	} else {
		config = config.SetEndpoint(fmt.Sprintf("eflo-controller-vpc.%s.aliyuncs.com", region))
	}
	// set protocol
	scheme := "HTTPS"
	if strings.Contains(region, "test") {
		// must use HTTP in lingjun test regions
		scheme = "HTTP"
	}
	if e := os.Getenv("ALICLOUD_CLIENT_SCHEME"); e != "" {
		scheme = e
	}
	config = config.SetProtocol(scheme)
	// init client
	return efloclient.NewClient(config)
}

func parseVolumeHandle(volumeHandle string) (string, string) {
	parts := strings.Split(volumeHandle, volumeHandleDelimiter)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return parts[0], ""
}

func getMountTarget(client *nasclient.Client, fsID, networkType string) (string, error) {
	resp, err := client.DescribeFileSystems(&nasclient.DescribeFileSystemsRequest{
		FileSystemId: &fsID,
	})
	if err != nil {
		return "", fmt.Errorf("nas:DescribeFileSystems failed: %w", err)
	}
	klog.V(4).InfoS("nas:DescribeFileSystems succeeded", "response", resp.Body)
	filesystems := resp.Body.FileSystems
	if filesystems == nil || len(filesystems.FileSystem) == 0 || filesystems.FileSystem[0] == nil {
		return "", nil
	}
	mountTargets := filesystems.FileSystem[0].MountTargets
	if mountTargets == nil {
		return "", nil
	}
	for _, mt := range mountTargets.MountTarget {
		t := tea.StringValue(mt.NetworkType)
		// mount targets with empty network type is vsc type
		if t == "" {
			t = networkTypeVSC
		}
		if t == networkType {
			mountTarget := tea.StringValue(mt.MountTargetDomain)
			status := tea.StringValue(mt.Status)
			klog.V(2).InfoS("Found cpfs mount target", "filesystem", fsID, "networkType", networkType, "mountTarget", mountTarget, "status", status)
			if status == "Active" {
				return mountTarget, nil
			}
		}
	}
	return "", fmt.Errorf("no active %s mount target found", networkType)
}
