// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskDefaultKMSKeyIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeDiskDefaultKMSKeyIdRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDiskDefaultKMSKeyIdRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *DescribeDiskDefaultKMSKeyIdRequest
	GetResourceOwnerId() *int64
}

type DescribeDiskDefaultKMSKeyIdRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDiskDefaultKMSKeyIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskDefaultKMSKeyIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskDefaultKMSKeyIdRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDiskDefaultKMSKeyIdRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiskDefaultKMSKeyIdRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDiskDefaultKMSKeyIdRequest) SetOwnerId(v int64) *DescribeDiskDefaultKMSKeyIdRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDiskDefaultKMSKeyIdRequest) SetRegionId(v string) *DescribeDiskDefaultKMSKeyIdRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskDefaultKMSKeyIdRequest) SetResourceOwnerId(v int64) *DescribeDiskDefaultKMSKeyIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDiskDefaultKMSKeyIdRequest) Validate() error {
	return dara.Validate(s)
}
