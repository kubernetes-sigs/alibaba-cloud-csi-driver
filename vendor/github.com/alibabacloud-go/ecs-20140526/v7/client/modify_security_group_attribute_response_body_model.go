// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySecurityGroupAttributeResponseBody
	GetRequestId() *string
}

type ModifySecurityGroupAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySecurityGroupAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySecurityGroupAttributeResponseBody) SetRequestId(v string) *ModifySecurityGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecurityGroupAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
