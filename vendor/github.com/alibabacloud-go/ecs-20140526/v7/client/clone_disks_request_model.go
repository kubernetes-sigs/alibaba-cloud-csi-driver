// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneDisksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v []*CloneDisksRequestArn) *CloneDisksRequest
	GetArn() []*CloneDisksRequestArn
	SetBurstingEnabled(v bool) *CloneDisksRequest
	GetBurstingEnabled() *bool
	SetClientToken(v string) *CloneDisksRequest
	GetClientToken() *string
	SetDiskCategory(v string) *CloneDisksRequest
	GetDiskCategory() *string
	SetDiskName(v string) *CloneDisksRequest
	GetDiskName() *string
	SetDryRun(v string) *CloneDisksRequest
	GetDryRun() *string
	SetEncrypted(v bool) *CloneDisksRequest
	GetEncrypted() *bool
	SetKmsKeyId(v string) *CloneDisksRequest
	GetKmsKeyId() *string
	SetMultiAttach(v string) *CloneDisksRequest
	GetMultiAttach() *string
	SetOwnerId(v int64) *CloneDisksRequest
	GetOwnerId() *int64
	SetPerformanceLevel(v string) *CloneDisksRequest
	GetPerformanceLevel() *string
	SetProvisionedIops(v int64) *CloneDisksRequest
	GetProvisionedIops() *int64
	SetRegionId(v string) *CloneDisksRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CloneDisksRequest
	GetResourceGroupId() *string
	SetResourceOwnerId(v int64) *CloneDisksRequest
	GetResourceOwnerId() *int64
	SetSize(v int32) *CloneDisksRequest
	GetSize() *int32
	SetSourceDiskId(v string) *CloneDisksRequest
	GetSourceDiskId() *string
	SetTag(v []*CloneDisksRequestTag) *CloneDisksRequest
	GetTag() []*CloneDisksRequestTag
}

