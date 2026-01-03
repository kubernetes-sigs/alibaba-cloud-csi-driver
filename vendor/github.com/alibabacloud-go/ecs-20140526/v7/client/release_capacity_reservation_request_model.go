// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseCapacityReservationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivatePoolOptions(v *ReleaseCapacityReservationRequestPrivatePoolOptions) *ReleaseCapacityReservationRequest
	GetPrivatePoolOptions() *ReleaseCapacityReservationRequestPrivatePoolOptions
	SetDryRun(v bool) *ReleaseCapacityReservationRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *ReleaseCapacityReservationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReleaseCapacityReservationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ReleaseCapacityReservationRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ReleaseCapacityReservationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReleaseCapacityReservationRequest
	GetResourceOwnerId() *int64
}

type ReleaseCapacityReservationRequest struct {
	PrivatePoolOptions *ReleaseCapacityReservationRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	// Specifies whether to perform only a dry run, without performing the actual request. Set the value to false. This indicates that the system directly releases the capacity reservation.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the capacity reservation. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ReleaseCapacityReservationRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseCapacityReservationRequest) GoString() string {
	return s.String()
}

func (s *ReleaseCapacityReservationRequest) GetPrivatePoolOptions() *ReleaseCapacityReservationRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *ReleaseCapacityReservationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ReleaseCapacityReservationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReleaseCapacityReservationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReleaseCapacityReservationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReleaseCapacityReservationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReleaseCapacityReservationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReleaseCapacityReservationRequest) SetPrivatePoolOptions(v *ReleaseCapacityReservationRequestPrivatePoolOptions) *ReleaseCapacityReservationRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *ReleaseCapacityReservationRequest) SetDryRun(v bool) *ReleaseCapacityReservationRequest {
	s.DryRun = &v
	return s
}

func (s *ReleaseCapacityReservationRequest) SetOwnerAccount(v string) *ReleaseCapacityReservationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseCapacityReservationRequest) SetOwnerId(v int64) *ReleaseCapacityReservationRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseCapacityReservationRequest) SetRegionId(v string) *ReleaseCapacityReservationRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseCapacityReservationRequest) SetResourceOwnerAccount(v string) *ReleaseCapacityReservationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseCapacityReservationRequest) SetResourceOwnerId(v int64) *ReleaseCapacityReservationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseCapacityReservationRequest) Validate() error {
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReleaseCapacityReservationRequestPrivatePoolOptions struct {
	// The ID of the capacity reservation.
	//
	// This parameter is required.
	//
	// example:
	//
	// crp-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ReleaseCapacityReservationRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s ReleaseCapacityReservationRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *ReleaseCapacityReservationRequestPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *ReleaseCapacityReservationRequestPrivatePoolOptions) SetId(v string) *ReleaseCapacityReservationRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *ReleaseCapacityReservationRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}
