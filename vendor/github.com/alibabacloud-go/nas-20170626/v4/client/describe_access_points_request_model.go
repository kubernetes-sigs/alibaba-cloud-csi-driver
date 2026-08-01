// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessPointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroup(v string) *DescribeAccessPointsRequest
	GetAccessGroup() *string
	SetFileSystemId(v string) *DescribeAccessPointsRequest
	GetFileSystemId() *string
	SetMaxResults(v int32) *DescribeAccessPointsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAccessPointsRequest
	GetNextToken() *string
	SetTag(v []*DescribeAccessPointsRequestTag) *DescribeAccessPointsRequest
	GetTag() []*DescribeAccessPointsRequestTag
}

type DescribeAccessPointsRequest struct {
	// The permission group name.
	//
	// This parameter is required if the file system is a General-purpose NAS file system.
	//
	// Default permission group: DEFAULT_VPC_GROUP_NAME (the default VPC permission group).
	//
	// example:
	//
	// DEFAULT_VPC_GROUP_NAME
	AccessGroup *string `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	// The file system ID.
	//
	// example:
	//
	// 174494****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The number of results for each query.
	//
	// Valid values: 10 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. Set the value to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// MTY4NzcxOTcwMjAzMDk2Nzc0MyM4MDM4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of access point tags.
	Tag []*DescribeAccessPointsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAccessPointsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsRequest) GetAccessGroup() *string {
	return s.AccessGroup
}

func (s *DescribeAccessPointsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeAccessPointsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAccessPointsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAccessPointsRequest) GetTag() []*DescribeAccessPointsRequestTag {
	return s.Tag
}

func (s *DescribeAccessPointsRequest) SetAccessGroup(v string) *DescribeAccessPointsRequest {
	s.AccessGroup = &v
	return s
}

func (s *DescribeAccessPointsRequest) SetFileSystemId(v string) *DescribeAccessPointsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeAccessPointsRequest) SetMaxResults(v int32) *DescribeAccessPointsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAccessPointsRequest) SetNextToken(v string) *DescribeAccessPointsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAccessPointsRequest) SetTag(v []*DescribeAccessPointsRequestTag) *DescribeAccessPointsRequest {
	s.Tag = v
	return s
}

func (s *DescribeAccessPointsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccessPointsRequestTag struct {
	// The tag key.
	//
	// Limits:
	//
	// - Valid values of N: 1 to 20.
	//
	// - The tag key can be up to 128 characters in length.
	//
	// - The tag key cannot start with aliyun or acs:.
	//
	// - The tag key cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// Limits:
	//
	// - Valid values of N: 1 to 20.
	//
	// - The tag value can be up to 128 characters in length.
	//
	// - The tag value cannot start with aliyun or acs:.
	//
	// - The tag value cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAccessPointsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeAccessPointsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeAccessPointsRequestTag) SetKey(v string) *DescribeAccessPointsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeAccessPointsRequestTag) SetValue(v string) *DescribeAccessPointsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeAccessPointsRequestTag) Validate() error {
	return dara.Validate(s)
}
