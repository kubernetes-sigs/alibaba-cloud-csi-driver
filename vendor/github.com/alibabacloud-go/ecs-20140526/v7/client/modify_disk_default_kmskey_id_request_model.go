// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskDefaultKMSKeyIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKMSKeyId(v string) *ModifyDiskDefaultKMSKeyIdRequest
	GetKMSKeyId() *string
	SetOwnerId(v int64) *ModifyDiskDefaultKMSKeyIdRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDiskDefaultKMSKeyIdRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *ModifyDiskDefaultKMSKeyIdRequest
	GetResourceOwnerId() *int64
}

type ModifyDiskDefaultKMSKeyIdRequest struct {
	// The ID of the new KMS key.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/2679950.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDiskDefaultKMSKeyIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskDefaultKMSKeyIdRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskDefaultKMSKeyIdRequest) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *ModifyDiskDefaultKMSKeyIdRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDiskDefaultKMSKeyIdRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDiskDefaultKMSKeyIdRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDiskDefaultKMSKeyIdRequest) SetKMSKeyId(v string) *ModifyDiskDefaultKMSKeyIdRequest {
	s.KMSKeyId = &v
	return s
}

func (s *ModifyDiskDefaultKMSKeyIdRequest) SetOwnerId(v int64) *ModifyDiskDefaultKMSKeyIdRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDiskDefaultKMSKeyIdRequest) SetRegionId(v string) *ModifyDiskDefaultKMSKeyIdRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDiskDefaultKMSKeyIdRequest) SetResourceOwnerId(v int64) *ModifyDiskDefaultKMSKeyIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDiskDefaultKMSKeyIdRequest) Validate() error {
	return dara.Validate(s)
}
