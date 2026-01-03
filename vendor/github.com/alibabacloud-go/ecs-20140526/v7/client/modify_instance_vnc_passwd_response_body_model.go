// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVncPasswdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceVncPasswdResponseBody
	GetRequestId() *string
}

type ModifyInstanceVncPasswdResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceVncPasswdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVncPasswdResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVncPasswdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceVncPasswdResponseBody) SetRequestId(v string) *ModifyInstanceVncPasswdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceVncPasswdResponseBody) Validate() error {
	return dara.Validate(s)
}
