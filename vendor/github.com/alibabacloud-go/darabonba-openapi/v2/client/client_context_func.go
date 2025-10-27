// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	"encoding/hex"

	spi "github.com/alibabacloud-go/alibabacloud-gateway-spi/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Description:
//
// # Encapsulate the request and invoke the network
//
// @param action - api name
//
// @param version - product version
//
// @param protocol - http or https
//
// @param method - e.g. GET
//
// @param authType - authorization type e.g. AK
//
// @param bodyType - response body type e.g. String
//
// @param request - object of OpenApiRequest
//
// @param runtime - which controls some details of call api, such as retry times
//
// @return the response
func (client *Client) DoRPCRequestWithCtx(ctx context.Context, action *string, version *string, protocol *string, method *string, authType *string, bodyType *string, request *openapiutil.OpenApiRequest, runtime *dara.RuntimeOptions) (_result map[string]interface{}, _err error) {
	_runtime := dara.NewRuntimeObject(map[string]interface{}{
		"key":            dara.ToString(dara.Default(dara.StringValue(runtime.Key), dara.StringValue(client.Key))),
		"cert":           dara.ToString(dara.Default(dara.StringValue(runtime.Cert), dara.StringValue(client.Cert))),
		"ca":             dara.ToString(dara.Default(dara.StringValue(runtime.Ca), dara.StringValue(client.Ca))),
		"readTimeout":    dara.ForceInt(dara.Default(dara.IntValue(runtime.ReadTimeout), dara.IntValue(client.ReadTimeout))),
		"connectTimeout": dara.ForceInt(dara.Default(dara.IntValue(runtime.ConnectTimeout), dara.IntValue(client.ConnectTimeout))),
		"httpProxy":      dara.ToString(dara.Default(dara.StringValue(runtime.HttpProxy), dara.StringValue(client.HttpProxy))),
		"httpsProxy":     dara.ToString(dara.Default(dara.StringValue(runtime.HttpsProxy), dara.StringValue(client.HttpsProxy))),
		"noProxy":        dara.ToString(dara.Default(dara.StringValue(runtime.NoProxy), dara.StringValue(client.NoProxy))),
		"socks5Proxy":    dara.ToString(dara.Default(dara.StringValue(runtime.Socks5Proxy), dara.StringValue(client.Socks5Proxy))),
		"socks5NetWork":  dara.ToString(dara.Default(dara.StringValue(runtime.Socks5NetWork), dara.StringValue(client.Socks5NetWork))),
		"maxIdleConns":   dara.ForceInt(dara.Default(dara.IntValue(runtime.MaxIdleConns), dara.IntValue(client.MaxIdleConns))),
		"retryOptions":   client.RetryOptions,
		"ignoreSSL":      dara.BoolValue(runtime.IgnoreSSL),
		"httpClient":     client.HttpClient,
		"tlsMinVersion":  dara.StringValue(client.TlsMinVersion),
	})

	var retryPolicyContext *dara.RetryPolicyContext
	var request_ *dara.Request
	var response_ *dara.Response
	var _resultErr error
	retriesAttempted := int(0)
	retryPolicyContext = &dara.RetryPolicyContext{
		RetriesAttempted: retriesAttempted,
	}

	_result = make(map[string]interface{})
	for dara.ShouldRetry(_runtime.RetryOptions, retryPolicyContext) {
		_resultErr = nil
		_backoffDelayTime := dara.GetBackoffDelay(_runtime.RetryOptions, retryPolicyContext)
		dara.Sleep(_backoffDelayTime)

		request_ = dara.NewRequest()
		request_.Protocol = dara.String(dara.ToString(dara.Default(dara.StringValue(client.Protocol), dara.StringValue(protocol))))
		request_.Method = method
		request_.Pathname = dara.String("/")
		globalQueries := make(map[string]*string)
		globalHeaders := make(map[string]*string)
		if !dara.IsNil(client.GlobalParameters) {
			globalParams := client.GlobalParameters
			if !dara.IsNil(globalParams.Queries) {
				globalQueries = globalParams.Queries
			}

			if !dara.IsNil(globalParams.Headers) {
				globalHeaders = globalParams.Headers
			}

		}

		extendsHeaders := make(map[string]*string)
		extendsQueries := make(map[string]*string)
		if !dara.IsNil(runtime.ExtendsParameters) {
			extendsParameters := runtime.ExtendsParameters
			if !dara.IsNil(extendsParameters.Headers) {
				extendsHeaders = extendsParameters.Headers
			}

			if !dara.IsNil(extendsParameters.Queries) {
				extendsQueries = extendsParameters.Queries
			}

		}

		request_.Query = dara.Merge(map[string]*string{
			"Action":         action,
			"Format":         dara.String("json"),
			"Version":        version,
			"Timestamp":      openapiutil.GetTimestamp(),
			"SignatureNonce": openapiutil.GetNonce(),
		}, globalQueries,
			extendsQueries,
			request.Query)
		headers, _err := client.GetRpcHeaders()
		if _err != nil {
			retriesAttempted++
			retryPolicyContext = &dara.RetryPolicyContext{
				RetriesAttempted: retriesAttempted,
				HttpRequest:      request_,
				HttpResponse:     response_,
				Exception:        _err,
			}
			_resultErr = _err
			continue
		}

		if dara.IsNil(headers) {
			// endpoint is setted in product client
			request_.Headers = dara.Merge(map[string]*string{
				"host":          client.Endpoint,
				"x-acs-version": version,
				"x-acs-action":  action,
				"user-agent":    openapiutil.GetUserAgent(client.UserAgent),
			}, globalHeaders,
				extendsHeaders,
				request.Headers)
		} else {
			request_.Headers = dara.Merge(map[string]*string{
				"host":          client.Endpoint,
				"x-acs-version": version,
				"x-acs-action":  action,
				"user-agent":    openapiutil.GetUserAgent(client.UserAgent),
			}, globalHeaders,
				extendsHeaders,
				request.Headers,
				headers)
		}

		if !dara.IsNil(request.Body) {
			m := dara.ToMap(request.Body)
			tmp := dara.ToMap(openapiutil.Query(m))
			request_.Body = dara.ToReader(dara.ToFormString(tmp))
			request_.Headers["content-type"] = dara.String("application/x-www-form-urlencoded")
		}

		if dara.StringValue(authType) != "Anonymous" {
			if dara.IsNil(client.Credential) {
				_err = &ClientError{
					Code:    dara.String("InvalidCredentials"),
					Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
				}
				if dara.BoolValue(client.DisableSDKError) != true {
					_err = dara.TeaSDKError(_err)
				}
				return _result, _err
			}

			credentialModel, _err := client.Credential.GetCredential()
			if _err != nil {
				retriesAttempted++
				retryPolicyContext = &dara.RetryPolicyContext{
					RetriesAttempted: retriesAttempted,
					HttpRequest:      request_,
					HttpResponse:     response_,
					Exception:        _err,
				}
				_resultErr = _err
				continue
			}

			if !dara.IsNil(credentialModel.ProviderName) {
				request_.Headers["x-acs-credentials-provider"] = credentialModel.ProviderName
			}

			credentialType := dara.StringValue(credentialModel.Type)
			if credentialType == "bearer" {
				bearerToken := dara.StringValue(credentialModel.BearerToken)
				request_.Query["BearerToken"] = dara.String(bearerToken)
				request_.Query["SignatureType"] = dara.String("BEARERTOKEN")
			} else if credentialType == "id_token" {
				idToken := dara.StringValue(credentialModel.SecurityToken)
				request_.Headers["x-acs-zero-trust-idtoken"] = dara.String(idToken)
			} else {
				accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
				accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
				securityToken := dara.StringValue(credentialModel.SecurityToken)
				if !dara.IsNil(dara.String(securityToken)) && securityToken != "" {
					request_.Query["SecurityToken"] = dara.String(securityToken)
				}

				request_.Query["SignatureMethod"] = dara.String("HMAC-SHA1")
				request_.Query["SignatureVersion"] = dara.String("1.0")
				request_.Query["AccessKeyId"] = dara.String(accessKeyId)
				var t map[string]interface{}
				if !dara.IsNil(request.Body) {
					t = dara.ToMap(request.Body)
				}

				signedParam := dara.Merge(request_.Query,
					openapiutil.Query(t))
				request_.Query["Signature"] = openapiutil.GetRPCSignature(signedParam, request_.Method, dara.String(accessKeySecret))
			}

		}

		response_, _err := dara.DoRequestWithCtx(ctx, request_, _runtime)
		if _err != nil {
			retriesAttempted++
			retryPolicyContext = &dara.RetryPolicyContext{
				RetriesAttempted: retriesAttempted,
				HttpRequest:      request_,
				HttpResponse:     response_,
				Exception:        _err,
			}
			_resultErr = _err
			continue
		}

		_result, _err = doRPCRequestWithCtx_opResponse(response_, client, bodyType)
		if _err != nil {
			retriesAttempted++
			retryPolicyContext = &dara.RetryPolicyContext{
				RetriesAttempted: retriesAttempted,
				HttpRequest:      request_,
				HttpResponse:     response_,
				Exception:        _err,
			}
			_resultErr = _err
			continue
		}

		return _result, _err
	}
	if dara.BoolValue(client.DisableSDKError) != true {
		_resultErr = dara.TeaSDKError(_resultErr)
	}
	return _result, _resultErr
}

