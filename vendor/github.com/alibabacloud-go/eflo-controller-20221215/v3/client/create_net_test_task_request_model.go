// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetTestTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateNetTestTaskRequest
	GetClusterId() *string
	SetClusterName(v string) *CreateNetTestTaskRequest
	GetClusterName() *string
	SetCommTest(v *CreateNetTestTaskRequestCommTest) *CreateNetTestTaskRequest
	GetCommTest() *CreateNetTestTaskRequestCommTest
	SetDelayTest(v *CreateNetTestTaskRequestDelayTest) *CreateNetTestTaskRequest
	GetDelayTest() *CreateNetTestTaskRequestDelayTest
	SetNetTestType(v string) *CreateNetTestTaskRequest
	GetNetTestType() *string
	SetNetworkMode(v string) *CreateNetTestTaskRequest
	GetNetworkMode() *string
	SetPort(v string) *CreateNetTestTaskRequest
	GetPort() *string
	SetTrafficTest(v *CreateNetTestTaskRequestTrafficTest) *CreateNetTestTaskRequest
	GetTrafficTest() *CreateNetTestTaskRequestTrafficTest
}

type CreateNetTestTaskRequest struct {
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
	CommTest *CreateNetTestTaskRequestCommTest `json:"CommTest,omitempty" xml:"CommTest,omitempty" type:"Struct"`
	// Specify when NetTestType is DelayTest.
	DelayTest *CreateNetTestTaskRequestDelayTest `json:"DelayTest,omitempty" xml:"DelayTest,omitempty" type:"Struct"`
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
	TrafficTest *CreateNetTestTaskRequestTrafficTest `json:"TrafficTest,omitempty" xml:"TrafficTest,omitempty" type:"Struct"`
}

func (s CreateNetTestTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetTestTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateNetTestTaskRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateNetTestTaskRequest) GetCommTest() *CreateNetTestTaskRequestCommTest {
	return s.CommTest
}

func (s *CreateNetTestTaskRequest) GetDelayTest() *CreateNetTestTaskRequestDelayTest {
	return s.DelayTest
}

func (s *CreateNetTestTaskRequest) GetNetTestType() *string {
	return s.NetTestType
}

func (s *CreateNetTestTaskRequest) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *CreateNetTestTaskRequest) GetPort() *string {
	return s.Port
}

func (s *CreateNetTestTaskRequest) GetTrafficTest() *CreateNetTestTaskRequestTrafficTest {
	return s.TrafficTest
}

func (s *CreateNetTestTaskRequest) SetClusterId(v string) *CreateNetTestTaskRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateNetTestTaskRequest) SetClusterName(v string) *CreateNetTestTaskRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateNetTestTaskRequest) SetCommTest(v *CreateNetTestTaskRequestCommTest) *CreateNetTestTaskRequest {
	s.CommTest = v
	return s
}

func (s *CreateNetTestTaskRequest) SetDelayTest(v *CreateNetTestTaskRequestDelayTest) *CreateNetTestTaskRequest {
	s.DelayTest = v
	return s
}

func (s *CreateNetTestTaskRequest) SetNetTestType(v string) *CreateNetTestTaskRequest {
	s.NetTestType = &v
	return s
}

func (s *CreateNetTestTaskRequest) SetNetworkMode(v string) *CreateNetTestTaskRequest {
	s.NetworkMode = &v
	return s
}

func (s *CreateNetTestTaskRequest) SetPort(v string) *CreateNetTestTaskRequest {
	s.Port = &v
	return s
}

func (s *CreateNetTestTaskRequest) SetTrafficTest(v *CreateNetTestTaskRequestTrafficTest) *CreateNetTestTaskRequest {
	s.TrafficTest = v
	return s
}

