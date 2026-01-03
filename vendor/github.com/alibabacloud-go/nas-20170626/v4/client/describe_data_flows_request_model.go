// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataFlowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DescribeDataFlowsRequest
	GetFileSystemId() *string
	SetFilters(v []*DescribeDataFlowsRequestFilters) *DescribeDataFlowsRequest
	GetFilters() []*DescribeDataFlowsRequestFilters
	SetMaxResults(v int64) *DescribeDataFlowsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeDataFlowsRequest
	GetNextToken() *string
}

type DescribeDataFlowsRequest struct {
	// The ID of the file system.
	//
	// 	- The IDs of CPFS file systems must start with `cpfs-`. Example: cpfs-125487\\*\\*\\*\\*.
	//
	// 	- The IDs of CPFS for Lingjun file systems must start with `bmcpfs-`. Example: bmcpfs-0015\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64ya****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The filter that is used to query dataflows.
	//
	// if can be null:
	// true
	Filters []*DescribeDataFlowsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The number of results for each query.
	//
	// Valid values: 10 to 100. Default value: 20.
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
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeDataFlowsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeDataFlowsRequest) GetFilters() []*DescribeDataFlowsRequestFilters {
	return s.Filters
}

func (s *DescribeDataFlowsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeDataFlowsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDataFlowsRequest) SetFileSystemId(v string) *DescribeDataFlowsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeDataFlowsRequest) SetFilters(v []*DescribeDataFlowsRequestFilters) *DescribeDataFlowsRequest {
	s.Filters = v
	return s
}

func (s *DescribeDataFlowsRequest) SetMaxResults(v int64) *DescribeDataFlowsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDataFlowsRequest) SetNextToken(v string) *DescribeDataFlowsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDataFlowsRequest) Validate() error {
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

type DescribeDataFlowsRequestFilters struct {
	// The filter name. Valid value:
	//
	// 	- DataFlowIds: filters dataflow tasks by dataflow ID.
	//
	// 	- FsetIds: filters dataflows by fileset ID.
	//
	// 	- FileSystemPath: filters dataflows based on the path of a fileset in a CPFS file system.
	//
	// 	- SourceStorage: filters dataflows based on the access path of the source storage.
	//
	// 	- ThroughputList: filters dataflows based on dataflow throughput.
	//
	// 	- Description: filters dataflows based on the fileset description.
	//
	// 	- Status: filters dataflows based on dataflow status.
	//
	// example:
	//
	// FsetIds
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter. This parameter does not support wildcards.
	//
	// 	- If Key is set to DataFlowIds, set Value to a dataflow ID or a part of the dataflow ID. You can specify a dataflow ID or a group of dataflow IDs. You can specify a maximum of 10 dataflow IDs. Example: `df-194433a5be31****` or `df-194433a512a2****,df-234533a5be31****`.
	//
	// 	- If Key is set to FsetIds, set Value to a fileset ID or a part of the fileset ID. You can specify a fileset ID or a group of fileset IDs. You can specify a maximum of 10 fileset IDs. For example, `fset-1902718ea0ae****` or `fset-235718ea0ae****,fset-5122718ea0ae****`.
	//
	// 	- If Key is set to FileSystemPath, set Value to the path or a part of the path of a fileset in a CPFS file system. The value of the parameter must be 1 to 1,024 characters in length.
	//
	// 	- If Key is set to SourceStorage, set Value to the access path or a part of the access path of the source storage. The path can be up to 1,024 characters in length.
	//
	// 	- If Key is set to ThroughputList, set Value to the dataflow throughput. Combined query is supported.
	//
	// 	- If Key is set to Description, set Value to a dataflow description or a part of the dataflow description.
	//
	// 	- If Key is set to Status, set Value to the dataflow status.
	//
	// 	- If Key is set to SourceStoragePath, set Value to the access path or a part of the access path of the source storage. The path can be up to 1,024 characters in length.
	//
	// example:
	//
	// FsetIds
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDataFlowsRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataFlowsRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeDataFlowsRequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeDataFlowsRequestFilters) GetValue() *string {
	return s.Value
}

func (s *DescribeDataFlowsRequestFilters) SetKey(v string) *DescribeDataFlowsRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeDataFlowsRequestFilters) SetValue(v string) *DescribeDataFlowsRequestFilters {
	s.Value = &v
	return s
}

func (s *DescribeDataFlowsRequestFilters) Validate() error {
	return dara.Validate(s)
}
