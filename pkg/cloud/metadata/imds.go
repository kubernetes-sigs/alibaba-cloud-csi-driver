package metadata

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata/imds"
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

type IMDSMetadata struct {
	idDoc InstanceIdentityDocument
}

var ErrInvalidIdentityDoc = errors.New("invalid ECS instance identity document")

func NewECSMetadata(ctx context.Context, httpRT http.RoundTripper) (*IMDSMetadata, error) {
	m := &IMDSMetadata{}

	imdsClient := imds.NewClient(httpRT)

	data, err := imdsClient.Fetch(ctx, ECSIdentityPath)
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

func (m *IMDSMetadata) Get(key MetadataKey) (string, error) {
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

type IMDSFetcer struct {
	httpRT http.RoundTripper
}

func (f *IMDSFetcer) ID() fetcherID { return imdsFetcherID }

func (f *IMDSFetcer) FetchFor(ctx *mcontext, key MetadataKey) (middleware, error) {
	switch key {
	case RegionID, ZoneID, InstanceID, InstanceType, AccountID:
	default:
		return nil, ErrUnknownMetadataKey
	}

	ecs, err := NewECSMetadata(ctx, f.httpRT)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			ctx.logger.Info("Hint: ECS metadata is only available when running on Alibaba Cloud ECS. " +
				"Set " + DISABLE_IMDS_ENV + " environment variable to disable ECS metadata for faster initialization.")
		}
		return nil, err
	}
	return newImmutable(strProvider{ecs}, "IMDS"), nil
}
