// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachKeyPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *DetachKeyPairRequest
	GetInstanceIds() *string
	SetKeyPairName(v string) *DetachKeyPairRequest
	GetKeyPairName() *string
	SetOwnerId(v int64) *DetachKeyPairRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DetachKeyPairRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DetachKeyPairRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DetachKeyPairRequest
	GetResourceOwnerId() *int64
}

type DetachKeyPairRequest struct {
	// The IDs of instances from which you want to unbind the SSH key pair. The value can be a JSON array that consists of up to 50 instance IDs. Separate multiple instance IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["i-bp1d6tsvznfghy7y****", "i-bp1ippxbaql9zet7****", … "i-bp1ib7bcz07l****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The name of the SSH key pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the SSH key pair. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DetachKeyPairRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachKeyPairRequest) GoString() string {
	return s.String()
}

func (s *DetachKeyPairRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DetachKeyPairRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DetachKeyPairRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DetachKeyPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachKeyPairRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DetachKeyPairRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DetachKeyPairRequest) SetInstanceIds(v string) *DetachKeyPairRequest {
	s.InstanceIds = &v
	return s
}

func (s *DetachKeyPairRequest) SetKeyPairName(v string) *DetachKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *DetachKeyPairRequest) SetOwnerId(v int64) *DetachKeyPairRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachKeyPairRequest) SetRegionId(v string) *DetachKeyPairRequest {
	s.RegionId = &v
	return s
}

func (s *DetachKeyPairRequest) SetResourceOwnerAccount(v string) *DetachKeyPairRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachKeyPairRequest) SetResourceOwnerId(v int64) *DetachKeyPairRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DetachKeyPairRequest) Validate() error {
	return dara.Validate(s)
}
