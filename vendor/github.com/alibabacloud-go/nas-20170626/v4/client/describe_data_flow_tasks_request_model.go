// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataFlowTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DescribeDataFlowTasksRequest
	GetFileSystemId() *string
	SetFilters(v []*DescribeDataFlowTasksRequestFilters) *DescribeDataFlowTasksRequest
	GetFilters() []*DescribeDataFlowTasksRequestFilters
	SetMaxResults(v int64) *DescribeDataFlowTasksRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeDataFlowTasksRequest
	GetNextToken() *string
	SetWithReports(v bool) *DescribeDataFlowTasksRequest
	GetWithReports() *bool
}

type DescribeDataFlowTasksRequest struct {
	// The ID of the file system.
	//
	// 	- The IDs of CPFS file systems must start with `cpfs-`. Example: cpfs-099394bd928c\\*\\*\\*\\*.
	//
	// 	- The IDs of CPFS for LINGJUN file systems must start with `bmcpfs-`. Example: bmcpfs-290w65p03ok64ya\\*\\*\\*\\*.
	//
	// >  CPFS is not supported on the international site.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The details about filters.
	//
	// if can be null:
	// false
	Filters []*DescribeDataFlowTasksRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The number of results for each query.
	//
	// Valid values: 10 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// TGlzdFJlc291cmNlU****mVzJjE1MTI2NjY4NzY5MTAzOTEmMiZORnI4NDhVeEtrUT0=
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	WithReports *bool   `json:"WithReports,omitempty" xml:"WithReports,omitempty"`
}

func (s DescribeDataFlowTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeDataFlowTasksRequest) GetFilters() []*DescribeDataFlowTasksRequestFilters {
	return s.Filters
}

func (s *DescribeDataFlowTasksRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeDataFlowTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDataFlowTasksRequest) GetWithReports() *bool {
	return s.WithReports
}

func (s *DescribeDataFlowTasksRequest) SetFileSystemId(v string) *DescribeDataFlowTasksRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeDataFlowTasksRequest) SetFilters(v []*DescribeDataFlowTasksRequestFilters) *DescribeDataFlowTasksRequest {
	s.Filters = v
	return s
}

func (s *DescribeDataFlowTasksRequest) SetMaxResults(v int64) *DescribeDataFlowTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDataFlowTasksRequest) SetNextToken(v string) *DescribeDataFlowTasksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDataFlowTasksRequest) SetWithReports(v bool) *DescribeDataFlowTasksRequest {
	s.WithReports = &v
	return s
}

func (s *DescribeDataFlowTasksRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataFlowTasksRequestFilters struct {
	// The filter name.
	//
	// Valid values:
	//
	// 	- DataFlowIds: filters data flow tasks by data flow ID.
	//
	// 	- TaskIds: filters data flow tasks by task ID.
	//
	// 	- Originator: filters data flow tasks by task initiator.
	//
	// 	- TaskActions: filters data flow tasks by task type.
	//
	// 	- DataTypes: filters data flow tasks by data type.
	//
	// 	- Status: filters data flow tasks by data flow status.
	//
	// 	- CreateTimeBegin: filters data flow tasks that are created after a specified time.
	//
	// 	- CreateTimeEnd: filters data flow tasks that are created before a specified time.
	//
	// 	- StartTimeBegin: filters data flow tasks that are started after a specified time.
	//
	// 	- StartTimeEnd: filters data flow tasks that are started before a specified time.
	//
	// 	- EndTimeBegin: filters data flow tasks that are stopped after a specified time.
	//
	// 	- EndTimeEnd: filters data flow tasks that are stopped before a specified time.
	//
	// example:
	//
	// DataFlowIds
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value. This parameter does not support wildcards.
	//
	// 	- If Key is set to DataFlowIds, set Value to a data flow ID or a part of the data flow ID. You can specify a data flow ID or a group of data flow IDs. You can specify a maximum of 10 data flow IDs. Example: `df-194433a5be31****` or `df-194433a512a2****,df-234533a5be31****`.
	//
	// 	- If Key is set to TaskId, set Value to a data flow task ID or a part of the data flow task ID. You can specify a data flow task ID or a group of data flow task IDs. You can specify a maximum of 10 data flow task IDs. Example: `task-38aa8e890f45****` or `task-38aa8e890f45****,task-29ae8e890f45****`.
	//
	// 	- If Key is set to TaskActions, set Value to the type of data flow task. The task type can be **Import**, **Export**, **Evict**, **Inventory**, **StreamImport**, or **StreamExport**. Combined query is supported. CPFS for LINGJUN supports only the Import, Export, StreamImport, and StreamExport tasks. Only CPFS for LINGJUN V2.6.0 and later support the StreamImport and StreamExport tasks.
	//
	// 	- If Key is set to DataTypes, set Value to the data type of the data flow task. The data type can be MetaAndData, Metadata, or Data. Combined query is supported.
	//
	// 	- If Key is set to Originator, set Value to the initiator of the data flow task. The initiator can be User or System.
	//
	// 	- If Key is set to Status, set Value to the status of the data flow task. The status can be Pending, Executing, Failed, Completed, Canceling, or Canceled. Combined query is supported.
	//
	// 	- If Key is set to CreateTimeBegin, set Value to the beginning of the time range to create the data flow task. Time format: `yyyy-MM-ddThh:mmZ`.
	//
	// 	- If Key is set to CreateTimeEnd, set Value to the end of the time range to create the data flow task. Time format: `yyyy-MM-ddThh:mmZ`.
	//
	// 	- If Key is set to StartTimeBegin, set Value to the beginning of the time range to start the data flow task. Time format: `yyyy-MM-ddThh:mmZ`.
	//
	// 	- If Key is set to StartTimeEnd, set Value to the end of the time range to start the data flow task. Time format: `yyyy-MM-ddThh:mmZ`.
	//
	// 	- If Key is set to EndTimeBegin, set Value to the beginning of the time range to stop the data flow task. Time format: `yyyy-MM-ddThh:mmZ`.
	//
	// 	- If Key is set to EndTimeEnd, set Value to the end of the time range to stop the data flow task. Time format: `yyyy-MM-ddThh:mmZ`.
	//
	// example:
	//
	// dfid-12345678
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDataFlowTasksRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowTasksRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowTasksRequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeDataFlowTasksRequestFilters) GetValue() *string {
	return s.Value
}

func (s *DescribeDataFlowTasksRequestFilters) SetKey(v string) *DescribeDataFlowTasksRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeDataFlowTasksRequestFilters) SetValue(v string) *DescribeDataFlowTasksRequestFilters {
	s.Value = &v
	return s
}

func (s *DescribeDataFlowTasksRequestFilters) Validate() error {
	return dara.Validate(s)
}
