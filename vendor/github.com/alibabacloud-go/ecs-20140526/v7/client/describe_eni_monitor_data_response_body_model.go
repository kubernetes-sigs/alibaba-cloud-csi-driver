// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEniMonitorDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorData(v *DescribeEniMonitorDataResponseBodyMonitorData) *DescribeEniMonitorDataResponseBody
	GetMonitorData() *DescribeEniMonitorDataResponseBodyMonitorData
	SetRequestId(v string) *DescribeEniMonitorDataResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeEniMonitorDataResponseBody
	GetTotalCount() *int32
}

type DescribeEniMonitorDataResponseBody struct {
	// The monitoring data of the secondary ENI.
	MonitorData *DescribeEniMonitorDataResponseBodyMonitorData `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEniMonitorDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEniMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEniMonitorDataResponseBody) GetMonitorData() *DescribeEniMonitorDataResponseBodyMonitorData {
	return s.MonitorData
}

func (s *DescribeEniMonitorDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEniMonitorDataResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEniMonitorDataResponseBody) SetMonitorData(v *DescribeEniMonitorDataResponseBodyMonitorData) *DescribeEniMonitorDataResponseBody {
	s.MonitorData = v
	return s
}

func (s *DescribeEniMonitorDataResponseBody) SetRequestId(v string) *DescribeEniMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEniMonitorDataResponseBody) SetTotalCount(v int32) *DescribeEniMonitorDataResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEniMonitorDataResponseBody) Validate() error {
	if s.MonitorData != nil {
		if err := s.MonitorData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEniMonitorDataResponseBodyMonitorData struct {
	EniMonitorData []*DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData `json:"EniMonitorData,omitempty" xml:"EniMonitorData,omitempty" type:"Repeated"`
}

func (s DescribeEniMonitorDataResponseBodyMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEniMonitorDataResponseBodyMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeEniMonitorDataResponseBodyMonitorData) GetEniMonitorData() []*DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData {
	return s.EniMonitorData
}

func (s *DescribeEniMonitorDataResponseBodyMonitorData) SetEniMonitorData(v []*DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) *DescribeEniMonitorDataResponseBodyMonitorData {
	s.EniMonitorData = v
	return s
}

func (s *DescribeEniMonitorDataResponseBodyMonitorData) Validate() error {
	if s.EniMonitorData != nil {
		for _, item := range s.EniMonitorData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData struct {
	// The number of received packets that were dropped by the secondary ENI over the internal network.
	//
	// example:
	//
	// 0
	DropPacketRx *string `json:"DropPacketRx,omitempty" xml:"DropPacketRx,omitempty"`
	// The number of sent packets that were dropped by the secondary ENI over the internal network.
	//
	// example:
	//
	// 0
	DropPacketTx *string `json:"DropPacketTx,omitempty" xml:"DropPacketTx,omitempty"`
	// The ID of the secondary ENI.
	//
	// example:
	//
	// eni-bp19da36d6xdwey****
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The average rate at which the secondary ENI received data over the internal network. Unit: Kbit/s.
	//
	// example:
	//
	// 0
	IntranetRx *string `json:"IntranetRx,omitempty" xml:"IntranetRx,omitempty"`
	// The average rate at which the secondary ENI sent data over the internal network. Unit: Kbit/s.
	//
	// example:
	//
	// 0
	IntranetTx *string `json:"IntranetTx,omitempty" xml:"IntranetTx,omitempty"`
	// The number of packets received by the secondary ENI over the internal network.
	//
	// example:
	//
	// 0
	PacketRx *string `json:"PacketRx,omitempty" xml:"PacketRx,omitempty"`
	// The number of packets sent by the secondary ENI over the internal network.
	//
	// example:
	//
	// 0
	PacketTx *string `json:"PacketTx,omitempty" xml:"PacketTx,omitempty"`
	// The timestamp of the monitoring data. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-05-21T03:22:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) GetDropPacketRx() *string {
	return s.DropPacketRx
}

func (s *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) GetDropPacketTx() *string {
	return s.DropPacketTx
}

func (s *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) GetEniId() *string {
	return s.EniId
}

func (s *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) GetIntranetRx() *string {
	return s.IntranetRx
}

func (s *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) GetIntranetTx() *string {
	return s.IntranetTx
}

func (s *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) GetPacketRx() *string {
	return s.PacketRx
}

func (s *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) GetPacketTx() *string {
	return s.PacketTx
}

func (s *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) SetDropPacketRx(v string) *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData {
	s.DropPacketRx = &v
	return s
}

func (s *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) SetDropPacketTx(v string) *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData {
	s.DropPacketTx = &v
	return s
}

func (s *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) SetEniId(v string) *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData {
	s.EniId = &v
	return s
}

func (s *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) SetIntranetRx(v string) *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData {
	s.IntranetRx = &v
	return s
}

func (s *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) SetIntranetTx(v string) *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData {
	s.IntranetTx = &v
	return s
}

func (s *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) SetPacketRx(v string) *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData {
	s.PacketRx = &v
	return s
}

func (s *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) SetPacketTx(v string) *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData {
	s.PacketTx = &v
	return s
}

func (s *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) SetTimeStamp(v string) *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeEniMonitorDataResponseBodyMonitorDataEniMonitorData) Validate() error {
	return dara.Validate(s)
}
