// This file is auto-generated, don't edit it. Thanks.
package utils

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	mathRand "math/rand"
	"net/http"
	"net/textproto"
	"net/url"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	models "github.com/alibabacloud-go/darabonba-openapi/v2/models"
	"github.com/alibabacloud-go/tea/dara"
	"github.com/tjfoc/gmsm/sm3"
)

var defaultUserAgent = fmt.Sprintf("AlibabaCloud (%s; %s) Golang/%s Core/%s TeaDSL/2", runtime.GOOS, runtime.GOARCH, strings.Trim(runtime.Version(), "go"), "0.01")
var seqId int64 = 0
var processStartTime int64 = time.Now().UnixNano() / 1e6

const (
	PEM_BEGIN = "-----BEGIN RSA PRIVATE KEY-----\n"
	PEM_END   = "\n-----END RSA PRIVATE KEY-----"
)

type Config = models.Config
type GlobalParameters = models.GlobalParameters
type Params = models.Params
type OpenApiRequest = models.OpenApiRequest

type Sorter struct {
	Keys []string
	Vals []string
}

func newSorter(m map[string]string) *Sorter {
	hs := &Sorter{
		Keys: make([]string, 0, len(m)),
		Vals: make([]string, 0, len(m)),
	}

	for k, v := range m {
		hs.Keys = append(hs.Keys, k)
		hs.Vals = append(hs.Vals, v)
	}
	return hs
}

// Sort is an additional function for function SignHeader.
func (hs *Sorter) Sort() {
	sort.Sort(hs)
}

// Len is an additional function for function SignHeader.
func (hs *Sorter) Len() int {
	return len(hs.Vals)
}

// Less is an additional function for function SignHeader.
func (hs *Sorter) Less(i, j int) bool {
	return bytes.Compare([]byte(hs.Keys[i]), []byte(hs.Keys[j])) < 0
}

// Swap is an additional function for function SignHeader.
func (hs *Sorter) Swap(i, j int) {
	hs.Vals[i], hs.Vals[j] = hs.Vals[j], hs.Vals[i]
	hs.Keys[i], hs.Keys[j] = hs.Keys[j], hs.Keys[i]
}

// Description:
//
// This is for OpenApi Util

// Description:
//
// # Convert all params of body other than type of readable into content
//
// @param body - source Model
//
// @param content - target Model
//
// @return void
func Convert(body interface{}, content interface{}) {
	res := make(map[string]interface{})
	val := reflect.ValueOf(body).Elem()
	dataType := val.Type()
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		name, _ := field.Tag.Lookup("json")
		name = strings.Split(name, ",omitempty")[0]
		_, ok := val.Field(i).Interface().(io.Reader)
		if !ok {
			res[name] = val.Field(i).Interface()
		}
	}
	byt, _ := json.Marshal(res)
	json.Unmarshal(byt, content)
}

// Description:
//
// # Get throttling param
//
// @param the - response headers
//
// @return time left
func GetThrottlingTimeLeft(headers map[string]*string) (_result *int64) {
	rateLimitForUserApi := headers["x-ratelimit-user-api"]
	rateLimitForUser := headers["x-ratelimit-user"]
	timeLeftForUserApi := getTimeLeft(rateLimitForUserApi)
	timeLeftForUser := getTimeLeft(rateLimitForUser)
	if dara.Int64Value(timeLeftForUserApi) > dara.Int64Value(timeLeftForUser) {
		return timeLeftForUserApi
	} else {
		return timeLeftForUser
	}
}

// Description:
//
// # Hash the raw data with signatureAlgorithm
//
// @param raw - hashing data
//
// @param signatureAlgorithm - the autograph method
//
// @return hashed bytes
func Hash(raw []byte, signatureAlgorithm *string) (_result []byte) {
	signType := dara.StringValue(signatureAlgorithm)
	if signType == "ACS3-HMAC-SHA256" || signType == "ACS3-RSA-SHA256" {
		h := sha256.New()
		h.Write(raw)
		return h.Sum(nil)
	} else if signType == "ACS3-HMAC-SM3" {
		h := sm3.New()
		h.Write(raw)
		return h.Sum(nil)
	}
	return nil
}

