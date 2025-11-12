// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMessage(v string) *ApproveOperationResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ApproveOperationResponseBody
	GetRequestId() *string
}

type ApproveOperationResponseBody struct {
	// The error message.
	//
	// example:
	//
	// Resource not found
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8F065DDD-6996-5973-9691-9EC57BD0072E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveOperationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApproveOperationResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveOperationResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ApproveOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApproveOperationResponseBody) SetErrorMessage(v string) *ApproveOperationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ApproveOperationResponseBody) SetRequestId(v string) *ApproveOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApproveOperationResponseBody) Validate() error {
	return dara.Validate(s)
}
