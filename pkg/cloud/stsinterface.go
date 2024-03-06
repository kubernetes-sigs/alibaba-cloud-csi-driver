package cloud

import "github.com/aliyun/alibaba-cloud-sdk-go/services/sts"

type STSInterface interface {
	GetCallerIdentity(request *sts.GetCallerIdentityRequest) (response *sts.GetCallerIdentityResponse, err error)
}