func getGID() uint64 {
	// https://blog.sgmansfield.com/2015/12/goroutine-ids/
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

// Description:
//
// # Generate a nonce string
//
// @return the nonce string
func GetNonce() (_result *string) {
	routineId := getGID()
	currentTime := time.Now().UnixNano() / 1e6
	seq := atomic.AddInt64(&seqId, 1)
	randNum := mathRand.Int63()
	msg := fmt.Sprintf("%d-%d-%d-%d-%d", processStartTime, routineId, currentTime, seq, randNum)
	h := md5.New()
	h.Write([]byte(msg))
	ret := hex.EncodeToString(h.Sum(nil))
	return &ret
}

// Description:
//
// # Get the string to be signed according to request
//
// @param request - which contains signed messages
//
// @return the signed string
func GetStringToSign(request *dara.Request) (_result *string) {
	return dara.String(getStringToSign(request))
}

func getStringToSign(request *dara.Request) string {
	// sort QueryParams by key
	var queryKeys []string
	resource := dara.StringValue(request.Pathname)
	queryParams := request.Query
	for key := range queryParams {
		queryKeys = append(queryKeys, key)
	}
	sort.Strings(queryKeys)
	tmp := ""
	for i := 0; i < len(queryKeys); i++ {
		queryKey := queryKeys[i]
		v := dara.StringValue(queryParams[queryKey])
		if v != "" {
			tmp = tmp + "&" + queryKey + "=" + v
		} else {
			tmp = tmp + "&" + queryKey
		}
	}
	if tmp != "" {
		tmp = strings.TrimLeft(tmp, "&")
		resource = resource + "?" + tmp
	}
	return getSignedStr(request, resource)
}

func getSignedStr(req *dara.Request, canonicalizedResource string) string {
	temp := make(map[string]string)

	for k, v := range req.Headers {
		if strings.HasPrefix(strings.ToLower(k), "x-acs-") {
			temp[strings.ToLower(k)] = dara.StringValue(v)
		}
	}
	hs := newSorter(temp)

	// Sort the temp by the ascending order
	hs.Sort()

	// Get the canonicalizedOSSHeaders
	canonicalizedOSSHeaders := ""
	for i := range hs.Keys {
		canonicalizedOSSHeaders += hs.Keys[i] + ":" + hs.Vals[i] + "\n"
	}

	// Give other parameters values
	// when sign URL, date is expires
	date := dara.StringValue(req.Headers["date"])
	accept := dara.StringValue(req.Headers["accept"])
	contentType := dara.StringValue(req.Headers["content-type"])
	contentMd5 := dara.StringValue(req.Headers["content-md5"])

	signStr := dara.StringValue(req.Method) + "\n" + accept + "\n" + contentMd5 + "\n" + contentType + "\n" + date + "\n" + canonicalizedOSSHeaders + canonicalizedResource
	return signStr
}

// Description:
//
// # Get signature according to stringToSign, secret
//
// @param stringToSign - the signed string
//
// @param secret - accesskey secret
//
// @return the signature
func GetROASignature(stringToSign *string, secret *string) (_result *string) {
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(dara.StringValue(secret)))
	io.WriteString(h, dara.StringValue(stringToSign))
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return dara.String(signedStr)
}

// Description:
//
// # Parse filter into a form string
//
// @param filter - object
//
// @return the string
func ToForm(filter map[string]interface{}) (_result *string) {
	tmp := make(map[string]interface{})
	byt, _ := json.Marshal(filter)
	d := json.NewDecoder(bytes.NewReader(byt))
	d.UseNumber()
	_ = d.Decode(&tmp)

	result := make(map[string]*string)
	for key, value := range tmp {
		filterValue := reflect.ValueOf(value)
		flatRepeatedList(filterValue, result, key)
	}

	m := make(map[string]interface{})
	for key, value := range result {
		m[key] = dara.StringValue(value)
	}
	return dara.String(dara.ToFormString(m))
}

