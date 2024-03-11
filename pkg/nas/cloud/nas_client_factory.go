package cloud

import (
	"os"
	"strconv"

	sdkv1 "github.com/aliyun/alibaba-cloud-sdk-go/services/nas"
	"github.com/sirupsen/logrus"
	"go.uber.org/ratelimit"
)

const defaultQps = 2

type NasClientFactory struct {
	// ratelimiter only takes effect on v2 client
	limiter ratelimit.Limiter
}

func NewNasClientFactory() *NasClientFactory {
	qps := defaultQps
	value, ok := os.LookupEnv("NAS_LIMIT_PERSECOND")
	if ok {
		v, err := strconv.Atoi(value)
		if err != nil {
			logrus.Errorf("invalid NAS_LIMIT_PERSECOND %q", value)
		} else {
			qps = v
			logrus.Infof("set nas QPS to %d", qps)
		}
	}
	return &NasClientFactory{
		limiter: ratelimit.New(qps),
	}
}

// V2 gets a NAS OpenAPI client sourced from github.com/alibabacloud-go/nas-20170626.
// As github.com/aliyun/alibaba-cloud-sdk-go/services/nas won't be updated with new NAS APIs (e.g., access points),
// we will fully migrate to github.com/alibabacloud-go/nas-20170626 in the future.
func (fac *NasClientFactory) V2(region string) (*NasClientV2, error) {
	client, err := newNasClientV2(region)
	if err != nil {
		return nil, err
	}
	return &NasClientV2{
		region:  region,
		limiter: fac.limiter,
		client:  client,
	}, nil
}

// Deprecated: NAS openapi client provided by github.com/aliyun/alibaba-cloud-sdk-go/services/nas.
func (fac *NasClientFactory) V1(region string) (*sdkv1.Client, error) {
	return newNasClientV1(region)
}
