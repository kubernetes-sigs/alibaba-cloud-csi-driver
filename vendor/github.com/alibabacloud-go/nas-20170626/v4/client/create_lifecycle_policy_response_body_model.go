// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLifecyclePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLifecyclePolicyId(v string) *CreateLifecyclePolicyResponseBody
	GetLifecyclePolicyId() *string
	SetRequestId(v string) *CreateLifecyclePolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateLifecyclePolicyResponseBody
	GetSuccess() *bool
}

type CreateLifecyclePolicyResponseBody struct {
	// The ID of the lifecycle policy.
	//
	// example:
	//
	// lsp-bp1234567890ab****
	LifecyclePolicyId *string `json:"LifecyclePolicyId,omitempty" xml:"LifecyclePolicyId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BC7C825C-5F65-4B56-BEF6-98C56C7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
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

func (s CreateLifecyclePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLifecyclePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLifecyclePolicyResponseBody) GetLifecyclePolicyId() *string {
	return s.LifecyclePolicyId
}

func (s *CreateLifecyclePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLifecyclePolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateLifecyclePolicyResponseBody) SetLifecyclePolicyId(v string) *CreateLifecyclePolicyResponseBody {
	s.LifecyclePolicyId = &v
	return s
}

func (s *CreateLifecyclePolicyResponseBody) SetRequestId(v string) *CreateLifecyclePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLifecyclePolicyResponseBody) SetSuccess(v bool) *CreateLifecyclePolicyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateLifecyclePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
