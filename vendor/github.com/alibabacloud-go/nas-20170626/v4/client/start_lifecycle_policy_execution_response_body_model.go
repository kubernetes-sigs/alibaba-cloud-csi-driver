// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLifecyclePolicyExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartLifecyclePolicyExecutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartLifecyclePolicyExecutionResponseBody
	GetSuccess() *bool
}

type StartLifecyclePolicyExecutionResponseBody struct {
	// example:
	//
	// 70EACC9C-D07A-4A34-ADA4-77506C42****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartLifecyclePolicyExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartLifecyclePolicyExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StartLifecyclePolicyExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartLifecyclePolicyExecutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartLifecyclePolicyExecutionResponseBody) SetRequestId(v string) *StartLifecyclePolicyExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartLifecyclePolicyExecutionResponseBody) SetSuccess(v bool) *StartLifecyclePolicyExecutionResponseBody {
	s.Success = &v
	return s
}

func (s *StartLifecyclePolicyExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
