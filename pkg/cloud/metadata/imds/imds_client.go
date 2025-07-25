package imds

import (
	"context"
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
	ECSTokenEndpoint    = ECSMetadataEndpoint + "api/token"
)

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

func (ctx *ecsFetchContext) fetch(client *http.Client) ([]byte, error) {
	if ctx.token == "" {
		err := ctx.fetchToken(client)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch token: %w", err)
		}
	}
	ctx.req.Header.Set("X-aliyun-ecs-metadata-token", ctx.token)
	resp, err := client.Do(ctx.req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			ctx.token = "" // the next retry should fetch a new token
		}
		return nil, HttpStatusError{StatusCode: resp.StatusCode}
	}
	return io.ReadAll(resp.Body)
}

type Client struct {
	http http.Client
}

func NewClient(httpRT http.RoundTripper) *Client {
	if httpRT == nil {
		httpRT = http.DefaultTransport
	}
	return &Client{
		http: http.Client{
			Transport: httpRT,
		},
	}
}

func (c *Client) Fetch(ctx context.Context, path string) ([]byte, error) {
	logger := klog.FromContext(ctx)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ECSMetadataEndpoint+path, nil)
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
	var data []byte
	err = wait.PollUntilContextCancel(ctx, 1*time.Second, true, func(ctx context.Context) (bool, error) {
		if lastErr != nil {
			logger.Error(lastErr, "retrying ECS metadata")
		}
		data, lastErr = fetchCtx.fetch(&c.http)
		if lastErr != nil {
			var httpErr HttpStatusError
			if errors.As(lastErr, &httpErr) && !httpErr.Retryable() {
				return false, lastErr // no retry
			}
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch from ECS IMDS: %w", lastErr)
	}
	return data, nil

}
