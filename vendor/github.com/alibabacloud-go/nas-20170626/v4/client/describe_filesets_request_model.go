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
	// The ID of the file system.
	//
	// 	- The IDs of CPFS file systems must start with `cpfs-`. Example: cpfs-099394bd928c\\*\\*\\*\\*.
	//
	// 	- The IDs of CPFS for LINGJUN file systems must start with `bmcpfs-`. Example: bmcpfs-290w65p03ok64ya\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64ya****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The filter that is used to query filesets.
	Filters []*DescribeFilesetsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
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
	// The condition by which the results are sorted. Valid values:
	//
	// 	- FileCountLimit: the file quantity quota
	//
	// 	- SizeLimit: the capacity quota
	//
	// 	- FileCountUsage: the usage of the file quantity quota
	//
	// 	- SpaceUsage: the capacity usage
	//
	// example:
	//
	// FileCountLimit
	OrderByField *string `json:"OrderByField,omitempty" xml:"OrderByField,omitempty"`
	// The order in which you want to sort the results. Valid values:
	//
	// 	- asc (default): ascending order
	//
	// 	- desc: descending order
	//
	// >  This parameter takes effect only if you specify the OrderByField parameter.
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
	// The filter name. Valid values:
	//
	// 	- FsetIds: filters filesets by fileset ID.
	//
	// 	- FileSystemPath: filters filesets based on the path of a fileset in a CPFS file system.
	//
	// 	- Description: filters filesets based on the fileset description.
	//
	// 	- QuotaExists: filters filesets based on whether quotas exist.
	//
	// >  Only CPFS for LINGJUN V2.7.0 and later support the QuotaExists parameter.
	//
	// example:
	//
	// FsetIds
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value. This parameter does not support wildcards.
	//
	// 	- If Key is set to FsetIds, set Value to a fileset ID or a part of the fileset ID. You can specify a fileset ID or a group of fileset IDs. You can specify a maximum of 10 fileset IDs. Example: `fset-1902718ea0ae****` or `fset-1902718ea0ae****,fset-3212718ea0ae****`.
	//
	// 	- If Key is set to FileSystemPath, set Value to the path or a part of the path of a fileset in a CPFS file system. The value must be 2 to 1024 characters in length. The value must be encoded in UTF-8.
	//
	// 	- If Key is set to Description, set Value to a fileset description or a part of the fileset description.
	//
	// 	- If Key is set to QuotaExists, set Value to true or false. If you do not specify the parameter, all filesets are returned.
	//
	// example:
	//
	// fset-1902718ea0ae****,fset-3212718ea0ae****
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
