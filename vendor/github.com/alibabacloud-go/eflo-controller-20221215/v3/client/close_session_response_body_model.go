// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloseSessionResponseBody
	GetRequestId() *string
	SetSessionId(v string) *CloseSessionResponseBody
	GetSessionId() *string
	SetState(v string) *CloseSessionResponseBody
	GetState() *string
}

type CloseSessionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 07AA3A1F-321E-50D8-B834-88C411331C94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// i206495551737511455528
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// status of session
	//
	// example:
	//
	// Inactive
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CloseSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CloseSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseSessionResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *CloseSessionResponseBody) GetState() *string {
	return s.State
}

func (s *CloseSessionResponseBody) SetRequestId(v string) *CloseSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseSessionResponseBody) SetSessionId(v string) *CloseSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *CloseSessionResponseBody) SetState(v string) *CloseSessionResponseBody {
	s.State = &v
	return s
}

func (s *CloseSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
