// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoSnapshotPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotPoliciesRequest
	GetAutoSnapshotPolicyId() *string
	SetFileSystemType(v string) *DescribeAutoSnapshotPoliciesRequest
	GetFileSystemType() *string
	SetPageNumber(v int32) *DescribeAutoSnapshotPoliciesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAutoSnapshotPoliciesRequest
	GetPageSize() *int32
}

type DescribeAutoSnapshotPoliciesRequest struct {
	// The ID of the automatic snapshot policy.
	//
	// example:
	//
	// sp-extreme-233e6****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The type of the file system.
	//
	// Valid value: extreme, which indicates Extreme File Storage NAS (NAS) file systems.
	//
	// example:
	//
	// extreme
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
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
}

func (s DescribeAutoSnapshotPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPoliciesRequest) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *DescribeAutoSnapshotPoliciesRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DescribeAutoSnapshotPoliciesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoSnapshotPoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAutoSnapshotPoliciesRequest) SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotPoliciesRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesRequest) SetFileSystemType(v string) *DescribeAutoSnapshotPoliciesRequest {
	s.FileSystemType = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesRequest) SetPageNumber(v int32) *DescribeAutoSnapshotPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesRequest) SetPageSize(v int32) *DescribeAutoSnapshotPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
