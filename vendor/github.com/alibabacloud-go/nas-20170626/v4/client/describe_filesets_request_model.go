// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFilesetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DescribeFilesetsRequest
	GetFileSystemId() *string
	SetFilters(v []*DescribeFilesetsRequestFilters) *DescribeFilesetsRequest
	GetFilters() []*DescribeFilesetsRequestFilters
	SetMaxResults(v int64) *DescribeFilesetsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeFilesetsRequest
	GetNextToken() *string
	SetOrderByField(v string) *DescribeFilesetsRequest
	GetOrderByField() *string
	SetSortOrder(v string) *DescribeFilesetsRequest
	GetSortOrder() *string
}

type DescribeFilesetsRequest struct {
	// The file system ID.
	//
	// - CPFS: The ID must start with `cpfs-`, such as cpfs-099394bd928c****.
	//
	// - CPFS for Lingjun: The ID must start with `bmcpfs-`, such as bmcpfs-290w65p03ok64ya****.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64ya****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The filter key information for the filesets to query.
	Filters []*DescribeFilesetsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The number of results per query.
	//
	// Valid values: 10 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If the response is truncated, you can use this token in the next request to retrieve the remaining results.
	//
	// example:
	//
	// TGlzdFJlc291cmNlU****mVzJjE1MTI2NjY4NzY5MTAzOTEmMiZORnI4NDhVeEtrUT0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The field by which to sort the results.
	//
	// - FileCountLimit: the quota file count limit.
	//
	// - SizeLimit: the quota capacity limit.
	//
	// - FileCountUsage: the file count usage.
	//
	// - SpaceUsage: the capacity usage.
	//
	// example:
	//
	// FileCountLimit
	OrderByField *string `json:"OrderByField,omitempty" xml:"OrderByField,omitempty"`
	// The sort order. Valid values:
	//
	// - asc (default): ascending order, which sorts results from smallest to largest.
	//
	// - desc: descending order, which sorts results from largest to smallest.
	//
	// > This parameter takes effect only when the OrderByField parameter is specified.
	//
	// example:
	//
	// asc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s DescribeFilesetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFilesetsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeFilesetsRequest) GetFilters() []*DescribeFilesetsRequestFilters {
	return s.Filters
}

func (s *DescribeFilesetsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeFilesetsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeFilesetsRequest) GetOrderByField() *string {
	return s.OrderByField
}

func (s *DescribeFilesetsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *DescribeFilesetsRequest) SetFileSystemId(v string) *DescribeFilesetsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFilesetsRequest) SetFilters(v []*DescribeFilesetsRequestFilters) *DescribeFilesetsRequest {
	s.Filters = v
	return s
}

func (s *DescribeFilesetsRequest) SetMaxResults(v int64) *DescribeFilesetsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeFilesetsRequest) SetNextToken(v string) *DescribeFilesetsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeFilesetsRequest) SetOrderByField(v string) *DescribeFilesetsRequest {
	s.OrderByField = &v
	return s
}

func (s *DescribeFilesetsRequest) SetSortOrder(v string) *DescribeFilesetsRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeFilesetsRequest) Validate() error {
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

type DescribeFilesetsRequestFilters struct {
	// The name of the filter key. Valid values:
	//
	// - FsetIds: filters by fileset ID.
	//
	// - FileSystemPath: filters by the path of the fileset in the CPFS file system.
	//
	// - Description: filters by the description of the fileset.
	//
	// - QuotaExists: filters by whether a quota exists.
	//
	// > Only CPFS for Lingjun 2.7.0 and later support filtering by the QuotaExists parameter.
	//
	// example:
	//
	// FsetIds
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter key. Wildcards are not supported for this parameter.
	//
	// - If Key is set to FsetIds, Value is set to a fileset ID. You can specify one or more fileset IDs, up to 10. Separate multiple values with commas (,). Example: `fset-1902718ea0ae****` or `fset-1902718ea0ae****,fset-3212718ea0ae****`.
	//
	// - If Key is set to FileSystemPath, Value is set to the full path or a partial path of the fileset in the CPFS file system. The value must be 2 to 1,024 characters in length and encoded in UTF-8.
	//
	// - If Key is set to Description, Value is set to the full description or a partial description of the fileset.
	//
	// - If Key is set to QuotaExists, Value is set to true or false. If you leave this parameter empty, all filesets are returned.
	//
	// example:
	//
	// fset-1902718ea0ae****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeFilesetsRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesetsRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeFilesetsRequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeFilesetsRequestFilters) GetValue() *string {
	return s.Value
}

func (s *DescribeFilesetsRequestFilters) SetKey(v string) *DescribeFilesetsRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeFilesetsRequestFilters) SetValue(v string) *DescribeFilesetsRequestFilters {
	s.Value = &v
	return s
}

func (s *DescribeFilesetsRequestFilters) Validate() error {
	return dara.Validate(s)
}
