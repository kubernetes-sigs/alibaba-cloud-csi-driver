// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataFlowSubTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DescribeDataFlowSubTasksRequest
	GetFileSystemId() *string
	SetFilters(v []*DescribeDataFlowSubTasksRequestFilters) *DescribeDataFlowSubTasksRequest
	GetFilters() []*DescribeDataFlowSubTasksRequestFilters
	SetMaxResults(v int64) *DescribeDataFlowSubTasksRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeDataFlowSubTasksRequest
	GetNextToken() *string
}

type DescribeDataFlowSubTasksRequest struct {
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-370lx1ev9ss27o0****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The filter that is used to query data streaming tasks.
	//
	// if can be null:
	// false
	Filters []*DescribeDataFlowSubTasksRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The number of results for each query.
	//
	// 	- Valid values: 20 to 100.
	//
	// 	- Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// iWk0AQAAAAAvY2FzZS8=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeDataFlowSubTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowSubTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowSubTasksRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeDataFlowSubTasksRequest) GetFilters() []*DescribeDataFlowSubTasksRequestFilters {
	return s.Filters
}

func (s *DescribeDataFlowSubTasksRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeDataFlowSubTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDataFlowSubTasksRequest) SetFileSystemId(v string) *DescribeDataFlowSubTasksRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeDataFlowSubTasksRequest) SetFilters(v []*DescribeDataFlowSubTasksRequestFilters) *DescribeDataFlowSubTasksRequest {
	s.Filters = v
	return s
}

func (s *DescribeDataFlowSubTasksRequest) SetMaxResults(v int64) *DescribeDataFlowSubTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDataFlowSubTasksRequest) SetNextToken(v string) *DescribeDataFlowSubTasksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDataFlowSubTasksRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeDataFlowSubTasksRequestFilters struct {
	// The filter name.
	//
	// Valid values:
	//
	// 	- DataFlowIds: filters data flow subtasks by data flow ID.
	//
	// 	- DataFlowTaskIds: filters data flow subtasks by data flow task ID.
	//
	// 	- DataFlowSubTaskIds: filters data flow subtasks by data streaming task ID.
	//
	// 	- Status: filters data flow subtasks by status.
	//
	// 	- SrcFilePath: filters data flow subtasks by source file path.
	//
	// 	- DstFilePath: filters data flow subtasks by destination file path.
	//
	// example:
	//
	// DataFlowSubTaskIds
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value. This parameter does not support wildcards.
	//
	// 	- If Key is set to DataFlowIds, set Value to a data flow ID or a part of the data flow ID. You can specify a data flow ID or a group of data flow IDs. You can specify a maximum of 10 data flow IDs. Example: `df-194433a5be31****` or `df-194433a5be31****,df-244433a5be31****`.
	//
	// 	- If Key is set to DataFlowTaskIds, set Value to a data flow task ID or a part of the data flow task ID. You can specify a data flow task ID or a group of data flow task IDs. You can specify a maximum of 10 data flow task IDs. Example:  `task-38aa8e890f45****` or `task-38aa8e890f45****,task-27aa8e890f45****`.
	//
	// 	- If Key is set to DataFlowSubTaskIds, set Value to a data streaming task ID or a part of the data streaming task ID. You can specify a data streaming task ID or a group of data streaming task IDs. You can specify a maximum of 10 data streaming task IDs. Example: ` subTaskId-370kyfmyknxcyzw***	- `or `subTaskId-370kyfmyknxcyzw****,subTaskId-280kyfmyknxcyzw****`.
	//
	// 	- If Key is set to Status, set Value to the status of the data flow task. The status can be EXPIRED, CREATED, RUNNING, COMPLETE, CANCELING, FAILED, or CANCELED. Combined query is supported.
	//
	// 	- If Key is set to SrcFilePath, set Value to the path of the source file. The path can be up to 1,023 characters in length.
	//
	// 	- If Key is set to DstFilePath, set Value to the path of the destination file. The path can be up to 1,023 characters in length.
	//
	// example:
	//
	// subTaskId-370kyfmyknxcyzw****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDataFlowSubTasksRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowSubTasksRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowSubTasksRequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeDataFlowSubTasksRequestFilters) GetValue() *string {
	return s.Value
}

func (s *DescribeDataFlowSubTasksRequestFilters) SetKey(v string) *DescribeDataFlowSubTasksRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeDataFlowSubTasksRequestFilters) SetValue(v string) *DescribeDataFlowSubTasksRequestFilters {
	s.Value = &v
	return s
}

func (s *DescribeDataFlowSubTasksRequestFilters) Validate() error {
	return dara.Validate(s)
}
