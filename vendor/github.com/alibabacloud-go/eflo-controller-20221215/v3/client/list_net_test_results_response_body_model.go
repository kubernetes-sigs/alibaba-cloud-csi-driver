// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetTestResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListNetTestResultsResponseBody
	GetMaxResults() *int64
	SetNetTestResults(v []*ListNetTestResultsResponseBodyNetTestResults) *ListNetTestResultsResponseBody
	GetNetTestResults() []*ListNetTestResultsResponseBodyNetTestResults
	SetNextToken(v string) *ListNetTestResultsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListNetTestResultsResponseBody
	GetRequestId() *string
}

type ListNetTestResultsResponseBody struct {
	// The number of entries to return on each page. Maximum value: 100.
	//
	// Default value:
	//
	// 	- If you do not configure this parameter or if you set this parameter to a value less than 20, the default value is 20.
	//
	// 	- If you set this parameter to a value that is greater than 100, the default value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The results.
	NetTestResults []*ListNetTestResultsResponseBodyNetTestResults `json:"NetTestResults,omitempty" xml:"NetTestResults,omitempty" type:"Repeated"`
	// The token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 3a6b93229825ac667104463b56790c91
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3C683243-7915-57FB-9570-A2932C1C0F78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNetTestResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListNetTestResultsResponseBody) GetNetTestResults() []*ListNetTestResultsResponseBodyNetTestResults {
	return s.NetTestResults
}

func (s *ListNetTestResultsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNetTestResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNetTestResultsResponseBody) SetMaxResults(v int64) *ListNetTestResultsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNetTestResultsResponseBody) SetNetTestResults(v []*ListNetTestResultsResponseBodyNetTestResults) *ListNetTestResultsResponseBody {
	s.NetTestResults = v
	return s
}

