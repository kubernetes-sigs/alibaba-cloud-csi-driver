// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLifecyclePolicyLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLifecyclePolicyLogs(v []*DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) *DescribeLifecyclePolicyLogsResponseBody
	GetLifecyclePolicyLogs() []*DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs
	SetPageNumber(v int32) *DescribeLifecyclePolicyLogsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLifecyclePolicyLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLifecyclePolicyLogsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeLifecyclePolicyLogsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *DescribeLifecyclePolicyLogsResponseBody
	GetTotalCount() *int32
}

type DescribeLifecyclePolicyLogsResponseBody struct {
	// The execution logs of the lifecycle policy.
	LifecyclePolicyLogs []*DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs `json:"LifecyclePolicyLogs,omitempty" xml:"LifecyclePolicyLogs,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of log entries per page.
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
	// The request status.
	//
	// Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of log entries.
	//
	// example:
	//
	// 36
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLifecyclePolicyLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecyclePolicyLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePolicyLogsResponseBody) GetLifecyclePolicyLogs() []*DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs {
	return s.LifecyclePolicyLogs
}

func (s *DescribeLifecyclePolicyLogsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLifecyclePolicyLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLifecyclePolicyLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLifecyclePolicyLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeLifecyclePolicyLogsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeLifecyclePolicyLogsResponseBody) SetLifecyclePolicyLogs(v []*DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) *DescribeLifecyclePolicyLogsResponseBody {
	s.LifecyclePolicyLogs = v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBody) SetPageNumber(v int32) *DescribeLifecyclePolicyLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBody) SetPageSize(v int32) *DescribeLifecyclePolicyLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBody) SetRequestId(v string) *DescribeLifecyclePolicyLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBody) SetSuccess(v bool) *DescribeLifecyclePolicyLogsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBody) SetTotalCount(v int32) *DescribeLifecyclePolicyLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBody) Validate() error {
	if s.LifecyclePolicyLogs != nil {
		for _, item := range s.LifecyclePolicyLogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs struct {
	// The time when the task was created. The time follows the ISO 8601 standard in UTC. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2025-10-20T02:25:07Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The file data expiration and deletion rules. A maximum of one rule can be configured.
	DeleteRules []*DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsDeleteRules `json:"DeleteRules,omitempty" xml:"DeleteRules,omitempty" type:"Repeated"`
	// The execution paths of the specified retrieval task.
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// The file data retrieval rules.
	RetrieveRules []*DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsRetrieveRules `json:"RetrieveRules,omitempty" xml:"RetrieveRules,omitempty" type:"Repeated"`
	// The status of the data retrieval task. Valid values:
	//
	// - PENDING: Being created.
	//
	// - RUNNING: Running.
	//
	// - STOPPED: Stopped.
	//
	// - FINISHED: Finished.
	//
	// - FAILED: Failed.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tiered storage type. Valid values:
	//
	// - InfrequentAccess: IA storage class (default).
	//
	// - Archive: Archive storage.
	//
	// example:
	//
	// InfrequentAccess
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The task overview.
	//
	// example:
	//
	// Total tasks: 100000, success tasks: 100000
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The file data transit rules.
	TransitRules []*DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsTransitRules `json:"TransitRules,omitempty" xml:"TransitRules,omitempty" type:"Repeated"`
}

func (s DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) GetDeleteRules() []*DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsDeleteRules {
	return s.DeleteRules
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) GetPaths() []*string {
	return s.Paths
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) GetRetrieveRules() []*DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsRetrieveRules {
	return s.RetrieveRules
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) GetStatus() *string {
	return s.Status
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) GetSummary() *string {
	return s.Summary
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) GetTransitRules() []*DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsTransitRules {
	return s.TransitRules
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) SetCreateTime(v string) *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs {
	s.CreateTime = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) SetDeleteRules(v []*DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsDeleteRules) *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs {
	s.DeleteRules = v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) SetPaths(v []*string) *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs {
	s.Paths = v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) SetRetrieveRules(v []*DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsRetrieveRules) *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs {
	s.RetrieveRules = v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) SetStatus(v string) *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs {
	s.Status = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) SetStorageType(v string) *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs {
	s.StorageType = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) SetSummary(v string) *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs {
	s.Summary = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) SetTransitRules(v []*DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsTransitRules) *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs {
	s.TransitRules = v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogs) Validate() error {
	if s.DeleteRules != nil {
		for _, item := range s.DeleteRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsDeleteRules struct {
	// The attribute of the rule.
	//
	// Valid values:
	//
	// - Atime: the access time of the file.
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The threshold of the rule.
	//
	// Valid values:
	//
	// - If Attribute is set to Atime, this parameter specifies the number of days since the file was last accessed. Valid values: 1 to 365.
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsDeleteRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsDeleteRules) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsDeleteRules) GetAttribute() *string {
	return s.Attribute
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsDeleteRules) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsDeleteRules) SetAttribute(v string) *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsDeleteRules {
	s.Attribute = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsDeleteRules) SetThreshold(v string) *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsDeleteRules {
	s.Threshold = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsDeleteRules) Validate() error {
	return dara.Validate(s)
}

type DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsRetrieveRules struct {
	// The attribute of the rule. Valid values:
	//
	// - RetrieveType: the retrieval method.
	//
	// example:
	//
	// RetrieveType
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The threshold of the rule. Valid values:
	//
	// - RetrieveType
	//
	//     - AfterVisit: Supported when LifecyclePolicyType=Auto. Indicates best-effort recall on visit.
	//
	//     - All: Supported when LifecyclePolicyType=OnDemand. Indicates retrieving all data.
	//
	// example:
	//
	// All
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsRetrieveRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsRetrieveRules) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsRetrieveRules) GetAttribute() *string {
	return s.Attribute
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsRetrieveRules) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsRetrieveRules) SetAttribute(v string) *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsRetrieveRules {
	s.Attribute = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsRetrieveRules) SetThreshold(v string) *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsRetrieveRules {
	s.Threshold = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsRetrieveRules) Validate() error {
	return dara.Validate(s)
}

type DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsTransitRules struct {
	// The attribute of the rule.
	//
	// Valid values:
	//
	// - Atime: the access time of the file.
	//
	// example:
	//
	// Atime
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The threshold of the rule.
	//
	// Valid values:
	//
	// - If Attribute is set to Atime, this parameter specifies the number of days since the file was last accessed. Valid values: 1 to 365.
	//
	// example:
	//
	// 3
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsTransitRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsTransitRules) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsTransitRules) GetAttribute() *string {
	return s.Attribute
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsTransitRules) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsTransitRules) SetAttribute(v string) *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsTransitRules {
	s.Attribute = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsTransitRules) SetThreshold(v string) *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsTransitRules {
	s.Threshold = &v
	return s
}

func (s *DescribeLifecyclePolicyLogsResponseBodyLifecyclePolicyLogsTransitRules) Validate() error {
	return dara.Validate(s)
}
