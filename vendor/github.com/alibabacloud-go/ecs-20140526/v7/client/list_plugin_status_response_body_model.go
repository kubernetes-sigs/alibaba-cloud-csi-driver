// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstancePluginStatusSet(v *ListPluginStatusResponseBodyInstancePluginStatusSet) *ListPluginStatusResponseBody
	GetInstancePluginStatusSet() *ListPluginStatusResponseBodyInstancePluginStatusSet
	SetNextToken(v string) *ListPluginStatusResponseBody
	GetNextToken() *string
	SetPageNumber(v int64) *ListPluginStatusResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListPluginStatusResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListPluginStatusResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListPluginStatusResponseBody
	GetTotalCount() *int64
}

type ListPluginStatusResponseBody struct {
	// The states of Cloud Assistant plug-ins on the instances.
	InstancePluginStatusSet *ListPluginStatusResponseBodyInstancePluginStatusSet `json:"InstancePluginStatusSet,omitempty" xml:"InstancePluginStatusSet,omitempty" type:"Struct"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
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
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPluginStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPluginStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListPluginStatusResponseBody) GetInstancePluginStatusSet() *ListPluginStatusResponseBodyInstancePluginStatusSet {
	return s.InstancePluginStatusSet
}

func (s *ListPluginStatusResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPluginStatusResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListPluginStatusResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPluginStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPluginStatusResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListPluginStatusResponseBody) SetInstancePluginStatusSet(v *ListPluginStatusResponseBodyInstancePluginStatusSet) *ListPluginStatusResponseBody {
	s.InstancePluginStatusSet = v
	return s
}

func (s *ListPluginStatusResponseBody) SetNextToken(v string) *ListPluginStatusResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPluginStatusResponseBody) SetPageNumber(v int64) *ListPluginStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPluginStatusResponseBody) SetPageSize(v int64) *ListPluginStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPluginStatusResponseBody) SetRequestId(v string) *ListPluginStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPluginStatusResponseBody) SetTotalCount(v int64) *ListPluginStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPluginStatusResponseBody) Validate() error {
	if s.InstancePluginStatusSet != nil {
		if err := s.InstancePluginStatusSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPluginStatusResponseBodyInstancePluginStatusSet struct {
	InstancePluginStatus []*ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatus `json:"InstancePluginStatus,omitempty" xml:"InstancePluginStatus,omitempty" type:"Repeated"`
}

func (s ListPluginStatusResponseBodyInstancePluginStatusSet) String() string {
	return dara.Prettify(s)
}

func (s ListPluginStatusResponseBodyInstancePluginStatusSet) GoString() string {
	return s.String()
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSet) GetInstancePluginStatus() []*ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatus {
	return s.InstancePluginStatus
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSet) SetInstancePluginStatus(v []*ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatus) *ListPluginStatusResponseBodyInstancePluginStatusSet {
	s.InstancePluginStatus = v
	return s
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSet) Validate() error {
	if s.InstancePluginStatus != nil {
		for _, item := range s.InstancePluginStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatus struct {
	// The ID of the instance.
	//
	// example:
	//
	// i-xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The queried Cloud Assistant plug-ins.
	PluginStatusSet *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSet `json:"PluginStatusSet,omitempty" xml:"PluginStatusSet,omitempty" type:"Struct"`
}

func (s ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatus) String() string {
	return dara.Prettify(s)
}

func (s ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatus) GoString() string {
	return s.String()
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatus) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatus) GetPluginStatusSet() *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSet {
	return s.PluginStatusSet
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatus) SetInstanceId(v string) *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatus {
	s.InstanceId = &v
	return s
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatus) SetPluginStatusSet(v *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSet) *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatus {
	s.PluginStatusSet = v
	return s
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatus) Validate() error {
	if s.PluginStatusSet != nil {
		if err := s.PluginStatusSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSet struct {
	PluginStatus []*ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus `json:"PluginStatus,omitempty" xml:"PluginStatus,omitempty" type:"Repeated"`
}

func (s ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSet) String() string {
	return dara.Prettify(s)
}

func (s ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSet) GoString() string {
	return s.String()
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSet) GetPluginStatus() []*ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus {
	return s.PluginStatus
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSet) SetPluginStatus(v []*ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus) *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSet {
	s.PluginStatus = v
	return s
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSet) Validate() error {
	if s.PluginStatus != nil {
		for _, item := range s.PluginStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus struct {
	// The first time when Cloud Assistant reported the state of the plug-in.
	//
	// example:
	//
	// 2020-01-19T09:15:46Z
	FirstHeartbeatTime *string `json:"FirstHeartbeatTime,omitempty" xml:"FirstHeartbeatTime,omitempty"`
	// The last time when Cloud Assistant reported the state of the plug-in.
	//
	// example:
	//
	// 2020-01-19T09:15:46Z
	LastHeartbeatTime *string `json:"LastHeartbeatTime,omitempty" xml:"LastHeartbeatTime,omitempty"`
	// The name of the plug-in.
	//
	// example:
	//
	// testName
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	// The state of the Cloud Assistant plug-in. Valid values:
	//
	// 	- NotInstalled: The plug-in is not installed.
	//
	// 	- Installed: The one-time plug-in is installed.
	//
	// 	- Running: The long-running plug-in is running.
	//
	// 	- Stopped: The long-running plug-in is not running.
	//
	// 	- Crashed: The plug-in is abnormal.
	//
	// 	- Removed: The plug-in is uninstalled.
	//
	// 	- Unknown: The state of the plug-in is unknown.
	//
	// example:
	//
	// Running
	PluginStatus *string `json:"PluginStatus,omitempty" xml:"PluginStatus,omitempty"`
	// The version of the plug-in.
	//
	// example:
	//
	// 1.1
	PluginVersion *string `json:"PluginVersion,omitempty" xml:"PluginVersion,omitempty"`
}

func (s ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus) String() string {
	return dara.Prettify(s)
}

func (s ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus) GoString() string {
	return s.String()
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus) GetFirstHeartbeatTime() *string {
	return s.FirstHeartbeatTime
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus) GetLastHeartbeatTime() *string {
	return s.LastHeartbeatTime
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus) GetPluginName() *string {
	return s.PluginName
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus) GetPluginStatus() *string {
	return s.PluginStatus
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus) GetPluginVersion() *string {
	return s.PluginVersion
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus) SetFirstHeartbeatTime(v string) *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus {
	s.FirstHeartbeatTime = &v
	return s
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus) SetLastHeartbeatTime(v string) *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus {
	s.LastHeartbeatTime = &v
	return s
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus) SetPluginName(v string) *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus {
	s.PluginName = &v
	return s
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus) SetPluginStatus(v string) *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus {
	s.PluginStatus = &v
	return s
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus) SetPluginVersion(v string) *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus {
	s.PluginVersion = &v
	return s
}

func (s *ListPluginStatusResponseBodyInstancePluginStatusSetInstancePluginStatusPluginStatusSetPluginStatus) Validate() error {
	return dara.Validate(s)
}
