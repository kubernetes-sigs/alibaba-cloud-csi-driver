// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSimulatedSystemEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v []*string) *CancelSimulatedSystemEventsRequest
	GetEventId() []*string
	SetOwnerAccount(v string) *CancelSimulatedSystemEventsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CancelSimulatedSystemEventsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CancelSimulatedSystemEventsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CancelSimulatedSystemEventsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelSimulatedSystemEventsRequest
	GetResourceOwnerId() *int64
}

type CancelSimulatedSystemEventsRequest struct {
	// The IDs of simulated system events. You can specify up to 100 event IDs in a single request.
	//
	// This parameter is required.
	//
	// example:
	//
	// e-xhskHun1256****
	EventId      []*string `json:"EventId,omitempty" xml:"EventId,omitempty" type:"Repeated"`
	OwnerAccount *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s CancelSimulatedSystemEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelSimulatedSystemEventsRequest) GoString() string {
	return s.String()
}

func (s *CancelSimulatedSystemEventsRequest) GetEventId() []*string {
	return s.EventId
}

func (s *CancelSimulatedSystemEventsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CancelSimulatedSystemEventsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelSimulatedSystemEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelSimulatedSystemEventsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelSimulatedSystemEventsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelSimulatedSystemEventsRequest) SetEventId(v []*string) *CancelSimulatedSystemEventsRequest {
	s.EventId = v
	return s
}

func (s *CancelSimulatedSystemEventsRequest) SetOwnerAccount(v string) *CancelSimulatedSystemEventsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CancelSimulatedSystemEventsRequest) SetOwnerId(v int64) *CancelSimulatedSystemEventsRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelSimulatedSystemEventsRequest) SetRegionId(v string) *CancelSimulatedSystemEventsRequest {
	s.RegionId = &v
	return s
}

func (s *CancelSimulatedSystemEventsRequest) SetResourceOwnerAccount(v string) *CancelSimulatedSystemEventsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelSimulatedSystemEventsRequest) SetResourceOwnerId(v int64) *CancelSimulatedSystemEventsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelSimulatedSystemEventsRequest) Validate() error {
	return dara.Validate(s)
}
