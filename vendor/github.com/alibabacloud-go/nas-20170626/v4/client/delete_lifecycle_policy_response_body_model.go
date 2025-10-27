// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLifecyclePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLifecyclePolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteLifecyclePolicyResponseBody
	GetSuccess() *bool
}

type DeleteLifecyclePolicyResponseBody struct {
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

func (s DeleteLifecyclePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLifecyclePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLifecyclePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLifecyclePolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteLifecyclePolicyResponseBody) SetRequestId(v string) *DeleteLifecyclePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLifecyclePolicyResponseBody) SetSuccess(v bool) *DeleteLifecyclePolicyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteLifecyclePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
