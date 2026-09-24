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
	// ListAccesspoints enumerates accesspoints, supporting the server-side Filters
	// (AccessPointId / AgenticSpaceId / AccessGroup). Unlike the DescribeAccessPoints
	// OpenAPI, its response items carry AgenticSpaceId, so this is the only call that
	// can prove which AgenticSpace an accesspoint is bound to.
	ListAccesspoints(ctx context.Context, req *sdk.ListAccessPointsRequest) (*sdk.ListAccessPointsResponse, error)
	DescribeFileSystems(ctx context.Context, filesystemID string) (*sdk.DescribeFileSystemsResponse, error)

	// AgenticSpace CRUD
	CreateAgenticSpace(ctx context.Context, req *sdk.CreateAgenticSpaceRequest) (*sdk.CreateAgenticSpaceResponse, error)
	DescribeAgenticSpaces(ctx context.Context, req *sdk.DescribeAgenticSpacesRequest) (*sdk.DescribeAgenticSpacesResponse, error)
	GetAgenticSpace(ctx context.Context, req *sdk.GetAgenticSpaceRequest) (*sdk.GetAgenticSpaceResponse, error)
	DeleteAgenticSpace(ctx context.Context, req *sdk.DeleteAgenticSpaceRequest) (*sdk.DeleteAgenticSpaceResponse, error)
	SetAgenticSpaceQuota(ctx context.Context, req *sdk.SetAgenticSpaceQuotaRequest) (*sdk.SetAgenticSpaceQuotaResponse, error)
}
