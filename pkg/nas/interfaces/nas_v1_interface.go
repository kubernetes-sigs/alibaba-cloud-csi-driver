package interfaces

import "github.com/aliyun/alibaba-cloud-sdk-go/services/nas"

type NasV1Interface interface {
	CreateFileSystem(request *nas.CreateFileSystemRequest) (response *nas.CreateFileSystemResponse, err error)
	CreateMountTarget(request *nas.CreateMountTargetRequest) (response *nas.CreateMountTargetResponse, err error)
	DeleteFileSystem(request *nas.DeleteFileSystemRequest) (response *nas.DeleteFileSystemResponse, err error)
	DeleteMountTarget(request *nas.DeleteMountTargetRequest) (response *nas.DeleteMountTargetResponse, err error)
	DescribeFileSystems(request *nas.DescribeFileSystemsRequest) (response *nas.DescribeFileSystemsResponse, err error)
	DescribeMountTargets(request *nas.DescribeMountTargetsRequest) (response *nas.DescribeMountTargetsResponse, err error)
	TagResources(request *nas.TagResourcesRequest) (response *nas.TagResourcesResponse, err error)
}
