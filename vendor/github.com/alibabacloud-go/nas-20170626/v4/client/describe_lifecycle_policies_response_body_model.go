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
	if s.LifecyclePolicies != nil {
		for _, item := range s.LifecyclePolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	// Description
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
	// The name of the lifecycle policy.
	//
	// example:
	//
	// lifecyclepolicy_01
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
	// The type of the lifecycle policy.
	//
	// example:
	//
	// Auto
	LifecyclePolicyType *string `json:"LifecyclePolicyType,omitempty" xml:"LifecyclePolicyType,omitempty"`
	// The management rule that is associated with the lifecycle policy.
	//
	// Valid values:
	//
	// 	- DEFAULT_ATIME_14: Files that are not accessed in the last 14 days.
	//
	// 	- DEFAULT_ATIME_30: Files that are not accessed in the last 30 days.
	//
	// 	- DEFAULT_ATIME_60: Files that are not accessed in the last 60 days.
	//
	// 	- DEFAULT_ATIME_90: Files that are not accessed in the last 90 days.
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
	// The absolute paths of directories with which the lifecycle policy is associated.
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// File data retrieval rules.
	RetrieveRules []*DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesRetrieveRules `json:"RetrieveRules,omitempty" xml:"RetrieveRules,omitempty" type:"Repeated"`
	// The storage class.
	//
	// 	- InfrequentAccess: the IA storage class.
	//
	// 	- Archive: the Archive storage class.
	//
	// example:
	//
	// InfrequentAccess
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// Data transition rules.
	TransitRules []*DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesTransitRules `json:"TransitRules,omitempty" xml:"TransitRules,omitempty" type:"Repeated"`
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

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) GetDescription() *string {
	return s.Description
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) GetLifecyclePolicyId() *string {
	return s.LifecyclePolicyId
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) GetLifecyclePolicyName() *string {
	return s.LifecyclePolicyName
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) GetLifecyclePolicyType() *string {
	return s.LifecyclePolicyType
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

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) GetRetrieveRules() []*DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesRetrieveRules {
	return s.RetrieveRules
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) GetTransitRules() []*DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesTransitRules {
	return s.TransitRules
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetCreateTime(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.CreateTime = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetDescription(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.Description = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetFileSystemId(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.FileSystemId = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetLifecyclePolicyId(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.LifecyclePolicyId = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetLifecyclePolicyName(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.LifecyclePolicyName = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetLifecyclePolicyType(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.LifecyclePolicyType = &v
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

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetRetrieveRules(v []*DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesRetrieveRules) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.RetrieveRules = v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetStorageType(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.StorageType = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) SetTransitRules(v []*DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesTransitRules) *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies {
	s.TransitRules = v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePolicies) Validate() error {
	if s.RetrieveRules != nil {
		for _, item := range s.RetrieveRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TransitRules != nil {
		for _, item := range s.TransitRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesRetrieveRules struct {
	// Attribute of the rule.
	//
	// example:
	//
	// RetrieveType
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// Threshold for the rule.
	//
	// example:
	//
	// All
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesRetrieveRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesRetrieveRules) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesRetrieveRules) GetAttribute() *string {
	return s.Attribute
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesRetrieveRules) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesRetrieveRules) SetAttribute(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesRetrieveRules {
	s.Attribute = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesRetrieveRules) SetThreshold(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesRetrieveRules {
	s.Threshold = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesRetrieveRules) Validate() error {
	return dara.Validate(s)
}

type DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesTransitRules struct {
	// Attribute of the rule.
	//
	// example:
	//
	// Atime
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// Threshold for the rule.
	//
	// example:
	//
	// 3
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesTransitRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesTransitRules) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesTransitRules) GetAttribute() *string {
	return s.Attribute
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesTransitRules) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesTransitRules) SetAttribute(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesTransitRules {
	s.Attribute = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesTransitRules) SetThreshold(v string) *DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesTransitRules {
	s.Threshold = &v
	return s
}

func (s *DescribeLifecyclePoliciesResponseBodyLifecyclePoliciesTransitRules) Validate() error {
	return dara.Validate(s)
}
