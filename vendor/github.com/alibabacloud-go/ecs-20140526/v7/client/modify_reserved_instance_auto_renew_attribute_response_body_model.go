// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReservedInstanceAutoRenewAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyReservedInstanceAutoRenewAttributeResponseBody
	GetRequestId() *string
}

type ModifyReservedInstanceAutoRenewAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2A4EA075-CB5B-41B7-B0EB-70D339F6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyReservedInstanceAutoRenewAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyReservedInstanceAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyReservedInstanceAutoRenewAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyReservedInstanceAutoRenewAttributeResponseBody) SetRequestId(v string) *ModifyReservedInstanceAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyReservedInstanceAutoRenewAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
