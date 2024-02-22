package metadata

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/wait"
)

const (
	ECSMetadataEndpoint = "http://100.100.100.200/latest/"
	ECSIdentityEndpoint = ECSMetadataEndpoint + "dynamic/instance-identity/document"
)

type InstanceIdentityDocument struct {
	ZoneID       string `json:"zone-id"`
	RegionID     string `json:"region-id"`
	InstanceID   string `json:"instance-id"`
	InstanceType string `json:"instance-type"`
	SerialNumber string `json:"serial-number"`
	AliUID       string `json:"owner-account-id"`
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
	err = wait.PollImmediateUntil(1*time.Second, func() (bool, error) {
		if lastErr != nil {
			logrus.Warnf("retrying ECS metadata: %v", lastErr)
		}
		resp, err := httpClient.Do(req)
		if err != nil {
			lastErr = err
			return false, nil
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			lastErr = fmt.Errorf("unexpected status code %d", resp.StatusCode)
			if resp.StatusCode >= 500 {
				return false, nil // retry on server errors
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
	}, ctx.Done())
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
	case AliUID:
		return m.idDoc.AliUID, nil
	default:
		return "", ErrUnknownMetadataKey
	}
}

type EcsFetcher struct {
	httpRT http.RoundTripper
}

func (f *EcsFetcher) FetchFor(key MetadataKey) (MetadataProvider, error) {
	switch key {
	case RegionID, ZoneID, InstanceID, InstanceType, AliUID:
	default:
		return nil, ErrUnknownMetadataKey
	}

	ecs, err := NewECSMetadata(f.httpRT)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			logrus.Warnf("Hint: ECS metadata is only available when running on Alibaba Cloud ECS. "+
				"Set %s environment variable to disable ECS metadata for faster initialization.", DISABLE_ECS_ENV)
		}
		return nil, err
	}
	return newImmutableProvider(ecs, "ECS"), nil
}
