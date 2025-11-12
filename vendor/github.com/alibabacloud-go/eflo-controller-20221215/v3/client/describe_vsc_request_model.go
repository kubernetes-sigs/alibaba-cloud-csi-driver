// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVscRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVscId(v string) *DescribeVscRequest
	GetVscId() *string
}

type DescribeVscRequest struct {
	// The VSC ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsc-001
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s DescribeVscRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscRequest) GoString() string {
	return s.String()
}

func (s *DescribeVscRequest) GetVscId() *string {
	return s.VscId
}

func (s *DescribeVscRequest) SetVscId(v string) *DescribeVscRequest {
	s.VscId = &v
	return s
}

func (s *DescribeVscRequest) Validate() error {
	return dara.Validate(s)
}
