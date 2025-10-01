// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DeleteLogAnalysisRequest
	GetFileSystemId() *string
	SetRegionId(v string) *DeleteLogAnalysisRequest
	GetRegionId() *string
}

type DeleteLogAnalysisRequest struct {
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 174494xxxx
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteLogAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogAnalysisRequest) GoString() string {
	return s.String()
}

func (s *DeleteLogAnalysisRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteLogAnalysisRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLogAnalysisRequest) SetFileSystemId(v string) *DeleteLogAnalysisRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteLogAnalysisRequest) SetRegionId(v string) *DeleteLogAnalysisRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLogAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
