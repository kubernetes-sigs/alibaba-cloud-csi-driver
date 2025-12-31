package metadata

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"k8s.io/utils/ptr"
)

type OpenAPIMetadata struct {
	instance *ecs20140526.DescribeInstancesResponseBodyInstancesInstance
}

func NewOpenAPIMetadata(c cloud.ECSv2Interface, instanceId string) (*OpenAPIMetadata, error) {
	ids, err := json.Marshal([]string{instanceId})
	if err != nil {
		panic(err)
	}
	instanceRequest := ecs20140526.DescribeInstancesRequest{
		InstanceIds: ptr.To(string(ids)),
	}

	instanceResponse, err := c.DescribeInstances(&instanceRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to describe instance %s: %w", instanceId, err)
	}
	if instanceResponse.Body == nil || instanceResponse.Body.Instances == nil {
		return nil, fmt.Errorf("no instances field: %s", instanceId)
	}
	instances := instanceResponse.Body.Instances.Instance
	if len(instances) != 1 || instances[0] == nil {
		return nil, fmt.Errorf("unexpected instance count %d for %s", len(instances), instanceId)
	}
	return &OpenAPIMetadata{instance: instances[0]}, nil
}

func (m *OpenAPIMetadata) get(key MetadataKey) *string {
	switch key {
	case ZoneID:
		return m.instance.ZoneId
	case InstanceID:
		return m.instance.InstanceId
	case InstanceType:
		return m.instance.InstanceType
	}
	return nil
}

func (m *OpenAPIMetadata) Get(key MetadataKey) (string, error) {
	v := m.get(key)
	if v == nil {
		return "", ErrUnknownMetadataKey
	}
	return *v, nil
}

type OpenAPIFetcher struct {
	ecsClient cloud.ECSv2Interface
	mPre      MetadataProvider
}

func (f *OpenAPIFetcher) FetchFor(key MetadataKey) (middleware, error) {
	switch key {
	case InstanceID, ZoneID, InstanceType:
	default:
		return nil, ErrUnknownMetadataKey
	}

	instanceId, err := f.mPre.Get(InstanceID)
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
	p, err := NewOpenAPIMetadata(f.ecsClient, instanceId)
	if err != nil {
		return nil, err
	}
	return newImmutable(strProvider{p}, "OpenAPI"), nil
}
