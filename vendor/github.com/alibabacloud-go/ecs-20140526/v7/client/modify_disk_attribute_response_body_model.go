// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDiskAttributeResponseBody
	GetRequestId() *string
}

type ModifyDiskAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDiskAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDiskAttributeResponseBody) SetRequestId(v string) *ModifyDiskAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDiskAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