func (s *ListNetTestResultsResponseBody) SetNextToken(v string) *ListNetTestResultsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNetTestResultsResponseBody) SetRequestId(v string) *ListNetTestResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetTestResultsResponseBody) Validate() error {
	if s.NetTestResults != nil {
		for _, item := range s.NetTestResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNetTestResultsResponseBodyNetTestResults struct {
	// The cluster ID.
	//
	// example:
	//
	// i110667211718265012218
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Returned when NetTestType is CommTest.
	CommTest *ListNetTestResultsResponseBodyNetTestResultsCommTest `json:"CommTest,omitempty" xml:"CommTest,omitempty" type:"Struct"`
	// The creation time.
	//
	// example:
	//
	// 2024-01-19T02:18:12Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Returned when NetTestType is DelayTest.
	DelayTest *ListNetTestResultsResponseBodyNetTestResultsDelayTest `json:"DelayTest,omitempty" xml:"DelayTest,omitempty" type:"Struct"`
	// The finish time.
	//
	// example:
	//
	// 2024-10-30T02:07Z
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The type of the network test.
	//
	// example:
	//
	// NetDiag
	NetTestType *string `json:"NetTestType,omitempty" xml:"NetTestType,omitempty"`
	// The network mode.
	//
	// example:
	//
	// 01
	NetworkMode *string `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	// The port number.
	//
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The status of the network test task. Valid values:\\
	//
	// ● InProgress\\
	//
	// ● Finished\\
	//
	// ● Failed
	//
	// example:
	//
	// InProgress
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The test ID. The unique identifier of the resource test task.
	//
	// example:
	//
	// String	i-uf6i0tv2refv8wz*****
	TestId *string `json:"TestId,omitempty" xml:"TestId,omitempty"`
	// Returned when NetTestType is TrafficTest.
	TrafficTest *ListNetTestResultsResponseBodyNetTestResultsTrafficTest `json:"TrafficTest,omitempty" xml:"TrafficTest,omitempty" type:"Struct"`
}

func (s ListNetTestResultsResponseBodyNetTestResults) String() string {
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResults) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetCommTest() *ListNetTestResultsResponseBodyNetTestResultsCommTest {
	return s.CommTest
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetDelayTest() *ListNetTestResultsResponseBodyNetTestResultsDelayTest {
	return s.DelayTest
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetFinishedTime() *string {
	return s.FinishedTime
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetNetTestType() *string {
	return s.NetTestType
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetPort() *string {
	return s.Port
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetStatus() *string {
	return s.Status
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetTestId() *string {
	return s.TestId
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetTrafficTest() *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	return s.TrafficTest
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetClusterId(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.ClusterId = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetClusterName(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.ClusterName = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetCommTest(v *ListNetTestResultsResponseBodyNetTestResultsCommTest) *ListNetTestResultsResponseBodyNetTestResults {
	s.CommTest = v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetCreationTime(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.CreationTime = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetDelayTest(v *ListNetTestResultsResponseBodyNetTestResultsDelayTest) *ListNetTestResultsResponseBodyNetTestResults {
	s.DelayTest = v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetFinishedTime(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.FinishedTime = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetNetTestType(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.NetTestType = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetNetworkMode(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.NetworkMode = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetPort(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.Port = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetStatus(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.Status = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetTestId(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.TestId = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetTrafficTest(v *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) *ListNetTestResultsResponseBodyNetTestResults {
	s.TrafficTest = v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) Validate() error {
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

type ListNetTestResultsResponseBodyNetTestResultsCommTest struct {
	// The number of GPUs.
	//
	// example:
	//
	// 4
	GPUNum *string `json:"GPUNum,omitempty" xml:"GPUNum,omitempty"`
	// The hosts of the test node.
	Hosts []*ListNetTestResultsResponseBodyNetTestResultsCommTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// The communication library model.
	//
	// example:
	//
	// AllToAll
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The CommTest type, which can be ACCL or NCCL.
	//
	// example:
	//
	// ACCL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsCommTest) String() string {
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsCommTest) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) GetGPUNum() *string {
	return s.GPUNum
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) GetHosts() []*ListNetTestResultsResponseBodyNetTestResultsCommTestHosts {
	return s.Hosts
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) GetModel() *string {
	return s.Model
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) GetType() *string {
	return s.Type
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) SetGPUNum(v string) *ListNetTestResultsResponseBodyNetTestResultsCommTest {
	s.GPUNum = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) SetHosts(v []*ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) *ListNetTestResultsResponseBodyNetTestResultsCommTest {
	s.Hosts = v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) SetModel(v string) *ListNetTestResultsResponseBodyNetTestResultsCommTest {
	s.Model = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) SetType(v string) *ListNetTestResultsResponseBodyNetTestResultsCommTest {
	s.Type = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) Validate() error {
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

type ListNetTestResultsResponseBodyNetTestResultsCommTestHosts struct {
	// The IP address of the node.
	//
	// example:
	//
	// 10.51.16.21
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// e01-cn-wwo3eteze19
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// www.xinjiaoyu.com
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) String() string {
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) GetIP() *string {
	return s.IP
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) GetServerName() *string {
	return s.ServerName
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) SetIP(v string) *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts {
	s.IP = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) SetResourceId(v string) *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts {
	s.ResourceId = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) SetServerName(v string) *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts {
	s.ServerName = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) Validate() error {
	return dara.Validate(s)
}

type ListNetTestResultsResponseBodyNetTestResultsDelayTest struct {
	// The hosts.
	Hosts []*ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsDelayTest) String() string {
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsDelayTest) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTest) GetHosts() []*ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts {
	return s.Hosts
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTest) SetHosts(v []*ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) *ListNetTestResultsResponseBodyNetTestResultsDelayTest {
	s.Hosts = v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTest) Validate() error {
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

type ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts struct {
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
	// pgm-bp174z988a27wre71o.pg.rds.aliyuncs.com
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// e01-cn-wwo3eteze19
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// WrF
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) String() string {
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) GetBond() *string {
	return s.Bond
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) GetIP() *string {
	return s.IP
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) GetServerName() *string {
	return s.ServerName
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) SetBond(v string) *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts {
	s.Bond = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) SetIP(v string) *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts {
	s.IP = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) SetResourceId(v string) *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts {
	s.ResourceId = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) SetServerName(v string) *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts {
	s.ServerName = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) Validate() error {
	return dara.Validate(s)
}

type ListNetTestResultsResponseBodyNetTestResultsTrafficTest struct {
	// The clients.
	Clients []*ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	// The running duration of the pipeline job. Unit: seconds.
	//
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// If the protocol is RDMA, can be True or False. If the protocol is TCP, this field is empty.
	//
	// example:
	//
	// True
	GDR *string `json:"GDR,omitempty" xml:"GDR,omitempty"`
	// The network protocol, which can be RDMA or TCP.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// If the protocol is TCP, the number of concurrent connections. If the protocol is RDMA, the configured QP value.
	//
	// example:
	//
	// RDMA
	QP *int64 `json:"QP,omitempty" xml:"QP,omitempty"`
	// If the TrafficModel is Fullmesh, this parameter is empty.
	Servers []*ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	// The traffic model, which can be MTON or Fullmesh.
	//
	// example:
	//
	// Fullmesh
	TrafficModel *string `json:"TrafficModel,omitempty" xml:"TrafficModel,omitempty"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTest) String() string {
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTest) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) GetClients() []*ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients {
	return s.Clients
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) GetDuration() *int64 {
	return s.Duration
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) GetGDR() *string {
	return s.GDR
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) GetProtocol() *string {
	return s.Protocol
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) GetQP() *int64 {
	return s.QP
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) GetServers() []*ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers {
	return s.Servers
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) GetTrafficModel() *string {
	return s.TrafficModel
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetClients(v []*ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.Clients = v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetDuration(v int64) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.Duration = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetGDR(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.GDR = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetProtocol(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.Protocol = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetQP(v int64) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.QP = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetServers(v []*ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.Servers = v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetTrafficModel(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.TrafficModel = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) Validate() error {
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

type ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients struct {
	// The bonding of network interface card.
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 74.73.100.1
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// e01-cn-20p36bqet39
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// prod-gf-cn.juequling.com
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) String() string {
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) GetBond() *string {
	return s.Bond
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) GetIP() *string {
	return s.IP
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) GetServerName() *string {
	return s.ServerName
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) SetBond(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients {
	s.Bond = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) SetIP(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients {
	s.IP = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) SetResourceId(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients {
	s.ResourceId = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) SetServerName(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients {
	s.ServerName = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) Validate() error {
	return dara.Validate(s)
}

type ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers struct {
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
	// 10.1.168.183
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// e01-cn-wwo3eteze19
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// prod-gf-cn.juequling.com
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) String() string {
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) GetBond() *string {
	return s.Bond
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) GetIP() *string {
	return s.IP
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) GetServerName() *string {
	return s.ServerName
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) SetBond(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers {
	s.Bond = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) SetIP(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers {
	s.IP = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) SetResourceId(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers {
	s.ResourceId = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) SetServerName(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers {
	s.ServerName = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) Validate() error {
	return dara.Validate(s)
}
