// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePublicIpAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *ReleasePublicIpAddressRequest
	GetDryRun() *bool
	SetInstanceId(v string) *ReleasePublicIpAddressRequest
	GetInstanceId() *string
	SetPublicIpAddress(v string) *ReleasePublicIpAddressRequest
	GetPublicIpAddress() *string
	SetRegionId(v string) *ReleasePublicIpAddressRequest
	GetRegionId() *string
}

type ReleasePublicIpAddressRequest struct {
	// > This parameter is unavailable.
	//
	// example:
	//
	// hide
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The public IP address of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ``121.40.**.**``
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReleasePublicIpAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleasePublicIpAddressRequest) GoString() string {
	return s.String()
}

func (s *ReleasePublicIpAddressRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ReleasePublicIpAddressRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleasePublicIpAddressRequest) GetPublicIpAddress() *string {
	return s.PublicIpAddress
}

func (s *ReleasePublicIpAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReleasePublicIpAddressRequest) SetDryRun(v bool) *ReleasePublicIpAddressRequest {
	s.DryRun = &v
	return s
}

func (s *ReleasePublicIpAddressRequest) SetInstanceId(v string) *ReleasePublicIpAddressRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleasePublicIpAddressRequest) SetPublicIpAddress(v string) *ReleasePublicIpAddressRequest {
	s.PublicIpAddress = &v
	return s
}

func (s *ReleasePublicIpAddressRequest) SetRegionId(v string) *ReleasePublicIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *ReleasePublicIpAddressRequest) Validate() error {
	return dara.Validate(s)
}
