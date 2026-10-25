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
	SetEriQuantity(v int32) *DescribeNodeTypeResponseBody
	GetEriQuantity() *int32
	SetNetworkCardQuantity(v int32) *DescribeNodeTypeResponseBody
	GetNetworkCardQuantity() *int32
	SetNetworkCards(v []*DescribeNodeTypeResponseBodyNetworkCards) *DescribeNodeTypeResponseBody
	GetNetworkCards() []*DescribeNodeTypeResponseBodyNetworkCards
	SetPrimaryEniQueueNumber(v int32) *DescribeNodeTypeResponseBody
	GetPrimaryEniQueueNumber() *int32
	SetQueuePairNumber(v int32) *DescribeNodeTypeResponseBody
	GetQueuePairNumber() *int32
	SetRequestId(v string) *DescribeNodeTypeResponseBody
	GetRequestId() *string
	SetSecondaryEniQueueNumber(v int32) *DescribeNodeTypeResponseBody
	GetSecondaryEniQueueNumber() *int32
	SetTotalEniQueueQuantity(v int32) *DescribeNodeTypeResponseBody
	GetTotalEniQueueQuantity() *int32
}

type DescribeNodeTypeResponseBody struct {
	// The maximum number of cloud disks that can be mounted.
	//
	// example:
	//
	// 17
	DiskQuantity *int32 `json:"DiskQuantity,omitempty" xml:"DiskQuantity,omitempty"`
	// The maximum number of network interface controllers (NICs) in high-density mode.
	//
	// example:
	//
	// 63
	EniHighDenseQuantity *int32 `json:"EniHighDenseQuantity,omitempty" xml:"EniHighDenseQuantity,omitempty"`
	// The maximum number of IPv6 addresses supported by a NIC.
	//
	// example:
	//
	// 256
	EniIpv6AddressQuantity *int32 `json:"EniIpv6AddressQuantity,omitempty" xml:"EniIpv6AddressQuantity,omitempty"`
	// The maximum number of IPv4 addresses supported by a NIC.
	//
	// example:
	//
	// 256
	EniPrivateIpAddressQuantity *int32 `json:"EniPrivateIpAddressQuantity,omitempty" xml:"EniPrivateIpAddressQuantity,omitempty"`
	// The maximum number of elastic network interfaces (ENIs) supported, including primary and secondary ENIs.
	//
	// example:
	//
	// 22
	EniQuantity *int32 `json:"EniQuantity,omitempty" xml:"EniQuantity,omitempty"`
	// The number of Elastic RDMA Interfaces (ERIs) for the current specification.
	//
	// example:
	//
	// 24
	EriQuantity *int32 `json:"EriQuantity,omitempty" xml:"EriQuantity,omitempty"`
	// The number of physical network interface controllers (NICs) supported by the instance type.
	//
	// example:
	//
	// 2
	NetworkCardQuantity *int32 `json:"NetworkCardQuantity,omitempty" xml:"NetworkCardQuantity,omitempty"`
	// The list of physical network interface controllers (NICs), including the details of each physical NIC.
	NetworkCards []*DescribeNodeTypeResponseBodyNetworkCards `json:"NetworkCards,omitempty" xml:"NetworkCards,omitempty" type:"Repeated"`
	// The default number of queues for the primary NIC.
	//
	// example:
	//
	// 64
	PrimaryEniQueueNumber *int32 `json:"PrimaryEniQueueNumber,omitempty" xml:"PrimaryEniQueueNumber,omitempty"`
	// The number of queue pairs (QPs) for all Elastic RDMA Interfaces (ERIs) of the current specification.
	//
	// example:
	//
	// 8
	QueuePairNumber *int32 `json:"QueuePairNumber,omitempty" xml:"QueuePairNumber,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The default number of queues for secondary elastic network interfaces (ENIs).
	//
	// example:
	//
	// 32
	SecondaryEniQueueNumber *int32 `json:"SecondaryEniQueueNumber,omitempty" xml:"SecondaryEniQueueNumber,omitempty"`
	// The total queue quota for elastic network interface controllers (NICs) of the current specification.
	//
	// example:
	//
	// 2000
	TotalEniQueueQuantity *int32 `json:"TotalEniQueueQuantity,omitempty" xml:"TotalEniQueueQuantity,omitempty"`
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

func (s *DescribeNodeTypeResponseBody) GetEriQuantity() *int32 {
	return s.EriQuantity
}

func (s *DescribeNodeTypeResponseBody) GetNetworkCardQuantity() *int32 {
	return s.NetworkCardQuantity
}

func (s *DescribeNodeTypeResponseBody) GetNetworkCards() []*DescribeNodeTypeResponseBodyNetworkCards {
	return s.NetworkCards
}

func (s *DescribeNodeTypeResponseBody) GetPrimaryEniQueueNumber() *int32 {
	return s.PrimaryEniQueueNumber
}

func (s *DescribeNodeTypeResponseBody) GetQueuePairNumber() *int32 {
	return s.QueuePairNumber
}

func (s *DescribeNodeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNodeTypeResponseBody) GetSecondaryEniQueueNumber() *int32 {
	return s.SecondaryEniQueueNumber
}

func (s *DescribeNodeTypeResponseBody) GetTotalEniQueueQuantity() *int32 {
	return s.TotalEniQueueQuantity
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

func (s *DescribeNodeTypeResponseBody) SetEriQuantity(v int32) *DescribeNodeTypeResponseBody {
	s.EriQuantity = &v
	return s
}

func (s *DescribeNodeTypeResponseBody) SetNetworkCardQuantity(v int32) *DescribeNodeTypeResponseBody {
	s.NetworkCardQuantity = &v
	return s
}

func (s *DescribeNodeTypeResponseBody) SetNetworkCards(v []*DescribeNodeTypeResponseBodyNetworkCards) *DescribeNodeTypeResponseBody {
	s.NetworkCards = v
	return s
}

func (s *DescribeNodeTypeResponseBody) SetPrimaryEniQueueNumber(v int32) *DescribeNodeTypeResponseBody {
	s.PrimaryEniQueueNumber = &v
	return s
}

func (s *DescribeNodeTypeResponseBody) SetQueuePairNumber(v int32) *DescribeNodeTypeResponseBody {
	s.QueuePairNumber = &v
	return s
}

func (s *DescribeNodeTypeResponseBody) SetRequestId(v string) *DescribeNodeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNodeTypeResponseBody) SetSecondaryEniQueueNumber(v int32) *DescribeNodeTypeResponseBody {
	s.SecondaryEniQueueNumber = &v
	return s
}

func (s *DescribeNodeTypeResponseBody) SetTotalEniQueueQuantity(v int32) *DescribeNodeTypeResponseBody {
	s.TotalEniQueueQuantity = &v
	return s
}

func (s *DescribeNodeTypeResponseBody) Validate() error {
	if s.NetworkCards != nil {
		for _, item := range s.NetworkCards {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNodeTypeResponseBodyNetworkCards struct {
	// The index of the physical network interface controller (NIC), starting from 0. You can specify this index to attach elastic network interfaces (ENIs) to different underlying physical network interface controllers (NICs).
	//
	// example:
	//
	// 0
	NetworkCardIndex *int32 `json:"NetworkCardIndex,omitempty" xml:"NetworkCardIndex,omitempty"`
}

func (s DescribeNodeTypeResponseBodyNetworkCards) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeTypeResponseBodyNetworkCards) GoString() string {
	return s.String()
}

func (s *DescribeNodeTypeResponseBodyNetworkCards) GetNetworkCardIndex() *int32 {
	return s.NetworkCardIndex
}

func (s *DescribeNodeTypeResponseBodyNetworkCards) SetNetworkCardIndex(v int32) *DescribeNodeTypeResponseBodyNetworkCards {
	s.NetworkCardIndex = &v
	return s
}

func (s *DescribeNodeTypeResponseBodyNetworkCards) Validate() error {
	return dara.Validate(s)
}
