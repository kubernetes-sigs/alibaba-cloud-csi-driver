// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoProvisioningGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoProvisioningGroupId(v string) *CreateAutoProvisioningGroupResponseBody
	GetAutoProvisioningGroupId() *string
	SetLaunchResults(v *CreateAutoProvisioningGroupResponseBodyLaunchResults) *CreateAutoProvisioningGroupResponseBody
	GetLaunchResults() *CreateAutoProvisioningGroupResponseBodyLaunchResults
	SetRequestId(v string) *CreateAutoProvisioningGroupResponseBody
	GetRequestId() *string
}

type CreateAutoProvisioningGroupResponseBody struct {
	// The ID of the auto provisioning group.
	//
	// example:
	//
	// apg-sn54avj8htgvtyh8****
	AutoProvisioningGroupId *string `json:"AutoProvisioningGroupId,omitempty" xml:"AutoProvisioningGroupId,omitempty"`
	// The instances created by the auto provisioning group. The values of the parameters in this array are returned only when AutoProvisioningGroupType is set to `instant`.
	LaunchResults *CreateAutoProvisioningGroupResponseBodyLaunchResults `json:"LaunchResults,omitempty" xml:"LaunchResults,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 745CEC9F-0DD7-4451-9FE7-8B752F39****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAutoProvisioningGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupResponseBody) GetAutoProvisioningGroupId() *string {
	return s.AutoProvisioningGroupId
}

func (s *CreateAutoProvisioningGroupResponseBody) GetLaunchResults() *CreateAutoProvisioningGroupResponseBodyLaunchResults {
	return s.LaunchResults
}

func (s *CreateAutoProvisioningGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAutoProvisioningGroupResponseBody) SetAutoProvisioningGroupId(v string) *CreateAutoProvisioningGroupResponseBody {
	s.AutoProvisioningGroupId = &v
	return s
}

func (s *CreateAutoProvisioningGroupResponseBody) SetLaunchResults(v *CreateAutoProvisioningGroupResponseBodyLaunchResults) *CreateAutoProvisioningGroupResponseBody {
	s.LaunchResults = v
	return s
}

func (s *CreateAutoProvisioningGroupResponseBody) SetRequestId(v string) *CreateAutoProvisioningGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAutoProvisioningGroupResponseBody) Validate() error {
	if s.LaunchResults != nil {
		if err := s.LaunchResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAutoProvisioningGroupResponseBodyLaunchResults struct {
	LaunchResult []*CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult `json:"LaunchResult,omitempty" xml:"LaunchResult,omitempty" type:"Repeated"`
}

func (s CreateAutoProvisioningGroupResponseBodyLaunchResults) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupResponseBodyLaunchResults) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResults) GetLaunchResult() []*CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult {
	return s.LaunchResult
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResults) SetLaunchResult(v []*CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult) *CreateAutoProvisioningGroupResponseBodyLaunchResults {
	s.LaunchResult = v
	return s
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResults) Validate() error {
	if s.LaunchResult != nil {
		for _, item := range s.LaunchResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult struct {
	// The number of created instances.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The error code returned when the instance cannot be created.
	//
	// example:
	//
	// InvalidParameter
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the instance cannot be created.
	//
	// example:
	//
	// Specific parameter is not valid.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The IDs of created instances.
	InstanceIds *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResultInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	// The instance type of the instance.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The bidding policy for the pay-as-you-go instance. Valid values:
	//
	// 	- NoSpot: The instance is a regular pay-as-you-go instance.
	//
	// 	- SpotWithPriceLimit: The instance is a spot instance for which you specify the maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instance is a spot instance for which the market price at the time of purchase is used as the bid price.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The zone ID of the instance.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult) GetInstanceIds() *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResultInstanceIds {
	return s.InstanceIds
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult) SetAmount(v int32) *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult {
	s.Amount = &v
	return s
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult) SetErrorCode(v string) *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult {
	s.ErrorCode = &v
	return s
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult) SetErrorMsg(v string) *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult {
	s.ErrorMsg = &v
	return s
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult) SetInstanceIds(v *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResultInstanceIds) *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult {
	s.InstanceIds = v
	return s
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult) SetInstanceType(v string) *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult {
	s.InstanceType = &v
	return s
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult) SetSpotStrategy(v string) *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult {
	s.SpotStrategy = &v
	return s
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult) SetZoneId(v string) *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult {
	s.ZoneId = &v
	return s
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResult) Validate() error {
	if s.InstanceIds != nil {
		if err := s.InstanceIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResultInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResultInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResultInstanceIds) GoString() string {
	return s.String()
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResultInstanceIds) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResultInstanceIds) SetInstanceId(v []*string) *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResultInstanceIds {
	s.InstanceId = v
	return s
}

func (s *CreateAutoProvisioningGroupResponseBodyLaunchResultsLaunchResultInstanceIds) Validate() error {
	return dara.Validate(s)
}
