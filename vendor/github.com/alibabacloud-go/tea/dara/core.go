package dara

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/alibabacloud-go/debug/debug"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/utils"

	"golang.org/x/net/proxy"
)

type RuntimeOptions = util.RuntimeOptions
type ExtendsParameters = util.ExtendsParameters

var debugLog = debug.Init("dara")
type HttpRequest interface {
}

type HttpResponse interface {
}

type HttpClient interface {
	Call(request *http.Request, transport *http.Transport) (response *http.Response, err error)
}

type daraClient struct {
	sync.Mutex
	httpClient *http.Client
	ifInit     bool
}

func (client *daraClient) Call(request *http.Request, transport *http.Transport) (response *http.Response, err error) {
	response, err = client.httpClient.Do(request)
	return
}

var hookDo = func(fn func(req *http.Request, transport *http.Transport) (*http.Response, error)) func(req *http.Request, transport *http.Transport) (*http.Response, error) {
	return fn
}

var basicTypes = []string{
	"int", "int16", "int64", "int32", "float32", "float64", "string", "bool", "uint64", "uint32", "uint16",
}

// Verify whether the parameters meet the requirements
var validateParams = []string{"require", "pattern", "maxLength", "minLength", "maximum", "minimum", "maxItems", "minItems"}

var clientPool = &sync.Map{}

// Request is used wrap http request
type Request struct {
	Protocol *string
	Port     *int
	Method   *string
	Pathname *string
	Domain   *string
	Headers  map[string]*string
	Query    map[string]*string
	Body     io.Reader
}

// Response is use d wrap http response
type Response struct {
	Body          io.ReadCloser
	StatusCode    *int
	StatusMessage *string
	Headers       map[string]*string
}

// RuntimeObject is used for converting http configuration
type RuntimeObject struct {
	IgnoreSSL         *bool                  `json:"ignoreSSL" xml:"ignoreSSL"`
	ReadTimeout       *int                   `json:"readTimeout" xml:"readTimeout"`
	ConnectTimeout    *int                   `json:"connectTimeout" xml:"connectTimeout"`
	LocalAddr         *string                `json:"localAddr" xml:"localAddr"`
	HttpProxy         *string                `json:"httpProxy" xml:"httpProxy"`
	HttpsProxy        *string                `json:"httpsProxy" xml:"httpsProxy"`
	NoProxy           *string                `json:"noProxy" xml:"noProxy"`
	MaxIdleConns      *int                   `json:"maxIdleConns" xml:"maxIdleConns"`
	Key               *string                `json:"key" xml:"key"`
	Cert              *string                `json:"cert" xml:"cert"`
	Ca                *string                `json:"ca" xml:"ca"`
	Socks5Proxy       *string                `json:"socks5Proxy" xml:"socks5Proxy"`
	Socks5NetWork     *string                `json:"socks5NetWork" xml:"socks5NetWork"`
	Listener          utils.ProgressListener `json:"listener" xml:"listener"`
	Tracker           *utils.ReaderTracker   `json:"tracker" xml:"tracker"`
	Logger            *utils.Logger          `json:"logger" xml:"logger"`
	RetryOptions      *RetryOptions          `json:"retryOptions" xml:"retryOptions"`
	ExtendsParameters *ExtendsParameters     `json:"extendsParameters,omitempty" xml:"extendsParameters,omitempty"`
	HttpClient
}

func (r *RuntimeObject) getClientTag(domain string) string {
	return strconv.FormatBool(BoolValue(r.IgnoreSSL)) + strconv.Itoa(IntValue(r.ReadTimeout)) +
		strconv.Itoa(IntValue(r.ConnectTimeout)) + StringValue(r.LocalAddr) + StringValue(r.HttpProxy) +
		StringValue(r.HttpsProxy) + StringValue(r.NoProxy) + StringValue(r.Socks5Proxy) + StringValue(r.Socks5NetWork) + domain
}

