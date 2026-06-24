package cloud

import (
	"context"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	"github.com/alibabacloud-go/tea/dara"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

type ECSInterface interface {
	CreateDisk(request *ecs.CreateDiskRequest) (response *ecs.CreateDiskResponse, err error)
	DescribeAvailableResource(request *ecs.DescribeAvailableResourceRequest) (response *ecs.DescribeAvailableResourceResponse, err error)
	DescribeInstanceHistoryEvents(request *ecs.DescribeInstanceHistoryEventsRequest) (response *ecs.DescribeInstanceHistoryEventsResponse, err error)
	AttachDisk(request *ecs.AttachDiskRequest) (response *ecs.AttachDiskResponse, err error)
	DetachDisk(request *ecs.DetachDiskRequest) (response *ecs.DetachDiskResponse, err error)
	DeleteDisk(request *ecs.DeleteDiskRequest) (response *ecs.DeleteDiskResponse, err error)
	DescribeInstances(request *ecs.DescribeInstancesRequest) (response *ecs.DescribeInstancesResponse, err error)
	DescribeDisks(request *ecs.DescribeDisksRequest) (response *ecs.DescribeDisksResponse, err error)
	ResizeDisk(request *ecs.ResizeDiskRequest) (response *ecs.ResizeDiskResponse, err error)
	CreateSnapshot(request *ecs.CreateSnapshotRequest) (response *ecs.CreateSnapshotResponse, err error)
	DescribeSnapshots(request *ecs.DescribeSnapshotsRequest) (response *ecs.DescribeSnapshotsResponse, err error)
}

type ECSv2Interface interface {
	DescribeInstancesWithContext(ctx context.Context, request *ecs20140526.DescribeInstancesRequest, o *dara.RuntimeOptions) (response *ecs20140526.DescribeInstancesResponse, err error)
	DescribeInstanceTypesWithContext(ctx context.Context, request *ecs20140526.DescribeInstanceTypesRequest, o *dara.RuntimeOptions) (response *ecs20140526.DescribeInstanceTypesResponse, err error)
	DescribeAvailableResourceWithContext(ctx context.Context, request *ecs20140526.DescribeAvailableResourceRequest, o *dara.RuntimeOptions) (response *ecs20140526.DescribeAvailableResourceResponse, err error)
	DescribeDisksWithContext(ctx context.Context, request *ecs20140526.DescribeDisksRequest, o *dara.RuntimeOptions) (response *ecs20140526.DescribeDisksResponse, err error)
	TagResourcesWithContext(ctx context.Context, request *ecs20140526.TagResourcesRequest, o *dara.RuntimeOptions) (response *ecs20140526.TagResourcesResponse, err error)
	UntagResourcesWithContext(ctx context.Context, request *ecs20140526.UntagResourcesRequest, o *dara.RuntimeOptions) (response *ecs20140526.UntagResourcesResponse, err error)
	ModifyDiskSpecWithContext(ctx context.Context, request *ecs20140526.ModifyDiskSpecRequest, o *dara.RuntimeOptions) (response *ecs20140526.ModifyDiskSpecResponse, err error)
	ModifyDiskAttributeWithContext(ctx context.Context, request *ecs20140526.ModifyDiskAttributeRequest, o *dara.RuntimeOptions) (response *ecs20140526.ModifyDiskAttributeResponse, err error)
	DescribeTasksWithContext(ctx context.Context, request *ecs20140526.DescribeTasksRequest, o *dara.RuntimeOptions) (response *ecs20140526.DescribeTasksResponse, err error)
}
