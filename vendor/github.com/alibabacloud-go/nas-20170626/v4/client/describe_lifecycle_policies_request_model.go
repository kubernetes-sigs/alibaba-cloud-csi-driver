// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLifecyclePoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DescribeLifecyclePoliciesRequest
	GetFileSystemId() *string
	SetLifecyclePolicyName(v string) *DescribeLifecyclePoliciesRequest
	GetLifecyclePolicyName() *string
	SetPageNumber(v int32) *DescribeLifecyclePoliciesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLifecyclePoliciesRequest
	GetPageSize() *int32
	SetStorageType(v string) *DescribeLifecyclePoliciesRequest
	GetStorageType() *string
}

type DescribeLifecyclePoliciesRequest struct {
	// The ID of the file system.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The name of the lifecycle policy. The name must meet the following conventions:
	//
	// The name must be 3 to 64 characters in length and must start with a letter. It can contain letters, digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// lifecyclepolicy_01
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The storage class.
	//
	// 	- InfrequentAccess: the Infrequent Access (IA) storage class.
	//
	// 	- Archive: the Archive storage class.
	//
	// >  If the StorageType parameter is not specified, data retrieval tasks of all types are returned.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// InfrequentAccess
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeLifecyclePoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecyclePoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePoliciesRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeLifecyclePoliciesRequest) GetLifecyclePolicyName() *string {
	return s.LifecyclePolicyName
}

func (s *DescribeLifecyclePoliciesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLifecyclePoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLifecyclePoliciesRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeLifecyclePoliciesRequest) SetFileSystemId(v string) *DescribeLifecyclePoliciesRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetLifecyclePolicyName(v string) *DescribeLifecyclePoliciesRequest {
	s.LifecyclePolicyName = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetPageNumber(v int32) *DescribeLifecyclePoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetPageSize(v int32) *DescribeLifecyclePoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetStorageType(v string) *DescribeLifecyclePoliciesRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) Validate() error {
	return dara.Validate(s)
}
