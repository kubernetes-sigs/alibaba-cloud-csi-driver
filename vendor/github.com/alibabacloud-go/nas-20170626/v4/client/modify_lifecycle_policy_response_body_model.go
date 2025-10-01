// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLifecyclePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyLifecyclePolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyLifecyclePolicyResponseBody
	GetSuccess() *bool
}

type ModifyLifecyclePolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BC7C825C-5F65-4B56-BEF6-98C56C7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyLifecyclePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLifecyclePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLifecyclePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLifecyclePolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyLifecyclePolicyResponseBody) SetRequestId(v string) *ModifyLifecyclePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLifecyclePolicyResponseBody) SetSuccess(v bool) *ModifyLifecyclePolicyResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyLifecyclePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
