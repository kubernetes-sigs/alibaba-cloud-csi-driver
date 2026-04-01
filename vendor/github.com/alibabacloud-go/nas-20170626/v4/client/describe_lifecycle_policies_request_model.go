// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLifecyclePoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribeLifecyclePoliciesRequest
	GetDescription() *string
	SetFileSystemId(v string) *DescribeLifecyclePoliciesRequest
	GetFileSystemId() *string
	SetLifecyclePolicyId(v string) *DescribeLifecyclePoliciesRequest
	GetLifecyclePolicyId() *string
	SetLifecyclePolicyName(v string) *DescribeLifecyclePoliciesRequest
	GetLifecyclePolicyName() *string
	SetLifecyclePolicyType(v string) *DescribeLifecyclePoliciesRequest
	GetLifecyclePolicyType() *string
	SetPageNumber(v int32) *DescribeLifecyclePoliciesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLifecyclePoliciesRequest
	GetPageSize() *int32
	SetPath(v string) *DescribeLifecyclePoliciesRequest
	GetPath() *string
	SetStorageType(v string) *DescribeLifecyclePoliciesRequest
	GetStorageType() *string
}

type DescribeLifecyclePoliciesRequest struct {
	// The description of the policy.
	//
	// >  Only CPFS for Lingjun supports this parameter.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the file system.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the lifecycle policy.
	//
	// example:
	//
	// lc-xxx
	LifecyclePolicyId *string `json:"LifecyclePolicyId,omitempty" xml:"LifecyclePolicyId,omitempty"`
	// The name of the lifecycle policy. The naming rules are as follows:
	//
	// The name must be 3 to 64 characters in length and must start with a letter. It can contain letters, digits, underscores (_), and hyphens (-).
	//
	// >  Optional for General-purpose NAS file systems. If this parameter is provided, it takes precedence over LifecyclePolicyId. If left empty, LifecyclePolicyId is used.
	//
	// example:
	//
	// lifecyclepolicy_01
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
	// The type of the lifecycle policy.
	//
	// Valid values:
	//
	// 	- Auto: The job is automatically triggered.
	//
	// 	- OnDemand: On-demand execution.
	//
	// >  Only CPFS for Lingjun supports this parameter.
	//
	// example:
	//
	// Auto
	LifecyclePolicyType *string `json:"LifecyclePolicyType,omitempty" xml:"LifecyclePolicyType,omitempty"`
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
	// Filter by path.
	//
	// >  Only CPFS for Lingjun supports this parameter.
	//
	// example:
	//
	// /abc/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The storage class.
	//
	// 	- InfrequentAccess: the Infrequent Access (IA) storage class.
	//
	// 	- Archive: the Archive storage class.
	//
	// >  If StorageType is not specified, all lifecycle policies are returned.
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

func (s *DescribeLifecyclePoliciesRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeLifecyclePoliciesRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeLifecyclePoliciesRequest) GetLifecyclePolicyId() *string {
	return s.LifecyclePolicyId
}

func (s *DescribeLifecyclePoliciesRequest) GetLifecyclePolicyName() *string {
	return s.LifecyclePolicyName
}

func (s *DescribeLifecyclePoliciesRequest) GetLifecyclePolicyType() *string {
	return s.LifecyclePolicyType
}

func (s *DescribeLifecyclePoliciesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLifecyclePoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLifecyclePoliciesRequest) GetPath() *string {
	return s.Path
}

func (s *DescribeLifecyclePoliciesRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeLifecyclePoliciesRequest) SetDescription(v string) *DescribeLifecyclePoliciesRequest {
	s.Description = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetFileSystemId(v string) *DescribeLifecyclePoliciesRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetLifecyclePolicyId(v string) *DescribeLifecyclePoliciesRequest {
	s.LifecyclePolicyId = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetLifecyclePolicyName(v string) *DescribeLifecyclePoliciesRequest {
	s.LifecyclePolicyName = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetLifecyclePolicyType(v string) *DescribeLifecyclePoliciesRequest {
	s.LifecyclePolicyType = &v
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

func (s *DescribeLifecyclePoliciesRequest) SetPath(v string) *DescribeLifecyclePoliciesRequest {
	s.Path = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetStorageType(v string) *DescribeLifecyclePoliciesRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) Validate() error {
	return dara.Validate(s)
}
