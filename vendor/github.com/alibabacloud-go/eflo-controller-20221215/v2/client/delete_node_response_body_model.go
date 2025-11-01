// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNodeResponseBody
	GetRequestId() *string
}

type DeleteNodeResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8F065DDD-6996-5973-9691-9EC57BD0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNodeResponseBody) SetRequestId(v string) *DeleteNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