// NewRuntimeObject is used for shortly create runtime object
func NewRuntimeObject(runtime map[string]interface{}) *RuntimeObject {
	if runtime == nil {
		return &RuntimeObject{}
	}

	runtimeObject := &RuntimeObject{
		IgnoreSSL:      TransInterfaceToBool(runtime["ignoreSSL"]),
		ReadTimeout:    TransInterfaceToInt(runtime["readTimeout"]),
		ConnectTimeout: TransInterfaceToInt(runtime["connectTimeout"]),
		LocalAddr:      TransInterfaceToString(runtime["localAddr"]),
		HttpProxy:      TransInterfaceToString(runtime["httpProxy"]),
		HttpsProxy:     TransInterfaceToString(runtime["httpsProxy"]),
		NoProxy:        TransInterfaceToString(runtime["noProxy"]),
		MaxIdleConns:   TransInterfaceToInt(runtime["maxIdleConns"]),
		Socks5Proxy:    TransInterfaceToString(runtime["socks5Proxy"]),
		Socks5NetWork:  TransInterfaceToString(runtime["socks5NetWork"]),
		Key:            TransInterfaceToString(runtime["key"]),
		Cert:           TransInterfaceToString(runtime["cert"]),
		Ca:             TransInterfaceToString(runtime["ca"]),
	}
	if runtime["listener"] != nil {
		runtimeObject.Listener = runtime["listener"].(utils.ProgressListener)
	}
	if runtime["tracker"] != nil {
		runtimeObject.Tracker = runtime["tracker"].(*utils.ReaderTracker)
	}
	if runtime["logger"] != nil {
		runtimeObject.Logger = runtime["logger"].(*utils.Logger)
	}
	if runtime["httpClient"] != nil {
		runtimeObject.HttpClient = runtime["httpClient"].(HttpClient)
	}
	if runtime["retryOptions"] != nil {
		runtimeObject.RetryOptions = runtime["retryOptions"].(*RetryOptions)
	}
	return runtimeObject
}

// NewRequest is used shortly create Request
func NewRequest() (req *Request) {
	return &Request{
		Headers: map[string]*string{},
		Query:   map[string]*string{},
	}
}

// NewResponse is create response with http response
func NewResponse(httpResponse *http.Response) (res *Response) {
	res = &Response{}
	res.Body = httpResponse.Body
	res.Headers = make(map[string]*string)
	res.StatusCode = Int(httpResponse.StatusCode)
	res.StatusMessage = String(httpResponse.Status)
	return
}

// Convert is use convert map[string]interface object to struct
func Convert(in interface{}, out interface{}) error {
	byt, _ := json.Marshal(in)
	decoder := jsonParser.NewDecoder(bytes.NewReader(byt))
	decoder.UseNumber()
	err := decoder.Decode(&out)
	return err
}

// Recover is used to format error
func Recover(in interface{}) error {
	if in == nil {
		return nil
	}
	return errors.New(fmt.Sprint(in))
}

// ReadBody is used read response body
func (response *Response) ReadBody() (body []byte, err error) {
	var buffer [512]byte
	defer response.Body.Close()
	result := bytes.NewBuffer(nil)

	for {
		n, err := response.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}
	return result.Bytes(), nil
}

func getDaraClient(tag string) *daraClient {
	client, ok := clientPool.Load(tag)
	if client == nil && !ok {
		client = &daraClient{
			httpClient: &http.Client{},
			ifInit:     false,
		}
		clientPool.Store(tag, client)
	}
	return client.(*daraClient)
}

