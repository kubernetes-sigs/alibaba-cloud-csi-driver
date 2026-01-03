// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeUserDataResponseBody
	GetInstanceId() *string
	SetRegionId(v string) *DescribeUserDataResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeUserDataResponseBody
	GetRequestId() *string
	SetUserData(v string) *DescribeUserDataResponseBody
	GetUserData() *string
}

type DescribeUserDataResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// i-bp14bnftyqhxg9ij****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The user data of the instance.
	//
	// >  If no user data is configured for the instance, an empty string is returned.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY321ABC
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s DescribeUserDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserDataResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUserDataResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserDataResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *DescribeUserDataResponseBody) SetInstanceId(v string) *DescribeUserDataResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserDataResponseBody) SetRegionId(v string) *DescribeUserDataResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeUserDataResponseBody) SetRequestId(v string) *DescribeUserDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserDataResponseBody) SetUserData(v string) *DescribeUserDataResponseBody {
	s.UserData = &v
	return s
}

func (s *DescribeUserDataResponseBody) Validate() error {
	return dara.Validate(s)
}
