// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeRoleWithOIDCRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDurationSeconds(v int64) *AssumeRoleWithOIDCRequest
	GetDurationSeconds() *int64
	SetOIDCProviderArn(v string) *AssumeRoleWithOIDCRequest
	GetOIDCProviderArn() *string
	SetOIDCToken(v string) *AssumeRoleWithOIDCRequest
	GetOIDCToken() *string
	SetPolicy(v string) *AssumeRoleWithOIDCRequest
	GetPolicy() *string
	SetRoleArn(v string) *AssumeRoleWithOIDCRequest
	GetRoleArn() *string
	SetRoleSessionName(v string) *AssumeRoleWithOIDCRequest
	GetRoleSessionName() *string
}

type AssumeRoleWithOIDCRequest struct {
	// The validity period of the STS token. Unit: seconds.
	//
	// Default value: 3600. Minimum value: 900. Maximum value: the value of the `MaxSessionDuration` parameter.
	//
	// For more information about how to specify `MaxSessionDuration`, see [CreateRole](https://help.aliyun.com/document_detail/28710.html) or [UpdateRole](https://help.aliyun.com/document_detail/28712.html).
	//
	// example:
	//
	// 3600
	DurationSeconds *int64 `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the OIDC IdP.
	//
	// You can view the ARN in the RAM console or by calling operations.
	//
	// 	- For more information about how to view the ARN in the RAM console, see [View the information about an OIDC IdP](https://help.aliyun.com/document_detail/327123.html).
	//
	// 	- For more information about how to view the ARN by calling operations, see [GetOIDCProvider](https://help.aliyun.com/document_detail/327126.html) or [ListOIDCProviders](https://help.aliyun.com/document_detail/327127.html).
	//
	// example:
	//
	// acs:ram::113511544585****:oidc-provider/TestOidcIdp
	OIDCProviderArn *string `json:"OIDCProviderArn,omitempty" xml:"OIDCProviderArn,omitempty"`
	// The OIDC token that is issued by the external IdP.
	//
	// The OIDC token must be 4 to 20,000 characters in length.
	//
	// > You must enter the original OIDC token. You do not need to enter the Base64-encoded OIDC token.
	//
	// example:
	//
	// eyJraWQiOiJKQzl3eHpyaHFKMGd0****
	OIDCToken *string `json:"OIDCToken,omitempty" xml:"OIDCToken,omitempty"`
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
	// {"Statement": [{"Action": ["*"],"Effect": "Allow","Resource": ["*"]}],"Version":"1"}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The ARN of the RAM role.
	//
	// You can view the ARN in the RAM console or by calling operations.
	//
	// 	- For more information about how to view the ARN in the RAM console, see [How do I view the ARN of the RAM role?](https://help.aliyun.com/document_detail/39744.html)
	//
	// 	- For more information about how to view the ARN by calling operations, see [ListRoles](https://help.aliyun.com/document_detail/28713.html) or [GetRole](https://help.aliyun.com/document_detail/28711.html).
	//
	// example:
	//
	// acs:ram::113511544585****:role/testoidc
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The custom name of the role session.
	//
	// Set this parameter based on your business requirements. In most cases, this parameter is set to the identity of the user who calls the operation, for example, the username. In ActionTrail logs, you can distinguish the users who assume the same RAM role to perform operations based on the value of the RoleSessionName parameter. This way, you can perform user-specific auditing.
	//
	// The value can contain letters, digits, periods (.), at signs (@), hyphens (-), and underscores (_).
	//
	// The value must be 2 to 64 characters in length.
	//
	// example:
	//
	// TestOidcAssumedRoleSession
	RoleSessionName *string `json:"RoleSessionName,omitempty" xml:"RoleSessionName,omitempty"`
}

func (s AssumeRoleWithOIDCRequest) String() string {
	return dara.Prettify(s)
}

func (s AssumeRoleWithOIDCRequest) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithOIDCRequest) GetDurationSeconds() *int64 {
	return s.DurationSeconds
}

func (s *AssumeRoleWithOIDCRequest) GetOIDCProviderArn() *string {
	return s.OIDCProviderArn
}

func (s *AssumeRoleWithOIDCRequest) GetOIDCToken() *string {
	return s.OIDCToken
}

func (s *AssumeRoleWithOIDCRequest) GetPolicy() *string {
	return s.Policy
}

func (s *AssumeRoleWithOIDCRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *AssumeRoleWithOIDCRequest) GetRoleSessionName() *string {
	return s.RoleSessionName
}

func (s *AssumeRoleWithOIDCRequest) SetDurationSeconds(v int64) *AssumeRoleWithOIDCRequest {
	s.DurationSeconds = &v
	return s
}

func (s *AssumeRoleWithOIDCRequest) SetOIDCProviderArn(v string) *AssumeRoleWithOIDCRequest {
	s.OIDCProviderArn = &v
	return s
}

func (s *AssumeRoleWithOIDCRequest) SetOIDCToken(v string) *AssumeRoleWithOIDCRequest {
	s.OIDCToken = &v
	return s
}

func (s *AssumeRoleWithOIDCRequest) SetPolicy(v string) *AssumeRoleWithOIDCRequest {
	s.Policy = &v
	return s
}

func (s *AssumeRoleWithOIDCRequest) SetRoleArn(v string) *AssumeRoleWithOIDCRequest {
	s.RoleArn = &v
	return s
}

func (s *AssumeRoleWithOIDCRequest) SetRoleSessionName(v string) *AssumeRoleWithOIDCRequest {
	s.RoleSessionName = &v
	return s
}

func (s *AssumeRoleWithOIDCRequest) Validate() error {
	return dara.Validate(s)
}