// DoRequest is used send request to server
func DoRequest(request *Request, runtimeObject *RuntimeObject) (response *Response, err error) {
	if runtimeObject == nil {
		runtimeObject = &RuntimeObject{}
	}
	fieldMap := make(map[string]string)
	utils.InitLogMsg(fieldMap)
	defer func() {
		if runtimeObject.Logger != nil {
			runtimeObject.Logger.PrintLog(fieldMap, err)
		}
	}()
	if request.Method == nil {
		request.Method = String("GET")
	}

	if request.Protocol == nil {
		request.Protocol = String("http")
	} else {
		request.Protocol = String(strings.ToLower(StringValue(request.Protocol)))
	}

	requestURL := ""
	request.Domain = request.Headers["host"]
	if request.Port != nil {
		request.Domain = String(fmt.Sprintf("%s:%d", StringValue(request.Domain), IntValue(request.Port)))
	}
	requestURL = fmt.Sprintf("%s://%s%s", StringValue(request.Protocol), StringValue(request.Domain), StringValue(request.Pathname))
	queryParams := request.Query
	// sort QueryParams by key
	q := url.Values{}
	for key, value := range queryParams {
		q.Add(key, StringValue(value))
	}
	querystring := q.Encode()
	if len(querystring) > 0 {
		if strings.Contains(requestURL, "?") {
			requestURL = fmt.Sprintf("%s&%s", requestURL, querystring)
		} else {
			requestURL = fmt.Sprintf("%s?%s", requestURL, querystring)
		}
	}
	debugLog("> %s %s", StringValue(request.Method), requestURL)

	httpRequest, err := http.NewRequest(StringValue(request.Method), requestURL, request.Body)
	if err != nil {
		return
	}
	httpRequest.Host = StringValue(request.Domain)

	var client HttpClient
	if runtimeObject.HttpClient == nil {
		client = getDaraClient(runtimeObject.getClientTag(StringValue(request.Domain)))
	} else {
		client = runtimeObject.HttpClient
	}

	trans, err := getHttpTransport(request, runtimeObject)
	if err != nil {
		return
	}
	if defaultClient, ok := client.(*daraClient); ok {
		defaultClient.Lock()
		if !defaultClient.ifInit || defaultClient.httpClient.Transport == nil {
			defaultClient.httpClient.Transport = trans
		}
		defaultClient.httpClient.Timeout = time.Duration(IntValue(runtimeObject.ReadTimeout)) * time.Millisecond
		defaultClient.ifInit = true
		defaultClient.Unlock()
	}

	for key, value := range request.Headers {
		if value == nil || key == "content-length" {
			continue
		} else if key == "host" {
			httpRequest.Header["Host"] = []string{*value}
			delete(httpRequest.Header, "host")
		} else if key == "user-agent" {
			httpRequest.Header["User-Agent"] = []string{*value}
			delete(httpRequest.Header, "user-agent")
		} else {
			httpRequest.Header[key] = []string{*value}
		}
		debugLog("> %s: %s", key, StringValue(value))
	}
	contentlength, _ := strconv.Atoi(StringValue(request.Headers["content-length"]))
	event := utils.NewProgressEvent(utils.TransferStartedEvent, 0, int64(contentlength), 0)
	utils.PublishProgress(runtimeObject.Listener, event)

	putMsgToMap(fieldMap, httpRequest)
	startTime := time.Now()
	fieldMap["{start_time}"] = startTime.Format("2006-01-02 15:04:05")
	res, err := hookDo(client.Call)(httpRequest, trans)
	fieldMap["{cost}"] = time.Since(startTime).String()
	completedBytes := int64(0)
	if runtimeObject.Tracker != nil {
		completedBytes = runtimeObject.Tracker.CompletedBytes
	}
	if err != nil {
		event = utils.NewProgressEvent(utils.TransferFailedEvent, completedBytes, int64(contentlength), 0)
		utils.PublishProgress(runtimeObject.Listener, event)
		return
	}

	event = utils.NewProgressEvent(utils.TransferCompletedEvent, completedBytes, int64(contentlength), 0)
	utils.PublishProgress(runtimeObject.Listener, event)

	response = NewResponse(res)
	fieldMap["{code}"] = strconv.Itoa(res.StatusCode)
	fieldMap["{res_headers}"] = Stringify(res.Header)
	debugLog("< HTTP/1.1 %s", res.Status)
	for key, value := range res.Header {
		debugLog("< %s: %s", key, strings.Join(value, ""))
		if len(value) != 0 {
			response.Headers[strings.ToLower(key)] = String(value[0])
		}
	}
	return
}

