// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenewalPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceInfo(v *DescribeRenewalPriceResponseBodyPriceInfo) *DescribeRenewalPriceResponseBody
	GetPriceInfo() *DescribeRenewalPriceResponseBodyPriceInfo
	SetRequestId(v string) *DescribeRenewalPriceResponseBody
	GetRequestId() *string
}

type DescribeRenewalPriceResponseBody struct {
	// Details about the prices and promotion rules.
	PriceInfo *DescribeRenewalPriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRenewalPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBody) GetPriceInfo() *DescribeRenewalPriceResponseBodyPriceInfo {
	return s.PriceInfo
}

func (s *DescribeRenewalPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRenewalPriceResponseBody) SetPriceInfo(v *DescribeRenewalPriceResponseBodyPriceInfo) *DescribeRenewalPriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribeRenewalPriceResponseBody) SetRequestId(v string) *DescribeRenewalPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRenewalPriceResponseBody) Validate() error {
	if s.PriceInfo != nil {
		if err := s.PriceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRenewalPriceResponseBodyPriceInfo struct {
	// The price.
	Price *DescribeRenewalPriceResponseBodyPriceInfoPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// The information about the promotion rules.
	Rules *DescribeRenewalPriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) GetPrice() *DescribeRenewalPriceResponseBodyPriceInfoPrice {
	return s.Price
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) GetRules() *DescribeRenewalPriceResponseBodyPriceInfoRules {
	return s.Rules
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) SetPrice(v *DescribeRenewalPriceResponseBodyPriceInfoPrice) *DescribeRenewalPriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) SetRules(v *DescribeRenewalPriceResponseBodyPriceInfoRules) *DescribeRenewalPriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) Validate() error {
	if s.Price != nil {
		if err := s.Price.Validate(); err != nil {
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

type DescribeRenewalPriceResponseBodyPriceInfoPrice struct {
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
	// Details about the resource prices.
	DetailInfos *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfos `json:"DetailInfos,omitempty" xml:"DetailInfos,omitempty" type:"Struct"`
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
	// The transaction price, which is equal to the original price minus the discount.
	//
	// example:
	//
	// 3712.8
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfoPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) GetDetailInfos() *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfos {
	return s.DetailInfos
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribeRenewalPriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) SetDetailInfos(v *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfos) *DescribeRenewalPriceResponseBodyPriceInfoPrice {
	s.DetailInfos = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribeRenewalPriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribeRenewalPriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribeRenewalPriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) Validate() error {
	if s.DetailInfos != nil {
		if err := s.DetailInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfos struct {
	DetailInfo []*DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Repeated"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfos) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfos) GetDetailInfo() []*DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	return s.DetailInfo
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfos) SetDetailInfo(v []*DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfos {
	s.DetailInfo = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfos) Validate() error {
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

type DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo struct {
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
	// The name of the resource that corresponds to the price.
	//
	// example:
	//
	// instance
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The pricing rules.
	SubRules *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules `json:"SubRules,omitempty" xml:"SubRules,omitempty" type:"Struct"`
	// The transaction price.
	//
	// example:
	//
	// 3712.8
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GetResource() *string {
	return s.Resource
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GetSubRules() *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules {
	return s.SubRules
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetDiscountPrice(v float32) *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetOriginalPrice(v float32) *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetResource(v string) *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.Resource = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetSubRules(v *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules) *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.SubRules = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetTradePrice(v float32) *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.TradePrice = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) Validate() error {
	if s.SubRules != nil {
		if err := s.SubRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules struct {
	Rule []*DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules) GetRule() []*DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule {
	return s.Rule
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules) SetRule(v []*DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules {
	s.Rule = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules) Validate() error {
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

type DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule struct {
	// The description of the pricing rule.
	//
	// example:
	//
	// Receive a 15% discount on a 1-year subscription.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the pricing rule.
	//
	// example:
	//
	// 1234567890
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) GetDescription() *string {
	return s.Description
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) SetDescription(v string) *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule {
	s.Description = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) SetRuleId(v int64) *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) Validate() error {
	return dara.Validate(s)
}

type DescribeRenewalPriceResponseBodyPriceInfoRules struct {
	Rule []*DescribeRenewalPriceResponseBodyPriceInfoRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfoRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoRules) GetRule() []*DescribeRenewalPriceResponseBodyPriceInfoRulesRule {
	return s.Rule
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoRules) SetRule(v []*DescribeRenewalPriceResponseBodyPriceInfoRulesRule) *DescribeRenewalPriceResponseBodyPriceInfoRules {
	s.Rule = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoRules) Validate() error {
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

type DescribeRenewalPriceResponseBodyPriceInfoRulesRule struct {
	// The description of the promotion rule.
	//
	// example:
	//
	// Receive a 15% discount on a 1-year subscription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the promotion rule.
	//
	// example:
	//
	// 1234567890
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfoRulesRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfoRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoRulesRule) GetDescription() *string {
	return s.Description
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoRulesRule) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoRulesRule) SetDescription(v string) *DescribeRenewalPriceResponseBodyPriceInfoRulesRule {
	s.Description = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoRulesRule) SetRuleId(v int64) *DescribeRenewalPriceResponseBodyPriceInfoRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoRulesRule) Validate() error {
	return dara.Validate(s)
}
