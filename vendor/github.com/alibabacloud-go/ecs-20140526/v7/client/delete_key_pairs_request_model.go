// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKeyPairsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPairNames(v string) *DeleteKeyPairsRequest
	GetKeyPairNames() *string
	SetOwnerId(v int64) *DeleteKeyPairsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteKeyPairsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteKeyPairsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteKeyPairsRequest
	GetResourceOwnerId() *int64
}

type DeleteKeyPairsRequest struct {
	// The names of SSH key pairs. The value can be a JSON array that consists of up to 100 SSH key pair names. Separate the names with commas (,).
	//
	// >  Before you delete an SSH key pair, you can call the [DescribeKeyPairs](https://help.aliyun.com/document_detail/51773.html) operation to query existing key pairs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["skp-bp67acfmxazb41****", "skp-bp67acfmxazb42****", … "skp-bp67acfmxazb4p3***"]
	KeyPairNames *string `json:"KeyPairNames,omitempty" xml:"KeyPairNames,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent list of regions.
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

func (s DeleteKeyPairsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKeyPairsRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsRequest) GetKeyPairNames() *string {
	return s.KeyPairNames
}

func (s *DeleteKeyPairsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteKeyPairsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteKeyPairsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteKeyPairsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteKeyPairsRequest) SetKeyPairNames(v string) *DeleteKeyPairsRequest {
	s.KeyPairNames = &v
	return s
}

func (s *DeleteKeyPairsRequest) SetOwnerId(v int64) *DeleteKeyPairsRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteKeyPairsRequest) SetRegionId(v string) *DeleteKeyPairsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteKeyPairsRequest) SetResourceOwnerAccount(v string) *DeleteKeyPairsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteKeyPairsRequest) SetResourceOwnerId(v int64) *DeleteKeyPairsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteKeyPairsRequest) Validate() error {
	return dara.Validate(s)
}
