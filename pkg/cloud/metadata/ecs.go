package metadata

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
)

const (
	ECSMetadataEndpoint = "http://100.100.100.200/latest/"
	ECSIdentityEndpoint = ECSMetadataEndpoint + "dynamic/instance-identity/document"
	ECSTokenEndpoint    = ECSMetadataEndpoint + "api/token"
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

type HttpStatusError struct {
	StatusCode int
}

func (e HttpStatusError) Error() string {
	return fmt.Sprintf("unexpected status error: %d", e.StatusCode)
}

func (e HttpStatusError) Retryable() bool {
	return e.StatusCode >= 500 || e.StatusCode == http.StatusTooManyRequests ||
		e.StatusCode == http.StatusUnauthorized // re-fetch token on 401
}

var ErrInvalidIdentityDoc = errors.New("invalid ECS instance identity document")

type ecsFetchContext struct {
	token    string
	tokenReq *http.Request
	req      *http.Request
}

func (ctx *ecsFetchContext) fetchToken(client *http.Client) error {
	resp, err := client.Do(ctx.tokenReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return HttpStatusError{resp.StatusCode}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	ctx.token = string(body)
	return nil
}

func (ctx *ecsFetchContext) fetch(client *http.Client) (InstanceIdentityDocument, error) {
	var doc InstanceIdentityDocument
	if ctx.token == "" {
		err := ctx.fetchToken(client)
		if err != nil {
			return doc, fmt.Errorf("failed to fetch token: %w", err)
		}
	}
	ctx.req.Header.Set("X-aliyun-ecs-metadata-token", ctx.token)
	resp, err := client.Do(ctx.req)
	if err != nil {
		return doc, err
	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			ctx.token = "" // the next retry should fetch a new token
		}
		return doc, HttpStatusError{StatusCode: resp.StatusCode}
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&doc); err != nil {
		var jerr *json.SyntaxError
		if errors.As(err, &jerr) {
			return doc, fmt.Errorf("%w: %w", ErrInvalidIdentityDoc, err)
		}
		return doc, fmt.Errorf("failed to read body: %w", err)
	}
	if doc.RegionID == "" {
		return doc, fmt.Errorf("%w: missing region-id", ErrInvalidIdentityDoc)
	}
	return doc, nil
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
	tokenReq, err := http.NewRequestWithContext(ctx, http.MethodPut, ECSTokenEndpoint, nil)
	if err != nil {
		return nil, err
	}
	tokenReq.Header.Set("X-aliyun-ecs-metadata-token-ttl-seconds", "63")
	fetchCtx := &ecsFetchContext{
		tokenReq: tokenReq,
		req:      req,
	}

	var lastErr error
	err = wait.PollUntilContextCancel(ctx, 1*time.Second, true, func(ctx context.Context) (bool, error) {
		if lastErr != nil {
			klog.Warningf("retrying ECS metadata: %v", lastErr)
		}
		m.idDoc, lastErr = fetchCtx.fetch(&httpClient)
		if lastErr == nil {
			return true, nil
		}

		var httpErr HttpStatusError
		if errors.As(lastErr, &httpErr) && !httpErr.Retryable() {
			return false, lastErr // no retry
		}
		if errors.Is(lastErr, ErrInvalidIdentityDoc) {
			return false, lastErr // no retry, maybe not on ECS
		}
		return false, nil
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
