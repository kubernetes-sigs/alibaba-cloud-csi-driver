// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"fmt"

	"github.com/alibabacloud-go/tea/dara"
)

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
