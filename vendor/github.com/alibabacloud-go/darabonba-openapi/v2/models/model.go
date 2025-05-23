package models

import (
	"io"
	"github.com/alibabacloud-go/tea/dara"
	credential "github.com/aliyun/credentials-go/credentials"
  )

// Description:
// 
// This is for OpenApi Util
type iGlobalParameters interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GlobalParameters
	GetHeaders() map[string]*string 
	SetQueries(v map[string]*string) *GlobalParameters
	GetQueries() map[string]*string 
  }
  
  type GlobalParameters struct {
	dara.Model
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Queries map[string]*string `json:"queries,omitempty" xml:"queries,omitempty"`
  }
  
  func (s GlobalParameters) String() string {
	return dara.Prettify(s)
  }
  
  func (s GlobalParameters) GoString() string {
	return s.String()
  }
  
  func (s *GlobalParameters) GetHeaders() map[string]*string  {
	return s.Headers
  }
  
  func (s *GlobalParameters) GetQueries() map[string]*string  {
	return s.Queries
  }
  
  func (s *GlobalParameters) SetHeaders(v map[string]*string) *GlobalParameters {
	s.Headers = v
	return s
  }
  
  func (s *GlobalParameters) SetQueries(v map[string]*string) *GlobalParameters {
	s.Queries = v
	return s
  }
  
  type iConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKeyId(v string) *Config
	GetAccessKeyId() *string 
	SetAccessKeySecret(v string) *Config
	GetAccessKeySecret() *string 
	SetSecurityToken(v string) *Config
	GetSecurityToken() *string 
	SetBearerToken(v string) *Config
	GetBearerToken() *string 
	SetProtocol(v string) *Config
	GetProtocol() *string 
	SetMethod(v string) *Config
	GetMethod() *string 
	SetRegionId(v string) *Config
	GetRegionId() *string 
	SetReadTimeout(v int) *Config
	GetReadTimeout() *int 
	SetConnectTimeout(v int) *Config
	GetConnectTimeout() *int 
	SetHttpProxy(v string) *Config
	GetHttpProxy() *string 
	SetHttpsProxy(v string) *Config
	GetHttpsProxy() *string 
	SetCredential(v credential.Credential) *Config
	GetCredential() credential.Credential 
	SetEndpoint(v string) *Config
	GetEndpoint() *string 
	SetNoProxy(v string) *Config
	GetNoProxy() *string 
	SetMaxIdleConns(v int) *Config
	GetMaxIdleConns() *int 
	SetNetwork(v string) *Config
	GetNetwork() *string 
	SetUserAgent(v string) *Config
	GetUserAgent() *string 
	SetSuffix(v string) *Config
	GetSuffix() *string 
	SetSocks5Proxy(v string) *Config
	GetSocks5Proxy() *string 
	SetSocks5NetWork(v string) *Config
	GetSocks5NetWork() *string 
	SetEndpointType(v string) *Config
	GetEndpointType() *string 
	SetOpenPlatformEndpoint(v string) *Config
	GetOpenPlatformEndpoint() *string 
	SetType(v string) *Config
	GetType() *string 
	SetSignatureVersion(v string) *Config
	GetSignatureVersion() *string 
	SetSignatureAlgorithm(v string) *Config
	GetSignatureAlgorithm() *string 
	SetGlobalParameters(v *GlobalParameters) *Config
	GetGlobalParameters() *GlobalParameters 
	SetKey(v string) *Config
	GetKey() *string 
	SetCert(v string) *Config
	GetCert() *string 
	SetCa(v string) *Config
	GetCa() *string 
	SetDisableHttp2(v bool) *Config
	GetDisableHttp2() *bool 
	SetRetryOptions(v *dara.RetryOptions) *Config
	GetRetryOptions() *dara.RetryOptions 
  }
  
  // Description:
  // 
  // Model for initing client
  type Config struct {
	dara.Model
	// accesskey id
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// accesskey secret
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// security token
	SecurityToken *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
	// bearer token
	// 
	// example:
	// 
	// the-bearer-token
	BearerToken *string `json:"bearerToken,omitempty" xml:"bearerToken,omitempty"`
	// http protocol
	// 
	// example:
	// 
	// http
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// http method
	// 
	// example:
	// 
	// GET
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// region id
	// 
	// example:
	// 
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// read timeout
	// 
	// example:
	// 
	// 10
	ReadTimeout *int `json:"readTimeout,omitempty" xml:"readTimeout,omitempty"`
	// connect timeout
	// 
	// example:
	// 
	// 10
	ConnectTimeout *int `json:"connectTimeout,omitempty" xml:"connectTimeout,omitempty"`
	// http proxy
	// 
	// example:
	// 
	// http://localhost
	HttpProxy *string `json:"httpProxy,omitempty" xml:"httpProxy,omitempty"`
	// https proxy
	// 
	// example:
	// 
	// https://localhost
	HttpsProxy *string `json:"httpsProxy,omitempty" xml:"httpsProxy,omitempty"`
	// credential
	Credential credential.Credential `json:"credential,omitempty" xml:"credential,omitempty"`
	// endpoint
	// 
	// example:
	// 
	// cs.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// proxy white list
	// 
	// example:
	// 
	// http://localhost
	NoProxy *string `json:"noProxy,omitempty" xml:"noProxy,omitempty"`
	// max idle conns
	// 
	// example:
	// 
	// 3
	MaxIdleConns *int `json:"maxIdleConns,omitempty" xml:"maxIdleConns,omitempty"`
	// network for endpoint
	// 
	// example:
	// 
	// public
	Network *string `json:"network,omitempty" xml:"network,omitempty"`
	// user agent
	// 
	// example:
	// 
	// Alibabacloud/1
	UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
	// suffix for endpoint
	// 
	// example:
	// 
	// aliyun
	Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty"`
	// socks5 proxy
	Socks5Proxy *string `json:"socks5Proxy,omitempty" xml:"socks5Proxy,omitempty"`
	// socks5 network
	// 
	// example:
	// 
	// TCP
	Socks5NetWork *string `json:"socks5NetWork,omitempty" xml:"socks5NetWork,omitempty"`
	// endpoint type
	// 
	// example:
	// 
	// internal
	EndpointType *string `json:"endpointType,omitempty" xml:"endpointType,omitempty"`
	// OpenPlatform endpoint
	// 
	// example:
	// 
	// openplatform.aliyuncs.com
	OpenPlatformEndpoint *string `json:"openPlatformEndpoint,omitempty" xml:"openPlatformEndpoint,omitempty"`
	// Deprecated
	// 
	// credential type
	// 
	// example:
	// 
	// access_key
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Signature Version
	// 
	// example:
	// 
	// v1
	SignatureVersion *string `json:"signatureVersion,omitempty" xml:"signatureVersion,omitempty"`
	// Signature Algorithm
	// 
	// example:
	// 
	// ACS3-HMAC-SHA256
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty" xml:"signatureAlgorithm,omitempty"`
	// Global Parameters
	GlobalParameters *GlobalParameters `json:"globalParameters,omitempty" xml:"globalParameters,omitempty"`
	// privite key for client certificate
	// 
	// example:
	// 
	// MIIEvQ
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// client certificate
	// 
	// example:
	// 
	// -----BEGIN CERTIFICATE-----
	// 
	// xxx-----END CERTIFICATE-----
	Cert *string `json:"cert,omitempty" xml:"cert,omitempty"`
	// server certificate
	// 
	// example:
	// 
	// -----BEGIN CERTIFICATE-----
	// 
	// xxx-----END CERTIFICATE-----
	Ca *string `json:"ca,omitempty" xml:"ca,omitempty"`
	// disable HTTP/2
	// 
	// example:
	// 
	// false
	DisableHttp2 *bool `json:"disableHttp2,omitempty" xml:"disableHttp2,omitempty"`
	// retry options
	RetryOptions *dara.RetryOptions `json:"retryOptions,omitempty" xml:"retryOptions,omitempty"`
	// http client
	HttpClient  dara.HttpClient `json:"httpClient,omitempty" xml:"httpClient,omitempty"`
  }
  
  func (s Config) String() string {
	return dara.Prettify(s)
  }
  
  func (s Config) GoString() string {
	return s.String()
  }
  
  func (s *Config) GetAccessKeyId() *string  {
	return s.AccessKeyId
  }
  
  func (s *Config) GetAccessKeySecret() *string  {
	return s.AccessKeySecret
  }
  
  func (s *Config) GetSecurityToken() *string  {
	return s.SecurityToken
  }
  
  func (s *Config) GetBearerToken() *string  {
	return s.BearerToken
  }
  
  func (s *Config) GetProtocol() *string  {
	return s.Protocol
  }
  
  func (s *Config) GetMethod() *string  {
	return s.Method
  }
  
  func (s *Config) GetRegionId() *string  {
	return s.RegionId
  }
  
  func (s *Config) GetReadTimeout() *int  {
	return s.ReadTimeout
  }
  
  func (s *Config) GetConnectTimeout() *int  {
	return s.ConnectTimeout
  }
  
  func (s *Config) GetHttpProxy() *string  {
	return s.HttpProxy
  }
  
  func (s *Config) GetHttpsProxy() *string  {
	return s.HttpsProxy
  }
  
  func (s *Config) GetCredential() credential.Credential  {
	return s.Credential
  }
  
  func (s *Config) GetEndpoint() *string  {
	return s.Endpoint
  }
  
  func (s *Config) GetNoProxy() *string  {
	return s.NoProxy
  }
  
  func (s *Config) GetMaxIdleConns() *int  {
	return s.MaxIdleConns
  }
  
  func (s *Config) GetNetwork() *string  {
	return s.Network
  }
  
  func (s *Config) GetUserAgent() *string  {
	return s.UserAgent
  }
  
  func (s *Config) GetSuffix() *string  {
	return s.Suffix
  }
  
  func (s *Config) GetSocks5Proxy() *string  {
	return s.Socks5Proxy
  }
  
  func (s *Config) GetSocks5NetWork() *string  {
	return s.Socks5NetWork
  }
  
  func (s *Config) GetEndpointType() *string  {
	return s.EndpointType
  }
  
  func (s *Config) GetOpenPlatformEndpoint() *string  {
	return s.OpenPlatformEndpoint
  }
  
  func (s *Config) GetType() *string  {
	return s.Type
  }
  
  func (s *Config) GetSignatureVersion() *string  {
	return s.SignatureVersion
  }
  
  func (s *Config) GetSignatureAlgorithm() *string  {
	return s.SignatureAlgorithm
  }
  
  func (s *Config) GetGlobalParameters() *GlobalParameters  {
	return s.GlobalParameters
  }
  
  func (s *Config) GetKey() *string  {
	return s.Key
  }
  
  func (s *Config) GetCert() *string  {
	return s.Cert
  }
  
  func (s *Config) GetCa() *string  {
	return s.Ca
  }
  
  func (s *Config) GetDisableHttp2() *bool  {
	return s.DisableHttp2
  }
  
  func (s *Config) GetRetryOptions() *dara.RetryOptions  {
	return s.RetryOptions
  }

  func (s *Config) GetHttpClient() dara.HttpClient {
	return s.HttpClient
  }
  
  func (s *Config) SetAccessKeyId(v string) *Config {
	s.AccessKeyId = &v
	return s
  }
  
  func (s *Config) SetAccessKeySecret(v string) *Config {
	s.AccessKeySecret = &v
	return s
  }
  
  func (s *Config) SetSecurityToken(v string) *Config {
	s.SecurityToken = &v
	return s
  }
  
  func (s *Config) SetBearerToken(v string) *Config {
	s.BearerToken = &v
	return s
  }
  
  func (s *Config) SetProtocol(v string) *Config {
	s.Protocol = &v
	return s
  }
  
  func (s *Config) SetMethod(v string) *Config {
	s.Method = &v
	return s
  }
  
  func (s *Config) SetRegionId(v string) *Config {
	s.RegionId = &v
	return s
  }
  
  func (s *Config) SetReadTimeout(v int) *Config {
	s.ReadTimeout = &v
	return s
  }
  
  func (s *Config) SetConnectTimeout(v int) *Config {
	s.ConnectTimeout = &v
	return s
  }
  
  func (s *Config) SetHttpProxy(v string) *Config {
	s.HttpProxy = &v
	return s
  }
  
  func (s *Config) SetHttpsProxy(v string) *Config {
	s.HttpsProxy = &v
	return s
  }
  
  func (s *Config) SetCredential(v credential.Credential) *Config {
	s.Credential = v
	return s
  }
  
  func (s *Config) SetEndpoint(v string) *Config {
	s.Endpoint = &v
	return s
  }
  
  func (s *Config) SetNoProxy(v string) *Config {
	s.NoProxy = &v
	return s
  }
  
  func (s *Config) SetMaxIdleConns(v int) *Config {
	s.MaxIdleConns = &v
	return s
  }
  
  func (s *Config) SetNetwork(v string) *Config {
	s.Network = &v
	return s
  }
  
  func (s *Config) SetUserAgent(v string) *Config {
	s.UserAgent = &v
	return s
  }
  
  func (s *Config) SetSuffix(v string) *Config {
	s.Suffix = &v
	return s
  }
  
  func (s *Config) SetSocks5Proxy(v string) *Config {
	s.Socks5Proxy = &v
	return s
  }
  
  func (s *Config) SetSocks5NetWork(v string) *Config {
	s.Socks5NetWork = &v
	return s
  }
  
  func (s *Config) SetEndpointType(v string) *Config {
	s.EndpointType = &v
	return s
  }
  
  func (s *Config) SetOpenPlatformEndpoint(v string) *Config {
	s.OpenPlatformEndpoint = &v
	return s
  }
  
  func (s *Config) SetType(v string) *Config {
	s.Type = &v
	return s
  }
  
  func (s *Config) SetSignatureVersion(v string) *Config {
	s.SignatureVersion = &v
	return s
  }
  
  func (s *Config) SetSignatureAlgorithm(v string) *Config {
	s.SignatureAlgorithm = &v
	return s
  }
  
  func (s *Config) SetGlobalParameters(v *GlobalParameters) *Config {
	s.GlobalParameters = v
	return s
  }
  
  func (s *Config) SetKey(v string) *Config {
	s.Key = &v
	return s
  }
  
  func (s *Config) SetCert(v string) *Config {
	s.Cert = &v
	return s
  }
  
  func (s *Config) SetCa(v string) *Config {
	s.Ca = &v
	return s
  }
  
  func (s *Config) SetDisableHttp2(v bool) *Config {
	s.DisableHttp2 = &v
	return s
  }
  
  func (s *Config) SetRetryOptions(v *dara.RetryOptions) *Config {
	s.RetryOptions = v
	return s
  }

  func (s *Config) SetHttpClient(v dara.HttpClient) *Config {
	s.HttpClient = v
	return s
  }
  
  type iParams interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *Params
	GetAction() *string 
	SetVersion(v string) *Params
	GetVersion() *string 
	SetProtocol(v string) *Params
	GetProtocol() *string 
	SetPathname(v string) *Params
	GetPathname() *string 
	SetMethod(v string) *Params
	GetMethod() *string 
	SetAuthType(v string) *Params
	GetAuthType() *string 
	SetBodyType(v string) *Params
	GetBodyType() *string 
	SetReqBodyType(v string) *Params
	GetReqBodyType() *string 
	SetStyle(v string) *Params
	GetStyle() *string 
  }
  
  type Params struct {
	dara.Model
	Action *string `json:"action,omitempty" xml:"action,omitempty" require:"true"`
	Version *string `json:"version,omitempty" xml:"version,omitempty" require:"true"`
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
	Pathname *string `json:"pathname,omitempty" xml:"pathname,omitempty" require:"true"`
	Method *string `json:"method,omitempty" xml:"method,omitempty" require:"true"`
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty" require:"true"`
	BodyType *string `json:"bodyType,omitempty" xml:"bodyType,omitempty" require:"true"`
	ReqBodyType *string `json:"reqBodyType,omitempty" xml:"reqBodyType,omitempty" require:"true"`
	Style *string `json:"style,omitempty" xml:"style,omitempty"`
  }
  
  func (s Params) String() string {
	return dara.Prettify(s)
  }
  
  func (s Params) GoString() string {
	return s.String()
  }
  
  func (s *Params) GetAction() *string  {
	return s.Action
  }
  
  func (s *Params) GetVersion() *string  {
	return s.Version
  }
  
  func (s *Params) GetProtocol() *string  {
	return s.Protocol
  }
  
  func (s *Params) GetPathname() *string  {
	return s.Pathname
  }
  
  func (s *Params) GetMethod() *string  {
	return s.Method
  }
  
  func (s *Params) GetAuthType() *string  {
	return s.AuthType
  }
  
  func (s *Params) GetBodyType() *string  {
	return s.BodyType
  }
  
  func (s *Params) GetReqBodyType() *string  {
	return s.ReqBodyType
  }
  
  func (s *Params) GetStyle() *string  {
	return s.Style
  }
  
  func (s *Params) SetAction(v string) *Params {
	s.Action = &v
	return s
  }
  
  func (s *Params) SetVersion(v string) *Params {
	s.Version = &v
	return s
  }
  
  func (s *Params) SetProtocol(v string) *Params {
	s.Protocol = &v
	return s
  }
  
  func (s *Params) SetPathname(v string) *Params {
	s.Pathname = &v
	return s
  }
  
  func (s *Params) SetMethod(v string) *Params {
	s.Method = &v
	return s
  }
  
  func (s *Params) SetAuthType(v string) *Params {
	s.AuthType = &v
	return s
  }
  
  func (s *Params) SetBodyType(v string) *Params {
	s.BodyType = &v
	return s
  }
  
  func (s *Params) SetReqBodyType(v string) *Params {
	s.ReqBodyType = &v
	return s
  }
  
  func (s *Params) SetStyle(v string) *Params {
	s.Style = &v
	return s
  }
  
  type iOpenApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenApiRequest
	GetHeaders() map[string]*string 
	SetQuery(v map[string]*string) *OpenApiRequest
	GetQuery() map[string]*string 
	SetBody(v interface{}) *OpenApiRequest
	GetBody() interface{} 
	SetStream(v io.Reader) *OpenApiRequest
	GetStream() io.Reader 
	SetHostMap(v map[string]*string) *OpenApiRequest
	GetHostMap() map[string]*string 
	SetEndpointOverride(v string) *OpenApiRequest
	GetEndpointOverride() *string 
  }
  
  type OpenApiRequest struct {
	dara.Model
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Query map[string]*string `json:"query,omitempty" xml:"query,omitempty"`
	Body interface{} `json:"body,omitempty" xml:"body,omitempty"`
	Stream io.Reader `json:"stream,omitempty" xml:"stream,omitempty"`
	HostMap map[string]*string `json:"hostMap,omitempty" xml:"hostMap,omitempty"`
	EndpointOverride *string `json:"endpointOverride,omitempty" xml:"endpointOverride,omitempty"`
  }
  
  func (s OpenApiRequest) String() string {
	return dara.Prettify(s)
  }
  
  func (s OpenApiRequest) GoString() string {
	return s.String()
  }
  
  func (s *OpenApiRequest) GetHeaders() map[string]*string  {
	return s.Headers
  }
  
  func (s *OpenApiRequest) GetQuery() map[string]*string  {
	return s.Query
  }
  
  func (s *OpenApiRequest) GetBody() interface{}  {
	return s.Body
  }
  
  func (s *OpenApiRequest) GetStream() io.Reader  {
	return s.Stream
  }
  
  func (s *OpenApiRequest) GetHostMap() map[string]*string  {
	return s.HostMap
  }
  
  func (s *OpenApiRequest) GetEndpointOverride() *string  {
	return s.EndpointOverride
  }
  
  func (s *OpenApiRequest) SetHeaders(v map[string]*string) *OpenApiRequest {
	s.Headers = v
	return s
  }
  
  func (s *OpenApiRequest) SetQuery(v map[string]*string) *OpenApiRequest {
	s.Query = v
	return s
  }
  
  func (s *OpenApiRequest) SetBody(v interface{}) *OpenApiRequest {
	s.Body = v
	return s
  }
  
  func (s *OpenApiRequest) SetStream(v io.Reader) *OpenApiRequest {
	s.Stream = v
	return s
  }
  
  func (s *OpenApiRequest) SetHostMap(v map[string]*string) *OpenApiRequest {
	s.HostMap = v
	return s
  }
  
  func (s *OpenApiRequest) SetEndpointOverride(v string) *OpenApiRequest {
	s.EndpointOverride = &v
	return s
  }
  