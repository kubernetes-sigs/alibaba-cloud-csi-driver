// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBandwidthPackageSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyBandwidthPackageSpecResponseBody
	GetRequestId() *string
}

type ModifyBandwidthPackageSpecResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBandwidthPackageSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBandwidthPackageSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBandwidthPackageSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBandwidthPackageSpecResponseBody) SetRequestId(v string) *ModifyBandwidthPackageSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBandwidthPackageSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
