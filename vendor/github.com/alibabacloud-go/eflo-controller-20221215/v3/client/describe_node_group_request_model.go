// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeGroupId(v string) *DescribeNodeGroupRequest
	GetNodeGroupId() *string
}

type DescribeNodeGroupRequest struct {
	// This parameter is required.
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s DescribeNodeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeGroupRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *DescribeNodeGroupRequest) SetNodeGroupId(v string) *DescribeNodeGroupRequest {
	s.NodeGroupId = &v
	return s
}

func (s *DescribeNodeGroupRequest) Validate() error {
	return dara.Validate(s)
}
