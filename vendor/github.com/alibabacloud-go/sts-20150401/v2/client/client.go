// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AssumeRoleRequest struct {
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
	// The external ID of the RAM role.
	//
	// This parameter is provided by an external party and is used to prevent the confused deputy problem. For more information, see [Use ExternalId to prevent the confused deputy problem](https://help.aliyun.com/document_detail/2361741.html).
	//
	// The value must be 2 to 1,224 characters in length and can contain letters, digits, and the following special characters: `= , . @ : / - _`. The regular expression for this parameter is `[\\w+=,.@:\\/-]*`.
	//
	// example:
	//
	// abcd1234
	ExternalId *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	// The policy that specifies the permissions of the returned STS token. You can use this parameter to grant the STS token fewer permissions than the permissions granted to the RAM role.
	//
	// 	- If you specify this parameter, the permissions of the returned STS token are the permissions that are included in the value of this parameter and owned by the RAM role.
	//
	// 	- If you do not specify this parameter, the returned STS token has all the permissions of the RAM role.
	//
	// The value must be 1 to 2,048 characters in length.
	//
	// For more information about policy elements and sample policies, see [Policy elements](https://help.aliyun.com/document_detail/93738.html) and [Overview of sample policies](https://help.aliyun.com/document_detail/210969.html).
	//
	// example:
	//
	// {"Statement": [{"Action": ["*"],"Effect": "Allow","Resource": ["*"]}],"Version":"1"}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the RAM role.
	//
	// The trusted entity of the RAM role is an Alibaba Cloud account. For more information, see [Create a RAM role for a trusted Alibaba Cloud account](https://help.aliyun.com/document_detail/93691.html) or [CreateRole](https://help.aliyun.com/document_detail/28710.html).
	//
	// Format: `acs:ram::<account_id>:role/<role_name>`.
	//
	// You can view the ARN in the RAM console or by calling operations. The following items describe the validity periods of storage addresses:
	//
	// 	- For more information about how to view the ARN in the RAM console, see [How do I find the ARN of the RAM role?](https://help.aliyun.com/document_detail/39744.html)
	//
	// 	- For more information about how to view the ARN by calling operations, see [ListRoles](https://help.aliyun.com/document_detail/28713.html) or [GetRole](https://help.aliyun.com/document_detail/28711.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The custom name of the role session.
	//
	// Set this parameter based on your business requirements. In most cases, you can set this parameter to the identity of the API caller. For example, you can specify a username. You can specify `RoleSessionName` to identify API callers that assume the same RAM role in ActionTrail logs. This allows you to track the users that perform the operations.
	//
	// The value must be 2 to 64 characters in length and can contain letters, digits, and the following special characters: `. @ - _`.
	//
	// This parameter is required.
	//
	// example:
	//
	// alice
	RoleSessionName *string `json:"RoleSessionName,omitempty" xml:"RoleSessionName,omitempty"`
	SourceIdentity  *string `json:"SourceIdentity,omitempty" xml:"SourceIdentity,omitempty"`
}

func (s AssumeRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleRequest) GoString() string {
	return s.String()
}

func (s *AssumeRoleRequest) SetDurationSeconds(v int64) *AssumeRoleRequest {
	s.DurationSeconds = &v
	return s
}

func (s *AssumeRoleRequest) SetExternalId(v string) *AssumeRoleRequest {
	s.ExternalId = &v
	return s
}

func (s *AssumeRoleRequest) SetPolicy(v string) *AssumeRoleRequest {
	s.Policy = &v
	return s
}

func (s *AssumeRoleRequest) SetRoleArn(v string) *AssumeRoleRequest {
	s.RoleArn = &v
	return s
}

func (s *AssumeRoleRequest) SetRoleSessionName(v string) *AssumeRoleRequest {
	s.RoleSessionName = &v
	return s
}

func (s *AssumeRoleRequest) SetSourceIdentity(v string) *AssumeRoleRequest {
	s.SourceIdentity = &v
	return s
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
	return tea.Prettify(s)
}

func (s AssumeRoleResponseBody) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s AssumeRoleResponseBodyAssumedRoleUser) GoString() string {
	return s.String()
}

func (s *AssumeRoleResponseBodyAssumedRoleUser) SetArn(v string) *AssumeRoleResponseBodyAssumedRoleUser {
	s.Arn = &v
	return s
}

