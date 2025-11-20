// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoreFailedNodeTasks(v bool) *StopNodesShrinkRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNodesShrink(v string) *StopNodesShrinkRequest
	GetNodesShrink() *string
}

type StopNodesShrinkRequest struct {
	// Specifies whether to allow skipping failed nodes. Default value: False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The nodes.
	NodesShrink *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s StopNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StopNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopNodesShrinkRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *StopNodesShrinkRequest) GetNodesShrink() *string {
	return s.NodesShrink
}

func (s *StopNodesShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *StopNodesShrinkRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *StopNodesShrinkRequest) SetNodesShrink(v string) *StopNodesShrinkRequest {
	s.NodesShrink = &v
	return s
}

func (s *StopNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
