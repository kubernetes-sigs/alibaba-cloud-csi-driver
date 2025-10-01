// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAccessRuleResponseBody
	GetRequestId() *string
}

type DeleteAccessRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5B4511A7-C99E-4071-AA8C-32E2529D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAccessRuleResponseBody) SetRequestId(v string) *DeleteAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccessRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
