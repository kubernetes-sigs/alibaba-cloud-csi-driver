package cloud

import "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

type ECSInterface interface {
	CreateDisk(request *ecs.CreateDiskRequest) (response *ecs.CreateDiskResponse, err error)
	DescribeAvailableResource(request *ecs.DescribeAvailableResourceRequest) (response *ecs.DescribeAvailableResourceResponse, err error)
	DeleteDisk(request *ecs.DeleteDiskRequest) (response *ecs.DeleteDiskResponse, err error)
	DescribeInstances(request *ecs.DescribeInstancesRequest) (response *ecs.DescribeInstancesResponse, err error)
	DescribeInstanceTypes(request *ecs.DescribeInstanceTypesRequest) (response *ecs.DescribeInstanceTypesResponse, err error)
	DescribeDisks(request *ecs.DescribeDisksRequest) (response *ecs.DescribeDisksResponse, err error)
	ResizeDisk(request *ecs.ResizeDiskRequest) (response *ecs.ResizeDiskResponse, err error)
	DescribeSnapshots(request *ecs.DescribeSnapshotsRequest) (response *ecs.DescribeSnapshotsResponse, err error)
}
