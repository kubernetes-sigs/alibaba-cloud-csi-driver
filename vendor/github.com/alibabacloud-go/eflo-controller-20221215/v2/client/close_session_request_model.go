// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionId(v string) *CloseSessionRequest
	GetSessionId() *string
	SetSessionToken(v string) *CloseSessionRequest
	GetSessionToken() *string
}

type CloseSessionRequest struct {
	// The session ID.
	//
	// example:
	//
	// i207023871669364793713
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// Session token.
	//
	// example:
	//
	// 03f53c719015a9ad4f4f55d66cac2dac161b18e8065ca75a3220b89de389c980
	SessionToken *string `json:"SessionToken,omitempty" xml:"SessionToken,omitempty"`
}

func (s CloseSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseSessionRequest) GoString() string {
	return s.String()
}

func (s *CloseSessionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *CloseSessionRequest) GetSessionToken() *string {
	return s.SessionToken
}

func (s *CloseSessionRequest) SetSessionId(v string) *CloseSessionRequest {
	s.SessionId = &v
	return s
}

func (s *CloseSessionRequest) SetSessionToken(v string) *CloseSessionRequest {
	s.SessionToken = &v
	return s
}

func (s *CloseSessionRequest) Validate() error {
	return dara.Validate(s)
}