// Description:
//
// # Encapsulate the request and invoke the network
//
// @param action - api name
//
// @param version - product version
//
// @param protocol - http or https
//
// @param method - e.g. GET
//
// @param authType - authorization type e.g. AK
//
// @param pathname - pathname of every api
//
// @param bodyType - response body type e.g. String
//
// @param request - object of OpenApiRequest
//
// @param runtime - which controls some details of call api, such as retry times
//
// @return the response
func (client *Client) DoROARequestWithCtx(ctx context.Context, action *string, version *string, protocol *string, method *string, authType *string, pathname *string, bodyType *string, request *openapiutil.OpenApiRequest, runtime *dara.RuntimeOptions) (_result map[string]interface{}, _err error) {
	_runtime := dara.NewRuntimeObject(map[string]interface{}{
		"key":            dara.ToString(dara.Default(dara.StringValue(runtime.Key), dara.StringValue(client.Key))),
		"cert":           dara.ToString(dara.Default(dara.StringValue(runtime.Cert), dara.StringValue(client.Cert))),
		"ca":             dara.ToString(dara.Default(dara.StringValue(runtime.Ca), dara.StringValue(client.Ca))),
		"readTimeout":    dara.ForceInt(dara.Default(dara.IntValue(runtime.ReadTimeout), dara.IntValue(client.ReadTimeout))),
		"connectTimeout": dara.ForceInt(dara.Default(dara.IntValue(runtime.ConnectTimeout), dara.IntValue(client.ConnectTimeout))),
		"httpProxy":      dara.ToString(dara.Default(dara.StringValue(runtime.HttpProxy), dara.StringValue(client.HttpProxy))),
		"httpsProxy":     dara.ToString(dara.Default(dara.StringValue(runtime.HttpsProxy), dara.StringValue(client.HttpsProxy))),
		"noProxy":        dara.ToString(dara.Default(dara.StringValue(runtime.NoProxy), dara.StringValue(client.NoProxy))),
		"socks5Proxy":    dara.ToString(dara.Default(dara.StringValue(runtime.Socks5Proxy), dara.StringValue(client.Socks5Proxy))),
		"socks5NetWork":  dara.ToString(dara.Default(dara.StringValue(runtime.Socks5NetWork), dara.StringValue(client.Socks5NetWork))),
		"maxIdleConns":   dara.ForceInt(dara.Default(dara.IntValue(runtime.MaxIdleConns), dara.IntValue(client.MaxIdleConns))),
		"retryOptions":   client.RetryOptions,
		"ignoreSSL":      dara.BoolValue(runtime.IgnoreSSL),
		"httpClient":     client.HttpClient,
		"tlsMinVersion":  dara.StringValue(client.TlsMinVersion),
	})

	var retryPolicyContext *dara.RetryPolicyContext
	var request_ *dara.Request
	var response_ *dara.Response
	var _resultErr error
	retriesAttempted := int(0)
	retryPolicyContext = &dara.RetryPolicyContext{
		RetriesAttempted: retriesAttempted,
	}

	_result = make(map[string]interface{})
	for dara.ShouldRetry(_runtime.RetryOptions, retryPolicyContext) {
		_resultErr = nil
		_backoffDelayTime := dara.GetBackoffDelay(_runtime.RetryOptions, retryPolicyContext)
		dara.Sleep(_backoffDelayTime)

		request_ = dara.NewRequest()
		request_.Protocol = dara.String(dara.ToString(dara.Default(dara.StringValue(client.Protocol), dara.StringValue(protocol))))
		request_.Method = method
		request_.Pathname = pathname
		globalQueries := make(map[string]*string)
		globalHeaders := make(map[string]*string)
		if !dara.IsNil(client.GlobalParameters) {
			globalParams := client.GlobalParameters
			if !dara.IsNil(globalParams.Queries) {
				globalQueries = globalParams.Queries
			}

			if !dara.IsNil(globalParams.Headers) {
				globalHeaders = globalParams.Headers
			}

		}

		extendsHeaders := make(map[string]*string)
		extendsQueries := make(map[string]*string)
		if !dara.IsNil(runtime.ExtendsParameters) {
			extendsParameters := runtime.ExtendsParameters
			if !dara.IsNil(extendsParameters.Headers) {
				extendsHeaders = extendsParameters.Headers
			}

			if !dara.IsNil(extendsParameters.Queries) {
				extendsQueries = extendsParameters.Queries
			}

		}

		request_.Headers = dara.Merge(map[string]*string{
			"date":                    openapiutil.GetDateUTCString(),
			"host":                    client.Endpoint,
			"accept":                  dara.String("application/json"),
			"x-acs-signature-nonce":   openapiutil.GetNonce(),
			"x-acs-signature-method":  dara.String("HMAC-SHA1"),
			"x-acs-signature-version": dara.String("1.0"),
			"x-acs-version":           version,
			"x-acs-action":            action,
			"user-agent":              openapiutil.GetUserAgent(client.UserAgent),
		}, globalHeaders,
			extendsHeaders,
			request.Headers)
		if !dara.IsNil(request.Body) {
			request_.Body = dara.ToReader(dara.Stringify(request.Body))
			request_.Headers["content-type"] = dara.String("application/json; charset=utf-8")
		}

		request_.Query = dara.Merge(globalQueries,
			extendsQueries)
		if !dara.IsNil(request.Query) {
			request_.Query = dara.Merge(request_.Query,
				request.Query)
		}

		if dara.StringValue(authType) != "Anonymous" {
			if dara.IsNil(client.Credential) {
				_err = &ClientError{
					Code:    dara.String("InvalidCredentials"),
					Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
				}
				if dara.BoolValue(client.DisableSDKError) != true {
					_err = dara.TeaSDKError(_err)
				}
				return _result, _err
			}

			credentialModel, _err := client.Credential.GetCredential()
			if _err != nil {
				retriesAttempted++
				retryPolicyContext = &dara.RetryPolicyContext{
					RetriesAttempted: retriesAttempted,
					HttpRequest:      request_,
					HttpResponse:     response_,
					Exception:        _err,
				}
				_resultErr = _err
				continue
			}

			if !dara.IsNil(credentialModel.ProviderName) {
				request_.Headers["x-acs-credentials-provider"] = credentialModel.ProviderName
			}

			credentialType := dara.StringValue(credentialModel.Type)
			if credentialType == "bearer" {
				bearerToken := dara.StringValue(credentialModel.BearerToken)
				request_.Headers["x-acs-bearer-token"] = dara.String(bearerToken)
				request_.Headers["x-acs-signature-type"] = dara.String("BEARERTOKEN")
			} else if credentialType == "id_token" {
				idToken := dara.StringValue(credentialModel.SecurityToken)
				request_.Headers["x-acs-zero-trust-idtoken"] = dara.String(idToken)
			} else {
				accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
				accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
				securityToken := dara.StringValue(credentialModel.SecurityToken)
				if !dara.IsNil(dara.String(securityToken)) && securityToken != "" {
					request_.Headers["x-acs-accesskey-id"] = dara.String(accessKeyId)
					request_.Headers["x-acs-security-token"] = dara.String(securityToken)
				}

				stringToSign := dara.StringValue(openapiutil.GetStringToSign(request_))
				request_.Headers["authorization"] = dara.String("acs " + accessKeyId + ":" + dara.StringValue(openapiutil.GetROASignature(dara.String(stringToSign), dara.String(accessKeySecret))))
			}

		}

		response_, _err := dara.DoRequestWithCtx(ctx, request_, _runtime)
		if _err != nil {
			retriesAttempted++
			retryPolicyContext = &dara.RetryPolicyContext{
				RetriesAttempted: retriesAttempted,
				HttpRequest:      request_,
				HttpResponse:     response_,
				Exception:        _err,
			}
			_resultErr = _err
			continue
		}

		_result, _err = doROARequestWithCtx_opResponse(response_, client, bodyType)
		if _err != nil {
			retriesAttempted++
			retryPolicyContext = &dara.RetryPolicyContext{
				RetriesAttempted: retriesAttempted,
				HttpRequest:      request_,
				HttpResponse:     response_,
				Exception:        _err,
			}
			_resultErr = _err
			continue
		}

		return _result, _err
	}
	if dara.BoolValue(client.DisableSDKError) != true {
		_resultErr = dara.TeaSDKError(_resultErr)
	}
	return _result, _resultErr
}

