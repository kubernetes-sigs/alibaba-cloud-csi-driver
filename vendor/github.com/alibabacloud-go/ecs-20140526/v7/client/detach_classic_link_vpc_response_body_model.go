// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachClassicLinkVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachClassicLinkVpcResponseBody
	GetRequestId() *string
}

type DetachClassicLinkVpcResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachClassicLinkVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachClassicLinkVpcResponseBody) GoString() string {
	return s.String()
}

func (s *DetachClassicLinkVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachClassicLinkVpcResponseBody) SetRequestId(v string) *DetachClassicLinkVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachClassicLinkVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
