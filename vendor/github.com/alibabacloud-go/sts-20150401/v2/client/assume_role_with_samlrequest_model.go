// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeRoleWithSAMLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDurationSeconds(v int64) *AssumeRoleWithSAMLRequest
	GetDurationSeconds() *int64
	SetPolicy(v string) *AssumeRoleWithSAMLRequest
	GetPolicy() *string
	SetRoleArn(v string) *AssumeRoleWithSAMLRequest
	GetRoleArn() *string
	SetSAMLAssertion(v string) *AssumeRoleWithSAMLRequest
	GetSAMLAssertion() *string
	SetSAMLProviderArn(v string) *AssumeRoleWithSAMLRequest
	GetSAMLProviderArn() *string
}

type AssumeRoleWithSAMLRequest struct {
	// The validity period of the STS token. Unit: seconds.
	//
	// Minimum value: 900. Maximum value: the value of the `MaxSessionDuration` parameter. Default value: 3600.
	//
	// You can call the CreateRole or UpdateRole operation to configure the `MaxSessionDuration` parameter. For more information, see [CreateRole](https://help.aliyun.com/document_detail/28710.html) or [UpdateRole](https://help.aliyun.com/document_detail/28712.html).
	//
	// example:
	//
	// 3600
	DurationSeconds *int64 `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
	// The policy that specifies the permissions of the returned STS token. You can use this parameter to grant the STS token fewer permissions than the permissions granted to the RAM role.
	//
	// 	- If you specify this parameter, the permissions of the returned STS token are the permissions that are included in the value of this parameter and owned by the RAM role.
	//
	// 	- If you do not specify this parameter, the returned STS token has all the permissions of the RAM role.
	//
	// The value must be 1 to 2,048 characters in length.
	//
	// example:
	//
	// url_encoded_policy
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The ARN of the RAM role.
	//
	// The trust entity of the RAM role is a SAML IdP. For more information, see [Create a RAM role for a trusted IdP](https://help.aliyun.com/document_detail/116805.html) or [CreateRole](https://help.aliyun.com/document_detail/28710.html).
	//
	// Format: `acs:ram::<account_id>:role/<role_name>`.
	//
	// You can view the ARN in the RAM console or by calling operations.
	//
	// 	- For more information about how to view the ARN in the RAM console, see [How do I view the ARN of the RAM role?](https://help.aliyun.com/document_detail/39744.html).
	//
	// 	- For more information about how to view the ARN by calling operations, see [ListRoles](https://help.aliyun.com/document_detail/28713.html) or [GetRole](https://help.aliyun.com/document_detail/28711.html).
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The Base64-encoded SAML assertion.
	//
	// The value must be 4 to 100,000 characters in length.
	//
	// > A complete SAML response rather than a single SAMLAssertion field must be retrieved from the external IdP.
	//
	// example:
	//
	// base64_encoded_saml_assertion
	SAMLAssertion *string `json:"SAMLAssertion,omitempty" xml:"SAMLAssertion,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the SAML IdP that is created in the RAM console.
	//
	// Format: `acs:ram::<account_id>:saml-provider/<saml_provider_id>`.
	//
	// You can view the ARN in the RAM console or by calling operations.
	//
	// 	- For more information about how to view the ARN in the RAM console, see [How do I view the ARN of a RAM role?](https://help.aliyun.com/document_detail/116795.html)
	//
	// 	- For more information about how to view the ARN by calling operations, see [GetSAMLProvider](https://help.aliyun.com/document_detail/186833.html) or [ListSAMLProviders](https://help.aliyun.com/document_detail/186851.html).
	//
	// example:
	//
	// acs:ram::123456789012****:saml-provider/company1
	SAMLProviderArn *string `json:"SAMLProviderArn,omitempty" xml:"SAMLProviderArn,omitempty"`
}

func (s AssumeRoleWithSAMLRequest) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleWithSAMLRequest) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithSAMLRequest) GetDurationSeconds() *int64 {
	return s.DurationSeconds
}

func (s *AssumeRoleWithSAMLRequest) GetPolicy() *string {
	return s.Policy
}

func (s *AssumeRoleWithSAMLRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *AssumeRoleWithSAMLRequest) GetSAMLAssertion() *string {
	return s.SAMLAssertion
}

func (s *AssumeRoleWithSAMLRequest) GetSAMLProviderArn() *string {
	return s.SAMLProviderArn
}

func (s *AssumeRoleWithSAMLRequest) SetDurationSeconds(v int64) *AssumeRoleWithSAMLRequest {
	s.DurationSeconds = &v
	return s
}

func (s *AssumeRoleWithSAMLRequest) SetPolicy(v string) *AssumeRoleWithSAMLRequest {
	s.Policy = &v
	return s
}

func (s *AssumeRoleWithSAMLRequest) SetRoleArn(v string) *AssumeRoleWithSAMLRequest {
	s.RoleArn = &v
	return s
}

func (s *AssumeRoleWithSAMLRequest) SetSAMLAssertion(v string) *AssumeRoleWithSAMLRequest {
	s.SAMLAssertion = &v
	return s
}

func (s *AssumeRoleWithSAMLRequest) SetSAMLProviderArn(v string) *AssumeRoleWithSAMLRequest {
	s.SAMLProviderArn = &v
	return s
}

func (s *AssumeRoleWithSAMLRequest) Validate() error {
	return dara.Validate(s)
}
