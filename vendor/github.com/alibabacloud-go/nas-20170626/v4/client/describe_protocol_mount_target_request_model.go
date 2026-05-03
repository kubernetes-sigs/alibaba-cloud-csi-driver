// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProtocolMountTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeProtocolMountTargetRequest
	GetClientToken() *string
	SetFileSystemId(v string) *DescribeProtocolMountTargetRequest
	GetFileSystemId() *string
	SetFilters(v []*DescribeProtocolMountTargetRequestFilters) *DescribeProtocolMountTargetRequest
	GetFilters() []*DescribeProtocolMountTargetRequestFilters
	SetMaxResults(v int64) *DescribeProtocolMountTargetRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeProtocolMountTargetRequest
	GetNextToken() *string
	SetProtocolServiceIds(v string) *DescribeProtocolMountTargetRequest
	GetProtocolServiceIds() *string
}

type DescribeProtocolMountTargetRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence?](https://help.aliyun.com/document_detail/25693.html)
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The filter that is used to query the export directories of the protocol service.
	Filters []*DescribeProtocolMountTargetRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The number of results for each query.
	//
	// 	- Value values: 10 to 100.
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
	// aBcdg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Protocol service ID list
	//
	// example:
	//
	// ptc-123xxx
	ProtocolServiceIds *string `json:"ProtocolServiceIds,omitempty" xml:"ProtocolServiceIds,omitempty"`
}

func (s DescribeProtocolMountTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProtocolMountTargetRequest) GoString() string {
	return s.String()
}

func (s *DescribeProtocolMountTargetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeProtocolMountTargetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeProtocolMountTargetRequest) GetFilters() []*DescribeProtocolMountTargetRequestFilters {
	return s.Filters
}

func (s *DescribeProtocolMountTargetRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeProtocolMountTargetRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeProtocolMountTargetRequest) GetProtocolServiceIds() *string {
	return s.ProtocolServiceIds
}

func (s *DescribeProtocolMountTargetRequest) SetClientToken(v string) *DescribeProtocolMountTargetRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeProtocolMountTargetRequest) SetFileSystemId(v string) *DescribeProtocolMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeProtocolMountTargetRequest) SetFilters(v []*DescribeProtocolMountTargetRequestFilters) *DescribeProtocolMountTargetRequest {
	s.Filters = v
	return s
}

func (s *DescribeProtocolMountTargetRequest) SetMaxResults(v int64) *DescribeProtocolMountTargetRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeProtocolMountTargetRequest) SetNextToken(v string) *DescribeProtocolMountTargetRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeProtocolMountTargetRequest) SetProtocolServiceIds(v string) *DescribeProtocolMountTargetRequest {
	s.ProtocolServiceIds = &v
	return s
}

func (s *DescribeProtocolMountTargetRequest) Validate() error {
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

type DescribeProtocolMountTargetRequestFilters struct {
	// The filter name.
	//
	// 	- ProtocolServiceIds: Filters export directories by protocol service ID.
	//
	// 	- ExportIds: Filters export directories by export directory ID.
	//
	// 	- VpcIds: Filters export directories by VPC ID.
	//
	// 	- FsetIds: Filters export directories by fileset ID.
	//
	// 	- Paths: Filters export directories based on the path of the file system corresponding to the mount target.
	//
	// 	- AccessGroupNames: Filters export directories by permission group name.
	//
	// example:
	//
	// ExportIds
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value. This parameter does not support wildcards.
	//
	// 	- If Key is set to ProtocolServiceIds, set Value to a protocol service ID. You can specify a maximum of 10 protocol service IDs. Example: `ptc-12345678` or `ptc-12345678,ptc-12345679`.
	//
	// 	- If Key is set to ExportIds, set Value to an export directory ID. You can specify a maximum of 10 export directory IDs. Example: `exp-12345678` or `exp-12345678,exp-12345679`.
	//
	// 	- If Key is set to VpcIds, set Value to a VPC ID of the protocol service. You can specify a maximum of 10 VPC IDs. Example: `vpc-12345678` or `vpc-12345678,vpc-12345679`.
	//
	// 	- If Key is set to FsetIds, set Value to a fileset ID. You can specify a maximum of 10 fileset IDs. Example: `fset-12345678` or `fset-12345678,fset-12345679`.
	//
	// 	- If Key is set to Paths, set Value to a path of the file system corresponding to the mount target. You can specify a maximum of 10 paths. Example: `/cpfs/mnt_1/` or `/cpfs/mnt_1/,/cpfs/mnt_2/`.
	//
	// 	- If Key is set to AccessGroupNames, set Value to a permission group name for the protocol service. You can specify a maximum of 10 permission group names. Example: `ag-12345678` or `ag-12345678,ag-12345679`.
	//
	// example:
	//
	// exp-19abf5beab8d****, exp-19acf6beaf7d****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeProtocolMountTargetRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeProtocolMountTargetRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeProtocolMountTargetRequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeProtocolMountTargetRequestFilters) GetValue() *string {
	return s.Value
}

func (s *DescribeProtocolMountTargetRequestFilters) SetKey(v string) *DescribeProtocolMountTargetRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeProtocolMountTargetRequestFilters) SetValue(v string) *DescribeProtocolMountTargetRequestFilters {
	s.Value = &v
	return s
}

func (s *DescribeProtocolMountTargetRequestFilters) Validate() error {
	return dara.Validate(s)
}
