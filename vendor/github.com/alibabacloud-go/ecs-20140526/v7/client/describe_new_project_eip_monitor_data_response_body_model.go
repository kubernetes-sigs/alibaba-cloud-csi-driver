// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNewProjectEipMonitorDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEipMonitorDatas(v *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatas) *DescribeNewProjectEipMonitorDataResponseBody
	GetEipMonitorDatas() *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatas
	SetRequestId(v string) *DescribeNewProjectEipMonitorDataResponseBody
	GetRequestId() *string
}

type DescribeNewProjectEipMonitorDataResponseBody struct {
	EipMonitorDatas *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatas `json:"EipMonitorDatas,omitempty" xml:"EipMonitorDatas,omitempty" type:"Struct"`
	RequestId       *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNewProjectEipMonitorDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNewProjectEipMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNewProjectEipMonitorDataResponseBody) GetEipMonitorDatas() *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatas {
	return s.EipMonitorDatas
}

func (s *DescribeNewProjectEipMonitorDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNewProjectEipMonitorDataResponseBody) SetEipMonitorDatas(v *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatas) *DescribeNewProjectEipMonitorDataResponseBody {
	s.EipMonitorDatas = v
	return s
}

func (s *DescribeNewProjectEipMonitorDataResponseBody) SetRequestId(v string) *DescribeNewProjectEipMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNewProjectEipMonitorDataResponseBody) Validate() error {
	if s.EipMonitorDatas != nil {
		if err := s.EipMonitorDatas.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatas struct {
	EipMonitorData []*DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData `json:"EipMonitorData,omitempty" xml:"EipMonitorData,omitempty" type:"Repeated"`
}

func (s DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatas) GoString() string {
	return s.String()
}

func (s *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatas) GetEipMonitorData() []*DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	return s.EipMonitorData
}

func (s *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatas) SetEipMonitorData(v []*DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatas {
	s.EipMonitorData = v
	return s
}

func (s *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatas) Validate() error {
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

type DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData struct {
	EipBandwidth *int32  `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	EipFlow      *int32  `json:"EipFlow,omitempty" xml:"EipFlow,omitempty"`
	EipPackets   *int32  `json:"EipPackets,omitempty" xml:"EipPackets,omitempty"`
	EipRX        *int32  `json:"EipRX,omitempty" xml:"EipRX,omitempty"`
	EipTX        *int32  `json:"EipTX,omitempty" xml:"EipTX,omitempty"`
	TimeStamp    *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GetEipBandwidth() *int32 {
	return s.EipBandwidth
}

func (s *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GetEipFlow() *int32 {
	return s.EipFlow
}

func (s *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GetEipPackets() *int32 {
	return s.EipPackets
}

func (s *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GetEipRX() *int32 {
	return s.EipRX
}

func (s *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GetEipTX() *int32 {
	return s.EipTX
}

func (s *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) SetEipBandwidth(v int32) *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	s.EipBandwidth = &v
	return s
}

func (s *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) SetEipFlow(v int32) *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	s.EipFlow = &v
	return s
}

func (s *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) SetEipPackets(v int32) *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	s.EipPackets = &v
	return s
}

func (s *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) SetEipRX(v int32) *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	s.EipRX = &v
	return s
}

func (s *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) SetEipTX(v int32) *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	s.EipTX = &v
	return s
}

func (s *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) SetTimeStamp(v string) *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeNewProjectEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) Validate() error {
	return dara.Validate(s)
}
