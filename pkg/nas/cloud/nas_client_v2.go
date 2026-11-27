package cloud

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/dara"
	"github.com/alibabacloud-go/tea/tea"
	alicred_old "github.com/aliyun/credentials-go/credentials"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/credentials"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/interfaces"
	utilshttp "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/http"
	"golang.org/x/time/rate"
	"k8s.io/klog/v2"
)

const (
	// connTimeout is the SDK ConnectTimeout, in milliseconds per the darabonba convention. The
	// runtime uses it as the deadline of the whole HTTP call rather than only the dial, so this
	// bounds the HTTP exchange at 10s, not credential resolution before it.
	// ReadTimeout is deliberately left unset: it is added to that deadline,
	// and leaving it at 0 keeps the 10s HTTP bound authoritative.
	connTimeout = 10000
)

func NewNasClientV2(region string) (*sdk.Client, error) {
	headers := utilshttp.MustParseHeaderEnv("NAS_HEADERS")
	var headersV2 map[string]*string
	if headers != nil {
		headersV2 = utilshttp.MustToV2SDKHeaders(headers)
	}
	config := new(openapi.Config).
		SetUserAgent(KubernetesAlicloudIdentity).
		SetRegionId(region).
		SetConnectTimeout(connTimeout).
		SetGlobalParameters(&openapi.GlobalParameters{
			Queries: map[string]*string{
				"RegionId": &region,
			},
			Headers: headersV2,
		})
	// set credential
	provider, err := credentials.NewProvider()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch credential: %w", err)
	}
	config = config.SetCredential(alicred_old.FromCredentialsProvider(provider.GetProviderName(), provider))
	// set endpoint
	ep := os.Getenv("NAS_ENDPOINT")
	if ep == "" {
		ep = GetEndpointForRegion(region)
	}
	config = config.SetEndpoint(ep)
	// set protocol
	scheme := "HTTPS"
	if e := os.Getenv("ALICLOUD_CLIENT_SCHEME"); e != "" {
		scheme = e
	}
	config = config.SetProtocol(scheme)
	// init client
	return sdk.NewClient(config)
}

type NasClientV2 struct {
	region string
	// limiter is the fallback bucket for an action without a dedicated entry.
	limiter *rate.Limiter
	// actionLimiters holds one adaptive bucket per OpenAPI action. The NAS
	// quota is enforced per (region, API, account), so sharing a single bucket
	// across different actions serializes quotas the service keeps separate.
	actionLimiters map[string]*actionLimiter
	client         cloud.NasInterface
}

// Reported here rather than at the factory's return statement when the two drift apart.
var _ interfaces.NasClientV2Interface = (*NasClientV2)(nil)

var (
	// longThrottleLatency defines threshold for logging requests. All requests being
	// throttled (via the provided rateLimiter) for more than longThrottleLatency will
	// be logged.
	longThrottleLatency = 250 * time.Millisecond
)

// ErrRateLimiterWait means the request never passed the rate limiter and was not sent.
var ErrRateLimiterWait = errors.New("error while waiting for rate limiter")

func (c *NasClientV2) wait(ctx context.Context, logger klog.Logger, action string) error {
	// The aggregate bucket is a safety cap and keeps NAS_LIMIT_PERSECOND's
	// historical meaning; the action bucket is what matches the service's
	// per-API quota.
	t0 := time.Now()
	if err := c.limiter.Wait(ctx); err != nil {
		return fmt.Errorf("%w: %w", ErrRateLimiterWait, err)
	}
	if t := time.Since(t0); t > longThrottleLatency {
		logger.V(3).Info("throttled NAS request", "action", action, "scope", "aggregate", "elapsed", t)
	}
	al, ok := c.actionLimiters[action]
	if !ok || al == nil {
		return nil
	}
	t0 = time.Now()
	if err := al.Wait(ctx); err != nil {
		return fmt.Errorf("%w: %w", ErrRateLimiterWait, err)
	}
	if t := time.Since(t0); t > longThrottleLatency {
		logger.V(3).Info("throttled NAS request", "action", action, "scope", "action", "elapsed", t)
	}
	return nil
}

const (
	// throttleRetryAttempts bounds the retries of one throttled call. The call
	// is also bounded by the request context, which carries the CSI timeout.
	throttleRetryAttempts = 5
	throttleRetryBase     = 200 * time.Millisecond
	throttleRetryMax      = 3 * time.Second
)

