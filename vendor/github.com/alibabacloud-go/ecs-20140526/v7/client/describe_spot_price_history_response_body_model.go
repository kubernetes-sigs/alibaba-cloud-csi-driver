// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSpotPriceHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrency(v string) *DescribeSpotPriceHistoryResponseBody
	GetCurrency() *string
	SetNextOffset(v int32) *DescribeSpotPriceHistoryResponseBody
	GetNextOffset() *int32
	SetRequestId(v string) *DescribeSpotPriceHistoryResponseBody
	GetRequestId() *string
	SetSpotPrices(v *DescribeSpotPriceHistoryResponseBodySpotPrices) *DescribeSpotPriceHistoryResponseBody
	GetSpotPrices() *DescribeSpotPriceHistoryResponseBodySpotPrices
}

type DescribeSpotPriceHistoryResponseBody struct {
	// The instance type of the spot instance.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The network type of the spot instance.
	//
	// example:
	//
	// 1000
	NextOffset *int32 `json:"NextOffset,omitempty" xml:"NextOffset,omitempty"`
	// The instance type of the spot instance.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SpotPrices *DescribeSpotPriceHistoryResponseBodySpotPrices `json:"SpotPrices,omitempty" xml:"SpotPrices,omitempty" type:"Struct"`
}

func (s DescribeSpotPriceHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpotPriceHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSpotPriceHistoryResponseBody) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeSpotPriceHistoryResponseBody) GetNextOffset() *int32 {
	return s.NextOffset
}

func (s *DescribeSpotPriceHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSpotPriceHistoryResponseBody) GetSpotPrices() *DescribeSpotPriceHistoryResponseBodySpotPrices {
	return s.SpotPrices
}

func (s *DescribeSpotPriceHistoryResponseBody) SetCurrency(v string) *DescribeSpotPriceHistoryResponseBody {
	s.Currency = &v
	return s
}

func (s *DescribeSpotPriceHistoryResponseBody) SetNextOffset(v int32) *DescribeSpotPriceHistoryResponseBody {
	s.NextOffset = &v
	return s
}

func (s *DescribeSpotPriceHistoryResponseBody) SetRequestId(v string) *DescribeSpotPriceHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSpotPriceHistoryResponseBody) SetSpotPrices(v *DescribeSpotPriceHistoryResponseBodySpotPrices) *DescribeSpotPriceHistoryResponseBody {
	s.SpotPrices = v
	return s
}

func (s *DescribeSpotPriceHistoryResponseBody) Validate() error {
	if s.SpotPrices != nil {
		if err := s.SpotPrices.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSpotPriceHistoryResponseBodySpotPrices struct {
	SpotPriceType []*DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType `json:"SpotPriceType,omitempty" xml:"SpotPriceType,omitempty" type:"Repeated"`
}

func (s DescribeSpotPriceHistoryResponseBodySpotPrices) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpotPriceHistoryResponseBodySpotPrices) GoString() string {
	return s.String()
}

func (s *DescribeSpotPriceHistoryResponseBodySpotPrices) GetSpotPriceType() []*DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType {
	return s.SpotPriceType
}

func (s *DescribeSpotPriceHistoryResponseBodySpotPrices) SetSpotPriceType(v []*DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType) *DescribeSpotPriceHistoryResponseBodySpotPrices {
	s.SpotPriceType = v
	return s
}

func (s *DescribeSpotPriceHistoryResponseBodySpotPrices) Validate() error {
	if s.SpotPriceType != nil {
		for _, item := range s.SpotPriceType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType struct {
	InstanceType *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	IoOptimized  *string  `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	NetworkType  *string  `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OriginPrice  *float32 `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
	SpotPrice    *float32 `json:"SpotPrice,omitempty" xml:"SpotPrice,omitempty"`
	Timestamp    *string  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	ZoneId       *string  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType) String() string {
	return dara.Prettify(s)
}

func (s DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType) GoString() string {
	return s.String()
}

func (s *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType) GetOriginPrice() *float32 {
	return s.OriginPrice
}

func (s *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType) GetSpotPrice() *float32 {
	return s.SpotPrice
}

func (s *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType) GetTimestamp() *string {
	return s.Timestamp
}

func (s *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType) SetInstanceType(v string) *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType {
	s.InstanceType = &v
	return s
}

func (s *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType) SetIoOptimized(v string) *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType {
	s.IoOptimized = &v
	return s
}

func (s *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType) SetNetworkType(v string) *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType {
	s.NetworkType = &v
	return s
}

func (s *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType) SetOriginPrice(v float32) *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType {
	s.OriginPrice = &v
	return s
}

func (s *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType) SetSpotPrice(v float32) *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType {
	s.SpotPrice = &v
	return s
}

func (s *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType) SetTimestamp(v string) *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType {
	s.Timestamp = &v
	return s
}

func (s *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType) SetZoneId(v string) *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType {
	s.ZoneId = &v
	return s
}

func (s *DescribeSpotPriceHistoryResponseBodySpotPricesSpotPriceType) Validate() error {
	return dara.Validate(s)
}
