// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDirQuotasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DescribeDirQuotasRequest
	GetFileSystemId() *string
	SetPageNumber(v int32) *DescribeDirQuotasRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDirQuotasRequest
	GetPageSize() *int32
	SetPath(v string) *DescribeDirQuotasRequest
	GetPath() *string
}

type DescribeDirQuotasRequest struct {
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The absolute path of a directory.
	//
	// If you do not specify this parameter, all directories for which quotas are created are returned.
	//
	// example:
	//
	// /data/sub1
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DescribeDirQuotasRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDirQuotasRequest) GoString() string {
	return s.String()
}

func (s *DescribeDirQuotasRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeDirQuotasRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDirQuotasRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDirQuotasRequest) GetPath() *string {
	return s.Path
}

func (s *DescribeDirQuotasRequest) SetFileSystemId(v string) *DescribeDirQuotasRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeDirQuotasRequest) SetPageNumber(v int32) *DescribeDirQuotasRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDirQuotasRequest) SetPageSize(v int32) *DescribeDirQuotasRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDirQuotasRequest) SetPath(v string) *DescribeDirQuotasRequest {
	s.Path = &v
	return s
}

func (s *DescribeDirQuotasRequest) Validate() error {
	return dara.Validate(s)
}