func (s *CreateNetTestTaskRequest) Validate() error {
	if s.CommTest != nil {
		if err := s.CommTest.Validate(); err != nil {
			return err
		}
	}
	if s.DelayTest != nil {
		if err := s.DelayTest.Validate(); err != nil {
			return err
		}
	}
	if s.TrafficTest != nil {
		if err := s.TrafficTest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateNetTestTaskRequestCommTest struct {
	// The number of GPUs.
	//
	// example:
	//
	// 1
	GPUNum *int64 `json:"GPUNum,omitempty" xml:"GPUNum,omitempty"`
	// The host IDs.
	Hosts []*CreateNetTestTaskRequestCommTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// The communication library model.
	//
	// example:
	//
	// intention_v4
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The CommTest type, which can be ACCL or NCCL.
	//
	// example:
	//
	// ACCL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateNetTestTaskRequestCommTest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetTestTaskRequestCommTest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestCommTest) GetGPUNum() *int64 {
	return s.GPUNum
}

func (s *CreateNetTestTaskRequestCommTest) GetHosts() []*CreateNetTestTaskRequestCommTestHosts {
	return s.Hosts
}

func (s *CreateNetTestTaskRequestCommTest) GetModel() *string {
	return s.Model
}

func (s *CreateNetTestTaskRequestCommTest) GetType() *string {
	return s.Type
}

func (s *CreateNetTestTaskRequestCommTest) SetGPUNum(v int64) *CreateNetTestTaskRequestCommTest {
	s.GPUNum = &v
	return s
}

func (s *CreateNetTestTaskRequestCommTest) SetHosts(v []*CreateNetTestTaskRequestCommTestHosts) *CreateNetTestTaskRequestCommTest {
	s.Hosts = v
	return s
}

func (s *CreateNetTestTaskRequestCommTest) SetModel(v string) *CreateNetTestTaskRequestCommTest {
	s.Model = &v
	return s
}

func (s *CreateNetTestTaskRequestCommTest) SetType(v string) *CreateNetTestTaskRequestCommTest {
	s.Type = &v
	return s
}

func (s *CreateNetTestTaskRequestCommTest) Validate() error {
	if s.Hosts != nil {
		for _, item := range s.Hosts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNetTestTaskRequestCommTestHosts struct {
	// The IP address.
	//
	// example:
	//
	// 169.253.253.15
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-tw-bqisacl3z6l
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// i111670831721110797708
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// VBw
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s CreateNetTestTaskRequestCommTestHosts) String() string {
	return dara.Prettify(s)
}

func (s CreateNetTestTaskRequestCommTestHosts) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestCommTestHosts) GetIP() *string {
	return s.IP
}

func (s *CreateNetTestTaskRequestCommTestHosts) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateNetTestTaskRequestCommTestHosts) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateNetTestTaskRequestCommTestHosts) GetServerName() *string {
	return s.ServerName
}

func (s *CreateNetTestTaskRequestCommTestHosts) SetIP(v string) *CreateNetTestTaskRequestCommTestHosts {
	s.IP = &v
	return s
}

func (s *CreateNetTestTaskRequestCommTestHosts) SetNodeId(v string) *CreateNetTestTaskRequestCommTestHosts {
	s.NodeId = &v
	return s
}

func (s *CreateNetTestTaskRequestCommTestHosts) SetResourceId(v string) *CreateNetTestTaskRequestCommTestHosts {
	s.ResourceId = &v
	return s
}

func (s *CreateNetTestTaskRequestCommTestHosts) SetServerName(v string) *CreateNetTestTaskRequestCommTestHosts {
	s.ServerName = &v
	return s
}

func (s *CreateNetTestTaskRequestCommTestHosts) Validate() error {
	return dara.Validate(s)
}

type CreateNetTestTaskRequestDelayTest struct {
	// The hosts of the test node.
	Hosts []*CreateNetTestTaskRequestDelayTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
}

func (s CreateNetTestTaskRequestDelayTest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetTestTaskRequestDelayTest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestDelayTest) GetHosts() []*CreateNetTestTaskRequestDelayTestHosts {
	return s.Hosts
}

func (s *CreateNetTestTaskRequestDelayTest) SetHosts(v []*CreateNetTestTaskRequestDelayTestHosts) *CreateNetTestTaskRequestDelayTest {
	s.Hosts = v
	return s
}

func (s *CreateNetTestTaskRequestDelayTest) Validate() error {
	if s.Hosts != nil {
		for _, item := range s.Hosts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNetTestTaskRequestDelayTestHosts struct {
	// The bonding of network interface card.
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// The IP address of the node.
	//
	// example:
	//
	// 125.210.225.48
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-cn-fou43an0a05
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// e01-cn-bcd3u1aee06
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// NQU
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s CreateNetTestTaskRequestDelayTestHosts) String() string {
	return dara.Prettify(s)
}

func (s CreateNetTestTaskRequestDelayTestHosts) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestDelayTestHosts) GetBond() *string {
	return s.Bond
}

func (s *CreateNetTestTaskRequestDelayTestHosts) GetIP() *string {
	return s.IP
}

func (s *CreateNetTestTaskRequestDelayTestHosts) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateNetTestTaskRequestDelayTestHosts) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateNetTestTaskRequestDelayTestHosts) GetServerName() *string {
	return s.ServerName
}

