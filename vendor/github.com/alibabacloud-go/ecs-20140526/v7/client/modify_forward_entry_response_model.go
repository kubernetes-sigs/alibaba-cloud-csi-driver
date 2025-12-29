// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyForwardEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyForwardEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyForwardEntryResponse
	GetStatusCode() *int32
	SetBody(v *ModifyForwardEntryResponseBody) *ModifyForwardEntryResponse
	GetBody() *ModifyForwardEntryResponseBody
}

type ModifyForwardEntryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyForwardEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyForwardEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyForwardEntryResponse) GoString() string {
	return s.String()
}

func (s *ModifyForwardEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyForwardEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyForwardEntryResponse) GetBody() *ModifyForwardEntryResponseBody {
	return s.Body
}

func (s *ModifyForwardEntryResponse) SetHeaders(v map[string]*string) *ModifyForwardEntryResponse {
	s.Headers = v
	return s
}

func (s *ModifyForwardEntryResponse) SetStatusCode(v int32) *ModifyForwardEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyForwardEntryResponse) SetBody(v *ModifyForwardEntryResponseBody) *ModifyForwardEntryResponse {
	s.Body = v
	return s
}

func (s *ModifyForwardEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
