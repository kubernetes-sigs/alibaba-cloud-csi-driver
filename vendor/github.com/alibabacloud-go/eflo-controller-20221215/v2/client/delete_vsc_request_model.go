// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVscRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVscRequest
	GetClientToken() *string
	SetVscId(v string) *DeleteVscRequest
	GetVscId() *string
}

type DeleteVscRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the VSC that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsc-001
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s DeleteVscRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVscRequest) GoString() string {
	return s.String()
}

func (s *DeleteVscRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVscRequest) GetVscId() *string {
	return s.VscId
}

func (s *DeleteVscRequest) SetClientToken(v string) *DeleteVscRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVscRequest) SetVscId(v string) *DeleteVscRequest {
	s.VscId = &v
	return s
}

func (s *DeleteVscRequest) Validate() error {
	return dara.Validate(s)
}
