// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoreFailedNodeTasks(v bool) *StopNodesRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNodes(v []*string) *StopNodesRequest
	GetNodes() []*string
}

type StopNodesRequest struct {
	// Specifies whether to allow skipping failed nodes. Default value: False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The nodes.
	Nodes []*string `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s StopNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s StopNodesRequest) GoString() string {
	return s.String()
}

func (s *StopNodesRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *StopNodesRequest) GetNodes() []*string {
	return s.Nodes
}

func (s *StopNodesRequest) SetIgnoreFailedNodeTasks(v bool) *StopNodesRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *StopNodesRequest) SetNodes(v []*string) *StopNodesRequest {
	s.Nodes = v
	return s
}

func (s *StopNodesRequest) Validate() error {
	return dara.Validate(s)
}
