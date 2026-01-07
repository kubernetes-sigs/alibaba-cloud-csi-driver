// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeManagedInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*DescribeManagedInstancesResponseBodyInstances) *DescribeManagedInstancesResponseBody
	GetInstances() []*DescribeManagedInstancesResponseBodyInstances
	SetNextToken(v string) *DescribeManagedInstancesResponseBody
	GetNextToken() *string
	SetPageNumber(v int64) *DescribeManagedInstancesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeManagedInstancesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeManagedInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeManagedInstancesResponseBody
	GetTotalCount() *int64
}

type DescribeManagedInstancesResponseBody struct {
	// The queried managed instances.
	Instances []*DescribeManagedInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 77115469-F2C5-4ECA-94F7-FA04F2FD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of queried managed instances.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeManagedInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeManagedInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeManagedInstancesResponseBody) GetInstances() []*DescribeManagedInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeManagedInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeManagedInstancesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeManagedInstancesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeManagedInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeManagedInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeManagedInstancesResponseBody) SetInstances(v []*DescribeManagedInstancesResponseBodyInstances) *DescribeManagedInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeManagedInstancesResponseBody) SetNextToken(v string) *DescribeManagedInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeManagedInstancesResponseBody) SetPageNumber(v int64) *DescribeManagedInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeManagedInstancesResponseBody) SetPageSize(v int64) *DescribeManagedInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeManagedInstancesResponseBody) SetRequestId(v string) *DescribeManagedInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeManagedInstancesResponseBody) SetTotalCount(v int64) *DescribeManagedInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeManagedInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeManagedInstancesResponseBodyInstances struct {
	// The ID of the activation code.
	//
	// example:
	//
	// 3704F543-F768-43FA-9864-897F75B3****
	ActivationId *string `json:"ActivationId,omitempty" xml:"ActivationId,omitempty"`
	// The version number of Cloud Assistant Agent.
	//
	// example:
	//
	// 2.2.0.102
	AgentVersion *string `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// Indicates whether the managed instance is connected. Valid values:
	//
	// 	- true: The managed instance is connected. You can manage the instance by using Cloud Assistant.
	//
	// 	- false: The managed instance is not connected. The managed instance may be down or Cloud Assistant Agent may be incorrectly installed.
	//
	// example:
	//
	// true
	Connected *bool `json:"Connected,omitempty" xml:"Connected,omitempty"`
	// The hostname of the managed instance.
	//
	// example:
	//
	// demo
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The ID of the managed instance.
	//
	// example:
	//
	// mi-hz018jrc1o0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the managed instance.
	//
	// example:
	//
	// webAPP-linux-01
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the managed instance.
	//
	// example:
	//
	// ``40.65.**.**``
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
	// 1
	InvocationCount *int64 `json:"InvocationCount,omitempty" xml:"InvocationCount,omitempty"`
	// The time when the last Cloud Assistant task was executed.
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
	// Linux
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The version information of the operating system.
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
	// The tags of the managed instance.
	Tags []*DescribeManagedInstancesResponseBodyInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeManagedInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeManagedInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeManagedInstancesResponseBodyInstances) GetActivationId() *string {
	return s.ActivationId
}

func (s *DescribeManagedInstancesResponseBodyInstances) GetAgentVersion() *string {
	return s.AgentVersion
}

func (s *DescribeManagedInstancesResponseBodyInstances) GetConnected() *bool {
	return s.Connected
}

func (s *DescribeManagedInstancesResponseBodyInstances) GetHostname() *string {
	return s.Hostname
}

func (s *DescribeManagedInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeManagedInstancesResponseBodyInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeManagedInstancesResponseBodyInstances) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeManagedInstancesResponseBodyInstances) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeManagedInstancesResponseBodyInstances) GetInvocationCount() *int64 {
	return s.InvocationCount
}

func (s *DescribeManagedInstancesResponseBodyInstances) GetLastInvokedTime() *string {
	return s.LastInvokedTime
}

func (s *DescribeManagedInstancesResponseBodyInstances) GetMachineId() *string {
	return s.MachineId
}

func (s *DescribeManagedInstancesResponseBodyInstances) GetOsType() *string {
	return s.OsType
}

func (s *DescribeManagedInstancesResponseBodyInstances) GetOsVersion() *string {
	return s.OsVersion
}

func (s *DescribeManagedInstancesResponseBodyInstances) GetRegistrationTime() *string {
	return s.RegistrationTime
}

func (s *DescribeManagedInstancesResponseBodyInstances) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeManagedInstancesResponseBodyInstances) GetTags() []*DescribeManagedInstancesResponseBodyInstancesTags {
	return s.Tags
}

func (s *DescribeManagedInstancesResponseBodyInstances) SetActivationId(v string) *DescribeManagedInstancesResponseBodyInstances {
	s.ActivationId = &v
	return s
}

func (s *DescribeManagedInstancesResponseBodyInstances) SetAgentVersion(v string) *DescribeManagedInstancesResponseBodyInstances {
	s.AgentVersion = &v
	return s
}

func (s *DescribeManagedInstancesResponseBodyInstances) SetConnected(v bool) *DescribeManagedInstancesResponseBodyInstances {
	s.Connected = &v
	return s
}

func (s *DescribeManagedInstancesResponseBodyInstances) SetHostname(v string) *DescribeManagedInstancesResponseBodyInstances {
	s.Hostname = &v
	return s
}

func (s *DescribeManagedInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeManagedInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeManagedInstancesResponseBodyInstances) SetInstanceName(v string) *DescribeManagedInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeManagedInstancesResponseBodyInstances) SetInternetIp(v string) *DescribeManagedInstancesResponseBodyInstances {
	s.InternetIp = &v
	return s
}

func (s *DescribeManagedInstancesResponseBodyInstances) SetIntranetIp(v string) *DescribeManagedInstancesResponseBodyInstances {
	s.IntranetIp = &v
	return s
}

func (s *DescribeManagedInstancesResponseBodyInstances) SetInvocationCount(v int64) *DescribeManagedInstancesResponseBodyInstances {
	s.InvocationCount = &v
	return s
}

func (s *DescribeManagedInstancesResponseBodyInstances) SetLastInvokedTime(v string) *DescribeManagedInstancesResponseBodyInstances {
	s.LastInvokedTime = &v
	return s
}

func (s *DescribeManagedInstancesResponseBodyInstances) SetMachineId(v string) *DescribeManagedInstancesResponseBodyInstances {
	s.MachineId = &v
	return s
}

func (s *DescribeManagedInstancesResponseBodyInstances) SetOsType(v string) *DescribeManagedInstancesResponseBodyInstances {
	s.OsType = &v
	return s
}

func (s *DescribeManagedInstancesResponseBodyInstances) SetOsVersion(v string) *DescribeManagedInstancesResponseBodyInstances {
	s.OsVersion = &v
	return s
}

func (s *DescribeManagedInstancesResponseBodyInstances) SetRegistrationTime(v string) *DescribeManagedInstancesResponseBodyInstances {
	s.RegistrationTime = &v
	return s
}

func (s *DescribeManagedInstancesResponseBodyInstances) SetResourceGroupId(v string) *DescribeManagedInstancesResponseBodyInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeManagedInstancesResponseBodyInstances) SetTags(v []*DescribeManagedInstancesResponseBodyInstancesTags) *DescribeManagedInstancesResponseBodyInstances {
	s.Tags = v
	return s
}

func (s *DescribeManagedInstancesResponseBodyInstances) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeManagedInstancesResponseBodyInstancesTags struct {
	// The key of tag N of the managed instance. Valid values of N: 1 to 20. The tag key cannot be an empty string.
	//
	// If a single tag is specified to query resources, up to 1,000 resources that have this tag added are returned. If multiple tags are specified to query resources, up to 1,000 resources that have all these tags added are returned. To query more than 1,000 resources that have the specified tags, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of tag N of the managed instance. Valid values of N: 1 to 20. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeManagedInstancesResponseBodyInstancesTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeManagedInstancesResponseBodyInstancesTags) GoString() string {
	return s.String()
}

func (s *DescribeManagedInstancesResponseBodyInstancesTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeManagedInstancesResponseBodyInstancesTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeManagedInstancesResponseBodyInstancesTags) SetTagKey(v string) *DescribeManagedInstancesResponseBodyInstancesTags {
	s.TagKey = &v
	return s
}

func (s *DescribeManagedInstancesResponseBodyInstancesTags) SetTagValue(v string) *DescribeManagedInstancesResponseBodyInstancesTags {
	s.TagValue = &v
	return s
}

func (s *DescribeManagedInstancesResponseBodyInstancesTags) Validate() error {
	return dara.Validate(s)
}
