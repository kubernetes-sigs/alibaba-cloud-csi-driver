// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVpcResponseBody
	GetRequestId() *string
	SetRouteTableId(v string) *CreateVpcResponseBody
	GetRouteTableId() *string
	SetVRouterId(v string) *CreateVpcResponseBody
	GetVRouterId() *string
	SetVpcId(v string) *CreateVpcResponseBody
	GetVpcId() *string
}

type CreateVpcResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	VRouterId    *string `json:"VRouterId,omitempty" xml:"VRouterId,omitempty"`
	VpcId        *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVpcResponseBody) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *CreateVpcResponseBody) GetVRouterId() *string {
	return s.VRouterId
}

func (s *CreateVpcResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateVpcResponseBody) SetRequestId(v string) *CreateVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcResponseBody) SetRouteTableId(v string) *CreateVpcResponseBody {
	s.RouteTableId = &v
	return s
}

func (s *CreateVpcResponseBody) SetVRouterId(v string) *CreateVpcResponseBody {
	s.VRouterId = &v
	return s
}

func (s *CreateVpcResponseBody) SetVpcId(v string) *CreateVpcResponseBody {
	s.VpcId = &v
	return s
}

func (s *CreateVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
