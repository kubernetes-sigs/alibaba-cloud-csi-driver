// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceTypes(v *DescribeInstanceTypesResponseBodyInstanceTypes) *DescribeInstanceTypesResponseBody
	GetInstanceTypes() *DescribeInstanceTypesResponseBodyInstanceTypes
	SetNextToken(v string) *DescribeInstanceTypesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeInstanceTypesResponseBody
	GetRequestId() *string
}

type DescribeInstanceTypesResponseBody struct {
	InstanceTypes *DescribeInstanceTypesResponseBodyInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Struct"`
	// The query token returned in this call.
	//
	// example:
	//
	// e71d8a535bd9cc11
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 00827261-20B7-4562-83F2-4DF39876A45A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBody) GetInstanceTypes() *DescribeInstanceTypesResponseBodyInstanceTypes {
	return s.InstanceTypes
}

func (s *DescribeInstanceTypesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceTypesResponseBody) SetInstanceTypes(v *DescribeInstanceTypesResponseBodyInstanceTypes) *DescribeInstanceTypesResponseBody {
	s.InstanceTypes = v
	return s
}

func (s *DescribeInstanceTypesResponseBody) SetNextToken(v string) *DescribeInstanceTypesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceTypesResponseBody) SetRequestId(v string) *DescribeInstanceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceTypesResponseBody) Validate() error {
	if s.InstanceTypes != nil {
		if err := s.InstanceTypes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceTypesResponseBodyInstanceTypes struct {
	InstanceType []*DescribeInstanceTypesResponseBodyInstanceTypesInstanceType `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypes) GetInstanceType() []*DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	return s.InstanceType
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypes) SetInstanceType(v []*DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) *DescribeInstanceTypesResponseBodyInstanceTypes {
	s.InstanceType = v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypes) Validate() error {
	if s.InstanceType != nil {
		for _, item := range s.InstanceType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceTypesResponseBodyInstanceTypesInstanceType struct {
	Attributes                  *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributes         `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Struct"`
	BaselineCredit              *int32                                                                        `json:"BaselineCredit,omitempty" xml:"BaselineCredit,omitempty"`
	Clock                       *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeClock              `json:"Clock,omitempty" xml:"Clock,omitempty" type:"Struct"`
	CpuArchitecture             *string                                                                       `json:"CpuArchitecture,omitempty" xml:"CpuArchitecture,omitempty"`
	CpuCoreCount                *int32                                                                        `json:"CpuCoreCount,omitempty" xml:"CpuCoreCount,omitempty"`
	CpuOptions                  *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions         `json:"CpuOptions,omitempty" xml:"CpuOptions,omitempty" type:"Struct"`
	CpuSpeedFrequency           *float32                                                                      `json:"CpuSpeedFrequency,omitempty" xml:"CpuSpeedFrequency,omitempty"`
	CpuTurboFrequency           *float32                                                                      `json:"CpuTurboFrequency,omitempty" xml:"CpuTurboFrequency,omitempty"`
	DiskQuantity                *int32                                                                        `json:"DiskQuantity,omitempty" xml:"DiskQuantity,omitempty"`
	EnhancedNetwork             *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeEnhancedNetwork    `json:"EnhancedNetwork,omitempty" xml:"EnhancedNetwork,omitempty" type:"Struct"`
	EniIpv6AddressQuantity      *int32                                                                        `json:"EniIpv6AddressQuantity,omitempty" xml:"EniIpv6AddressQuantity,omitempty"`
	EniPrivateIpAddressQuantity *int32                                                                        `json:"EniPrivateIpAddressQuantity,omitempty" xml:"EniPrivateIpAddressQuantity,omitempty"`
	EniQuantity                 *int32                                                                        `json:"EniQuantity,omitempty" xml:"EniQuantity,omitempty"`
	EniTotalQuantity            *int32                                                                        `json:"EniTotalQuantity,omitempty" xml:"EniTotalQuantity,omitempty"`
	EniTrunkSupported           *bool                                                                         `json:"EniTrunkSupported,omitempty" xml:"EniTrunkSupported,omitempty"`
	EriQuantity                 *int32                                                                        `json:"EriQuantity,omitempty" xml:"EriQuantity,omitempty"`
	GPUAmount                   *int32                                                                        `json:"GPUAmount,omitempty" xml:"GPUAmount,omitempty"`
	GPUMemorySize               *float32                                                                      `json:"GPUMemorySize,omitempty" xml:"GPUMemorySize,omitempty"`
	GPUSpec                     *string                                                                       `json:"GPUSpec,omitempty" xml:"GPUSpec,omitempty"`
	InitialCredit               *int32                                                                        `json:"InitialCredit,omitempty" xml:"InitialCredit,omitempty"`
	InstanceBandwidthRx         *int32                                                                        `json:"InstanceBandwidthRx,omitempty" xml:"InstanceBandwidthRx,omitempty"`
	InstanceBandwidthTx         *int32                                                                        `json:"InstanceBandwidthTx,omitempty" xml:"InstanceBandwidthTx,omitempty"`
	InstanceCategory            *string                                                                       `json:"InstanceCategory,omitempty" xml:"InstanceCategory,omitempty"`
	InstanceFamilyLevel         *string                                                                       `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	InstancePpsRx               *int64                                                                        `json:"InstancePpsRx,omitempty" xml:"InstancePpsRx,omitempty"`
	InstancePpsTx               *int64                                                                        `json:"InstancePpsTx,omitempty" xml:"InstancePpsTx,omitempty"`
	InstanceTypeFamily          *string                                                                       `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	InstanceTypeId              *string                                                                       `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty"`
	JumboFrameSupport           *bool                                                                         `json:"JumboFrameSupport,omitempty" xml:"JumboFrameSupport,omitempty"`
	LocalStorageAmount          *int32                                                                        `json:"LocalStorageAmount,omitempty" xml:"LocalStorageAmount,omitempty"`
	LocalStorageCapacity        *int64                                                                        `json:"LocalStorageCapacity,omitempty" xml:"LocalStorageCapacity,omitempty"`
	LocalStorageCategory        *string                                                                       `json:"LocalStorageCategory,omitempty" xml:"LocalStorageCategory,omitempty"`
	MaximumQueueNumberPerEni    *int32                                                                        `json:"MaximumQueueNumberPerEni,omitempty" xml:"MaximumQueueNumberPerEni,omitempty"`
	MemorySize                  *float32                                                                      `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	NetworkCardQuantity         *int32                                                                        `json:"NetworkCardQuantity,omitempty" xml:"NetworkCardQuantity,omitempty"`
	NetworkCards                *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCards       `json:"NetworkCards,omitempty" xml:"NetworkCards,omitempty" type:"Struct"`
	NetworkEncryptionSupport    *bool                                                                         `json:"NetworkEncryptionSupport,omitempty" xml:"NetworkEncryptionSupport,omitempty"`
	NetworkInfo                 *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfo        `json:"NetworkInfo,omitempty" xml:"NetworkInfo,omitempty" type:"Struct"`
	NvmeSupport                 *string                                                                       `json:"NvmeSupport,omitempty" xml:"NvmeSupport,omitempty"`
	PhysicalProcessorModel      *string                                                                       `json:"PhysicalProcessorModel,omitempty" xml:"PhysicalProcessorModel,omitempty"`
	PrimaryEniQueueNumber       *int32                                                                        `json:"PrimaryEniQueueNumber,omitempty" xml:"PrimaryEniQueueNumber,omitempty"`
	QueuePairNumber             *int32                                                                        `json:"QueuePairNumber,omitempty" xml:"QueuePairNumber,omitempty"`
	SecondaryEniQueueNumber     *int32                                                                        `json:"SecondaryEniQueueNumber,omitempty" xml:"SecondaryEniQueueNumber,omitempty"`
	SupportedBootModes          *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeSupportedBootModes `json:"SupportedBootModes,omitempty" xml:"SupportedBootModes,omitempty" type:"Struct"`
	TotalEniQueueQuantity       *int32                                                                        `json:"TotalEniQueueQuantity,omitempty" xml:"TotalEniQueueQuantity,omitempty"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetAttributes() *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributes {
	return s.Attributes
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetBaselineCredit() *int32 {
	return s.BaselineCredit
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetClock() *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeClock {
	return s.Clock
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetCpuArchitecture() *string {
	return s.CpuArchitecture
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetCpuCoreCount() *int32 {
	return s.CpuCoreCount
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetCpuOptions() *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions {
	return s.CpuOptions
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetCpuSpeedFrequency() *float32 {
	return s.CpuSpeedFrequency
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetCpuTurboFrequency() *float32 {
	return s.CpuTurboFrequency
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetDiskQuantity() *int32 {
	return s.DiskQuantity
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetEnhancedNetwork() *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeEnhancedNetwork {
	return s.EnhancedNetwork
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetEniIpv6AddressQuantity() *int32 {
	return s.EniIpv6AddressQuantity
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetEniPrivateIpAddressQuantity() *int32 {
	return s.EniPrivateIpAddressQuantity
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetEniQuantity() *int32 {
	return s.EniQuantity
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetEniTotalQuantity() *int32 {
	return s.EniTotalQuantity
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetEniTrunkSupported() *bool {
	return s.EniTrunkSupported
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetEriQuantity() *int32 {
	return s.EriQuantity
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetGPUAmount() *int32 {
	return s.GPUAmount
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetGPUMemorySize() *float32 {
	return s.GPUMemorySize
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetGPUSpec() *string {
	return s.GPUSpec
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetInitialCredit() *int32 {
	return s.InitialCredit
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetInstanceBandwidthRx() *int32 {
	return s.InstanceBandwidthRx
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetInstanceBandwidthTx() *int32 {
	return s.InstanceBandwidthTx
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetInstanceCategory() *string {
	return s.InstanceCategory
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetInstanceFamilyLevel() *string {
	return s.InstanceFamilyLevel
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetInstancePpsRx() *int64 {
	return s.InstancePpsRx
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetInstancePpsTx() *int64 {
	return s.InstancePpsTx
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetInstanceTypeId() *string {
	return s.InstanceTypeId
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetJumboFrameSupport() *bool {
	return s.JumboFrameSupport
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetLocalStorageAmount() *int32 {
	return s.LocalStorageAmount
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetLocalStorageCapacity() *int64 {
	return s.LocalStorageCapacity
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetLocalStorageCategory() *string {
	return s.LocalStorageCategory
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetMaximumQueueNumberPerEni() *int32 {
	return s.MaximumQueueNumberPerEni
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetMemorySize() *float32 {
	return s.MemorySize
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetNetworkCardQuantity() *int32 {
	return s.NetworkCardQuantity
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetNetworkCards() *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCards {
	return s.NetworkCards
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetNetworkEncryptionSupport() *bool {
	return s.NetworkEncryptionSupport
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetNetworkInfo() *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfo {
	return s.NetworkInfo
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetNvmeSupport() *string {
	return s.NvmeSupport
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetPhysicalProcessorModel() *string {
	return s.PhysicalProcessorModel
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetPrimaryEniQueueNumber() *int32 {
	return s.PrimaryEniQueueNumber
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetQueuePairNumber() *int32 {
	return s.QueuePairNumber
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetSecondaryEniQueueNumber() *int32 {
	return s.SecondaryEniQueueNumber
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetSupportedBootModes() *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeSupportedBootModes {
	return s.SupportedBootModes
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) GetTotalEniQueueQuantity() *int32 {
	return s.TotalEniQueueQuantity
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetAttributes(v *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributes) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.Attributes = v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetBaselineCredit(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.BaselineCredit = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetClock(v *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeClock) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.Clock = v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetCpuArchitecture(v string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.CpuArchitecture = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetCpuCoreCount(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.CpuCoreCount = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetCpuOptions(v *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.CpuOptions = v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetCpuSpeedFrequency(v float32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.CpuSpeedFrequency = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetCpuTurboFrequency(v float32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.CpuTurboFrequency = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetDiskQuantity(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.DiskQuantity = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetEnhancedNetwork(v *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeEnhancedNetwork) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.EnhancedNetwork = v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetEniIpv6AddressQuantity(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.EniIpv6AddressQuantity = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetEniPrivateIpAddressQuantity(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.EniPrivateIpAddressQuantity = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetEniQuantity(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.EniQuantity = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetEniTotalQuantity(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.EniTotalQuantity = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetEniTrunkSupported(v bool) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.EniTrunkSupported = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetEriQuantity(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.EriQuantity = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetGPUAmount(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.GPUAmount = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetGPUMemorySize(v float32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.GPUMemorySize = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetGPUSpec(v string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.GPUSpec = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetInitialCredit(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InitialCredit = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstanceBandwidthRx(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstanceBandwidthRx = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstanceBandwidthTx(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstanceBandwidthTx = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstanceCategory(v string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstanceCategory = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstanceFamilyLevel(v string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstancePpsRx(v int64) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstancePpsRx = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstancePpsTx(v int64) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstancePpsTx = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstanceTypeFamily(v string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstanceTypeId(v string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstanceTypeId = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetJumboFrameSupport(v bool) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.JumboFrameSupport = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetLocalStorageAmount(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.LocalStorageAmount = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetLocalStorageCapacity(v int64) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.LocalStorageCapacity = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetLocalStorageCategory(v string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.LocalStorageCategory = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetMaximumQueueNumberPerEni(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.MaximumQueueNumberPerEni = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetMemorySize(v float32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.MemorySize = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetNetworkCardQuantity(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.NetworkCardQuantity = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetNetworkCards(v *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCards) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.NetworkCards = v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetNetworkEncryptionSupport(v bool) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.NetworkEncryptionSupport = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetNetworkInfo(v *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfo) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.NetworkInfo = v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetNvmeSupport(v string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.NvmeSupport = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetPhysicalProcessorModel(v string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.PhysicalProcessorModel = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetPrimaryEniQueueNumber(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.PrimaryEniQueueNumber = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetQueuePairNumber(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.QueuePairNumber = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetSecondaryEniQueueNumber(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.SecondaryEniQueueNumber = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetSupportedBootModes(v *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeSupportedBootModes) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.SupportedBootModes = v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) SetTotalEniQueueQuantity(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.TotalEniQueueQuantity = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceType) Validate() error {
	if s.Attributes != nil {
		if err := s.Attributes.Validate(); err != nil {
			return err
		}
	}
	if s.Clock != nil {
		if err := s.Clock.Validate(); err != nil {
			return err
		}
	}
	if s.CpuOptions != nil {
		if err := s.CpuOptions.Validate(); err != nil {
			return err
		}
	}
	if s.EnhancedNetwork != nil {
		if err := s.EnhancedNetwork.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkCards != nil {
		if err := s.NetworkCards.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkInfo != nil {
		if err := s.NetworkInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SupportedBootModes != nil {
		if err := s.SupportedBootModes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributes struct {
	Attribute []*DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributesAttribute `json:"Attribute,omitempty" xml:"Attribute,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributes) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributes) GetAttribute() []*DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributesAttribute {
	return s.Attribute
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributes) SetAttribute(v []*DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributesAttribute) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributes {
	s.Attribute = v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributes) Validate() error {
	if s.Attribute != nil {
		for _, item := range s.Attribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributesAttribute struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributesAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributesAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributesAttribute) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributesAttribute) GetValue() *string {
	return s.Value
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributesAttribute) SetName(v string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributesAttribute {
	s.Name = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributesAttribute) SetValue(v string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributesAttribute {
	s.Value = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeAttributesAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeClock struct {
	PtpSupport *string `json:"PtpSupport,omitempty" xml:"PtpSupport,omitempty"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeClock) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeClock) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeClock) GetPtpSupport() *string {
	return s.PtpSupport
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeClock) SetPtpSupport(v string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeClock {
	s.PtpSupport = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeClock) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions struct {
	Core                        *int32                                                                                      `json:"Core,omitempty" xml:"Core,omitempty"`
	CoreFactor                  *int32                                                                                      `json:"CoreFactor,omitempty" xml:"CoreFactor,omitempty"`
	HyperThreadingAdjustable    *bool                                                                                       `json:"HyperThreadingAdjustable,omitempty" xml:"HyperThreadingAdjustable,omitempty"`
	NestedVirtualizationSupport *string                                                                                     `json:"NestedVirtualizationSupport,omitempty" xml:"NestedVirtualizationSupport,omitempty"`
	SupportedTopologyTypes      *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptionsSupportedTopologyTypes `json:"SupportedTopologyTypes,omitempty" xml:"SupportedTopologyTypes,omitempty" type:"Struct"`
	ThreadsPerCore              *int32                                                                                      `json:"ThreadsPerCore,omitempty" xml:"ThreadsPerCore,omitempty"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions) GetCore() *int32 {
	return s.Core
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions) GetCoreFactor() *int32 {
	return s.CoreFactor
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions) GetHyperThreadingAdjustable() *bool {
	return s.HyperThreadingAdjustable
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions) GetNestedVirtualizationSupport() *string {
	return s.NestedVirtualizationSupport
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions) GetSupportedTopologyTypes() *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptionsSupportedTopologyTypes {
	return s.SupportedTopologyTypes
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions) GetThreadsPerCore() *int32 {
	return s.ThreadsPerCore
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions) SetCore(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions {
	s.Core = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions) SetCoreFactor(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions {
	s.CoreFactor = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions) SetHyperThreadingAdjustable(v bool) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions {
	s.HyperThreadingAdjustable = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions) SetNestedVirtualizationSupport(v string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions {
	s.NestedVirtualizationSupport = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions) SetSupportedTopologyTypes(v *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptionsSupportedTopologyTypes) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions {
	s.SupportedTopologyTypes = v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions) SetThreadsPerCore(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions {
	s.ThreadsPerCore = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptions) Validate() error {
	if s.SupportedTopologyTypes != nil {
		if err := s.SupportedTopologyTypes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptionsSupportedTopologyTypes struct {
	SupportedTopologyType []*string `json:"SupportedTopologyType,omitempty" xml:"SupportedTopologyType,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptionsSupportedTopologyTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptionsSupportedTopologyTypes) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptionsSupportedTopologyTypes) GetSupportedTopologyType() []*string {
	return s.SupportedTopologyType
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptionsSupportedTopologyTypes) SetSupportedTopologyType(v []*string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptionsSupportedTopologyTypes {
	s.SupportedTopologyType = v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeCpuOptionsSupportedTopologyTypes) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeEnhancedNetwork struct {
	RssSupport          *bool  `json:"RssSupport,omitempty" xml:"RssSupport,omitempty"`
	SriovSupport        *bool  `json:"SriovSupport,omitempty" xml:"SriovSupport,omitempty"`
	VfQueueNumberPerEni *int32 `json:"VfQueueNumberPerEni,omitempty" xml:"VfQueueNumberPerEni,omitempty"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeEnhancedNetwork) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeEnhancedNetwork) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeEnhancedNetwork) GetRssSupport() *bool {
	return s.RssSupport
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeEnhancedNetwork) GetSriovSupport() *bool {
	return s.SriovSupport
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeEnhancedNetwork) GetVfQueueNumberPerEni() *int32 {
	return s.VfQueueNumberPerEni
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeEnhancedNetwork) SetRssSupport(v bool) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeEnhancedNetwork {
	s.RssSupport = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeEnhancedNetwork) SetSriovSupport(v bool) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeEnhancedNetwork {
	s.SriovSupport = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeEnhancedNetwork) SetVfQueueNumberPerEni(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeEnhancedNetwork {
	s.VfQueueNumberPerEni = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeEnhancedNetwork) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCards struct {
	NetworkCardInfo []*DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCardsNetworkCardInfo `json:"NetworkCardInfo,omitempty" xml:"NetworkCardInfo,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCards) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCards) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCards) GetNetworkCardInfo() []*DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCardsNetworkCardInfo {
	return s.NetworkCardInfo
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCards) SetNetworkCardInfo(v []*DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCardsNetworkCardInfo) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCards {
	s.NetworkCardInfo = v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCards) Validate() error {
	if s.NetworkCardInfo != nil {
		for _, item := range s.NetworkCardInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCardsNetworkCardInfo struct {
	NetworkCardIndex *int32 `json:"NetworkCardIndex,omitempty" xml:"NetworkCardIndex,omitempty"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCardsNetworkCardInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCardsNetworkCardInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCardsNetworkCardInfo) GetNetworkCardIndex() *int32 {
	return s.NetworkCardIndex
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCardsNetworkCardInfo) SetNetworkCardIndex(v int32) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCardsNetworkCardInfo {
	s.NetworkCardIndex = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkCardsNetworkCardInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfo struct {
	BandwidthWeighting *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeighting `json:"BandwidthWeighting,omitempty" xml:"BandwidthWeighting,omitempty" type:"Struct"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfo) GetBandwidthWeighting() *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeighting {
	return s.BandwidthWeighting
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfo) SetBandwidthWeighting(v *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeighting) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfo {
	s.BandwidthWeighting = v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfo) Validate() error {
	if s.BandwidthWeighting != nil {
		if err := s.BandwidthWeighting.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeighting struct {
	WeightingInfos *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfos `json:"WeightingInfos,omitempty" xml:"WeightingInfos,omitempty" type:"Struct"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeighting) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeighting) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeighting) GetWeightingInfos() *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfos {
	return s.WeightingInfos
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeighting) SetWeightingInfos(v *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfos) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeighting {
	s.WeightingInfos = v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeighting) Validate() error {
	if s.WeightingInfos != nil {
		if err := s.WeightingInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfos struct {
	WeightingInfo []*DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo `json:"WeightingInfo,omitempty" xml:"WeightingInfo,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfos) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfos) GetWeightingInfo() []*DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo {
	return s.WeightingInfo
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfos) SetWeightingInfo(v []*DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfos {
	s.WeightingInfo = v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfos) Validate() error {
	if s.WeightingInfo != nil {
		for _, item := range s.WeightingInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo struct {
	EbsBandwidth      *int64  `json:"EbsBandwidth,omitempty" xml:"EbsBandwidth,omitempty"`
	EbsBurstBandwidth *int64  `json:"EbsBurstBandwidth,omitempty" xml:"EbsBurstBandwidth,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	VpcBandwidth      *int64  `json:"VpcBandwidth,omitempty" xml:"VpcBandwidth,omitempty"`
	VpcBurstBandwidth *int64  `json:"VpcBurstBandwidth,omitempty" xml:"VpcBurstBandwidth,omitempty"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo) GetEbsBandwidth() *int64 {
	return s.EbsBandwidth
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo) GetEbsBurstBandwidth() *int64 {
	return s.EbsBurstBandwidth
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo) GetName() *string {
	return s.Name
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo) GetVpcBandwidth() *int64 {
	return s.VpcBandwidth
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo) GetVpcBurstBandwidth() *int64 {
	return s.VpcBurstBandwidth
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo) SetEbsBandwidth(v int64) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo {
	s.EbsBandwidth = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo) SetEbsBurstBandwidth(v int64) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo {
	s.EbsBurstBandwidth = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo) SetName(v string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo {
	s.Name = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo) SetVpcBandwidth(v int64) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo {
	s.VpcBandwidth = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo) SetVpcBurstBandwidth(v int64) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo {
	s.VpcBurstBandwidth = &v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeNetworkInfoBandwidthWeightingWeightingInfosWeightingInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeSupportedBootModes struct {
	SupportedBootMode []*string `json:"SupportedBootMode,omitempty" xml:"SupportedBootMode,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeSupportedBootModes) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeSupportedBootModes) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeSupportedBootModes) GetSupportedBootMode() []*string {
	return s.SupportedBootMode
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeSupportedBootModes) SetSupportedBootMode(v []*string) *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeSupportedBootModes {
	s.SupportedBootMode = v
	return s
}

func (s *DescribeInstanceTypesResponseBodyInstanceTypesInstanceTypeSupportedBootModes) Validate() error {
	return dara.Validate(s)
}
