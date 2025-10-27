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
	// 	- The IDs of CPFS for LINGJUN file systems must start with `bmcpfs-`. Example: bmcpfs-0015\\*\\*\\*\\*.
	//
	// >  CPFS file systems are available only on the China site (aliyun.com).
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64ya****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The filter that is used to query data flows.
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
	return dara.Validate(s)
}

type DescribeDataFlowsRequestFilters struct {
	// The filter name. Valid values:
	//
	// 	- DataFlowIds: filters data flows by data flow ID.
	//
	// 	- FsetIds: filters data flows by fileset ID.
	//
	// 	- FileSystemPath: filters data flows based on the path of a fileset in a CPFS file system.
	//
	// 	- SourceStorage: filters data flows based on the access path of the source storage.
	//
	// 	- ThroughputList: filters data flows based on data flow throughput.
	//
	// 	- Description: filters data flows based on the fileset description.
	//
	// 	- Status: filters data flows based on data flow status.
	//
	// example:
	//
	// FsetIds
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value. This parameter does not support wildcards.
	//
	// 	- If Key is set to DataFlowIds, set Value to a data flow ID or a part of the data flow ID. You can specify a data flow ID or a group of data flow IDs. You can specify a maximum of 10 data flow IDs. Example: `df-194433a5be31****` or `df-194433a5be31****,df-184433a5be31****`.
	//
	// 	- If Key is set to FsetIds, set Value to a fileset ID or a part of the fileset ID. You can specify a fileset ID or a group of fileset IDs. You can specify a maximum of 10 fileset IDs. Example: `fset-1902718ea0ae****` or `fset-1902718ea0ae****,fset-1242718ea0ae****`.
	//
	// 	- If Key is set to FileSystemPath, set Value to the path or a part of the path of a fileset in a CPFS file system. The value of the parameter must be 1 to 1,024 characters in length.
	//
	// 	- If Key is set to SourceStorage, set Value to the access path or a part of the access path of the source storage. The path can be up to 1,024 characters in length.
	//
	// 	- If Key is set to ThroughputList, set Value to the data flow throughput. Combined query is supported.
	//
	// 	- If Key is set to Description, set Value to a data flow description or a part of the data flow description.
	//
	// 	- If Key is set to Status, set Value to the data flow status.
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
