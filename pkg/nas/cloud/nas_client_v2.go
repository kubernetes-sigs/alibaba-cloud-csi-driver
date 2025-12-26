package cloud

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	alicred_old "github.com/aliyun/credentials-go/credentials"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/wrap"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/credentials"
	utilshttp "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/http"
	"golang.org/x/time/rate"
	"k8s.io/klog/v2"
)

const (
	connTimeout = 10
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
	region  string
	limiter *rate.Limiter
	client  cloud.NasInterface
}

var (
	// longThrottleLatency defines threshold for logging requests. All requests being
	// throttled (via the provided rateLimiter) for more than longThrottleLatency will
	// be logged.
	longThrottleLatency = 250 * time.Millisecond
)

func (c *NasClientV2) wait(ctx context.Context, logger klog.Logger) error {
	t0 := time.Now()
	if err := c.limiter.Wait(ctx); err != nil {
		return fmt.Errorf("error while waiting for rate limiter: %w", err)
	}
	t := time.Since(t0)
	if t > longThrottleLatency {
		logger.V(3).Info("throttled NAS request", "elapsed", t)
	}
	return nil
}

func (c *NasClientV2) CreateDir(ctx context.Context, req *sdk.CreateDirRequest) error {
	logger := klog.FromContext(ctx)
	if err := c.wait(ctx, logger); err != nil {
		return err
	}
	_, err := wrap.V2(logger, c.client.CreateDir)(req)
	return err
}

var ErrNotSuccess = errors.New("response indicates a failure")

func (c *NasClientV2) SetDirQuota(ctx context.Context, req *sdk.SetDirQuotaRequest) error {
	logger := klog.FromContext(ctx)
	if err := c.wait(ctx, logger); err != nil {
		return err
	}
	resp, err := wrap.V2(logger, c.client.SetDirQuota)(req)
	if err == nil && resp.Body != nil && !tea.BoolValue(resp.Body.Success) {
		err = ErrNotSuccess
	}
	return err
}

func (c *NasClientV2) CancelDirQuota(ctx context.Context, req *sdk.CancelDirQuotaRequest) error {
	logger := klog.FromContext(ctx)
	if err := c.wait(ctx, logger); err != nil {
		return err
	}
	resp, err := wrap.V2(logger, c.client.CancelDirQuota)(req)
	if err == nil {
		if !tea.BoolValue(resp.Body.Success) {
			err = ErrNotSuccess
		}
	} else {
		if errors.Is(err, wrap.ErrorCode("InvalidParameter.QuotaNotExistOnPath")) {
			// ignore err if quota not exists
			err = nil
		}
	}
	return err
}

func (c *NasClientV2) GetRecycleBinAttribute(ctx context.Context, filesystemId string) (*sdk.GetRecycleBinAttributeResponse, error) {
	logger := klog.FromContext(ctx)
	if err := c.wait(ctx, logger); err != nil {
		return nil, err
	}
	req := &sdk.GetRecycleBinAttributeRequest{FileSystemId: &filesystemId}
	return wrap.V2(logger, c.client.GetRecycleBinAttribute)(req)
}

func (c *NasClientV2) CreateAccesspoint(ctx context.Context, req *sdk.CreateAccessPointRequest) (*sdk.CreateAccessPointResponse, error) {
	logger := klog.FromContext(ctx)
	if err := c.wait(ctx, logger); err != nil {
		return nil, err
	}
	return wrap.V2(logger, c.client.CreateAccessPoint)(req)
}

func (c *NasClientV2) DeleteAccesspoint(ctx context.Context, filesystemId, accessPointId string) error {
	logger := klog.FromContext(ctx)
	if err := c.wait(ctx, logger); err != nil {
		return nil
	}
	req := &sdk.DeleteAccessPointRequest{
		AccessPointId: &accessPointId,
		FileSystemId:  &filesystemId,
	}
	_, err := wrap.V2(logger, c.client.DeleteAccessPoint)(req)
	return err
}

func (c *NasClientV2) DescribeAccesspoint(ctx context.Context, filesystemId, accessPointId string) (*sdk.DescribeAccessPointResponse, error) {
	logger := klog.FromContext(ctx)
	if err := c.wait(ctx, logger); err != nil {
		return nil, err
	}
	return wrap.V2(logger, c.client.DescribeAccessPoint)(&sdk.DescribeAccessPointRequest{
		AccessPointId: &accessPointId,
		FileSystemId:  &filesystemId,
	})
}

func (c *NasClientV2) DescribeFileSystems(ctx context.Context, filesystemID string) (*sdk.DescribeFileSystemsResponse, error) {
	logger := klog.FromContext(ctx)
	if err := c.wait(ctx, logger); err != nil {
		return nil, err
	}
	return wrap.V2(logger, c.client.DescribeFileSystems)(&sdk.DescribeFileSystemsRequest{
		FileSystemId: &filesystemID,
	})
}