// Description:
//
// # Encapsulate the request and invoke the network with form body
//
// @param action - api name
//
// @param version - product version
//
// @param protocol - http or https
//
// @param method - e.g. GET
//
// @param authType - authorization type e.g. AK
//
// @param pathname - pathname of every api
//
// @param bodyType - response body type e.g. String
//
// @param request - object of OpenApiRequest
//
// @param runtime - which controls some details of call api, such as retry times
//
// @return the response
func (client *Client) DoROAFormRequestWithCtx(ctx context.Context, action *string, version *string, protocol *string, method *string, authType *string, pathname *string, bodyType *string, request *openapiutil.OpenApiRequest, runtime *dara.RuntimeOptions) (_result map[string]interface{}, _err error) {
	_runtime := dara.NewRuntimeObject(map[string]interface{}{
		"key":            dara.ToString(dara.Default(dara.StringValue(runtime.Key), dara.StringValue(client.Key))),
		"cert":           dara.ToString(dara.Default(dara.StringValue(runtime.Cert), dara.StringValue(client.Cert))),
		"ca":             dara.ToString(dara.Default(dara.StringValue(runtime.Ca), dara.StringValue(client.Ca))),
		"readTimeout":    dara.ForceInt(dara.Default(dara.IntValue(runtime.ReadTimeout), dara.IntValue(client.ReadTimeout))),
		"connectTimeout": dara.ForceInt(dara.Default(dara.IntValue(runtime.ConnectTimeout), dara.IntValue(client.ConnectTimeout))),
		"httpProxy":      dara.ToString(dara.Default(dara.StringValue(runtime.HttpProxy), dara.StringValue(client.HttpProxy))),
		"httpsProxy":     dara.ToString(dara.Default(dara.StringValue(runtime.HttpsProxy), dara.StringValue(client.HttpsProxy))),
		"noProxy":        dara.ToString(dara.Default(dara.StringValue(runtime.NoProxy), dara.StringValue(client.NoProxy))),
		"socks5Proxy":    dara.ToString(dara.Default(dara.StringValue(runtime.Socks5Proxy), dara.StringValue(client.Socks5Proxy))),
		"socks5NetWork":  dara.ToString(dara.Default(dara.StringValue(runtime.Socks5NetWork), dara.StringValue(client.Socks5NetWork))),
		"maxIdleConns":   dara.ForceInt(dara.Default(dara.IntValue(runtime.MaxIdleConns), dara.IntValue(client.MaxIdleConns))),
		"retryOptions":   client.RetryOptions,
		"ignoreSSL":      dara.BoolValue(runtime.IgnoreSSL),
		"httpClient":     client.HttpClient,
		"tlsMinVersion":  dara.StringValue(client.TlsMinVersion),
	})

	var retryPolicyContext *dara.RetryPolicyContext
	var request_ *dara.Request
	var response_ *dara.Response
	var _resultErr error
	retriesAttempted := int(0)
	retryPolicyContext = &dara.RetryPolicyContext{
		RetriesAttempted: retriesAttempted,
	}

	_result = make(map[string]interface{})
	for dara.ShouldRetry(_runtime.RetryOptions, retryPolicyContext) {
		_resultErr = nil
		_backoffDelayTime := dara.GetBackoffDelay(_runtime.RetryOptions, retryPolicyContext)
		dara.Sleep(_backoffDelayTime)

		request_ = dara.NewRequest()
		request_.Protocol = dara.String(dara.ToString(dara.Default(dara.StringValue(client.Protocol), dara.StringValue(protocol))))
		request_.Method = method
		request_.Pathname = pathname
		globalQueries := make(map[string]*string)
		globalHeaders := make(map[string]*string)
		if !dara.IsNil(client.GlobalParameters) {
			globalParams := client.GlobalParameters
			if !dara.IsNil(globalParams.Queries) {
				globalQueries = globalParams.Queries
			}

			if !dara.IsNil(globalParams.Headers) {
				globalHeaders = globalParams.Headers
			}

		}

		extendsHeaders := make(map[string]*string)
		extendsQueries := make(map[string]*string)
		if !dara.IsNil(runtime.ExtendsParameters) {
			extendsParameters := runtime.ExtendsParameters
			if !dara.IsNil(extendsParameters.Headers) {
				extendsHeaders = extendsParameters.Headers
			}

			if !dara.IsNil(extendsParameters.Queries) {
				extendsQueries = extendsParameters.Queries
			}

		}

		request_.Headers = dara.Merge(map[string]*string{
			"date":                    openapiutil.GetDateUTCString(),
			"host":                    client.Endpoint,
			"accept":                  dara.String("application/json"),
			"x-acs-signature-nonce":   openapiutil.GetNonce(),
			"x-acs-signature-method":  dara.String("HMAC-SHA1"),
			"x-acs-signature-version": dara.String("1.0"),
			"x-acs-version":           version,
			"x-acs-action":            action,
			"user-agent":              openapiutil.GetUserAgent(client.UserAgent),
		}, globalHeaders,
			extendsHeaders,
			request.Headers)
		if !dara.IsNil(request.Body) {
			m := dara.ToMap(request.Body)
			request_.Body = dara.ToReader(openapiutil.ToForm(m))
			request_.Headers["content-type"] = dara.String("application/x-www-form-urlencoded")
		}

		request_.Query = dara.Merge(globalQueries,
			extendsQueries)
		if !dara.IsNil(request.Query) {
			request_.Query = dara.Merge(request_.Query,
				request.Query)
		}

		if dara.StringValue(authType) != "Anonymous" {
			if dara.IsNil(client.Credential) {
				_err = &ClientError{
					Code:    dara.String("InvalidCredentials"),
					Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
				}
				if dara.BoolValue(client.DisableSDKError) != true {
					_err = dara.TeaSDKError(_err)
				}
				return _result, _err
			}

			credentialModel, _err := client.Credential.GetCredential()
			if _err != nil {
				retriesAttempted++
				retryPolicyContext = &dara.RetryPolicyContext{
					RetriesAttempted: retriesAttempted,
					HttpRequest:      request_,
					HttpResponse:     response_,
					Exception:        _err,
				}
				_resultErr = _err
				continue
			}

			if !dara.IsNil(credentialModel.ProviderName) {
				request_.Headers["x-acs-credentials-provider"] = credentialModel.ProviderName
			}

			credentialType := dara.StringValue(credentialModel.Type)
			if credentialType == "bearer" {
				bearerToken := dara.StringValue(credentialModel.BearerToken)
				request_.Headers["x-acs-bearer-token"] = dara.String(bearerToken)
				request_.Headers["x-acs-signature-type"] = dara.String("BEARERTOKEN")
			} else if credentialType == "id_token" {
				idToken := dara.StringValue(credentialModel.SecurityToken)
				request_.Headers["x-acs-zero-trust-idtoken"] = dara.String(idToken)
			} else {
				accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
				accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
				securityToken := dara.StringValue(credentialModel.SecurityToken)
				if !dara.IsNil(dara.String(securityToken)) && securityToken != "" {
					request_.Headers["x-acs-accesskey-id"] = dara.String(accessKeyId)
					request_.Headers["x-acs-security-token"] = dara.String(securityToken)
				}

				stringToSign := dara.StringValue(openapiutil.GetStringToSign(request_))
				request_.Headers["authorization"] = dara.String("acs " + accessKeyId + ":" + dara.StringValue(openapiutil.GetROASignature(dara.String(stringToSign), dara.String(accessKeySecret))))
			}

		}

		response_, _err := dara.DoRequestWithCtx(ctx, request_, _runtime)
		if _err != nil {
			retriesAttempted++
			retryPolicyContext = &dara.RetryPolicyContext{
				RetriesAttempted: retriesAttempted,
				HttpRequest:      request_,
				HttpResponse:     response_,
				Exception:        _err,
			}
			_resultErr = _err
			continue
		}

		_result, _err = doROAFormRequestWithCtx_opResponse(response_, client, bodyType)
		if _err != nil {
			retriesAttempted++
			retryPolicyContext = &dara.RetryPolicyContext{
				RetriesAttempted: retriesAttempted,
				HttpRequest:      request_,
				HttpResponse:     response_,
				Exception:        _err,
			}
			_resultErr = _err
			continue
		}

		return _result, _err
	}
	if dara.BoolValue(client.DisableSDKError) != true {
		_resultErr = dara.TeaSDKError(_resultErr)
	}
	return _result, _resultErr
}

