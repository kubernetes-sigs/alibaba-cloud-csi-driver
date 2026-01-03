// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStorageSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteStorageSetResponseBody
	GetRequestId() *string
}

type DeleteStorageSetResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStorageSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStorageSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStorageSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStorageSetResponseBody) SetRequestId(v string) *DeleteStorageSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStorageSetResponseBody) Validate() error {
	return dara.Validate(s)
}
