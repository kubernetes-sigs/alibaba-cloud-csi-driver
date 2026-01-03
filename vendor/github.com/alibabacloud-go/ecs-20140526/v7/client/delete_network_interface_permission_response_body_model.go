// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkInterfacePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNetworkInterfacePermissionResponseBody
	GetRequestId() *string
}

type DeleteNetworkInterfacePermissionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkInterfacePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkInterfacePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkInterfacePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNetworkInterfacePermissionResponseBody) SetRequestId(v string) *DeleteNetworkInterfacePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNetworkInterfacePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
