// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKeyPairsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludePublicKey(v bool) *DescribeKeyPairsRequest
	GetIncludePublicKey() *bool
	SetKeyPairFingerPrint(v string) *DescribeKeyPairsRequest
	GetKeyPairFingerPrint() *string
	SetKeyPairName(v string) *DescribeKeyPairsRequest
	GetKeyPairName() *string
	SetOwnerId(v int64) *DescribeKeyPairsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeKeyPairsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeKeyPairsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeKeyPairsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeKeyPairsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeKeyPairsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeKeyPairsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeKeyPairsRequestTag) *DescribeKeyPairsRequest
	GetTag() []*DescribeKeyPairsRequestTag
}

type DescribeKeyPairsRequest struct {
	// Specifies whether to include PublicKey in the response. Default value: false.
	//
	// example:
	//
	// false
	IncludePublicKey *bool `json:"IncludePublicKey,omitempty" xml:"IncludePublicKey,omitempty"`
	// The fingerprint of the key pair. The message-digest algorithm 5 (MD5) is used based on the public key fingerprint format defined in RFC 4716. For more information, see [RFC 4716](https://tools.ietf.org/html/rfc4716).
	//
	// example:
	//
	// ABC1234567
	KeyPairFingerPrint *string `json:"KeyPairFingerPrint,omitempty" xml:"KeyPairFingerPrint,omitempty"`
	// The name of the key pair. You can use the asterisk (\\*) symbol as a wildcard in regular expressions to perform a fuzzy search for key pairs. Sample patterns:
	//
	// 	- `*SshKey`: queries key pairs whose names end with SshKey, including the key pair named SshKey.
	//
	// 	- `SshKey*`: queries key pairs whose names start with SshKey, including the key pair named SshKey.
	//
	// 	- `*SshKey*`: queries key pairs whose names include SshKey, including the key pair named SshKey.
	//
	// 	- `SshKey`: queries the key pair named SshKey.
	//
	// example:
	//
	// *SshKey*
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the key pair. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. If this parameter is specified to query resources, up to 1,000 resources that belong to the specified resource group can be displayed in the response.
	//
	// >  Resources in the default resource group are displayed in the response regardless of how this parameter is set.
	//
	// example:
	//
	// rg-amnhr7u7c7hj****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*DescribeKeyPairsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeKeyPairsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyPairsRequest) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsRequest) GetIncludePublicKey() *bool {
	return s.IncludePublicKey
}

func (s *DescribeKeyPairsRequest) GetKeyPairFingerPrint() *string {
	return s.KeyPairFingerPrint
}

func (s *DescribeKeyPairsRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DescribeKeyPairsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeKeyPairsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeKeyPairsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeKeyPairsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeKeyPairsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeKeyPairsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeKeyPairsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeKeyPairsRequest) GetTag() []*DescribeKeyPairsRequestTag {
	return s.Tag
}

func (s *DescribeKeyPairsRequest) SetIncludePublicKey(v bool) *DescribeKeyPairsRequest {
	s.IncludePublicKey = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetKeyPairFingerPrint(v string) *DescribeKeyPairsRequest {
	s.KeyPairFingerPrint = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetKeyPairName(v string) *DescribeKeyPairsRequest {
	s.KeyPairName = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetOwnerId(v int64) *DescribeKeyPairsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetPageNumber(v int32) *DescribeKeyPairsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetPageSize(v int32) *DescribeKeyPairsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetRegionId(v string) *DescribeKeyPairsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetResourceGroupId(v string) *DescribeKeyPairsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetResourceOwnerAccount(v string) *DescribeKeyPairsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetResourceOwnerId(v int64) *DescribeKeyPairsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeKeyPairsRequest) SetTag(v []*DescribeKeyPairsRequestTag) *DescribeKeyPairsRequest {
	s.Tag = v
	return s
}

func (s *DescribeKeyPairsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeKeyPairsRequestTag struct {
	// The key of tag N of the key pair. Valid values of N: 1 to 20.
	//
	// If a single tag is specified to query resources, up to 1,000 resources that are bound with this tag can be displayed in the response. If multiple tags are specified to query resources, up to 1,000 resources that are bound with all these tags can be displayed in the response. To query more than 1,000 resources that have specified tags added, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the key pair. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeKeyPairsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyPairsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeKeyPairsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeKeyPairsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeKeyPairsRequestTag) SetKey(v string) *DescribeKeyPairsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeKeyPairsRequestTag) SetValue(v string) *DescribeKeyPairsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeKeyPairsRequestTag) Validate() error {
	return dara.Validate(s)
}
