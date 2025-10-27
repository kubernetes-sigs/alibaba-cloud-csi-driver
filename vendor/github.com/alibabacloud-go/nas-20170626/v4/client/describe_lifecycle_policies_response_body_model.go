// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLifecyclePoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLifecyclePolicies(v []*DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) *DescribeLifecyclePoliciesResponseBody
	GetLifecyclePolicies() []*DescribeLifecyclePoliciesResponseBodyLifecyclePolicies
	SetPageNumber(v int32) *DescribeLifecyclePoliciesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLifecyclePoliciesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLifecyclePoliciesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeLifecyclePoliciesResponseBody
	GetTotalCount() *int32
}

type DescribeLifecyclePoliciesResponseBody struct {
	// The queried lifecycle policies.
	LifecyclePolicies []*DescribeLifecyclePoliciesResponseBodyLifecyclePolicies `json:"LifecyclePolicies,omitempty" xml:"LifecyclePolicies,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BC7C825C-5F65-4B56-BEF6-98C56C7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of lifecycle policies.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLifecyclePoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecyclePoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePoliciesResponseBody) GetLifecyclePolicies() []*DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	return s.LifecyclePolicies
}

func (s *DescribeLifecyclePoliciesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLifecyclePoliciesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLifecyclePoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLifecyclePoliciesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeLifecyclePoliciesResponseBody) SetLifecyclePolicies(v []*DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) *DescribeLifecyclePoliciesResponseBody {
	s.LifecyclePolicies = v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBody) SetPageNumber(v int32) *DescribeLifecyclePoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBody) SetPageSize(v int32) *DescribeLifecyclePoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBody) SetRequestId(v string) *DescribeLifecyclePoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBody) SetTotalCount(v int32) *DescribeLifecyclePoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLifecyclePoliciesResponseBodyLifecyclePolicies struct {
	// The time when the lifecycle policy was created.
	//
	// The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// example:
	//
	// 2019-10-30T10:08:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the file system.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The name of the lifecycle policy.
	//
	// example:
	//
	// lifecyclepolicy_01
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
	// The management rule that is associated with the lifecycle policy.
	//
	// Valid values:
	//
	// 	- DEFAULT_ATIME_14: Files that are not accessed in the last 14 days are dumped to the IA storage medium.
	//
	// 	- DEFAULT_ATIME_30: Files that are not accessed in the last 30 days are dumped to the IA storage medium.
	//
	// 	- DEFAULT_ATIME_60: Files that are not accessed in the last 60 days are dumped to the IA storage medium.
	//
	// 	- DEFAULT_ATIME_90: Files that are not accessed in the last 90 days are dumped to the IA storage medium.
	//
	// example:
	//
	// DEFAULT_ATIME_14
	LifecycleRuleName *string `json:"LifecycleRuleName,omitempty" xml:"LifecycleRuleName,omitempty"`
	// The absolute path of a directory with which the lifecycle policy is associated.
	//
	// example:
	//
	// /pathway/to/folder
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The absolute paths to multiple directories associated with the lifecycle policy.
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// The storage type of the data that is dumped to the IA storage medium.
	//
	// Default value: InfrequentAccess (IA).
	//
	// example:
	//
	// InfrequentAccess
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) GetLifecyclePolicyName() *string {
	return s.LifecyclePolicyName
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) GetLifecycleRuleName() *string {
	return s.LifecycleRuleName
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) GetPath() *string {
	return s.Path
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) GetPaths() []*string {
	return s.Paths
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetCreateTime(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.CreateTime = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetFileSystemId(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.FileSystemId = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetLifecyclePolicyName(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.LifecyclePolicyName = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetLifecycleRuleName(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.LifecycleRuleName = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetPath(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.Path = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetPaths(v []*string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.Paths = v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetStorageType(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.StorageType = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) Validate() error {
	return dara.Validate(s)
}
