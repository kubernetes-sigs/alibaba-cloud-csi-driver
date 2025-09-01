// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *CreateSessionRequest
	GetNodeId() *string
	SetSessionType(v string) *CreateSessionRequest
	GetSessionType() *string
	SetStartTime(v string) *CreateSessionRequest
	GetStartTime() *string
}

type CreateSessionRequest struct {
	// The instance ID.
	//
	// example:
	//
	// e01-cn-kvw44e6dn04
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The type of the session corresponding to the session package.
	//
	// example:
	//
	// Valid values: Sol (default): based on serial port Assistant: based on cloud assistant
	SessionType *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	// The start time. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1669340937156
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateSessionRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateSessionRequest) GetSessionType() *string {
	return s.SessionType
}

func (s *CreateSessionRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateSessionRequest) SetNodeId(v string) *CreateSessionRequest {
	s.NodeId = &v
	return s
}

func (s *CreateSessionRequest) SetSessionType(v string) *CreateSessionRequest {
	s.SessionType = &v
	return s
}

func (s *CreateSessionRequest) SetStartTime(v string) *CreateSessionRequest {
	s.StartTime = &v
	return s
}

func (s *CreateSessionRequest) Validate() error {
	return dara.Validate(s)
}
