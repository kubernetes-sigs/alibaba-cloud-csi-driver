// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHyperNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHyperNodeId(v string) *DeleteHyperNodeRequest
	GetHyperNodeId() *string
}

type DeleteHyperNodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e01-cn-zvp2tgykr08
	HyperNodeId *string `json:"HyperNodeId,omitempty" xml:"HyperNodeId,omitempty"`
}

func (s DeleteHyperNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHyperNodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteHyperNodeRequest) GetHyperNodeId() *string {
	return s.HyperNodeId
}

func (s *DeleteHyperNodeRequest) SetHyperNodeId(v string) *DeleteHyperNodeRequest {
	s.HyperNodeId = &v
	return s
}

func (s *DeleteHyperNodeRequest) Validate() error {
	return dara.Validate(s)
}
