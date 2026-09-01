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
	// Ensures the idempotence of the request. Generate a unique parameter value from your client to ensure that the value is unique among different requests.
	//
	// ClientToken supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system uses the RequestId of the API request as the ClientToken. The RequestId may vary for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The filter keys for querying protocol service export directories.
	Filters []*DescribeProtocolMountTargetRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The maximum number of results to return per query.
	//
	//  - Valid values: 10 to 100.
	//
	//
	// - Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to initiate the next request when the response is truncated. You can use this token to retrieve the remaining results from where the truncation occurred.
	//
	// example:
	//
	// aBcdg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of protocol service IDs.
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
	// The name of the filter key.
	//
	// - ProtocolServiceIds: filters by protocol service ID.
	//
	// - ExportIds: filters by export directory ID.
	//
	// - VpcIds: filters by VPC ID.
	//
	// - FsetIds: filters by fileset ID.
	//
	// - Paths: filters by the file system path that corresponds to the mount target.
	//
	// - AccessGroupNames: filters by permission group name.
	//
	// example:
	//
	// ExportIds
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter key. Wildcards are not supported.
	//
	// - If Key is set to ProtocolServiceIds, set Value to a protocol service ID. You can specify up to 10 protocol service IDs. Example: `ptc-12345678` or `ptc-12345678,ptc-12345679`.
	//
	// - If Key is set to ExportIds, set Value to an export directory ID. You can specify up to 10 export directory IDs. Example: `exp-12345678` or `exp-12345678,exp-12345679`.
	//
	// - If Key is set to VpcIds, set Value to the VPC ID of the protocol service. You can specify up to 10 VPC IDs. Example: `vpc-12345678` or `vpc-12345678,vpc-12345679`.
	//
	// - If Key is set to FsetIds, set Value to a fileset ID. You can specify up to 10 fileset IDs. Example: `fset-12345678` or `fset-12345678,fset-12345679`.
	//
	// - If Key is set to Paths, set Value to the file system directory that corresponds to the mount target. You can specify up to 10 paths. Example: `/cpfs/mnt_1/` or `/cpfs/mnt_1/,/cpfs/mnt_2/`.
	//
	// - If Key is set to AccessGroupNames, set Value to the permission group name of the protocol service. You can specify up to 10 permission group names. Example: `ag-12345678` or `ag-12345678,ag-12345679`.
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
