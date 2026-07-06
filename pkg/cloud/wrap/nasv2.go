package wrap

import (
	"context"

	nas "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/dara"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
)

// loggingNAS decorates a cloud.NasInterface so that every call is logged,
// metered and its error transformed (see logV2). Because it implements the
// interface, the compiler forces a wrapper for each method.
type loggingNAS struct {
	inner cloud.NasInterface
}

var _ cloud.NasInterface = (*loggingNAS)(nil)

// NAS wraps a NAS v2 client with call-boundary logging, metrics and error
// transformation. Wrap once at construction and pass the result wherever a
// cloud.NasInterface is expected.
func NAS(inner cloud.NasInterface) cloud.NasInterface {
	return &loggingNAS{inner: inner}
}

func (c *loggingNAS) CancelDirQuotaWithContext(ctx context.Context, req *nas.CancelDirQuotaRequest, o *dara.RuntimeOptions) (*nas.CancelDirQuotaResponse, error) {
	return logV2(ctx, "nas", "CancelDirQuota", req, o, c.inner.CancelDirQuotaWithContext)
}

func (c *loggingNAS) CreateAccessPointWithContext(ctx context.Context, req *nas.CreateAccessPointRequest, o *dara.RuntimeOptions) (*nas.CreateAccessPointResponse, error) {
	return logV2(ctx, "nas", "CreateAccessPoint", req, o, c.inner.CreateAccessPointWithContext)
}

func (c *loggingNAS) CreateDirWithContext(ctx context.Context, req *nas.CreateDirRequest, o *dara.RuntimeOptions) (*nas.CreateDirResponse, error) {
	return logV2(ctx, "nas", "CreateDir", req, o, c.inner.CreateDirWithContext)
}

func (c *loggingNAS) DeleteAccessPointWithContext(ctx context.Context, req *nas.DeleteAccessPointRequest, o *dara.RuntimeOptions) (*nas.DeleteAccessPointResponse, error) {
	return logV2(ctx, "nas", "DeleteAccessPoint", req, o, c.inner.DeleteAccessPointWithContext)
}

func (c *loggingNAS) DescribeAccessPointWithContext(ctx context.Context, req *nas.DescribeAccessPointRequest, o *dara.RuntimeOptions) (*nas.DescribeAccessPointResponse, error) {
	return logV2(ctx, "nas", "DescribeAccessPoint", req, o, c.inner.DescribeAccessPointWithContext)
}

func (c *loggingNAS) DescribeFileSystemsWithContext(ctx context.Context, req *nas.DescribeFileSystemsRequest, o *dara.RuntimeOptions) (*nas.DescribeFileSystemsResponse, error) {
	return logV2(ctx, "nas", "DescribeFileSystems", req, o, c.inner.DescribeFileSystemsWithContext)
}

func (c *loggingNAS) GetRecycleBinAttributeWithContext(ctx context.Context, req *nas.GetRecycleBinAttributeRequest, o *dara.RuntimeOptions) (*nas.GetRecycleBinAttributeResponse, error) {
	return logV2(ctx, "nas", "GetRecycleBinAttribute", req, o, c.inner.GetRecycleBinAttributeWithContext)
}

func (c *loggingNAS) SetDirQuotaWithContext(ctx context.Context, req *nas.SetDirQuotaRequest, o *dara.RuntimeOptions) (*nas.SetDirQuotaResponse, error) {
	return logV2(ctx, "nas", "SetDirQuota", req, o, c.inner.SetDirQuotaWithContext)
}
