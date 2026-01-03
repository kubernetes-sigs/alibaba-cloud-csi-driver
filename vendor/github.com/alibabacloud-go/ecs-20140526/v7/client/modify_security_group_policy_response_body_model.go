// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySecurityGroupPolicyResponseBody
	GetRequestId() *string
}

type ModifySecurityGroupPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySecurityGroupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySecurityGroupPolicyResponseBody) SetRequestId(v string) *ModifySecurityGroupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecurityGroupPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
