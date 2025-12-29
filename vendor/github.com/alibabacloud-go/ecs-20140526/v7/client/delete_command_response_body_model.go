// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCommandResponseBody
	GetRequestId() *string
}

type DeleteCommandResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCommandResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCommandResponseBody) SetRequestId(v string) *DeleteCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCommandResponseBody) Validate() error {
	return dara.Validate(s)
}
