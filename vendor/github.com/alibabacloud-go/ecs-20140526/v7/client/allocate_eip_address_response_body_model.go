// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateEipAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v string) *AllocateEipAddressResponseBody
	GetAllocationId() *string
	SetEipAddress(v string) *AllocateEipAddressResponseBody
	GetEipAddress() *string
	SetRequestId(v string) *AllocateEipAddressResponseBody
	GetRequestId() *string
}

type AllocateEipAddressResponseBody struct {
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	EipAddress   *string `json:"EipAddress,omitempty" xml:"EipAddress,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateEipAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateEipAddressResponseBody) GetAllocationId() *string {
	return s.AllocationId
}

func (s *AllocateEipAddressResponseBody) GetEipAddress() *string {
	return s.EipAddress
}

func (s *AllocateEipAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateEipAddressResponseBody) SetAllocationId(v string) *AllocateEipAddressResponseBody {
	s.AllocationId = &v
	return s
}

func (s *AllocateEipAddressResponseBody) SetEipAddress(v string) *AllocateEipAddressResponseBody {
	s.EipAddress = &v
	return s
}

func (s *AllocateEipAddressResponseBody) SetRequestId(v string) *AllocateEipAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateEipAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
