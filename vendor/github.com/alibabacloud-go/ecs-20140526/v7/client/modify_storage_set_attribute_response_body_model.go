// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStorageSetAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyStorageSetAttributeResponseBody
	GetRequestId() *string
}

type ModifyStorageSetAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73369
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyStorageSetAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyStorageSetAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyStorageSetAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyStorageSetAttributeResponseBody) SetRequestId(v string) *ModifyStorageSetAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyStorageSetAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
