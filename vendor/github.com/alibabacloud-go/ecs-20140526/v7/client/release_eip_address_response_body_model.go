// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseEipAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseEipAddressResponseBody
	GetRequestId() *string
}

type ReleaseEipAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseEipAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseEipAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseEipAddressResponseBody) SetRequestId(v string) *ReleaseEipAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseEipAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
