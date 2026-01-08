// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPoint(v *CreateAccessPointResponseBodyAccessPoint) *CreateAccessPointResponseBody
	GetAccessPoint() *CreateAccessPointResponseBodyAccessPoint
	SetRequestId(v string) *CreateAccessPointResponseBody
	GetRequestId() *string
}

type CreateAccessPointResponseBody struct {
	// The access point.
	AccessPoint *CreateAccessPointResponseBodyAccessPoint `json:"AccessPoint,omitempty" xml:"AccessPoint,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessPointResponseBody) GetAccessPoint() *CreateAccessPointResponseBodyAccessPoint {
	return s.AccessPoint
}

func (s *CreateAccessPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccessPointResponseBody) SetAccessPoint(v *CreateAccessPointResponseBodyAccessPoint) *CreateAccessPointResponseBody {
	s.AccessPoint = v
	return s
}

func (s *CreateAccessPointResponseBody) SetRequestId(v string) *CreateAccessPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccessPointResponseBody) Validate() error {
	if s.AccessPoint != nil {
		if err := s.AccessPoint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAccessPointResponseBodyAccessPoint struct {
	// The domain name of the access point.
	//
	// example:
	//
	// ap-ie15ydanoz.001014****-w****.cn-hangzhou.nas.aliyuncs.com
	AccessPointDomain *string `json:"AccessPointDomain,omitempty" xml:"AccessPointDomain,omitempty"`
	// The ID of the access point.
	//
	// example:
	//
	// ap-ie15yd****
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
}

func (s CreateAccessPointResponseBodyAccessPoint) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessPointResponseBodyAccessPoint) GoString() string {
	return s.String()
}

func (s *CreateAccessPointResponseBodyAccessPoint) GetAccessPointDomain() *string {
	return s.AccessPointDomain
}

func (s *CreateAccessPointResponseBodyAccessPoint) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *CreateAccessPointResponseBodyAccessPoint) SetAccessPointDomain(v string) *CreateAccessPointResponseBodyAccessPoint {
	s.AccessPointDomain = &v
	return s
}

func (s *CreateAccessPointResponseBodyAccessPoint) SetAccessPointId(v string) *CreateAccessPointResponseBodyAccessPoint {
	s.AccessPointId = &v
	return s
}

func (s *CreateAccessPointResponseBodyAccessPoint) Validate() error {
	return dara.Validate(s)
}
