// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAutoReleaseTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceAutoReleaseTimeResponseBody
	GetRequestId() *string
}

type ModifyInstanceAutoReleaseTimeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceAutoReleaseTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAutoReleaseTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAutoReleaseTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceAutoReleaseTimeResponseBody) SetRequestId(v string) *ModifyInstanceAutoReleaseTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceAutoReleaseTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
