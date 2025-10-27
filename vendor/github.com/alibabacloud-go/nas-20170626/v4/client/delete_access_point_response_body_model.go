// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAccessPointResponseBody
	GetRequestId() *string
}

type DeleteAccessPointResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAccessPointResponseBody) SetRequestId(v string) *DeleteAccessPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccessPointResponseBody) Validate() error {
	return dara.Validate(s)
}
