// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedHostAutoRenewAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDedicatedHostAutoRenewAttributeResponseBody
	GetRequestId() *string
}

type ModifyDedicatedHostAutoRenewAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2A4EA075-CB5B-41B7-B0EB-70D339F6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDedicatedHostAutoRenewAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedHostAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostAutoRenewAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDedicatedHostAutoRenewAttributeResponseBody) SetRequestId(v string) *ModifyDedicatedHostAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDedicatedHostAutoRenewAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
