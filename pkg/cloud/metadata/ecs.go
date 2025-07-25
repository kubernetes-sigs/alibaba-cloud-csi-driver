package metadata

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata/imds"
	"k8s.io/klog/v2"
)

const (
	ECSIdentityPath = "dynamic/instance-identity/document"
)

type InstanceIdentityDocument struct {
	ZoneID         string `json:"zone-id"`
	RegionID       string `json:"region-id"`
	InstanceID     string `json:"instance-id"`
	InstanceType   string `json:"instance-type"`
	SerialNumber   string `json:"serial-number"`
	OwnerAccountID string `json:"owner-account-id"`
}

type ECSMetadata struct {
	idDoc InstanceIdentityDocument
}

var ErrInvalidIdentityDoc = errors.New("invalid ECS instance identity document")

func NewECSMetadata(httpRT http.RoundTripper) (*ECSMetadata, error) {
	m := &ECSMetadata{}

	imdsClient := imds.NewClient(httpRT)

	data, err := imdsClient.Fetch(context.Background(), ECSIdentityPath)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &m.idDoc); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidIdentityDoc, err)
	}
	if m.idDoc.RegionID == "" {
		return nil, fmt.Errorf("%w: missing region-id", ErrInvalidIdentityDoc)
	}
	return m, nil
}

func (m *ECSMetadata) Get(key MetadataKey) (string, error) {
	switch key {
	case RegionID:
		return m.idDoc.RegionID, nil
	case ZoneID:
		return m.idDoc.ZoneID, nil
	case InstanceID:
		return m.idDoc.InstanceID, nil
	case InstanceType:
		return m.idDoc.InstanceType, nil
	case AccountID:
		return m.idDoc.OwnerAccountID, nil
	default:
		return "", ErrUnknownMetadataKey
	}
}

type EcsFetcher struct {
	httpRT http.RoundTripper
}

func (f *EcsFetcher) FetchFor(key MetadataKey) (MetadataProvider, error) {
	switch key {
	case RegionID, ZoneID, InstanceID, InstanceType, AccountID:
	default:
		return nil, ErrUnknownMetadataKey
	}

	ecs, err := NewECSMetadata(f.httpRT)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			klog.Warningf("Hint: ECS metadata is only available when running on Alibaba Cloud ECS. "+
				"Set %s environment variable to disable ECS metadata for faster initialization.", DISABLE_ECS_ENV)
		}
		return nil, err
	}
	return newImmutableProvider(ecs, "ECS"), nil
}
