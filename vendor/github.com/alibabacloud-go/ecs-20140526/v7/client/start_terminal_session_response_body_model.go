// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTerminalSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartTerminalSessionResponseBody
	GetRequestId() *string
	SetSecurityToken(v string) *StartTerminalSessionResponseBody
	GetSecurityToken() *string
	SetSessionId(v string) *StartTerminalSessionResponseBody
	GetSessionId() *string
	SetWebSocketUrl(v string) *StartTerminalSessionResponseBody
	GetWebSocketUrl() *string
}

type StartTerminalSessionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EB5173B0-8E80-564E-AAD1-3135412*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security token included in the WebSocket request header. The system uses this token to authenticate the request.
	//
	// example:
	//
	// d86c2df2-d19c-4bd8-b817-a19ef123****
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The session ID.
	//
	// example:
	//
	// s-hz023od0x9****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The URL of the WebSocket session that is used to connect to the instance. The URL includes the session ID (`SessionId`) and the authentication token (`SecurityToken`).
	//
	// example:
	//
	// wss://cn-hangzhou.axt.aliyuncs.com/session?sessionId=s-hz023od0x9****&token=d86c2df2-d19c-4bd8-b817-a19ef123****
	WebSocketUrl *string `json:"WebSocketUrl,omitempty" xml:"WebSocketUrl,omitempty"`
}

func (s StartTerminalSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartTerminalSessionResponseBody) GoString() string {
	return s.String()
}

func (s *StartTerminalSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartTerminalSessionResponseBody) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *StartTerminalSessionResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *StartTerminalSessionResponseBody) GetWebSocketUrl() *string {
	return s.WebSocketUrl
}

func (s *StartTerminalSessionResponseBody) SetRequestId(v string) *StartTerminalSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTerminalSessionResponseBody) SetSecurityToken(v string) *StartTerminalSessionResponseBody {
	s.SecurityToken = &v
	return s
}

func (s *StartTerminalSessionResponseBody) SetSessionId(v string) *StartTerminalSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *StartTerminalSessionResponseBody) SetWebSocketUrl(v string) *StartTerminalSessionResponseBody {
	s.WebSocketUrl = &v
	return s
}

func (s *StartTerminalSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
