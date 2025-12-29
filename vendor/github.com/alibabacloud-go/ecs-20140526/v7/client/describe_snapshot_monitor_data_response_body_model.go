// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotMonitorDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorData(v *DescribeSnapshotMonitorDataResponseBodyMonitorData) *DescribeSnapshotMonitorDataResponseBody
	GetMonitorData() *DescribeSnapshotMonitorDataResponseBodyMonitorData
	SetRequestId(v string) *DescribeSnapshotMonitorDataResponseBody
	GetRequestId() *string
}

type DescribeSnapshotMonitorDataResponseBody struct {
	// The monitoring data of snapshot sizes.
	MonitorData *DescribeSnapshotMonitorDataResponseBodyMonitorData `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9F8163A8-F5DE-47A2-A572-4E062D223E09
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSnapshotMonitorDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotMonitorDataResponseBody) GetMonitorData() *DescribeSnapshotMonitorDataResponseBodyMonitorData {
	return s.MonitorData
}

func (s *DescribeSnapshotMonitorDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSnapshotMonitorDataResponseBody) SetMonitorData(v *DescribeSnapshotMonitorDataResponseBodyMonitorData) *DescribeSnapshotMonitorDataResponseBody {
	s.MonitorData = v
	return s
}

func (s *DescribeSnapshotMonitorDataResponseBody) SetRequestId(v string) *DescribeSnapshotMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotMonitorDataResponseBody) Validate() error {
	if s.MonitorData != nil {
		if err := s.MonitorData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSnapshotMonitorDataResponseBodyMonitorData struct {
	DataPoint []*DescribeSnapshotMonitorDataResponseBodyMonitorDataDataPoint `json:"DataPoint,omitempty" xml:"DataPoint,omitempty" type:"Repeated"`
}

func (s DescribeSnapshotMonitorDataResponseBodyMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotMonitorDataResponseBodyMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotMonitorDataResponseBodyMonitorData) GetDataPoint() []*DescribeSnapshotMonitorDataResponseBodyMonitorDataDataPoint {
	return s.DataPoint
}

func (s *DescribeSnapshotMonitorDataResponseBodyMonitorData) SetDataPoint(v []*DescribeSnapshotMonitorDataResponseBodyMonitorDataDataPoint) *DescribeSnapshotMonitorDataResponseBodyMonitorData {
	s.DataPoint = v
	return s
}

func (s *DescribeSnapshotMonitorDataResponseBodyMonitorData) Validate() error {
	if s.DataPoint != nil {
		for _, item := range s.DataPoint {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSnapshotMonitorDataResponseBodyMonitorDataDataPoint struct {
	// The total size of snapshots. Unit: bytes.
	//
	// example:
	//
	// 243036848128
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The timestamp that corresponds to a snapshot size.
	//
	// example:
	//
	// 2019-05-10T04:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeSnapshotMonitorDataResponseBodyMonitorDataDataPoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotMonitorDataResponseBodyMonitorDataDataPoint) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotMonitorDataResponseBodyMonitorDataDataPoint) GetSize() *int64 {
	return s.Size
}

func (s *DescribeSnapshotMonitorDataResponseBodyMonitorDataDataPoint) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeSnapshotMonitorDataResponseBodyMonitorDataDataPoint) SetSize(v int64) *DescribeSnapshotMonitorDataResponseBodyMonitorDataDataPoint {
	s.Size = &v
	return s
}

func (s *DescribeSnapshotMonitorDataResponseBodyMonitorDataDataPoint) SetTimeStamp(v string) *DescribeSnapshotMonitorDataResponseBodyMonitorDataDataPoint {
	s.TimeStamp = &v
	return s
}

func (s *DescribeSnapshotMonitorDataResponseBodyMonitorDataDataPoint) Validate() error {
	return dara.Validate(s)
}
