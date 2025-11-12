// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVscResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVscResponseBody
	GetRequestId() *string
}

type DeleteVscResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVscResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVscResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVscResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVscResponseBody) SetRequestId(v string) *DeleteVscResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVscResponseBody) Validate() error {
	return dara.Validate(s)
}
