package cloud

import (
	"context"
	"errors"
	"fmt"
	"os"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	alicred_old "github.com/aliyun/credentials-go/credentials"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/credentials"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/interfaces"
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

func (c *NasClientV2) CreateDir(req *sdk.CreateDirRequest) error {
	if err := c.limiter.Wait(context.TODO()); err != nil {
		return fmt.Errorf("error while waiting for rate limiter: %w", err)
	}
	resp, err := c.client.CreateDir(req)
	logger := klog.Background().WithValues("request", req, "response", resp)
	if err == nil {
		logger.V(2).Info("nas:CreateDir succeeded")
	} else {
		logger.Error(err, "nas:CreateDir failed")
	}
	return err
}

func (c *NasClientV2) SetDirQuota(req *sdk.SetDirQuotaRequest) error {
	if err := c.limiter.Wait(context.TODO()); err != nil {
		return fmt.Errorf("error while waiting for rate limiter: %w", err)
	}
	resp, err := c.client.SetDirQuota(req)
	if err == nil && resp.Body != nil && !tea.BoolValue(resp.Body.Success) {
		err = errors.New("response indicates a failure")
	}
	logger := klog.Background().WithValues("request", req, "response", resp)
	if err == nil {
		logger.V(2).Info("nas:SetDirQuota succeeded")
	} else {
		logger.Error(err, "nas:SetDirQuota failed")
	}
	return err
}

func (c *NasClientV2) CancelDirQuota(req *sdk.CancelDirQuotaRequest) error {
	if err := c.limiter.Wait(context.TODO()); err != nil {
		return fmt.Errorf("error while waiting for rate limiter: %w", err)
	}
	resp, err := c.client.CancelDirQuota(req)
	if err == nil {
		if !tea.BoolValue(resp.Body.Success) {
			err = errors.New("response indicates a failure")
		}
	} else {
		_err, ok := err.(*tea.SDKError)
		if ok && tea.StringValue(_err.Code) == "InvalidParameter.QuotaNotExistOnPath" {
			// ignore err if quota not exists
			err = nil
		}
	}
	logger := klog.Background().WithValues("request", req, "response", resp)
	if err == nil {
		logger.V(2).Info("nas:CancelDirQuota succeeded")
	} else {
		logger.Error(err, "nas:CancelDirQuota failed")
	}
	return err
}

func (c *NasClientV2) GetRecycleBinAttribute(filesystemId string) (*sdk.GetRecycleBinAttributeResponse, error) {
	if err := c.limiter.Wait(context.TODO()); err != nil {
		return nil, fmt.Errorf("error while waiting for rate limiter: %w", err)
	}
	req := &sdk.GetRecycleBinAttributeRequest{FileSystemId: &filesystemId}
	return c.client.GetRecycleBinAttribute(req)
}

func (c *NasClientV2) CreateAccesspoint(req *sdk.CreateAccessPointRequest) (*sdk.CreateAccessPointResponse, error) {
	if err := c.limiter.Wait(context.TODO()); err != nil {
		return nil, fmt.Errorf("error while waiting for rate limiter: %w", err)
	}
	resp, err := c.client.CreateAccessPoint(req)
	logger := klog.Background().WithValues("request", req, "response", resp)
	if err == nil {
		logger.V(2).Info("nas:CreateAccessPoint succeeded")
	} else {
		logger.Error(err, "nas:CreateAccessPoint failed")
	}
	return resp, err
}

func (c *NasClientV2) DeleteAccesspoint(filesystemId, accessPointId string) error {
	if err := c.limiter.Wait(context.TODO()); err != nil {
		return fmt.Errorf("error while waiting for rate limiter: %w", err)
	}
	req := &sdk.DeleteAccessPointRequest{
		AccessPointId: &accessPointId,
		FileSystemId:  &filesystemId,
	}
	resp, err := c.client.DeleteAccessPoint(req)
	logger := klog.Background().WithValues("request", req, "response", resp)
	if err == nil {
		logger.V(2).Info("nas:DeleteAccessPoint succeeded")
	} else {
		logger.Error(err, "nas:DeleteAccessPoint failed")
	}
	return err
}

func (c *NasClientV2) DescribeAccesspoint(filesystemId, accessPointId string) (*sdk.DescribeAccessPointResponse, error) {
	if err := c.limiter.Wait(context.TODO()); err != nil {
		return nil, fmt.Errorf("error while waiting for rate limiter: %w", err)
	}
	return c.client.DescribeAccessPoint(&sdk.DescribeAccessPointRequest{
		AccessPointId: &accessPointId,
		FileSystemId:  &filesystemId,
	})
}

func (c *NasClientV2) DescribeFileSystems(filesystemID string) (*sdk.DescribeFileSystemsResponse, error) {
	if err := c.limiter.Wait(context.TODO()); err != nil {
		return nil, fmt.Errorf("error while waiting for rate limiter: %w", err)
	}
	return c.client.DescribeFileSystems(&sdk.DescribeFileSystemsRequest{
		FileSystemId: &filesystemID,
	})
}

func IsAccessPointNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	sdkErr, ok := err.(*tea.SDKError)
	return ok && tea.StringValue(sdkErr.Code) == "NotFound"
}