// Description:
//
// # Encapsulate the request and invoke the network
//
// @param action - api name
//
// @param version - product version
//
// @param protocol - http or https
//
// @param method - e.g. GET
//
// @param authType - authorization type e.g. AK
//
// @param bodyType - response body type e.g. String
//
// @param request - object of OpenApiRequest
//
// @param runtime - which controls some details of call api, such as retry times
//
// @return the response
func (client *Client) DoRequestWithCtx(ctx context.Context, params *openapiutil.Params, request *openapiutil.OpenApiRequest, runtime *dara.RuntimeOptions) (_result map[string]interface{}, _err error) {
	_runtime := dara.NewRuntimeObject(map[string]interface{}{
		"key":            dara.ToString(dara.Default(dara.StringValue(runtime.Key), dara.StringValue(client.Key))),
		"cert":           dara.ToString(dara.Default(dara.StringValue(runtime.Cert), dara.StringValue(client.Cert))),
		"ca":             dara.ToString(dara.Default(dara.StringValue(runtime.Ca), dara.StringValue(client.Ca))),
		"readTimeout":    dara.ForceInt(dara.Default(dara.IntValue(runtime.ReadTimeout), dara.IntValue(client.ReadTimeout))),
		"connectTimeout": dara.ForceInt(dara.Default(dara.IntValue(runtime.ConnectTimeout), dara.IntValue(client.ConnectTimeout))),
		"httpProxy":      dara.ToString(dara.Default(dara.StringValue(runtime.HttpProxy), dara.StringValue(client.HttpProxy))),
		"httpsProxy":     dara.ToString(dara.Default(dara.StringValue(runtime.HttpsProxy), dara.StringValue(client.HttpsProxy))),
		"noProxy":        dara.ToString(dara.Default(dara.StringValue(runtime.NoProxy), dara.StringValue(client.NoProxy))),
		"socks5Proxy":    dara.ToString(dara.Default(dara.StringValue(runtime.Socks5Proxy), dara.StringValue(client.Socks5Proxy))),
		"socks5NetWork":  dara.ToString(dara.Default(dara.StringValue(runtime.Socks5NetWork), dara.StringValue(client.Socks5NetWork))),
		"maxIdleConns":   dara.ForceInt(dara.Default(dara.IntValue(runtime.MaxIdleConns), dara.IntValue(client.MaxIdleConns))),
		"retryOptions":   client.RetryOptions,
		"ignoreSSL":      dara.BoolValue(runtime.IgnoreSSL),
		"httpClient":     client.HttpClient,
		"tlsMinVersion":  dara.StringValue(client.TlsMinVersion),
	})

	var retryPolicyContext *dara.RetryPolicyContext
	var request_ *dara.Request
	var response_ *dara.Response
	var _resultErr error
	retriesAttempted := int(0)
	retryPolicyContext = &dara.RetryPolicyContext{
		RetriesAttempted: retriesAttempted,
	}

	_result = make(map[string]interface{})
	for dara.ShouldRetry(_runtime.RetryOptions, retryPolicyContext) {
		_resultErr = nil
		_backoffDelayTime := dara.GetBackoffDelay(_runtime.RetryOptions, retryPolicyContext)
		dara.Sleep(_backoffDelayTime)

		request_ = dara.NewRequest()
		request_.Protocol = dara.String(dara.ToString(dara.Default(dara.StringValue(client.Protocol), dara.StringValue(params.Protocol))))
		request_.Method = params.Method
		request_.Pathname = params.Pathname
		globalQueries := make(map[string]*string)
		globalHeaders := make(map[string]*string)
		if !dara.IsNil(client.GlobalParameters) {
			globalParams := client.GlobalParameters
			if !dara.IsNil(globalParams.Queries) {
				globalQueries = globalParams.Queries
			}

			if !dara.IsNil(globalParams.Headers) {
				globalHeaders = globalParams.Headers
			}

		}

		extendsHeaders := make(map[string]*string)
		extendsQueries := make(map[string]*string)
		if !dara.IsNil(runtime.ExtendsParameters) {
			extendsParameters := runtime.ExtendsParameters
			if !dara.IsNil(extendsParameters.Headers) {
				extendsHeaders = extendsParameters.Headers
			}

			if !dara.IsNil(extendsParameters.Queries) {
				extendsQueries = extendsParameters.Queries
			}

		}

		request_.Query = dara.Merge(globalQueries,
			extendsQueries,
			request.Query)
		// endpoint is setted in product client
		request_.Headers = dara.Merge(map[string]*string{
			"host":                  client.Endpoint,
			"x-acs-version":         params.Version,
			"x-acs-action":          params.Action,
			"user-agent":            openapiutil.GetUserAgent(client.UserAgent),
			"x-acs-date":            openapiutil.GetTimestamp(),
			"x-acs-signature-nonce": openapiutil.GetNonce(),
			"accept":                dara.String("application/json"),
		}, globalHeaders,
			extendsHeaders,
			request.Headers)
		if dara.StringValue(params.Style) == "RPC" {
			headers, _err := client.GetRpcHeaders()
			if _err != nil {
				retriesAttempted++
				retryPolicyContext = &dara.RetryPolicyContext{
					RetriesAttempted: retriesAttempted,
					HttpRequest:      request_,
					HttpResponse:     response_,
					Exception:        _err,
				}
				_resultErr = _err
				continue
			}

			if !dara.IsNil(headers) {
				request_.Headers = dara.Merge(request_.Headers,
					headers)
			}

		}

		signatureAlgorithm := dara.ToString(dara.Default(dara.StringValue(client.SignatureAlgorithm), "ACS3-HMAC-SHA256"))
		hashedRequestPayload := openapiutil.Hash(dara.BytesFromString("", "utf-8"), dara.String(signatureAlgorithm))
		if !dara.IsNil(request.Stream) {
			tmp, _err := dara.ReadAsBytes(request.Stream)
			if _err != nil {
				retriesAttempted++
				retryPolicyContext = &dara.RetryPolicyContext{
					RetriesAttempted: retriesAttempted,
					HttpRequest:      request_,
					HttpResponse:     response_,
					Exception:        _err,
				}
				_resultErr = _err
				continue
			}

			hashedRequestPayload = openapiutil.Hash(tmp, dara.String(signatureAlgorithm))
			request_.Body = dara.ToReader(tmp)
			request_.Headers["content-type"] = dara.String("application/octet-stream")
		} else {
			if !dara.IsNil(request.Body) {
				if dara.StringValue(params.ReqBodyType) == "byte" {
					byteObj := []byte(dara.ToString(request.Body))
					hashedRequestPayload = openapiutil.Hash(byteObj, dara.String(signatureAlgorithm))
					request_.Body = dara.ToReader(byteObj)
				} else if dara.StringValue(params.ReqBodyType) == "json" {
					jsonObj := dara.Stringify(request.Body)
					hashedRequestPayload = openapiutil.Hash(dara.ToBytes(jsonObj, "utf8"), dara.String(signatureAlgorithm))
					request_.Body = dara.ToReader(dara.String(jsonObj))
					request_.Headers["content-type"] = dara.String("application/json; charset=utf-8")
				} else {
					m := dara.ToMap(request.Body)
					formObj := dara.StringValue(openapiutil.ToForm(m))
					hashedRequestPayload = openapiutil.Hash(dara.ToBytes(formObj, "utf8"), dara.String(signatureAlgorithm))
					request_.Body = dara.ToReader(dara.String(formObj))
					request_.Headers["content-type"] = dara.String("application/x-www-form-urlencoded")
				}

			}

		}

		request_.Headers["x-acs-content-sha256"] = dara.String(hex.EncodeToString(hashedRequestPayload))
		if dara.StringValue(params.AuthType) != "Anonymous" {
			if dara.IsNil(client.Credential) {
				_err = &ClientError{
					Code:    dara.String("InvalidCredentials"),
					Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
				}
				if dara.BoolValue(client.DisableSDKError) != true {
					_err = dara.TeaSDKError(_err)
				}
				return _result, _err
			}

			credentialModel, _err := client.Credential.GetCredential()
			if _err != nil {
				retriesAttempted++
				retryPolicyContext = &dara.RetryPolicyContext{
					RetriesAttempted: retriesAttempted,
					HttpRequest:      request_,
					HttpResponse:     response_,
					Exception:        _err,
				}
				_resultErr = _err
				continue
			}

			if !dara.IsNil(credentialModel.ProviderName) {
				request_.Headers["x-acs-credentials-provider"] = credentialModel.ProviderName
			}

			authType := dara.StringValue(credentialModel.Type)
			if authType == "bearer" {
				bearerToken := dara.StringValue(credentialModel.BearerToken)
				request_.Headers["x-acs-bearer-token"] = dara.String(bearerToken)
				if dara.StringValue(params.Style) == "RPC" {
					request_.Query["SignatureType"] = dara.String("BEARERTOKEN")
				} else {
					request_.Headers["x-acs-signature-type"] = dara.String("BEARERTOKEN")
				}

			} else if authType == "id_token" {
				idToken := dara.StringValue(credentialModel.SecurityToken)
				request_.Headers["x-acs-zero-trust-idtoken"] = dara.String(idToken)
			} else {
				accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
				accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
				securityToken := dara.StringValue(credentialModel.SecurityToken)
				if !dara.IsNil(dara.String(securityToken)) && securityToken != "" {
					request_.Headers["x-acs-accesskey-id"] = dara.String(accessKeyId)
					request_.Headers["x-acs-security-token"] = dara.String(securityToken)
				}

				request_.Headers["Authorization"] = openapiutil.GetAuthorization(request_, dara.String(signatureAlgorithm), dara.String(hex.EncodeToString(hashedRequestPayload)), dara.String(accessKeyId), dara.String(accessKeySecret))
			}

		}

		response_, _err := dara.DoRequestWithCtx(ctx, request_, _runtime)
		if _err != nil {
			retriesAttempted++
			retryPolicyContext = &dara.RetryPolicyContext{
				RetriesAttempted: retriesAttempted,
				HttpRequest:      request_,
				HttpResponse:     response_,
				Exception:        _err,
			}
			_resultErr = _err
			continue
		}

		_result, _err = doRequestWithCtx_opResponse(response_, client, params)
		if _err != nil {
			retriesAttempted++
			retryPolicyContext = &dara.RetryPolicyContext{
				RetriesAttempted: retriesAttempted,
				HttpRequest:      request_,
				HttpResponse:     response_,
				Exception:        _err,
			}
			_resultErr = _err
			continue
		}

		return _result, _err
	}
	if dara.BoolValue(client.DisableSDKError) != true {
		_resultErr = dara.TeaSDKError(_resultErr)
	}
	return _result, _resultErr
}

