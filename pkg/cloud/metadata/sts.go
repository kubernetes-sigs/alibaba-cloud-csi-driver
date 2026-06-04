package metadata

import (
	"context"
	"fmt"

	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"k8s.io/klog/v2"
)

type StsMetadata struct {
	identity *sts20150401.GetCallerIdentityResponseBody
}

func NewStsMetadata(ctx context.Context, s cloud.STSInterface) (*StsMetadata, error) {
	resp, err := s.GetCallerIdentity()
	if err != nil {
		return nil, fmt.Errorf("failed to get caller identity: %w", err)
	}
	logger := klog.FromContext(ctx)
	logger.V(2).Info("GetCallerIdentity OK", "requestID", tea.StringValue(resp.Body.RequestId))
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
	stsClient cloud.STSInterface
}

func (f *StsFetcher) ID() fetcherID { return stsFetcherID }

func (f *StsFetcher) FetchFor(ctx *mcontext, key MetadataKey) (middleware, error) {
	switch key {
	case AccountID:
	default:
		return nil, ErrUnknownMetadataKey
	}

	p, err := NewStsMetadata(ctx, f.stsClient)
	if err != nil {
		return nil, err
	}
	return newImmutable(strProvider{p}, "Sts"), nil
}
