// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"encoding/hex"
	"fmt"
	spi "github.com/alibabacloud-go/alibabacloud-gateway-spi/client"
	models "github.com/alibabacloud-go/darabonba-openapi/v2/models"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
	credential "github.com/aliyun/credentials-go/credentials"
)

// Description:
//
// This is for OpenApi SDK
type Config = models.Config
type GlobalParameters = models.GlobalParameters
type Params = models.Params
type OpenApiRequest = models.OpenApiRequest
type iSSEResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SSEResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int) *SSEResponse
	GetStatusCode() *int
	SetEvent(v *dara.SSEEvent) *SSEResponse
	GetEvent() *dara.SSEEvent
}

type SSEResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	// HTTP Status Code
	StatusCode *int           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Event      *dara.SSEEvent `json:"event,omitempty" xml:"event,omitempty" require:"true"`
}

func (s SSEResponse) String() string {
	return dara.Prettify(s)
}

func (s SSEResponse) GoString() string {
	return s.String()
}

func (s *SSEResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SSEResponse) GetStatusCode() *int {
	return s.StatusCode
}

func (s *SSEResponse) GetEvent() *dara.SSEEvent {
	return s.Event
}

func (s *SSEResponse) SetHeaders(v map[string]*string) *SSEResponse {
	s.Headers = v
	return s
}

func (s *SSEResponse) SetStatusCode(v int) *SSEResponse {
	s.StatusCode = &v
	return s
}

func (s *SSEResponse) SetEvent(v *dara.SSEEvent) *SSEResponse {
	s.Event = v
	return s
}

func (s *SSEResponse) Validate() error {
	return dara.Validate(s)
}

type iAlibabaCloudError interface {
	Error() string
	GetRetryAfter() *int64
	GetData() map[string]interface{}
	GetAccessDeniedDetail() map[string]interface{}
	GetName() *string
	GetStack() *string
	GetStatusCode() *int
	GetCode() *string
	GetMessage() *string
	GetDescription() *string
	GetRequestId() *string
}

type AlibabaCloudError struct {
	RetryAfter         *int64                 ``
	Data               map[string]interface{} ``
	AccessDeniedDetail map[string]interface{} ``
	Name               *string                ``
	Stack              *string                ``
	StatusCode         *int                   ``
	Code               *string                ``
	Message            *string                ``
	Description        *string                ``
	RequestId          *string                ``
}

func (err AlibabaCloudError) Error() string {
	if err.Message == nil {
		str := fmt.Sprintf("AlibabaCloudError:\n   Name: %s\n   Code: %s\n",
			dara.StringValue(err.Name), dara.StringValue(err.Code))
		err.Message = dara.String(str)
	}
	return dara.StringValue(err.Message)
}

func (s *AlibabaCloudError) GetRetryAfter() *int64 {
	return s.RetryAfter
}

func (s *AlibabaCloudError) GetData() map[string]interface{} {
	return s.Data
}

func (s *AlibabaCloudError) GetAccessDeniedDetail() map[string]interface{} {
	return s.AccessDeniedDetail
}

func (s *AlibabaCloudError) GetName() *string {
	return s.Name
}

func (s *AlibabaCloudError) GetStack() *string {
	return s.Stack
}

func (s *AlibabaCloudError) GetStatusCode() *int {
	return s.StatusCode
}

func (s *AlibabaCloudError) GetCode() *string {
	return s.Code
}

func (s *AlibabaCloudError) GetMessage() *string {
	return s.Message
}

func (s *AlibabaCloudError) GetDescription() *string {
	return s.Description
}

func (s *AlibabaCloudError) GetRequestId() *string {
	return s.RequestId
}

type iClientError interface {
	Error() string
	GetStatusCode() *int
	GetCode() *string
	GetMessage() *string
	GetDescription() *string
	GetRequestId() *string
	GetRetryAfter() *int64
	GetData() map[string]interface{}
	GetName() *string
	GetStack() *string
	GetAccessDeniedDetail() map[string]interface{}
}

