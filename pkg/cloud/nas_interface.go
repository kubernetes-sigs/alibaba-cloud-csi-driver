package cloud

import (
	"context"

	nas "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/dara"
)

type NasInterface interface {
	CancelDirQuotaWithContext(ctx context.Context, request *nas.CancelDirQuotaRequest, o *dara.RuntimeOptions) (*nas.CancelDirQuotaResponse, error)
	CreateAccessPointWithContext(ctx context.Context, request *nas.CreateAccessPointRequest, o *dara.RuntimeOptions) (*nas.CreateAccessPointResponse, error)
	CreateAgenticSpaceWithContext(ctx context.Context, request *nas.CreateAgenticSpaceRequest, o *dara.RuntimeOptions) (*nas.CreateAgenticSpaceResponse, error)
	CreateDirWithContext(ctx context.Context, request *nas.CreateDirRequest, o *dara.RuntimeOptions) (*nas.CreateDirResponse, error)
	DeleteAccessPointWithContext(ctx context.Context, request *nas.DeleteAccessPointRequest, o *dara.RuntimeOptions) (*nas.DeleteAccessPointResponse, error)
	DeleteAgenticSpaceWithContext(ctx context.Context, request *nas.DeleteAgenticSpaceRequest, o *dara.RuntimeOptions) (*nas.DeleteAgenticSpaceResponse, error)
	DescribeAccessPointWithContext(ctx context.Context, request *nas.DescribeAccessPointRequest, o *dara.RuntimeOptions) (*nas.DescribeAccessPointResponse, error)
	DescribeAgenticSpacesWithContext(ctx context.Context, request *nas.DescribeAgenticSpacesRequest, o *dara.RuntimeOptions) (*nas.DescribeAgenticSpacesResponse, error)
	DescribeFileSystemsWithContext(ctx context.Context, request *nas.DescribeFileSystemsRequest, o *dara.RuntimeOptions) (*nas.DescribeFileSystemsResponse, error)
	GetAgenticSpaceWithContext(ctx context.Context, request *nas.GetAgenticSpaceRequest, o *dara.RuntimeOptions) (*nas.GetAgenticSpaceResponse, error)
	GetRecycleBinAttributeWithContext(ctx context.Context, request *nas.GetRecycleBinAttributeRequest, o *dara.RuntimeOptions) (*nas.GetRecycleBinAttributeResponse, error)
	ListAccessPointsWithContext(ctx context.Context, request *nas.ListAccessPointsRequest, o *dara.RuntimeOptions) (*nas.ListAccessPointsResponse, error)
	SetAgenticSpaceQuotaWithContext(ctx context.Context, request *nas.SetAgenticSpaceQuotaRequest, o *dara.RuntimeOptions) (*nas.SetAgenticSpaceQuotaResponse, error)
	SetDirQuotaWithContext(ctx context.Context, request *nas.SetDirQuotaRequest, o *dara.RuntimeOptions) (*nas.SetDirQuotaResponse, error)
}
