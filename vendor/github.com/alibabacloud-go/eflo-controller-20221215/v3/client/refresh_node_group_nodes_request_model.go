// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshNodeGroupNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxDisruptiveAction(v string) *RefreshNodeGroupNodesRequest
	GetMaxDisruptiveAction() *string
	SetNodeGroupId(v string) *RefreshNodeGroupNodesRequest
	GetNodeGroupId() *string
	SetNodeIds(v []*string) *RefreshNodeGroupNodesRequest
	GetNodeIds() []*string
}

type RefreshNodeGroupNodesRequest struct {
	// The maximum disruptive action level allowed for the refresh operation. The system independently evaluates the action level required to refresh each drifted property of a node and performs the refresh within the specified action level constraint. If the action level required for a property exceeds the specified level, that property is skipped. Action levels in increasing order of disruption: Refresh < Reboot < Reimage.
	//
	// - Refresh (default): refreshes the configuration in place without restarting or reimaging. Currently applicable only to the RamRoleName property.
	//
	// - Reboot (not currently supported): allows restarting the node for the configuration to take effect. Supported properties include system cloud disk type and all properties supported by Refresh.
	//
	// - Reimage (not currently supported): allows reimaging the node for the configuration to take effect. Supported properties include image ID and all properties supported by Reboot.
	//
	// example:
	//
	// Refresh
	MaxDisruptiveAction *string `json:"MaxDisruptiveAction,omitempty" xml:"MaxDisruptiveAction,omitempty"`
	// The node group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-3525
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The filter scope for node refresh. If not specified, all nodes in the node group are included. <warning>If the instance type is a hypernode, pass the TrayNode ID, not the HyperNodeId.</warning>
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
}

func (s RefreshNodeGroupNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshNodeGroupNodesRequest) GoString() string {
	return s.String()
}

func (s *RefreshNodeGroupNodesRequest) GetMaxDisruptiveAction() *string {
	return s.MaxDisruptiveAction
}

func (s *RefreshNodeGroupNodesRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *RefreshNodeGroupNodesRequest) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *RefreshNodeGroupNodesRequest) SetMaxDisruptiveAction(v string) *RefreshNodeGroupNodesRequest {
	s.MaxDisruptiveAction = &v
	return s
}

func (s *RefreshNodeGroupNodesRequest) SetNodeGroupId(v string) *RefreshNodeGroupNodesRequest {
	s.NodeGroupId = &v
	return s
}

func (s *RefreshNodeGroupNodesRequest) SetNodeIds(v []*string) *RefreshNodeGroupNodesRequest {
	s.NodeIds = v
	return s
}

func (s *RefreshNodeGroupNodesRequest) Validate() error {
	return dara.Validate(s)
}
