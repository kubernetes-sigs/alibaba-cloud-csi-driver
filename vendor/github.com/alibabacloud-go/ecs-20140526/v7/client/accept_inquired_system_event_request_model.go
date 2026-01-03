// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptInquiredSystemEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChoice(v string) *AcceptInquiredSystemEventRequest
	GetChoice() *string
	SetEventId(v string) *AcceptInquiredSystemEventRequest
	GetEventId() *string
	SetOwnerAccount(v string) *AcceptInquiredSystemEventRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AcceptInquiredSystemEventRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AcceptInquiredSystemEventRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AcceptInquiredSystemEventRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AcceptInquiredSystemEventRequest
	GetResourceOwnerId() *int64
}

type AcceptInquiredSystemEventRequest struct {
	// example:
	//
	// hide
	Choice *string `json:"Choice,omitempty" xml:"Choice,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e-2zeielxl1qzq8slb****
	EventId      *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AcceptInquiredSystemEventRequest) String() string {
	return dara.Prettify(s)
}

func (s AcceptInquiredSystemEventRequest) GoString() string {
	return s.String()
}

func (s *AcceptInquiredSystemEventRequest) GetChoice() *string {
	return s.Choice
}

func (s *AcceptInquiredSystemEventRequest) GetEventId() *string {
	return s.EventId
}

func (s *AcceptInquiredSystemEventRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AcceptInquiredSystemEventRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AcceptInquiredSystemEventRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AcceptInquiredSystemEventRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AcceptInquiredSystemEventRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AcceptInquiredSystemEventRequest) SetChoice(v string) *AcceptInquiredSystemEventRequest {
	s.Choice = &v
	return s
}

func (s *AcceptInquiredSystemEventRequest) SetEventId(v string) *AcceptInquiredSystemEventRequest {
	s.EventId = &v
	return s
}

func (s *AcceptInquiredSystemEventRequest) SetOwnerAccount(v string) *AcceptInquiredSystemEventRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AcceptInquiredSystemEventRequest) SetOwnerId(v int64) *AcceptInquiredSystemEventRequest {
	s.OwnerId = &v
	return s
}

func (s *AcceptInquiredSystemEventRequest) SetRegionId(v string) *AcceptInquiredSystemEventRequest {
	s.RegionId = &v
	return s
}

func (s *AcceptInquiredSystemEventRequest) SetResourceOwnerAccount(v string) *AcceptInquiredSystemEventRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AcceptInquiredSystemEventRequest) SetResourceOwnerId(v int64) *AcceptInquiredSystemEventRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AcceptInquiredSystemEventRequest) Validate() error {
	return dara.Validate(s)
}
