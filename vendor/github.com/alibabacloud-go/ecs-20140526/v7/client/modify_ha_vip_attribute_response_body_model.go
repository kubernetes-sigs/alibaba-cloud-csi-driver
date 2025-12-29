// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHaVipAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHaVipAttributeResponseBody
	GetRequestId() *string
}

type ModifyHaVipAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHaVipAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHaVipAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHaVipAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHaVipAttributeResponseBody) SetRequestId(v string) *ModifyHaVipAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHaVipAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
