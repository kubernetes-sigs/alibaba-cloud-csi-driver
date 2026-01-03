// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDiskDefaultKMSKeyIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ResetDiskDefaultKMSKeyIdRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ResetDiskDefaultKMSKeyIdRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *ResetDiskDefaultKMSKeyIdRequest
	GetResourceOwnerId() *int64
}

type ResetDiskDefaultKMSKeyIdRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region for which you want to disable Account-level EBS Default Encryption. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/2679950.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ResetDiskDefaultKMSKeyIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetDiskDefaultKMSKeyIdRequest) GoString() string {
	return s.String()
}

func (s *ResetDiskDefaultKMSKeyIdRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResetDiskDefaultKMSKeyIdRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetDiskDefaultKMSKeyIdRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ResetDiskDefaultKMSKeyIdRequest) SetOwnerId(v int64) *ResetDiskDefaultKMSKeyIdRequest {
	s.OwnerId = &v
	return s
}

func (s *ResetDiskDefaultKMSKeyIdRequest) SetRegionId(v string) *ResetDiskDefaultKMSKeyIdRequest {
	s.RegionId = &v
	return s
}

func (s *ResetDiskDefaultKMSKeyIdRequest) SetResourceOwnerId(v int64) *ResetDiskDefaultKMSKeyIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResetDiskDefaultKMSKeyIdRequest) Validate() error {
	return dara.Validate(s)
}