func getHttpTransport(req *Request, runtime *RuntimeObject) (*http.Transport, error) {
	trans := new(http.Transport)
	httpProxy, err := getHttpProxy(StringValue(req.Protocol), StringValue(req.Domain), runtime)
	if err != nil {
		return nil, err
	}
	if strings.ToLower(*req.Protocol) == "https" {
		if BoolValue(runtime.IgnoreSSL) != true {
			trans.TLSClientConfig = &tls.Config{
				InsecureSkipVerify: false,
			}
			if runtime.Key != nil && runtime.Cert != nil && StringValue(runtime.Key) != "" && StringValue(runtime.Cert) != "" {
				cert, err := tls.X509KeyPair([]byte(StringValue(runtime.Cert)), []byte(StringValue(runtime.Key)))
				if err != nil {
					return nil, err
				}
				trans.TLSClientConfig.Certificates = []tls.Certificate{cert}
			}
			if runtime.Ca != nil && StringValue(runtime.Ca) != "" {
				clientCertPool := x509.NewCertPool()
				ok := clientCertPool.AppendCertsFromPEM([]byte(StringValue(runtime.Ca)))
				if !ok {
					return nil, errors.New("Failed to parse root certificate")
				}
				trans.TLSClientConfig.RootCAs = clientCertPool
			}
		} else {
			trans.TLSClientConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
		}
	}
	if httpProxy != nil {
		trans.Proxy = http.ProxyURL(httpProxy)
		if httpProxy.User != nil {
			password, _ := httpProxy.User.Password()
			auth := httpProxy.User.Username() + ":" + password
			basic := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
			req.Headers["Proxy-Authorization"] = String(basic)
		}
	}
	if runtime.Socks5Proxy != nil && StringValue(runtime.Socks5Proxy) != "" {
		socks5Proxy, err := getSocks5Proxy(runtime)
		if err != nil {
			return nil, err
		}
		if socks5Proxy != nil {
			var auth *proxy.Auth
			if socks5Proxy.User != nil {
				password, _ := socks5Proxy.User.Password()
				auth = &proxy.Auth{
					User:     socks5Proxy.User.Username(),
					Password: password,
				}
			}
			dialer, err := proxy.SOCKS5(strings.ToLower(StringValue(runtime.Socks5NetWork)), socks5Proxy.String(), auth,
				&net.Dialer{
					Timeout:   time.Duration(IntValue(runtime.ConnectTimeout)) * time.Millisecond,
					DualStack: true,
					LocalAddr: getLocalAddr(StringValue(runtime.LocalAddr)),
				})
			if err != nil {
				return nil, err
			}
			trans.Dial = dialer.Dial
		}
	} else {
		trans.DialContext = setDialContext(runtime)
	}
	if runtime.MaxIdleConns != nil && *runtime.MaxIdleConns > 0 {
		trans.MaxIdleConns = IntValue(runtime.MaxIdleConns)
		trans.MaxIdleConnsPerHost = IntValue(runtime.MaxIdleConns)
	}
	return trans, nil
}

func putMsgToMap(fieldMap map[string]string, request *http.Request) {
	fieldMap["{host}"] = request.Host
	fieldMap["{method}"] = request.Method
	fieldMap["{uri}"] = request.URL.RequestURI()
	fieldMap["{pid}"] = strconv.Itoa(os.Getpid())
	fieldMap["{version}"] = strings.Split(request.Proto, "/")[1]
	hostname, _ := os.Hostname()
	fieldMap["{hostname}"] = hostname
	fieldMap["{req_headers}"] = Stringify(request.Header)
	fieldMap["{target}"] = request.URL.Path + request.URL.RawQuery
}

func getNoProxy(protocol string, runtime *RuntimeObject) []string {
	var urls []string
	if runtime.NoProxy != nil && StringValue(runtime.NoProxy) != "" {
		urls = strings.Split(StringValue(runtime.NoProxy), ",")
	} else if rawurl := os.Getenv("NO_PROXY"); rawurl != "" {
		urls = strings.Split(rawurl, ",")
	} else if rawurl := os.Getenv("no_proxy"); rawurl != "" {
		urls = strings.Split(rawurl, ",")
	}

	return urls
}

func ToReader(obj interface{}) io.Reader {
	switch obj.(type) {
	case string:
		tmp := obj.(string)
		return strings.NewReader(tmp)
	case *string:
		tmp := obj.(*string)
		return strings.NewReader(StringValue(tmp))
	case []byte:
		return strings.NewReader(string(obj.([]byte)))
	case io.Reader:
		return obj.(io.Reader)
	default:
		panic("Invalid Body. Please set a valid Body.")
	}
}

func ToWriter(obj interface{}) io.Writer {
	switch obj.(type) {
	case string:
		var buf bytes.Buffer
		buf.WriteString(obj.(string))
		return &buf
	case *string:
		var buf bytes.Buffer
		tmp := obj.(*string)
		buf.WriteString(*tmp)
		return &buf
	case []byte:
		var buf bytes.Buffer
		buf.Write(obj.([]byte))
		return &buf
	case io.Writer:
		return obj.(io.Writer)
	case *bytes.Buffer:
		return obj.(*bytes.Buffer)
	case *os.File:
		return obj.(*os.File)
	default:
		panic("Invalid Writer. Please provide a valid Writer.")
	}
}

