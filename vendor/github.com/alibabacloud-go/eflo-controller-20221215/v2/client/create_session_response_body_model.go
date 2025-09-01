// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSessionResponseBody
	GetRequestId() *string
	SetServerSn(v string) *CreateSessionResponseBody
	GetServerSn() *string
	SetSessionId(v string) *CreateSessionResponseBody
	GetSessionId() *string
	SetSessionToken(v string) *CreateSessionResponseBody
	GetSessionToken() *string
	SetWssEndpoint(v string) *CreateSessionResponseBody
	GetWssEndpoint() *string
}

type CreateSessionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The node ID.
	//
	// example:
	//
	// 2A59143F1
	ServerSn *string `json:"ServerSn,omitempty" xml:"ServerSn,omitempty"`
	// The session ID.
	//
	// example:
	//
	// i207023871669364793713
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The session credential.
	//
	// example:
	//
	// 03f53c719015a9ad4f4f55d66cac2dac161b18e8065ca75a3220b89de389c980
	SessionToken *string `json:"SessionToken,omitempty" xml:"SessionToken,omitempty"`
	// The WebSocket address.
	//
	// example:
	//
	// ws://x.x.x.x:xx/calypso_web_console
	WssEndpoint *string `json:"WssEndpoint,omitempty" xml:"WssEndpoint,omitempty"`
}

func (s CreateSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSessionResponseBody) GetServerSn() *string {
	return s.ServerSn
}

func (s *CreateSessionResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateSessionResponseBody) GetSessionToken() *string {
	return s.SessionToken
}

func (s *CreateSessionResponseBody) GetWssEndpoint() *string {
	return s.WssEndpoint
}

func (s *CreateSessionResponseBody) SetRequestId(v string) *CreateSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSessionResponseBody) SetServerSn(v string) *CreateSessionResponseBody {
	s.ServerSn = &v
	return s
}

func (s *CreateSessionResponseBody) SetSessionId(v string) *CreateSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *CreateSessionResponseBody) SetSessionToken(v string) *CreateSessionResponseBody {
	s.SessionToken = &v
	return s
}

func (s *CreateSessionResponseBody) SetWssEndpoint(v string) *CreateSessionResponseBody {
	s.WssEndpoint = &v
	return s
}

func (s *CreateSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
