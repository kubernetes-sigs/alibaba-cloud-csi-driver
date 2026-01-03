// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedHostAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDedicatedHostAttributeResponseBody
	GetRequestId() *string
}

type ModifyDedicatedHostAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2A4EA075-CB5B-41B7-B0EB-70D339F6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDedicatedHostAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedHostAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDedicatedHostAttributeResponseBody) SetRequestId(v string) *ModifyDedicatedHostAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDedicatedHostAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
