// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrefixListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPrefixListResponseBody
	GetRequestId() *string
}

type ModifyPrefixListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 38793DB8-A4B2-4AEC-BFD3-111234E9188D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPrefixListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrefixListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPrefixListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPrefixListResponseBody) SetRequestId(v string) *ModifyPrefixListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPrefixListResponseBody) Validate() error {
	return dara.Validate(s)
}
