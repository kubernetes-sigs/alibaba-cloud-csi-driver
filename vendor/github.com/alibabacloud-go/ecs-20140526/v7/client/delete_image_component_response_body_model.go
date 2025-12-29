// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageComponentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteImageComponentResponseBody
	GetRequestId() *string
}

type DeleteImageComponentResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImageComponentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageComponentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageComponentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteImageComponentResponseBody) SetRequestId(v string) *DeleteImageComponentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImageComponentResponseBody) Validate() error {
	return dara.Validate(s)
}
