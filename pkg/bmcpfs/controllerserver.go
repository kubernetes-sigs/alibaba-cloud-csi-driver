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
	"strings"
	"time"

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
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
)

type controllerServer struct {
	common.GenericControllerServer
	vscManager     internal.VscManager
	attachDetacher internal.CPFSAttachDetacher
	nasClient      *nasclient.Client
}

func newControllerServer(region string) (*controllerServer, error) {
	efloClient, err := newEfloClient(region)
	if err != nil {
		return nil, err
	}
	nasClient, err := cloud.NewNasClientV2(region)
	if err != nil {
		return nil, err
	}
	return &controllerServer{
		vscManager:     internal.NewVscManager(efloClient),
		attachDetacher: internal.NewCPFSAttachDetacher(nasClient),
		nasClient:      nasClient,
	}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	// TODO:  set default timeout for context
	if !strings.HasPrefix(req.NodeId, LingjunNodeIDPrefix) {
		klog.Info("Use VPC MountTarget for common node", "nodeId", req.NodeId)
		if req.VolumeContext[_vpcMountTarget] == "" {
			return nil, status.Errorf(codes.InvalidArgument, "missing %q config in volume context", _vpcMountTarget)
		}
		// TODO: try to use existing vpc mount target
		// getMountTarget(cs.nasClient, req.VolumeId, "vpc")
		return &csi.ControllerPublishVolumeResponse{
			PublishContext: map[string]string{
				_vpcMountTarget: req.VolumeContext[_vpcMountTarget],
			},
		}, nil
	}

	// Get VscMountTarget of filesystem
	mt := req.VolumeContext[_vscMountTarget]
	if mt == "" {
		var err error
		mt, err = getMountTarget(cs.nasClient, req.VolumeId, "")
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get VscMountTarget: %v", err)
		}
	}

	// Create Primary vsc for Lingjun node
	lingjunInstanceId := strings.TrimPrefix(req.NodeId, LingjunNodeIDPrefix)
	if LingjunNodeIDPrefix == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid node id")
	}
	// TODO: add cache as we only need to create once for a node
	vsc, err := cs.vscManager.GetPrimaryVscOf(lingjunInstanceId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var vscId string
	if vsc == nil {
		// create vsc
		vscId, err = cs.vscManager.CreatePrimaryVscFor(lingjunInstanceId)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	} else {
		vscId = vsc.VscID
		if vsc.Status != internal.VscStatusAvailable && vsc.Status != internal.VscStatusCreating {
			return nil, status.Errorf(codes.Internal, "unexpected vsc status: %v", vsc.Status)
		}
	}
	klog.Info("Use VSC MountTarget for lingjun node", "nodeId", req.NodeId, "vscId", vscId)
	if vsc == nil || vsc.Status != internal.VscStatusAvailable {
		err = wait.ExponentialBackoffWithContext(ctx, wait.Backoff{
			Duration: time.Millisecond * 400,
			Factor:   2,
			Steps:    6,
		}, func(ctx context.Context) (done bool, err error) {
			vsc, err := cs.vscManager.GetVsc(vscId)
			if err != nil {
				return false, status.Error(codes.Internal, err.Error())
			}
			switch vsc.Status {
			case internal.VscStatusAvailable:
				return true, nil
			case internal.VscStatusCreating:
				return false, nil
			default:
				return false, status.Errorf(codes.Internal, "unexpected vsc status: %v", vsc.Status)
			}
		})
		if err != nil {
			if wait.Interrupted(err) {
				return nil, status.Error(codes.DeadlineExceeded, "Failed to wait for the VSC to become Availabel")
			}
			return nil, err
		}
	}

	// Attach CPFS to VSC
	err = cs.attachDetacher.Attach(ctx, req.VolumeId, vscId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	klog.InfoS("Succeeded to attach CPFS to VSC", "vscMountTarget", mt, "vscId", vscId, "node", req.NodeId)
	return &csi.ControllerPublishVolumeResponse{
		PublishContext: map[string]string{
			_vscId:          vscId,
			_vscMountTarget: mt,
		},
	}, nil
}

var KubernetesAlicloudIdentity = fmt.Sprintf("Kubernetes.Alicloud/CsiProvision.Bmcpfs-%s", version.VERSION)

func newEfloClient(region string) (*efloclient.Client, error) {
	config := new(openapi.Config).
		SetUserAgent(KubernetesAlicloudIdentity).
		SetRegionId(region).
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
	ep := os.Getenv("EFLO_ENDPOINT")
	if ep != "" {
		config = config.SetEndpoint(ep)
	}
	// set protocol
	scheme := strings.ToUpper(os.Getenv("ALICLOUD_CLIENT_SCHEME"))
	if scheme != "HTTP" {
		scheme = "HTTPS"
	}
	config = config.SetProtocol(scheme)
	// init client
	return efloclient.NewClient(config)
}

func getMountTarget(client *nasclient.Client, fsId, networkType string) (string, error) {
	resp, err := client.DescribeFileSystems(&nasclient.DescribeFileSystemsRequest{
		FileSystemId: &fsId,
	})
	if err != nil {
		return "", fmt.Errorf("nas:DescribeFileSystems failed: %w", err)
	}
	klog.V(2).InfoS("nas:DescribeFileSystems succeeded", "response", resp.Body)
	filesystems := resp.Body.FileSystems
	if filesystems == nil || len(filesystems.FileSystem) == 0 || filesystems.FileSystem[0] == nil {
		return "", nil
	}
	mountTargets := filesystems.FileSystem[0].MountTargets
	if mountTargets == nil {
		return "", nil
	}
	for _, mt := range mountTargets.MountTarget {
		if tea.StringValue(mt.NetworkType) == networkType {
			if tea.StringValue(mt.Status) == "Active" {
				return tea.StringValue(mt.MountTargetDomain), nil
			}
		}
	}
	return "", fmt.Errorf("no active %s mount target found", networkType)
}
