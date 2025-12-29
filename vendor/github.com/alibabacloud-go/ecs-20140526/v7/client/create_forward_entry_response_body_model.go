// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateForwardEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForwardEntryId(v string) *CreateForwardEntryResponseBody
	GetForwardEntryId() *string
	SetRequestId(v string) *CreateForwardEntryResponseBody
	GetRequestId() *string
}

type CreateForwardEntryResponseBody struct {
	ForwardEntryId *string `json:"ForwardEntryId,omitempty" xml:"ForwardEntryId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateForwardEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardEntryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateForwardEntryResponseBody) GetForwardEntryId() *string {
	return s.ForwardEntryId
}

func (s *CreateForwardEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateForwardEntryResponseBody) SetForwardEntryId(v string) *CreateForwardEntryResponseBody {
	s.ForwardEntryId = &v
	return s
}

func (s *CreateForwardEntryResponseBody) SetRequestId(v string) *CreateForwardEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateForwardEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
