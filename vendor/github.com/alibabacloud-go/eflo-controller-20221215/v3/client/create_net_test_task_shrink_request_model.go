// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetTestTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateNetTestTaskShrinkRequest
	GetClusterId() *string
	SetClusterName(v string) *CreateNetTestTaskShrinkRequest
	GetClusterName() *string
	SetCommTestShrink(v string) *CreateNetTestTaskShrinkRequest
	GetCommTestShrink() *string
	SetDelayTestShrink(v string) *CreateNetTestTaskShrinkRequest
	GetDelayTestShrink() *string
	SetNetTestType(v string) *CreateNetTestTaskShrinkRequest
	GetNetTestType() *string
	SetNetworkMode(v string) *CreateNetTestTaskShrinkRequest
	GetNetworkMode() *string
	SetPort(v string) *CreateNetTestTaskShrinkRequest
	GetPort() *string
	SetTrafficTestShrink(v string) *CreateNetTestTaskShrinkRequest
	GetTrafficTestShrink() *string
}

type CreateNetTestTaskShrinkRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// i119982311660892626523
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// Eflo-YJ-Test-Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Specify when NetTestType is CommTest.
	CommTestShrink *string `json:"CommTest,omitempty" xml:"CommTest,omitempty"`
	// Specify when NetTestType is DelayTest.
	DelayTestShrink *string `json:"DelayTest,omitempty" xml:"DelayTest,omitempty"`
	// The type of the network test. Valid values: DelayTest, TrafficTest, and CommTest.
	//
	// example:
	//
	// DelayTest
	NetTestType *string `json:"NetTestType,omitempty" xml:"NetTestType,omitempty"`
	// The network mode.
	//
	// example:
	//
	// 2
	NetworkMode *string `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	// The port number.
	//
	// example:
	//
	// 23604
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// If the TrafficModel is Fullmesh, leave this parameter empty.
	TrafficTestShrink *string `json:"TrafficTest,omitempty" xml:"TrafficTest,omitempty"`
}

func (s CreateNetTestTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetTestTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateNetTestTaskShrinkRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateNetTestTaskShrinkRequest) GetCommTestShrink() *string {
	return s.CommTestShrink
}

func (s *CreateNetTestTaskShrinkRequest) GetDelayTestShrink() *string {
	return s.DelayTestShrink
}

func (s *CreateNetTestTaskShrinkRequest) GetNetTestType() *string {
	return s.NetTestType
}

func (s *CreateNetTestTaskShrinkRequest) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *CreateNetTestTaskShrinkRequest) GetPort() *string {
	return s.Port
}

func (s *CreateNetTestTaskShrinkRequest) GetTrafficTestShrink() *string {
	return s.TrafficTestShrink
}

func (s *CreateNetTestTaskShrinkRequest) SetClusterId(v string) *CreateNetTestTaskShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetClusterName(v string) *CreateNetTestTaskShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetCommTestShrink(v string) *CreateNetTestTaskShrinkRequest {
	s.CommTestShrink = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetDelayTestShrink(v string) *CreateNetTestTaskShrinkRequest {
	s.DelayTestShrink = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetNetTestType(v string) *CreateNetTestTaskShrinkRequest {
	s.NetTestType = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetNetworkMode(v string) *CreateNetTestTaskShrinkRequest {
	s.NetworkMode = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetPort(v string) *CreateNetTestTaskShrinkRequest {
	s.Port = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetTrafficTestShrink(v string) *CreateNetTestTaskShrinkRequest {
	s.TrafficTestShrink = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