type CloneDisksRequest struct {
	Arn []*CloneDisksRequestArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// example:
	//
	// true
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cloud_essd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// example:
	//
	// MyDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// example:
	//
	// true
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// example:
	//
	// key-szz67b2f696f4wh9yeg5d
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Disabled
	MultiAttach *string `json:"MultiAttach,omitempty" xml:"MultiAttach,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// example:
	//
	// 10
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-bp199lyny9b3****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d-bp1d6tsvznfghy7y****
	SourceDiskId *string                 `json:"SourceDiskId,omitempty" xml:"SourceDiskId,omitempty"`
	Tag          []*CloneDisksRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CloneDisksRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneDisksRequest) GoString() string {
	return s.String()
}

func (s *CloneDisksRequest) GetArn() []*CloneDisksRequestArn {
	return s.Arn
}

func (s *CloneDisksRequest) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CloneDisksRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CloneDisksRequest) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *CloneDisksRequest) GetDiskName() *string {
	return s.DiskName
}

func (s *CloneDisksRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *CloneDisksRequest) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CloneDisksRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *CloneDisksRequest) GetMultiAttach() *string {
	return s.MultiAttach
}

func (s *CloneDisksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloneDisksRequest) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CloneDisksRequest) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CloneDisksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CloneDisksRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CloneDisksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloneDisksRequest) GetSize() *int32 {
	return s.Size
}

func (s *CloneDisksRequest) GetSourceDiskId() *string {
	return s.SourceDiskId
}

func (s *CloneDisksRequest) GetTag() []*CloneDisksRequestTag {
	return s.Tag
}

func (s *CloneDisksRequest) SetArn(v []*CloneDisksRequestArn) *CloneDisksRequest {
	s.Arn = v
	return s
}

func (s *CloneDisksRequest) SetBurstingEnabled(v bool) *CloneDisksRequest {
	s.BurstingEnabled = &v
	return s
}

func (s *CloneDisksRequest) SetClientToken(v string) *CloneDisksRequest {
	s.ClientToken = &v
	return s
}

func (s *CloneDisksRequest) SetDiskCategory(v string) *CloneDisksRequest {
	s.DiskCategory = &v
	return s
}

func (s *CloneDisksRequest) SetDiskName(v string) *CloneDisksRequest {
	s.DiskName = &v
	return s
}

func (s *CloneDisksRequest) SetDryRun(v string) *CloneDisksRequest {
	s.DryRun = &v
	return s
}

func (s *CloneDisksRequest) SetEncrypted(v bool) *CloneDisksRequest {
	s.Encrypted = &v
	return s
}

func (s *CloneDisksRequest) SetKmsKeyId(v string) *CloneDisksRequest {
	s.KmsKeyId = &v
	return s
}

func (s *CloneDisksRequest) SetMultiAttach(v string) *CloneDisksRequest {
	s.MultiAttach = &v
	return s
}

func (s *CloneDisksRequest) SetOwnerId(v int64) *CloneDisksRequest {
	s.OwnerId = &v
	return s
}

func (s *CloneDisksRequest) SetPerformanceLevel(v string) *CloneDisksRequest {
	s.PerformanceLevel = &v
	return s
}

func (s *CloneDisksRequest) SetProvisionedIops(v int64) *CloneDisksRequest {
	s.ProvisionedIops = &v
	return s
}

func (s *CloneDisksRequest) SetRegionId(v string) *CloneDisksRequest {
	s.RegionId = &v
	return s
}

func (s *CloneDisksRequest) SetResourceGroupId(v string) *CloneDisksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CloneDisksRequest) SetResourceOwnerId(v int64) *CloneDisksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloneDisksRequest) SetSize(v int32) *CloneDisksRequest {
	s.Size = &v
	return s
}

func (s *CloneDisksRequest) SetSourceDiskId(v string) *CloneDisksRequest {
	s.SourceDiskId = &v
	return s
}

func (s *CloneDisksRequest) SetTag(v []*CloneDisksRequestTag) *CloneDisksRequest {
	s.Tag = v
	return s
}

func (s *CloneDisksRequest) Validate() error {
	if s.Arn != nil {
		for _, item := range s.Arn {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type CloneDisksRequestArn struct {
	// example:
	//
	// null
	AssumeRoleFor *string `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// example:
	//
	// null
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// example:
	//
	// null
	Rolearn *string `json:"Rolearn,omitempty" xml:"Rolearn,omitempty"`
}

func (s CloneDisksRequestArn) String() string {
	return dara.Prettify(s)
}

func (s CloneDisksRequestArn) GoString() string {
	return s.String()
}

func (s *CloneDisksRequestArn) GetAssumeRoleFor() *string {
	return s.AssumeRoleFor
}

func (s *CloneDisksRequestArn) GetRoleType() *string {
	return s.RoleType
}

func (s *CloneDisksRequestArn) GetRolearn() *string {
	return s.Rolearn
}

func (s *CloneDisksRequestArn) SetAssumeRoleFor(v string) *CloneDisksRequestArn {
	s.AssumeRoleFor = &v
	return s
}

func (s *CloneDisksRequestArn) SetRoleType(v string) *CloneDisksRequestArn {
	s.RoleType = &v
	return s
}

func (s *CloneDisksRequestArn) SetRolearn(v string) *CloneDisksRequestArn {
	s.Rolearn = &v
	return s
}

func (s *CloneDisksRequestArn) Validate() error {
	return dara.Validate(s)
}

type CloneDisksRequestTag struct {
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CloneDisksRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CloneDisksRequestTag) GoString() string {
	return s.String()
}

func (s *CloneDisksRequestTag) GetKey() *string {
	return s.Key
}

func (s *CloneDisksRequestTag) GetValue() *string {
	return s.Value
}

func (s *CloneDisksRequestTag) SetKey(v string) *CloneDisksRequestTag {
	s.Key = &v
	return s
}

func (s *CloneDisksRequestTag) SetValue(v string) *CloneDisksRequestTag {
	s.Value = &v
	return s
}

func (s *CloneDisksRequestTag) Validate() error {
	return dara.Validate(s)
}
