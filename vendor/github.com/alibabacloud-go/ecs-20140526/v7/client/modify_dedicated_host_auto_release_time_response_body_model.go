// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedHostAutoReleaseTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDedicatedHostAutoReleaseTimeResponseBody
	GetRequestId() *string
}

type ModifyDedicatedHostAutoReleaseTimeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDedicatedHostAutoReleaseTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedHostAutoReleaseTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostAutoReleaseTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDedicatedHostAutoReleaseTimeResponseBody) SetRequestId(v string) *ModifyDedicatedHostAutoReleaseTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDedicatedHostAutoReleaseTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
