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
	// The file system ID.
	//
	// - General-purpose CPFS: must start with `cpfs-`, such as cpfs-099394bd928c****.
	//
	// - CPFS for Lingjun: must start with `bmcpfs-`, such as bmcpfs-290w65p03ok64ya****.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64ya****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The filter conditions.
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
	// The pagination token that is used in the next request to retrieve a new page of results. If the return results are truncated, you can use NextToken to initiate a new request to retrieve the content after the truncation point.
	//
	// example:
	//
	// TGlzdFJlc291cmNlU****mVzJjE1MTI2NjY4NzY5MTAzOTEmMiZORnI4NDhVeEtrUT0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Specifies whether to query report information.
	//
	// - True (default): queries reports.
	//
	// - False: does not query reports.
	//
	// >- Setting this parameter to False can speed up queries.
	//
	// > - Only CPFS for Lingjun is supported.
	//
	// example:
	//
	// True
	WithReports *bool `json:"WithReports,omitempty" xml:"WithReports,omitempty"`
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
	// The name of the filter key.
	//
	// Valid values:
	//
	// - DataFlowIds: filters by data flow ID.
	//
	// - TaskIds: filters by data flow task ID.
	//
	// - Originator: filters by the initiator of the data flow task.
	//
	// - TaskActions: filters by the type of the data flow task.
	//
	// - DataTypes: filters by the data type of the data flow task.
	//
	// - Status: filters by data flow status.
	//
	// - CreateTimeBegin: filters data flow tasks created after the specified time.
	//
	// - CreateTimeEnd: filters data flow tasks created before the specified time.
	//
	// - StartTimeBegin: filters data flow tasks started after the specified time.
	//
	// - StartTimeEnd: filters data flow tasks started before the specified time.
	//
	// - EndTimeBegin: filters data flow tasks ended after the specified time.
	//
	// - EndTimeEnd: filters data flow tasks ended before the specified time.
	//
	// example:
	//
	// DataFlowIds
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter key. Wildcards are not supported.
	//
	// - If Key is set to DataFlowIds, Value is set to a data flow ID or part of a data flow ID. You can specify one or more data flow IDs. A maximum of 10 data flow IDs can be specified. Example: `df-194433a5be31****` or `df-194433a512a2****,df-234533a5be31****`.
	//
	// - If Key is set to TaskId, Value is set to a data flow task ID or part of a data flow task ID. You can specify one or more data flow task IDs. A maximum of 10 data flow task IDs can be specified. Example: `task-38aa8e890f45****` or `task-38aa8e890f45****,task-29ae8e890f45****`.
	//
	// - If Key is set to TaskActions, Value is set to the type of the data flow task, including **Import**, **Export**, **Evict**, **Inventory**, **StreamImport**, and **StreamExport**. Combined queries are supported. CPFS for Lingjun supports only Import, Export, StreamImport, and StreamExport. StreamImport and StreamExport are supported only by CPFS for Lingjun 2.6.0 and later.
	//
	// - If Key is set to DataTypes, Value is set to the data type of the data flow task, including MetaAndData, Metadata, and Data. Combined queries are supported.
	//
	// - If Key is set to Originator, Value is set to the initiator of the data flow task, including User and System.
	//
	// - If Key is set to Status, Value is set to the status of the data flow task, including Pending, Executing, Failed, Completed, Canceling, and Canceled. Combined queries are supported.
	//
	// - If Key is set to CreateTimeBegin, Value is set to the earliest creation time of data flow tasks. Format: `yyyy-MM-ddThh:mmZ`.
	//
	// - If Key is set to CreateTimeEnd, Value is set to the latest creation time of data flow tasks. Format: `yyyy-MM-ddThh:mmZ`.
	//
	// - If Key is set to StartTimeBegin, Value is set to the earliest start time of data flow tasks. Format: `yyyy-MM-ddThh:mmZ`.
	//
	// - If Key is set to StartTimeEnd, Value is set to the latest start time of data flow tasks. Format: `yyyy-MM-ddThh:mmZ`.
	//
	// - If Key is set to EndTimeBegin, Value is set to the earliest end time of data flow tasks. Format: `yyyy-MM-ddThh:mmZ`.
	//
	// - If Key is set to EndTimeEnd, Value is set to the latest end time of data flow tasks. Format: `yyyy-MM-ddThh:mmZ`.
	//
	// example:
	//
	// df-194433a5be31****
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
