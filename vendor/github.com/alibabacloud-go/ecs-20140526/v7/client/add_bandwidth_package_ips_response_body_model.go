// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBandwidthPackageIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddBandwidthPackageIpsResponseBody
	GetRequestId() *string
}

type AddBandwidthPackageIpsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddBandwidthPackageIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddBandwidthPackageIpsResponseBody) GoString() string {
	return s.String()
}

func (s *AddBandwidthPackageIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddBandwidthPackageIpsResponseBody) SetRequestId(v string) *AddBandwidthPackageIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddBandwidthPackageIpsResponseBody) Validate() error {
	return dara.Validate(s)
}