func flatRepeatedList(dataValue reflect.Value, result map[string]*string, prefix string) {
	if !dataValue.IsValid() {
		return
	}

	dataType := dataValue.Type()
	if dataType.Kind().String() == "slice" {
		handleRepeatedParams(dataValue, result, prefix)
	} else if dataType.Kind().String() == "map" {
		handleMap(dataValue, result, prefix)
	} else {
		result[prefix] = dara.String(fmt.Sprintf("%v", dataValue.Interface()))
	}
}

func handleRepeatedParams(repeatedFieldValue reflect.Value, result map[string]*string, prefix string) {
	if repeatedFieldValue.IsValid() && !repeatedFieldValue.IsNil() {
		for m := 0; m < repeatedFieldValue.Len(); m++ {
			elementValue := repeatedFieldValue.Index(m)
			key := prefix + "." + strconv.Itoa(m+1)
			fieldValue := reflect.ValueOf(elementValue.Interface())
			if fieldValue.Kind().String() == "map" {
				handleMap(fieldValue, result, key)
			} else {
				result[key] = dara.String(fmt.Sprintf("%v", fieldValue.Interface()))
			}
		}
	}
}

func handleMap(valueField reflect.Value, result map[string]*string, prefix string) {
	var byt []byte
	if valueField.IsValid() && valueField.String() != "" {
		valueFieldType := valueField.Type()
		if valueFieldType.Kind().String() == "map" {
			byt, _ = json.Marshal(valueField.Interface())
			cache := make(map[string]interface{})
			d := json.NewDecoder(bytes.NewReader(byt))
			d.UseNumber()
			_ = d.Decode(&cache)
			for key, value := range cache {
				pre := ""
				if prefix != "" {
					pre = prefix + "." + key
				} else {
					pre = key
				}
				fieldValue := reflect.ValueOf(value)
				flatRepeatedList(fieldValue, result, pre)
			}
		}
	}
}

// Description:
//
// # Get timestamp
//
// @return the timestamp string
func GetTimestamp() (_result *string) {
	gmt := time.FixedZone("GMT", 0)
	return dara.String(time.Now().In(gmt).Format("2006-01-02T15:04:05Z"))
}

// Description:
//
// # Get UTC string
//
// @return the UTC string
func GetDateUTCString() (_result *string) {
	return dara.String(time.Now().UTC().Format(http.TimeFormat))
}

// Description:
//
// Parse filter into a object which's type is map[string]string
//
// @param filter - query param
//
// @return the object
func Query(filter interface{}) (_result map[string]*string) {
	tmp := make(map[string]interface{})
	byt, _ := json.Marshal(filter)
	d := json.NewDecoder(bytes.NewReader(byt))
	d.UseNumber()
	_ = d.Decode(&tmp)

	result := make(map[string]*string)
	for key, value := range tmp {
		filterValue := reflect.ValueOf(value)
		flatRepeatedList(filterValue, result, key)
	}

	return result
}

// Description:
//
// # Get signature according to signedParams, method and secret
//
// @param signedParams - params which need to be signed
//
// @param method - http method e.g. GET
//
// @param secret - AccessKeySecret
//
// @return the signature
func GetRPCSignature(signedParams map[string]*string, method *string, secret *string) (_result *string) {
	stringToSign := buildRpcStringToSign(signedParams, dara.StringValue(method))
	signature := sign(stringToSign, dara.StringValue(secret), "&")
	return dara.String(signature)
}

// Description:
//
// # Parse array into a string with specified style
//
// @param array - the array
//
// @param prefix - the prefix string
//
// @return the string
func ArrayToStringWithSpecifiedStyle(array interface{}, prefix *string, style *string) (_result *string) {
	if dara.IsNil(array) {
		return dara.String("")
	}

	sty := dara.StringValue(style)
	if sty == "repeatList" {
		tmp := map[string]interface{}{
			dara.StringValue(prefix): array,
		}
		return flatRepeatList(tmp)
	} else if sty == "simple" || sty == "spaceDelimited" || sty == "pipeDelimited" {
		return flatArray(array, sty)
	} else if sty == "json" {
		return dara.String(dara.Stringify(array))
	}
	return dara.String("")
}

