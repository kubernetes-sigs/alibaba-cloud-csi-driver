// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReInitDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReInitDiskResponseBody
	GetRequestId() *string
}

type ReInitDiskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReInitDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReInitDiskResponseBody) GoString() string {
	return s.String()
}

func (s *ReInitDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReInitDiskResponseBody) SetRequestId(v string) *ReInitDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReInitDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
