// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeTaskAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeTaskAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeTaskAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTaskAttributeRequest
	GetResourceOwnerId() *int64
	SetTaskId(v string) *DescribeTaskAttributeRequest
	GetTaskId() *string
}

type DescribeTaskAttributeRequest struct {
	// RAM用户的虚拟账号ID。
	//
	// example:
	//
	// 155780923770
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the task. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 资源主账号的账号名称。
	//
	// example:
	//
	// ECSforCloud
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// 资源主账号的ID，亦即UID。
	//
	// example:
	//
	// 155780923770
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the task. You can call the [DescribeTasks](https://help.aliyun.com/document_detail/25622.html) operation to query the list of task IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-ce946ntx4wr****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTaskAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTaskAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTaskAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTaskAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTaskAttributeRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTaskAttributeRequest) SetOwnerId(v int64) *DescribeTaskAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTaskAttributeRequest) SetRegionId(v string) *DescribeTaskAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTaskAttributeRequest) SetResourceOwnerAccount(v string) *DescribeTaskAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTaskAttributeRequest) SetResourceOwnerId(v int64) *DescribeTaskAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTaskAttributeRequest) SetTaskId(v string) *DescribeTaskAttributeRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskAttributeRequest) Validate() error {
	return dara.Validate(s)
}
