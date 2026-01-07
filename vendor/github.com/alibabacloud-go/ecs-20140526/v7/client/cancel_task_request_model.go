// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *CancelTaskRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CancelTaskRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CancelTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelTaskRequest
	GetResourceOwnerId() *int64
	SetTaskId(v string) *CancelTaskRequest
	GetTaskId() *string
}

type CancelTaskRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the task. You can call the [DescribeTasks](https://help.aliyun.com/document_detail/25622.html) operation to query the list of task IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-bp198jigq7l0h5ac****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CancelTaskRequest) SetOwnerId(v int64) *CancelTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelTaskRequest) SetRegionId(v string) *CancelTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CancelTaskRequest) SetResourceOwnerAccount(v string) *CancelTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelTaskRequest) SetResourceOwnerId(v int64) *CancelTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelTaskRequest) SetTaskId(v string) *CancelTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CancelTaskRequest) Validate() error {
	return dara.Validate(s)
}
