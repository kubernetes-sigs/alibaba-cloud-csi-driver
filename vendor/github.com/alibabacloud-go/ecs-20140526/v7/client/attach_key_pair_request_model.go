// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachKeyPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *AttachKeyPairRequest
	GetInstanceIds() *string
	SetKeyPairName(v string) *AttachKeyPairRequest
	GetKeyPairName() *string
	SetOwnerId(v int64) *AttachKeyPairRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AttachKeyPairRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AttachKeyPairRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AttachKeyPairRequest
	GetResourceOwnerId() *int64
}

type AttachKeyPairRequest struct {
	// The IDs of instances to which you want to bind the SSH key pair. The value can be a JSON array that consists of up to 50 instance IDs. Separate multiple instance IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["i-bp1gtjxuuvwj17zr****", "i-bp17b7zrsbjwvmfy****", … "i-bp1h6jmbefj1ytos****"]
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

func (s AttachKeyPairRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachKeyPairRequest) GoString() string {
	return s.String()
}

func (s *AttachKeyPairRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *AttachKeyPairRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *AttachKeyPairRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AttachKeyPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachKeyPairRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AttachKeyPairRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AttachKeyPairRequest) SetInstanceIds(v string) *AttachKeyPairRequest {
	s.InstanceIds = &v
	return s
}

func (s *AttachKeyPairRequest) SetKeyPairName(v string) *AttachKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *AttachKeyPairRequest) SetOwnerId(v int64) *AttachKeyPairRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachKeyPairRequest) SetRegionId(v string) *AttachKeyPairRequest {
	s.RegionId = &v
	return s
}

func (s *AttachKeyPairRequest) SetResourceOwnerAccount(v string) *AttachKeyPairRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachKeyPairRequest) SetResourceOwnerId(v int64) *AttachKeyPairRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AttachKeyPairRequest) Validate() error {
	return dara.Validate(s)
}