func (s *CreateNetTestTaskRequestDelayTestHosts) SetBond(v string) *CreateNetTestTaskRequestDelayTestHosts {
	s.Bond = &v
	return s
}

func (s *CreateNetTestTaskRequestDelayTestHosts) SetIP(v string) *CreateNetTestTaskRequestDelayTestHosts {
	s.IP = &v
	return s
}

func (s *CreateNetTestTaskRequestDelayTestHosts) SetNodeId(v string) *CreateNetTestTaskRequestDelayTestHosts {
	s.NodeId = &v
	return s
}

func (s *CreateNetTestTaskRequestDelayTestHosts) SetResourceId(v string) *CreateNetTestTaskRequestDelayTestHosts {
	s.ResourceId = &v
	return s
}

func (s *CreateNetTestTaskRequestDelayTestHosts) SetServerName(v string) *CreateNetTestTaskRequestDelayTestHosts {
	s.ServerName = &v
	return s
}

func (s *CreateNetTestTaskRequestDelayTestHosts) Validate() error {
	return dara.Validate(s)
}

type CreateNetTestTaskRequestTrafficTest struct {
	// The client IDs.
	Clients []*CreateNetTestTaskRequestTrafficTestClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	// The running duration of the pipeline job. Unit: seconds.
	//
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// If the protocol is RDMA, enter True or False. If the protocol is TCP, leave this field empty.
	//
	// example:
	//
	// False
	GDR *bool `json:"GDR,omitempty" xml:"GDR,omitempty"`
	// The network protocol, which can be RDMA or TCP.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// If the protocol is TCP, enter the number of concurrent connections. If the protocol is RDMA, enter the configured QP value.
	//
	// example:
	//
	// 1
	QP *int64 `json:"QP,omitempty" xml:"QP,omitempty"`
	// The services.
	Servers []*CreateNetTestTaskRequestTrafficTestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	// The traffic model, which can be MTON or Fullmesh.
	//
	// example:
	//
	// Fullmesh
	TrafficModel *string `json:"TrafficModel,omitempty" xml:"TrafficModel,omitempty"`
}

func (s CreateNetTestTaskRequestTrafficTest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetTestTaskRequestTrafficTest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestTrafficTest) GetClients() []*CreateNetTestTaskRequestTrafficTestClients {
	return s.Clients
}

func (s *CreateNetTestTaskRequestTrafficTest) GetDuration() *int64 {
	return s.Duration
}

func (s *CreateNetTestTaskRequestTrafficTest) GetGDR() *bool {
	return s.GDR
}

func (s *CreateNetTestTaskRequestTrafficTest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateNetTestTaskRequestTrafficTest) GetQP() *int64 {
	return s.QP
}

func (s *CreateNetTestTaskRequestTrafficTest) GetServers() []*CreateNetTestTaskRequestTrafficTestServers {
	return s.Servers
}

func (s *CreateNetTestTaskRequestTrafficTest) GetTrafficModel() *string {
	return s.TrafficModel
}

