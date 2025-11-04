package metadata

import (
	"fmt"

	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"k8s.io/klog/v2"
)

type StsMetadata struct {
	identity *sts20150401.GetCallerIdentityResponseBody
}

func NewStsMetadata(s cloud.STSInterface) (*StsMetadata, error) {
	resp, err := s.GetCallerIdentity()
	if err != nil {
		return nil, fmt.Errorf("failed to get caller identity: %w", err)
	}
	klog.V(2).InfoS("GetCallerIdentity OK", "requestID", tea.StringValue(resp.Body.RequestId))
	return &StsMetadata{identity: resp.Body}, nil
}

func (m *StsMetadata) Get(key MetadataKey) (string, error) {
	switch key {
	case AccountID:
		if m.identity.AccountId != nil {
			return *m.identity.AccountId, nil
		}
	}
	return "", ErrUnknownMetadataKey
}

type StsFetcher struct {
	stsClient func(regionID string) (cloud.STSInterface, error)
	mPre      MetadataProvider
}

func (f *StsFetcher) FetchFor(key MetadataKey) (MetadataProvider, error) {
	switch key {
	case AccountID:
	default:
		return nil, ErrUnknownMetadataKey
	}

	regionId, err := f.mPre.Get(RegionID)
	if err != nil {
		return nil, fmt.Errorf("region ID is not available: %w", err)
	}
	client, err := f.stsClient(regionId)
	if err != nil {
		return nil, fmt.Errorf("failed to create STS client: %w", err)
	}
	p, err := NewStsMetadata(client)
	if err != nil {
		return nil, err
	}
	return newImmutableProvider(p, "Sts"), nil
}
