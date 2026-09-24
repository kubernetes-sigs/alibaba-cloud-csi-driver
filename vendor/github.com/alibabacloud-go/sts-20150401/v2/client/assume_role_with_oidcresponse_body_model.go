// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeRoleWithOIDCResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssumedRoleUser(v *AssumeRoleWithOIDCResponseBodyAssumedRoleUser) *AssumeRoleWithOIDCResponseBody
	GetAssumedRoleUser() *AssumeRoleWithOIDCResponseBodyAssumedRoleUser
	SetCredentials(v *AssumeRoleWithOIDCResponseBodyCredentials) *AssumeRoleWithOIDCResponseBody
	GetCredentials() *AssumeRoleWithOIDCResponseBodyCredentials
	SetOIDCTokenInfo(v *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) *AssumeRoleWithOIDCResponseBody
	GetOIDCTokenInfo() *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo
	SetRequestId(v string) *AssumeRoleWithOIDCResponseBody
	GetRequestId() *string
	SetSourceIdentity(v string) *AssumeRoleWithOIDCResponseBody
	GetSourceIdentity() *string
}

type AssumeRoleWithOIDCResponseBody struct {
	// The temporary identity that you use to assume the RAM role.
	AssumedRoleUser *AssumeRoleWithOIDCResponseBodyAssumedRoleUser `json:"AssumedRoleUser,omitempty" xml:"AssumedRoleUser,omitempty" type:"Struct"`
	// The access credentials.
	Credentials *AssumeRoleWithOIDCResponseBodyCredentials `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Struct"`
	// The information about the OIDC token.
	OIDCTokenInfo *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo `json:"OIDCTokenInfo,omitempty" xml:"OIDCTokenInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 3D57EAD2-8723-1F26-B69C-F8707D8B565D
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceIdentity *string `json:"SourceIdentity,omitempty" xml:"SourceIdentity,omitempty"`
}

func (s AssumeRoleWithOIDCResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleWithOIDCResponseBody) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithOIDCResponseBody) GetAssumedRoleUser() *AssumeRoleWithOIDCResponseBodyAssumedRoleUser {
	return s.AssumedRoleUser
}

func (s *AssumeRoleWithOIDCResponseBody) GetCredentials() *AssumeRoleWithOIDCResponseBodyCredentials {
	return s.Credentials
}

func (s *AssumeRoleWithOIDCResponseBody) GetOIDCTokenInfo() *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo {
	return s.OIDCTokenInfo
}

func (s *AssumeRoleWithOIDCResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssumeRoleWithOIDCResponseBody) GetSourceIdentity() *string {
	return s.SourceIdentity
}

func (s *AssumeRoleWithOIDCResponseBody) SetAssumedRoleUser(v *AssumeRoleWithOIDCResponseBodyAssumedRoleUser) *AssumeRoleWithOIDCResponseBody {
	s.AssumedRoleUser = v
	return s
}

func (s *AssumeRoleWithOIDCResponseBody) SetCredentials(v *AssumeRoleWithOIDCResponseBodyCredentials) *AssumeRoleWithOIDCResponseBody {
	s.Credentials = v
	return s
}

func (s *AssumeRoleWithOIDCResponseBody) SetOIDCTokenInfo(v *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) *AssumeRoleWithOIDCResponseBody {
	s.OIDCTokenInfo = v
	return s
}

func (s *AssumeRoleWithOIDCResponseBody) SetRequestId(v string) *AssumeRoleWithOIDCResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBody) SetSourceIdentity(v string) *AssumeRoleWithOIDCResponseBody {
	s.SourceIdentity = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBody) Validate() error {
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
	if s.OIDCTokenInfo != nil {
		if err := s.OIDCTokenInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AssumeRoleWithOIDCResponseBodyAssumedRoleUser struct {
	// The ARN of the temporary identity that you use to assume the RAM role.
	//
	// example:
	//
	// acs:ram::113511544585****:role/testoidc/TestOidcAssumedRoleSession
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the temporary identity that you use to assume the RAM role.
	//
	// example:
	//
	// 33157794895460****
	AssumedRoleId *string `json:"AssumedRoleId,omitempty" xml:"AssumedRoleId,omitempty"`
}

func (s AssumeRoleWithOIDCResponseBodyAssumedRoleUser) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleWithOIDCResponseBodyAssumedRoleUser) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithOIDCResponseBodyAssumedRoleUser) GetArn() *string {
	return s.Arn
}

func (s *AssumeRoleWithOIDCResponseBodyAssumedRoleUser) GetAssumedRoleId() *string {
	return s.AssumedRoleId
}

func (s *AssumeRoleWithOIDCResponseBodyAssumedRoleUser) SetArn(v string) *AssumeRoleWithOIDCResponseBodyAssumedRoleUser {
	s.Arn = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyAssumedRoleUser) SetAssumedRoleId(v string) *AssumeRoleWithOIDCResponseBodyAssumedRoleUser {
	s.AssumedRoleId = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyAssumedRoleUser) Validate() error {
	return dara.Validate(s)
}

type AssumeRoleWithOIDCResponseBodyCredentials struct {
	// The AccessKey ID.
	//
	// example:
	//
	// STS.NUgYrLnoC37mZZCNnAbez****
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The AccessKey secret.
	//
	// example:
	//
	// CVwjCkNzTMupZ8NbTCxCBRq3K16jtcWFTJAyBEv2****
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The time when the STS token expires. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-10-20T04:27:09Z
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The STS token.
	//
	// > Alibaba Cloud STS does not impose limits on the length of STS tokens. We strongly recommend that you do not specify a maximum length for STS tokens.
	//
	// example:
	//
	// CAIShwJ1q6Ft5B2yfSjIr5bSEsj4g7BihPWGWHz****
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AssumeRoleWithOIDCResponseBodyCredentials) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleWithOIDCResponseBodyCredentials) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithOIDCResponseBodyCredentials) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *AssumeRoleWithOIDCResponseBodyCredentials) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *AssumeRoleWithOIDCResponseBodyCredentials) GetExpiration() *string {
	return s.Expiration
}

func (s *AssumeRoleWithOIDCResponseBodyCredentials) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AssumeRoleWithOIDCResponseBodyCredentials) SetAccessKeyId(v string) *AssumeRoleWithOIDCResponseBodyCredentials {
	s.AccessKeyId = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyCredentials) SetAccessKeySecret(v string) *AssumeRoleWithOIDCResponseBodyCredentials {
	s.AccessKeySecret = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyCredentials) SetExpiration(v string) *AssumeRoleWithOIDCResponseBodyCredentials {
	s.Expiration = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyCredentials) SetSecurityToken(v string) *AssumeRoleWithOIDCResponseBodyCredentials {
	s.SecurityToken = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyCredentials) Validate() error {
	return dara.Validate(s)
}

type AssumeRoleWithOIDCResponseBodyOIDCTokenInfo struct {
	// The audience. If multiple audiences are returned, the audiences are separated by commas (,).
	//
	// The audience is represented by the `aud` field in the OIDC Token.
	//
	// example:
	//
	// 496271242565057****
	ClientIds *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	// The time when the OIDC token expires.
	//
	// example:
	//
	// 2021-10-20T04:27:09Z
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The time when the OIDC token was issued.
	//
	// example:
	//
	// 2021-10-20T03:27:09Z
	IssuanceTime *string `json:"IssuanceTime,omitempty" xml:"IssuanceTime,omitempty"`
	// The URL of the issuer,
	//
	// which is represented by the `iss` field in the OIDC Token.
	//
	// example:
	//
	// https://dev-xxxxxx.okta.com
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The subject,
	//
	// which is represented by the `sub` field in the OIDC Token.
	//
	// example:
	//
	// KryrkIdjylZb7agUgCEf****
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The verification information about the OIDC token. For more information, see [Manage an OIDC IdP](https://help.aliyun.com/document_detail/327123.html).
	//
	// example:
	//
	// Success
	VerificationInfo *string `json:"VerificationInfo,omitempty" xml:"VerificationInfo,omitempty"`
}

func (s AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) GetClientIds() *string {
	return s.ClientIds
}

func (s *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) GetExpirationTime() *string {
	return s.ExpirationTime
}

func (s *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) GetIssuanceTime() *string {
	return s.IssuanceTime
}

func (s *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) GetIssuer() *string {
	return s.Issuer
}

func (s *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) GetSubject() *string {
	return s.Subject
}

func (s *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) GetVerificationInfo() *string {
	return s.VerificationInfo
}

func (s *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) SetClientIds(v string) *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo {
	s.ClientIds = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) SetExpirationTime(v string) *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo {
	s.ExpirationTime = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) SetIssuanceTime(v string) *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo {
	s.IssuanceTime = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) SetIssuer(v string) *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo {
	s.Issuer = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) SetSubject(v string) *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo {
	s.Subject = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) SetVerificationInfo(v string) *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo {
	s.VerificationInfo = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) Validate() error {
	return dara.Validate(s)
}
