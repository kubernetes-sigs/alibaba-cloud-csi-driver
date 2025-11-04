package cloud

import sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"

type STSInterface interface {
	GetCallerIdentity() (response *sts20150401.GetCallerIdentityResponse, err error)
}