// Description:
//
// # Get the authorization
//
// @param request - request params
//
// @param signatureAlgorithm - the autograph method
//
// @param payload - the hashed request
//
// @param accessKey - the accessKey string
//
// @param accessKeySecret - the accessKeySecret string
//
// @return authorization string
func GetAuthorization(request *dara.Request, signatureAlgorithm *string, payload *string, accessKey *string, accessKeySecret *string) (_result *string) {
	canonicalURI := dara.StringValue(request.Pathname)
	if canonicalURI == "" {
		canonicalURI = "/"
	}

	canonicalURI = strings.Replace(canonicalURI, "+", "%20", -1)
	canonicalURI = strings.Replace(canonicalURI, "*", "%2A", -1)
	canonicalURI = strings.Replace(canonicalURI, "%7E", "~", -1)

	method := dara.StringValue(request.Method)
	canonicalQueryString := getCanonicalQueryString(request.Query)
	canonicalheaders, signedHeaders := getCanonicalHeaders(request.Headers)

	canonicalRequest := method + "\n" + canonicalURI + "\n" + canonicalQueryString + "\n" + canonicalheaders + "\n" +
		strings.Join(signedHeaders, ";") + "\n" + dara.StringValue(payload)
	signType := dara.StringValue(signatureAlgorithm)
	StringToSign := signType + "\n" + hex.EncodeToString(Hash([]byte(canonicalRequest), signatureAlgorithm))
	signature := hex.EncodeToString(SignatureMethod(dara.StringValue(accessKeySecret), StringToSign, signType))
	auth := signType + " Credential=" + dara.StringValue(accessKey) + ",SignedHeaders=" +
		strings.Join(signedHeaders, ";") + ",Signature=" + signature
	return dara.String(auth)
}

func GetUserAgent(userAgent *string) (_result *string) {
	if userAgent != nil && dara.StringValue(userAgent) != "" {
		return dara.String(defaultUserAgent + " " + dara.StringValue(userAgent))
	}
	return dara.String(defaultUserAgent)
}

func SignatureMethod(secret, source, signatureAlgorithm string) []byte {
	if signatureAlgorithm == "ACS3-HMAC-SHA256" {
		h := hmac.New(sha256.New, []byte(secret))
		h.Write([]byte(source))
		return h.Sum(nil)
	} else if signatureAlgorithm == "ACS3-HMAC-SM3" {
		h := hmac.New(sm3.New, []byte(secret))
		h.Write([]byte(source))
		return h.Sum(nil)
	} else if signatureAlgorithm == "ACS3-RSA-SHA256" {
		return rsaSign(source, secret)
	}
	return nil
}

func flatRepeatList(filter map[string]interface{}) (_result *string) {
	tmp := make(map[string]interface{})
	byt, _ := json.Marshal(filter)
	d := json.NewDecoder(bytes.NewReader(byt))
	d.UseNumber()
	_ = d.Decode(&tmp)

	result := make(map[string]*string)
	for key, value := range tmp {
		filterValue := reflect.ValueOf(value)
		flatRepeatedList(filterValue, result, key)
	}

	res := make(map[string]string)
	for k, v := range result {
		res[k] = dara.StringValue(v)
	}
	hs := newSorter(res)

	hs.Sort()

	// Get the canonicalizedOSSHeaders
	t := ""
	for i := range hs.Keys {
		if i == len(hs.Keys)-1 {
			t += hs.Keys[i] + "=" + hs.Vals[i]
		} else {
			t += hs.Keys[i] + "=" + hs.Vals[i] + "&&"
		}
	}
	return dara.String(t)
}

func flatArray(array interface{}, sty string) *string {
	t := reflect.ValueOf(array)
	strs := make([]string, 0)
	for i := 0; i < t.Len(); i++ {
		tmp := t.Index(i)
		if tmp.Kind() == reflect.Ptr || tmp.Kind() == reflect.Interface {
			tmp = tmp.Elem()
		}

		if tmp.Kind() == reflect.Ptr {
			tmp = tmp.Elem()
		}
		if tmp.Kind() == reflect.String {
			strs = append(strs, tmp.String())
		} else {
			inter := tmp.Interface()
			byt, _ := json.Marshal(inter)
			strs = append(strs, string(byt))
		}
	}
	str := ""
	if sty == "simple" {
		str = strings.Join(strs, ",")
	} else if sty == "spaceDelimited" {
		str = strings.Join(strs, " ")
	} else if sty == "pipeDelimited" {
		str = strings.Join(strs, "|")
	}
	return dara.String(str)
}

