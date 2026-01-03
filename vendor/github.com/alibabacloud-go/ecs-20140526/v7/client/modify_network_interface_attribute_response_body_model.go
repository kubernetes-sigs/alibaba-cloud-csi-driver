// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNetworkInterfaceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyNetworkInterfaceAttributeResponseBody
	GetRequestId() *string
}

type ModifyNetworkInterfaceAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNetworkInterfaceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyNetworkInterfaceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNetworkInterfaceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyNetworkInterfaceAttributeResponseBody) SetRequestId(v string) *ModifyNetworkInterfaceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyNetworkInterfaceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
