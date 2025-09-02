// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeType(v string) *DescribeNodeTypeRequest
	GetNodeType() *string
}

type DescribeNodeTypeRequest struct {
	// example:
	//
	// standard
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s DescribeNodeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeTypeRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeNodeTypeRequest) SetNodeType(v string) *DescribeNodeTypeRequest {
	s.NodeType = &v
	return s
}

func (s *DescribeNodeTypeRequest) Validate() error {
	return dara.Validate(s)
}
