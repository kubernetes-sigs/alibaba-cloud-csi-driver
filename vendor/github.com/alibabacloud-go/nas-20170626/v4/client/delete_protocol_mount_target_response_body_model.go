// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProtocolMountTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteProtocolMountTargetResponseBody
	GetRequestId() *string
}

type DeleteProtocolMountTargetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProtocolMountTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProtocolMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProtocolMountTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProtocolMountTargetResponseBody) SetRequestId(v string) *DeleteProtocolMountTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProtocolMountTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
