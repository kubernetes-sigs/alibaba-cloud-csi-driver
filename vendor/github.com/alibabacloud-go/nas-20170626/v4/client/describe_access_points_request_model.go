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
}

type DescribeAccessPointsRequest struct {
	// The name of the permission group.
	//
	// This parameter is required for a General-purpose File Storage NAS (NAS) file system.
	//
	// The default permission group for virtual private clouds (VPCs) is named DEFAULT_VPC_GROUP_NAME.
	//
	// example:
	//
	// DEFAULT_VPC_GROUP_NAME
	AccessGroup *string `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	// The ID of the file system.
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
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// MTY4NzcxOTcwMjAzMDk2Nzc0MyM4MDM4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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

func (s *DescribeAccessPointsRequest) Validate() error {
	return dara.Validate(s)
}
