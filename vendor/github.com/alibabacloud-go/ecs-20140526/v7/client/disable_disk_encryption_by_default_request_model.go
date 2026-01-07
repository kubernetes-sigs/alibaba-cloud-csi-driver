// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDiskEncryptionByDefaultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DisableDiskEncryptionByDefaultRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DisableDiskEncryptionByDefaultRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *DisableDiskEncryptionByDefaultRequest
	GetResourceOwnerId() *int64
}

type DisableDiskEncryptionByDefaultRequest struct {
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

func (s DisableDiskEncryptionByDefaultRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableDiskEncryptionByDefaultRequest) GoString() string {
	return s.String()
}

func (s *DisableDiskEncryptionByDefaultRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisableDiskEncryptionByDefaultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableDiskEncryptionByDefaultRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DisableDiskEncryptionByDefaultRequest) SetOwnerId(v int64) *DisableDiskEncryptionByDefaultRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableDiskEncryptionByDefaultRequest) SetRegionId(v string) *DisableDiskEncryptionByDefaultRequest {
	s.RegionId = &v
	return s
}

func (s *DisableDiskEncryptionByDefaultRequest) SetResourceOwnerId(v int64) *DisableDiskEncryptionByDefaultRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisableDiskEncryptionByDefaultRequest) Validate() error {
	return dara.Validate(s)
}
