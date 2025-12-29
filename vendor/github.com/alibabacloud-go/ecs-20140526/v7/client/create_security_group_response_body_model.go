// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSecurityGroupResponseBody
	GetRequestId() *string
	SetSecurityGroupId(v string) *CreateSecurityGroupResponseBody
	GetSecurityGroupId() *string
}

type CreateSecurityGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-bp1fg655nh68xyz9****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s CreateSecurityGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecurityGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSecurityGroupResponseBody) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateSecurityGroupResponseBody) SetRequestId(v string) *CreateSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecurityGroupResponseBody) SetSecurityGroupId(v string) *CreateSecurityGroupResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateSecurityGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
