// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *DeleteNodeRequest
	GetNodeId() *string
}

type DeleteNodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e01-cn-kvw****dn04
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DeleteNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DeleteNodeRequest) SetNodeId(v string) *DeleteNodeRequest {
	s.NodeId = &v
	return s
}

func (s *DeleteNodeRequest) Validate() error {
	return dara.Validate(s)
}
