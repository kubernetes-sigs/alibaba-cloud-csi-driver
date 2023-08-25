package disk

import "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"

type ECSInterface interface {
	DescribeAvailableResource(request *ecs.DescribeAvailableResourceRequest) (response *ecs.DescribeAvailableResourceResponse, err error)
}
