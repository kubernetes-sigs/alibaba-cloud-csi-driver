// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *RebootNodesRequest
	GetClusterId() *string
	SetIgnoreFailedNodeTasks(v bool) *RebootNodesRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNodes(v []*string) *RebootNodesRequest
	GetNodes() []*string
}

type RebootNodesRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// i15b480fbd2fcdbc2869cd80
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to allow skipping failed nodes. Default value: False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The nodes.
	Nodes []*string `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s RebootNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootNodesRequest) GoString() string {
	return s.String()
}

func (s *RebootNodesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *RebootNodesRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *RebootNodesRequest) GetNodes() []*string {
	return s.Nodes
}

func (s *RebootNodesRequest) SetClusterId(v string) *RebootNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *RebootNodesRequest) SetIgnoreFailedNodeTasks(v bool) *RebootNodesRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *RebootNodesRequest) SetNodes(v []*string) *RebootNodesRequest {
	s.Nodes = v
	return s
}

func (s *RebootNodesRequest) Validate() error {
	return dara.Validate(s)
}
