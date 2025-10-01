// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *CreateLogAnalysisRequest
	GetFileSystemId() *string
	SetRegionId(v string) *CreateLogAnalysisRequest
	GetRegionId() *string
}

type CreateLogAnalysisRequest struct {
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

func (s CreateLogAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLogAnalysisRequest) GoString() string {
	return s.String()
}

func (s *CreateLogAnalysisRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateLogAnalysisRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLogAnalysisRequest) SetFileSystemId(v string) *CreateLogAnalysisRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateLogAnalysisRequest) SetRegionId(v string) *CreateLogAnalysisRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLogAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
