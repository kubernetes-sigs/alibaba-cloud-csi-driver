// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachClassicLinkVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachClassicLinkVpcResponseBody
	GetRequestId() *string
}

type AttachClassicLinkVpcResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachClassicLinkVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachClassicLinkVpcResponseBody) GoString() string {
	return s.String()
}

func (s *AttachClassicLinkVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachClassicLinkVpcResponseBody) SetRequestId(v string) *AttachClassicLinkVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachClassicLinkVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
