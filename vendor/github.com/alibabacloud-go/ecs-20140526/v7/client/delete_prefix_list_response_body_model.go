// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrefixListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePrefixListResponseBody
	GetRequestId() *string
}

type DeletePrefixListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 38793DB8-A4B2-4AEC-BFD3-111234E9188D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePrefixListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrefixListResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrefixListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrefixListResponseBody) SetRequestId(v string) *DeletePrefixListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrefixListResponseBody) Validate() error {
	return dara.Validate(s)
}
