// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachNetworkInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachNetworkInterfaceResponseBody
	GetRequestId() *string
}

type AttachNetworkInterfaceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachNetworkInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *AttachNetworkInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachNetworkInterfaceResponseBody) SetRequestId(v string) *AttachNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachNetworkInterfaceResponseBody) Validate() error {
	return dara.Validate(s)
}
