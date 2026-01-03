// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeregisterManagedInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstance(v *DeregisterManagedInstanceResponseBodyInstance) *DeregisterManagedInstanceResponseBody
	GetInstance() *DeregisterManagedInstanceResponseBodyInstance
	SetRequestId(v string) *DeregisterManagedInstanceResponseBody
	GetRequestId() *string
}

type DeregisterManagedInstanceResponseBody struct {
	// Details of the managed instances.
	Instance *DeregisterManagedInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4ECEEE12-56F1-4FBC-9AB1-890F74942176
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeregisterManagedInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeregisterManagedInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeregisterManagedInstanceResponseBody) GetInstance() *DeregisterManagedInstanceResponseBodyInstance {
	return s.Instance
}

func (s *DeregisterManagedInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeregisterManagedInstanceResponseBody) SetInstance(v *DeregisterManagedInstanceResponseBodyInstance) *DeregisterManagedInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *DeregisterManagedInstanceResponseBody) SetRequestId(v string) *DeregisterManagedInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeregisterManagedInstanceResponseBody) Validate() error {
	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeregisterManagedInstanceResponseBodyInstance struct {
	// The activation code ID.
	//
	// example:
	//
	// 4ECEEE12-56F1-4FBC-9AB1-890F7494****
	ActivationId *string `json:"ActivationId,omitempty" xml:"ActivationId,omitempty"`
	// The version number of Cloud Assistant Agent.
	//
	// example:
	//
	// 2.2.0.102
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// The hostname of the managed instance.
	//
	// example:
	//
	// test-Hostname
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The managed instance ID.
	//
	// example:
	//
	// mi-hz01axdfas****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the managed instance.
	//
	// example:
	//
	// test-InstanceName-001
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the managed instance.
	//
	// example:
	//
	// ``47.8.**.**``
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The internal IP address of the managed instance.
	//
	// example:
	//
	// ``10.0.**.**``
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The number of times that Cloud Assistant tasks were executed on the managed instance.
	//
	// example:
	//
	// 2
	InvocationCount *int64 `json:"InvocationCount,omitempty" xml:"InvocationCount,omitempty"`
	// The time when the Cloud Assistant task was last executed.
	//
	// example:
	//
	// 2021-01-20T09:00:40Z
	LastInvokedTime *string `json:"LastInvokedTime,omitempty" xml:"LastInvokedTime,omitempty"`
	// The machine code of the managed instance.
	//
	// example:
	//
	// e03231b37ab14e53b5795ad625fc****
	MachineId *string `json:"MachineId,omitempty" xml:"MachineId,omitempty"`
	// The operating system type of the managed instance.
	//
	// example:
	//
	// linux
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The version information about the operating system.
	//
	// example:
	//
	// Linux_#38~18.04.1-Ubuntu SMP Wed Jan 6 18:26:30 UTC 2021_x86_64
	OsVersion *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	// The time when the managed instance was registered.
	//
	// example:
	//
	// 2021-01-20T08:57:56Z
	RegistrationTime *string `json:"RegistrationTime,omitempty" xml:"RegistrationTime,omitempty"`
	// The ID of the resource group to which the managed instance belongs.
	//
	// example:
	//
	// rg-123******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeregisterManagedInstanceResponseBodyInstance) String() string {
	return dara.Prettify(s)
}

func (s DeregisterManagedInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *DeregisterManagedInstanceResponseBodyInstance) GetActivationId() *string {
	return s.ActivationId
}

func (s *DeregisterManagedInstanceResponseBodyInstance) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *DeregisterManagedInstanceResponseBodyInstance) GetHostname() *string {
	return s.Hostname
}

func (s *DeregisterManagedInstanceResponseBodyInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeregisterManagedInstanceResponseBodyInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DeregisterManagedInstanceResponseBodyInstance) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DeregisterManagedInstanceResponseBodyInstance) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DeregisterManagedInstanceResponseBodyInstance) GetInvocationCount() *int64 {
	return s.InvocationCount
}

func (s *DeregisterManagedInstanceResponseBodyInstance) GetLastInvokedTime() *string {
	return s.LastInvokedTime
}

func (s *DeregisterManagedInstanceResponseBodyInstance) GetMachineId() *string {
	return s.MachineId
}

func (s *DeregisterManagedInstanceResponseBodyInstance) GetOsType() *string {
	return s.OsType
}

func (s *DeregisterManagedInstanceResponseBodyInstance) GetOsVersion() *string {
	return s.OsVersion
}

func (s *DeregisterManagedInstanceResponseBodyInstance) GetRegistrationTime() *string {
	return s.RegistrationTime
}

func (s *DeregisterManagedInstanceResponseBodyInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeregisterManagedInstanceResponseBodyInstance) SetActivationId(v string) *DeregisterManagedInstanceResponseBodyInstance {
	s.ActivationId = &v
	return s
}

func (s *DeregisterManagedInstanceResponseBodyInstance) SetAgentVersion(v string) *DeregisterManagedInstanceResponseBodyInstance {
	s.AgentVersion = &v
	return s
}

func (s *DeregisterManagedInstanceResponseBodyInstance) SetHostname(v string) *DeregisterManagedInstanceResponseBodyInstance {
	s.Hostname = &v
	return s
}

func (s *DeregisterManagedInstanceResponseBodyInstance) SetInstanceId(v string) *DeregisterManagedInstanceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *DeregisterManagedInstanceResponseBodyInstance) SetInstanceName(v string) *DeregisterManagedInstanceResponseBodyInstance {
	s.InstanceName = &v
	return s
}

func (s *DeregisterManagedInstanceResponseBodyInstance) SetInternetIp(v string) *DeregisterManagedInstanceResponseBodyInstance {
	s.InternetIp = &v
	return s
}

func (s *DeregisterManagedInstanceResponseBodyInstance) SetIntranetIp(v string) *DeregisterManagedInstanceResponseBodyInstance {
	s.IntranetIp = &v
	return s
}

func (s *DeregisterManagedInstanceResponseBodyInstance) SetInvocationCount(v int64) *DeregisterManagedInstanceResponseBodyInstance {
	s.InvocationCount = &v
	return s
}

func (s *DeregisterManagedInstanceResponseBodyInstance) SetLastInvokedTime(v string) *DeregisterManagedInstanceResponseBodyInstance {
	s.LastInvokedTime = &v
	return s
}

func (s *DeregisterManagedInstanceResponseBodyInstance) SetMachineId(v string) *DeregisterManagedInstanceResponseBodyInstance {
	s.MachineId = &v
	return s
}

func (s *DeregisterManagedInstanceResponseBodyInstance) SetOsType(v string) *DeregisterManagedInstanceResponseBodyInstance {
	s.OsType = &v
	return s
}

func (s *DeregisterManagedInstanceResponseBodyInstance) SetOsVersion(v string) *DeregisterManagedInstanceResponseBodyInstance {
	s.OsVersion = &v
	return s
}

func (s *DeregisterManagedInstanceResponseBodyInstance) SetRegistrationTime(v string) *DeregisterManagedInstanceResponseBodyInstance {
	s.RegistrationTime = &v
	return s
}

func (s *DeregisterManagedInstanceResponseBodyInstance) SetResourceGroupId(v string) *DeregisterManagedInstanceResponseBodyInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DeregisterManagedInstanceResponseBodyInstance) Validate() error {
	return dara.Validate(s)
}
