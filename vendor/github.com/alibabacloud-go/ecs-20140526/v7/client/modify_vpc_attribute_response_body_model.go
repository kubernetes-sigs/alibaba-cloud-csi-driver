// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVpcAttributeResponseBody
	GetRequestId() *string
}

type ModifyVpcAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpcAttributeResponseBody) SetRequestId(v string) *ModifyVpcAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpcAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
