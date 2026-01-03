// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageSharePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyImageSharePermissionResponseBody
	GetRequestId() *string
}

type ModifyImageSharePermissionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyImageSharePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageSharePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyImageSharePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyImageSharePermissionResponseBody) SetRequestId(v string) *ModifyImageSharePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyImageSharePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
