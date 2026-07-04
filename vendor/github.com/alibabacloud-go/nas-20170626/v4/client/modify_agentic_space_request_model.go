// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAgenticSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgenticSpaceId(v string) *ModifyAgenticSpaceRequest
	GetAgenticSpaceId() *string
	SetClientToken(v string) *ModifyAgenticSpaceRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyAgenticSpaceRequest
	GetDescription() *string
	SetDryRun(v bool) *ModifyAgenticSpaceRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *ModifyAgenticSpaceRequest
	GetFileSystemId() *string
}

type ModifyAgenticSpaceRequest struct {
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
	// Agentic Space Description
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
}

func (s ModifyAgenticSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgenticSpaceRequest) GoString() string {
	return s.String()
}

func (s *ModifyAgenticSpaceRequest) GetAgenticSpaceId() *string {
	return s.AgenticSpaceId
}

func (s *ModifyAgenticSpaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAgenticSpaceRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyAgenticSpaceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyAgenticSpaceRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyAgenticSpaceRequest) SetAgenticSpaceId(v string) *ModifyAgenticSpaceRequest {
	s.AgenticSpaceId = &v
	return s
}

func (s *ModifyAgenticSpaceRequest) SetClientToken(v string) *ModifyAgenticSpaceRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAgenticSpaceRequest) SetDescription(v string) *ModifyAgenticSpaceRequest {
	s.Description = &v
	return s
}

func (s *ModifyAgenticSpaceRequest) SetDryRun(v bool) *ModifyAgenticSpaceRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyAgenticSpaceRequest) SetFileSystemId(v string) *ModifyAgenticSpaceRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyAgenticSpaceRequest) Validate() error {
	return dara.Validate(s)
}