type ClientError struct {
	StatusCode         *int                   ``
	Code               *string                ``
	Message            *string                ``
	Description        *string                ``
	RequestId          *string                ``
	RetryAfter         *int64                 ``
	Data               map[string]interface{} ``
	Name               *string                ``
	Stack              *string                ``
	AccessDeniedDetail map[string]interface{} ``
}

func (err ClientError) Error() string {
	if err.Message == nil {
		str := fmt.Sprintf("ClientError:\n   Name: %s\n   Code: %s\n",
			dara.StringValue(err.Name), dara.StringValue(err.Code))
		err.Message = dara.String(str)
	}
	return dara.StringValue(err.Message)
}

func (s *ClientError) GetStatusCode() *int {
	return s.StatusCode
}

func (s *ClientError) GetCode() *string {
	return s.Code
}

func (s *ClientError) GetMessage() *string {
	return s.Message
}

func (s *ClientError) GetDescription() *string {
	return s.Description
}

func (s *ClientError) GetRequestId() *string {
	return s.RequestId
}

func (s *ClientError) GetRetryAfter() *int64 {
	return s.RetryAfter
}

func (s *ClientError) GetData() map[string]interface{} {
	return s.Data
}

func (s *ClientError) GetName() *string {
	return s.Name
}

func (s *ClientError) GetStack() *string {
	return s.Stack
}

func (s *ClientError) GetAccessDeniedDetail() map[string]interface{} {
	return s.AccessDeniedDetail
}

type iServerError interface {
	Error() string
	GetStatusCode() *int
	GetCode() *string
	GetMessage() *string
	GetDescription() *string
	GetRequestId() *string
	GetRetryAfter() *int64
	GetData() map[string]interface{}
	GetAccessDeniedDetail() map[string]interface{}
	GetName() *string
	GetStack() *string
}

type ServerError struct {
	StatusCode         *int                   ``
	Code               *string                ``
	Message            *string                ``
	Description        *string                ``
	RequestId          *string                ``
	RetryAfter         *int64                 ``
	Data               map[string]interface{} ``
	AccessDeniedDetail map[string]interface{} ``
	Name               *string                ``
	Stack              *string                ``
}

func (err ServerError) Error() string {
	if err.Message == nil {
		str := fmt.Sprintf("ServerError:\n   Name: %s\n   Code: %s\n",
			dara.StringValue(err.Name), dara.StringValue(err.Code))
		err.Message = dara.String(str)
	}
	return dara.StringValue(err.Message)
}

func (s *ServerError) GetStatusCode() *int {
	return s.StatusCode
}

func (s *ServerError) GetCode() *string {
	return s.Code
}

func (s *ServerError) GetMessage() *string {
	return s.Message
}

func (s *ServerError) GetDescription() *string {
	return s.Description
}

func (s *ServerError) GetRequestId() *string {
	return s.RequestId
}

func (s *ServerError) GetRetryAfter() *int64 {
	return s.RetryAfter
}

func (s *ServerError) GetData() map[string]interface{} {
	return s.Data
}

func (s *ServerError) GetAccessDeniedDetail() map[string]interface{} {
	return s.AccessDeniedDetail
}

func (s *ServerError) GetName() *string {
	return s.Name
}

func (s *ServerError) GetStack() *string {
	return s.Stack
}

type iThrottlingError interface {
	Error() string
	GetStatusCode() *int
	GetCode() *string
	GetMessage() *string
	GetDescription() *string
	GetRequestId() *string
	GetData() map[string]interface{}
	GetAccessDeniedDetail() map[string]interface{}
	GetName() *string
	GetStack() *string
	GetRetryAfter() *int64
}

type ThrottlingError struct {
	StatusCode         *int                   ``
	Code               *string                ``
	Message            *string                ``
	Description        *string                ``
	RequestId          *string                ``
	Data               map[string]interface{} ``
	AccessDeniedDetail map[string]interface{} ``
	Name               *string                ``
	Stack              *string                ``
	RetryAfter         *int64                 ``
}

