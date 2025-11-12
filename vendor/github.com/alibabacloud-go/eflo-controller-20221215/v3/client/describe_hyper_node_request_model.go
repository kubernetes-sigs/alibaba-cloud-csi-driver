// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHyperNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHyperNodeId(v string) *DescribeHyperNodeRequest
	GetHyperNodeId() *string
}

type DescribeHyperNodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e01-cn-zvp2tgykr08
	HyperNodeId *string `json:"HyperNodeId,omitempty" xml:"HyperNodeId,omitempty"`
}

func (s DescribeHyperNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHyperNodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeHyperNodeRequest) GetHyperNodeId() *string {
	return s.HyperNodeId
}

func (s *DescribeHyperNodeRequest) SetHyperNodeId(v string) *DescribeHyperNodeRequest {
	s.HyperNodeId = &v
	return s
}

func (s *DescribeHyperNodeRequest) Validate() error {
	return dara.Validate(s)
}
