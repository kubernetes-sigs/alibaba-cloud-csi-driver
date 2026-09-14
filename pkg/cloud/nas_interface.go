package cloud

import nas "github.com/alibabacloud-go/nas-20170626/v4/client"

type NasInterface interface {
	CancelDirQuota(request *nas.CancelDirQuotaRequest) (*nas.CancelDirQuotaResponse, error)
	CreateAccessPoint(request *nas.CreateAccessPointRequest) (*nas.CreateAccessPointResponse, error)
	CreateAgenticSpace(request *nas.CreateAgenticSpaceRequest) (*nas.CreateAgenticSpaceResponse, error)
	CreateDir(request *nas.CreateDirRequest) (*nas.CreateDirResponse, error)
	DeleteAccessPoint(request *nas.DeleteAccessPointRequest) (*nas.DeleteAccessPointResponse, error)
	DeleteAgenticSpace(request *nas.DeleteAgenticSpaceRequest) (*nas.DeleteAgenticSpaceResponse, error)
	DescribeAccessPoint(request *nas.DescribeAccessPointRequest) (*nas.DescribeAccessPointResponse, error)
	DescribeFileSystems(request *nas.DescribeFileSystemsRequest) (*nas.DescribeFileSystemsResponse, error)
	GetAgenticSpace(request *nas.GetAgenticSpaceRequest) (*nas.GetAgenticSpaceResponse, error)
	GetRecycleBinAttribute(request *nas.GetRecycleBinAttributeRequest) (*nas.GetRecycleBinAttributeResponse, error)
	ListAccessPoints(request *nas.ListAccessPointsRequest) (*nas.ListAccessPointsResponse, error)
	SetAgenticSpaceQuota(request *nas.SetAgenticSpaceQuotaRequest) (*nas.SetAgenticSpaceQuotaResponse, error)
	SetDirQuota(request *nas.SetDirQuotaRequest) (*nas.SetDirQuotaResponse, error)
}
