// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachNetworkInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachNetworkInterfaceResponseBody
	GetRequestId() *string
}

type DetachNetworkInterfaceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachNetworkInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *DetachNetworkInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachNetworkInterfaceResponseBody) SetRequestId(v string) *DetachNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachNetworkInterfaceResponseBody) Validate() error {
	return dara.Validate(s)
}
