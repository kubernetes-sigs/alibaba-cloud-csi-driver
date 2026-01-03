// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSavingsPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *CreateSavingsPlanResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateSavingsPlanResponseBody
	GetRequestId() *string
	SetSavingsPlanId(v string) *CreateSavingsPlanResponseBody
	GetSavingsPlanId() *string
}

type CreateSavingsPlanResponseBody struct {
	OrderId       *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SavingsPlanId *string `json:"SavingsPlanId,omitempty" xml:"SavingsPlanId,omitempty"`
}

func (s CreateSavingsPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSavingsPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSavingsPlanResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateSavingsPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSavingsPlanResponseBody) GetSavingsPlanId() *string {
	return s.SavingsPlanId
}

func (s *CreateSavingsPlanResponseBody) SetOrderId(v string) *CreateSavingsPlanResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateSavingsPlanResponseBody) SetRequestId(v string) *CreateSavingsPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSavingsPlanResponseBody) SetSavingsPlanId(v string) *CreateSavingsPlanResponseBody {
	s.SavingsPlanId = &v
	return s
}

func (s *CreateSavingsPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
