// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssumedRoleUser(v *AssumeRoleResponseBodyAssumedRoleUser) *AssumeRoleResponseBody
	GetAssumedRoleUser() *AssumeRoleResponseBodyAssumedRoleUser
	SetCredentials(v *AssumeRoleResponseBodyCredentials) *AssumeRoleResponseBody
	GetCredentials() *AssumeRoleResponseBodyCredentials
	SetRequestId(v string) *AssumeRoleResponseBody
	GetRequestId() *string
	SetSourceIdentity(v string) *AssumeRoleResponseBody
	GetSourceIdentity() *string
}

type AssumeRoleResponseBody struct {
	// The temporary identity that you use to assume the RAM role.
	AssumedRoleUser *AssumeRoleResponseBodyAssumedRoleUser `json:"AssumedRoleUser,omitempty" xml:"AssumedRoleUser,omitempty" type:"Struct"`
	// The STS credentials.
	Credentials *AssumeRoleResponseBodyCredentials `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 6894B13B-6D71-4EF5-88FA-F32781734A7F
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceIdentity *string `json:"SourceIdentity,omitempty" xml:"SourceIdentity,omitempty"`
}

func (s AssumeRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleResponseBody) GoString() string {
	return s.String()
}

func (s *AssumeRoleResponseBody) GetAssumedRoleUser() *AssumeRoleResponseBodyAssumedRoleUser {
	return s.AssumedRoleUser
}

func (s *AssumeRoleResponseBody) GetCredentials() *AssumeRoleResponseBodyCredentials {
	return s.Credentials
}

func (s *AssumeRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssumeRoleResponseBody) GetSourceIdentity() *string {
	return s.SourceIdentity
}

func (s *AssumeRoleResponseBody) SetAssumedRoleUser(v *AssumeRoleResponseBodyAssumedRoleUser) *AssumeRoleResponseBody {
	s.AssumedRoleUser = v
	return s
}

func (s *AssumeRoleResponseBody) SetCredentials(v *AssumeRoleResponseBodyCredentials) *AssumeRoleResponseBody {
	s.Credentials = v
	return s
}

func (s *AssumeRoleResponseBody) SetRequestId(v string) *AssumeRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssumeRoleResponseBody) SetSourceIdentity(v string) *AssumeRoleResponseBody {
	s.SourceIdentity = &v
	return s
}

func (s *AssumeRoleResponseBody) Validate() error {
	if s.AssumedRoleUser != nil {
		if err := s.AssumedRoleUser.Validate(); err != nil {
			return err
		}
	}
	if s.Credentials != nil {
		if err := s.Credentials.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AssumeRoleResponseBodyAssumedRoleUser struct {
	// The ARN of the temporary identity that you use to assume the RAM role.
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole/alice
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the temporary identity that you use to assume the RAM role.
	//
	// example:
	//
	// 34458433936495****:alice
	AssumedRoleId *string `json:"AssumedRoleId,omitempty" xml:"AssumedRoleId,omitempty"`
}

func (s AssumeRoleResponseBodyAssumedRoleUser) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleResponseBodyAssumedRoleUser) GoString() string {
	return s.String()
}

func (s *AssumeRoleResponseBodyAssumedRoleUser) GetArn() *string {
	return s.Arn
}

func (s *AssumeRoleResponseBodyAssumedRoleUser) GetAssumedRoleId() *string {
	return s.AssumedRoleId
}

func (s *AssumeRoleResponseBodyAssumedRoleUser) SetArn(v string) *AssumeRoleResponseBodyAssumedRoleUser {
	s.Arn = &v
	return s
}

func (s *AssumeRoleResponseBodyAssumedRoleUser) SetAssumedRoleId(v string) *AssumeRoleResponseBodyAssumedRoleUser {
	s.AssumedRoleId = &v
	return s
}

func (s *AssumeRoleResponseBodyAssumedRoleUser) Validate() error {
	return dara.Validate(s)
}

type AssumeRoleResponseBodyCredentials struct {
	// The AccessKey ID.
	//
	// example:
	//
	// STS.L4aBSCSJVMuKg5U1****
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The AccessKey secret.
	//
	// example:
	//
	// wyLTSmsyPGP1ohvvw8xYgB29dlGI8KMiH2pK****
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The time when the STS token expires. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-04-09T11:52:19Z
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The STS token.
	//
	// > Alibaba Cloud STS does not impose limits on the length of STS tokens. We strongly recommend that you do not specify a maximum length for STS tokens.
	//
	// example:
	//
	// ********
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AssumeRoleResponseBodyCredentials) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleResponseBodyCredentials) GoString() string {
	return s.String()
}

func (s *AssumeRoleResponseBodyCredentials) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *AssumeRoleResponseBodyCredentials) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *AssumeRoleResponseBodyCredentials) GetExpiration() *string {
	return s.Expiration
}

func (s *AssumeRoleResponseBodyCredentials) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AssumeRoleResponseBodyCredentials) SetAccessKeyId(v string) *AssumeRoleResponseBodyCredentials {
	s.AccessKeyId = &v
	return s
}

func (s *AssumeRoleResponseBodyCredentials) SetAccessKeySecret(v string) *AssumeRoleResponseBodyCredentials {
	s.AccessKeySecret = &v
	return s
}

func (s *AssumeRoleResponseBodyCredentials) SetExpiration(v string) *AssumeRoleResponseBodyCredentials {
	s.Expiration = &v
	return s
}

func (s *AssumeRoleResponseBodyCredentials) SetSecurityToken(v string) *AssumeRoleResponseBodyCredentials {
	s.SecurityToken = &v
	return s
}

func (s *AssumeRoleResponseBodyCredentials) Validate() error {
	return dara.Validate(s)
}
