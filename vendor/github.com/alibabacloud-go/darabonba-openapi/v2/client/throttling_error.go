// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"fmt"

	"github.com/alibabacloud-go/tea/dara"
)

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
