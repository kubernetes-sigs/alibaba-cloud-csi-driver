// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecycleBinAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRecycleBinAttributeResponseBody
	GetRequestId() *string
}

type UpdateRecycleBinAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BC7C825C-5F65-4B56-BEF6-98C5****546E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecycleBinAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecycleBinAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecycleBinAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRecycleBinAttributeResponseBody) SetRequestId(v string) *UpdateRecycleBinAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecycleBinAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
