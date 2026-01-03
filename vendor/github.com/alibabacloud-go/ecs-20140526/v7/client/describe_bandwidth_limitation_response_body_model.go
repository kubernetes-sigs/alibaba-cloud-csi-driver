// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandwidthLimitationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidths(v *DescribeBandwidthLimitationResponseBodyBandwidths) *DescribeBandwidthLimitationResponseBody
	GetBandwidths() *DescribeBandwidthLimitationResponseBodyBandwidths
	SetRequestId(v string) *DescribeBandwidthLimitationResponseBody
	GetRequestId() *string
}

type DescribeBandwidthLimitationResponseBody struct {
	// Details about the maximum public bandwidth.
	Bandwidths *DescribeBandwidthLimitationResponseBodyBandwidths `json:"Bandwidths,omitempty" xml:"Bandwidths,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBandwidthLimitationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthLimitationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthLimitationResponseBody) GetBandwidths() *DescribeBandwidthLimitationResponseBodyBandwidths {
	return s.Bandwidths
}

func (s *DescribeBandwidthLimitationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBandwidthLimitationResponseBody) SetBandwidths(v *DescribeBandwidthLimitationResponseBodyBandwidths) *DescribeBandwidthLimitationResponseBody {
	s.Bandwidths = v
	return s
}

func (s *DescribeBandwidthLimitationResponseBody) SetRequestId(v string) *DescribeBandwidthLimitationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBandwidthLimitationResponseBody) Validate() error {
	if s.Bandwidths != nil {
		if err := s.Bandwidths.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBandwidthLimitationResponseBodyBandwidths struct {
	Bandwidth []*DescribeBandwidthLimitationResponseBodyBandwidthsBandwidth `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty" type:"Repeated"`
}

func (s DescribeBandwidthLimitationResponseBodyBandwidths) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthLimitationResponseBodyBandwidths) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthLimitationResponseBodyBandwidths) GetBandwidth() []*DescribeBandwidthLimitationResponseBodyBandwidthsBandwidth {
	return s.Bandwidth
}

func (s *DescribeBandwidthLimitationResponseBodyBandwidths) SetBandwidth(v []*DescribeBandwidthLimitationResponseBodyBandwidthsBandwidth) *DescribeBandwidthLimitationResponseBodyBandwidths {
	s.Bandwidth = v
	return s
}

func (s *DescribeBandwidthLimitationResponseBodyBandwidths) Validate() error {
	if s.Bandwidth != nil {
		for _, item := range s.Bandwidth {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBandwidthLimitationResponseBodyBandwidthsBandwidth struct {
	// The billing method for network usage. Valid values:
	//
	// 	- PayByBandwidth
	//
	// 	- PayByTraffic
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum public bandwidth.
	//
	// example:
	//
	// 100
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// The minimum public bandwidth.
	//
	// example:
	//
	// 0
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
	// The unit of the public bandwidth.
	//
	// example:
	//
	// Mbps
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeBandwidthLimitationResponseBodyBandwidthsBandwidth) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthLimitationResponseBodyBandwidthsBandwidth) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthLimitationResponseBodyBandwidthsBandwidth) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeBandwidthLimitationResponseBodyBandwidthsBandwidth) GetMax() *int32 {
	return s.Max
}

func (s *DescribeBandwidthLimitationResponseBodyBandwidthsBandwidth) GetMin() *int32 {
	return s.Min
}

func (s *DescribeBandwidthLimitationResponseBodyBandwidthsBandwidth) GetUnit() *string {
	return s.Unit
}

func (s *DescribeBandwidthLimitationResponseBodyBandwidthsBandwidth) SetInternetChargeType(v string) *DescribeBandwidthLimitationResponseBodyBandwidthsBandwidth {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeBandwidthLimitationResponseBodyBandwidthsBandwidth) SetMax(v int32) *DescribeBandwidthLimitationResponseBodyBandwidthsBandwidth {
	s.Max = &v
	return s
}

func (s *DescribeBandwidthLimitationResponseBodyBandwidthsBandwidth) SetMin(v int32) *DescribeBandwidthLimitationResponseBodyBandwidthsBandwidth {
	s.Min = &v
	return s
}

func (s *DescribeBandwidthLimitationResponseBodyBandwidthsBandwidth) SetUnit(v string) *DescribeBandwidthLimitationResponseBodyBandwidthsBandwidth {
	s.Unit = &v
	return s
}

func (s *DescribeBandwidthLimitationResponseBodyBandwidthsBandwidth) Validate() error {
	return dara.Validate(s)
}
