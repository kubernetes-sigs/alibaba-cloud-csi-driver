// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySecurityGroupRuleResponseBody
	GetRequestId() *string
}

type ModifySecurityGroupRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySecurityGroupRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySecurityGroupRuleResponseBody) SetRequestId(v string) *ModifySecurityGroupRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecurityGroupRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