// Description:
//
// # Encapsulate the request and invoke the network
//
// @param action - api name
//
// @param version - product version
//
// @param protocol - http or https
//
// @param method - e.g. GET
//
// @param authType - authorization type e.g. AK
//
// @param bodyType - response body type e.g. String
//
// @param request - object of OpenApiRequest
//
// @param runtime - which controls some details of call api, such as retry times
//
// @return the response
func (client *Client) ExecuteWithCtx(ctx context.Context, params *openapiutil.Params, request *openapiutil.OpenApiRequest, runtime *dara.RuntimeOptions) (_result map[string]interface{}, _err error) {
	_runtime := dara.NewRuntimeObject(map[string]interface{}{
		"key":            dara.ToString(dara.Default(dara.StringValue(runtime.Key), dara.StringValue(client.Key))),
		"cert":           dara.ToString(dara.Default(dara.StringValue(runtime.Cert), dara.StringValue(client.Cert))),
		"ca":             dara.ToString(dara.Default(dara.StringValue(runtime.Ca), dara.StringValue(client.Ca))),
		"readTimeout":    dara.ForceInt(dara.Default(dara.IntValue(runtime.ReadTimeout), dara.IntValue(client.ReadTimeout))),
		"connectTimeout": dara.ForceInt(dara.Default(dara.IntValue(runtime.ConnectTimeout), dara.IntValue(client.ConnectTimeout))),
		"httpProxy":      dara.ToString(dara.Default(dara.StringValue(runtime.HttpProxy), dara.StringValue(client.HttpProxy))),
		"httpsProxy":     dara.ToString(dara.Default(dara.StringValue(runtime.HttpsProxy), dara.StringValue(client.HttpsProxy))),
		"noProxy":        dara.ToString(dara.Default(dara.StringValue(runtime.NoProxy), dara.StringValue(client.NoProxy))),
		"socks5Proxy":    dara.ToString(dara.Default(dara.StringValue(runtime.Socks5Proxy), dara.StringValue(client.Socks5Proxy))),
		"socks5NetWork":  dara.ToString(dara.Default(dara.StringValue(runtime.Socks5NetWork), dara.StringValue(client.Socks5NetWork))),
		"maxIdleConns":   dara.ForceInt(dara.Default(dara.IntValue(runtime.MaxIdleConns), dara.IntValue(client.MaxIdleConns))),
		"retryOptions":   client.RetryOptions,
		"ignoreSSL":      dara.BoolValue(runtime.IgnoreSSL),
		"httpClient":     client.HttpClient,
		"tlsMinVersion":  dara.StringValue(client.TlsMinVersion),
		"disableHttp2":   dara.ForceBoolean(dara.Default(dara.BoolValue(client.DisableHttp2), false)),
	})

	var retryPolicyContext *dara.RetryPolicyContext
	var request_ *dara.Request
	var response_ *dara.Response
	var _resultErr error
	retriesAttempted := int(0)
	retryPolicyContext = &dara.RetryPolicyContext{
		RetriesAttempted: retriesAttempted,
	}

	_result = make(map[string]interface{})
	for dara.ShouldRetry(_runtime.RetryOptions, retryPolicyContext) {
		_resultErr = nil
		_backoffDelayTime := dara.GetBackoffDelay(_runtime.RetryOptions, retryPolicyContext)
		dara.Sleep(_backoffDelayTime)

		request_ = dara.NewRequest()
		// spi = new Gateway();//Gateway implements SPI，这一步在产品 SDK 中实例化
		headers, _err := client.GetRpcHeaders()
		if _err != nil {
			retriesAttempted++
			retryPolicyContext = &dara.RetryPolicyContext{
				RetriesAttempted: retriesAttempted,
				HttpRequest:      request_,
				HttpResponse:     response_,
				Exception:        _err,
			}
			_resultErr = _err
			continue
		}

		globalQueries := make(map[string]*string)
		globalHeaders := make(map[string]*string)
		if !dara.IsNil(client.GlobalParameters) {
			globalParams := client.GlobalParameters
			if !dara.IsNil(globalParams.Queries) {
				globalQueries = globalParams.Queries
			}

			if !dara.IsNil(globalParams.Headers) {
				globalHeaders = globalParams.Headers
			}

		}

		extendsHeaders := make(map[string]*string)
		extendsQueries := make(map[string]*string)
		if !dara.IsNil(runtime.ExtendsParameters) {
			extendsParameters := runtime.ExtendsParameters
			if !dara.IsNil(extendsParameters.Headers) {
				extendsHeaders = extendsParameters.Headers
			}

			if !dara.IsNil(extendsParameters.Queries) {
				extendsQueries = extendsParameters.Queries
			}

		}

		requestContext := &spi.InterceptorContextRequest{
			Headers: dara.Merge(globalHeaders,
				extendsHeaders,
				request.Headers,
				headers),
			Query: dara.Merge(globalQueries,
				extendsQueries,
				request.Query),
			Body:               request.Body,
			Stream:             request.Stream,
			HostMap:            request.HostMap,
			Pathname:           params.Pathname,
			ProductId:          client.ProductId,
			Action:             params.Action,
			Version:            params.Version,
			Protocol:           dara.String(dara.ToString(dara.Default(dara.StringValue(client.Protocol), dara.StringValue(params.Protocol)))),
			Method:             dara.String(dara.ToString(dara.Default(dara.StringValue(client.Method), dara.StringValue(params.Method)))),
			AuthType:           params.AuthType,
			BodyType:           params.BodyType,
			ReqBodyType:        params.ReqBodyType,
			Style:              params.Style,
			Credential:         client.Credential,
			SignatureVersion:   client.SignatureVersion,
			SignatureAlgorithm: client.SignatureAlgorithm,
			UserAgent:          openapiutil.GetUserAgent(client.UserAgent),
		}
		configurationContext := &spi.InterceptorContextConfiguration{
			RegionId:     client.RegionId,
			Endpoint:     dara.String(dara.ToString(dara.Default(dara.StringValue(request.EndpointOverride), dara.StringValue(client.Endpoint)))),
			EndpointRule: client.EndpointRule,
			EndpointMap:  client.EndpointMap,
			EndpointType: client.EndpointType,
			Network:      client.Network,
			Suffix:       client.Suffix,
		}
		interceptorContext := &spi.InterceptorContext{
			Request:       requestContext,
			Configuration: configurationContext,
		}
		attributeMap := &spi.AttributeMap{}
		if !dara.IsNil(client.AttributeMap) {
			attributeMap = client.AttributeMap
		}

		// 1. spi.modifyConfiguration(context: SPI.InterceptorContext, attributeMap: SPI.AttributeMap);
		_err = client.Spi.ModifyConfiguration(interceptorContext, attributeMap)
		if _err != nil {
			retriesAttempted++
			retryPolicyContext = &dara.RetryPolicyContext{
				RetriesAttempted: retriesAttempted,
				HttpRequest:      request_,
				HttpResponse:     response_,
				Exception:        _err,
			}
			_resultErr = _err
			continue
		}

		// 2. spi.modifyRequest(context: SPI.InterceptorContext, attributeMap: SPI.AttributeMap);
		_err = client.Spi.ModifyRequest(interceptorContext, attributeMap)
		if _err != nil {
			retriesAttempted++
			retryPolicyContext = &dara.RetryPolicyContext{
				RetriesAttempted: retriesAttempted,
				HttpRequest:      request_,
				HttpResponse:     response_,
				Exception:        _err,
			}
			_resultErr = _err
			continue
		}

		request_.Protocol = interceptorContext.Request.Protocol
		request_.Method = interceptorContext.Request.Method
		request_.Pathname = interceptorContext.Request.Pathname
		request_.Query = interceptorContext.Request.Query
		request_.Body = interceptorContext.Request.Stream
		request_.Headers = interceptorContext.Request.Headers
		response_, _err := dara.DoRequestWithCtx(ctx, request_, _runtime)
		if _err != nil {
			retriesAttempted++
			retryPolicyContext = &dara.RetryPolicyContext{
				RetriesAttempted: retriesAttempted,
				HttpRequest:      request_,
				HttpResponse:     response_,
				Exception:        _err,
			}
			_resultErr = _err
			continue
		}

		responseContext := &spi.InterceptorContextResponse{
			StatusCode: response_.StatusCode,
			Headers:    response_.Headers,
			Body:       response_.Body,
		}
		interceptorContext.Response = responseContext
		// 3. spi.modifyResponse(context: SPI.InterceptorContext, attributeMap: SPI.AttributeMap);
		_err = client.Spi.ModifyResponse(interceptorContext, attributeMap)
		if _err != nil {
			retriesAttempted++
			retryPolicyContext = &dara.RetryPolicyContext{
				RetriesAttempted: retriesAttempted,
				HttpRequest:      request_,
				HttpResponse:     response_,
				Exception:        _err,
			}
			_resultErr = _err
			continue
		}

		_err = dara.Convert(map[string]interface{}{
			"headers":    interceptorContext.Response.Headers,
			"statusCode": dara.IntValue(interceptorContext.Response.StatusCode),
			"body":       interceptorContext.Response.DeserializedBody,
		}, &_result)

		return _result, _err
	}
	if dara.BoolValue(client.DisableSDKError) != true {
		_resultErr = dara.TeaSDKError(_resultErr)
	}
	return _result, _resultErr
}

