package cloud

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type Common interface {
	ProcessCommonRequest(request *requests.CommonRequest) (response *responses.CommonResponse, err error)
}