func (s *AssumeRoleResponseBodyAssumedRoleUser) SetAssumedRoleId(v string) *AssumeRoleResponseBodyAssumedRoleUser {
	s.AssumedRoleId = &v
	return s
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
	return tea.Prettify(s)
}

func (s AssumeRoleResponseBodyCredentials) GoString() string {
	return s.String()
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

type AssumeRoleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssumeRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssumeRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleResponse) GoString() string {
	return s.String()
}

func (s *AssumeRoleResponse) SetHeaders(v map[string]*string) *AssumeRoleResponse {
	s.Headers = v
	return s
}

func (s *AssumeRoleResponse) SetStatusCode(v int32) *AssumeRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *AssumeRoleResponse) SetBody(v *AssumeRoleResponseBody) *AssumeRoleResponse {
	s.Body = v
	return s
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
	return tea.Prettify(s)
}

func (s AssumeRoleWithOIDCRequest) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s AssumeRoleWithOIDCResponseBody) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s AssumeRoleWithOIDCResponseBodyAssumedRoleUser) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithOIDCResponseBodyAssumedRoleUser) SetArn(v string) *AssumeRoleWithOIDCResponseBodyAssumedRoleUser {
	s.Arn = &v
	return s
}

func (s *AssumeRoleWithOIDCResponseBodyAssumedRoleUser) SetAssumedRoleId(v string) *AssumeRoleWithOIDCResponseBodyAssumedRoleUser {
	s.AssumedRoleId = &v
	return s
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
	return tea.Prettify(s)
}

func (s AssumeRoleWithOIDCResponseBodyCredentials) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s AssumeRoleWithOIDCResponseBodyOIDCTokenInfo) GoString() string {
	return s.String()
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

type AssumeRoleWithOIDCResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssumeRoleWithOIDCResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssumeRoleWithOIDCResponse) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithOIDCResponse) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithOIDCResponse) SetHeaders(v map[string]*string) *AssumeRoleWithOIDCResponse {
	s.Headers = v
	return s
}

func (s *AssumeRoleWithOIDCResponse) SetStatusCode(v int32) *AssumeRoleWithOIDCResponse {
	s.StatusCode = &v
	return s
}

func (s *AssumeRoleWithOIDCResponse) SetBody(v *AssumeRoleWithOIDCResponseBody) *AssumeRoleWithOIDCResponse {
	s.Body = v
	return s
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
	return tea.Prettify(s)
}

func (s AssumeRoleWithSAMLRequest) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s AssumeRoleWithSAMLResponseBody) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s AssumeRoleWithSAMLResponseBodyAssumedRoleUser) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithSAMLResponseBodyAssumedRoleUser) SetArn(v string) *AssumeRoleWithSAMLResponseBodyAssumedRoleUser {
	s.Arn = &v
	return s
}

func (s *AssumeRoleWithSAMLResponseBodyAssumedRoleUser) SetAssumedRoleId(v string) *AssumeRoleWithSAMLResponseBodyAssumedRoleUser {
	s.AssumedRoleId = &v
	return s
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
	return tea.Prettify(s)
}

func (s AssumeRoleWithSAMLResponseBodyCredentials) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s AssumeRoleWithSAMLResponseBodySAMLAssertionInfo) GoString() string {
	return s.String()
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

type AssumeRoleWithSAMLResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssumeRoleWithSAMLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssumeRoleWithSAMLResponse) String() string {
	return tea.Prettify(s)
}

func (s AssumeRoleWithSAMLResponse) GoString() string {
	return s.String()
}

func (s *AssumeRoleWithSAMLResponse) SetHeaders(v map[string]*string) *AssumeRoleWithSAMLResponse {
	s.Headers = v
	return s
}

func (s *AssumeRoleWithSAMLResponse) SetStatusCode(v int32) *AssumeRoleWithSAMLResponse {
	s.StatusCode = &v
	return s
}

func (s *AssumeRoleWithSAMLResponse) SetBody(v *AssumeRoleWithSAMLResponseBody) *AssumeRoleWithSAMLResponse {
	s.Body = v
	return s
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
	return tea.Prettify(s)
}

func (s GetCallerIdentityResponseBody) GoString() string {
	return s.String()
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

type GetCallerIdentityResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCallerIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCallerIdentityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCallerIdentityResponse) GoString() string {
	return s.String()
}

