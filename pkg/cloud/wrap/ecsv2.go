package wrap

import (
	"context"

	ecs "github.com/alibabacloud-go/ecs-20140526/v7/client"
	"github.com/alibabacloud-go/tea/dara"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
)

// loggingECSv2 decorates a cloud.ECSv2Interface so that every call is logged and
// its error transformed (see logV2). Because it implements the interface, the
// compiler forces a wrapper for each method.
type loggingECSv2 struct {
	inner cloud.ECSv2Interface
}

var _ cloud.ECSv2Interface = (*loggingECSv2)(nil)

// ECS wraps an ECS v2 client with call-boundary logging, metrics and error
// transformation. Wrap once at construction and pass the result wherever a
// cloud.ECSv2Interface is expected.
func ECS(inner cloud.ECSv2Interface) cloud.ECSv2Interface {
	return &loggingECSv2{
		inner: inner,
	}
}

func (c *loggingECSv2) DescribeInstancesWithContext(ctx context.Context, req *ecs.DescribeInstancesRequest, o *dara.RuntimeOptions) (*ecs.DescribeInstancesResponse, error) {
	return logV2(ctx, "ecs", "DescribeInstances", req, o, c.inner.DescribeInstancesWithContext)
}

func (c *loggingECSv2) DescribeInstanceTypesWithContext(ctx context.Context, req *ecs.DescribeInstanceTypesRequest, o *dara.RuntimeOptions) (*ecs.DescribeInstanceTypesResponse, error) {
	return logV2(ctx, "ecs", "DescribeInstanceTypes", req, o, c.inner.DescribeInstanceTypesWithContext)
}

func (c *loggingECSv2) DescribeAvailableResourceWithContext(ctx context.Context, req *ecs.DescribeAvailableResourceRequest, o *dara.RuntimeOptions) (*ecs.DescribeAvailableResourceResponse, error) {
	return logV2(ctx, "ecs", "DescribeAvailableResource", req, o, c.inner.DescribeAvailableResourceWithContext)
}

func (c *loggingECSv2) DescribeDisksWithContext(ctx context.Context, req *ecs.DescribeDisksRequest, o *dara.RuntimeOptions) (*ecs.DescribeDisksResponse, error) {
	return logV2(ctx, "ecs", "DescribeDisks", req, o, c.inner.DescribeDisksWithContext)
}

func (c *loggingECSv2) TagResourcesWithContext(ctx context.Context, req *ecs.TagResourcesRequest, o *dara.RuntimeOptions) (*ecs.TagResourcesResponse, error) {
	return logV2(ctx, "ecs", "TagResources", req, o, c.inner.TagResourcesWithContext)
}

func (c *loggingECSv2) UntagResourcesWithContext(ctx context.Context, req *ecs.UntagResourcesRequest, o *dara.RuntimeOptions) (*ecs.UntagResourcesResponse, error) {
	return logV2(ctx, "ecs", "UntagResources", req, o, c.inner.UntagResourcesWithContext)
}

func (c *loggingECSv2) ModifyDiskSpecWithContext(ctx context.Context, req *ecs.ModifyDiskSpecRequest, o *dara.RuntimeOptions) (*ecs.ModifyDiskSpecResponse, error) {
	return logV2(ctx, "ecs", "ModifyDiskSpec", req, o, c.inner.ModifyDiskSpecWithContext)
}

func (c *loggingECSv2) ModifyDiskAttributeWithContext(ctx context.Context, req *ecs.ModifyDiskAttributeRequest, o *dara.RuntimeOptions) (*ecs.ModifyDiskAttributeResponse, error) {
	return logV2(ctx, "ecs", "ModifyDiskAttribute", req, o, c.inner.ModifyDiskAttributeWithContext)
}

func (c *loggingECSv2) DescribeTasksWithContext(ctx context.Context, req *ecs.DescribeTasksRequest, o *dara.RuntimeOptions) (*ecs.DescribeTasksResponse, error) {
	return logV2(ctx, "ecs", "DescribeTasks", req, o, c.inner.DescribeTasksWithContext)
}
