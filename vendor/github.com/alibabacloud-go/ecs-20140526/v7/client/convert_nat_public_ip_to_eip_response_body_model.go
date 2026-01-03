// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertNatPublicIpToEipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConvertNatPublicIpToEipResponseBody
	GetRequestId() *string
}

type ConvertNatPublicIpToEipResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConvertNatPublicIpToEipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConvertNatPublicIpToEipResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertNatPublicIpToEipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConvertNatPublicIpToEipResponseBody) SetRequestId(v string) *ConvertNatPublicIpToEipResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertNatPublicIpToEipResponseBody) Validate() error {
	return dara.Validate(s)
}
