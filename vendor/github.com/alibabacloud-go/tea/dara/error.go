package dara

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"github.com/alibabacloud-go/tea/tea"
)

type BaseError interface {
	error
	GetName() *string
	GetCode() *string
}

type ResponseError interface {
	BaseError
	GetRetryAfter() *int64
	GetStatusCode() *int
	GetAccessDeniedDetail() map[string]interface{}
	GetDescription() *string
	GetData() map[string]interface{}
}

// SDKError struct is used save error code and message
type SDKError struct {
	BaseError
	Code               *string
	Name               *string
	StatusCode         *int
	Message            *string
	Data               *string
	Stack              *string
	errMsg             *string
	Description        *string
	AccessDeniedDetail map[string]interface{}
}

// CastError is used for cast type fails
type CastError struct {
	Message *string
}

func TeaSDKError(err error) error {
	if(err == nil) {
		return nil
	}

	if te, ok := err.(*SDKError); ok {
		return tea.NewSDKError(map[string]interface{}{
			"code": StringValue(te.Code),
			"statusCode": IntValue(te.StatusCode),
			"message": StringValue(te.Message),
			"data": te.Data,
			"description": StringValue(te.Description),
			"accessDeniedDetail": te.AccessDeniedDetail,
		})
	}

	if respErr, ok := err.(ResponseError); ok { 
		return tea.NewSDKError(map[string]interface{}{
			"code": StringValue(respErr.GetCode()),
			"statusCode": IntValue(respErr.GetStatusCode()),
			"message": respErr.Error(),
			"description": StringValue(respErr.GetDescription()),
			"data": respErr.GetData(),
			"accessDeniedDetail": respErr.GetAccessDeniedDetail(),
		})
	}

	if baseErr, ok := err.(BaseError); ok { 
		return tea.NewSDKError(map[string]interface{}{
			"code": StringValue(baseErr.GetCode()),
			"message": baseErr.Error(),
		})
	}

	return err
}

// NewSDKError is used for shortly create SDKError object
func NewSDKError(obj map[string]interface{}) *SDKError {
	err := &SDKError{}
	err.Name = String("BaseError")
	if val, ok := obj["code"].(int); ok {
		err.Code = String(strconv.Itoa(val))
	} else if val, ok := obj["code"].(string); ok {
		err.Code = String(val)
	}
 
	if obj["message"] != nil {
		err.Message = String(obj["message"].(string))
	}

	if obj["name"] != nil {

	}

	if obj["description"] != nil {
		err.Description = String(obj["description"].(string))
	}
	if detail := obj["accessDeniedDetail"]; detail != nil {
		r := reflect.ValueOf(detail)
		if r.Kind().String() == "map" {
			res := make(map[string]interface{})
			tmp := r.MapKeys()
			for _, key := range tmp {
				res[key.String()] = r.MapIndex(key).Interface()
			}
			err.AccessDeniedDetail = res
		}
	}
	if data := obj["data"]; data != nil {
		r := reflect.ValueOf(data)
		if r.Kind().String() == "map" {
			res := make(map[string]interface{})
			tmp := r.MapKeys()
			for _, key := range tmp {
				res[key.String()] = r.MapIndex(key).Interface()
			}
			if statusCode := res["statusCode"]; statusCode != nil {
				if code, ok := statusCode.(int); ok {
					err.StatusCode = Int(code)
				} else if tmp, ok := statusCode.(string); ok {
					code, err_ := strconv.Atoi(tmp)
					if err_ == nil {
						err.StatusCode = Int(code)
					}
				} else if code, ok := statusCode.(*int); ok {
					err.StatusCode = code
				}
			}
		}
		byt := bytes.NewBuffer([]byte{})
		jsonEncoder := json.NewEncoder(byt)
		jsonEncoder.SetEscapeHTML(false)
		jsonEncoder.Encode(data)
		err.Data = String(string(bytes.TrimSpace(byt.Bytes())))
	}

	if statusCode, ok := obj["statusCode"].(int); ok {
		err.StatusCode = Int(statusCode)
	} else if status, ok := obj["statusCode"].(string); ok {
		statusCode, err_ := strconv.Atoi(status)
		if err_ == nil {
			err.StatusCode = Int(statusCode)
		}
	}

	return err
}

func (err *SDKError) ErrorName() *string {
	return err.Name
}

func (err *SDKError) ErrorMessage() *string {
	return err.Message
}

func (err *SDKError) GetCode() *string {
	return err.Code
}

// Set ErrMsg by msg
func (err *SDKError) SetErrMsg(msg string) {
	err.errMsg = String(msg)
}

func (err *SDKError) Error() string {
	if err.errMsg == nil {
		str := fmt.Sprintf("SDKError:\n   StatusCode: %d\n   Code: %s\n   Message: %s\n   Data: %s\n",
			IntValue(err.StatusCode), StringValue(err.Code), StringValue(err.Message), StringValue(err.Data))
		err.SetErrMsg(str)
	}
	return StringValue(err.errMsg)
}

// Return message of CastError
func (err *CastError) Error() string {
	return StringValue(err.Message)
}

func Retryable(err error) *bool {
	if err == nil {
		return Bool(false)
	}
	if realErr, ok := err.(*SDKError); ok {
		if realErr.StatusCode == nil {
			return Bool(false)
		}
		code := IntValue(realErr.StatusCode)
		return Bool(code >= http.StatusInternalServerError)
	}
	return Bool(true)
}

// NewCastError is used for cast type fails
func NewCastError(message *string) (err error) {
	return &CastError{
		Message: message,
	}
}
