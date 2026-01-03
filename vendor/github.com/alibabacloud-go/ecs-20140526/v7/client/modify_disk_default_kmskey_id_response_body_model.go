// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskDefaultKMSKeyIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDiskDefaultKMSKeyIdResponseBody
	GetRequestId() *string
}

type ModifyDiskDefaultKMSKeyIdResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDiskDefaultKMSKeyIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskDefaultKMSKeyIdResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskDefaultKMSKeyIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDiskDefaultKMSKeyIdResponseBody) SetRequestId(v string) *ModifyDiskDefaultKMSKeyIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDiskDefaultKMSKeyIdResponseBody) Validate() error {
	return dara.Validate(s)
}
