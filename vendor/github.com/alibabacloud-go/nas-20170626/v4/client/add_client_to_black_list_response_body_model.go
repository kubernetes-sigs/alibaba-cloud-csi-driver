// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddClientToBlackListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddClientToBlackListResponseBody
	GetRequestId() *string
}

type AddClientToBlackListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A70BEE5D-76D3-49FB-B58F-1F398211A5C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddClientToBlackListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddClientToBlackListResponseBody) GoString() string {
	return s.String()
}

func (s *AddClientToBlackListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddClientToBlackListResponseBody) SetRequestId(v string) *AddClientToBlackListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddClientToBlackListResponseBody) Validate() error {
	return dara.Validate(s)
}