func (client *Client) CallSSEApiWithCtx(ctx context.Context, params *openapiutil.Params, request *openapiutil.OpenApiRequest, runtime *dara.RuntimeOptions, _yield chan *SSEResponse, _yieldErr chan error) {
	defer close(_yield)
	defer close(_yieldErr)
	_runtime := dara.NewRuntimeObject(map[string]interface{}{
		"key":            dara.ToString(dara.Default(dara.StringValue(runtime.Key), dara.StringValue(client.Key))),
		"cert":           dara.ToString(dara.Default(dara.StringValue(runtime.Cert), dara.StringValue(client.Cert))),
		"ca":             dara.ToString(dara.Default(dara.StringValue(runtime.Ca), dara.StringValue(client.Ca))),
		"readTimeout":    dara.ForceInt(dara.Default(dara.IntValue(runtime.ReadTimeout), dara.IntValue(client.ReadTimeout))),
		"connectTimeout": dara.ForceInt(dara.Default(dara.IntValue(runtime.ConnectTimeout), dara.IntValue(client.ConnectTimeout))),
		"httpProxy":      dara.ToString(dara.Default(dara.StringValue(runtime.HttpProxy), dara.StringValue(client.HttpProxy))),
		"httpsProxy":     dara.ToString(dara.Default(dara.StringValue(runtime.HttpsProxy), dara.StringValue(client.HttpsProxy))),
		"noProxy":        dara.ToString(dara.Default(dara.StringValue(runtime.NoProxy), dara.StringValue(client.NoProxy))),
		"socks5Proxy":    dara.ToString(dara.Default(dara.StringValue(runtime.Socks5Proxy), dara.StringValue(client.Socks5Proxy))),
		"socks5NetWork":  dara.ToString(dara.Default(dara.StringValue(runtime.Socks5NetWork), dara.StringValue(client.Socks5NetWork))),
		"maxIdleConns":   dara.ForceInt(dara.Default(dara.IntValue(runtime.MaxIdleConns), dara.IntValue(client.MaxIdleConns))),
		"retryOptions":   client.RetryOptions,
		"ignoreSSL":      dara.BoolValue(runtime.IgnoreSSL),
		"httpClient":     client.HttpClient,
		"tlsMinVersion":  dara.StringValue(client.TlsMinVersion),
	})

	var retryPolicyContext *dara.RetryPolicyContext
	var request_ *dara.Request
	var response_ *dara.Response
	var _resultErr error
	retriesAttempted := int(0)
	retryPolicyContext = &dara.RetryPolicyContext{
		RetriesAttempted: retriesAttempted,
	}

	for dara.ShouldRetry(_runtime.RetryOptions, retryPolicyContext) {
		_resultErr = nil
		_backoffDelayTime := dara.GetBackoffDelay(_runtime.RetryOptions, retryPolicyContext)
		dara.Sleep(_backoffDelayTime)

		request_ = dara.NewRequest()
		request_.Protocol = dara.String(dara.ToString(dara.Default(dara.StringValue(client.Protocol), dara.StringValue(params.Protocol))))
		request_.Method = params.Method
		request_.Pathname = params.Pathname
		globalQueries := make(map[string]*string)
		globalHeaders := make(map[string]*string)
		if !dara.IsNil(client.GlobalParameters) {
			globalParams := client.GlobalParameters
			if !dara.IsNil(globalParams.Queries) {
				globalQueries = globalParams.Queries
			}

			if !dara.IsNil(globalParams.Headers) {
				globalHeaders = globalParams.Headers
			}

		}

		extendsHeaders := make(map[string]*string)
		extendsQueries := make(map[string]*string)
		if !dara.IsNil(runtime.ExtendsParameters) {
			extendsParameters := runtime.ExtendsParameters
			if !dara.IsNil(extendsParameters.Headers) {
				extendsHeaders = extendsParameters.Headers
			}

			if !dara.IsNil(extendsParameters.Queries) {
				extendsQueries = extendsParameters.Queries
			}

		}

		request_.Query = dara.Merge(globalQueries,
			extendsQueries,
			request.Query)
		// endpoint is setted in product client
		request_.Headers = dara.Merge(map[string]*string{
			"host":                  client.Endpoint,
			"x-acs-version":         params.Version,
			"x-acs-action":          params.Action,
			"user-agent":            openapiutil.GetUserAgent(client.UserAgent),
			"x-acs-date":            openapiutil.GetTimestamp(),
			"x-acs-signature-nonce": openapiutil.GetNonce(),
			"accept":                dara.String("application/json"),
		}, extendsHeaders,
			globalHeaders,
			request.Headers)
		if dara.StringValue(params.Style) == "RPC" {
			headers, _err := client.GetRpcHeaders()
			if _err != nil {
				retriesAttempted++
				retryPolicyContext = &dara.RetryPolicyContext{
					RetriesAttempted: retriesAttempted,
					HttpRequest:      request_,
					HttpResponse:     response_,
					Exception:        _err,
				}
				_resultErr = _err
				continue
			}

			if !dara.IsNil(headers) {
				request_.Headers = dara.Merge(request_.Headers,
					headers)
			}

		}

		signatureAlgorithm := dara.ToString(dara.Default(dara.StringValue(client.SignatureAlgorithm), "ACS3-HMAC-SHA256"))
		hashedRequestPayload := openapiutil.Hash(dara.BytesFromString("", "utf-8"), dara.String(signatureAlgorithm))
		if !dara.IsNil(request.Stream) {
			tmp, _err := dara.ReadAsBytes(request.Stream)
			if _err != nil {
				retriesAttempted++
				retryPolicyContext = &dara.RetryPolicyContext{
					RetriesAttempted: retriesAttempted,
					HttpRequest:      request_,
					HttpResponse:     response_,
					Exception:        _err,
				}
				_resultErr = _err
				continue
			}

			hashedRequestPayload = openapiutil.Hash(tmp, dara.String(signatureAlgorithm))
			request_.Body = dara.ToReader(tmp)
			request_.Headers["content-type"] = dara.String("application/octet-stream")
		} else {
			if !dara.IsNil(request.Body) {
				if dara.StringValue(params.ReqBodyType) == "byte" {
					byteObj := []byte(dara.ToString(request.Body))
					hashedRequestPayload = openapiutil.Hash(byteObj, dara.String(signatureAlgorithm))
					request_.Body = dara.ToReader(byteObj)
				} else if dara.StringValue(params.ReqBodyType) == "json" {
					jsonObj := dara.Stringify(request.Body)
					hashedRequestPayload = openapiutil.Hash(dara.ToBytes(jsonObj, "utf8"), dara.String(signatureAlgorithm))
					request_.Body = dara.ToReader(dara.String(jsonObj))
					request_.Headers["content-type"] = dara.String("application/json; charset=utf-8")
				} else {
					m := dara.ToMap(request.Body)
					formObj := dara.StringValue(openapiutil.ToForm(m))
					hashedRequestPayload = openapiutil.Hash(dara.ToBytes(formObj, "utf8"), dara.String(signatureAlgorithm))
					request_.Body = dara.ToReader(dara.String(formObj))
					request_.Headers["content-type"] = dara.String("application/x-www-form-urlencoded")
				}

			}

		}

		request_.Headers["x-acs-content-sha256"] = dara.String(hex.EncodeToString(hashedRequestPayload))
		if dara.StringValue(params.AuthType) != "Anonymous" {
			credentialModel, _err := client.Credential.GetCredential()
			if _err != nil {
				retriesAttempted++
				retryPolicyContext = &dara.RetryPolicyContext{
					RetriesAttempted: retriesAttempted,
					HttpRequest:      request_,
					HttpResponse:     response_,
					Exception:        _err,
				}
				_resultErr = _err
				continue
			}

			if !dara.IsNil(credentialModel.ProviderName) {
				request_.Headers["x-acs-credentials-provider"] = credentialModel.ProviderName
			}

			authType := dara.StringValue(credentialModel.Type)
			if authType == "bearer" {
				bearerToken := dara.StringValue(credentialModel.BearerToken)
				request_.Headers["x-acs-bearer-token"] = dara.String(bearerToken)
			} else if authType == "id_token" {
				idToken := dara.StringValue(credentialModel.SecurityToken)
				request_.Headers["x-acs-zero-trust-idtoken"] = dara.String(idToken)
			} else {
				accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
				accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
				securityToken := dara.StringValue(credentialModel.SecurityToken)
				if !dara.IsNil(dara.String(securityToken)) && securityToken != "" {
					request_.Headers["x-acs-accesskey-id"] = dara.String(accessKeyId)
					request_.Headers["x-acs-security-token"] = dara.String(securityToken)
				}

				request_.Headers["Authorization"] = openapiutil.GetAuthorization(request_, dara.String(signatureAlgorithm), dara.String(hex.EncodeToString(hashedRequestPayload)), dara.String(accessKeyId), dara.String(accessKeySecret))
			}

		}

		response_, _err := dara.DoRequestWithCtx(ctx, request_, _runtime)
		if _err != nil {
			retriesAttempted++
			retryPolicyContext = &dara.RetryPolicyContext{
				RetriesAttempted: retriesAttempted,
				HttpRequest:      request_,
				HttpResponse:     response_,
				Exception:        _err,
			}
			_resultErr = _err
			continue
		}
		callSSEApiWithCtx_opResponse(ctx, _yield, _yieldErr, response_)
		return
	}
	_yieldErr <- _resultErr
	return
}

func (client *Client) CallApiWithCtx(ctx context.Context, params *openapiutil.Params, request *openapiutil.OpenApiRequest, runtime *dara.RuntimeOptions) (_result map[string]interface{}, _err error) {
	if dara.IsNil(params) {
		_err = &ClientError{
			Code:    dara.String("ParameterMissing"),
			Message: dara.String("'params' can not be unset"),
		}
		return _result, _err
	}

	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		if dara.IsNil(client.SignatureAlgorithm) || dara.StringValue(client.SignatureAlgorithm) != "v2" {
			_body, _err := client.DoRequestWithCtx(ctx, params, request, runtime)
			if _err != nil {
				return _result, _err
			}
			_result = _body
			return _result, _err
		} else if (dara.StringValue(params.Style) == "ROA") && (dara.StringValue(params.ReqBodyType) == "json") {
			_body, _err := client.DoROARequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, request, runtime)
			if _err != nil {
				return _result, _err
			}
			_result = _body
			return _result, _err
		} else if dara.StringValue(params.Style) == "ROA" {
			_body, _err := client.DoROAFormRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, request, runtime)
			if _err != nil {
				return _result, _err
			}
			_result = _body
			return _result, _err
		} else {
			_body, _err := client.DoRPCRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, request, runtime)
			if _err != nil {
				return _result, _err
			}
			_result = _body
			return _result, _err
		}

	} else {
		_body, _err := client.ExecuteWithCtx(ctx, params, request, runtime)
		if _err != nil {
			return _result, _err
		}
		_result = _body
		return _result, _err
	}

}

