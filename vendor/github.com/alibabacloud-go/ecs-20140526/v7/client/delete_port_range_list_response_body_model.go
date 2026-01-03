// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePortRangeListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePortRangeListResponseBody
	GetRequestId() *string
}

type DeletePortRangeListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePortRangeListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePortRangeListResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePortRangeListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePortRangeListResponseBody) SetRequestId(v string) *DeletePortRangeListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePortRangeListResponseBody) Validate() error {
	return dara.Validate(s)
}