func (err ThrottlingError) Error() string {
	if err.Message == nil {
		str := fmt.Sprintf("ThrottlingError:\n   Name: %s\n   Code: %s\n",
			dara.StringValue(err.Name), dara.StringValue(err.Code))
		err.Message = dara.String(str)
	}
	return dara.StringValue(err.Message)
}

func (s *ThrottlingError) GetStatusCode() *int {
	return s.StatusCode
}

func (s *ThrottlingError) GetCode() *string {
	return s.Code
}

func (s *ThrottlingError) GetMessage() *string {
	return s.Message
}

func (s *ThrottlingError) GetDescription() *string {
	return s.Description
}

func (s *ThrottlingError) GetRequestId() *string {
	return s.RequestId
}

func (s *ThrottlingError) GetData() map[string]interface{} {
	return s.Data
}

func (s *ThrottlingError) GetAccessDeniedDetail() map[string]interface{} {
	return s.AccessDeniedDetail
}

func (s *ThrottlingError) GetName() *string {
	return s.Name
}

func (s *ThrottlingError) GetStack() *string {
	return s.Stack
}

func (s *ThrottlingError) GetRetryAfter() *int64 {
	return s.RetryAfter
}

type Client struct {
	DisableSDKError      *bool
	Endpoint             *string
	RegionId             *string
	Protocol             *string
	Method               *string
	UserAgent            *string
	EndpointRule         *string
	EndpointMap          map[string]*string
	Suffix               *string
	ReadTimeout          *int
	ConnectTimeout       *int
	HttpProxy            *string
	HttpsProxy           *string
	Socks5Proxy          *string
	Socks5NetWork        *string
	NoProxy              *string
	Network              *string
	ProductId            *string
	MaxIdleConns         *int
	EndpointType         *string
	OpenPlatformEndpoint *string
	Credential           credential.Credential
	SignatureVersion     *string
	SignatureAlgorithm   *string
	Headers              map[string]*string
	Spi                  spi.ClientInterface
	GlobalParameters     *openapiutil.GlobalParameters
	Key                  *string
	Cert                 *string
	Ca                   *string
	DisableHttp2         *bool
	RetryOptions         *dara.RetryOptions
	HttpClient           dara.HttpClient
}

