// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiskQuantity(v int32) *DescribeNodeTypeResponseBody
	GetDiskQuantity() *int32
	SetEniHighDenseQuantity(v int32) *DescribeNodeTypeResponseBody
	GetEniHighDenseQuantity() *int32
	SetEniIpv6AddressQuantity(v int32) *DescribeNodeTypeResponseBody
	GetEniIpv6AddressQuantity() *int32
	SetEniPrivateIpAddressQuantity(v int32) *DescribeNodeTypeResponseBody
	GetEniPrivateIpAddressQuantity() *int32
	SetEniQuantity(v int32) *DescribeNodeTypeResponseBody
	GetEniQuantity() *int32
	SetRequestId(v string) *DescribeNodeTypeResponseBody
	GetRequestId() *string
}

type DescribeNodeTypeResponseBody struct {
	DiskQuantity *int32 `json:"DiskQuantity,omitempty" xml:"DiskQuantity,omitempty"`
	// example:
	//
	// 63
	EniHighDenseQuantity *int32 `json:"EniHighDenseQuantity,omitempty" xml:"EniHighDenseQuantity,omitempty"`
	// example:
	//
	// 256
	EniIpv6AddressQuantity *int32 `json:"EniIpv6AddressQuantity,omitempty" xml:"EniIpv6AddressQuantity,omitempty"`
	// example:
	//
	// 256
	EniPrivateIpAddressQuantity *int32 `json:"EniPrivateIpAddressQuantity,omitempty" xml:"EniPrivateIpAddressQuantity,omitempty"`
	// example:
	//
	// 22
	EniQuantity *int32 `json:"EniQuantity,omitempty" xml:"EniQuantity,omitempty"`
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNodeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeTypeResponseBody) GetDiskQuantity() *int32 {
	return s.DiskQuantity
}

func (s *DescribeNodeTypeResponseBody) GetEniHighDenseQuantity() *int32 {
	return s.EniHighDenseQuantity
}

func (s *DescribeNodeTypeResponseBody) GetEniIpv6AddressQuantity() *int32 {
	return s.EniIpv6AddressQuantity
}

func (s *DescribeNodeTypeResponseBody) GetEniPrivateIpAddressQuantity() *int32 {
	return s.EniPrivateIpAddressQuantity
}

func (s *DescribeNodeTypeResponseBody) GetEniQuantity() *int32 {
	return s.EniQuantity
}

func (s *DescribeNodeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNodeTypeResponseBody) SetDiskQuantity(v int32) *DescribeNodeTypeResponseBody {
	s.DiskQuantity = &v
	return s
}

func (s *DescribeNodeTypeResponseBody) SetEniHighDenseQuantity(v int32) *DescribeNodeTypeResponseBody {
	s.EniHighDenseQuantity = &v
	return s
}

func (s *DescribeNodeTypeResponseBody) SetEniIpv6AddressQuantity(v int32) *DescribeNodeTypeResponseBody {
	s.EniIpv6AddressQuantity = &v
	return s
}

func (s *DescribeNodeTypeResponseBody) SetEniPrivateIpAddressQuantity(v int32) *DescribeNodeTypeResponseBody {
	s.EniPrivateIpAddressQuantity = &v
	return s
}

func (s *DescribeNodeTypeResponseBody) SetEniQuantity(v int32) *DescribeNodeTypeResponseBody {
	s.EniQuantity = &v
	return s
}

func (s *DescribeNodeTypeResponseBody) SetRequestId(v string) *DescribeNodeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNodeTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