func (s *CreateNetTestTaskRequestTrafficTest) SetClients(v []*CreateNetTestTaskRequestTrafficTestClients) *CreateNetTestTaskRequestTrafficTest {
	s.Clients = v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTest) SetDuration(v int64) *CreateNetTestTaskRequestTrafficTest {
	s.Duration = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTest) SetGDR(v bool) *CreateNetTestTaskRequestTrafficTest {
	s.GDR = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTest) SetProtocol(v string) *CreateNetTestTaskRequestTrafficTest {
	s.Protocol = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTest) SetQP(v int64) *CreateNetTestTaskRequestTrafficTest {
	s.QP = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTest) SetServers(v []*CreateNetTestTaskRequestTrafficTestServers) *CreateNetTestTaskRequestTrafficTest {
	s.Servers = v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTest) SetTrafficModel(v string) *CreateNetTestTaskRequestTrafficTest {
	s.TrafficModel = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTest) Validate() error {
	if s.Clients != nil {
		for _, item := range s.Clients {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Servers != nil {
		for _, item := range s.Servers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNetTestTaskRequestTrafficTestClients struct {
	// The bonding of network interface card.
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// The IP address of the node.
	//
	// example:
	//
	// 192.168.1.1
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-tw-w5elqg7pw18
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// e01-cn-20s41p6cx01
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// xMv
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s CreateNetTestTaskRequestTrafficTestClients) String() string {
	return dara.Prettify(s)
}

func (s CreateNetTestTaskRequestTrafficTestClients) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestTrafficTestClients) GetBond() *string {
	return s.Bond
}

func (s *CreateNetTestTaskRequestTrafficTestClients) GetIP() *string {
	return s.IP
}

func (s *CreateNetTestTaskRequestTrafficTestClients) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateNetTestTaskRequestTrafficTestClients) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateNetTestTaskRequestTrafficTestClients) GetServerName() *string {
	return s.ServerName
}

func (s *CreateNetTestTaskRequestTrafficTestClients) SetBond(v string) *CreateNetTestTaskRequestTrafficTestClients {
	s.Bond = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestClients) SetIP(v string) *CreateNetTestTaskRequestTrafficTestClients {
	s.IP = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestClients) SetNodeId(v string) *CreateNetTestTaskRequestTrafficTestClients {
	s.NodeId = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestClients) SetResourceId(v string) *CreateNetTestTaskRequestTrafficTestClients {
	s.ResourceId = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestClients) SetServerName(v string) *CreateNetTestTaskRequestTrafficTestClients {
	s.ServerName = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestClients) Validate() error {
	return dara.Validate(s)
}

type CreateNetTestTaskRequestTrafficTestServers struct {
	// The bonding of network interface card.
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// The IP address of the node.
	//
	// example:
	//
	// 47.121.110.190
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-tw-bqisacl3z6l
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// e01-cn-wwo3etaqu0b
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// xMv
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s CreateNetTestTaskRequestTrafficTestServers) String() string {
	return dara.Prettify(s)
}

func (s CreateNetTestTaskRequestTrafficTestServers) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestTrafficTestServers) GetBond() *string {
	return s.Bond
}

func (s *CreateNetTestTaskRequestTrafficTestServers) GetIP() *string {
	return s.IP
}

func (s *CreateNetTestTaskRequestTrafficTestServers) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateNetTestTaskRequestTrafficTestServers) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateNetTestTaskRequestTrafficTestServers) GetServerName() *string {
	return s.ServerName
}

func (s *CreateNetTestTaskRequestTrafficTestServers) SetBond(v string) *CreateNetTestTaskRequestTrafficTestServers {
	s.Bond = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestServers) SetIP(v string) *CreateNetTestTaskRequestTrafficTestServers {
	s.IP = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestServers) SetNodeId(v string) *CreateNetTestTaskRequestTrafficTestServers {
	s.NodeId = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestServers) SetResourceId(v string) *CreateNetTestTaskRequestTrafficTestServers {
	s.ResourceId = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestServers) SetServerName(v string) *CreateNetTestTaskRequestTrafficTestServers {
	s.ServerName = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestServers) Validate() error {
	return dara.Validate(s)
}