func (s *GetCallerIdentityResponse) SetHeaders(v map[string]*string) *GetCallerIdentityResponse {
	s.Headers = v
	return s
}

func (s *GetCallerIdentityResponse) SetStatusCode(v int32) *GetCallerIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCallerIdentityResponse) SetBody(v *GetCallerIdentityResponseBody) *GetCallerIdentityResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-2-pop":          tea.String("sts.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("sts.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("sts.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("sts.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("sts.aliyuncs.com"),
		"cn-edge-1":                   tea.String("sts.aliyuncs.com"),
		"cn-fujian":                   tea.String("sts.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("sts.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("sts.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("sts.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("sts.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("sts.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("sts.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("sts.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("sts.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("sts.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("sts.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("sts-vpc.cn-north-2-gov-1.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("sts.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("sts.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("sts.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("sts.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("sts.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("sts-vpc.cn-shenzhen-finance-1.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("sts.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("sts.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("sts.aliyuncs.com"),
		"cn-wuhan":                    tea.String("sts.aliyuncs.com"),
		"cn-yushanfang":               tea.String("sts.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("sts.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("sts.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("sts.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("sts.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("sts.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("sts.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("sts"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a Security Token Service (STS) token to assume a Resource Access Management (RAM) role.
//
// Description:
//
// ### Prerequisites
//
// You cannot use an Alibaba Cloud account to call this operation. The requester of this operation can only be a RAM user or RAM role. Make sure that the AliyunSTSAssumeRoleAccess policy is attached to the requester. After this policy is attached to the requester, the requester has the management permissions on STS.
//
// If you do not attach the AliyunSTSAssumeRoleAccess policy to the requester, the following error message is returned:
//
// `You are not authorized to do this action. You should be authorized by RAM.`
//
// You can refer to the following information to troubleshoot the error:
//
//   - Cause of the error: The policy that is required to assume a RAM role is not attached to the requester. To resolve this issue, attach the AliyunSTSAssumeRoleAccess policy or a custom policy to the requester. For more information, see [Can I specify the RAM role that a RAM user can assume?](https://help.aliyun.com/document_detail/39744.html) and [Grant permissions to a RAM user](https://help.aliyun.com/document_detail/116146.html).
//
//   - Cause of the error: The requester is not authorized to assume the RAM role. To resolve this issue, add the requester to the Principal element in the trust policy of the RAM role For more information, see [Edit the trust policy of a RAM role](https://help.aliyun.com/document_detail/116819.html).
//
// ### Best practices
//
// An STS token is valid for a period of time after it is issued, and the number of STS tokens that can be issued within an interval is also limited. Therefore, we recommend that you configure a proper validity period for an STS token and repeatedly use the token within this period. This prevents frequent issuing of STS tokens from adversely affecting your services if a large number of requests are sent. For more information about the limit, see [Is the number of STS API requests limited?](https://help.aliyun.com/document_detail/39744.html) You can configure the `DurationSeconds` parameter to specify a validity period for an STS token.
//
// When you upload or download Object Storage Service (OSS) objects on mobile devices, a large number of STS API requests are sent. In this case, repeated use of an STS token may not meet your business requirements. To avoid the limit on STS API requests from affecting access to OSS, you can **add a signature to the URL of an OSS object**. For more information, see [Add signatures to URLs](https://help.aliyun.com/document_detail/31952.html) and [Obtain signature information from the server and upload data to OSS](https://help.aliyun.com/document_detail/31926.html).
//
// @param request - AssumeRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssumeRoleResponse
func (client *Client) AssumeRoleWithOptions(request *AssumeRoleRequest, runtime *util.RuntimeOptions) (_result *AssumeRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DurationSeconds)) {
		query["DurationSeconds"] = request.DurationSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.ExternalId)) {
		query["ExternalId"] = request.ExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.RoleArn)) {
		query["RoleArn"] = request.RoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.RoleSessionName)) {
		query["RoleSessionName"] = request.RoleSessionName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIdentity)) {
		query["SourceIdentity"] = request.SourceIdentity
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssumeRole"),
		Version:     tea.String("2015-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssumeRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a Security Token Service (STS) token to assume a Resource Access Management (RAM) role.
//
// Description:
//
// ### Prerequisites
//
// You cannot use an Alibaba Cloud account to call this operation. The requester of this operation can only be a RAM user or RAM role. Make sure that the AliyunSTSAssumeRoleAccess policy is attached to the requester. After this policy is attached to the requester, the requester has the management permissions on STS.
//
// If you do not attach the AliyunSTSAssumeRoleAccess policy to the requester, the following error message is returned:
//
// `You are not authorized to do this action. You should be authorized by RAM.`
//
// You can refer to the following information to troubleshoot the error:
//
//   - Cause of the error: The policy that is required to assume a RAM role is not attached to the requester. To resolve this issue, attach the AliyunSTSAssumeRoleAccess policy or a custom policy to the requester. For more information, see [Can I specify the RAM role that a RAM user can assume?](https://help.aliyun.com/document_detail/39744.html) and [Grant permissions to a RAM user](https://help.aliyun.com/document_detail/116146.html).
//
//   - Cause of the error: The requester is not authorized to assume the RAM role. To resolve this issue, add the requester to the Principal element in the trust policy of the RAM role For more information, see [Edit the trust policy of a RAM role](https://help.aliyun.com/document_detail/116819.html).
//
// ### Best practices
//
// An STS token is valid for a period of time after it is issued, and the number of STS tokens that can be issued within an interval is also limited. Therefore, we recommend that you configure a proper validity period for an STS token and repeatedly use the token within this period. This prevents frequent issuing of STS tokens from adversely affecting your services if a large number of requests are sent. For more information about the limit, see [Is the number of STS API requests limited?](https://help.aliyun.com/document_detail/39744.html) You can configure the `DurationSeconds` parameter to specify a validity period for an STS token.
//
// When you upload or download Object Storage Service (OSS) objects on mobile devices, a large number of STS API requests are sent. In this case, repeated use of an STS token may not meet your business requirements. To avoid the limit on STS API requests from affecting access to OSS, you can **add a signature to the URL of an OSS object**. For more information, see [Add signatures to URLs](https://help.aliyun.com/document_detail/31952.html) and [Obtain signature information from the server and upload data to OSS](https://help.aliyun.com/document_detail/31926.html).
//
// @param request - AssumeRoleRequest
//
// @return AssumeRoleResponse
func (client *Client) AssumeRole(request *AssumeRoleRequest) (_result *AssumeRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssumeRoleResponse{}
	_body, _err := client.AssumeRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a Security Token Service (STS) token to assume a Resource Access Management (RAM) role during role-based single sign-on (SSO) by using OpenID Connect (OIDC).
//
// Description:
//
// ### Prerequisites
//
//   - An OIDC token is obtained from an external identity provider (IdP).
//
//   - An OIDC IdP is created in the RAM console. For more information, see [Create an OIDC IdP](https://help.aliyun.com/document_detail/327123.html) or [CreateOIDCProvider](https://help.aliyun.com/document_detail/327135.html).
//
//   - A RAM role whose trusted entity is an OIDC IdP is created in the RAM console. For more information, see [Create a RAM role for a trusted IdP](https://help.aliyun.com/document_detail/116805.html) or [CreateRole](https://help.aliyun.com/document_detail/28710.html).
//
// @param request - AssumeRoleWithOIDCRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssumeRoleWithOIDCResponse
func (client *Client) AssumeRoleWithOIDCWithOptions(request *AssumeRoleWithOIDCRequest, runtime *util.RuntimeOptions) (_result *AssumeRoleWithOIDCResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DurationSeconds)) {
		query["DurationSeconds"] = request.DurationSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.OIDCProviderArn)) {
		query["OIDCProviderArn"] = request.OIDCProviderArn
	}

	if !tea.BoolValue(util.IsUnset(request.OIDCToken)) {
		query["OIDCToken"] = request.OIDCToken
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.RoleArn)) {
		query["RoleArn"] = request.RoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.RoleSessionName)) {
		query["RoleSessionName"] = request.RoleSessionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssumeRoleWithOIDC"),
		Version:     tea.String("2015-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssumeRoleWithOIDCResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a Security Token Service (STS) token to assume a Resource Access Management (RAM) role during role-based single sign-on (SSO) by using OpenID Connect (OIDC).
//
// Description:
//
// ### Prerequisites
//
//   - An OIDC token is obtained from an external identity provider (IdP).
//
//   - An OIDC IdP is created in the RAM console. For more information, see [Create an OIDC IdP](https://help.aliyun.com/document_detail/327123.html) or [CreateOIDCProvider](https://help.aliyun.com/document_detail/327135.html).
//
//   - A RAM role whose trusted entity is an OIDC IdP is created in the RAM console. For more information, see [Create a RAM role for a trusted IdP](https://help.aliyun.com/document_detail/116805.html) or [CreateRole](https://help.aliyun.com/document_detail/28710.html).
//
// @param request - AssumeRoleWithOIDCRequest
//
// @return AssumeRoleWithOIDCResponse
func (client *Client) AssumeRoleWithOIDC(request *AssumeRoleWithOIDCRequest) (_result *AssumeRoleWithOIDCResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssumeRoleWithOIDCResponse{}
	_body, _err := client.AssumeRoleWithOIDCWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a Security Token Service (STS) token to assume a Resource Access Management (RAM) role during role-based single sign-on (SSO) by using Security Assertion Markup Language (SAML).
//
// Description:
//
// ###
//
//   - A SAML response is obtained from an external identity provider (IdP).
//
//   - A SAML IdP is created in the RAM console. For more information, see [Create a SAML IdP](https://help.aliyun.com/document_detail/116083.html) or [CreateSAMLProvider](https://help.aliyun.com/document_detail/186846.html).
//
//   - A RAM role whose trusted entity is a SAML IdP is created in the RAM console. For more information, see [Create a RAM role for a trusted IdP](https://help.aliyun.com/document_detail/116805.html) or [CreateRole](https://help.aliyun.com/document_detail/28710.html).
//
// @param request - AssumeRoleWithSAMLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssumeRoleWithSAMLResponse
func (client *Client) AssumeRoleWithSAMLWithOptions(request *AssumeRoleWithSAMLRequest, runtime *util.RuntimeOptions) (_result *AssumeRoleWithSAMLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DurationSeconds)) {
		query["DurationSeconds"] = request.DurationSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.Policy)) {
		query["Policy"] = request.Policy
	}

	if !tea.BoolValue(util.IsUnset(request.RoleArn)) {
		query["RoleArn"] = request.RoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.SAMLAssertion)) {
		query["SAMLAssertion"] = request.SAMLAssertion
	}

	if !tea.BoolValue(util.IsUnset(request.SAMLProviderArn)) {
		query["SAMLProviderArn"] = request.SAMLProviderArn
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AssumeRoleWithSAML"),
		Version:     tea.String("2015-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssumeRoleWithSAMLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a Security Token Service (STS) token to assume a Resource Access Management (RAM) role during role-based single sign-on (SSO) by using Security Assertion Markup Language (SAML).
//
// Description:
//
// ###
//
//   - A SAML response is obtained from an external identity provider (IdP).
//
//   - A SAML IdP is created in the RAM console. For more information, see [Create a SAML IdP](https://help.aliyun.com/document_detail/116083.html) or [CreateSAMLProvider](https://help.aliyun.com/document_detail/186846.html).
//
//   - A RAM role whose trusted entity is a SAML IdP is created in the RAM console. For more information, see [Create a RAM role for a trusted IdP](https://help.aliyun.com/document_detail/116805.html) or [CreateRole](https://help.aliyun.com/document_detail/28710.html).
//
// @param request - AssumeRoleWithSAMLRequest
//
// @return AssumeRoleWithSAMLResponse
func (client *Client) AssumeRoleWithSAML(request *AssumeRoleWithSAMLRequest) (_result *AssumeRoleWithSAMLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssumeRoleWithSAMLResponse{}
	_body, _err := client.AssumeRoleWithSAMLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The ID of the Alibaba Cloud account to which the current requester belongs.
//
// @param request - GetCallerIdentityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCallerIdentityResponse
func (client *Client) GetCallerIdentityWithOptions(runtime *util.RuntimeOptions) (_result *GetCallerIdentityResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetCallerIdentity"),
		Version:     tea.String("2015-04-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCallerIdentityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the Alibaba Cloud account to which the current requester belongs.
//
// @return GetCallerIdentityResponse
func (client *Client) GetCallerIdentity() (_result *GetCallerIdentityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCallerIdentityResponse{}
	_body, _err := client.GetCallerIdentityWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
