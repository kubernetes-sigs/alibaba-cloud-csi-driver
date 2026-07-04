// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticSpaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAgenticSpaceResponseBody
	GetRequestId() *string
}

type DeleteAgenticSpaceResponseBody struct {
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAgenticSpaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgenticSpaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgenticSpaceResponseBody) SetRequestId(v string) *DeleteAgenticSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgenticSpaceResponseBody) Validate() error {
	return dara.Validate(s)
}
