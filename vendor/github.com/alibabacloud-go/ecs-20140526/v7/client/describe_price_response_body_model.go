// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceInfo(v *DescribePriceResponseBodyPriceInfo) *DescribePriceResponseBody
	GetPriceInfo() *DescribePriceResponseBodyPriceInfo
	SetRequestId(v string) *DescribePriceResponseBody
	GetRequestId() *string
}

type DescribePriceResponseBody struct {
	// The information about the prices and promotion rules.
	PriceInfo *DescribePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBody) GetPriceInfo() *DescribePriceResponseBodyPriceInfo {
	return s.PriceInfo
}

func (s *DescribePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePriceResponseBody) SetPriceInfo(v *DescribePriceResponseBodyPriceInfo) *DescribePriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribePriceResponseBody) SetRequestId(v string) *DescribePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePriceResponseBody) Validate() error {
	if s.PriceInfo != nil {
		if err := s.PriceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePriceResponseBodyPriceInfo struct {
	// The price.
	Price *DescribePriceResponseBodyPriceInfoPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// The related price.
	RelatedPrice *DescribePriceResponseBodyPriceInfoRelatedPrice `json:"RelatedPrice,omitempty" xml:"RelatedPrice,omitempty" type:"Struct"`
	// The information about the promotion rules.
	Rules *DescribePriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s DescribePriceResponseBodyPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfo) GetPrice() *DescribePriceResponseBodyPriceInfoPrice {
	return s.Price
}

func (s *DescribePriceResponseBodyPriceInfo) GetRelatedPrice() *DescribePriceResponseBodyPriceInfoRelatedPrice {
	return s.RelatedPrice
}

func (s *DescribePriceResponseBodyPriceInfo) GetRules() *DescribePriceResponseBodyPriceInfoRules {
	return s.Rules
}

func (s *DescribePriceResponseBodyPriceInfo) SetPrice(v *DescribePriceResponseBodyPriceInfoPrice) *DescribePriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) SetRelatedPrice(v *DescribePriceResponseBodyPriceInfoRelatedPrice) *DescribePriceResponseBodyPriceInfo {
	s.RelatedPrice = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) SetRules(v *DescribePriceResponseBodyPriceInfoRules) *DescribePriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) Validate() error {
	if s.Price != nil {
		if err := s.Price.Validate(); err != nil {
			return err
		}
	}
	if s.RelatedPrice != nil {
		if err := s.RelatedPrice.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		if err := s.Rules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePriceResponseBodyPriceInfoPrice struct {
	// The currency unit.
	//
	// Alibaba Cloud China site (aliyun.com): CNY.
	//
	// Alibaba Cloud International site (alibabacloud.com): USD.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The information about the price.
	//
	// >  This parameter is returned only when ResourceType is set to instance.
	DetailInfos *DescribePriceResponseBodyPriceInfoPriceDetailInfos `json:"DetailInfos,omitempty" xml:"DetailInfos,omitempty" type:"Struct"`
	// The discount.
	//
	// example:
	//
	// 655.2
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	//
	// example:
	//
	// 4368
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The hourly price of the reserved instance for which the No Upfront or Partial Upfront payment option is used.
	//
	// example:
	//
	// 1
	ReservedInstanceHourPrice *float32 `json:"ReservedInstanceHourPrice,omitempty" xml:"ReservedInstanceHourPrice,omitempty"`
	// The transaction price of the order. The transaction price is equal to the original price minus the discount.
	//
	// example:
	//
	// 3712.8
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetDetailInfos() *DescribePriceResponseBodyPriceInfoPriceDetailInfos {
	return s.DetailInfos
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetReservedInstanceHourPrice() *float32 {
	return s.ReservedInstanceHourPrice
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribePriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetDetailInfos(v *DescribePriceResponseBodyPriceInfoPriceDetailInfos) *DescribePriceResponseBodyPriceInfoPrice {
	s.DetailInfos = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetReservedInstanceHourPrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.ReservedInstanceHourPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) Validate() error {
	if s.DetailInfos != nil {
		if err := s.DetailInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePriceResponseBodyPriceInfoPriceDetailInfos struct {
	DetailInfo []*DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyPriceInfoPriceDetailInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPriceDetailInfos) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfos) GetDetailInfo() []*DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	return s.DetailInfo
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfos) SetDetailInfo(v []*DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) *DescribePriceResponseBodyPriceInfoPriceDetailInfos {
	s.DetailInfo = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfos) Validate() error {
	if s.DetailInfo != nil {
		for _, item := range s.DetailInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo struct {
	// The discount.
	//
	// example:
	//
	// 655.2
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	//
	// example:
	//
	// 4368
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The resource name. Valid values:
	//
	// 	- InstanceType
	//
	// 	- bandwidth
	//
	// 	- image
	//
	// 	- SystemDisk
	//
	// 	- DataDisk
	//
	// example:
	//
	// instance
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// Details about the pricing rules.
	SubRules *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules `json:"SubRules,omitempty" xml:"SubRules,omitempty" type:"Struct"`
	// The transaction price.
	//
	// example:
	//
	// 3712.8
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GetResource() *string {
	return s.Resource
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GetSubRules() *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules {
	return s.SubRules
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetDiscountPrice(v float32) *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.DiscountPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetOriginalPrice(v float32) *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetResource(v string) *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.Resource = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetSubRules(v *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules) *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.SubRules = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetTradePrice(v float32) *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.TradePrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) Validate() error {
	if s.SubRules != nil {
		if err := s.SubRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules struct {
	Rule []*DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules) GetRule() []*DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule {
	return s.Rule
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules) SetRule(v []*DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules {
	s.Rule = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules) Validate() error {
	if s.Rule != nil {
		for _, item := range s.Rule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule struct {
	// The description of the pricing rule.
	//
	// example:
	//
	// Receive a 15% discount on a 1-year subscription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the pricing rule.
	//
	// example:
	//
	// 587
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) GetDescription() *string {
	return s.Description
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) SetDescription(v string) *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule {
	s.Description = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) SetRuleId(v int64) *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyPriceInfoRelatedPrice struct {
	// The Alibaba Cloud Marketplace image price.
	MarketplaceImagePrice *DescribePriceResponseBodyPriceInfoRelatedPriceMarketplaceImagePrice `json:"MarketplaceImagePrice,omitempty" xml:"MarketplaceImagePrice,omitempty" type:"Struct"`
}

func (s DescribePriceResponseBodyPriceInfoRelatedPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoRelatedPrice) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoRelatedPrice) GetMarketplaceImagePrice() *DescribePriceResponseBodyPriceInfoRelatedPriceMarketplaceImagePrice {
	return s.MarketplaceImagePrice
}

func (s *DescribePriceResponseBodyPriceInfoRelatedPrice) SetMarketplaceImagePrice(v *DescribePriceResponseBodyPriceInfoRelatedPriceMarketplaceImagePrice) *DescribePriceResponseBodyPriceInfoRelatedPrice {
	s.MarketplaceImagePrice = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoRelatedPrice) Validate() error {
	if s.MarketplaceImagePrice != nil {
		if err := s.MarketplaceImagePrice.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePriceResponseBodyPriceInfoRelatedPriceMarketplaceImagePrice struct {
	// The currency unit.
	//
	// China site (aliyun.com): CNY
	//
	// International site (alibabacloud.com): USD
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The discount.
	//
	// example:
	//
	// 0
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	//
	// example:
	//
	// 100
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The transaction price, which is equal to the original price minus the discount.
	//
	// example:
	//
	// 100
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoRelatedPriceMarketplaceImagePrice) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoRelatedPriceMarketplaceImagePrice) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoRelatedPriceMarketplaceImagePrice) GetCurrency() *string {
	return s.Currency
}

func (s *DescribePriceResponseBodyPriceInfoRelatedPriceMarketplaceImagePrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribePriceResponseBodyPriceInfoRelatedPriceMarketplaceImagePrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribePriceResponseBodyPriceInfoRelatedPriceMarketplaceImagePrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribePriceResponseBodyPriceInfoRelatedPriceMarketplaceImagePrice) SetCurrency(v string) *DescribePriceResponseBodyPriceInfoRelatedPriceMarketplaceImagePrice {
	s.Currency = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoRelatedPriceMarketplaceImagePrice) SetDiscountPrice(v float32) *DescribePriceResponseBodyPriceInfoRelatedPriceMarketplaceImagePrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoRelatedPriceMarketplaceImagePrice) SetOriginalPrice(v float32) *DescribePriceResponseBodyPriceInfoRelatedPriceMarketplaceImagePrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoRelatedPriceMarketplaceImagePrice) SetTradePrice(v float32) *DescribePriceResponseBodyPriceInfoRelatedPriceMarketplaceImagePrice {
	s.TradePrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoRelatedPriceMarketplaceImagePrice) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyPriceInfoRules struct {
	Rule []*DescribePriceResponseBodyPriceInfoRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyPriceInfoRules) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoRules) GetRule() []*DescribePriceResponseBodyPriceInfoRulesRule {
	return s.Rule
}

func (s *DescribePriceResponseBodyPriceInfoRules) SetRule(v []*DescribePriceResponseBodyPriceInfoRulesRule) *DescribePriceResponseBodyPriceInfoRules {
	s.Rule = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoRules) Validate() error {
	if s.Rule != nil {
		for _, item := range s.Rule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePriceResponseBodyPriceInfoRulesRule struct {
	// The description of the promotion rule.
	//
	// example:
	//
	// Receive a 15% discount on a 1-year subscription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the pricing rule.
	//
	// example:
	//
	// 587
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoRulesRule) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoRulesRule) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoRulesRule) GetDescription() *string {
	return s.Description
}

func (s *DescribePriceResponseBodyPriceInfoRulesRule) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribePriceResponseBodyPriceInfoRulesRule) SetDescription(v string) *DescribePriceResponseBodyPriceInfoRulesRule {
	s.Description = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoRulesRule) SetRuleId(v int64) *DescribePriceResponseBodyPriceInfoRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoRulesRule) Validate() error {
	return dara.Validate(s)
}
