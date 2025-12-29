// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveBandwidthPackageIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveBandwidthPackageIpsResponseBody
	GetRequestId() *string
}

type RemoveBandwidthPackageIpsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveBandwidthPackageIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveBandwidthPackageIpsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveBandwidthPackageIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveBandwidthPackageIpsResponseBody) SetRequestId(v string) *RemoveBandwidthPackageIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveBandwidthPackageIpsResponseBody) Validate() error {
	return dara.Validate(s)
}
