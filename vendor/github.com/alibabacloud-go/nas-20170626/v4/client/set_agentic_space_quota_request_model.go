// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAgenticSpaceQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgenticSpaceId(v string) *SetAgenticSpaceQuotaRequest
	GetAgenticSpaceId() *string
	SetClientToken(v string) *SetAgenticSpaceQuotaRequest
	GetClientToken() *string
	SetDryRun(v bool) *SetAgenticSpaceQuotaRequest
	GetDryRun() *bool
	SetFileCountLimit(v int64) *SetAgenticSpaceQuotaRequest
	GetFileCountLimit() *int64
	SetFileSystemId(v string) *SetAgenticSpaceQuotaRequest
	GetFileSystemId() *string
	SetSizeLimit(v int64) *SetAgenticSpaceQuotaRequest
	GetSizeLimit() *int64
}

type SetAgenticSpaceQuotaRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// agentic-229oypxjgpau2****
	AgenticSpaceId *string `json:"AgenticSpaceId,omitempty" xml:"AgenticSpaceId,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// 10000
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// example:
	//
	// 10737418240
	SizeLimit *int64 `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s SetAgenticSpaceQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAgenticSpaceQuotaRequest) GoString() string {
	return s.String()
}

func (s *SetAgenticSpaceQuotaRequest) GetAgenticSpaceId() *string {
	return s.AgenticSpaceId
}

func (s *SetAgenticSpaceQuotaRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SetAgenticSpaceQuotaRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *SetAgenticSpaceQuotaRequest) GetFileCountLimit() *int64 {
	return s.FileCountLimit
}

func (s *SetAgenticSpaceQuotaRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *SetAgenticSpaceQuotaRequest) GetSizeLimit() *int64 {
	return s.SizeLimit
}

func (s *SetAgenticSpaceQuotaRequest) SetAgenticSpaceId(v string) *SetAgenticSpaceQuotaRequest {
	s.AgenticSpaceId = &v
	return s
}

func (s *SetAgenticSpaceQuotaRequest) SetClientToken(v string) *SetAgenticSpaceQuotaRequest {
	s.ClientToken = &v
	return s
}

func (s *SetAgenticSpaceQuotaRequest) SetDryRun(v bool) *SetAgenticSpaceQuotaRequest {
	s.DryRun = &v
	return s
}

func (s *SetAgenticSpaceQuotaRequest) SetFileCountLimit(v int64) *SetAgenticSpaceQuotaRequest {
	s.FileCountLimit = &v
	return s
}

func (s *SetAgenticSpaceQuotaRequest) SetFileSystemId(v string) *SetAgenticSpaceQuotaRequest {
	s.FileSystemId = &v
	return s
}

func (s *SetAgenticSpaceQuotaRequest) SetSizeLimit(v int64) *SetAgenticSpaceQuotaRequest {
	s.SizeLimit = &v
	return s
}

func (s *SetAgenticSpaceQuotaRequest) Validate() error {
	return dara.Validate(s)
}