// isThrottlingError reports a request the service rejected because a quota was
// exceeded: HTTP 400/403 with Code Throttling.User, Throttling.Api or
// Throttling.CloudBox. The request did not take effect, so retrying it cannot
// duplicate a resource -- unlike a timeout, where the outcome is unknown.
func isThrottlingError(err error) bool {
	if err == nil {
		return false
	}
	// wrap.NAS turns *tea.SDKError into briefAliError, which exposes the code
	// through ErrorCode(); the raw SDK error is also accepted so a client that
	// is not wrapped is classified the same way.
	var e interface{ ErrorCode() string }
	if errors.As(err, &e) {
		return strings.HasPrefix(strings.ToLower(strings.TrimSpace(e.ErrorCode())), "throttling")
	}
	var teaErr *tea.SDKError
	if errors.As(err, &teaErr) {
		return strings.HasPrefix(strings.ToLower(strings.TrimSpace(tea.StringValue(teaErr.Code))), "throttling")
	}
	return false
}

// throttleBackoff is exponential with jitter: concurrent workers that are
// throttled together must not retry in lockstep.
func throttleBackoff(attempt int) time.Duration {
	d := throttleRetryBase << attempt
	if d > throttleRetryMax || d <= 0 {
		d = throttleRetryMax
	}
	return d/2 + time.Duration(rand.Int64N(int64(d/2)+1))
}

// doCall paces one OpenAPI call with its action bucket, retries a throttled
// call, and feeds both outcomes to the adaptive limiter.
func doCall[T any](ctx context.Context, c *NasClientV2, action string, call func() (T, error)) (T, error) {
	var zero T
	logger := klog.FromContext(ctx)
	for attempt := 0; ; attempt++ {
		if err := c.wait(ctx, logger, action); err != nil {
			return zero, err
		}
		resp, err := call()
		if err == nil {
			c.onSuccess(action)
			return resp, nil
		}
		if !isThrottlingError(err) || attempt >= throttleRetryAttempts {
			return zero, err
		}
		c.onThrottle(action)
		delay := throttleBackoff(attempt)
		logger.V(2).Info("NAS OpenAPI throttled, retrying",
			"action", action, "attempt", attempt+1, "delay", delay, "error", err.Error())
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return zero, ctx.Err()
		case <-timer.C:
		}
	}
}

func doCallErr(ctx context.Context, c *NasClientV2, action string, call func() error) error {
	_, err := doCall(ctx, c, action, func() (struct{}, error) {
		return struct{}{}, call()
	})
	return err
}

func (c *NasClientV2) onSuccess(action string) {
	if al, ok := c.actionLimiters[action]; ok && al != nil {
		al.OnSuccess()
	}
}

func (c *NasClientV2) onThrottle(action string) {
	if al, ok := c.actionLimiters[action]; ok && al != nil {
		al.OnThrottle()
	}
}

func (c *NasClientV2) CreateDir(ctx context.Context, req *sdk.CreateDirRequest) error {
	return doCallErr(ctx, c, "CreateDir", func() error {
		_, err := c.client.CreateDirWithContext(ctx, req, &dara.RuntimeOptions{})
		return err
	})
}

var ErrNotSuccess = errors.New("response indicates a failure")

func (c *NasClientV2) SetDirQuota(ctx context.Context, req *sdk.SetDirQuotaRequest) error {
	return doCallErr(ctx, c, "SetDirQuota", func() error {
		resp, err := c.client.SetDirQuotaWithContext(ctx, req, &dara.RuntimeOptions{})
		if err == nil && resp.Body != nil && !tea.BoolValue(resp.Body.Success) {
			err = ErrNotSuccess
		}
		return err
	})
}

func (c *NasClientV2) CancelDirQuota(ctx context.Context, req *sdk.CancelDirQuotaRequest) error {
	return doCallErr(ctx, c, "CancelDirQuota", func() error {
		resp, err := c.client.CancelDirQuotaWithContext(ctx, req, &dara.RuntimeOptions{})
		if err == nil {
			if resp.Body != nil && !tea.BoolValue(resp.Body.Success) {
				err = ErrNotSuccess
			}
		} else if cloud.ErrorCodeV2(err) == "InvalidParameter.QuotaNotExistOnPath" {
			// ignore err if quota not exists
			err = nil
		}
		return err
	})
}

func (c *NasClientV2) GetRecycleBinAttribute(ctx context.Context, filesystemId string) (*sdk.GetRecycleBinAttributeResponse, error) {
	req := &sdk.GetRecycleBinAttributeRequest{FileSystemId: &filesystemId}
	return doCall(ctx, c, "GetRecycleBinAttribute", func() (*sdk.GetRecycleBinAttributeResponse, error) {
		return c.client.GetRecycleBinAttributeWithContext(ctx, req, &dara.RuntimeOptions{})
	})
}

