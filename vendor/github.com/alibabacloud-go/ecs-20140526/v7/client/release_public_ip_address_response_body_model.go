// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePublicIpAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRemainTimes(v string) *ReleasePublicIpAddressResponseBody
	GetRemainTimes() *string
	SetRequestId(v string) *ReleasePublicIpAddressResponseBody
	GetRequestId() *string
}

type ReleasePublicIpAddressResponseBody struct {
	// > This parameter is unavailable.
	//
	// example:
	//
	// hide
	RemainTimes *string `json:"RemainTimes,omitempty" xml:"RemainTimes,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleasePublicIpAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleasePublicIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ReleasePublicIpAddressResponseBody) GetRemainTimes() *string {
	return s.RemainTimes
}

func (s *ReleasePublicIpAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleasePublicIpAddressResponseBody) SetRemainTimes(v string) *ReleasePublicIpAddressResponseBody {
	s.RemainTimes = &v
	return s
}

func (s *ReleasePublicIpAddressResponseBody) SetRequestId(v string) *ReleasePublicIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleasePublicIpAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
