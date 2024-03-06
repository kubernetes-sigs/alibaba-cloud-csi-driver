package metadata

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
)

type OpenAPIMetadata struct {
	instance *ecs.Instance
	identity *sts.GetCallerIdentityResponse
}

func NewOpenAPIMetadata(c cloud.ECSInterface, s cloud.STSInterface, regionId, instanceId string) (*OpenAPIMetadata, error) {
	instanceRequest := ecs.CreateDescribeInstancesRequest()

	instanceRequest.RegionId = regionId
	ids, err := json.Marshal([]string{instanceId})
	if err != nil {
		panic(err)
	}
	instanceRequest.InstanceIds = string(ids)

	instanceResponse, err := c.DescribeInstances(instanceRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to describe instance %s: %w", instanceId, err)
	}
	if len(instanceResponse.Instances.Instance) != 1 {
		return nil, fmt.Errorf("instance not found: %s", instanceId)
	}

	identityRequest := sts.CreateGetCallerIdentityRequest()
	identityRequest.Scheme = "https"
	identityResponse, err := s.GetCallerIdentity(identityRequest)
	if err == nil {
		return &OpenAPIMetadata{instance: &instanceResponse.Instances.Instance[0], identity: identityResponse}, nil
	}

	return &OpenAPIMetadata{instance: &instanceResponse.Instances.Instance[0], identity: nil}, nil
}

func (m *OpenAPIMetadata) Get(key MetadataKey) (string, error) {
	switch key {
	case RegionID:
		return m.instance.RegionId, nil
	case ZoneID:
		return m.instance.ZoneId, nil
	case InstanceID:
		return m.instance.InstanceId, nil
	case InstanceType:
		return m.instance.InstanceType, nil
	case AccountID:
		if m.identity != nil {
			return m.identity.AccountId, nil
		}
	}
	return "", ErrUnknownMetadataKey
}

type OpenAPIFetcher struct {
	ecsClient cloud.ECSInterface
	stsClient cloud.STSInterface
	mPre      MetadataProvider
}

func (f *OpenAPIFetcher) FetchFor(key MetadataKey) (MetadataProvider, error) {
	switch key {
	case RegionID, InstanceID, ZoneID, InstanceType, AccountID:
	default:
		return nil, ErrUnknownMetadataKey
	}

	regionId, err := f.mPre.Get(RegionID)
	if err != nil {
		return nil, fmt.Errorf("region ID is not available: %w", err)
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
	p, err := NewOpenAPIMetadata(f.ecsClient, f.stsClient, regionId, instanceId)
	if err != nil {
		return nil, err
	}
	return newImmutableProvider(p, "OpenAPI"), nil
}
