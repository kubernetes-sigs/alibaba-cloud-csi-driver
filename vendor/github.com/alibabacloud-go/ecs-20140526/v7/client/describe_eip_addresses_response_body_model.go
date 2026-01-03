// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEipAddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEipAddresses(v *DescribeEipAddressesResponseBodyEipAddresses) *DescribeEipAddressesResponseBody
	GetEipAddresses() *DescribeEipAddressesResponseBodyEipAddresses
	SetPageNumber(v int32) *DescribeEipAddressesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEipAddressesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeEipAddressesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeEipAddressesResponseBody
	GetTotalCount() *int32
}

type DescribeEipAddressesResponseBody struct {
	EipAddresses *DescribeEipAddressesResponseBodyEipAddresses `json:"EipAddresses,omitempty" xml:"EipAddresses,omitempty" type:"Struct"`
	PageNumber   *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEipAddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipAddressesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponseBody) GetEipAddresses() *DescribeEipAddressesResponseBodyEipAddresses {
	return s.EipAddresses
}

func (s *DescribeEipAddressesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEipAddressesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEipAddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEipAddressesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEipAddressesResponseBody) SetEipAddresses(v *DescribeEipAddressesResponseBodyEipAddresses) *DescribeEipAddressesResponseBody {
	s.EipAddresses = v
	return s
}

func (s *DescribeEipAddressesResponseBody) SetPageNumber(v int32) *DescribeEipAddressesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEipAddressesResponseBody) SetPageSize(v int32) *DescribeEipAddressesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEipAddressesResponseBody) SetRequestId(v string) *DescribeEipAddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEipAddressesResponseBody) SetTotalCount(v int32) *DescribeEipAddressesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEipAddressesResponseBody) Validate() error {
	if s.EipAddresses != nil {
		if err := s.EipAddresses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEipAddressesResponseBodyEipAddresses struct {
	EipAddress []*DescribeEipAddressesResponseBodyEipAddressesEipAddress `json:"EipAddress,omitempty" xml:"EipAddress,omitempty" type:"Repeated"`
}

func (s DescribeEipAddressesResponseBodyEipAddresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipAddressesResponseBodyEipAddresses) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponseBodyEipAddresses) GetEipAddress() []*DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	return s.EipAddress
}

func (s *DescribeEipAddressesResponseBodyEipAddresses) SetEipAddress(v []*DescribeEipAddressesResponseBodyEipAddressesEipAddress) *DescribeEipAddressesResponseBodyEipAddresses {
	s.EipAddress = v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddresses) Validate() error {
	if s.EipAddress != nil {
		for _, item := range s.EipAddress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEipAddressesResponseBodyEipAddressesEipAddress struct {
	AllocationId       *string                                                               `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	AllocationTime     *string                                                               `json:"AllocationTime,omitempty" xml:"AllocationTime,omitempty"`
	Bandwidth          *string                                                               `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ChargeType         *string                                                               `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	EipBandwidth       *string                                                               `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	ExpiredTime        *string                                                               `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	InstanceId         *string                                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType       *string                                                               `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetChargeType *string                                                               `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	IpAddress          *string                                                               `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	OperationLocks     *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	RegionId           *string                                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status             *string                                                               `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddress) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetAllocationTime() *string {
	return s.AllocationTime
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetEipBandwidth() *string {
	return s.EipBandwidth
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetOperationLocks() *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks {
	return s.OperationLocks
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) GetStatus() *string {
	return s.Status
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetAllocationId(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.AllocationId = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetAllocationTime(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.AllocationTime = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetBandwidth(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.Bandwidth = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetChargeType(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.ChargeType = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetEipBandwidth(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.EipBandwidth = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetExpiredTime(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetInstanceId(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.InstanceId = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetInstanceType(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.InstanceType = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetInternetChargeType(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetIpAddress(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.IpAddress = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetOperationLocks(v *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.OperationLocks = v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetRegionId(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.RegionId = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) SetStatus(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddress {
	s.Status = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddress) Validate() error {
	if s.OperationLocks != nil {
		if err := s.OperationLocks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks struct {
	LockReason []*DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason `json:"LockReason,omitempty" xml:"LockReason,omitempty" type:"Repeated"`
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks) GetLockReason() []*DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason {
	return s.LockReason
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks) SetLockReason(v []*DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason) *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks {
	s.LockReason = v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocks) Validate() error {
	if s.LockReason != nil {
		for _, item := range s.LockReason {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason struct {
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason) GoString() string {
	return s.String()
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason) SetLockReason(v string) *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason {
	s.LockReason = &v
	return s
}

func (s *DescribeEipAddressesResponseBodyEipAddressesEipAddressOperationLocksLockReason) Validate() error {
	return dara.Validate(s)
}
