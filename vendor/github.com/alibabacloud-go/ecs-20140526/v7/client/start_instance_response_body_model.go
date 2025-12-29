// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartInstanceResponseBody
	GetRequestId() *string
}

type StartInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartInstanceResponseBody) SetRequestId(v string) *StartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
