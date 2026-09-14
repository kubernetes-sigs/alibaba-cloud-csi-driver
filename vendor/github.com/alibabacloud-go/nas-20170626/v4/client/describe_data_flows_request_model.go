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
	// The file system ID.
	//
	// - CPFS: must start with `cpfs-`, such as cpfs-125487\\*\\*\\*\\*.
	//
	// - CPFS for Lingjun: must start with `bmcpfs-`, such as bmcpfs-0015\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The filter keys for querying data flows.
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
	// The pagination token that is used in the next request to retrieve a new page of results. If the return results are truncated, use NextToken to obtain content starting from the truncation point.
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
	// The name of the filter key. Valid values:
	//
	// - DataFlowIds: filters by data flow ID.
	//
	// - FsetIds: filters by Fileset ID.
	//
	// - FileSystemPath: filters by the path of the Fileset in the CPFS file system.
	//
	// - SourceStorage: filters by the access path of the source storage.
	//
	// - ThroughputList: filters by the transmission bandwidth of the data flow.
	//
	// - Description: filters by the description of the Fileset.
	//
	// - Status: filters by data flow status.
	//
	// example:
	//
	// FsetIds
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter key. Wildcards are not supported for this parameter.
	//
	// - If Key is set to DataFlowIds, Value is set to a data flow ID or part of a data flow ID. You can specify one or more data flow IDs. A maximum of 10 data flow IDs can be specified. Example: `df-194433a5be31****` or `df-194433a512a2****,df-234533a5be31****`.
	//
	// - If Key is set to FsetIds, Value is set to a Fileset ID or part of a Fileset ID. You can specify one or more Fileset IDs. A maximum of 10 Fileset IDs can be specified. Example: `fset-1902718ea0ae****` or `fset-235718ea0ae****,fset-5122718ea0ae****`.
	//
	// - If Key is set to FileSystemPath, Value is set to a path or part of a path in the CPFS file system. The value must be 1 to 1024 characters in length.
	//
	// - If Key is set to SourceStorage, Value is set to the access path of the source storage. The maximum length is 1024 characters.
	//
	// - If Key is set to ThroughputList, Value is set to the transmission bandwidth of the data flow. Combined queries are supported.
	//
	// - If Key is set to Description, Value is set to the description or part of the description of the data flow.
	//
	// - If Key is set to Status, Value is set to the data flow status.
	//
	// - If Key is set to SourceStoragePath, Value is set to the access path of the source storage or part of the access path. The maximum length is 1024 characters.
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
