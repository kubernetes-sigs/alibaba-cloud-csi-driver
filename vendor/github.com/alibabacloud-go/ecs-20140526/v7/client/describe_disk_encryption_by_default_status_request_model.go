// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskEncryptionByDefaultStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeDiskEncryptionByDefaultStatusRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDiskEncryptionByDefaultStatusRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *DescribeDiskEncryptionByDefaultStatusRequest
	GetResourceOwnerId() *int64
}

type DescribeDiskEncryptionByDefaultStatusRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s DescribeDiskEncryptionByDefaultStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskEncryptionByDefaultStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskEncryptionByDefaultStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDiskEncryptionByDefaultStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiskEncryptionByDefaultStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDiskEncryptionByDefaultStatusRequest) SetOwnerId(v int64) *DescribeDiskEncryptionByDefaultStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDiskEncryptionByDefaultStatusRequest) SetRegionId(v string) *DescribeDiskEncryptionByDefaultStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskEncryptionByDefaultStatusRequest) SetResourceOwnerId(v int64) *DescribeDiskEncryptionByDefaultStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDiskEncryptionByDefaultStatusRequest) Validate() error {
	return dara.Validate(s)
}
