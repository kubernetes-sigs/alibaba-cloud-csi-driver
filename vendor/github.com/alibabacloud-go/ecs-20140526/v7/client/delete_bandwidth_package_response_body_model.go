// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBandwidthPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBandwidthPackageResponseBody
	GetRequestId() *string
}

type DeleteBandwidthPackageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBandwidthPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBandwidthPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBandwidthPackageResponseBody) SetRequestId(v string) *DeleteBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBandwidthPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
