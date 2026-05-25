// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTerminalSessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeTerminalSessionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeTerminalSessionsResponseBody
	GetRequestId() *string
	SetSessions(v *DescribeTerminalSessionsResponseBodySessions) *DescribeTerminalSessionsResponseBody
	GetSessions() *DescribeTerminalSessionsResponseBodySessions
}

type DescribeTerminalSessionsResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sessions  *DescribeTerminalSessionsResponseBodySessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Struct"`
}

func (s DescribeTerminalSessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTerminalSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTerminalSessionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTerminalSessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTerminalSessionsResponseBody) GetSessions() *DescribeTerminalSessionsResponseBodySessions {
	return s.Sessions
}

func (s *DescribeTerminalSessionsResponseBody) SetNextToken(v string) *DescribeTerminalSessionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeTerminalSessionsResponseBody) SetRequestId(v string) *DescribeTerminalSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTerminalSessionsResponseBody) SetSessions(v *DescribeTerminalSessionsResponseBodySessions) *DescribeTerminalSessionsResponseBody {
	s.Sessions = v
	return s
}

func (s *DescribeTerminalSessionsResponseBody) Validate() error {
	if s.Sessions != nil {
		if err := s.Sessions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTerminalSessionsResponseBodySessions struct {
	Session []*DescribeTerminalSessionsResponseBodySessionsSession `json:"Session,omitempty" xml:"Session,omitempty" type:"Repeated"`
}

func (s DescribeTerminalSessionsResponseBodySessions) String() string {
	return dara.Prettify(s)
}

func (s DescribeTerminalSessionsResponseBodySessions) GoString() string {
	return s.String()
}

func (s *DescribeTerminalSessionsResponseBodySessions) GetSession() []*DescribeTerminalSessionsResponseBodySessionsSession {
	return s.Session
}

func (s *DescribeTerminalSessionsResponseBodySessions) SetSession(v []*DescribeTerminalSessionsResponseBodySessionsSession) *DescribeTerminalSessionsResponseBodySessions {
	s.Session = v
	return s
}

func (s *DescribeTerminalSessionsResponseBodySessions) Validate() error {
	if s.Session != nil {
		for _, item := range s.Session {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTerminalSessionsResponseBodySessionsSession struct {
	ClientIP     *string                                                         `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	Connections  *DescribeTerminalSessionsResponseBodySessionsSessionConnections `json:"Connections,omitempty" xml:"Connections,omitempty" type:"Struct"`
	CreationTime *string                                                         `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	IdentityType *string                                                         `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	PortNumber   *int32                                                          `json:"PortNumber,omitempty" xml:"PortNumber,omitempty"`
	PrincipalId  *string                                                         `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	SessionId    *string                                                         `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TargetServer *string                                                         `json:"TargetServer,omitempty" xml:"TargetServer,omitempty"`
	Username     *string                                                         `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeTerminalSessionsResponseBodySessionsSession) String() string {
	return dara.Prettify(s)
}

func (s DescribeTerminalSessionsResponseBodySessionsSession) GoString() string {
	return s.String()
}

func (s *DescribeTerminalSessionsResponseBodySessionsSession) GetClientIP() *string {
	return s.ClientIP
}

func (s *DescribeTerminalSessionsResponseBodySessionsSession) GetConnections() *DescribeTerminalSessionsResponseBodySessionsSessionConnections {
	return s.Connections
}

func (s *DescribeTerminalSessionsResponseBodySessionsSession) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeTerminalSessionsResponseBodySessionsSession) GetIdentityType() *string {
	return s.IdentityType
}

func (s *DescribeTerminalSessionsResponseBodySessionsSession) GetPortNumber() *int32 {
	return s.PortNumber
}

func (s *DescribeTerminalSessionsResponseBodySessionsSession) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *DescribeTerminalSessionsResponseBodySessionsSession) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeTerminalSessionsResponseBodySessionsSession) GetTargetServer() *string {
	return s.TargetServer
}

func (s *DescribeTerminalSessionsResponseBodySessionsSession) GetUsername() *string {
	return s.Username
}

func (s *DescribeTerminalSessionsResponseBodySessionsSession) SetClientIP(v string) *DescribeTerminalSessionsResponseBodySessionsSession {
	s.ClientIP = &v
	return s
}

func (s *DescribeTerminalSessionsResponseBodySessionsSession) SetConnections(v *DescribeTerminalSessionsResponseBodySessionsSessionConnections) *DescribeTerminalSessionsResponseBodySessionsSession {
	s.Connections = v
	return s
}

func (s *DescribeTerminalSessionsResponseBodySessionsSession) SetCreationTime(v string) *DescribeTerminalSessionsResponseBodySessionsSession {
	s.CreationTime = &v
	return s
}

func (s *DescribeTerminalSessionsResponseBodySessionsSession) SetIdentityType(v string) *DescribeTerminalSessionsResponseBodySessionsSession {
	s.IdentityType = &v
	return s
}

func (s *DescribeTerminalSessionsResponseBodySessionsSession) SetPortNumber(v int32) *DescribeTerminalSessionsResponseBodySessionsSession {
	s.PortNumber = &v
	return s
}

func (s *DescribeTerminalSessionsResponseBodySessionsSession) SetPrincipalId(v string) *DescribeTerminalSessionsResponseBodySessionsSession {
	s.PrincipalId = &v
	return s
}

func (s *DescribeTerminalSessionsResponseBodySessionsSession) SetSessionId(v string) *DescribeTerminalSessionsResponseBodySessionsSession {
	s.SessionId = &v
	return s
}

func (s *DescribeTerminalSessionsResponseBodySessionsSession) SetTargetServer(v string) *DescribeTerminalSessionsResponseBodySessionsSession {
	s.TargetServer = &v
	return s
}

func (s *DescribeTerminalSessionsResponseBodySessionsSession) SetUsername(v string) *DescribeTerminalSessionsResponseBodySessionsSession {
	s.Username = &v
	return s
}

func (s *DescribeTerminalSessionsResponseBodySessionsSession) Validate() error {
	if s.Connections != nil {
		if err := s.Connections.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTerminalSessionsResponseBodySessionsSessionConnections struct {
	Connection []*DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection `json:"Connection,omitempty" xml:"Connection,omitempty" type:"Repeated"`
}

func (s DescribeTerminalSessionsResponseBodySessionsSessionConnections) String() string {
	return dara.Prettify(s)
}

func (s DescribeTerminalSessionsResponseBodySessionsSessionConnections) GoString() string {
	return s.String()
}

func (s *DescribeTerminalSessionsResponseBodySessionsSessionConnections) GetConnection() []*DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection {
	return s.Connection
}

func (s *DescribeTerminalSessionsResponseBodySessionsSessionConnections) SetConnection(v []*DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection) *DescribeTerminalSessionsResponseBodySessionsSessionConnections {
	s.Connection = v
	return s
}

func (s *DescribeTerminalSessionsResponseBodySessionsSessionConnections) Validate() error {
	if s.Connection != nil {
		for _, item := range s.Connection {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection struct {
	ClosedReason *string `json:"ClosedReason,omitempty" xml:"ClosedReason,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FailedDetail *string `json:"FailedDetail,omitempty" xml:"FailedDetail,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection) String() string {
	return dara.Prettify(s)
}

func (s DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection) GoString() string {
	return s.String()
}

func (s *DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection) GetClosedReason() *string {
	return s.ClosedReason
}

func (s *DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection) GetFailedDetail() *string {
	return s.FailedDetail
}

func (s *DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection) GetStatus() *string {
	return s.Status
}

func (s *DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection) SetClosedReason(v string) *DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection {
	s.ClosedReason = &v
	return s
}

func (s *DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection) SetEndTime(v string) *DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection {
	s.EndTime = &v
	return s
}

func (s *DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection) SetFailedDetail(v string) *DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection {
	s.FailedDetail = &v
	return s
}

func (s *DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection) SetInstanceId(v string) *DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection {
	s.InstanceId = &v
	return s
}

func (s *DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection) SetStartTime(v string) *DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection {
	s.StartTime = &v
	return s
}

func (s *DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection) SetStatus(v string) *DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection {
	s.Status = &v
	return s
}

func (s *DescribeTerminalSessionsResponseBodySessionsSessionConnectionsConnection) Validate() error {
	return dara.Validate(s)
}
