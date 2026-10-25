// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCallerIdentityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *GetCallerIdentityResponseBody
	GetAccountId() *string
	SetArn(v string) *GetCallerIdentityResponseBody
	GetArn() *string
	SetIdentityType(v string) *GetCallerIdentityResponseBody
	GetIdentityType() *string
	SetPrincipalId(v string) *GetCallerIdentityResponseBody
	GetPrincipalId() *string
	SetRequestId(v string) *GetCallerIdentityResponseBody
	GetRequestId() *string
	SetRoleId(v string) *GetCallerIdentityResponseBody
	GetRoleId() *string
	SetUserId(v string) *GetCallerIdentityResponseBody
	GetUserId() *string
}

type GetCallerIdentityResponseBody struct {
	// example:
	//
	// 196813200012****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// acs:ram::196813200012****:user/admin
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// example:
	//
	// RAMUser
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// example:
	//
	// 28877424437521****
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// example:
	//
	// 3C87BF47-3724-5443-ADC1-5AEAD9A03EB1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 33537620082992****
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// example:
	//
	// 216959339000****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetCallerIdentityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCallerIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *GetCallerIdentityResponseBody) GetAccountId() *string {
	return s.AccountId
}

func (s *GetCallerIdentityResponseBody) GetArn() *string {
	return s.Arn
}

func (s *GetCallerIdentityResponseBody) GetIdentityType() *string {
	return s.IdentityType
}

func (s *GetCallerIdentityResponseBody) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *GetCallerIdentityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCallerIdentityResponseBody) GetRoleId() *string {
	return s.RoleId
}

func (s *GetCallerIdentityResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetCallerIdentityResponseBody) SetAccountId(v string) *GetCallerIdentityResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetCallerIdentityResponseBody) SetArn(v string) *GetCallerIdentityResponseBody {
	s.Arn = &v
	return s
}

func (s *GetCallerIdentityResponseBody) SetIdentityType(v string) *GetCallerIdentityResponseBody {
	s.IdentityType = &v
	return s
}

func (s *GetCallerIdentityResponseBody) SetPrincipalId(v string) *GetCallerIdentityResponseBody {
	s.PrincipalId = &v
	return s
}

func (s *GetCallerIdentityResponseBody) SetRequestId(v string) *GetCallerIdentityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCallerIdentityResponseBody) SetRoleId(v string) *GetCallerIdentityResponseBody {
	s.RoleId = &v
	return s
}

func (s *GetCallerIdentityResponseBody) SetUserId(v string) *GetCallerIdentityResponseBody {
	s.UserId = &v
	return s
}

func (s *GetCallerIdentityResponseBody) Validate() error {
	return dara.Validate(s)
}
