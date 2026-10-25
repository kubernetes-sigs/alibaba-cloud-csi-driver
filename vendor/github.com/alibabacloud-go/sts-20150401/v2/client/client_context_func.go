// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

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
func (client *Client) AssumeRoleWithContext(ctx context.Context, request *AssumeRoleRequest, runtime *dara.RuntimeOptions) (_result *AssumeRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DurationSeconds) {
		query["DurationSeconds"] = request.DurationSeconds
	}

	if !dara.IsNil(request.ExternalId) {
		query["ExternalId"] = request.ExternalId
	}

	if !dara.IsNil(request.Policy) {
		query["Policy"] = request.Policy
	}

	if !dara.IsNil(request.RoleArn) {
		query["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.RoleSessionName) {
		query["RoleSessionName"] = request.RoleSessionName
	}

	if !dara.IsNil(request.SourceIdentity) {
		query["SourceIdentity"] = request.SourceIdentity
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssumeRole"),
		Version:     dara.String("2015-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssumeRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) AssumeRoleWithOIDCWithContext(ctx context.Context, request *AssumeRoleWithOIDCRequest, runtime *dara.RuntimeOptions) (_result *AssumeRoleWithOIDCResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DurationSeconds) {
		query["DurationSeconds"] = request.DurationSeconds
	}

	if !dara.IsNil(request.OIDCProviderArn) {
		query["OIDCProviderArn"] = request.OIDCProviderArn
	}

	if !dara.IsNil(request.OIDCToken) {
		query["OIDCToken"] = request.OIDCToken
	}

	if !dara.IsNil(request.Policy) {
		query["Policy"] = request.Policy
	}

	if !dara.IsNil(request.RoleArn) {
		query["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.RoleSessionName) {
		query["RoleSessionName"] = request.RoleSessionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssumeRoleWithOIDC"),
		Version:     dara.String("2015-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssumeRoleWithOIDCResponse{}
	_body, _err := client.DoRPCRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
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
func (client *Client) AssumeRoleWithSAMLWithContext(ctx context.Context, request *AssumeRoleWithSAMLRequest, runtime *dara.RuntimeOptions) (_result *AssumeRoleWithSAMLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DurationSeconds) {
		query["DurationSeconds"] = request.DurationSeconds
	}

	if !dara.IsNil(request.Policy) {
		query["Policy"] = request.Policy
	}

	if !dara.IsNil(request.RoleArn) {
		query["RoleArn"] = request.RoleArn
	}

	if !dara.IsNil(request.SAMLAssertion) {
		query["SAMLAssertion"] = request.SAMLAssertion
	}

	if !dara.IsNil(request.SAMLProviderArn) {
		query["SAMLProviderArn"] = request.SAMLProviderArn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssumeRoleWithSAML"),
		Version:     dara.String("2015-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssumeRoleWithSAMLResponse{}
	_body, _err := client.DoRPCRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
