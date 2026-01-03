// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrefixListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrefixListId(v string) *CreatePrefixListResponseBody
	GetPrefixListId() *string
	SetRequestId(v string) *CreatePrefixListResponseBody
	GetRequestId() *string
}

type CreatePrefixListResponseBody struct {
	// The ID of the prefix list.
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 38793DB8-A4B2-4AEC-BFD3-111234E9188D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePrefixListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePrefixListResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrefixListResponseBody) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *CreatePrefixListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePrefixListResponseBody) SetPrefixListId(v string) *CreatePrefixListResponseBody {
	s.PrefixListId = &v
	return s
}

func (s *CreatePrefixListResponseBody) SetRequestId(v string) *CreatePrefixListResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrefixListResponseBody) Validate() error {
	return dara.Validate(s)
}
