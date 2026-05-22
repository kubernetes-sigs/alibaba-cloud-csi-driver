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
	BPSRead      *int32  `json:"BPSRead,omitempty" xml:"BPSRead,omitempty"`
	BPSTotal     *int32  `json:"BPSTotal,omitempty" xml:"BPSTotal,omitempty"`
	BPSWrite     *int32  `json:"BPSWrite,omitempty" xml:"BPSWrite,omitempty"`
	DiskId       *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	IOPSRead     *int32  `json:"IOPSRead,omitempty" xml:"IOPSRead,omitempty"`
	IOPSTotal    *int32  `json:"IOPSTotal,omitempty" xml:"IOPSTotal,omitempty"`
	IOPSWrite    *int32  `json:"IOPSWrite,omitempty" xml:"IOPSWrite,omitempty"`
	LatencyRead  *int32  `json:"LatencyRead,omitempty" xml:"LatencyRead,omitempty"`
	LatencyWrite *int32  `json:"LatencyWrite,omitempty" xml:"LatencyWrite,omitempty"`
	TimeStamp    *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
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
