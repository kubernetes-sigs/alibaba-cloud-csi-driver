// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *CreateCommandResponseBody
	GetCommandId() *string
	SetRequestId(v string) *CreateCommandResponseBody
	GetRequestId() *string
}

type CreateCommandResponseBody struct {
	// The ID of the command.
	//
	// example:
	//
	// c-7d2a745b412b4601b2d47f6a768d****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCommandResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCommandResponseBody) GetCommandId() *string {
	return s.CommandId
}

func (s *CreateCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCommandResponseBody) SetCommandId(v string) *CreateCommandResponseBody {
	s.CommandId = &v
	return s
}

func (s *CreateCommandResponseBody) SetRequestId(v string) *CreateCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCommandResponseBody) Validate() error {
	return dara.Validate(s)
}
