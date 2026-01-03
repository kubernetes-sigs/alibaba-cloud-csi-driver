// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualBorderRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVirtualBorderRouterResponseBody
	GetRequestId() *string
	SetVbrId(v string) *CreateVirtualBorderRouterResponseBody
	GetVbrId() *string
}

type CreateVirtualBorderRouterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VbrId     *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
}

func (s CreateVirtualBorderRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualBorderRouterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirtualBorderRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVirtualBorderRouterResponseBody) GetVbrId() *string {
	return s.VbrId
}

func (s *CreateVirtualBorderRouterResponseBody) SetRequestId(v string) *CreateVirtualBorderRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVirtualBorderRouterResponseBody) SetVbrId(v string) *CreateVirtualBorderRouterResponseBody {
	s.VbrId = &v
	return s
}

func (s *CreateVirtualBorderRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