func buildRpcStringToSign(signedParam map[string]*string, method string) (stringToSign string) {
	signParams := make(map[string]string)
	for key, value := range signedParam {
		signParams[key] = dara.StringValue(value)
	}

	stringToSign = getUrlFormedMap(signParams)
	stringToSign = strings.Replace(stringToSign, "+", "%20", -1)
	stringToSign = strings.Replace(stringToSign, "*", "%2A", -1)
	stringToSign = strings.Replace(stringToSign, "%7E", "~", -1)
	stringToSign = url.QueryEscape(stringToSign)
	stringToSign = method + "&%2F&" + stringToSign
	return
}

func getUrlFormedMap(source map[string]string) (urlEncoded string) {
	urlEncoder := url.Values{}
	for key, value := range source {
		urlEncoder.Add(key, value)
	}
	urlEncoded = urlEncoder.Encode()
	return
}

func sign(stringToSign, accessKeySecret, secretSuffix string) string {
	secret := accessKeySecret + secretSuffix
	signedBytes := shaHmac1(stringToSign, secret)
	signedString := base64.StdEncoding.EncodeToString(signedBytes)
	return signedString
}

func shaHmac1(source, secret string) []byte {
	key := []byte(secret)
	hmac := hmac.New(sha1.New, key)
	hmac.Write([]byte(source))
	return hmac.Sum(nil)
}

func getTimeLeft(rateLimit *string) (_result *int64) {
	if rateLimit != nil {
		pairs := strings.Split(dara.StringValue(rateLimit), ",")
		for _, pair := range pairs {
			kv := strings.Split(pair, ":")
			if len(kv) == 2 {
				key, value := kv[0], kv[1]
				if key == "TimeLeft" {
					timeLeftValue, err := strconv.ParseInt(value, 10, 64)
					if err != nil {
						return nil
					}
					return dara.Int64(timeLeftValue)
				}
			}
		}
	}
	return nil
}

func rsaSign(content, secret string) []byte {
	h := crypto.SHA256.New()
	h.Write([]byte(content))
	hashed := h.Sum(nil)
	priv, err := parsePrivateKey(secret)
	if err != nil {
		return nil
	}
	sign, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed)
	if err != nil {
		return nil
	}
	return sign
}

func parsePrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	privateKey = formatPrivateKey(privateKey)
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("PrivateKey is invalid")
	}
	priKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	switch priKey.(type) {
	case *rsa.PrivateKey:
		return priKey.(*rsa.PrivateKey), nil
	default:
		return nil, nil
	}
}

func formatPrivateKey(privateKey string) string {
	if !strings.HasPrefix(privateKey, PEM_BEGIN) {
		privateKey = PEM_BEGIN + privateKey
	}

	if !strings.HasSuffix(privateKey, PEM_END) {
		privateKey += PEM_END
	}
	return privateKey
}

func getCanonicalHeaders(headers map[string]*string) (string, []string) {
	tmp := make(map[string]string)
	tmpHeader := http.Header{}
	for k, v := range headers {
		if strings.HasPrefix(strings.ToLower(k), "x-acs-") || strings.ToLower(k) == "host" ||
			strings.ToLower(k) == "content-type" {
			tmp[strings.ToLower(k)] = strings.TrimSpace(dara.StringValue(v))
			tmpHeader.Add(strings.ToLower(k), strings.TrimSpace(dara.StringValue(v)))
		}
	}
	hs := newSorter(tmp)

	// Sort the temp by the ascending order
	hs.Sort()
	canonicalheaders := ""
	for _, key := range hs.Keys {
		vals := tmpHeader[textproto.CanonicalMIMEHeaderKey(key)]
		sort.Strings(vals)
		canonicalheaders += key + ":" + strings.Join(vals, ",") + "\n"
	}

	return canonicalheaders, hs.Keys
}

