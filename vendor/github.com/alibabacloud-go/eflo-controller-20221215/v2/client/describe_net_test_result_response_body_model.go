// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetTestResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeNetTestResultResponseBody
	GetClusterId() *string
	SetClusterName(v string) *DescribeNetTestResultResponseBody
	GetClusterName() *string
	SetCommTest(v *DescribeNetTestResultResponseBodyCommTest) *DescribeNetTestResultResponseBody
	GetCommTest() *DescribeNetTestResultResponseBodyCommTest
	SetCreationTime(v string) *DescribeNetTestResultResponseBody
	GetCreationTime() *string
	SetDelayTest(v *DescribeNetTestResultResponseBodyDelayTest) *DescribeNetTestResultResponseBody
	GetDelayTest() *DescribeNetTestResultResponseBodyDelayTest
	SetFinishedTime(v string) *DescribeNetTestResultResponseBody
	GetFinishedTime() *string
	SetNetTestType(v string) *DescribeNetTestResultResponseBody
	GetNetTestType() *string
	SetPort(v string) *DescribeNetTestResultResponseBody
	GetPort() *string
	SetRequestId(v string) *DescribeNetTestResultResponseBody
	GetRequestId() *string
	SetResultDetial(v string) *DescribeNetTestResultResponseBody
	GetResultDetial() *string
	SetStatus(v string) *DescribeNetTestResultResponseBody
	GetStatus() *string
	SetTestId(v string) *DescribeNetTestResultResponseBody
	GetTestId() *string
	SetTrafficTest(v *DescribeNetTestResultResponseBodyTrafficTest) *DescribeNetTestResultResponseBody
	GetTrafficTest() *DescribeNetTestResultResponseBodyTrafficTest
}

type DescribeNetTestResultResponseBody struct {
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
	// Standard_Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Specify when NetTestType is CommTest.
	CommTest *DescribeNetTestResultResponseBodyCommTest `json:"CommTest,omitempty" xml:"CommTest,omitempty" type:"Struct"`
	// create time
	//
	// example:
	//
	// 2024-10-15T10:25:42+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Fill in when the network test type is a delay test.
	DelayTest *DescribeNetTestResultResponseBodyDelayTest `json:"DelayTest,omitempty" xml:"DelayTest,omitempty" type:"Struct"`
	// finish time
	//
	// example:
	//
	// 2024-10-16T02:04Z
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The type of the network test.
	//
	// example:
	//
	// DelayTest
	NetTestType *string `json:"NetTestType,omitempty" xml:"NetTestType,omitempty"`
	// Test port number.
	//
	// example:
	//
	// 23604
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8F065DDD-6996-5973-9691-9EC57BD0072E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// result detail
	//
	// example:
	//
	// {}
	ResultDetial *string `json:"ResultDetial,omitempty" xml:"ResultDetial,omitempty"`
	// status of session
	//
	// example:
	//
	// Failed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the test task. The unique identifier of a network test task.
	//
	// example:
	//
	// af35ce53-7c35-4277-834a-fbf49c316a96
	TestId *string `json:"TestId,omitempty" xml:"TestId,omitempty"`
	// Fill in when the network test type is a traffic test.
	TrafficTest *DescribeNetTestResultResponseBodyTrafficTest `json:"TrafficTest,omitempty" xml:"TrafficTest,omitempty" type:"Struct"`
}

func (s DescribeNetTestResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetTestResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeNetTestResultResponseBody) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeNetTestResultResponseBody) GetCommTest() *DescribeNetTestResultResponseBodyCommTest {
	return s.CommTest
}

func (s *DescribeNetTestResultResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeNetTestResultResponseBody) GetDelayTest() *DescribeNetTestResultResponseBodyDelayTest {
	return s.DelayTest
}

func (s *DescribeNetTestResultResponseBody) GetFinishedTime() *string {
	return s.FinishedTime
}

func (s *DescribeNetTestResultResponseBody) GetNetTestType() *string {
	return s.NetTestType
}

func (s *DescribeNetTestResultResponseBody) GetPort() *string {
	return s.Port
}

func (s *DescribeNetTestResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetTestResultResponseBody) GetResultDetial() *string {
	return s.ResultDetial
}

func (s *DescribeNetTestResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeNetTestResultResponseBody) GetTestId() *string {
	return s.TestId
}

func (s *DescribeNetTestResultResponseBody) GetTrafficTest() *DescribeNetTestResultResponseBodyTrafficTest {
	return s.TrafficTest
}

