// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVpcAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceVpcAttributeResponseBody
	GetRequestId() *string
}

type ModifyInstanceVpcAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceVpcAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVpcAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVpcAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceVpcAttributeResponseBody) SetRequestId(v string) *ModifyInstanceVpcAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceVpcAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
