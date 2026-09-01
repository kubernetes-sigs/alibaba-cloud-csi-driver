// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataInsightDirectoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *ListDataInsightDirectoriesRequest
	GetFileSystemId() *string
	SetMaxResults(v int32) *ListDataInsightDirectoriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataInsightDirectoriesRequest
	GetNextToken() *string
	SetParentDir(v string) *ListDataInsightDirectoriesRequest
	GetParentDir() *string
}

type ListDataInsightDirectoriesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-030wldnqm8evtpy****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// ""
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// /
	ParentDir *string `json:"ParentDir,omitempty" xml:"ParentDir,omitempty"`
}

func (s ListDataInsightDirectoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataInsightDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *ListDataInsightDirectoriesRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ListDataInsightDirectoriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataInsightDirectoriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataInsightDirectoriesRequest) GetParentDir() *string {
	return s.ParentDir
}

func (s *ListDataInsightDirectoriesRequest) SetFileSystemId(v string) *ListDataInsightDirectoriesRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListDataInsightDirectoriesRequest) SetMaxResults(v int32) *ListDataInsightDirectoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataInsightDirectoriesRequest) SetNextToken(v string) *ListDataInsightDirectoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataInsightDirectoriesRequest) SetParentDir(v string) *ListDataInsightDirectoriesRequest {
	s.ParentDir = &v
	return s
}

func (s *ListDataInsightDirectoriesRequest) Validate() error {
	return dara.Validate(s)
}
