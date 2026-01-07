// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSpotPriceHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeSpotPriceHistoryRequest
	GetEndTime() *string
	SetInstanceType(v string) *DescribeSpotPriceHistoryRequest
	GetInstanceType() *string
	SetIoOptimized(v string) *DescribeSpotPriceHistoryRequest
	GetIoOptimized() *string
	SetNetworkType(v string) *DescribeSpotPriceHistoryRequest
	GetNetworkType() *string
	SetOSType(v string) *DescribeSpotPriceHistoryRequest
	GetOSType() *string
	SetOffset(v int32) *DescribeSpotPriceHistoryRequest
	GetOffset() *int32
	SetOwnerAccount(v string) *DescribeSpotPriceHistoryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSpotPriceHistoryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSpotPriceHistoryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSpotPriceHistoryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSpotPriceHistoryRequest
	GetResourceOwnerId() *int64
	SetSpotDuration(v int32) *DescribeSpotPriceHistoryRequest
	GetSpotDuration() *int32
	SetStartTime(v string) *DescribeSpotPriceHistoryRequest
	GetStartTime() *string
	SetZoneId(v string) *DescribeSpotPriceHistoryRequest
	GetZoneId() *string
}

type DescribeSpotPriceHistoryRequest struct {
	// The end of the time range to query. Specify the time in the [ISO 8601 standard](https://help.aliyun.com/document_detail/25696.html) in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// This parameter is empty by default. If this parameter is empty, the current time is used.
	//
	// example:
	//
	// 2017-08-22T08:45:08Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC. The specified time can be up to 30 days earlier than the specified EndTime value.
	//
	// This parameter is empty by default. If this parameter is empty, the time that is 3 hours earlier than the specified EndTime value is used.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.t1.xsmall
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies whether the instance is I/O optimized. Valid values:
	//
	// 	- optimized: The instance is I/O optimized.
	//
	// 	- none: The instance is not I/O optimized.
	//
	// For instances of generation I instance families, the default value is none.
	//
	// For instances of other instance families, the default value is optimized.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The network type of the spot instance. Valid values:
	//
	// 	- classic: classic network
	//
	// 	- vpc: Virtual Private Cloud (VPC)
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The type of the operating system platform. Valid values:
	//
	// - linux
	//
	// - windows
	//
	// example:
	//
	// linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The line from which the query starts.
	//
	// Default value: 0
	//
	// example:
	//
	// 0
	Offset       *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The zone ID of the spot instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The protection period of the spot instance. Unit: hours. Default value: 1. Valid values:
	//
	// 	- 1: After a spot instance is created, Alibaba Cloud ensures that the instance is not automatically released within 1 hour. After the 1-hour protection period ends, the system compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// 	- 0: After a spot instance is created, Alibaba Cloud does not ensure that the instance runs for 1 hour. The system compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// Alibaba Cloud sends an ECS system event to notify you 5 minutes before the instance is released. Spot instances are billed by second. We recommend that you specify a protection period based on your business requirements.
	//
	// >  This parameter takes effect only if you set SpotStrategy to SpotWithPriceLimit or SpotAsPriceGo.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The beginning of the time range to query. The value of this parameter and the value of EndTime can be up to 30 days apart. Specify the time in the [ISO 8601 standard](https://help.aliyun.com/document_detail/25696.html) in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// This parameter is left empty by default. If this parameter is empty, the time that is 3 hours earlier than the value of EndTime is used.
	//
	// example:
	//
	// 2017-08-22T08:45:08Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The spot price (market price) of the spot instance.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeSpotPriceHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpotPriceHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeSpotPriceHistoryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSpotPriceHistoryRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeSpotPriceHistoryRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *DescribeSpotPriceHistoryRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeSpotPriceHistoryRequest) GetOSType() *string {
	return s.OSType
}

func (s *DescribeSpotPriceHistoryRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *DescribeSpotPriceHistoryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSpotPriceHistoryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSpotPriceHistoryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSpotPriceHistoryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSpotPriceHistoryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSpotPriceHistoryRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *DescribeSpotPriceHistoryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSpotPriceHistoryRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeSpotPriceHistoryRequest) SetEndTime(v string) *DescribeSpotPriceHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetInstanceType(v string) *DescribeSpotPriceHistoryRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetIoOptimized(v string) *DescribeSpotPriceHistoryRequest {
	s.IoOptimized = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetNetworkType(v string) *DescribeSpotPriceHistoryRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetOSType(v string) *DescribeSpotPriceHistoryRequest {
	s.OSType = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetOffset(v int32) *DescribeSpotPriceHistoryRequest {
	s.Offset = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetOwnerAccount(v string) *DescribeSpotPriceHistoryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetOwnerId(v int64) *DescribeSpotPriceHistoryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetRegionId(v string) *DescribeSpotPriceHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetResourceOwnerAccount(v string) *DescribeSpotPriceHistoryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetResourceOwnerId(v int64) *DescribeSpotPriceHistoryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetSpotDuration(v int32) *DescribeSpotPriceHistoryRequest {
	s.SpotDuration = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetStartTime(v string) *DescribeSpotPriceHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) SetZoneId(v string) *DescribeSpotPriceHistoryRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeSpotPriceHistoryRequest) Validate() error {
	return dara.Validate(s)
}
