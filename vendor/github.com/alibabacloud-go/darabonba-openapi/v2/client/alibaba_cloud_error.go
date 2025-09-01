// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"fmt"

	"github.com/alibabacloud-go/tea/dara"
)

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
