// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"fmt"

	"github.com/alibabacloud-go/tea/dara"
)

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
