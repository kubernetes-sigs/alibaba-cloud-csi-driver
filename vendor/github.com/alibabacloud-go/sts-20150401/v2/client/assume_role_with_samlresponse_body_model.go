// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeRoleWithSAMLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssumedRoleUser(v *AssumeRoleWithSAMLResponseBodyAssumedRoleUser) *AssumeRoleWithSAMLResponseBody
	GetAssumedRoleUser() *AssumeRoleWithSAMLResponseBodyAssumedRoleUser
	SetCredentials(v *AssumeRoleWithSAMLResponseBodyCredentials) *AssumeRoleWithSAMLResponseBody
	GetCredentials() *AssumeRoleWithSAMLResponseBodyCredentials
	SetRequestId(v string) *AssumeRoleWithSAMLResponseBody
	GetRequestId() *string
	SetSAMLAssertionInfo(v *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) *AssumeRoleWithSAMLResponseBody
	GetSAMLAssertionInfo() *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo
	SetSourceIdentity(v string) *AssumeRoleWithSAMLResponseBody
	GetSourceIdentity() *string
}

type AssumeRoleWithSAMLResponseBody struct {
	// The temporary identity that you use to assume the RAM role.
	AssumedRoleUser *AssumeRoleWithSAMLResponseBodyAssumedRoleUser `json:"AssumedRoleUser,omitempty" xml:"AssumedRoleUser,omitempty" type:"Struct"`
	// The STS credentials.
	Credentials *AssumeRoleWithSAMLResponseBodyCredentials `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 6894B13B-6D71-4EF5-88FA-F32781734A7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information in the SAML assertion.
	SAMLAssertionInfo *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo `json:"SAMLAssertionInfo,omitempty" xml:"SAMLAssertionInfo,omitempty" type:"Struct"`
	SourceIdentity    *string                                          `json:"SourceIdentity,omitempty" xml:"SourceIdentity,omitempty"`
}

func (s AssumeRoleWithSAMLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleWithSAMLResponseBody) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithSAMLResponseBody) GetAssumedRoleUser() *AssumeRoleWithSAMLResponseBodyAssumedRoleUser {
	return s.AssumedRoleUser
}

func (s *AssumeRoleWithSAMLResponseBody) GetCredentials() *AssumeRoleWithSAMLResponseBodyCredentials {
	return s.Credentials
}

func (s *AssumeRoleWithSAMLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssumeRoleWithSAMLResponseBody) GetSAMLAssertionInfo() *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo {
	return s.SAMLAssertionInfo
}

func (s *AssumeRoleWithSAMLResponseBody) GetSourceIdentity() *string {
	return s.SourceIdentity
}

func (s *AssumeRoleWithSAMLResponseBody) SetAssumedRoleUser(v *AssumeRoleWithSAMLResponseBodyAssumedRoleUser) *AssumeRoleWithSAMLResponseBody {
	s.AssumedRoleUser = v
	return s
}

func (s *AssumeRoleWithSAMLResponseBody) SetCredentials(v *AssumeRoleWithSAMLResponseBodyCredentials) *AssumeRoleWithSAMLResponseBody {
	s.Credentials = v
	return s
}

func (s *AssumeRoleWithSAMLResponseBody) SetRequestId(v string) *AssumeRoleWithSAMLResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBody) SetSAMLAssertionInfo(v *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) *AssumeRoleWithSAMLResponseBody {
	s.SAMLAssertionInfo = v
	return s
}

func (s *AssumeRoleWithSAMLResponseBody) SetSourceIdentity(v string) *AssumeRoleWithSAMLResponseBody {
	s.SourceIdentity = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBody) Validate() error {
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
	if s.SAMLAssertionInfo != nil {
		if err := s.SAMLAssertionInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AssumeRoleWithSAMLResponseBodyAssumedRoleUser struct {
	// The ARN of the temporary identity that you use to assume the RAM role.
	//
	// example:
	//
	// acs:sts::123456789012****:assumed-role/AdminRole/alice
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the temporary identity that you use to assume the RAM role.
	//
	// example:
	//
	// 34458433936495****:alice
	AssumedRoleId *string `json:"AssumedRoleId,omitempty" xml:"AssumedRoleId,omitempty"`
}

func (s AssumeRoleWithSAMLResponseBodyAssumedRoleUser) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleWithSAMLResponseBodyAssumedRoleUser) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithSAMLResponseBodyAssumedRoleUser) GetArn() *string {
	return s.Arn
}

func (s *AssumeRoleWithSAMLResponseBodyAssumedRoleUser) GetAssumedRoleId() *string {
	return s.AssumedRoleId
}

func (s *AssumeRoleWithSAMLResponseBodyAssumedRoleUser) SetArn(v string) *AssumeRoleWithSAMLResponseBodyAssumedRoleUser {
	s.Arn = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodyAssumedRoleUser) SetAssumedRoleId(v string) *AssumeRoleWithSAMLResponseBodyAssumedRoleUser {
	s.AssumedRoleId = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodyAssumedRoleUser) Validate() error {
	return dara.Validate(s)
}

type AssumeRoleWithSAMLResponseBodyCredentials struct {
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

func (s AssumeRoleWithSAMLResponseBodyCredentials) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleWithSAMLResponseBodyCredentials) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithSAMLResponseBodyCredentials) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *AssumeRoleWithSAMLResponseBodyCredentials) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *AssumeRoleWithSAMLResponseBodyCredentials) GetExpiration() *string {
	return s.Expiration
}

func (s *AssumeRoleWithSAMLResponseBodyCredentials) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AssumeRoleWithSAMLResponseBodyCredentials) SetAccessKeyId(v string) *AssumeRoleWithSAMLResponseBodyCredentials {
	s.AccessKeyId = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodyCredentials) SetAccessKeySecret(v string) *AssumeRoleWithSAMLResponseBodyCredentials {
	s.AccessKeySecret = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodyCredentials) SetExpiration(v string) *AssumeRoleWithSAMLResponseBodyCredentials {
	s.Expiration = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodyCredentials) SetSecurityToken(v string) *AssumeRoleWithSAMLResponseBodyCredentials {
	s.SecurityToken = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodyCredentials) Validate() error {
	return dara.Validate(s)
}

type AssumeRoleWithSAMLResponseBodySAMLAssertionInfo struct {
	// The value in the `Issuer` element in the SAML assertion.
	//
	// example:
	//
	// http://example.com/adfs/services/trust
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The `Recipient` attribute of the SubjectConfirmationData sub-element. SubjectConfirmationData is a sub-element of the `Subject` element in the SAML assertion.
	//
	// example:
	//
	// https://signin.aliyun.com/saml-role/SSO
	Recipient *string `json:"Recipient,omitempty" xml:"Recipient,omitempty"`
	// The value in the NameID sub-element of the `Subject` element in the SAML assertion.
	//
	// example:
	//
	// alice@example.com
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The Format attribute of the `NameID` element in the SAML assertion. If the Format attribute is prefixed with `urn:oasis:names:tc:SAML:2.0:nameid-format:`, the prefix is not included in the value of this parameter. For example, if the value of the Format attribute is urn:oasis:names:tc:SAML:2.0:nameid-format:persistent/transient, the value of this parameter is `persistent/transient`.
	//
	// example:
	//
	// persistent
	SubjectType *string `json:"SubjectType,omitempty" xml:"SubjectType,omitempty"`
}

func (s AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) GetIssuer() *string {
	return s.Issuer
}

func (s *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) GetRecipient() *string {
	return s.Recipient
}

func (s *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) GetSubject() *string {
	return s.Subject
}

func (s *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) GetSubjectType() *string {
	return s.SubjectType
}

func (s *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) SetIssuer(v string) *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo {
	s.Issuer = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) SetRecipient(v string) *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo {
	s.Recipient = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) SetSubject(v string) *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo {
	s.Subject = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) SetSubjectType(v string) *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo {
	s.SubjectType = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) Validate() error {
	return dara.Validate(s)
}
