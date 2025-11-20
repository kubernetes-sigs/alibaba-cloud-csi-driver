package cloud

import (
	eflo "github.com/alibabacloud-go/eflo-controller-20221215/v3/client"
)

type EFLOInterface interface {
	DescribeNodeType(req *eflo.DescribeNodeTypeRequest) (*eflo.DescribeNodeTypeResponse, error)
	DescribeNode(request *eflo.DescribeNodeRequest) (*eflo.DescribeNodeResponse, error)
}
