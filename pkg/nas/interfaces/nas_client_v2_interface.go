package interfaces

import sdk "github.com/alibabacloud-go/nas-20170626/v3/client"

type NasClientV2Interface interface {
	CreateDir(req *sdk.CreateDirRequest) error
	SetDirQuota(req *sdk.SetDirQuotaRequest) error
	CancelDirQuota(req *sdk.CancelDirQuotaRequest) error
	GetRecycleBinAttribute(filesystemId string) (*sdk.GetRecycleBinAttributeResponse, error)
	CreateAccesspoint(req *sdk.CreateAccessPointRequest) (*sdk.CreateAccessPointResponse, error)
	DeleteAccesspoint(filesystemId, accessPointId string) error
	DescribeAccesspoint(filesystemId, accessPointId string) (*sdk.DescribeAccessPointResponse, error)
	DescribeFileSystems(filesystemID string) (*sdk.DescribeFileSystemsResponse, error)
}
