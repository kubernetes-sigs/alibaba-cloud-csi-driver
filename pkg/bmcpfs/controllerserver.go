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
	"strings"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	efloclient "github.com/alibabacloud-go/eflo-controller-20221215/v2/client"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/bmcpfs/internal"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type controllerServer struct {
	common.GenericControllerServer
	vscManager internal.VscManager
}

func newControllerServer(region string) (*controllerServer, error) {
	efloClient, err := newEfloClient(region)
	if err != nil {
		return nil, err
	}
	return &controllerServer{
		vscManager: internal.NewVscManager(efloClient),
	}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	if !strings.HasPrefix(req.NodeId, LingjunNodeIDPrefix) {
		klog.Info("Use VPC MountTarget for common node", "nodeId", req.NodeId)
		if req.VolumeContext[SERVER] == "" {
			return nil, status.Error(codes.InvalidArgument, "missing 'server' config in volume context")
		}
		return &csi.ControllerPublishVolumeResponse{
			PublishContext: map[string]string{
				SERVER: req.VolumeContext[SERVER],
			},
		}, nil
	}

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
		}, func() (done bool, err error) {
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
			if errors.Is(err, wait.ErrWaitTimeout) {
				return nil, status.Error(codes.DeadlineExceeded, "Failed to wait for the VSC to become Availabel")
			}
			return nil, err
		}
	}

	// Attach CPFS to VSC

	return nil, status.Error(codes.Unimplemented, "")
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