// Description:
//
// # Init client with Config
//
// @param config - config contains the necessary information to create a client
func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	if dara.IsNil(config) {
		_err = &ClientError{
			Code:    dara.String("ParameterMissing"),
			Message: dara.String("'config' can not be unset"),
		}
		return _err
	}

	if (!dara.IsNil(config.AccessKeyId) && dara.StringValue(config.AccessKeyId) != "") && (!dara.IsNil(config.AccessKeySecret) && dara.StringValue(config.AccessKeySecret) != "") {
		if !dara.IsNil(config.SecurityToken) && dara.StringValue(config.SecurityToken) != "" {
			config.Type = dara.String("sts")
		} else {
			config.Type = dara.String("access_key")
		}

		credentialConfig := &credential.Config{
			AccessKeyId:     config.AccessKeyId,
			Type:            config.Type,
			AccessKeySecret: config.AccessKeySecret,
		}
		credentialConfig.SecurityToken = config.SecurityToken
		client.Credential, _err = credential.NewCredential(credentialConfig)
		if _err != nil {
			return _err
		}

	} else if !dara.IsNil(config.BearerToken) && dara.StringValue(config.BearerToken) != "" {
		cc := &credential.Config{
			Type:        dara.String("bearer"),
			BearerToken: config.BearerToken,
		}
		client.Credential, _err = credential.NewCredential(cc)
		if _err != nil {
			return _err
		}

	} else if !dara.IsNil(config.Credential) {
		client.Credential = config.Credential
	}

	client.Endpoint = config.Endpoint
	client.EndpointType = config.EndpointType
	client.Network = config.Network
	client.Suffix = config.Suffix
	client.Protocol = config.Protocol
	client.Method = config.Method
	client.RegionId = config.RegionId
	client.UserAgent = config.UserAgent
	client.ReadTimeout = config.ReadTimeout
	client.ConnectTimeout = config.ConnectTimeout
	client.HttpProxy = config.HttpProxy
	client.HttpsProxy = config.HttpsProxy
	client.NoProxy = config.NoProxy
	client.Socks5Proxy = config.Socks5Proxy
	client.Socks5NetWork = config.Socks5NetWork
	client.MaxIdleConns = config.MaxIdleConns
	client.SignatureVersion = config.SignatureVersion
	client.SignatureAlgorithm = config.SignatureAlgorithm
	client.GlobalParameters = config.GlobalParameters
	client.Key = config.Key
	client.Cert = config.Cert
	client.Ca = config.Ca
	client.DisableHttp2 = config.DisableHttp2
	client.RetryOptions = config.RetryOptions
	client.HttpClient = config.HttpClient
	return nil
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
func (client *Client) DoRPCRequest(action *string, version *string, protocol *string, method *string, authType *string, bodyType *string, request *openapiutil.OpenApiRequest, runtime *dara.RuntimeOptions) (_result map[string]interface{}, _err error) {
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
				extendsHeaders)
		} else {
			request_.Headers = dara.Merge(map[string]*string{
				"host":          client.Endpoint,
				"x-acs-version": version,
				"x-acs-action":  action,
				"user-agent":    openapiutil.GetUserAgent(client.UserAgent),
			}, globalHeaders,
				extendsHeaders,
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

			credentialType := dara.StringValue(credentialModel.Type)
			if credentialType == "bearer" {
				bearerToken := dara.StringValue(credentialModel.BearerToken)
				request_.Query["BearerToken"] = dara.String(bearerToken)
				request_.Query["SignatureType"] = dara.String("BEARERTOKEN")
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

		response_, _err := dara.DoRequest(request_, _runtime)
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

		_result, _err = doRPCRequest_opResponse(response_, client, bodyType)
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
func (client *Client) DoROARequest(action *string, version *string, protocol *string, method *string, authType *string, pathname *string, bodyType *string, request *openapiutil.OpenApiRequest, runtime *dara.RuntimeOptions) (_result map[string]interface{}, _err error) {
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

			credentialType := dara.StringValue(credentialModel.Type)
			if credentialType == "bearer" {
				bearerToken := dara.StringValue(credentialModel.BearerToken)
				request_.Headers["x-acs-bearer-token"] = dara.String(bearerToken)
				request_.Headers["x-acs-signature-type"] = dara.String("BEARERTOKEN")
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

		response_, _err := dara.DoRequest(request_, _runtime)
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

		_result, _err = doROARequest_opResponse(response_, client, bodyType)
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
func (client *Client) DoROARequestWithForm(action *string, version *string, protocol *string, method *string, authType *string, pathname *string, bodyType *string, request *openapiutil.OpenApiRequest, runtime *dara.RuntimeOptions) (_result map[string]interface{}, _err error) {
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

			credentialType := dara.StringValue(credentialModel.Type)
			if credentialType == "bearer" {
				bearerToken := dara.StringValue(credentialModel.BearerToken)
				request_.Headers["x-acs-bearer-token"] = dara.String(bearerToken)
				request_.Headers["x-acs-signature-type"] = dara.String("BEARERTOKEN")
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

		response_, _err := dara.DoRequest(request_, _runtime)
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

		_result, _err = doROARequestWithForm_opResponse(response_, client, bodyType)
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
func (client *Client) DoRequest(params *openapiutil.Params, request *openapiutil.OpenApiRequest, runtime *dara.RuntimeOptions) (_result map[string]interface{}, _err error) {
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

			authType := dara.StringValue(credentialModel.Type)
			if authType == "bearer" {
				bearerToken := dara.StringValue(credentialModel.BearerToken)
				request_.Headers["x-acs-bearer-token"] = dara.String(bearerToken)
				if dara.StringValue(params.Style) == "RPC" {
					request_.Query["SignatureType"] = dara.String("BEARERTOKEN")
				} else {
					request_.Headers["x-acs-signature-type"] = dara.String("BEARERTOKEN")
				}

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

		response_, _err := dara.DoRequest(request_, _runtime)
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

		_result, _err = doRequest_opResponse(response_, client, params)
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
func (client *Client) Execute(params *openapiutil.Params, request *openapiutil.OpenApiRequest, runtime *dara.RuntimeOptions) (_result map[string]interface{}, _err error) {
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
		"disableHttp2":   dara.ForceBoolean(dara.Default(dara.BoolValue(client.DisableHttp2), false)),
		"httpClient":     client.HttpClient,
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
		response_, _err := dara.DoRequest(request_, _runtime)
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

func (client *Client) CallSSEApi(params *openapiutil.Params, request *openapiutil.OpenApiRequest, runtime *dara.RuntimeOptions, _yield chan *SSEResponse, _yieldErr chan error) {
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

			authType := dara.StringValue(credentialModel.Type)
			if authType == "bearer" {
				bearerToken := dara.StringValue(credentialModel.BearerToken)
				request_.Headers["x-acs-bearer-token"] = dara.String(bearerToken)
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

		response_, _err := dara.DoRequest(request_, _runtime)
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

		callSSEApi_opResponse(_yield, _yieldErr, response_)
		_err = <-_yieldErr
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

		return
	}
	_yieldErr <- _resultErr
	return
}

func (client *Client) CallApi(params *openapiutil.Params, request *openapiutil.OpenApiRequest, runtime *dara.RuntimeOptions) (_result map[string]interface{}, _err error) {
	if dara.IsNil(params) {
		_err = &ClientError{
			Code:    dara.String("ParameterMissing"),
			Message: dara.String("'params' can not be unset"),
		}
		return _result, _err
	}

	if dara.IsNil(client.SignatureAlgorithm) || dara.StringValue(client.SignatureAlgorithm) != "v2" {
		_body, _err := client.DoRequest(params, request, runtime)
		if _err != nil {
			return _result, _err
		}
		_result = _body
		return _result, _err
	} else if (dara.StringValue(params.Style) == "ROA") && (dara.StringValue(params.ReqBodyType) == "json") {
		_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, request, runtime)
		if _err != nil {
			return _result, _err
		}
		_result = _body
		return _result, _err
	} else if dara.StringValue(params.Style) == "ROA" {
		_body, _err := client.DoROARequestWithForm(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, request, runtime)
		if _err != nil {
			return _result, _err
		}
		_result = _body
		return _result, _err
	} else {
		_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, request, runtime)
		if _err != nil {
			return _result, _err
		}
		_result = _body
		return _result, _err
	}

}

// Description:
//
// # Get accesskey id by using credential
//
// @return accesskey id
func (client *Client) GetAccessKeyId() (_result *string, _err error) {
	if dara.IsNil(client.Credential) {
		_result = dara.String("")
		return _result, _err
	}

	accessKeyIdTmp, _err := client.Credential.GetAccessKeyId()
	accessKeyId := dara.StringValue(accessKeyIdTmp)
	if _err != nil {
		return _result, _err
	}

	_result = dara.String(accessKeyId)
	return _result, _err
}

// Description:
//
// # Get accesskey secret by using credential
//
// @return accesskey secret
func (client *Client) GetAccessKeySecret() (_result *string, _err error) {
	if dara.IsNil(client.Credential) {
		_result = dara.String("")
		return _result, _err
	}

	secretTmp, _err := client.Credential.GetAccessKeySecret()
	secret := dara.StringValue(secretTmp)
	if _err != nil {
		return _result, _err
	}

	_result = dara.String(secret)
	return _result, _err
}

// Description:
//
// # Get security token by using credential
//
// @return security token
func (client *Client) GetSecurityToken() (_result *string, _err error) {
	if dara.IsNil(client.Credential) {
		_result = dara.String("")
		return _result, _err
	}

	tokenTmp, _err := client.Credential.GetSecurityToken()
	token := dara.StringValue(tokenTmp)
	if _err != nil {
		return _result, _err
	}

	_result = dara.String(token)
	return _result, _err
}

// Description:
//
// # Get bearer token by credential
//
// @return bearer token
func (client *Client) GetBearerToken() (_result *string, _err error) {
	if dara.IsNil(client.Credential) {
		_result = dara.String("")
		return _result, _err
	}

	token := dara.StringValue(client.Credential.GetBearerToken())
	_result = dara.String(token)
	return _result, _err
}

// Description:
//
// # Get credential type by credential
//
// @return credential type e.g. access_key
func (client *Client) GetType() (_result *string, _err error) {
	if dara.IsNil(client.Credential) {
		_result = dara.String("")
		return _result, _err
	}

	authType := dara.StringValue(client.Credential.GetType())
	_result = dara.String(authType)
	return _result, _err
}

// Description:
//
// # If the endpointRule and config.endpoint are empty, throw error
//
// @param config - config contains the necessary information to create a client
func (client *Client) CheckConfig(config *openapiutil.Config) (_err error) {
	if dara.IsNil(client.EndpointRule) && dara.IsNil(config.Endpoint) {
		_err = &ClientError{
			Code:    dara.String("ParameterMissing"),
			Message: dara.String("'config.endpoint' can not be empty"),
		}
		return _err
	}

	return _err
}

// Description:
//
// set gateway client
//
// @param spi - .
func (client *Client) SetGatewayClient(spi spi.ClientInterface) (_err error) {
	client.Spi = spi
	return _err
}

// Description:
//
// set RPC header for debug
//
// @param headers - headers for debug, this header can be used only once.
func (client *Client) SetRpcHeaders(headers map[string]*string) (_err error) {
	client.Headers = headers
	return _err
}

// Description:
//
// get RPC header for debug
func (client *Client) GetRpcHeaders() (_result map[string]*string, _err error) {
	headers := client.Headers
	client.Headers = nil
	_result = headers
	return _result, _err
}

func (client *Client) GetAccessDeniedDetail(err map[string]interface{}) (_result map[string]interface{}) {
	var accessDeniedDetail map[string]interface{}
	if !dara.IsNil(err["AccessDeniedDetail"]) {
		detail1 := dara.ToMap(err["AccessDeniedDetail"])
		accessDeniedDetail = detail1
	} else if !dara.IsNil(err["accessDeniedDetail"]) {
		detail2 := dara.ToMap(err["accessDeniedDetail"])
		accessDeniedDetail = detail2
	}

	_result = accessDeniedDetail
	return _result
}

func doRPCRequest_opResponse(response_ *dara.Response, client *Client, bodyType *string) (_result map[string]interface{}, _err error) {
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

func doROARequest_opResponse(response_ *dara.Response, client *Client, bodyType *string) (_result map[string]interface{}, _err error) {
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

func doROARequestWithForm_opResponse(response_ *dara.Response, client *Client, bodyType *string) (_result map[string]interface{}, _err error) {
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

func doRequest_opResponse(response_ *dara.Response, client *Client, params *openapiutil.Params) (_result map[string]interface{}, _err error) {
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

func callSSEApi_opResponse(_yield chan *SSEResponse, _yieldErr chan error, response_ *dara.Response) {
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
	dara.ReadAsSSE(response_.Body, events, _yieldErr)
	for event := range events {
		_yield <- &SSEResponse{
			StatusCode: response_.StatusCode,
			Headers:    response_.Headers,
			Event:      event,
		}
	}
	return
}
