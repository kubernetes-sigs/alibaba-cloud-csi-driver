// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgenticSpaceId(v string) *DeleteAgenticSpaceRequest
	GetAgenticSpaceId() *string
	SetClientToken(v string) *DeleteAgenticSpaceRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteAgenticSpaceRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *DeleteAgenticSpaceRequest
	GetFileSystemId() *string
}

type DeleteAgenticSpaceRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DeleteAgenticSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticSpaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgenticSpaceRequest) GetAgenticSpaceId() *string {
	return s.AgenticSpaceId
}

func (s *DeleteAgenticSpaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAgenticSpaceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteAgenticSpaceRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteAgenticSpaceRequest) SetAgenticSpaceId(v string) *DeleteAgenticSpaceRequest {
	s.AgenticSpaceId = &v
	return s
}

func (s *DeleteAgenticSpaceRequest) SetClientToken(v string) *DeleteAgenticSpaceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAgenticSpaceRequest) SetDryRun(v bool) *DeleteAgenticSpaceRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteAgenticSpaceRequest) SetFileSystemId(v string) *DeleteAgenticSpaceRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteAgenticSpaceRequest) Validate() error {
	return dara.Validate(s)
}