func (c *NasClientV2) CreateAccesspoint(ctx context.Context, req *sdk.CreateAccessPointRequest) (*sdk.CreateAccessPointResponse, error) {
	return doCall(ctx, c, "CreateAccessPoint", func() (*sdk.CreateAccessPointResponse, error) {
		return c.client.CreateAccessPointWithContext(ctx, req, &dara.RuntimeOptions{})
	})
}

func (c *NasClientV2) DeleteAccesspoint(ctx context.Context, filesystemId, accessPointId string) error {
	req := &sdk.DeleteAccessPointRequest{
		AccessPointId: &accessPointId,
		FileSystemId:  &filesystemId,
	}
	// Pass ctx all the way to HTTP, including after an uninterruptible credential
	// refresh. An expired context must never allow a late delete onto the wire.
	return doCallErr(ctx, c, "DeleteAccessPoint", func() error {
		_, err := c.client.DeleteAccessPointWithContext(ctx, req, &dara.RuntimeOptions{})
		return err
	})
}

func (c *NasClientV2) DescribeAccesspoint(ctx context.Context, filesystemId, accessPointId string) (*sdk.DescribeAccessPointResponse, error) {
	req := &sdk.DescribeAccessPointRequest{
		AccessPointId: &accessPointId,
		FileSystemId:  &filesystemId,
	}
	return doCall(ctx, c, "DescribeAccessPoint", func() (*sdk.DescribeAccessPointResponse, error) {
		return c.client.DescribeAccessPointWithContext(ctx, req, &dara.RuntimeOptions{})
	})
}

func (c *NasClientV2) ListAccesspoints(ctx context.Context, req *sdk.ListAccessPointsRequest) (*sdk.ListAccessPointsResponse, error) {
	return doCall(ctx, c, "ListAccessPoints", func() (*sdk.ListAccessPointsResponse, error) {
		return c.client.ListAccessPointsWithContext(ctx, req, &dara.RuntimeOptions{})
	})
}

func (c *NasClientV2) DescribeFileSystems(ctx context.Context, filesystemID string) (*sdk.DescribeFileSystemsResponse, error) {
	req := &sdk.DescribeFileSystemsRequest{FileSystemId: &filesystemID}
	return doCall(ctx, c, "DescribeFileSystems", func() (*sdk.DescribeFileSystemsResponse, error) {
		return c.client.DescribeFileSystemsWithContext(ctx, req, &dara.RuntimeOptions{})
	})
}

func (c *NasClientV2) CreateAgenticSpace(ctx context.Context, req *sdk.CreateAgenticSpaceRequest) (*sdk.CreateAgenticSpaceResponse, error) {
	return doCall(ctx, c, "CreateAgenticSpace", func() (*sdk.CreateAgenticSpaceResponse, error) {
		return c.client.CreateAgenticSpaceWithContext(ctx, req, &dara.RuntimeOptions{})
	})
}

func (c *NasClientV2) GetAgenticSpace(ctx context.Context, req *sdk.GetAgenticSpaceRequest) (*sdk.GetAgenticSpaceResponse, error) {
	return doCall(ctx, c, "GetAgenticSpace", func() (*sdk.GetAgenticSpaceResponse, error) {
		return c.client.GetAgenticSpaceWithContext(ctx, req, &dara.RuntimeOptions{})
	})
}

func (c *NasClientV2) DescribeAgenticSpaces(ctx context.Context, req *sdk.DescribeAgenticSpacesRequest) (*sdk.DescribeAgenticSpacesResponse, error) {
	return doCall(ctx, c, "DescribeAgenticSpaces", func() (*sdk.DescribeAgenticSpacesResponse, error) {
		return c.client.DescribeAgenticSpacesWithContext(ctx, req, &dara.RuntimeOptions{})
	})
}

func (c *NasClientV2) DeleteAgenticSpace(ctx context.Context, req *sdk.DeleteAgenticSpaceRequest) (*sdk.DeleteAgenticSpaceResponse, error) {
	return doCall(ctx, c, "DeleteAgenticSpace", func() (*sdk.DeleteAgenticSpaceResponse, error) {
		return c.client.DeleteAgenticSpaceWithContext(ctx, req, &dara.RuntimeOptions{})
	})
}

func (c *NasClientV2) SetAgenticSpaceQuota(ctx context.Context, req *sdk.SetAgenticSpaceQuotaRequest) (*sdk.SetAgenticSpaceQuotaResponse, error) {
	return doCall(ctx, c, "SetAgenticSpaceQuota", func() (*sdk.SetAgenticSpaceQuotaResponse, error) {
		return c.client.SetAgenticSpaceQuotaWithContext(ctx, req, &dara.RuntimeOptions{})
	})
}
