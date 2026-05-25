// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLaunchTemplateVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLaunchTemplateVersionSets(v *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSets) *DescribeLaunchTemplateVersionsResponseBody
	GetLaunchTemplateVersionSets() *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSets
	SetPageNumber(v int32) *DescribeLaunchTemplateVersionsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLaunchTemplateVersionsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLaunchTemplateVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeLaunchTemplateVersionsResponseBody
	GetTotalCount() *int32
}

type DescribeLaunchTemplateVersionsResponseBody struct {
	LaunchTemplateVersionSets *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSets `json:"LaunchTemplateVersionSets,omitempty" xml:"LaunchTemplateVersionSets,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3989ED0C-20A1-4351-A127-2067FF8390AX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of launch templates.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLaunchTemplateVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplateVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplateVersionsResponseBody) GetLaunchTemplateVersionSets() *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSets {
	return s.LaunchTemplateVersionSets
}

func (s *DescribeLaunchTemplateVersionsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLaunchTemplateVersionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLaunchTemplateVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLaunchTemplateVersionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeLaunchTemplateVersionsResponseBody) SetLaunchTemplateVersionSets(v *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSets) *DescribeLaunchTemplateVersionsResponseBody {
	s.LaunchTemplateVersionSets = v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBody) SetPageNumber(v int32) *DescribeLaunchTemplateVersionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBody) SetPageSize(v int32) *DescribeLaunchTemplateVersionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBody) SetRequestId(v string) *DescribeLaunchTemplateVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBody) SetTotalCount(v int32) *DescribeLaunchTemplateVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBody) Validate() error {
	if s.LaunchTemplateVersionSets != nil {
		if err := s.LaunchTemplateVersionSets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSets struct {
	LaunchTemplateVersionSet []*DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet `json:"LaunchTemplateVersionSet,omitempty" xml:"LaunchTemplateVersionSet,omitempty" type:"Repeated"`
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSets) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSets) GetLaunchTemplateVersionSet() []*DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet {
	return s.LaunchTemplateVersionSet
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSets) SetLaunchTemplateVersionSet(v []*DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSets {
	s.LaunchTemplateVersionSet = v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSets) Validate() error {
	if s.LaunchTemplateVersionSet != nil {
		for _, item := range s.LaunchTemplateVersionSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet struct {
	CreateTime         *string                                                                                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatedBy          *string                                                                                                        `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	DefaultVersion     *bool                                                                                                          `json:"DefaultVersion,omitempty" xml:"DefaultVersion,omitempty"`
	LaunchTemplateData *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData `json:"LaunchTemplateData,omitempty" xml:"LaunchTemplateData,omitempty" type:"Struct"`
	LaunchTemplateId   *string                                                                                                        `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	LaunchTemplateName *string                                                                                                        `json:"LaunchTemplateName,omitempty" xml:"LaunchTemplateName,omitempty"`
	ModifiedTime       *string                                                                                                        `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	VersionDescription *string                                                                                                        `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
	VersionNumber      *int64                                                                                                         `json:"VersionNumber,omitempty" xml:"VersionNumber,omitempty"`
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) GetDefaultVersion() *bool {
	return s.DefaultVersion
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) GetLaunchTemplateData() *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	return s.LaunchTemplateData
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) GetLaunchTemplateName() *string {
	return s.LaunchTemplateName
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) GetVersionNumber() *int64 {
	return s.VersionNumber
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) SetCreateTime(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet {
	s.CreateTime = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) SetCreatedBy(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet {
	s.CreatedBy = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) SetDefaultVersion(v bool) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet {
	s.DefaultVersion = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) SetLaunchTemplateData(v *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet {
	s.LaunchTemplateData = v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) SetLaunchTemplateId(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet {
	s.LaunchTemplateId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) SetLaunchTemplateName(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet {
	s.LaunchTemplateName = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) SetModifiedTime(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) SetVersionDescription(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet {
	s.VersionDescription = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) SetVersionNumber(v int64) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet {
	s.VersionNumber = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSet) Validate() error {
	if s.LaunchTemplateData != nil {
		if err := s.LaunchTemplateData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData struct {
	SystemDisk                  *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk        `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	AutoReleaseTime             *string                                                                                                                         `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	AutoRenew                   *bool                                                                                                                           `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod             *int32                                                                                                                          `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	CreditSpecification         *string                                                                                                                         `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	DataDisks                   *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisks         `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Struct"`
	DeletionProtection          *bool                                                                                                                           `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	DeploymentSetId             *string                                                                                                                         `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	Description                 *string                                                                                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableVmOsConfig            *bool                                                                                                                           `json:"EnableVmOsConfig,omitempty" xml:"EnableVmOsConfig,omitempty"`
	HostName                    *string                                                                                                                         `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HttpEndpoint                *string                                                                                                                         `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	HttpPutResponseHopLimit     *int32                                                                                                                          `json:"HttpPutResponseHopLimit,omitempty" xml:"HttpPutResponseHopLimit,omitempty"`
	HttpTokens                  *string                                                                                                                         `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
	ImageId                     *string                                                                                                                         `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageOptions                *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataImageOptions      `json:"ImageOptions,omitempty" xml:"ImageOptions,omitempty" type:"Struct"`
	ImageOwnerAlias             *string                                                                                                                         `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	InstanceChargeType          *string                                                                                                                         `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceName                *string                                                                                                                         `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType                *string                                                                                                                         `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetChargeType          *string                                                                                                                         `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandwidthIn      *int32                                                                                                                          `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	InternetMaxBandwidthOut     *int32                                                                                                                          `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	IoOptimized                 *string                                                                                                                         `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	Ipv6AddressCount            *int32                                                                                                                          `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	KeyPairName                 *string                                                                                                                         `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	NetworkInterfaces           *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfaces `json:"NetworkInterfaces,omitempty" xml:"NetworkInterfaces,omitempty" type:"Struct"`
	NetworkType                 *string                                                                                                                         `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PasswordInherit             *bool                                                                                                                           `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	Period                      *int32                                                                                                                          `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit                  *string                                                                                                                         `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PrivateIpAddress            *string                                                                                                                         `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	RamRoleName                 *string                                                                                                                         `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	ResourceGroupId             *string                                                                                                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityEnhancementStrategy *string                                                                                                                         `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	SecurityGroupId             *string                                                                                                                         `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupIds            *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityGroupIds  `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Struct"`
	SecurityOptions             *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityOptions   `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
	SpotDuration                *int32                                                                                                                          `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	SpotPriceLimit              *float32                                                                                                                        `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	SpotStrategy                *string                                                                                                                         `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	Tags                        *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTags              `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UserData                    *string                                                                                                                         `json:"UserData,omitempty" xml:"UserData,omitempty"`
	VSwitchId                   *string                                                                                                                         `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                       *string                                                                                                                         `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                      *string                                                                                                                         `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetSystemDisk() *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk {
	return s.SystemDisk
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetAutoReleaseTime() *string {
	return s.AutoReleaseTime
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetDataDisks() *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisks {
	return s.DataDisks
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetDescription() *string {
	return s.Description
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetEnableVmOsConfig() *bool {
	return s.EnableVmOsConfig
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetHostName() *string {
	return s.HostName
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetHttpEndpoint() *string {
	return s.HttpEndpoint
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetHttpPutResponseHopLimit() *int32 {
	return s.HttpPutResponseHopLimit
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetImageOptions() *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataImageOptions {
	return s.ImageOptions
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetImageOwnerAlias() *string {
	return s.ImageOwnerAlias
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetNetworkInterfaces() *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfaces {
	return s.NetworkInterfaces
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetSecurityGroupIds() *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityGroupIds {
	return s.SecurityGroupIds
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetSecurityOptions() *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityOptions {
	return s.SecurityOptions
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetTags() *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTags {
	return s.Tags
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetUserData() *string {
	return s.UserData
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetSystemDisk(v *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.SystemDisk = v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetAutoReleaseTime(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.AutoReleaseTime = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetAutoRenew(v bool) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.AutoRenew = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetAutoRenewPeriod(v int32) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.AutoRenewPeriod = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetCreditSpecification(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.CreditSpecification = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetDataDisks(v *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisks) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.DataDisks = v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetDeletionProtection(v bool) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetDeploymentSetId(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.DeploymentSetId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetDescription(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.Description = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetEnableVmOsConfig(v bool) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.EnableVmOsConfig = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetHostName(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.HostName = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetHttpEndpoint(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.HttpEndpoint = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetHttpPutResponseHopLimit(v int32) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.HttpPutResponseHopLimit = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetHttpTokens(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.HttpTokens = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetImageId(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.ImageId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetImageOptions(v *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataImageOptions) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.ImageOptions = v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetImageOwnerAlias(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.ImageOwnerAlias = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetInstanceChargeType(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetInstanceName(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.InstanceName = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetInstanceType(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.InstanceType = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetInternetChargeType(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetInternetMaxBandwidthIn(v int32) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetInternetMaxBandwidthOut(v int32) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetIoOptimized(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.IoOptimized = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetIpv6AddressCount(v int32) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.Ipv6AddressCount = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetKeyPairName(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.KeyPairName = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetNetworkInterfaces(v *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfaces) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.NetworkInterfaces = v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetNetworkType(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.NetworkType = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetPasswordInherit(v bool) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.PasswordInherit = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetPeriod(v int32) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.Period = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetPeriodUnit(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetPrivateIpAddress(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetRamRoleName(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.RamRoleName = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetResourceGroupId(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetSecurityEnhancementStrategy(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetSecurityGroupId(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetSecurityGroupIds(v *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityGroupIds) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetSecurityOptions(v *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityOptions) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.SecurityOptions = v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetSpotDuration(v int32) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.SpotDuration = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetSpotPriceLimit(v float32) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.SpotPriceLimit = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetSpotStrategy(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetTags(v *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTags) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.Tags = v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetUserData(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.UserData = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetVSwitchId(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.VSwitchId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetVpcId(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.VpcId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) SetZoneId(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData {
	s.ZoneId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateData) Validate() error {
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	if s.DataDisks != nil {
		if err := s.DataDisks.Validate(); err != nil {
			return err
		}
	}
	if s.ImageOptions != nil {
		if err := s.ImageOptions.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkInterfaces != nil {
		if err := s.NetworkInterfaces.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityGroupIds != nil {
		if err := s.SecurityGroupIds.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityOptions != nil {
		if err := s.SecurityOptions.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk struct {
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	BurstingEnabled      *bool   `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteWithInstance   *bool   `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DiskName             *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	Encrypted            *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	Iops                 *int32  `json:"Iops,omitempty" xml:"Iops,omitempty"`
	KMSKeyId             *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	PerformanceLevel     *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	ProvisionedIops      *int64  `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	Size                 *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) GetDescription() *string {
	return s.Description
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) GetIops() *int32 {
	return s.Iops
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) SetAutoSnapshotPolicyId(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) SetBurstingEnabled(v bool) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) SetCategory(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk {
	s.Category = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) SetDeleteWithInstance(v bool) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) SetDescription(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk {
	s.Description = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) SetDiskName(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk {
	s.DiskName = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) SetEncrypted(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk {
	s.Encrypted = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) SetIops(v int32) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk {
	s.Iops = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) SetKMSKeyId(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk {
	s.KMSKeyId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) SetPerformanceLevel(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) SetProvisionedIops(v int64) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) SetSize(v int32) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk {
	s.Size = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSystemDisk) Validate() error {
	return dara.Validate(s)
}

type DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisks struct {
	DataDisk []*DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisks) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisks) GetDataDisk() []*DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk {
	return s.DataDisk
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisks) SetDataDisk(v []*DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisks {
	s.DataDisk = v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisks) Validate() error {
	if s.DataDisk != nil {
		for _, item := range s.DataDisk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk struct {
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	BurstingEnabled      *bool   `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteWithInstance   *bool   `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Device               *string `json:"Device,omitempty" xml:"Device,omitempty"`
	DiskName             *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	Encrypted            *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	KMSKeyId             *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	PerformanceLevel     *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	ProvisionedIops      *int64  `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	Size                 *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	SnapshotId           *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) GetCategory() *string {
	return s.Category
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) GetDescription() *string {
	return s.Description
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) GetDevice() *string {
	return s.Device
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) SetAutoSnapshotPolicyId(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) SetBurstingEnabled(v bool) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) SetCategory(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk {
	s.Category = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) SetDeleteWithInstance(v bool) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) SetDescription(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk {
	s.Description = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) SetDevice(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk {
	s.Device = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) SetDiskName(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk {
	s.DiskName = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) SetEncrypted(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk {
	s.Encrypted = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) SetKMSKeyId(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk {
	s.KMSKeyId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) SetPerformanceLevel(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) SetProvisionedIops(v int64) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) SetSize(v int32) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk {
	s.Size = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) SetSnapshotId(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk {
	s.SnapshotId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataDataDisksDataDisk) Validate() error {
	return dara.Validate(s)
}

type DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataImageOptions struct {
	LoginAsNonRoot *bool `json:"LoginAsNonRoot,omitempty" xml:"LoginAsNonRoot,omitempty"`
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataImageOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataImageOptions) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataImageOptions) GetLoginAsNonRoot() *bool {
	return s.LoginAsNonRoot
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataImageOptions) SetLoginAsNonRoot(v bool) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataImageOptions {
	s.LoginAsNonRoot = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataImageOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfaces struct {
	NetworkInterface []*DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface `json:"NetworkInterface,omitempty" xml:"NetworkInterface,omitempty" type:"Repeated"`
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfaces) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfaces) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfaces) GetNetworkInterface() []*DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface {
	return s.NetworkInterface
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfaces) SetNetworkInterface(v []*DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfaces {
	s.NetworkInterface = v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfaces) Validate() error {
	if s.NetworkInterface != nil {
		for _, item := range s.NetworkInterface {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface struct {
	DeleteOnRelease             *bool                                                                                                                                                           `json:"DeleteOnRelease,omitempty" xml:"DeleteOnRelease,omitempty"`
	Description                 *string                                                                                                                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceType                *string                                                                                                                                                         `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	NetworkInterfaceName        *string                                                                                                                                                         `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	NetworkInterfaceTrafficMode *string                                                                                                                                                         `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	PrimaryIpAddress            *string                                                                                                                                                         `json:"PrimaryIpAddress,omitempty" xml:"PrimaryIpAddress,omitempty"`
	SecurityGroupId             *string                                                                                                                                                         `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupIds            *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterfaceSecurityGroupIds `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Struct"`
	VSwitchId                   *string                                                                                                                                                         `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) GetDeleteOnRelease() *bool {
	return s.DeleteOnRelease
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) GetDescription() *string {
	return s.Description
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) GetPrimaryIpAddress() *string {
	return s.PrimaryIpAddress
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) GetSecurityGroupIds() *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterfaceSecurityGroupIds {
	return s.SecurityGroupIds
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) SetDeleteOnRelease(v bool) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface {
	s.DeleteOnRelease = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) SetDescription(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface {
	s.Description = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) SetInstanceType(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface {
	s.InstanceType = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) SetNetworkInterfaceName(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface {
	s.NetworkInterfaceName = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) SetNetworkInterfaceTrafficMode(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) SetPrimaryIpAddress(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface {
	s.PrimaryIpAddress = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) SetSecurityGroupId(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) SetSecurityGroupIds(v *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterfaceSecurityGroupIds) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) SetVSwitchId(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface {
	s.VSwitchId = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterface) Validate() error {
	if s.SecurityGroupIds != nil {
		if err := s.SecurityGroupIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterfaceSecurityGroupIds struct {
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterfaceSecurityGroupIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterfaceSecurityGroupIds) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterfaceSecurityGroupIds) GetSecurityGroupId() []*string {
	return s.SecurityGroupId
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterfaceSecurityGroupIds) SetSecurityGroupId(v []*string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterfaceSecurityGroupIds {
	s.SecurityGroupId = v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataNetworkInterfacesNetworkInterfaceSecurityGroupIds) Validate() error {
	return dara.Validate(s)
}

type DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityGroupIds struct {
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityGroupIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityGroupIds) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityGroupIds) GetSecurityGroupId() []*string {
	return s.SecurityGroupId
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityGroupIds) SetSecurityGroupId(v []*string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityGroupIds {
	s.SecurityGroupId = v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityGroupIds) Validate() error {
	return dara.Validate(s)
}

type DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityOptions struct {
	TrustedSystemMode *string `json:"TrustedSystemMode,omitempty" xml:"TrustedSystemMode,omitempty"`
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityOptions) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityOptions) GetTrustedSystemMode() *string {
	return s.TrustedSystemMode
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityOptions) SetTrustedSystemMode(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityOptions {
	s.TrustedSystemMode = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataSecurityOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTags struct {
	InstanceTag []*DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTagsInstanceTag `json:"InstanceTag,omitempty" xml:"InstanceTag,omitempty" type:"Repeated"`
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTags) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTags) GetInstanceTag() []*DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTagsInstanceTag {
	return s.InstanceTag
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTags) SetInstanceTag(v []*DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTagsInstanceTag) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTags {
	s.InstanceTag = v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTags) Validate() error {
	if s.InstanceTag != nil {
		for _, item := range s.InstanceTag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTagsInstanceTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTagsInstanceTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTagsInstanceTag) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTagsInstanceTag) GetKey() *string {
	return s.Key
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTagsInstanceTag) GetValue() *string {
	return s.Value
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTagsInstanceTag) SetKey(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTagsInstanceTag {
	s.Key = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTagsInstanceTag) SetValue(v string) *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTagsInstanceTag {
	s.Value = &v
	return s
}

func (s *DescribeLaunchTemplateVersionsResponseBodyLaunchTemplateVersionSetsLaunchTemplateVersionSetLaunchTemplateDataTagsInstanceTag) Validate() error {
	return dara.Validate(s)
}
