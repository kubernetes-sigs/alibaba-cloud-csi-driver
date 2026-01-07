// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBusinessBehaviorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUserBusinessBehaviorResponseBody
	GetRequestId() *string
	SetStatusValue(v string) *DescribeUserBusinessBehaviorResponseBody
	GetStatusValue() *string
}

type DescribeUserBusinessBehaviorResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatusValue *string `json:"StatusValue,omitempty" xml:"StatusValue,omitempty"`
}

func (s DescribeUserBusinessBehaviorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBusinessBehaviorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserBusinessBehaviorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserBusinessBehaviorResponseBody) GetStatusValue() *string {
	return s.StatusValue
}

func (s *DescribeUserBusinessBehaviorResponseBody) SetRequestId(v string) *DescribeUserBusinessBehaviorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserBusinessBehaviorResponseBody) SetStatusValue(v string) *DescribeUserBusinessBehaviorResponseBody {
	s.StatusValue = &v
	return s
}

func (s *DescribeUserBusinessBehaviorResponseBody) Validate() error {
	return dara.Validate(s)
}
