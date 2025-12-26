package interfaces

import (
	"context"

	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
)

type NasClientV2Interface interface {
	CreateDir(ctx context.Context, req *sdk.CreateDirRequest) error
	SetDirQuota(ctx context.Context, req *sdk.SetDirQuotaRequest) error
	CancelDirQuota(ctx context.Context, req *sdk.CancelDirQuotaRequest) error
	GetRecycleBinAttribute(ctx context.Context, filesystemId string) (*sdk.GetRecycleBinAttributeResponse, error)
	CreateAccesspoint(ctx context.Context, req *sdk.CreateAccessPointRequest) (*sdk.CreateAccessPointResponse, error)
	DeleteAccesspoint(ctx context.Context, filesystemId, accessPointId string) error
	DescribeAccesspoint(ctx context.Context, filesystemId, accessPointId string) (*sdk.DescribeAccessPointResponse, error)
	DescribeFileSystems(ctx context.Context, filesystemID string) (*sdk.DescribeFileSystemsResponse, error)
}
