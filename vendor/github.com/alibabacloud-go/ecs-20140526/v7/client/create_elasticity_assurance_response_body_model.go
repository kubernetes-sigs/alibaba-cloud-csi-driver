// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateElasticityAssuranceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *CreateElasticityAssuranceResponseBody
	GetOrderId() *string
	SetPrivatePoolOptionsId(v string) *CreateElasticityAssuranceResponseBody
	GetPrivatePoolOptionsId() *string
	SetRequestId(v string) *CreateElasticityAssuranceResponseBody
	GetRequestId() *string
}

type CreateElasticityAssuranceResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 1234567890
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The elasticity assurance ID.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	PrivatePoolOptionsId *string `json:"PrivatePoolOptionsId,omitempty" xml:"PrivatePoolOptionsId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateElasticityAssuranceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticityAssuranceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateElasticityAssuranceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateElasticityAssuranceResponseBody) GetPrivatePoolOptionsId() *string {
	return s.PrivatePoolOptionsId
}

func (s *CreateElasticityAssuranceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateElasticityAssuranceResponseBody) SetOrderId(v string) *CreateElasticityAssuranceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateElasticityAssuranceResponseBody) SetPrivatePoolOptionsId(v string) *CreateElasticityAssuranceResponseBody {
	s.PrivatePoolOptionsId = &v
	return s
}

func (s *CreateElasticityAssuranceResponseBody) SetRequestId(v string) *CreateElasticityAssuranceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateElasticityAssuranceResponseBody) Validate() error {
	return dara.Validate(s)
}
