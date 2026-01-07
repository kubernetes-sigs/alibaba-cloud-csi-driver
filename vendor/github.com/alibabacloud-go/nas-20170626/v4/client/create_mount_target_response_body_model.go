// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMountTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMountTargetDomain(v string) *CreateMountTargetResponseBody
	GetMountTargetDomain() *string
	SetMountTargetExtra(v *CreateMountTargetResponseBodyMountTargetExtra) *CreateMountTargetResponseBody
	GetMountTargetExtra() *CreateMountTargetResponseBodyMountTargetExtra
	SetRequestId(v string) *CreateMountTargetResponseBody
	GetRequestId() *string
}

type CreateMountTargetResponseBody struct {
	// The IPv4 domain name of the mount target.
	//
	// example:
	//
	// 174494b666-x****.cn-hangzhou.nas.aliyuncs.com
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// The information about the mount target.
	MountTargetExtra *CreateMountTargetResponseBodyMountTargetExtra `json:"MountTargetExtra,omitempty" xml:"MountTargetExtra,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 70EACC9C-D07A-4A34-ADA4-77506C42****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMountTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMountTargetResponseBody) GetMountTargetDomain() *string {
	return s.MountTargetDomain
}

func (s *CreateMountTargetResponseBody) GetMountTargetExtra() *CreateMountTargetResponseBodyMountTargetExtra {
	return s.MountTargetExtra
}

func (s *CreateMountTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMountTargetResponseBody) SetMountTargetDomain(v string) *CreateMountTargetResponseBody {
	s.MountTargetDomain = &v
	return s
}

func (s *CreateMountTargetResponseBody) SetMountTargetExtra(v *CreateMountTargetResponseBodyMountTargetExtra) *CreateMountTargetResponseBody {
	s.MountTargetExtra = v
	return s
}

func (s *CreateMountTargetResponseBody) SetRequestId(v string) *CreateMountTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMountTargetResponseBody) Validate() error {
	if s.MountTargetExtra != nil {
		if err := s.MountTargetExtra.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMountTargetResponseBodyMountTargetExtra struct {
	// The dual-stack (IPv4 and IPv6) domain name of the mount target.
	//
	// example:
	//
	// 174494b666-x****.dualstack.cn-hangzhou.nas.aliyuncs.com
	DualStackMountTargetDomain *string `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
}

func (s CreateMountTargetResponseBodyMountTargetExtra) String() string {
	return dara.Prettify(s)
}

func (s CreateMountTargetResponseBodyMountTargetExtra) GoString() string {
	return s.String()
}

func (s *CreateMountTargetResponseBodyMountTargetExtra) GetDualStackMountTargetDomain() *string {
	return s.DualStackMountTargetDomain
}

func (s *CreateMountTargetResponseBodyMountTargetExtra) SetDualStackMountTargetDomain(v string) *CreateMountTargetResponseBodyMountTargetExtra {
	s.DualStackMountTargetDomain = &v
	return s
}

func (s *CreateMountTargetResponseBodyMountTargetExtra) Validate() error {
	return dara.Validate(s)
}
