// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticSpaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgenticSpaceId(v string) *CreateAgenticSpaceResponseBody
	GetAgenticSpaceId() *string
	SetRequestId(v string) *CreateAgenticSpaceResponseBody
	GetRequestId() *string
}

type CreateAgenticSpaceResponseBody struct {
	// example:
	//
	// agentic-229oypxjgpau2****
	AgenticSpaceId *string `json:"AgenticSpaceId,omitempty" xml:"AgenticSpaceId,omitempty"`
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAgenticSpaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgenticSpaceResponseBody) GetAgenticSpaceId() *string {
	return s.AgenticSpaceId
}

func (s *CreateAgenticSpaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgenticSpaceResponseBody) SetAgenticSpaceId(v string) *CreateAgenticSpaceResponseBody {
	s.AgenticSpaceId = &v
	return s
}

func (s *CreateAgenticSpaceResponseBody) SetRequestId(v string) *CreateAgenticSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgenticSpaceResponseBody) Validate() error {
	return dara.Validate(s)
}
