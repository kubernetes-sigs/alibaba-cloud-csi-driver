// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceMonitorDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorData(v *DescribeInstanceMonitorDataResponseBodyMonitorData) *DescribeInstanceMonitorDataResponseBody
	GetMonitorData() *DescribeInstanceMonitorDataResponseBodyMonitorData
	SetRequestId(v string) *DescribeInstanceMonitorDataResponseBody
	GetRequestId() *string
}

type DescribeInstanceMonitorDataResponseBody struct {
	// The monitoring data of the instance.
	MonitorData *DescribeInstanceMonitorDataResponseBodyMonitorData `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceMonitorDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMonitorDataResponseBody) GetMonitorData() *DescribeInstanceMonitorDataResponseBodyMonitorData {
	return s.MonitorData
}

func (s *DescribeInstanceMonitorDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceMonitorDataResponseBody) SetMonitorData(v *DescribeInstanceMonitorDataResponseBodyMonitorData) *DescribeInstanceMonitorDataResponseBody {
	s.MonitorData = v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBody) SetRequestId(v string) *DescribeInstanceMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBody) Validate() error {
	if s.MonitorData != nil {
		if err := s.MonitorData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceMonitorDataResponseBodyMonitorData struct {
	InstanceMonitorData []*DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData `json:"InstanceMonitorData,omitempty" xml:"InstanceMonitorData,omitempty" type:"Repeated"`
}

func (s DescribeInstanceMonitorDataResponseBodyMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMonitorDataResponseBodyMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorData) GetInstanceMonitorData() []*DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	return s.InstanceMonitorData
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorData) SetInstanceMonitorData(v []*DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) *DescribeInstanceMonitorDataResponseBodyMonitorData {
	s.InstanceMonitorData = v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorData) Validate() error {
	if s.InstanceMonitorData != nil {
		for _, item := range s.InstanceMonitorData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData struct {
	// The read bandwidth of the cloud disks (system disk and data disks). Unit: Byte/s.
	//
	// example:
	//
	// 1000
	BPSRead *int32 `json:"BPSRead,omitempty" xml:"BPSRead,omitempty"`
	// The write bandwidth of the cloud disks (system disk and data disks). Unit: Byte/s.
	//
	// example:
	//
	// 13585
	BPSWrite *int32 `json:"BPSWrite,omitempty" xml:"BPSWrite,omitempty"`
	// The vCPU utilization of the instance. Unit: percent (%).
	//
	// example:
	//
	// 2
	CPU *int32 `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// The overdrawn CPU credits of the burstable instance.
	//
	// example:
	//
	// 0.4
	CPUAdvanceCreditBalance *float32 `json:"CPUAdvanceCreditBalance,omitempty" xml:"CPUAdvanceCreditBalance,omitempty"`
	// The total number of CPU credits of the burstable instance.
	//
	// example:
	//
	// 120
	CPUCreditBalance *float32 `json:"CPUCreditBalance,omitempty" xml:"CPUCreditBalance,omitempty"`
	// The number of CPU credits consumed by the burstable instance.
	//
	// example:
	//
	// 30
	CPUCreditUsage *float32 `json:"CPUCreditUsage,omitempty" xml:"CPUCreditUsage,omitempty"`
	// The unpaid overdrawn CPU credits.
	//
	// example:
	//
	// 0.5
	CPUNotpaidSurplusCreditUsage *float32 `json:"CPUNotpaidSurplusCreditUsage,omitempty" xml:"CPUNotpaidSurplusCreditUsage,omitempty"`
	// The number of read I/O operations per second on the cloud disks (system disk and data disks).
	//
	// example:
	//
	// 1000
	IOPSRead *int32 `json:"IOPSRead,omitempty" xml:"IOPSRead,omitempty"`
	// The number of write I/O operations per second on the cloud disks (system disk and data disks).
	//
	// example:
	//
	// 200
	IOPSWrite *int32 `json:"IOPSWrite,omitempty" xml:"IOPSWrite,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-bp1a36962lrhj4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The public bandwidth of the instance. Unit: Kbit/s.
	//
	// example:
	//
	// 10
	InternetBandwidth *int32 `json:"InternetBandwidth,omitempty" xml:"InternetBandwidth,omitempty"`
	// The Internet traffic received by the instance during the period that is specified by the `Period` parameter. The period starts from the point in time that is specified by the `TimeStamp` parameter. Unit: Kbit.
	//
	// example:
	//
	// 122
	InternetRX *int32 `json:"InternetRX,omitempty" xml:"InternetRX,omitempty"`
	// The Internet traffic sent by the instance during the period that is specified by the `Period` parameter. The period starts from the point in time that is specified by the `TimeStamp` parameter. Unit: Kbit.
	//
	// example:
	//
	// 343
	InternetTX *int32 `json:"InternetTX,omitempty" xml:"InternetTX,omitempty"`
	// The internal bandwidth of the instance. Unit: Kbit/s.
	//
	// example:
	//
	// 10
	IntranetBandwidth *int32 `json:"IntranetBandwidth,omitempty" xml:"IntranetBandwidth,omitempty"`
	// The internal data traffic received by the instance during the period that is specified by the `Period` parameter. The period starts from the point in time that is specified by the `TimeStamp` parameter. Unit: Kbit.
	//
	// example:
	//
	// 122
	IntranetRX *int32 `json:"IntranetRX,omitempty" xml:"IntranetRX,omitempty"`
	// The internal data traffic sent by the instance during the period that is specified by the `Period` parameter. The period starts from the point in time that is specified by the `TimeStamp` parameter. Unit: Kbit.
	//
	// example:
	//
	// 343
	IntranetTX *int32 `json:"IntranetTX,omitempty" xml:"IntranetTX,omitempty"`
	// The timestamp of the monitoring data.
	//
	// example:
	//
	// 2014-10-30T05:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetBPSRead() *int32 {
	return s.BPSRead
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetBPSWrite() *int32 {
	return s.BPSWrite
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetCPU() *int32 {
	return s.CPU
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetCPUAdvanceCreditBalance() *float32 {
	return s.CPUAdvanceCreditBalance
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetCPUCreditBalance() *float32 {
	return s.CPUCreditBalance
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetCPUCreditUsage() *float32 {
	return s.CPUCreditUsage
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetCPUNotpaidSurplusCreditUsage() *float32 {
	return s.CPUNotpaidSurplusCreditUsage
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetIOPSRead() *int32 {
	return s.IOPSRead
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetIOPSWrite() *int32 {
	return s.IOPSWrite
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetInternetBandwidth() *int32 {
	return s.InternetBandwidth
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetInternetRX() *int32 {
	return s.InternetRX
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetInternetTX() *int32 {
	return s.InternetTX
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetIntranetBandwidth() *int32 {
	return s.IntranetBandwidth
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetIntranetRX() *int32 {
	return s.IntranetRX
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetIntranetTX() *int32 {
	return s.IntranetTX
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetBPSRead(v int32) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.BPSRead = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetBPSWrite(v int32) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.BPSWrite = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetCPU(v int32) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.CPU = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetCPUAdvanceCreditBalance(v float32) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.CPUAdvanceCreditBalance = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetCPUCreditBalance(v float32) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.CPUCreditBalance = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetCPUCreditUsage(v float32) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.CPUCreditUsage = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetCPUNotpaidSurplusCreditUsage(v float32) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.CPUNotpaidSurplusCreditUsage = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetIOPSRead(v int32) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.IOPSRead = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetIOPSWrite(v int32) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.IOPSWrite = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetInstanceId(v string) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetInternetBandwidth(v int32) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.InternetBandwidth = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetInternetRX(v int32) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.InternetRX = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetInternetTX(v int32) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.InternetTX = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetIntranetBandwidth(v int32) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.IntranetBandwidth = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetIntranetRX(v int32) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.IntranetRX = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetIntranetTX(v int32) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.IntranetTX = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) SetTimeStamp(v string) *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeInstanceMonitorDataResponseBodyMonitorDataInstanceMonitorData) Validate() error {
	return dara.Validate(s)
}
