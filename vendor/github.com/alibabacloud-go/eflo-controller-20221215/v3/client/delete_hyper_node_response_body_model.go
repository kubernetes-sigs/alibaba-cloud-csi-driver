// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHyperNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHyperNodeResponseBody
	GetRequestId() *string
}

type DeleteHyperNodeResponseBody struct {
	// example:
	//
	// 041724FC-2BD7-58B1-863B-B42022D4B351
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHyperNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHyperNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHyperNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHyperNodeResponseBody) SetRequestId(v string) *DeleteHyperNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHyperNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
