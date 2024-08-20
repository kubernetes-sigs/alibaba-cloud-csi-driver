package cloud

import (
	"errors"
	"fmt"
	"os"
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sdk "github.com/alibabacloud-go/nas-20170626/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/interfaces"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	utilshttp "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/http"
	"github.com/sirupsen/logrus"
	"go.uber.org/ratelimit"
)

func newNasClientV2(region string) (*sdk.Client, error) {
	headers := utilshttp.MustParseHeaderEnv("NAS_HEADERS")
	var headersV2 map[string]*string
	if headers != nil {
		headersV2 = utilshttp.MustToV2SDKHeaders(headers)
	}
	config := new(openapi.Config).
		SetUserAgent(KubernetesAlicloudIdentity).
		SetRegionId(region).
		SetGlobalParameters(&openapi.GlobalParameters{
			Queries: map[string]*string{
				"RegionId": &region,
			},
			Headers: headersV2,
		})
	// set credential
	cred, err := utils.GetCredentialV2()
	if err != nil {
		return nil, fmt.Errorf("init credential: %w", err)
	}
	config = config.SetCredential(cred)
	// set endpoint
	ep := os.Getenv("NAS_ENDPOINT")
	if ep == "" {
		ep = GetEndpointForRegion(region)
	}
	config = config.SetEndpoint(ep)
	// set protocol
	scheme := strings.ToUpper(os.Getenv("ALICLOUD_CLIENT_SCHEME"))
	if scheme != "HTTP" {
		scheme = "HTTPS"
	}
	config = config.SetProtocol(scheme)
	// init client
	return sdk.NewClient(config)
}

type NasClientV2 struct {
	region  string
	limiter ratelimit.Limiter
	client  interfaces.NasV2Interface
}

func (c *NasClientV2) CreateDir(req *sdk.CreateDirRequest) error {
	c.limiter.Take()
	resp, err := c.client.CreateDir(req)
	log := logrus.WithFields(logrus.Fields{
		"request":  req,
		"response": resp,
	})
	if err == nil {
		log.Info("nas:CreateDir succeeded")
	} else {
		log.Errorf("nas:CreateDir failed: %v", err)
	}
	return err
}

func (c *NasClientV2) SetDirQuota(req *sdk.SetDirQuotaRequest) error {
	c.limiter.Take()
	resp, err := c.client.SetDirQuota(req)
	if err == nil && resp.Body != nil && !tea.BoolValue(resp.Body.Success) {
		err = errors.New("response indicates a failure")
	}
	log := logrus.WithFields(logrus.Fields{
		"request":  req,
		"response": resp,
	})
	if err == nil {
		log.Info("nas:SetDirQuota succeeded")
	} else {
		log.Errorf("nas:SetDirQuota failed: %v", err)
	}
	return err
}

func (c *NasClientV2) CancelDirQuota(req *sdk.CancelDirQuotaRequest) error {
	c.limiter.Take()
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
	log := logrus.WithFields(logrus.Fields{
		"request":  req,
		"response": resp,
	})
	if err == nil {
		log.Info("nas:CancelDirQuota succeeded")
	} else {
		log.Errorf("nas:CancelDirQuota failed: %v", err)
	}
	return err
}

func (c *NasClientV2) GetRecycleBinAttribute(filesystemId string) (*sdk.GetRecycleBinAttributeResponse, error) {
	c.limiter.Take()
	req := &sdk.GetRecycleBinAttributeRequest{FileSystemId: &filesystemId}
	return c.client.GetRecycleBinAttribute(req)
}

func (c *NasClientV2) CreateAccesspoint(req *sdk.CreateAccessPointRequest) (*sdk.CreateAccessPointResponse, error) {
	c.limiter.Take()
	resp, err := c.client.CreateAccessPoint(req)
	log := logrus.WithFields(logrus.Fields{
		"request":  req,
		"response": resp,
	})
	if err == nil {
		log.Info("nas:CreateAccessPoint succeeded")
	} else {
		log.Errorf("nas:CreateAccessPoint failed: %v", err)
	}
	return resp, err
}

func (c *NasClientV2) DeleteAccesspoint(filesystemId, accessPointId string) error {
	c.limiter.Take()
	req := &sdk.DeleteAccessPointRequest{
		AccessPointId: &accessPointId,
		FileSystemId:  &filesystemId,
	}
	resp, err := c.client.DeleteAccessPoint(req)
	log := logrus.WithFields(logrus.Fields{
		"request":  req,
		"response": resp,
	})
	if err == nil {
		log.Info("nas:DeleteAccessPoint succeeded")
	} else {
		log.Errorf("nas:DeleteAccessPoint failed: %v", err)
	}
	return err
}

func (c *NasClientV2) DescribeAccesspoint(filesystemId, accessPointId string) (*sdk.DescribeAccessPointResponse, error) {
	c.limiter.Take()
	return c.client.DescribeAccessPoint(&sdk.DescribeAccessPointRequest{
		AccessPointId: &accessPointId,
		FileSystemId:  &filesystemId,
	})
}

func IsAccessPointNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	sdkErr, ok := err.(*tea.SDKError)
	return ok && tea.StringValue(sdkErr.Code) == "NotFound"
}
