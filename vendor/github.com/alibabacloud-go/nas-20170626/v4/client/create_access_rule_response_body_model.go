// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessRuleId(v string) *CreateAccessRuleResponseBody
	GetAccessRuleId() *string
	SetRequestId(v string) *CreateAccessRuleResponseBody
	GetRequestId() *string
}

type CreateAccessRuleResponseBody struct {
	// The rule ID.
	//
	// example:
	//
	// 1
	AccessRuleId *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A323836B-5BC6-45A6-8048-60675C23****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessRuleResponseBody) GetAccessRuleId() *string {
	return s.AccessRuleId
}

func (s *CreateAccessRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccessRuleResponseBody) SetAccessRuleId(v string) *CreateAccessRuleResponseBody {
	s.AccessRuleId = &v
	return s
}

func (s *CreateAccessRuleResponseBody) SetRequestId(v string) *CreateAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccessRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
