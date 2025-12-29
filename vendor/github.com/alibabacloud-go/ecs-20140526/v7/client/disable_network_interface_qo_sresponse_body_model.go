// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableNetworkInterfaceQoSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableNetworkInterfaceQoSResponseBody
	GetRequestId() *string
}

type DisableNetworkInterfaceQoSResponseBody struct {
	// example:
	//
	// 745CEC9F-0DD7-4451-9FE7-8B752F39****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableNetworkInterfaceQoSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableNetworkInterfaceQoSResponseBody) GoString() string {
	return s.String()
}

func (s *DisableNetworkInterfaceQoSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableNetworkInterfaceQoSResponseBody) SetRequestId(v string) *DisableNetworkInterfaceQoSResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableNetworkInterfaceQoSResponseBody) Validate() error {
	return dara.Validate(s)
}
