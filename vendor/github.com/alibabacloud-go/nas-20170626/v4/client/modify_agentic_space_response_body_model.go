// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAgenticSpaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAgenticSpaceResponseBody
	GetRequestId() *string
}

type ModifyAgenticSpaceResponseBody struct {
	// example:
	//
	// BC7C825C-5F65-4B56-BEF6-98C56C7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAgenticSpaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgenticSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAgenticSpaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAgenticSpaceResponseBody) SetRequestId(v string) *ModifyAgenticSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAgenticSpaceResponseBody) Validate() error {
	return dara.Validate(s)
}