func getCanonicalQueryString(query map[string]*string) string {
	canonicalQueryString := ""
	if dara.IsNil(query) {
		return canonicalQueryString
	}
	tmp := make(map[string]string)
	for k, v := range query {
		tmp[k] = dara.StringValue(v)
	}

	hs := newSorter(tmp)

	// Sort the temp by the ascending order
	hs.Sort()
	for i := range hs.Keys {
		if hs.Vals[i] != "" {
			canonicalQueryString += "&" + hs.Keys[i] + "=" + url.QueryEscape(hs.Vals[i])
		} else {
			canonicalQueryString += "&" + hs.Keys[i] + "="
		}
	}
	canonicalQueryString = strings.Replace(canonicalQueryString, "+", "%20", -1)
	canonicalQueryString = strings.Replace(canonicalQueryString, "*", "%2A", -1)
	canonicalQueryString = strings.Replace(canonicalQueryString, "%7E", "~", -1)

	if canonicalQueryString != "" {
		canonicalQueryString = strings.TrimLeft(canonicalQueryString, "&")
	}
	return canonicalQueryString
}

func ParseToMap(in interface{}) map[string]interface{} {
	if dara.IsNil(in) {
		return nil
	}

	tmp := make(map[string]interface{})
	byt, _ := json.Marshal(in)
	d := json.NewDecoder(bytes.NewReader(byt))
	d.UseNumber()
	err := d.Decode(&tmp)
	if err != nil {
		return nil
	}
	return tmp
}

func GetEndpointRules(product, regionId, endpointType, network, suffix *string) (_result *string, _err error) {
	if dara.StringValue(endpointType) == "regional" {
		if dara.StringValue(regionId) == "" {
			_err = fmt.Errorf("RegionId is empty, please set a valid RegionId")
			return dara.String(""), _err
		}
		_result = dara.String(strings.Replace("<product><suffix><network>.<region_id>.aliyuncs.com",
			"<region_id>", dara.StringValue(regionId), 1))
	} else {
		_result = dara.String("<product><suffix><network>.aliyuncs.com")
	}
	_result = dara.String(strings.Replace(dara.StringValue(_result),
		"<product>", strings.ToLower(dara.StringValue(product)), 1))
	if dara.StringValue(network) == "" || dara.StringValue(network) == "public" {
		_result = dara.String(strings.Replace(dara.StringValue(_result), "<network>", "", 1))
	} else {
		_result = dara.String(strings.Replace(dara.StringValue(_result),
			"<network>", "-"+dara.StringValue(network), 1))
	}
	if dara.StringValue(suffix) == "" {
		_result = dara.String(strings.Replace(dara.StringValue(_result), "<suffix>", "", 1))
	} else {
		_result = dara.String(strings.Replace(dara.StringValue(_result),
			"<suffix>", "-"+dara.StringValue(suffix), 1))
	}
	return _result, nil
}

func ToArray(in interface{}) []map[string]interface{} {
	tmp := make([]map[string]interface{}, 0)
	if dara.IsNil(in) {
		return nil
	}
	byt, _ := json.Marshal(in)
	d := json.NewDecoder(bytes.NewReader(byt))
	d.UseNumber()
	err := d.Decode(&tmp)
	if err != nil {
		return nil
	}
	return tmp
}

func GetEndpoint(endpoint *string, server *bool, endpointType *string) *string {
	if dara.StringValue(endpointType) == "internal" {
		strs := strings.Split(dara.StringValue(endpoint), ".")
		strs[0] += "-internal"
		endpoint = dara.String(strings.Join(strs, "."))
	}
	if dara.BoolValue(server) && dara.StringValue(endpointType) == "accelerate" {
		return dara.String("oss-accelerate.aliyuncs.com")
	}

	return endpoint
}

