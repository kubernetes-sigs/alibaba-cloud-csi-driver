// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNatGatewayResponseBody
	GetRequestId() *string
}

type DeleteNatGatewayResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNatGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNatGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNatGatewayResponseBody) SetRequestId(v string) *DeleteNatGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNatGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
