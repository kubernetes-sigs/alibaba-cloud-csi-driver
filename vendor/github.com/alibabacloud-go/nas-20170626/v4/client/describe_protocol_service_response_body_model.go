// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProtocolServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeProtocolServiceResponseBody
	GetNextToken() *string
	SetProtocolServices(v []*DescribeProtocolServiceResponseBodyProtocolServices) *DescribeProtocolServiceResponseBody
	GetProtocolServices() []*DescribeProtocolServiceResponseBodyProtocolServices
	SetRequestId(v string) *DescribeProtocolServiceResponseBody
	GetRequestId() *string
}

type DescribeProtocolServiceResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// aBcdeg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about protocol services.
	ProtocolServices []*DescribeProtocolServiceResponseBodyProtocolServices `json:"ProtocolServices,omitempty" xml:"ProtocolServices,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProtocolServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProtocolServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProtocolServiceResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeProtocolServiceResponseBody) GetProtocolServices() []*DescribeProtocolServiceResponseBodyProtocolServices {
	return s.ProtocolServices
}

func (s *DescribeProtocolServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProtocolServiceResponseBody) SetNextToken(v string) *DescribeProtocolServiceResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeProtocolServiceResponseBody) SetProtocolServices(v []*DescribeProtocolServiceResponseBodyProtocolServices) *DescribeProtocolServiceResponseBody {
	s.ProtocolServices = v
	return s
}

func (s *DescribeProtocolServiceResponseBody) SetRequestId(v string) *DescribeProtocolServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProtocolServiceResponseBody) Validate() error {
	if s.ProtocolServices != nil {
		for _, item := range s.ProtocolServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeProtocolServiceResponseBodyProtocolServices struct {
	// The time when the protocol service was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-12-12T07:28:38Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the protocol service.
	//
	// Limits:
	//
	// 	- The description must be 2 to 128 characters in length.
	//
	// 	- The description must start with a letter and cannot start with `http://` or `https://`.
	//
	// 	- The description can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the file system.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The base throughput of the protocol service. Unit: MB/s.
	//
	// example:
	//
	// 100
	InstanceBaseThroughput *int32 `json:"InstanceBaseThroughput,omitempty" xml:"InstanceBaseThroughput,omitempty"`
	// The burst throughput of the protocol service. Unit: MB/s.
	//
	// example:
	//
	// 100
	InstanceBurstThroughput *int32 `json:"InstanceBurstThroughput,omitempty" xml:"InstanceBurstThroughput,omitempty"`
	// The memory cache size of the protocol service. Unit: GiB.
	//
	// example:
	//
	// 0
	InstanceRAM *int32 `json:"InstanceRAM,omitempty" xml:"InstanceRAM,omitempty"`
	// The time when the protocol service was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-12-12T07:28:38Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The total number of CPFS directories and filesets exported in the protocol service.
	//
	// example:
	//
	// 5
	MountTargetCount *int32 `json:"MountTargetCount,omitempty" xml:"MountTargetCount,omitempty"`
	// The ID of the protocol service.
	//
	// example:
	//
	// ptc-197ed6a00f2b****
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
	// The specification of the protocol service.
	//
	// 	- Valid value: General.
	//
	// 	- Default value: General.
	//
	// example:
	//
	// General
	ProtocolSpec *string `json:"ProtocolSpec,omitempty" xml:"ProtocolSpec,omitempty"`
	// The throughput of the protocol service. Unit: MB/s.
	//
	// example:
	//
	// 500
	ProtocolThroughput *int32 `json:"ProtocolThroughput,omitempty" xml:"ProtocolThroughput,omitempty"`
	// The protocol type supported by the protocol service.
	//
	// Valid values:
	//
	// 	- NFS: The protocol service supports access over the Network File System (NFS) protocol.
	//
	// example:
	//
	// NFS
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The status of the protocol service.
	//
	// Valid values:
	//
	// 	- Creating: The protocol service is being created.
	//
	// 	- Starting: The protocol service is being started.
	//
	// 	- Running: The protocol service is running.
	//
	// 	- Updating: The protocol service is being updated.
	//
	// 	- Deleting: The protocol service is being deleted.
	//
	// 	- Stopping: The protocol service is being stopped.
	//
	// 	- Stopped: The protocol service is stopped.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeProtocolServiceResponseBodyProtocolServices) String() string {
	return dara.Prettify(s)
}

func (s DescribeProtocolServiceResponseBodyProtocolServices) GoString() string {
	return s.String()
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) GetDescription() *string {
	return s.Description
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) GetInstanceBaseThroughput() *int32 {
	return s.InstanceBaseThroughput
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) GetInstanceBurstThroughput() *int32 {
	return s.InstanceBurstThroughput
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) GetInstanceRAM() *int32 {
	return s.InstanceRAM
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) GetMountTargetCount() *int32 {
	return s.MountTargetCount
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) GetProtocolServiceId() *string {
	return s.ProtocolServiceId
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) GetProtocolSpec() *string {
	return s.ProtocolSpec
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) GetProtocolThroughput() *int32 {
	return s.ProtocolThroughput
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) GetStatus() *string {
	return s.Status
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetCreateTime(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.CreateTime = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetDescription(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.Description = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetFileSystemId(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.FileSystemId = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetInstanceBaseThroughput(v int32) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.InstanceBaseThroughput = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetInstanceBurstThroughput(v int32) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.InstanceBurstThroughput = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetInstanceRAM(v int32) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.InstanceRAM = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetModifyTime(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.ModifyTime = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetMountTargetCount(v int32) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.MountTargetCount = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetProtocolServiceId(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.ProtocolServiceId = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetProtocolSpec(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.ProtocolSpec = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetProtocolThroughput(v int32) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.ProtocolThroughput = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetProtocolType(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.ProtocolType = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) SetStatus(v string) *DescribeProtocolServiceResponseBodyProtocolServices {
	s.Status = &v
	return s
}

func (s *DescribeProtocolServiceResponseBodyProtocolServices) Validate() error {
	return dara.Validate(s)
}