func ToString(val interface{}) string {
	switch v := val.(type) {
	case []byte:
		return string(v) // 将 []byte 转换为字符串
	default:
		return fmt.Sprintf("%v", v) // 处理其他类型
	}
}

func getHttpProxy(protocol, host string, runtime *RuntimeObject) (proxy *url.URL, err error) {
	urls := getNoProxy(protocol, runtime)
	for _, url := range urls {
		if url == host {
			return nil, nil
		}
	}
	if protocol == "https" {
		if runtime.HttpsProxy != nil && StringValue(runtime.HttpsProxy) != "" {
			proxy, err = url.Parse(StringValue(runtime.HttpsProxy))
		} else if rawurl := os.Getenv("HTTPS_PROXY"); rawurl != "" {
			proxy, err = url.Parse(rawurl)
		} else if rawurl := os.Getenv("https_proxy"); rawurl != "" {
			proxy, err = url.Parse(rawurl)
		}
	} else {
		if runtime.HttpProxy != nil && StringValue(runtime.HttpProxy) != "" {
			proxy, err = url.Parse(StringValue(runtime.HttpProxy))
		} else if rawurl := os.Getenv("HTTP_PROXY"); rawurl != "" {
			proxy, err = url.Parse(rawurl)
		} else if rawurl := os.Getenv("http_proxy"); rawurl != "" {
			proxy, err = url.Parse(rawurl)
		}
	}

	return proxy, err
}

func getSocks5Proxy(runtime *RuntimeObject) (proxy *url.URL, err error) {
	if runtime.Socks5Proxy != nil && StringValue(runtime.Socks5Proxy) != "" {
		proxy, err = url.Parse(StringValue(runtime.Socks5Proxy))
	}
	return proxy, err
}

func getLocalAddr(localAddr string) (addr *net.TCPAddr) {
	if localAddr != "" {
		addr = &net.TCPAddr{
			IP: []byte(localAddr),
		}
	}
	return addr
}

func setDialContext(runtime *RuntimeObject) func(cxt context.Context, net, addr string) (c net.Conn, err error) {
	return func(ctx context.Context, network, address string) (net.Conn, error) {
		if runtime.LocalAddr != nil && StringValue(runtime.LocalAddr) != "" {
			netAddr := &net.TCPAddr{
				IP: []byte(StringValue(runtime.LocalAddr)),
			}
			return (&net.Dialer{
				Timeout:   time.Duration(IntValue(runtime.ConnectTimeout)) * time.Second,
				DualStack: true,
				LocalAddr: netAddr,
			}).DialContext(ctx, network, address)
		}
		return (&net.Dialer{
			Timeout:   time.Duration(IntValue(runtime.ConnectTimeout)) * time.Second,
			DualStack: true,
		}).DialContext(ctx, network, address)
	}
}

func ToObject(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	byt, _ := json.Marshal(obj)
	err := json.Unmarshal(byt, &result)
	if err != nil {
		return nil
	}
	return result
}

func AllowRetry(retry interface{}, retryTimes *int) *bool {
	if IntValue(retryTimes) == 0 {
		return Bool(true)
	}
	retryMap, ok := retry.(map[string]interface{})
	if !ok {
		return Bool(false)
	}
	retryable, ok := retryMap["retryable"].(bool)
	if !ok || !retryable {
		return Bool(false)
	}

	maxAttempts, ok := retryMap["maxAttempts"].(int)
	if !ok || maxAttempts < IntValue(retryTimes) {
		return Bool(false)
	}
	return Bool(true)
}

func Merge(args ...interface{}) map[string]*string {
	finalArg := make(map[string]*string)
	for _, obj := range args {
		switch obj.(type) {
		case map[string]*string:
			arg := obj.(map[string]*string)
			for key, value := range arg {
				if value != nil {
					finalArg[key] = value
				}
			}
		default:
			byt, _ := json.Marshal(obj)
			arg := make(map[string]string)
			err := json.Unmarshal(byt, &arg)
			if err != nil {
				return finalArg
			}
			for key, value := range arg {
				if value != "" {
					finalArg[key] = String(value)
				}
			}
		}
	}

	return finalArg
}

func IsNil(val interface{}) bool {
	defer func() {
		recover()
	}()
	if val == nil {
		return true
	}

	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
		return v.IsNil()
	}

	valType := reflect.TypeOf(val)
	valZero := reflect.Zero(valType)
	return valZero == v
}

