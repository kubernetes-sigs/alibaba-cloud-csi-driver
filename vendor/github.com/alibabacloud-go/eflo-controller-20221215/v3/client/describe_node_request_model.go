// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *DescribeNodeRequest
	GetNodeId() *string
}

type DescribeNodeRequest struct {
	// The node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// mock-sn-2060
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeNodeRequest) SetNodeId(v string) *DescribeNodeRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeNodeRequest) Validate() error {
	return dara.Validate(s)
}
