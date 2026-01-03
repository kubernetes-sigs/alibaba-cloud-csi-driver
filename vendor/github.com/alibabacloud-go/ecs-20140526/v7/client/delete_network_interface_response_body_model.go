// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkInterfaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNetworkInterfaceResponseBody
	GetRequestId() *string
}

type DeleteNetworkInterfaceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F3CD6886-D8D0-4FEE-B93E-1B73239673DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkInterfaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkInterfaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkInterfaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNetworkInterfaceResponseBody) SetRequestId(v string) *DeleteNetworkInterfaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNetworkInterfaceResponseBody) Validate() error {
	return dara.Validate(s)
}