func toJSONString(a interface{}) *string {
	switch v := a.(type) {
	case *string:
		return v
	case string:
		return dara.String(v)
	case []byte:
		return dara.String(string(v))
	case io.Reader:
		byt, err := ioutil.ReadAll(v)
		if err != nil {
			return nil
		}
		return dara.String(string(byt))
	}
	byt := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(byt)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(a); err != nil {
		return nil
	}
	return dara.String(string(bytes.TrimSpace(byt.Bytes())))
}

func StringifyMapValue(a map[string]interface{}) map[string]*string {
	res := make(map[string]*string)
	for key, value := range a {
		if value != nil {
			res[key] = toJSONString(value)
		}
	}
	return res
}

// MapToFlatStyle transforms a map to a flat style map where keys are prefixed with length info.
// Map keys are transformed from "key" to "#length#key" format.
func MapToFlatStyle(object interface{}) interface{} {
	if object == nil {
		return object
	}

	val := reflect.ValueOf(object)
	if !val.IsValid() {
		return object
	}

	// Handle pointer types
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return object
		}
		val = val.Elem()
	}

	// Handle slice/array (List)
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		result := make([]interface{}, val.Len())
		for i := 0; i < val.Len(); i++ {
			result[i] = MapToFlatStyle(val.Index(i).Interface())
		}
		return result
	}

	// Handle struct (TeaModel equivalent)
	if val.Kind() == reflect.Struct {
		// Create a pointer to the struct so we can modify it
		if reflect.TypeOf(object).Kind() == reflect.Ptr {
			// Already a pointer
			val = reflect.ValueOf(object).Elem()
		} else {
			// Make a copy and work with pointer
			ptrVal := reflect.New(val.Type())
			ptrVal.Elem().Set(val)
			val = ptrVal.Elem()
			object = ptrVal.Interface()
		}

		valType := val.Type()
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldType := valType.Field(i)

			// Skip unexported fields
			if !field.CanSet() {
				continue
			}

			fieldValue := field.Interface()

			// Check if field is a map
			if field.Kind() == reflect.Map {
				flatMap := make(map[string]interface{})
				iter := field.MapRange()
				for iter.Next() {
					key := iter.Key()
					value := iter.Value()
					keyStr := fmt.Sprintf("%v", key.Interface())
					flatKey := fmt.Sprintf("#%d#%s", len(keyStr), keyStr)
					flatMap[flatKey] = MapToFlatStyle(value.Interface())
				}

				// Set the flattened map back to the field
				newMapValue := reflect.MakeMap(field.Type())
				for k, v := range flatMap {
					keyVal := reflect.ValueOf(k)
					valVal := reflect.ValueOf(v)
					if valVal.IsValid() && valVal.Type().AssignableTo(field.Type().Elem()) {
						newMapValue.SetMapIndex(keyVal, valVal)
					} else if valVal.IsValid() {
						// Try to convert the value
						if field.Type().Elem().Kind() == reflect.Interface {
							newMapValue.SetMapIndex(keyVal, valVal)
						}
					}
				}
				if newMapValue.Len() > 0 {
					field.Set(newMapValue)
				}
			} else {
				// Recursively process other fields
				processed := MapToFlatStyle(fieldValue)
				if processed != nil && field.CanSet() {
					processedVal := reflect.ValueOf(processed)
					if processedVal.IsValid() {
						// Only set if types are compatible
						if processedVal.Type().AssignableTo(fieldType.Type) {
							field.Set(processedVal)
						} else if fieldType.Type.Kind() == reflect.Interface {
							field.Set(processedVal)
						} else if processedVal.Type().ConvertibleTo(fieldType.Type) {
							field.Set(processedVal.Convert(fieldType.Type))
						}
					}
				}
			}
		}
		return object
	}

	// Handle map
	if val.Kind() == reflect.Map {
		flatMap := make(map[string]interface{})
		iter := val.MapRange()
		for iter.Next() {
			key := iter.Key()
			value := iter.Value()
			keyStr := fmt.Sprintf("%v", key.Interface())
			flatKey := fmt.Sprintf("#%d#%s", len(keyStr), keyStr)
			flatMap[flatKey] = MapToFlatStyle(value.Interface())
		}
		return flatMap
	}

	return object
}
