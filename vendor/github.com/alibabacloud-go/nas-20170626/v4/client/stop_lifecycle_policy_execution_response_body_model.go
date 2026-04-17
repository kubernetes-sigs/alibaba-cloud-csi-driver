// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLifecyclePolicyExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopLifecyclePolicyExecutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopLifecyclePolicyExecutionResponseBody
	GetSuccess() *bool
}

type StopLifecyclePolicyExecutionResponseBody struct {
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopLifecyclePolicyExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopLifecyclePolicyExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StopLifecyclePolicyExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopLifecyclePolicyExecutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopLifecyclePolicyExecutionResponseBody) SetRequestId(v string) *StopLifecyclePolicyExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopLifecyclePolicyExecutionResponseBody) SetSuccess(v bool) *StopLifecyclePolicyExecutionResponseBody {
	s.Success = &v
	return s
}

func (s *StopLifecyclePolicyExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
