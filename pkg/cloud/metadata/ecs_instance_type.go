package metadata

import (
	"fmt"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"k8s.io/utils/ptr"
)

type ECSInstanceTypeMetadata struct {
	t *ecs20140526.DescribeInstanceTypesResponseBodyInstanceTypesInstanceType
}

func NewEcsInstanceTypeMetadata(c cloud.ECSv2Interface, instanceType string) (*ECSInstanceTypeMetadata, error) {
	req := ecs20140526.DescribeInstanceTypesRequest{
		InstanceTypes: []*string{ptr.To(instanceType)},
	}
	resp, err := c.DescribeInstanceTypes(&req)
	if err != nil {
		return nil, fmt.Errorf("failed to describe instance type %s: %w", instanceType, err)
	}
	if resp.Body == nil || resp.Body.InstanceTypes == nil {
		return nil, fmt.Errorf("no instance types field: %s", instanceType)
	}
	for _, t := range resp.Body.InstanceTypes.InstanceType {
		if t != nil && t.InstanceTypeId != nil && *t.InstanceTypeId == instanceType {
			return &ECSInstanceTypeMetadata{t: t}, nil
		}
	}
	return nil, fmt.Errorf("instance type %s not found in response", instanceType)
}

func (m *ECSInstanceTypeMetadata) GetAny(_ *mcontext, key MetadataKey) (any, error) {
	switch key {
	case diskQuantity:
		if m.t.DiskQuantity != nil {
			return *m.t.DiskQuantity, nil
		}
	}
	return nil, ErrUnknownMetadataKey
}

type ECSInstanceTypeFetcher struct {
	ecsClient cloud.ECSv2Interface
	mPre      middleware
}

func (f *ECSInstanceTypeFetcher) ID() fetcherID { return ecsInstanceTypeFetcherID }

func (f *ECSInstanceTypeFetcher) FetchFor(ctx *mcontext, key MetadataKey) (middleware, error) {
	switch key {
	case diskQuantity:
	default:
		return nil, ErrUnknownMetadataKey
	}

	kind, err := f.mPre.GetAny(ctx, machineKind)
	if err == nil && kind != MachineKindECS { // skip for non-ECS instances
		ctx.logger.V(1).Info("skip ECS DescribeInstanceTypes metadata fetcher", "machineKind", kind)
		return empty{}, nil
	}
	t, err := f.mPre.GetAny(ctx, InstanceType)
	if err != nil {
		return nil, fmt.Errorf("instance type is not available: %w", err)
	}
	p, err := NewEcsInstanceTypeMetadata(f.ecsClient, t.(string))
	if err != nil {
		return nil, err
	}
	return newImmutable(p, "ECS_Instance_Type"), nil
}