func (s *DescribeNetTestResultResponseBody) SetClusterId(v string) *DescribeNetTestResultResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetClusterName(v string) *DescribeNetTestResultResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetCommTest(v *DescribeNetTestResultResponseBodyCommTest) *DescribeNetTestResultResponseBody {
	s.CommTest = v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetCreationTime(v string) *DescribeNetTestResultResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetDelayTest(v *DescribeNetTestResultResponseBodyDelayTest) *DescribeNetTestResultResponseBody {
	s.DelayTest = v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetFinishedTime(v string) *DescribeNetTestResultResponseBody {
	s.FinishedTime = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetNetTestType(v string) *DescribeNetTestResultResponseBody {
	s.NetTestType = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetPort(v string) *DescribeNetTestResultResponseBody {
	s.Port = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetRequestId(v string) *DescribeNetTestResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetResultDetial(v string) *DescribeNetTestResultResponseBody {
	s.ResultDetial = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetStatus(v string) *DescribeNetTestResultResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetTestId(v string) *DescribeNetTestResultResponseBody {
	s.TestId = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetTrafficTest(v *DescribeNetTestResultResponseBodyTrafficTest) *DescribeNetTestResultResponseBody {
	s.TrafficTest = v
	return s
}

func (s *DescribeNetTestResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNetTestResultResponseBodyCommTest struct {
	// Number of GPUs
	//
	// example:
	//
	// 1
	GPUNum *string `json:"GPUNum,omitempty" xml:"GPUNum,omitempty"`
	// All hosts infomation
	Hosts []*DescribeNetTestResultResponseBodyCommTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
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

func (s DescribeNetTestResultResponseBodyCommTest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyCommTest) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyCommTest) GetGPUNum() *string {
	return s.GPUNum
}

func (s *DescribeNetTestResultResponseBodyCommTest) GetHosts() []*DescribeNetTestResultResponseBodyCommTestHosts {
	return s.Hosts
}

func (s *DescribeNetTestResultResponseBodyCommTest) GetModel() *string {
	return s.Model
}

func (s *DescribeNetTestResultResponseBodyCommTest) GetType() *string {
	return s.Type
}

func (s *DescribeNetTestResultResponseBodyCommTest) SetGPUNum(v string) *DescribeNetTestResultResponseBodyCommTest {
	s.GPUNum = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyCommTest) SetHosts(v []*DescribeNetTestResultResponseBodyCommTestHosts) *DescribeNetTestResultResponseBodyCommTest {
	s.Hosts = v
	return s
}

func (s *DescribeNetTestResultResponseBodyCommTest) SetModel(v string) *DescribeNetTestResultResponseBodyCommTest {
	s.Model = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyCommTest) SetType(v string) *DescribeNetTestResultResponseBodyCommTest {
	s.Type = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyCommTest) Validate() error {
	return dara.Validate(s)
}

type DescribeNetTestResultResponseBodyCommTestHosts struct {
	// The IP address.
	//
	// example:
	//
	// 169.253.253.15
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// i111670831721110797708
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Service name.
	//
	// example:
	//
	// VBw
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s DescribeNetTestResultResponseBodyCommTestHosts) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyCommTestHosts) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyCommTestHosts) GetIP() *string {
	return s.IP
}

func (s *DescribeNetTestResultResponseBodyCommTestHosts) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeNetTestResultResponseBodyCommTestHosts) GetServerName() *string {
	return s.ServerName
}

func (s *DescribeNetTestResultResponseBodyCommTestHosts) SetIP(v string) *DescribeNetTestResultResponseBodyCommTestHosts {
	s.IP = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyCommTestHosts) SetResourceId(v string) *DescribeNetTestResultResponseBodyCommTestHosts {
	s.ResourceId = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyCommTestHosts) SetServerName(v string) *DescribeNetTestResultResponseBodyCommTestHosts {
	s.ServerName = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyCommTestHosts) Validate() error {
	return dara.Validate(s)
}

type DescribeNetTestResultResponseBodyDelayTest struct {
	// All hosts infomation
	Hosts []*DescribeNetTestResultResponseBodyDelayTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
}

func (s DescribeNetTestResultResponseBodyDelayTest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyDelayTest) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyDelayTest) GetHosts() []*DescribeNetTestResultResponseBodyDelayTestHosts {
	return s.Hosts
}

func (s *DescribeNetTestResultResponseBodyDelayTest) SetHosts(v []*DescribeNetTestResultResponseBodyDelayTestHosts) *DescribeNetTestResultResponseBodyDelayTest {
	s.Hosts = v
	return s
}

func (s *DescribeNetTestResultResponseBodyDelayTest) Validate() error {
	return dara.Validate(s)
}

type DescribeNetTestResultResponseBodyDelayTestHosts struct {
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
	// 125.210.225.48
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// e01-cn-bcd3u1aee06
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Service name.
	//
	// example:
	//
	// NQU
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s DescribeNetTestResultResponseBodyDelayTestHosts) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyDelayTestHosts) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) GetBond() *string {
	return s.Bond
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) GetIP() *string {
	return s.IP
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) GetServerName() *string {
	return s.ServerName
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) SetBond(v string) *DescribeNetTestResultResponseBodyDelayTestHosts {
	s.Bond = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) SetIP(v string) *DescribeNetTestResultResponseBodyDelayTestHosts {
	s.IP = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) SetResourceId(v string) *DescribeNetTestResultResponseBodyDelayTestHosts {
	s.ResourceId = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) SetServerName(v string) *DescribeNetTestResultResponseBodyDelayTestHosts {
	s.ServerName = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) Validate() error {
	return dara.Validate(s)
}

type DescribeNetTestResultResponseBodyTrafficTest struct {
	// All clients information
	Clients []*DescribeNetTestResultResponseBodyTrafficTestClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	// Call duration, in seconds.
	//
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// When the protocol is RDMA, fill in True/False,
	//
	// when the protocol is TCP, this field is empty.
	//
	// example:
	//
	// False
	GDR *string `json:"GDR,omitempty" xml:"GDR,omitempty"`
	// Network protocol, either RDMA or TCP.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// When the protocol is TCP, fill in the number of concurrent connections; when the protocol is RDMA, fill in the configured QP value.
	//
	// example:
	//
	// 1
	QP *int64 `json:"QP,omitempty" xml:"QP,omitempty"`
	// Servers infomation.
	Servers []*DescribeNetTestResultResponseBodyTrafficTestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	// Traffic model, either MTON or Fullmesh.
	//
	// example:
	//
	// Fullmesh
	TrafficModel *string `json:"TrafficModel,omitempty" xml:"TrafficModel,omitempty"`
}

func (s DescribeNetTestResultResponseBodyTrafficTest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyTrafficTest) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) GetClients() []*DescribeNetTestResultResponseBodyTrafficTestClients {
	return s.Clients
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) GetGDR() *string {
	return s.GDR
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) GetQP() *int64 {
	return s.QP
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) GetServers() []*DescribeNetTestResultResponseBodyTrafficTestServers {
	return s.Servers
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) GetTrafficModel() *string {
	return s.TrafficModel
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetClients(v []*DescribeNetTestResultResponseBodyTrafficTestClients) *DescribeNetTestResultResponseBodyTrafficTest {
	s.Clients = v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetDuration(v int64) *DescribeNetTestResultResponseBodyTrafficTest {
	s.Duration = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetGDR(v string) *DescribeNetTestResultResponseBodyTrafficTest {
	s.GDR = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetProtocol(v string) *DescribeNetTestResultResponseBodyTrafficTest {
	s.Protocol = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetQP(v int64) *DescribeNetTestResultResponseBodyTrafficTest {
	s.QP = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetServers(v []*DescribeNetTestResultResponseBodyTrafficTestServers) *DescribeNetTestResultResponseBodyTrafficTest {
	s.Servers = v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetTrafficModel(v string) *DescribeNetTestResultResponseBodyTrafficTest {
	s.TrafficModel = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) Validate() error {
	return dara.Validate(s)
}

type DescribeNetTestResultResponseBodyTrafficTestClients struct {
	// Network interface bond port
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// IP address.
	//
	// example:
	//
	// 192.168.1.1
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
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

func (s DescribeNetTestResultResponseBodyTrafficTestClients) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyTrafficTestClients) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) GetBond() *string {
	return s.Bond
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) GetIP() *string {
	return s.IP
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) GetServerName() *string {
	return s.ServerName
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) SetBond(v string) *DescribeNetTestResultResponseBodyTrafficTestClients {
	s.Bond = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) SetIP(v string) *DescribeNetTestResultResponseBodyTrafficTestClients {
	s.IP = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) SetResourceId(v string) *DescribeNetTestResultResponseBodyTrafficTestClients {
	s.ResourceId = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) SetServerName(v string) *DescribeNetTestResultResponseBodyTrafficTestClients {
	s.ServerName = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) Validate() error {
	return dara.Validate(s)
}

type DescribeNetTestResultResponseBodyTrafficTestServers struct {
	// Network interface bond port
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 47.121.110.190
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// e01-cn-wwo3etaqu0b
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Service name.
	//
	// example:
	//
	// xMv
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s DescribeNetTestResultResponseBodyTrafficTestServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyTrafficTestServers) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) GetBond() *string {
	return s.Bond
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) GetIP() *string {
	return s.IP
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) GetServerName() *string {
	return s.ServerName
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) SetBond(v string) *DescribeNetTestResultResponseBodyTrafficTestServers {
	s.Bond = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) SetIP(v string) *DescribeNetTestResultResponseBodyTrafficTestServers {
	s.IP = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) SetResourceId(v string) *DescribeNetTestResultResponseBodyTrafficTestServers {
	s.ResourceId = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) SetServerName(v string) *DescribeNetTestResultResponseBodyTrafficTestServers {
	s.ServerName = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) Validate() error {
	return dara.Validate(s)
}
