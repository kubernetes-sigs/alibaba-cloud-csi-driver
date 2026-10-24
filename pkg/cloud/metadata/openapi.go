package metadata

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	"github.com/alibabacloud-go/tea/dara"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"k8s.io/utils/ptr"
)

type OpenAPIMetadata struct {
	instance *ecs20140526.DescribeInstancesResponseBodyInstancesInstance
}

func NewOpenAPIMetadata(ctx context.Context, c cloud.ECSv2Interface, instanceId string) (*OpenAPIMetadata, error) {
	ids, err := json.Marshal([]string{instanceId})
	if err != nil {
		panic(err)
	}
	instanceRequest := ecs20140526.DescribeInstancesRequest{
		InstanceIds: new(string(ids)),
	}
	if diskHighDensityModeEnabled() {
		// Request the high-density flag from AdditionalInfo. Only when the feature is
		// on, so a server that rejects the non-standard attribute never receives it.
		instanceRequest.AdditionalAttributes = []*string{new("DISK_HIGH_DENSITY_MODE")}
	}

	instanceResponse, err := c.DescribeInstancesWithContext(ctx, &instanceRequest, &dara.RuntimeOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to describe instance %s: %w", instanceId, err)
	}
	if instanceResponse.Body == nil || instanceResponse.Body.Instances == nil {
		return nil, fmt.Errorf("no instances field: %s", instanceId)
	}
	for _, i := range instanceResponse.Body.Instances.Instance {
		if i != nil && i.InstanceId != nil && *i.InstanceId == instanceId {
			return &OpenAPIMetadata{instance: i}, nil
		}
	}
	return nil, fmt.Errorf("instance %s not found in response", instanceId)
}

func (m *OpenAPIMetadata) get(key MetadataKey) *string {
	switch key {
	case ZoneID:
		return m.instance.ZoneId
	case InstanceType:
		return m.instance.InstanceType
	}
	return nil
}

func (m *OpenAPIMetadata) GetAny(_ *mcontext, key MetadataKey) (any, error) {
	if key == isHighDensityMode {
		ai := m.instance.AdditionalInfo
		return ai != nil && ptr.Deref(ai.EnableHighDensityMode, false), nil
	}
	if v := m.get(key); v != nil {
		return *v, nil
	}
	return nil, ErrUnknownMetadataKey
}

type OpenAPIFetcher struct {
	ecsClient cloud.ECSv2Interface
	mPre      middleware
}

func (f *OpenAPIFetcher) ID() fetcherID { return openAPIFetcherID }

func (f *OpenAPIFetcher) FetchFor(ctx *mcontext, key MetadataKey) (middleware, error) {
	switch key {
	case ZoneID, InstanceType, isHighDensityMode:
	default:
		return nil, ErrUnknownMetadataKey
	}

	kind, err := f.mPre.GetAny(ctx, machineKind)
	if err == nil && kind != MachineKindECS { // skip for non-ECS instances
		ctx.logger.V(1).Info("skip ECS DescribeInstances metadata fetcher", "machineKind", kind)
		return empty{}, nil
	}
	instanceId, err := f.mPre.GetAny(ctx, InstanceID)
	if err != nil {
		if err == ErrUnknownMetadataKey {
			nodeName := os.Getenv(KUBE_NODE_NAME_ENV)
			if strings.HasPrefix(nodeName, "i-") {
				// on some setup, node name is the instance id, try that if it looks like
				instanceId = nodeName
				err = nil
			}
		}
	}
	if err != nil {
		return nil, fmt.Errorf("instance ID is not available: %w", err)
	}
	p, err := NewOpenAPIMetadata(ctx, f.ecsClient, instanceId.(string))
	if err != nil {
		return nil, err
	}
	return newImmutable(p, "OpenAPI"), nil
}
