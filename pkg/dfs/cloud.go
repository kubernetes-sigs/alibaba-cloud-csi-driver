/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http:// www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dfs

import (
	"fmt"
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dfs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"k8s.io/klog/v2"
)

type VSC = dfs.Vsc

const (
	VSC_NORMARL  = "NORMAL"
	VSC_CREATING = "CREATING"
	VSC_INVALID  = "INVALID"

	VSCType_Primary = "Primary"
)

type DFSClient interface {
	Attach(fileSystemId, mountPointId, instanceId string) error
	Detach(fileSystemId, mountPointId, instanceId string) error
	DescribeAttachment(fileSystemId, mountPointId, instanceId string) (*VSC, error)
}

func initDFSEndpoint(region string) {
	ep := os.Getenv("DFS_ENDPOINT")
	if ep == "" {
		ep = fmt.Sprintf("dfs-vpc.%s.aliyuncs.com", region)
	}
	_ = endpoints.AddEndpointMapping(region, "Dfs", ep)
	klog.Infof("DFS endpoint: %s", ep)
}

type dfsClient struct {
	region string
	client *dfs.Client
}

func newDFSClient(region string) (DFSClient, error) {
	ac := utils.GetAccessControl()
	client, err := dfs.NewClientWithOptions(region, ac.Config, ac.Credential)
	if err != nil {
		return nil, err
	}
	return &dfsClient{
		region: region,
		client: client,
	}, nil
}

func (d *dfsClient) Attach(fileSystemId string, mountPointId string, instanceId string) error {
	attachRequest := dfs.CreateAttachVscMountPointRequest()
	attachRequest.InputRegionId = d.region
	attachRequest.FileSystemId = fileSystemId
	attachRequest.MountPointId = mountPointId
	attachRequest.InstanceIds = fmt.Sprintf("[%q]", instanceId)
	attachResponse, err := d.client.AttachVscMountPoint(attachRequest)
	if err != nil {
		return fmt.Errorf("AttachVscMountPoint failed: %w", err)
	}
	klog.InfoS("DFS AttachVscMountPoint succeeded", "request", *attachRequest, "requestId", attachResponse.RequestId)
	return nil
}

func (d *dfsClient) Detach(fileSystemId string, mountPointId string, instanceId string) error {
	detachRequest := dfs.CreateDetachVscMountPointRequest()
	detachRequest.InputRegionId = d.region
	detachRequest.FileSystemId = fileSystemId
	detachRequest.MountPointId = mountPointId
	detachRequest.InstanceIds = fmt.Sprintf("[%q]", instanceId)
	detachResponse, err := d.client.DetachVscMountPoint(detachRequest)
	if err != nil {
		return fmt.Errorf("DetachVscMountPoint failed: %w", err)
	}
	klog.InfoS("DFS DetachVscMountPoint succeeded", "request", *detachRequest, "requestId", detachResponse.RequestId)
	return nil
}

func (d *dfsClient) DescribeAttachment(fileSystemId string, mountPointId string, instanceId string) (*VSC, error) {
	describeRequest := dfs.CreateDescribeVscMountPointsRequest()
	describeRequest.InputRegionId = d.region
	describeRequest.FileSystemId = fileSystemId
	describeRequest.MountPointId = mountPointId
	describeRequest.InstanceId = instanceId
	describeResponse, err := d.client.DescribeVscMountPoints(describeRequest)
	if err != nil {
		return nil, fmt.Errorf("DescribeVscMountPoints failed: %w", err)
	}
	klog.V(4).InfoS("DFS DescribeVscMountPoints succeeded", "requestId", describeResponse.RequestId)

	if len(describeResponse.MountPoints) != 1 || len(describeResponse.MountPoints[0].Instances) != 1 {
		return nil, fmt.Errorf("Unexpected DescribeVscMountPoints response: %v", *describeResponse)
	}

	for _, vsc := range describeResponse.MountPoints[0].Instances[0].Vscs {
		if vsc.VscType == VSCType_Primary {
			return &vsc, nil
		}
	}
	return nil, nil
}
