package cloud

import (
	"context"
	"fmt"
	"strings"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

type OSSClient struct {
	Ossc          cloud.OSSInterface
	ClusterRegion string
}

// OSS location & OSS region mapping
// refer to https://help.aliyun.com/zh/oss/user-guide/regions-and-endpoints
var location2region map[string]string = map[string]string{
	"oss-cn-hzjbp":                  "cn-hangzhou-finance",
	"oss-cn-hzfinance":              "cn-hangzhou-finance",
	"oss-cn-shanghai-finance-1-pub": "cn-shanghai-finance-1",
	"oss-cn-szfinance":              "cn-shenzhen-finance-1",
	"oss-cn-beijing-finance-1-pub":  "cn-beijing-finance-1",
}

// get OSS client, region is cluster region-id
func NewOSSClient(region string) *OSSClient {
	// Init OSS Client
	ac := utils.GetAccessControl()
	provider := credentials.NewStaticCredentialsProvider(ac.AccessKeyID, ac.AccessKeySecret, ac.StsToken)
	cfg := oss.LoadDefaultConfig().WithCredentialsProvider(provider).WithRegion(region)
	ossc := oss.NewClient(cfg)
	return &OSSClient{Ossc: ossc, ClusterRegion: region}
}

func (c *OSSClient) UpdateToken() {
	ac := utils.GetAccessControl()
	provider := credentials.NewStaticCredentialsProvider(ac.AccessKeyID, ac.AccessKeySecret, ac.StsToken)
	cfg := oss.LoadDefaultConfig().WithCredentialsProvider(provider).WithRegion(c.ClusterRegion)
	ossc := oss.NewClient(cfg)
	c.Ossc = ossc
}

func getOSSRegion(location *string) (region string) {
	if location == nil {
		return
	}
	var ok bool
	region, ok = location2region[*location]
	if ok {
		return
	}
	region = strings.TrimPrefix(*location, "oss-")
	return region
}

func (c *OSSClient) GetBucketRegion(ctx context.Context, bucketName string) (region string, err error) {

	request := &oss.GetBucketInfoRequest{
		Bucket: oss.Ptr(bucketName),
	}
	result, err := c.Ossc.GetBucketInfo(ctx, request)
	if err != nil {
		return "", err
	}
	if result == nil || result.BucketInfo.Name == nil {
		return "", fmt.Errorf("bucket %s not found", bucketName)
	}

	region = getOSSRegion(result.BucketInfo.Location)
	return
}
