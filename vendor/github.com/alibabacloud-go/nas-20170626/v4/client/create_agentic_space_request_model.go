// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAzone(v string) *CreateAgenticSpaceRequest
	GetAzone() *string
	SetClientToken(v string) *CreateAgenticSpaceRequest
	GetClientToken() *string
	SetDescription(v string) *CreateAgenticSpaceRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateAgenticSpaceRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *CreateAgenticSpaceRequest
	GetFileSystemId() *string
	SetFileSystemPath(v string) *CreateAgenticSpaceRequest
	GetFileSystemPath() *string
	SetQuota(v *CreateAgenticSpaceRequestQuota) *CreateAgenticSpaceRequest
	GetQuota() *CreateAgenticSpaceRequestQuota
}

type CreateAgenticSpaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-i
	Azone *string `json:"Azone,omitempty" xml:"Azone,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// AgenticSpace Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// /path/
	FileSystemPath *string `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	// This parameter is required.
	Quota *CreateAgenticSpaceRequestQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
}

func (s CreateAgenticSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticSpaceRequest) GoString() string {
	return s.String()
}

func (s *CreateAgenticSpaceRequest) GetAzone() *string {
	return s.Azone
}

func (s *CreateAgenticSpaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAgenticSpaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAgenticSpaceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateAgenticSpaceRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateAgenticSpaceRequest) GetFileSystemPath() *string {
	return s.FileSystemPath
}

func (s *CreateAgenticSpaceRequest) GetQuota() *CreateAgenticSpaceRequestQuota {
	return s.Quota
}

func (s *CreateAgenticSpaceRequest) SetAzone(v string) *CreateAgenticSpaceRequest {
	s.Azone = &v
	return s
}

func (s *CreateAgenticSpaceRequest) SetClientToken(v string) *CreateAgenticSpaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAgenticSpaceRequest) SetDescription(v string) *CreateAgenticSpaceRequest {
	s.Description = &v
	return s
}

func (s *CreateAgenticSpaceRequest) SetDryRun(v bool) *CreateAgenticSpaceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAgenticSpaceRequest) SetFileSystemId(v string) *CreateAgenticSpaceRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateAgenticSpaceRequest) SetFileSystemPath(v string) *CreateAgenticSpaceRequest {
	s.FileSystemPath = &v
	return s
}

func (s *CreateAgenticSpaceRequest) SetQuota(v *CreateAgenticSpaceRequestQuota) *CreateAgenticSpaceRequest {
	s.Quota = v
	return s
}

func (s *CreateAgenticSpaceRequest) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgenticSpaceRequestQuota struct {
	// This parameter is required.
	//
	// example:
	//
	// 10000
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10737418240
	SizeLimit *int64 `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s CreateAgenticSpaceRequestQuota) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticSpaceRequestQuota) GoString() string {
	return s.String()
}

func (s *CreateAgenticSpaceRequestQuota) GetFileCountLimit() *int64 {
	return s.FileCountLimit
}

func (s *CreateAgenticSpaceRequestQuota) GetSizeLimit() *int64 {
	return s.SizeLimit
}

func (s *CreateAgenticSpaceRequestQuota) SetFileCountLimit(v int64) *CreateAgenticSpaceRequestQuota {
	s.FileCountLimit = &v
	return s
}

func (s *CreateAgenticSpaceRequestQuota) SetSizeLimit(v int64) *CreateAgenticSpaceRequestQuota {
	s.SizeLimit = &v
	return s
}

func (s *CreateAgenticSpaceRequestQuota) Validate() error {
	return dara.Validate(s)
}
