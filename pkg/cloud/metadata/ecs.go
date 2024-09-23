package metadata

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
)

const (
	ECSMetadataEndpoint = "http://100.100.100.200/latest/"
	ECSIdentityEndpoint = ECSMetadataEndpoint + "dynamic/instance-identity/document"
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

func NewECSMetadata(httpRT http.RoundTripper) (*ECSMetadata, error) {
	m := &ECSMetadata{}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	httpClient := http.Client{
		Transport: httpRT,
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ECSIdentityEndpoint, nil)
	if err != nil {
		return nil, err
	}

	var lastErr error
	err = wait.PollImmediateUntilWithContext(ctx, 1*time.Second, func(ctx context.Context) (bool, error) {
		if lastErr != nil {
			klog.Warningf("retrying ECS metadata: %v", lastErr)
		}
		resp, err := httpClient.Do(req)
		if err != nil {
			lastErr = err
			return false, nil
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			lastErr = fmt.Errorf("unexpected status code %d", resp.StatusCode)
			if resp.StatusCode >= 500 || resp.StatusCode == 429 {
				return false, nil // retry on server errors/too many requests
			} else {
				return false, lastErr
			}
		}
		if err := json.NewDecoder(resp.Body).Decode(&m.idDoc); err != nil {
			return false, err
		}
		if m.idDoc.RegionID == "" {
			return false, fmt.Errorf("missing region-id in ECS instance identity document, possibly not running on ECS")
		}
		return true, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ECS instance identity document: %w", lastErr)
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
