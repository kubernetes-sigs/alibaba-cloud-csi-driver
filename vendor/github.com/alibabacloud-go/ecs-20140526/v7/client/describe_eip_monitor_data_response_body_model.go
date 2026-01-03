// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEipMonitorDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEipMonitorDatas(v *DescribeEipMonitorDataResponseBodyEipMonitorDatas) *DescribeEipMonitorDataResponseBody
	GetEipMonitorDatas() *DescribeEipMonitorDataResponseBodyEipMonitorDatas
	SetRequestId(v string) *DescribeEipMonitorDataResponseBody
	GetRequestId() *string
}

type DescribeEipMonitorDataResponseBody struct {
	EipMonitorDatas *DescribeEipMonitorDataResponseBodyEipMonitorDatas `json:"EipMonitorDatas,omitempty" xml:"EipMonitorDatas,omitempty" type:"Struct"`
	RequestId       *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEipMonitorDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEipMonitorDataResponseBody) GetEipMonitorDatas() *DescribeEipMonitorDataResponseBodyEipMonitorDatas {
	return s.EipMonitorDatas
}

func (s *DescribeEipMonitorDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEipMonitorDataResponseBody) SetEipMonitorDatas(v *DescribeEipMonitorDataResponseBodyEipMonitorDatas) *DescribeEipMonitorDataResponseBody {
	s.EipMonitorDatas = v
	return s
}

func (s *DescribeEipMonitorDataResponseBody) SetRequestId(v string) *DescribeEipMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEipMonitorDataResponseBody) Validate() error {
	if s.EipMonitorDatas != nil {
		if err := s.EipMonitorDatas.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEipMonitorDataResponseBodyEipMonitorDatas struct {
	EipMonitorData []*DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData `json:"EipMonitorData,omitempty" xml:"EipMonitorData,omitempty" type:"Repeated"`
}

func (s DescribeEipMonitorDataResponseBodyEipMonitorDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipMonitorDataResponseBodyEipMonitorDatas) GoString() string {
	return s.String()
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatas) GetEipMonitorData() []*DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	return s.EipMonitorData
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatas) SetEipMonitorData(v []*DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) *DescribeEipMonitorDataResponseBodyEipMonitorDatas {
	s.EipMonitorData = v
	return s
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatas) Validate() error {
	if s.EipMonitorData != nil {
		for _, item := range s.EipMonitorData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData struct {
	EipBandwidth *int32  `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	EipFlow      *int32  `json:"EipFlow,omitempty" xml:"EipFlow,omitempty"`
	EipPackets   *int32  `json:"EipPackets,omitempty" xml:"EipPackets,omitempty"`
	EipRX        *int32  `json:"EipRX,omitempty" xml:"EipRX,omitempty"`
	EipTX        *int32  `json:"EipTX,omitempty" xml:"EipTX,omitempty"`
	TimeStamp    *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GetEipBandwidth() *int32 {
	return s.EipBandwidth
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GetEipFlow() *int32 {
	return s.EipFlow
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GetEipPackets() *int32 {
	return s.EipPackets
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GetEipRX() *int32 {
	return s.EipRX
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GetEipTX() *int32 {
	return s.EipTX
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) SetEipBandwidth(v int32) *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	s.EipBandwidth = &v
	return s
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) SetEipFlow(v int32) *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	s.EipFlow = &v
	return s
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) SetEipPackets(v int32) *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	s.EipPackets = &v
	return s
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) SetEipRX(v int32) *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	s.EipRX = &v
	return s
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) SetEipTX(v int32) *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	s.EipTX = &v
	return s
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) SetTimeStamp(v string) *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) Validate() error {
	return dara.Validate(s)
}
