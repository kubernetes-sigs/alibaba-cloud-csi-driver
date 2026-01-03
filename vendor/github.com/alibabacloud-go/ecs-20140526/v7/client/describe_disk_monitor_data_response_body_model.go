// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskMonitorDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorData(v *DescribeDiskMonitorDataResponseBodyMonitorData) *DescribeDiskMonitorDataResponseBody
	GetMonitorData() *DescribeDiskMonitorDataResponseBodyMonitorData
	SetRequestId(v string) *DescribeDiskMonitorDataResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDiskMonitorDataResponseBody
	GetTotalCount() *int32
}

type DescribeDiskMonitorDataResponseBody struct {
	// The monitoring data of the disk.
	MonitorData *DescribeDiskMonitorDataResponseBodyMonitorData `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of monitoring data entries returned.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDiskMonitorDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataResponseBody) GetMonitorData() *DescribeDiskMonitorDataResponseBodyMonitorData {
	return s.MonitorData
}

func (s *DescribeDiskMonitorDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiskMonitorDataResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDiskMonitorDataResponseBody) SetMonitorData(v *DescribeDiskMonitorDataResponseBodyMonitorData) *DescribeDiskMonitorDataResponseBody {
	s.MonitorData = v
	return s
}

func (s *DescribeDiskMonitorDataResponseBody) SetRequestId(v string) *DescribeDiskMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBody) SetTotalCount(v int32) *DescribeDiskMonitorDataResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBody) Validate() error {
	if s.MonitorData != nil {
		if err := s.MonitorData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDiskMonitorDataResponseBodyMonitorData struct {
	DiskMonitorData []*DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData `json:"DiskMonitorData,omitempty" xml:"DiskMonitorData,omitempty" type:"Repeated"`
}

func (s DescribeDiskMonitorDataResponseBodyMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskMonitorDataResponseBodyMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) GetDiskMonitorData() []*DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData {
	return s.DiskMonitorData
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) SetDiskMonitorData(v []*DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) *DescribeDiskMonitorDataResponseBodyMonitorData {
	s.DiskMonitorData = v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorData) Validate() error {
	if s.DiskMonitorData != nil {
		for _, item := range s.DiskMonitorData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData struct {
	// The read bandwidth of the disk. Unit: byte/s.
	//
	// example:
	//
	// 0
	BPSRead *int32 `json:"BPSRead,omitempty" xml:"BPSRead,omitempty"`
	// The total read and write bandwidth of the disk. Unit: byte/s.
	//
	// example:
	//
	// 204
	BPSTotal *int32 `json:"BPSTotal,omitempty" xml:"BPSTotal,omitempty"`
	// The write bandwidth of the disk. Unit: byte/s.
	//
	// example:
	//
	// 204
	BPSWrite *int32 `json:"BPSWrite,omitempty" xml:"BPSWrite,omitempty"`
	// The ID of the disk.
	//
	// example:
	//
	// d-bp1bq5g3dxxo1x4o****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The number of read I/O operations per second on the disk.
	//
	// example:
	//
	// 0
	IOPSRead *int32 `json:"IOPSRead,omitempty" xml:"IOPSRead,omitempty"`
	// The total number of read and write I/O operations per second on the disk.
	//
	// example:
	//
	// 0
	IOPSTotal *int32 `json:"IOPSTotal,omitempty" xml:"IOPSTotal,omitempty"`
	// The number of write I/O operations per second on the disk.
	//
	// example:
	//
	// 0
	IOPSWrite *int32 `json:"IOPSWrite,omitempty" xml:"IOPSWrite,omitempty"`
	// The read latency of the disk. Unit: microseconds.
	//
	// example:
	//
	// 0
	LatencyRead *int32 `json:"LatencyRead,omitempty" xml:"LatencyRead,omitempty"`
	// The write latency of the disk. Unit: microseconds.
	//
	// example:
	//
	// 0
	LatencyWrite *int32 `json:"LatencyWrite,omitempty" xml:"LatencyWrite,omitempty"`
	// The timestamp of the monitoring data. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2014-07-23T12:07:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) GetBPSRead() *int32 {
	return s.BPSRead
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) GetBPSTotal() *int32 {
	return s.BPSTotal
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) GetBPSWrite() *int32 {
	return s.BPSWrite
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) GetIOPSRead() *int32 {
	return s.IOPSRead
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) GetIOPSTotal() *int32 {
	return s.IOPSTotal
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) GetIOPSWrite() *int32 {
	return s.IOPSWrite
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) GetLatencyRead() *int32 {
	return s.LatencyRead
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) GetLatencyWrite() *int32 {
	return s.LatencyWrite
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) SetBPSRead(v int32) *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData {
	s.BPSRead = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) SetBPSTotal(v int32) *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData {
	s.BPSTotal = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) SetBPSWrite(v int32) *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData {
	s.BPSWrite = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) SetDiskId(v string) *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData {
	s.DiskId = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) SetIOPSRead(v int32) *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData {
	s.IOPSRead = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) SetIOPSTotal(v int32) *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData {
	s.IOPSTotal = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) SetIOPSWrite(v int32) *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData {
	s.IOPSWrite = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) SetLatencyRead(v int32) *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData {
	s.LatencyRead = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) SetLatencyWrite(v int32) *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData {
	s.LatencyWrite = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) SetTimeStamp(v string) *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDiskMonitorDataResponseBodyMonitorDataDiskMonitorData) Validate() error {
	return dara.Validate(s)
}
