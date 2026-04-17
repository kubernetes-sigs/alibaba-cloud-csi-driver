// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLifecyclePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLifecyclePolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateLifecyclePolicyResponseBody
	GetSuccess() *bool
}

type UpdateLifecyclePolicyResponseBody struct {
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateLifecyclePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLifecyclePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLifecyclePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLifecyclePolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateLifecyclePolicyResponseBody) SetRequestId(v string) *UpdateLifecyclePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLifecyclePolicyResponseBody) SetSuccess(v bool) *UpdateLifecyclePolicyResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateLifecyclePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
