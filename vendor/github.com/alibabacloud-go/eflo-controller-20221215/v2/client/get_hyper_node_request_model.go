// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHyperNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHyperNodeId(v string) *GetHyperNodeRequest
	GetHyperNodeId() *string
}

type GetHyperNodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e01-cn-zvp2tgykr08
	HyperNodeId *string `json:"HyperNodeId,omitempty" xml:"HyperNodeId,omitempty"`
}

func (s GetHyperNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHyperNodeRequest) GoString() string {
	return s.String()
}

func (s *GetHyperNodeRequest) GetHyperNodeId() *string {
	return s.HyperNodeId
}

func (s *GetHyperNodeRequest) SetHyperNodeId(v string) *GetHyperNodeRequest {
	s.HyperNodeId = &v
	return s
}

func (s *GetHyperNodeRequest) Validate() error {
	return dara.Validate(s)
}
