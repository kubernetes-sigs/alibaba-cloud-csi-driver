// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudAssistantStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceCloudAssistantStatusSet(v *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) *DescribeCloudAssistantStatusResponseBody
	GetInstanceCloudAssistantStatusSet() *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet
	SetNextToken(v string) *DescribeCloudAssistantStatusResponseBody
	GetNextToken() *string
	SetPageNumber(v int64) *DescribeCloudAssistantStatusResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeCloudAssistantStatusResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeCloudAssistantStatusResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeCloudAssistantStatusResponseBody
	GetTotalCount() *int64
}

type DescribeCloudAssistantStatusResponseBody struct {
	// Details about the installation status of Cloud Assistant on the instances.
	InstanceCloudAssistantStatusSet *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet `json:"InstanceCloudAssistantStatusSet,omitempty" xml:"InstanceCloudAssistantStatusSet,omitempty" type:"Struct"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instances.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudAssistantStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantStatusResponseBody) GetInstanceCloudAssistantStatusSet() *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet {
	return s.InstanceCloudAssistantStatusSet
}

func (s *DescribeCloudAssistantStatusResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCloudAssistantStatusResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCloudAssistantStatusResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCloudAssistantStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudAssistantStatusResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeCloudAssistantStatusResponseBody) SetInstanceCloudAssistantStatusSet(v *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) *DescribeCloudAssistantStatusResponseBody {
	s.InstanceCloudAssistantStatusSet = v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBody) SetNextToken(v string) *DescribeCloudAssistantStatusResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBody) SetPageNumber(v int64) *DescribeCloudAssistantStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBody) SetPageSize(v int64) *DescribeCloudAssistantStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBody) SetRequestId(v string) *DescribeCloudAssistantStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBody) SetTotalCount(v int64) *DescribeCloudAssistantStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBody) Validate() error {
	if s.InstanceCloudAssistantStatusSet != nil {
		if err := s.InstanceCloudAssistantStatusSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet struct {
	InstanceCloudAssistantStatus []*DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus `json:"InstanceCloudAssistantStatus,omitempty" xml:"InstanceCloudAssistantStatus,omitempty" type:"Repeated"`
}

func (s DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) GetInstanceCloudAssistantStatus() []*DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus {
	return s.InstanceCloudAssistantStatus
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) SetInstanceCloudAssistantStatus(v []*DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet {
	s.InstanceCloudAssistantStatus = v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSet) Validate() error {
	if s.InstanceCloudAssistantStatus != nil {
		for _, item := range s.InstanceCloudAssistantStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus struct {
	// The number of tasks that Cloud Assistant was running on the instance.
	//
	// example:
	//
	// 0
	ActiveTaskCount *int64 `json:"ActiveTaskCount,omitempty" xml:"ActiveTaskCount,omitempty"`
	// Indicates whether Cloud Assistant is running on the instance. Valid values:
	//
	// 	- true: Heartbeats are detected in the last 2 minutes.
	//
	// 	- false: No heartbeats are detected in the last 2 minutes.
	//
	// example:
	//
	// true
	CloudAssistantStatus *string `json:"CloudAssistantStatus,omitempty" xml:"CloudAssistantStatus,omitempty"`
	// The version number of Cloud Assistant Agent. This parameter is empty if Cloud Assistant Agent is not installed or is not running on the instance.
	//
	// example:
	//
	// 2.2.0.106
	CloudAssistantVersion *string `json:"CloudAssistantVersion,omitempty" xml:"CloudAssistantVersion,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-bp1iudwa5b1tqa****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of tasks that Cloud Assistant completed on the instance.
	//
	// example:
	//
	// 2
	InvocationCount *int64 `json:"InvocationCount,omitempty" xml:"InvocationCount,omitempty"`
	// The last heartbeat time of Cloud Assistant. The value is updated every minute on average. The interval can be 55, 60, or 65 seconds.
	//
	// example:
	//
	// 2021-03-15T09:00:00Z
	LastHeartbeatTime *string `json:"LastHeartbeatTime,omitempty" xml:"LastHeartbeatTime,omitempty"`
	// The time when commands were last run.
	//
	// example:
	//
	// 2021-03-15T08:00:00Z
	LastInvokedTime *string `json:"LastInvokedTime,omitempty" xml:"LastInvokedTime,omitempty"`
	// The operating system type of the instance. Valid values:
	//
	// 	- Windows
	//
	// 	- Linux
	//
	// 	- FreeBSD
	//
	// example:
	//
	// Linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// Indicates whether Cloud Assistant supports Session Manager on the instance. If Session Manager is not supported, the version of Cloud Assistant Agent is outdated. Update Cloud Assistant Agent to the latest version.
	//
	// To support Session Manager, the version of Cloud Assistant Agent cannot be earlier than the following versions:
	//
	// 	- Linux: 2.2.3.189
	//
	// 	- Windows: 2.1.3.189
	//
	// example:
	//
	// true
	SupportSessionManager *bool `json:"SupportSessionManager,omitempty" xml:"SupportSessionManager,omitempty"`
}

func (s DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) GetActiveTaskCount() *int64 {
	return s.ActiveTaskCount
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) GetCloudAssistantStatus() *string {
	return s.CloudAssistantStatus
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) GetCloudAssistantVersion() *string {
	return s.CloudAssistantVersion
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) GetInvocationCount() *int64 {
	return s.InvocationCount
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) GetLastHeartbeatTime() *string {
	return s.LastHeartbeatTime
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) GetLastInvokedTime() *string {
	return s.LastInvokedTime
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) GetOSType() *string {
	return s.OSType
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) GetSupportSessionManager() *bool {
	return s.SupportSessionManager
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) SetActiveTaskCount(v int64) *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus {
	s.ActiveTaskCount = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) SetCloudAssistantStatus(v string) *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus {
	s.CloudAssistantStatus = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) SetCloudAssistantVersion(v string) *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus {
	s.CloudAssistantVersion = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) SetInstanceId(v string) *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) SetInvocationCount(v int64) *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus {
	s.InvocationCount = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) SetLastHeartbeatTime(v string) *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus {
	s.LastHeartbeatTime = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) SetLastInvokedTime(v string) *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus {
	s.LastInvokedTime = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) SetOSType(v string) *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus {
	s.OSType = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) SetSupportSessionManager(v bool) *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus {
	s.SupportSessionManager = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBodyInstanceCloudAssistantStatusSetInstanceCloudAssistantStatus) Validate() error {
	return dara.Validate(s)
}
