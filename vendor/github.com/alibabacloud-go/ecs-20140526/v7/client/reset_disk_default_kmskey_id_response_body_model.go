// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDiskDefaultKMSKeyIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetDiskDefaultKMSKeyIdResponseBody
	GetRequestId() *string
}

type ResetDiskDefaultKMSKeyIdResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetDiskDefaultKMSKeyIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetDiskDefaultKMSKeyIdResponseBody) GoString() string {
	return s.String()
}

func (s *ResetDiskDefaultKMSKeyIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetDiskDefaultKMSKeyIdResponseBody) SetRequestId(v string) *ResetDiskDefaultKMSKeyIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetDiskDefaultKMSKeyIdResponseBody) Validate() error {
	return dara.Validate(s)
}