func isNil(a interface{}) bool {
	defer func() {
		recover()
	}()
	vi := reflect.ValueOf(a)
	return vi.IsNil()
}

func isNilOrZero(value interface{}) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		return v.IsNil()
	default:
		// Check for zero value
		return reflect.DeepEqual(value, reflect.Zero(v.Type()).Interface())
	}
}

func Default(inputValue interface{}, defaultValue interface{}) (_result interface{}) {
	if isNilOrZero(inputValue) {
		_result = defaultValue
		return _result
	}

	_result = inputValue
	return _result
}

func ToMap(args ...interface{}) map[string]interface{} {
	isNotNil := false
	finalArg := make(map[string]interface{})
	for _, obj := range args {
		if obj == nil {
			continue
		}

		if isNil(obj) {
			continue
		}
		isNotNil = true

		switch obj.(type) {
		case map[string]*string:
			arg := obj.(map[string]*string)
			for key, value := range arg {
				if value != nil {
					finalArg[key] = StringValue(value)
				}
			}
		case map[string]interface{}:
			arg := obj.(map[string]interface{})
			for key, value := range arg {
				if value != nil {
					finalArg[key] = value
				}
			}
		case *string:
			str := obj.(*string)
			arg := make(map[string]interface{})
			err := json.Unmarshal([]byte(StringValue(str)), &arg)
			if err == nil {
				for key, value := range arg {
					if value != nil {
						finalArg[key] = value
					}
				}
			}
			tmp := make(map[string]string)
			err = json.Unmarshal([]byte(StringValue(str)), &tmp)
			if err == nil {
				for key, value := range arg {
					if value != "" {
						finalArg[key] = value
					}
				}
			}
		case []byte:
			byt := obj.([]byte)
			arg := make(map[string]interface{})
			err := json.Unmarshal(byt, &arg)
			if err == nil {
				for key, value := range arg {
					if value != nil {
						finalArg[key] = value
					}
				}
				break
			}
		default:
			val := reflect.ValueOf(obj)
			if val.Kind().String() == "map" {
				tmp := val.MapKeys()
				for _, key := range tmp {
					finalArg[key.String()] = val.MapIndex(key).Interface()
				}
			} else {
				res := structToMap(val)
				for key, value := range res {
					if value != nil {
						finalArg[key] = value
					}
				}
			}
		}
	}

	if !isNotNil {
		return nil
	}
	return finalArg
}

func structToMap(dataValue reflect.Value) map[string]interface{} {
	out := make(map[string]interface{})
	if !dataValue.IsValid() {
		return out
	}
	if dataValue.Kind().String() == "ptr" {
		if dataValue.IsNil() {
			return out
		}
		dataValue = dataValue.Elem()
	}
	if !dataValue.IsValid() {
		return out
	}
	dataType := dataValue.Type()
	if dataType.Kind().String() != "struct" {
		return out
	}
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		name, containsNameTag := field.Tag.Lookup("json")
		if !containsNameTag {
			name = field.Name
		} else {
			strs := strings.Split(name, ",")
			name = strs[0]
		}
		fieldValue := dataValue.FieldByName(field.Name)
		if !fieldValue.IsValid() || fieldValue.IsNil() {
			continue
		}
		if field.Type.String() == "io.Reader" || field.Type.String() == "io.Writer" {
			continue
		} else if field.Type.Kind().String() == "struct" {
			out[name] = structToMap(fieldValue)
		} else if field.Type.Kind().String() == "ptr" &&
			field.Type.Elem().Kind().String() == "struct" {
			if fieldValue.Elem().IsValid() {
				out[name] = structToMap(fieldValue)
			}
		} else if field.Type.Kind().String() == "ptr" {
			if fieldValue.IsValid() && !fieldValue.IsNil() {
				out[name] = fieldValue.Elem().Interface()
			}
		} else if field.Type.Kind().String() == "slice" {
			tmp := make([]interface{}, 0)
			num := fieldValue.Len()
			for i := 0; i < num; i++ {
				value := fieldValue.Index(i)
				if !value.IsValid() {
					continue
				}
				if value.Type().Kind().String() == "ptr" &&
					value.Type().Elem().Kind().String() == "struct" {
					if value.IsValid() && !value.IsNil() {
						tmp = append(tmp, structToMap(value))
					}
				} else if value.Type().Kind().String() == "struct" {
					tmp = append(tmp, structToMap(value))
				} else if value.Type().Kind().String() == "ptr" {
					if value.IsValid() && !value.IsNil() {
						tmp = append(tmp, value.Elem().Interface())
					}
				} else {
					tmp = append(tmp, value.Interface())
				}
			}
			if len(tmp) > 0 {
				out[name] = tmp
			}
		} else {
			out[name] = fieldValue.Interface()
		}

	}
	return out
}

