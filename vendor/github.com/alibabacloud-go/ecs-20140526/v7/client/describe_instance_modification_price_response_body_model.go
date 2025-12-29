// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceModificationPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceInfo(v *DescribeInstanceModificationPriceResponseBodyPriceInfo) *DescribeInstanceModificationPriceResponseBody
	GetPriceInfo() *DescribeInstanceModificationPriceResponseBodyPriceInfo
	SetRequestId(v string) *DescribeInstanceModificationPriceResponseBody
	GetRequestId() *string
}

type DescribeInstanceModificationPriceResponseBody struct {
	// Details about the prices and promotion rules.
	PriceInfo *DescribeInstanceModificationPriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A3DC3196-379B-4F32-A2C5-B937134FAD8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceModificationPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceModificationPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceModificationPriceResponseBody) GetPriceInfo() *DescribeInstanceModificationPriceResponseBodyPriceInfo {
	return s.PriceInfo
}

func (s *DescribeInstanceModificationPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceModificationPriceResponseBody) SetPriceInfo(v *DescribeInstanceModificationPriceResponseBodyPriceInfo) *DescribeInstanceModificationPriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBody) SetRequestId(v string) *DescribeInstanceModificationPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBody) Validate() error {
	if s.PriceInfo != nil {
		if err := s.PriceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceModificationPriceResponseBodyPriceInfo struct {
	// The price.
	Price *DescribeInstanceModificationPriceResponseBodyPriceInfoPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// The information about the promotion rules.
	Rules *DescribeInstanceModificationPriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s DescribeInstanceModificationPriceResponseBodyPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceModificationPriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfo) GetPrice() *DescribeInstanceModificationPriceResponseBodyPriceInfoPrice {
	return s.Price
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfo) GetRules() *DescribeInstanceModificationPriceResponseBodyPriceInfoRules {
	return s.Rules
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfo) SetPrice(v *DescribeInstanceModificationPriceResponseBodyPriceInfoPrice) *DescribeInstanceModificationPriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfo) SetRules(v *DescribeInstanceModificationPriceResponseBodyPriceInfoRules) *DescribeInstanceModificationPriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfo) Validate() error {
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

type DescribeInstanceModificationPriceResponseBodyPriceInfoPrice struct {
	// The currency unit. Valid values:
	//
	// 	- Alibaba Cloud China site (aliyun.com): CNY
	//
	// 	- Alibaba Cloud International site (alibabacloud.com): USD
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The information about the price.
	//
	// >  This parameter is returned only when ResourceType is set to instance.
	DetailInfos *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfos `json:"DetailInfos,omitempty" xml:"DetailInfos,omitempty" type:"Struct"`
	// The discount.
	//
	// example:
	//
	// 61.320
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	//
	// example:
	//
	// 175.200
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The transaction price, which is equal to the original price minus the discount.
	//
	// example:
	//
	// 113.880
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeInstanceModificationPriceResponseBodyPriceInfoPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceModificationPriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPrice) GetDetailInfos() *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfos {
	return s.DetailInfos
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribeInstanceModificationPriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPrice) SetDetailInfos(v *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfos) *DescribeInstanceModificationPriceResponseBodyPriceInfoPrice {
	s.DetailInfos = v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribeInstanceModificationPriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribeInstanceModificationPriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribeInstanceModificationPriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPrice) Validate() error {
	if s.DetailInfos != nil {
		if err := s.DetailInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfos struct {
	DetailInfo []*DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Repeated"`
}

func (s DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfos) GoString() string {
	return s.String()
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfos) GetDetailInfo() []*DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	return s.DetailInfo
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfos) SetDetailInfo(v []*DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfos {
	s.DetailInfo = v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfos) Validate() error {
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

type DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo struct {
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
	// instanceType
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The pricing rules.
	SubRules *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules `json:"SubRules,omitempty" xml:"SubRules,omitempty" type:"Struct"`
	// The transaction price.
	//
	// example:
	//
	// 3712.8
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GetResource() *string {
	return s.Resource
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GetSubRules() *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules {
	return s.SubRules
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetDiscountPrice(v float32) *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetOriginalPrice(v float32) *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetResource(v string) *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.Resource = &v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetSubRules(v *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules) *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.SubRules = v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetTradePrice(v float32) *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.TradePrice = &v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) Validate() error {
	if s.SubRules != nil {
		if err := s.SubRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules struct {
	Rule []*DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules) GoString() string {
	return s.String()
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules) GetRule() []*DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule {
	return s.Rule
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules) SetRule(v []*DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules {
	s.Rule = v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRules) Validate() error {
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

type DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule struct {
	// The description of the pricing rule.
	//
	// example:
	//
	// If you subscribe to an instance for one year, you can receive a 15% discount off the list price.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the pricing rule.
	//
	// example:
	//
	// 315716429631488
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) SetDescription(v string) *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule {
	s.Description = &v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) SetRuleId(v int64) *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoSubRulesRule) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceModificationPriceResponseBodyPriceInfoRules struct {
	Rule []*DescribeInstanceModificationPriceResponseBodyPriceInfoRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeInstanceModificationPriceResponseBodyPriceInfoRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceModificationPriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoRules) GetRule() []*DescribeInstanceModificationPriceResponseBodyPriceInfoRulesRule {
	return s.Rule
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoRules) SetRule(v []*DescribeInstanceModificationPriceResponseBodyPriceInfoRulesRule) *DescribeInstanceModificationPriceResponseBodyPriceInfoRules {
	s.Rule = v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoRules) Validate() error {
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

type DescribeInstanceModificationPriceResponseBodyPriceInfoRulesRule struct {
	// The description of the promotion rule.
	//
	// example:
	//
	// Upgrade offers
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the promotion rule.
	//
	// example:
	//
	// 1234567890
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeInstanceModificationPriceResponseBodyPriceInfoRulesRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceModificationPriceResponseBodyPriceInfoRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoRulesRule) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoRulesRule) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoRulesRule) SetDescription(v string) *DescribeInstanceModificationPriceResponseBodyPriceInfoRulesRule {
	s.Description = &v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoRulesRule) SetRuleId(v int64) *DescribeInstanceModificationPriceResponseBodyPriceInfoRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribeInstanceModificationPriceResponseBodyPriceInfoRulesRule) Validate() error {
	return dara.Validate(s)
}
