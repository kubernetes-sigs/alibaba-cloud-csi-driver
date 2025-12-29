// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlanPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceInfo(v *DescribeSavingsPlanPriceResponseBodyPriceInfo) *DescribeSavingsPlanPriceResponseBody
	GetPriceInfo() *DescribeSavingsPlanPriceResponseBodyPriceInfo
	SetRequestId(v string) *DescribeSavingsPlanPriceResponseBody
	GetRequestId() *string
}

type DescribeSavingsPlanPriceResponseBody struct {
	PriceInfo *DescribeSavingsPlanPriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSavingsPlanPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlanPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlanPriceResponseBody) GetPriceInfo() *DescribeSavingsPlanPriceResponseBodyPriceInfo {
	return s.PriceInfo
}

func (s *DescribeSavingsPlanPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSavingsPlanPriceResponseBody) SetPriceInfo(v *DescribeSavingsPlanPriceResponseBodyPriceInfo) *DescribeSavingsPlanPriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribeSavingsPlanPriceResponseBody) SetRequestId(v string) *DescribeSavingsPlanPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSavingsPlanPriceResponseBody) Validate() error {
	if s.PriceInfo != nil {
		if err := s.PriceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSavingsPlanPriceResponseBodyPriceInfo struct {
	Price *DescribeSavingsPlanPriceResponseBodyPriceInfoPrice   `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	Rules []*DescribeSavingsPlanPriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeSavingsPlanPriceResponseBodyPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlanPriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlanPriceResponseBodyPriceInfo) GetPrice() *DescribeSavingsPlanPriceResponseBodyPriceInfoPrice {
	return s.Price
}

func (s *DescribeSavingsPlanPriceResponseBodyPriceInfo) GetRules() []*DescribeSavingsPlanPriceResponseBodyPriceInfoRules {
	return s.Rules
}

func (s *DescribeSavingsPlanPriceResponseBodyPriceInfo) SetPrice(v *DescribeSavingsPlanPriceResponseBodyPriceInfoPrice) *DescribeSavingsPlanPriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

func (s *DescribeSavingsPlanPriceResponseBodyPriceInfo) SetRules(v []*DescribeSavingsPlanPriceResponseBodyPriceInfoRules) *DescribeSavingsPlanPriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *DescribeSavingsPlanPriceResponseBodyPriceInfo) Validate() error {
	if s.Price != nil {
		if err := s.Price.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSavingsPlanPriceResponseBodyPriceInfoPrice struct {
	Currency      *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	TradePrice    *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeSavingsPlanPriceResponseBodyPriceInfoPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlanPriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlanPriceResponseBodyPriceInfoPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeSavingsPlanPriceResponseBodyPriceInfoPrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeSavingsPlanPriceResponseBodyPriceInfoPrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeSavingsPlanPriceResponseBodyPriceInfoPrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeSavingsPlanPriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribeSavingsPlanPriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribeSavingsPlanPriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribeSavingsPlanPriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeSavingsPlanPriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribeSavingsPlanPriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeSavingsPlanPriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribeSavingsPlanPriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeSavingsPlanPriceResponseBodyPriceInfoPrice) Validate() error {
	return dara.Validate(s)
}

type DescribeSavingsPlanPriceResponseBodyPriceInfoRules struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleId      *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeSavingsPlanPriceResponseBodyPriceInfoRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlanPriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlanPriceResponseBodyPriceInfoRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeSavingsPlanPriceResponseBodyPriceInfoRules) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeSavingsPlanPriceResponseBodyPriceInfoRules) SetDescription(v string) *DescribeSavingsPlanPriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *DescribeSavingsPlanPriceResponseBodyPriceInfoRules) SetRuleId(v string) *DescribeSavingsPlanPriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

func (s *DescribeSavingsPlanPriceResponseBodyPriceInfoRules) Validate() error {
	return dara.Validate(s)
}