func GetBackoffTime(backoff interface{}, retrytimes *int) *int {
	backoffMap, ok := backoff.(map[string]interface{})
	if !ok {
		return Int(0)
	}
	policy, ok := backoffMap["policy"].(string)
	if !ok || policy == "no" {
		return Int(0)
	}

	period, ok := backoffMap["period"].(int)
	if !ok || period == 0 {
		return Int(0)
	}

	maxTime := math.Pow(2.0, float64(IntValue(retrytimes)))
	return Int(rand.Intn(int(maxTime-1)) * period)
}

func Sleep(backoffTime int) {
	sleeptime := time.Duration(backoffTime) * time.Second
	time.Sleep(sleeptime)
}

// Determines whether realType is in filterTypes
func isFilterType(realType string, filterTypes []string) bool {
	for _, value := range filterTypes {
		if value == realType {
			return true
		}
	}
	return false
}

func TransInterfaceToBool(val interface{}) *bool {
	if val == nil {
		return nil
	}

	return Bool(val.(bool))
}

func TransInterfaceToInt(val interface{}) *int {
	if val == nil {
		return nil
	}

	return Int(val.(int))
}

func TransInterfaceToString(val interface{}) *string {
	if val == nil {
		return nil
	}

	return String(val.(string))
}

func Prettify(i interface{}) string {
	resp, _ := json.MarshalIndent(i, "", "   ")
	return string(resp)
}

func ToInt(a *int32) *int {
	return Int(int(Int32Value(a)))
}

func ToInt32(a *int) *int32 {
	return Int32(int32(IntValue(a)))
}

func ToBytes(s, encodingType string) []byte {
	switch encodingType {
	case "utf8":
		return []byte(s)
	case "base64":
		data, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			return nil
		}
		return data
	case "hex":
		data, err := hex.DecodeString(s)
		if err != nil {
			return nil
		}
		return data
	default:
		return nil
	}
}

func BytesFromString(str string, typeStr string) []byte {
	switch typeStr {
	case "utf8":
		return []byte(str)
	case "hex":
		bytes, err := hex.DecodeString(str)
		if err == nil {
			return bytes
		}
	case "base64":
		bytes, err := base64.StdEncoding.DecodeString(str)
		if err == nil {
			return bytes
		}
	}

	// 对于不支持的类型或解码失败，返回 nil
	return nil
}

func ForceInt(a interface{}) int {
	num, _ := a.(int)
	return num
}

func ForceBoolean(a interface{}) bool {
	b, _ := a.(bool)
	return b
}

func ForceInt64(a interface{}) int64 {
	b, _ := a.(int64)
	return b
}

func ForceUint64(a interface{}) uint64 {
	b, _ := a.(uint64)
	return b
}

// ForceInt32 attempts to assert that a is of type int32.
func ForceInt32(a interface{}) int32 {
	b, _ := a.(int32)
	return b
}

// ForceUInt32 attempts to assert that a is of type uint32.
func ForceUInt32(a interface{}) uint32 {
	b, _ := a.(uint32)
	return b
}

// ForceInt16 attempts to assert that a is of type int16.
func ForceInt16(a interface{}) int16 {
	b, _ := a.(int16)
	return b
}

// ForceUInt16 attempts to assert that a is of type uint16.
func ForceUInt16(a interface{}) uint16 {
	b, _ := a.(uint16)
	return b
}

// ForceInt8 attempts to assert that a is of type int8.
func ForceInt8(a interface{}) int8 {
	b, _ := a.(int8)
	return b
}

// ForceUInt8 attempts to assert that a is of type uint8.
func ForceUInt8(a interface{}) uint8 {
	b, _ := a.(uint8)
	return b
}

// ForceFloat32 attempts to assert that a is of type float32.
func ForceFloat32(a interface{}) float32 {
	b, _ := a.(float32)
	return b
}

// ForceFloat64 attempts to assert that a is of type float64.
func ForceFloat64(a interface{}) float64 {
	b, _ := a.(float64)
	return b
}
