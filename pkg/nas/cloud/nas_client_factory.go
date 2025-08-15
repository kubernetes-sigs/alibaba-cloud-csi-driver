package cloud

import (
	"os"
	"strconv"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/interfaces"
	"golang.org/x/time/rate"
	"k8s.io/klog/v2"
)

const defaultQps = 2

type NasClientFactory struct {
	// ratelimiter only takes effect on v2 client
	limiter *rate.Limiter
}

func NewNasClientFactory() *NasClientFactory {
	qps := defaultQps
	value, ok := os.LookupEnv("NAS_LIMIT_PERSECOND")
	if ok {
		v, err := strconv.Atoi(value)
		if err != nil {
			klog.Errorf("invalid NAS_LIMIT_PERSECOND %q", value)
		} else {
			qps = v
			klog.Infof("set nas QPS to %d", qps)
		}
	}
	return &NasClientFactory{
		limiter: rate.NewLimiter(rate.Limit(qps), 10),
	}
}

// V2 gets a NAS OpenAPI client sourced from github.com/alibabacloud-go/nas-20170626.
// As github.com/aliyun/alibaba-cloud-sdk-go/services/nas won't be updated with new NAS APIs (e.g., access points),
// we will fully migrate to github.com/alibabacloud-go/nas-20170626 in the future.
func (fac *NasClientFactory) V2(region string) (interfaces.NasClientV2Interface, error) {
	client, err := NewNasClientV2(region)
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
func (fac *NasClientFactory) V1(region string) (interfaces.NasV1Interface, error) {
	return newNasClientV1(region)
}
