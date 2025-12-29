// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyForwardEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExternalIp(v string) *ModifyForwardEntryRequest
	GetExternalIp() *string
	SetExternalPort(v string) *ModifyForwardEntryRequest
	GetExternalPort() *string
	SetForwardEntryId(v string) *ModifyForwardEntryRequest
	GetForwardEntryId() *string
	SetForwardTableId(v string) *ModifyForwardEntryRequest
	GetForwardTableId() *string
	SetInternalIp(v string) *ModifyForwardEntryRequest
	GetInternalIp() *string
	SetInternalPort(v string) *ModifyForwardEntryRequest
	GetInternalPort() *string
	SetIpProtocol(v string) *ModifyForwardEntryRequest
	GetIpProtocol() *string
	SetOwnerAccount(v string) *ModifyForwardEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyForwardEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyForwardEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyForwardEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyForwardEntryRequest
	GetResourceOwnerId() *int64
}

type ModifyForwardEntryRequest struct {
	ExternalIp   *string `json:"ExternalIp,omitempty" xml:"ExternalIp,omitempty"`
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// This parameter is required.
	ForwardEntryId *string `json:"ForwardEntryId,omitempty" xml:"ForwardEntryId,omitempty"`
	// This parameter is required.
	ForwardTableId *string `json:"ForwardTableId,omitempty" xml:"ForwardTableId,omitempty"`
	InternalIp     *string `json:"InternalIp,omitempty" xml:"InternalIp,omitempty"`
	InternalPort   *string `json:"InternalPort,omitempty" xml:"InternalPort,omitempty"`
	IpProtocol     *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyForwardEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyForwardEntryRequest) GoString() string {
	return s.String()
}

func (s *ModifyForwardEntryRequest) GetExternalIp() *string {
	return s.ExternalIp
}

func (s *ModifyForwardEntryRequest) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *ModifyForwardEntryRequest) GetForwardEntryId() *string {
	return s.ForwardEntryId
}

func (s *ModifyForwardEntryRequest) GetForwardTableId() *string {
	return s.ForwardTableId
}

func (s *ModifyForwardEntryRequest) GetInternalIp() *string {
	return s.InternalIp
}

func (s *ModifyForwardEntryRequest) GetInternalPort() *string {
	return s.InternalPort
}

func (s *ModifyForwardEntryRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifyForwardEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyForwardEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyForwardEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyForwardEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyForwardEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyForwardEntryRequest) SetExternalIp(v string) *ModifyForwardEntryRequest {
	s.ExternalIp = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetExternalPort(v string) *ModifyForwardEntryRequest {
	s.ExternalPort = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetForwardEntryId(v string) *ModifyForwardEntryRequest {
	s.ForwardEntryId = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetForwardTableId(v string) *ModifyForwardEntryRequest {
	s.ForwardTableId = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetInternalIp(v string) *ModifyForwardEntryRequest {
	s.InternalIp = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetInternalPort(v string) *ModifyForwardEntryRequest {
	s.InternalPort = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetIpProtocol(v string) *ModifyForwardEntryRequest {
	s.IpProtocol = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetOwnerAccount(v string) *ModifyForwardEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetOwnerId(v int64) *ModifyForwardEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetRegionId(v string) *ModifyForwardEntryRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetResourceOwnerAccount(v string) *ModifyForwardEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyForwardEntryRequest) SetResourceOwnerId(v int64) *ModifyForwardEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyForwardEntryRequest) Validate() error {
	return dara.Validate(s)
}
