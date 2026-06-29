package cloud

import (
	"context"

	nas "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/dara"
)

type NasInterface interface {
	CancelDirQuotaWithContext(ctx context.Context, request *nas.CancelDirQuotaRequest, o *dara.RuntimeOptions) (*nas.CancelDirQuotaResponse, error)
	CreateAccessPointWithContext(ctx context.Context, request *nas.CreateAccessPointRequest, o *dara.RuntimeOptions) (*nas.CreateAccessPointResponse, error)
	CreateDirWithContext(ctx context.Context, request *nas.CreateDirRequest, o *dara.RuntimeOptions) (*nas.CreateDirResponse, error)
	DeleteAccessPointWithContext(ctx context.Context, request *nas.DeleteAccessPointRequest, o *dara.RuntimeOptions) (*nas.DeleteAccessPointResponse, error)
	DescribeAccessPointWithContext(ctx context.Context, request *nas.DescribeAccessPointRequest, o *dara.RuntimeOptions) (*nas.DescribeAccessPointResponse, error)
	DescribeFileSystemsWithContext(ctx context.Context, request *nas.DescribeFileSystemsRequest, o *dara.RuntimeOptions) (*nas.DescribeFileSystemsResponse, error)
	GetRecycleBinAttributeWithContext(ctx context.Context, request *nas.GetRecycleBinAttributeRequest, o *dara.RuntimeOptions) (*nas.GetRecycleBinAttributeResponse, error)
	SetDirQuotaWithContext(ctx context.Context, request *nas.SetDirQuotaRequest, o *dara.RuntimeOptions) (*nas.SetDirQuotaResponse, error)
}
