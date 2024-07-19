package interfaces

import nas "github.com/alibabacloud-go/nas-20170626/v3/client"

type NasInterface interface {
	CancelDirQuota(request *nas.CancelDirQuotaRequest) (*nas.CancelDirQuotaResponse, error)
	CreateAccessPoint(request *nas.CreateAccessPointRequest) (*nas.CreateAccessPointResponse, error)
	CreateDir(request *nas.CreateDirRequest) (*nas.CreateDirResponse, error)
	DeleteAccessPoint(request *nas.DeleteAccessPointRequest) (*nas.DeleteAccessPointResponse, error)
	DescribeAccessPoint(request *nas.DescribeAccessPointRequest) (*nas.DescribeAccessPointResponse, error)
	GetRecycleBinAttribute(request *nas.GetRecycleBinAttributeRequest) (*nas.GetRecycleBinAttributeResponse, error)
	SetDirQuota(request *nas.SetDirQuotaRequest) (*nas.SetDirQuotaResponse, error)
}

type CreateDirWillSucceedRequest struct {
	nas.CreateDirRequest
}

type CreateDirWillFailRequest struct {
	nas.CreateDirRequest
}
