// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandwidthPackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackages(v *DescribeBandwidthPackagesResponseBodyBandwidthPackages) *DescribeBandwidthPackagesResponseBody
	GetBandwidthPackages() *DescribeBandwidthPackagesResponseBodyBandwidthPackages
	SetPageNumber(v int32) *DescribeBandwidthPackagesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBandwidthPackagesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeBandwidthPackagesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeBandwidthPackagesResponseBody
	GetTotalCount() *int32
}

type DescribeBandwidthPackagesResponseBody struct {
	BandwidthPackages *DescribeBandwidthPackagesResponseBodyBandwidthPackages `json:"BandwidthPackages,omitempty" xml:"BandwidthPackages,omitempty" type:"Struct"`
	PageNumber        *int32                                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount        *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBandwidthPackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackagesResponseBody) GetBandwidthPackages() *DescribeBandwidthPackagesResponseBodyBandwidthPackages {
	return s.BandwidthPackages
}

func (s *DescribeBandwidthPackagesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBandwidthPackagesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBandwidthPackagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBandwidthPackagesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBandwidthPackagesResponseBody) SetBandwidthPackages(v *DescribeBandwidthPackagesResponseBodyBandwidthPackages) *DescribeBandwidthPackagesResponseBody {
	s.BandwidthPackages = v
	return s
}

func (s *DescribeBandwidthPackagesResponseBody) SetPageNumber(v int32) *DescribeBandwidthPackagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBody) SetPageSize(v int32) *DescribeBandwidthPackagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBody) SetRequestId(v string) *DescribeBandwidthPackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBody) SetTotalCount(v int32) *DescribeBandwidthPackagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBody) Validate() error {
	if s.BandwidthPackages != nil {
		if err := s.BandwidthPackages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBandwidthPackagesResponseBodyBandwidthPackages struct {
	BandwidthPackage []*DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage `json:"BandwidthPackage,omitempty" xml:"BandwidthPackage,omitempty" type:"Repeated"`
}

func (s DescribeBandwidthPackagesResponseBodyBandwidthPackages) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackagesResponseBodyBandwidthPackages) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackages) GetBandwidthPackage() []*DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage {
	return s.BandwidthPackage
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackages) SetBandwidthPackage(v []*DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) *DescribeBandwidthPackagesResponseBodyBandwidthPackages {
	s.BandwidthPackage = v
	return s
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackages) Validate() error {
	if s.BandwidthPackage != nil {
		for _, item := range s.BandwidthPackage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage struct {
	Bandwidth          *string                                                                                  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BandwidthPackageId *string                                                                                  `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	BusinessStatus     *string                                                                                  `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	CreationTime       *string                                                                                  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description        *string                                                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	ISP                *string                                                                                  `json:"ISP,omitempty" xml:"ISP,omitempty"`
	InstanceChargeType *string                                                                                  `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InternetChargeType *string                                                                                  `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	IpCount            *string                                                                                  `json:"IpCount,omitempty" xml:"IpCount,omitempty"`
	Name               *string                                                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	NatGatewayId       *string                                                                                  `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	PublicIpAddresses  *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddresses `json:"PublicIpAddresses,omitempty" xml:"PublicIpAddresses,omitempty" type:"Struct"`
	RegionId           *string                                                                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status             *string                                                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	ZoneId             *string                                                                                  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) GetDescription() *string {
	return s.Description
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) GetISP() *string {
	return s.ISP
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) GetIpCount() *string {
	return s.IpCount
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) GetName() *string {
	return s.Name
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) GetPublicIpAddresses() *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddresses {
	return s.PublicIpAddresses
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) GetStatus() *string {
	return s.Status
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) SetBandwidth(v string) *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) SetBandwidthPackageId(v string) *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) SetBusinessStatus(v string) *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) SetCreationTime(v string) *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage {
	s.CreationTime = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) SetDescription(v string) *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage {
	s.Description = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) SetISP(v string) *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage {
	s.ISP = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) SetInstanceChargeType(v string) *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) SetInternetChargeType(v string) *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) SetIpCount(v string) *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage {
	s.IpCount = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) SetName(v string) *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage {
	s.Name = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) SetNatGatewayId(v string) *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) SetPublicIpAddresses(v *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddresses) *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage {
	s.PublicIpAddresses = v
	return s
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) SetRegionId(v string) *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage {
	s.RegionId = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) SetStatus(v string) *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage {
	s.Status = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) SetZoneId(v string) *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage {
	s.ZoneId = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackage) Validate() error {
	if s.PublicIpAddresses != nil {
		if err := s.PublicIpAddresses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddresses struct {
	PublicIpAddresse []*DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddressesPublicIpAddresse `json:"PublicIpAddresse,omitempty" xml:"PublicIpAddresse,omitempty" type:"Repeated"`
}

func (s DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddresses) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddresses) GetPublicIpAddresse() []*DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddressesPublicIpAddresse {
	return s.PublicIpAddresse
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddresses) SetPublicIpAddresse(v []*DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddressesPublicIpAddresse) *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddresses {
	s.PublicIpAddresse = v
	return s
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddresses) Validate() error {
	if s.PublicIpAddresse != nil {
		for _, item := range s.PublicIpAddresse {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddressesPublicIpAddresse struct {
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	IpAddress    *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
}

func (s DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddressesPublicIpAddresse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddressesPublicIpAddresse) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddressesPublicIpAddresse) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddressesPublicIpAddresse) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddressesPublicIpAddresse) SetAllocationId(v string) *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddressesPublicIpAddresse {
	s.AllocationId = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddressesPublicIpAddresse) SetIpAddress(v string) *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddressesPublicIpAddresse {
	s.IpAddress = &v
	return s
}

func (s *DescribeBandwidthPackagesResponseBodyBandwidthPackagesBandwidthPackagePublicIpAddressesPublicIpAddresse) Validate() error {
	return dara.Validate(s)
}
