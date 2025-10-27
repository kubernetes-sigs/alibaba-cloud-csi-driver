// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccessRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAccessRuleResponseBody
	GetRequestId() *string
}

type ModifyAccessRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6299428C-3861-435D-AE54-9B330A00****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccessRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccessRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccessRuleResponseBody) SetRequestId(v string) *ModifyAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccessRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