func doRPCRequestWithCtx_opResponse(response_ *dara.Response, client *Client, bodyType *string) (_result map[string]interface{}, _err error) {
	if (dara.IntValue(response_.StatusCode) >= 400) && (dara.IntValue(response_.StatusCode) < 600) {
		_res, _err := dara.ReadAsJSON(response_.Body)
		if _err != nil {
			return _result, _err
		}

		err := dara.ToMap(_res)
		requestId := dara.Default(err["RequestId"], err["requestId"])
		code := dara.Default(err["Code"], err["code"])
		if (dara.ToString(code) == "Throttling") || (dara.ToString(code) == "Throttling.User") || (dara.ToString(code) == "Throttling.Api") {
			_err = &ThrottlingError{
				StatusCode:  response_.StatusCode,
				Code:        dara.String(dara.ToString(code)),
				Message:     dara.String("code: " + dara.ToString(dara.IntValue(response_.StatusCode)) + ", " + dara.ToString(dara.Default(err["Message"], err["message"])) + " request id: " + dara.ToString(requestId)),
				Description: dara.String(dara.ToString(dara.Default(err["Description"], err["description"]))),
				RetryAfter:  openapiutil.GetThrottlingTimeLeft(response_.Headers),
				Data:        err,
				RequestId:   dara.String(dara.ToString(requestId)),
			}
			return _result, _err
		} else if (dara.IntValue(response_.StatusCode) >= 400) && (dara.IntValue(response_.StatusCode) < 500) {
			_err = &ClientError{
				StatusCode:         response_.StatusCode,
				Code:               dara.String(dara.ToString(code)),
				Message:            dara.String("code: " + dara.ToString(dara.IntValue(response_.StatusCode)) + ", " + dara.ToString(dara.Default(err["Message"], err["message"])) + " request id: " + dara.ToString(requestId)),
				Description:        dara.String(dara.ToString(dara.Default(err["Description"], err["description"]))),
				Data:               err,
				AccessDeniedDetail: client.GetAccessDeniedDetail(err),
				RequestId:          dara.String(dara.ToString(requestId)),
			}
			return _result, _err
		} else {
			_err = &ServerError{
				StatusCode:  response_.StatusCode,
				Code:        dara.String(dara.ToString(code)),
				Message:     dara.String("code: " + dara.ToString(dara.IntValue(response_.StatusCode)) + ", " + dara.ToString(dara.Default(err["Message"], err["message"])) + " request id: " + dara.ToString(requestId)),
				Description: dara.String(dara.ToString(dara.Default(err["Description"], err["description"]))),
				Data:        err,
				RequestId:   dara.String(dara.ToString(requestId)),
			}
			return _result, _err
		}

	}

	if dara.StringValue(bodyType) == "binary" {
		resp := map[string]interface{}{
			"body":       response_.Body,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}
		_result = resp
		return _result, _err
	} else if dara.StringValue(bodyType) == "byte" {
		byt, _err := dara.ReadAsBytes(response_.Body)
		if _err != nil {
			return _result, _err
		}

		_err = dara.Convert(map[string]interface{}{
			"body":       byt,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	} else if dara.StringValue(bodyType) == "string" {
		_str, _err := dara.ReadAsString(response_.Body)
		if _err != nil {
			return _result, _err
		}

		_err = dara.Convert(map[string]interface{}{
			"body":       _str,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	} else if dara.StringValue(bodyType) == "json" {
		obj, _err := dara.ReadAsJSON(response_.Body)
		if _err != nil {
			return _result, _err
		}

		res := dara.ToMap(obj)
		_err = dara.Convert(map[string]interface{}{
			"body":       res,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	} else if dara.StringValue(bodyType) == "array" {
		arr, _err := dara.ReadAsJSON(response_.Body)
		if _err != nil {
			return _result, _err
		}

		_err = dara.Convert(map[string]interface{}{
			"body":       arr,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	} else {
		_err = dara.Convert(map[string]interface{}{
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	}

}

func doROARequestWithCtx_opResponse(response_ *dara.Response, client *Client, bodyType *string) (_result map[string]interface{}, _err error) {
	if dara.IntValue(response_.StatusCode) == 204 {
		_err = dara.Convert(map[string]map[string]*string{
			"headers": response_.Headers,
		}, &_result)

		return _result, _err
	}

	if (dara.IntValue(response_.StatusCode) >= 400) && (dara.IntValue(response_.StatusCode) < 600) {
		_res, _err := dara.ReadAsJSON(response_.Body)
		if _err != nil {
			return _result, _err
		}

		err := dara.ToMap(_res)
		requestId := dara.ToString(dara.Default(err["RequestId"], err["requestId"]))
		requestId = dara.ToString(dara.Default(requestId, err["requestid"]))
		code := dara.ToString(dara.Default(err["Code"], err["code"]))
		if (code == "Throttling") || (code == "Throttling.User") || (code == "Throttling.Api") {
			_err = &ThrottlingError{
				StatusCode:  response_.StatusCode,
				Code:        dara.String(code),
				Message:     dara.String("code: " + dara.ToString(dara.IntValue(response_.StatusCode)) + ", " + dara.ToString(dara.Default(err["Message"], err["message"])) + " request id: " + requestId),
				Description: dara.String(dara.ToString(dara.Default(err["Description"], err["description"]))),
				RetryAfter:  openapiutil.GetThrottlingTimeLeft(response_.Headers),
				Data:        err,
				RequestId:   dara.String(requestId),
			}
			return _result, _err
		} else if (dara.IntValue(response_.StatusCode) >= 400) && (dara.IntValue(response_.StatusCode) < 500) {
			_err = &ClientError{
				StatusCode:         response_.StatusCode,
				Code:               dara.String(code),
				Message:            dara.String("code: " + dara.ToString(dara.IntValue(response_.StatusCode)) + ", " + dara.ToString(dara.Default(err["Message"], err["message"])) + " request id: " + requestId),
				Description:        dara.String(dara.ToString(dara.Default(err["Description"], err["description"]))),
				Data:               err,
				AccessDeniedDetail: client.GetAccessDeniedDetail(err),
				RequestId:          dara.String(requestId),
			}
			return _result, _err
		} else {
			_err = &ServerError{
				StatusCode:  response_.StatusCode,
				Code:        dara.String(code),
				Message:     dara.String("code: " + dara.ToString(dara.IntValue(response_.StatusCode)) + ", " + dara.ToString(dara.Default(err["Message"], err["message"])) + " request id: " + requestId),
				Description: dara.String(dara.ToString(dara.Default(err["Description"], err["description"]))),
				Data:        err,
				RequestId:   dara.String(requestId),
			}
			return _result, _err
		}

	}

	if dara.StringValue(bodyType) == "binary" {
		resp := map[string]interface{}{
			"body":       response_.Body,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}
		_result = resp
		return _result, _err
	} else if dara.StringValue(bodyType) == "byte" {
		byt, _err := dara.ReadAsBytes(response_.Body)
		if _err != nil {
			return _result, _err
		}

		_err = dara.Convert(map[string]interface{}{
			"body":       byt,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	} else if dara.StringValue(bodyType) == "string" {
		_str, _err := dara.ReadAsString(response_.Body)
		if _err != nil {
			return _result, _err
		}

		_err = dara.Convert(map[string]interface{}{
			"body":       _str,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	} else if dara.StringValue(bodyType) == "json" {
		obj, _err := dara.ReadAsJSON(response_.Body)
		if _err != nil {
			return _result, _err
		}

		res := dara.ToMap(obj)
		_err = dara.Convert(map[string]interface{}{
			"body":       res,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	} else if dara.StringValue(bodyType) == "array" {
		arr, _err := dara.ReadAsJSON(response_.Body)
		if _err != nil {
			return _result, _err
		}

		_err = dara.Convert(map[string]interface{}{
			"body":       arr,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	} else {
		_err = dara.Convert(map[string]interface{}{
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	}

}

func doROAFormRequestWithCtx_opResponse(response_ *dara.Response, client *Client, bodyType *string) (_result map[string]interface{}, _err error) {
	if dara.IntValue(response_.StatusCode) == 204 {
		_err = dara.Convert(map[string]map[string]*string{
			"headers": response_.Headers,
		}, &_result)

		return _result, _err
	}

	if (dara.IntValue(response_.StatusCode) >= 400) && (dara.IntValue(response_.StatusCode) < 600) {
		_res, _err := dara.ReadAsJSON(response_.Body)
		if _err != nil {
			return _result, _err
		}

		err := dara.ToMap(_res)
		requestId := dara.ToString(dara.Default(err["RequestId"], err["requestId"]))
		code := dara.ToString(dara.Default(err["Code"], err["code"]))
		if (code == "Throttling") || (code == "Throttling.User") || (code == "Throttling.Api") {
			_err = &ThrottlingError{
				StatusCode:  response_.StatusCode,
				Code:        dara.String(code),
				Message:     dara.String("code: " + dara.ToString(dara.IntValue(response_.StatusCode)) + ", " + dara.ToString(dara.Default(err["Message"], err["message"])) + " request id: " + requestId),
				Description: dara.String(dara.ToString(dara.Default(err["Description"], err["description"]))),
				RetryAfter:  openapiutil.GetThrottlingTimeLeft(response_.Headers),
				Data:        err,
				RequestId:   dara.String(requestId),
			}
			return _result, _err
		} else if (dara.IntValue(response_.StatusCode) >= 400) && (dara.IntValue(response_.StatusCode) < 500) {
			_err = &ClientError{
				StatusCode:         response_.StatusCode,
				Code:               dara.String(code),
				Message:            dara.String("code: " + dara.ToString(dara.IntValue(response_.StatusCode)) + ", " + dara.ToString(dara.Default(err["Message"], err["message"])) + " request id: " + requestId),
				Description:        dara.String(dara.ToString(dara.Default(err["Description"], err["description"]))),
				Data:               err,
				AccessDeniedDetail: client.GetAccessDeniedDetail(err),
				RequestId:          dara.String(requestId),
			}
			return _result, _err
		} else {
			_err = &ServerError{
				StatusCode:  response_.StatusCode,
				Code:        dara.String(code),
				Message:     dara.String("code: " + dara.ToString(dara.IntValue(response_.StatusCode)) + ", " + dara.ToString(dara.Default(err["Message"], err["message"])) + " request id: " + requestId),
				Description: dara.String(dara.ToString(dara.Default(err["Description"], err["description"]))),
				Data:        err,
				RequestId:   dara.String(requestId),
			}
			return _result, _err
		}

	}

	if dara.StringValue(bodyType) == "binary" {
		resp := map[string]interface{}{
			"body":       response_.Body,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}
		_result = resp
		return _result, _err
	} else if dara.StringValue(bodyType) == "byte" {
		byt, _err := dara.ReadAsBytes(response_.Body)
		if _err != nil {
			return _result, _err
		}

		_err = dara.Convert(map[string]interface{}{
			"body":       byt,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	} else if dara.StringValue(bodyType) == "string" {
		_str, _err := dara.ReadAsString(response_.Body)
		if _err != nil {
			return _result, _err
		}

		_err = dara.Convert(map[string]interface{}{
			"body":       _str,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	} else if dara.StringValue(bodyType) == "json" {
		obj, _err := dara.ReadAsJSON(response_.Body)
		if _err != nil {
			return _result, _err
		}

		res := dara.ToMap(obj)
		_err = dara.Convert(map[string]interface{}{
			"body":       res,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	} else if dara.StringValue(bodyType) == "array" {
		arr, _err := dara.ReadAsJSON(response_.Body)
		if _err != nil {
			return _result, _err
		}

		_err = dara.Convert(map[string]interface{}{
			"body":       arr,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	} else {
		_err = dara.Convert(map[string]interface{}{
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	}

}

func doRequestWithCtx_opResponse(response_ *dara.Response, client *Client, params *openapiutil.Params) (_result map[string]interface{}, _err error) {
	if (dara.IntValue(response_.StatusCode) >= 400) && (dara.IntValue(response_.StatusCode) < 600) {
		err := map[string]interface{}{}
		if !dara.IsNil(response_.Headers["content-type"]) && dara.StringValue(response_.Headers["content-type"]) == "text/xml;charset=utf-8" {
			_str, _err := dara.ReadAsString(response_.Body)
			if _err != nil {
				return _result, _err
			}

			respMap := dara.ParseXml(_str, nil)
			err = dara.ToMap(respMap["Error"])
		} else {
			_res, _err := dara.ReadAsJSON(response_.Body)
			if _err != nil {
				return _result, _err
			}

			err = dara.ToMap(_res)
		}

		requestId := dara.ToString(dara.Default(err["RequestId"], err["requestId"]))
		code := dara.ToString(dara.Default(err["Code"], err["code"]))
		if (code == "Throttling") || (code == "Throttling.User") || (code == "Throttling.Api") {
			_err = &ThrottlingError{
				StatusCode:  response_.StatusCode,
				Code:        dara.String(code),
				Message:     dara.String("code: " + dara.ToString(dara.IntValue(response_.StatusCode)) + ", " + dara.ToString(dara.Default(err["Message"], err["message"])) + " request id: " + requestId),
				Description: dara.String(dara.ToString(dara.Default(err["Description"], err["description"]))),
				RetryAfter:  openapiutil.GetThrottlingTimeLeft(response_.Headers),
				Data:        err,
				RequestId:   dara.String(requestId),
			}
			return _result, _err
		} else if (dara.IntValue(response_.StatusCode) >= 400) && (dara.IntValue(response_.StatusCode) < 500) {
			_err = &ClientError{
				StatusCode:         response_.StatusCode,
				Code:               dara.String(code),
				Message:            dara.String("code: " + dara.ToString(dara.IntValue(response_.StatusCode)) + ", " + dara.ToString(dara.Default(err["Message"], err["message"])) + " request id: " + requestId),
				Description:        dara.String(dara.ToString(dara.Default(err["Description"], err["description"]))),
				Data:               err,
				AccessDeniedDetail: client.GetAccessDeniedDetail(err),
				RequestId:          dara.String(requestId),
			}
			return _result, _err
		} else {
			_err = &ServerError{
				StatusCode:  response_.StatusCode,
				Code:        dara.String(code),
				Message:     dara.String("code: " + dara.ToString(dara.IntValue(response_.StatusCode)) + ", " + dara.ToString(dara.Default(err["Message"], err["message"])) + " request id: " + requestId),
				Description: dara.String(dara.ToString(dara.Default(err["Description"], err["description"]))),
				Data:        err,
				RequestId:   dara.String(requestId),
			}
			return _result, _err
		}

	}

	if dara.StringValue(params.BodyType) == "binary" {
		resp := map[string]interface{}{
			"body":       response_.Body,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}
		_result = resp
		return _result, _err
	} else if dara.StringValue(params.BodyType) == "byte" {
		byt, _err := dara.ReadAsBytes(response_.Body)
		if _err != nil {
			return _result, _err
		}

		_err = dara.Convert(map[string]interface{}{
			"body":       byt,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	} else if dara.StringValue(params.BodyType) == "string" {
		respStr, _err := dara.ReadAsString(response_.Body)
		if _err != nil {
			return _result, _err
		}

		_err = dara.Convert(map[string]interface{}{
			"body":       respStr,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	} else if dara.StringValue(params.BodyType) == "json" {
		obj, _err := dara.ReadAsJSON(response_.Body)
		if _err != nil {
			return _result, _err
		}

		res := dara.ToMap(obj)
		_err = dara.Convert(map[string]interface{}{
			"body":       res,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	} else if dara.StringValue(params.BodyType) == "array" {
		arr, _err := dara.ReadAsJSON(response_.Body)
		if _err != nil {
			return _result, _err
		}

		_err = dara.Convert(map[string]interface{}{
			"body":       arr,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	} else {
		anything, _err := dara.ReadAsString(response_.Body)
		if _err != nil {
			return _result, _err
		}

		_err = dara.Convert(map[string]interface{}{
			"body":       anything,
			"headers":    response_.Headers,
			"statusCode": dara.IntValue(response_.StatusCode),
		}, &_result)

		return _result, _err
	}

}

func callSSEApiWithCtx_opResponse(ctx context.Context, _yield chan *SSEResponse, _yieldErr chan error, response_ *dara.Response) {
	if (dara.IntValue(response_.StatusCode) >= 400) && (dara.IntValue(response_.StatusCode) < 600) {
		err := map[string]interface{}{}
		if !dara.IsNil(response_.Headers["content-type"]) && dara.StringValue(response_.Headers["content-type"]) == "text/xml;charset=utf-8" {
			_str, _err := dara.ReadAsString(response_.Body)
			if _err != nil {
				_yieldErr <- _err
				return
			}

			respMap := dara.ParseXml(_str, nil)
			err = dara.ToMap(respMap["Error"])
		} else {
			_res, _err := dara.ReadAsJSON(response_.Body)
			if _err != nil {
				_yieldErr <- _err
				return
			}

			err = dara.ToMap(_res)
		}

		err["statusCode"] = response_.StatusCode
		_err := dara.NewSDKError(map[string]interface{}{
			"code":               dara.ToString(dara.Default(err["Code"], err["code"])),
			"message":            "code: " + dara.ToString(dara.IntValue(response_.StatusCode)) + ", " + dara.ToString(dara.Default(err["Message"], err["message"])) + " request id: " + dara.ToString(dara.Default(err["RequestId"], err["requestId"])),
			"data":               err,
			"description":        dara.ToString(dara.Default(err["Description"], err["description"])),
			"accessDeniedDetail": dara.Default(err["AccessDeniedDetail"], err["accessDeniedDetail"]),
		})
		_yieldErr <- _err
		return
	}

	events := make(chan *dara.SSEEvent, 1)
	dara.ReadAsSSEWithContext(ctx, response_.Body, events, _yieldErr)
	for event := range events {
		_yield <- &SSEResponse{
			StatusCode: response_.StatusCode,
			Headers:    response_.Headers,
			Event:      event,
		}
	}
	return
}
