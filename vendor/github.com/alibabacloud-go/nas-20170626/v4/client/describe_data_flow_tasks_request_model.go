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
	// - CPFS General-purpose: The ID must start with `cpfs-`, such as cpfs-099394bd928c\\*\\*\\*\\*.
	//
	// - CPFS for AI Computing: The ID must start with `bmcpfs-`, such as bmcpfs-290w65p03ok64ya\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64ya****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// A collection of filters.
	//
	// if can be null:
	// false
	Filters []*DescribeDataFlowTasksRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The maximum number of results to return per page.
	//
	// Valid values: 10 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page of results. If the response is truncated, use this token in your next request to retrieve the subsequent page.
	//
	// example:
	//
	// TGlzdFJlc291cmNlU****mVzJjE1MTI2NjY4NzY5MTAzOTEmMiZORnI4NDhVeEtrUT0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Specifies whether to return report information.
	//
	// - True (default): Includes reports in the response.
	//
	// - False: Excludes reports from the response.
	//
	// > 	- Set this parameter to False to speed up the query.
	//
	// >
	//
	// > 	- This parameter is supported only in CPFS for AI Computing.
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
	// The filter key.
	//
	// Valid values:
	//
	// - DataFlowIds: Filters by data flow ID.
	//
	// - TaskIds: Filters by data flow task ID.
	//
	// - Originator: Filters by originator.
	//
	// - TaskActions: Filters by data flow task type.
	//
	// - DataTypes: Filters by data type.
	//
	// - Status: Filters by status.
	//
	// - CreateTimeBegin: Filters data flow tasks created after the specified time.
	//
	// - CreateTimeEnd: Filters data flow tasks created before the specified time.
	//
	// - StartTimeBegin: Filters data flow tasks that started after the specified time.
	//
	// - StartTimeEnd: Filters data flow tasks that started before the specified time.
	//
	// - EndTimeBegin: Filters data flow tasks that ended after the specified time.
	//
	// - EndTimeEnd: Filters data flow tasks that ended before the specified time.
	//
	// example:
	//
	// DataFlowIds
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value. This parameter does not support wildcards.
	//
	// - When `Key` is `DataFlowIds`, specify one or more data flow IDs. You can specify up to 10 data flow IDs, separated by commas. For example, `df-194433a5be31****` or `df-194433a512a2****,df-234533a5be31****`.
	//
	// - When `Key` is `TaskId`, specify one or more data flow task IDs. You can specify up to 10 data flow task IDs, separated by commas. For example, `task-38aa8e890f45****` or `task-38aa8e890f45****,task-29ae8e890f45****`.
	//
	// - When `Key` is `TaskActions`, specify the data flow task type. Valid values are **Import**, **Export**, **Evict**, **Inventory**, **StreamImport**, and **StreamExport**. You can specify multiple values. CPFS for AI Computing supports only Import, Export, StreamImport, and StreamExport. StreamImport and StreamExport are available only in CPFS for AI Computing 2.6.0 and later.
	//
	// - When `Key` is `DataTypes`, specify the data type of the data flow task. Valid values are MetaAndData, Metadata, and Data. You can specify multiple values.
	//
	// - When `Key` is `Originator`, specify the originator of the data flow task. Valid values are User and System.
	//
	// - When `Key` is `Status`, specify the status of the data flow task. Valid values are Pending, Executing, Failed, Completed, Canceling, and Canceled. You can specify multiple values.
	//
	// - When `Key` is `CreateTimeBegin`, specify the earliest creation time. Use the `yyyy-MM-ddTHH:mmZ` format.
	//
	// - When `Key` is `CreateTimeEnd`, specify the latest creation time. Use the `yyyy-MM-ddTHH:mmZ` format.
	//
	// - When `Key` is `StartTimeBegin`, specify the earliest start time. Use the `yyyy-MM-ddTHH:mmZ` format.
	//
	// - When `Key` is `StartTimeEnd`, specify the latest start time. Use the `yyyy-MM-ddTHH:mmZ` format.
	//
	// - When `Key` is `EndTimeBegin`, specify the earliest end time. Use the `yyyy-MM-ddTHH:mmZ` format.
	//
	// - When `Key` is `EndTimeEnd`, specify the latest end time. Use the `yyyy-MM-ddTHH:mmZ` format.
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
