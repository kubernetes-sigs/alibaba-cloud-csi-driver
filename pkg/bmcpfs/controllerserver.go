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
	"fmt"
	"os"
	"strconv"
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	efloclient "github.com/alibabacloud-go/eflo-controller-20221215/v2/client"
	nasclient "github.com/alibabacloud-go/nas-20170626/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/bmcpfs/internal"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
)

type controllerServer struct {
	common.GenericControllerServer
	vscManager     *internal.PrimaryVscManagerWithCache
	attachDetacher internal.CPFSAttachDetacher
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
		nasClient:      nasClient,
		skipDetach:     skipDetach,
	}, nil
}

func (cs *controllerServer) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: common.ControllerRPCCapabilities(
			csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
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
	// Get VscMountTarget of filesystem
	mt := req.VolumeContext[_vscMountTarget]
	if mt == "" {
		var err error
		mt, err = getMountTarget(cs.nasClient, cpfsID, networkTypeVSC)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get VscMountTarget: %v", err)
		}
	}

	// Get Primary vsc of Lingjun node
	lingjunInstanceId := strings.TrimPrefix(req.NodeId, LingjunNodeIDPrefix)
	if LingjunNodeIDPrefix == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid node id")
	}
	vscId, err := cs.vscManager.EnsurePrimaryVsc(ctx, lingjunInstanceId, false)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	klog.Info("Use VSC MountTarget for lingjun node", "nodeId", req.NodeId, "vscId", vscId)

	// Attach CPFS to VSC
	err = cs.attachDetacher.Attach(ctx, cpfsID, vscId)
	if err != nil {
		if autoSwitch, _ := strconv.ParseBool(req.VolumeContext[_mpAutoSwitch]); autoSwitch && internal.IsAttachNotSupportedError(err) {
			if req.VolumeContext[_vpcMountTarget] == "" {
				return nil, status.Errorf(codes.InvalidArgument, "missing %q config in volume context as vsc mountpoint not supported", _vpcMountTarget)
			}
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

	klog.InfoS("ControllerPublishVolume: attached cpfs to vsc", "vscMountTarget", mt, "vscId", vscId, "node", req.NodeId)
	return &csi.ControllerPublishVolumeResponse{
		PublishContext: map[string]string{
			_networkType:    networkTypeVSC,
			_vscId:          vscId,
			_vscMountTarget: mt,
		},
	}, nil
}

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (
	*csi.ControllerUnpublishVolumeResponse, error,
) {
	if !strings.HasPrefix(req.NodeId, LingjunNodeIDPrefix) || cs.skipDetach {
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}
	// Create Primary vsc for Lingjun node
	lingjunInstanceId := strings.TrimPrefix(req.NodeId, LingjunNodeIDPrefix)
	if LingjunNodeIDPrefix == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid node id")
	}
	vsc, err := cs.vscManager.GetPrimaryVscOf(lingjunInstanceId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
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
	cred, err := utils.GetCredentialV2()
	if err != nil {
		return nil, fmt.Errorf("init credential: %w", err)
	}
	config = config.SetCredential(cred)
	// set endpoint
	ep := os.Getenv("EFLO_CONTROLLER_ENDPOINT")
	if ep != "" {
		config = config.SetEndpoint(ep)
	} else {
		config = config.SetEndpoint(fmt.Sprintf("eflo-controller-vpc.%s.aliyuncs.com", region))
	}
	// set protocol
	scheme := strings.ToUpper(os.Getenv("ALICLOUD_CLIENT_SCHEME"))
	if strings.Contains(region, "test") {
		// must use HTTP in lingjun test regions
		scheme = "HTTP"
	}
	if scheme != "HTTP" {
		scheme = "HTTPS"
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

func getMountTarget(client *nasclient.Client, fsId, networkType string) (string, error) {
	resp, err := client.DescribeFileSystems(&nasclient.DescribeFileSystemsRequest{
		FileSystemId: &fsId,
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
			klog.V(2).InfoS("Found cpfs mount target", "filesystem", fsId, "networkType", networkType, "mountTarget", mountTarget, "status", status)
			if status == "Active" {
				return mountTarget, nil
			}
		}
	}
	return "", fmt.Errorf("no active %s mount target found", networkType)
}
