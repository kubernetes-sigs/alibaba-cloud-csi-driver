// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeGroupDriftedNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListNodeGroupDriftedNodesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodeGroupDriftedNodesResponseBody
	GetNextToken() *string
	SetNodes(v []*ListNodeGroupDriftedNodesResponseBodyNodes) *ListNodeGroupDriftedNodesResponseBody
	GetNodes() []*ListNodeGroupDriftedNodesResponseBodyNodes
	SetRequestId(v string) *ListNodeGroupDriftedNodesResponseBody
	GetRequestId() *string
}

type ListNodeGroupDriftedNodesResponseBody struct {
	// The maximum number of entries per page for a single query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token returned by this call. An empty value indicates that no more pages are available.
	//
	// example:
	//
	// 3a6b93229825ac667104463b56790c91
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of nodes that are inconsistent with the node group configuration (paginated).
	Nodes []*ListNodeGroupDriftedNodesResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNodeGroupDriftedNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodeGroupDriftedNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeGroupDriftedNodesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodeGroupDriftedNodesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodeGroupDriftedNodesResponseBody) GetNodes() []*ListNodeGroupDriftedNodesResponseBodyNodes {
	return s.Nodes
}

func (s *ListNodeGroupDriftedNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodeGroupDriftedNodesResponseBody) SetMaxResults(v int32) *ListNodeGroupDriftedNodesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNodeGroupDriftedNodesResponseBody) SetNextToken(v string) *ListNodeGroupDriftedNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNodeGroupDriftedNodesResponseBody) SetNodes(v []*ListNodeGroupDriftedNodesResponseBodyNodes) *ListNodeGroupDriftedNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *ListNodeGroupDriftedNodesResponseBody) SetRequestId(v string) *ListNodeGroupDriftedNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeGroupDriftedNodesResponseBody) Validate() error {
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNodeGroupDriftedNodesResponseBodyNodes struct {
	// The ID of the node.
	//
	// example:
	//
	// node-001
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The list of inconsistent properties for this node.
	PropertyDrifts []*ListNodeGroupDriftedNodesResponseBodyNodesPropertyDrifts `json:"PropertyDrifts,omitempty" xml:"PropertyDrifts,omitempty" type:"Repeated"`
}

func (s ListNodeGroupDriftedNodesResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s ListNodeGroupDriftedNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListNodeGroupDriftedNodesResponseBodyNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *ListNodeGroupDriftedNodesResponseBodyNodes) GetPropertyDrifts() []*ListNodeGroupDriftedNodesResponseBodyNodesPropertyDrifts {
	return s.PropertyDrifts
}

func (s *ListNodeGroupDriftedNodesResponseBodyNodes) SetNodeId(v string) *ListNodeGroupDriftedNodesResponseBodyNodes {
	s.NodeId = &v
	return s
}

func (s *ListNodeGroupDriftedNodesResponseBodyNodes) SetPropertyDrifts(v []*ListNodeGroupDriftedNodesResponseBodyNodesPropertyDrifts) *ListNodeGroupDriftedNodesResponseBodyNodes {
	s.PropertyDrifts = v
	return s
}

func (s *ListNodeGroupDriftedNodesResponseBodyNodes) Validate() error {
	if s.PropertyDrifts != nil {
		for _, item := range s.PropertyDrifts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNodeGroupDriftedNodesResponseBodyNodesPropertyDrifts struct {
	// The current value of the node property. Complex types are serialized as JSON strings.
	//
	// example:
	//
	// old-role
	ActualValue *string `json:"ActualValue,omitempty" xml:"ActualValue,omitempty"`
	// The minimum action required to apply the TargetValue: Refresh / Reboot / Reimage. For more information, refer to the MaxDisruptiveAction parameter description in the RefreshNodeGroupNodes operation.
	//
	// example:
	//
	// Refresh
	MinRequiredAction *string `json:"MinRequiredAction,omitempty" xml:"MinRequiredAction,omitempty"`
	// The property path in dot notation (such as a.b.c), compatible with both flat and nested properties.
	//
	// example:
	//
	// RamRoleName
	PropertyPath *string `json:"PropertyPath,omitempty" xml:"PropertyPath,omitempty"`
	// The target value of the node property. Complex types are serialized as JSON strings.
	//
	// example:
	//
	// new-role
	TargetValue *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s ListNodeGroupDriftedNodesResponseBodyNodesPropertyDrifts) String() string {
	return dara.Prettify(s)
}

func (s ListNodeGroupDriftedNodesResponseBodyNodesPropertyDrifts) GoString() string {
	return s.String()
}

func (s *ListNodeGroupDriftedNodesResponseBodyNodesPropertyDrifts) GetActualValue() *string {
	return s.ActualValue
}

func (s *ListNodeGroupDriftedNodesResponseBodyNodesPropertyDrifts) GetMinRequiredAction() *string {
	return s.MinRequiredAction
}

func (s *ListNodeGroupDriftedNodesResponseBodyNodesPropertyDrifts) GetPropertyPath() *string {
	return s.PropertyPath
}

func (s *ListNodeGroupDriftedNodesResponseBodyNodesPropertyDrifts) GetTargetValue() *string {
	return s.TargetValue
}

func (s *ListNodeGroupDriftedNodesResponseBodyNodesPropertyDrifts) SetActualValue(v string) *ListNodeGroupDriftedNodesResponseBodyNodesPropertyDrifts {
	s.ActualValue = &v
	return s
}

func (s *ListNodeGroupDriftedNodesResponseBodyNodesPropertyDrifts) SetMinRequiredAction(v string) *ListNodeGroupDriftedNodesResponseBodyNodesPropertyDrifts {
	s.MinRequiredAction = &v
	return s
}

func (s *ListNodeGroupDriftedNodesResponseBodyNodesPropertyDrifts) SetPropertyPath(v string) *ListNodeGroupDriftedNodesResponseBodyNodesPropertyDrifts {
	s.PropertyPath = &v
	return s
}

func (s *ListNodeGroupDriftedNodesResponseBodyNodesPropertyDrifts) SetTargetValue(v string) *ListNodeGroupDriftedNodesResponseBodyNodesPropertyDrifts {
	s.TargetValue = &v
	return s
}

func (s *ListNodeGroupDriftedNodesResponseBodyNodesPropertyDrifts) Validate() error {
	return dara.Validate(s)
}
