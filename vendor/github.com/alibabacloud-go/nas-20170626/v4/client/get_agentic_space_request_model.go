// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgenticSpaceId(v string) *GetAgenticSpaceRequest
	GetAgenticSpaceId() *string
	SetFileSystemId(v string) *GetAgenticSpaceRequest
	GetFileSystemId() *string
}

type GetAgenticSpaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// agentic-229oypxjgpau2****
	AgenticSpaceId *string `json:"AgenticSpaceId,omitempty" xml:"AgenticSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 06229oypxjgox0u****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s GetAgenticSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticSpaceRequest) GoString() string {
	return s.String()
}

func (s *GetAgenticSpaceRequest) GetAgenticSpaceId() *string {
	return s.AgenticSpaceId
}

func (s *GetAgenticSpaceRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *GetAgenticSpaceRequest) SetAgenticSpaceId(v string) *GetAgenticSpaceRequest {
	s.AgenticSpaceId = &v
	return s
}

func (s *GetAgenticSpaceRequest) SetFileSystemId(v string) *GetAgenticSpaceRequest {
	s.FileSystemId = &v
	return s
}

func (s *GetAgenticSpaceRequest) Validate() error {
	return dara.Validate(s)
}
