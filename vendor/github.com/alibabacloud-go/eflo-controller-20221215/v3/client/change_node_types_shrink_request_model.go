// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeNodeTypesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeIdsShrink(v string) *ChangeNodeTypesShrinkRequest
	GetNodeIdsShrink() *string
	SetNodeType(v string) *ChangeNodeTypesShrinkRequest
	GetNodeType() *string
}

type ChangeNodeTypesShrinkRequest struct {
	NodeIdsShrink *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	// example:
	//
	// standard
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s ChangeNodeTypesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeNodeTypesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ChangeNodeTypesShrinkRequest) GetNodeIdsShrink() *string {
	return s.NodeIdsShrink
}

func (s *ChangeNodeTypesShrinkRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *ChangeNodeTypesShrinkRequest) SetNodeIdsShrink(v string) *ChangeNodeTypesShrinkRequest {
	s.NodeIdsShrink = &v
	return s
}

func (s *ChangeNodeTypesShrinkRequest) SetNodeType(v string) *ChangeNodeTypesShrinkRequest {
	s.NodeType = &v
	return s
}

func (s *ChangeNodeTypesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
