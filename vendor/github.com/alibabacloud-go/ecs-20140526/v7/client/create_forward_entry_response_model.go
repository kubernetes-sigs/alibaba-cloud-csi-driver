// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateForwardEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateForwardEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateForwardEntryResponse
	GetStatusCode() *int32
	SetBody(v *CreateForwardEntryResponseBody) *CreateForwardEntryResponse
	GetBody() *CreateForwardEntryResponseBody
}

type CreateForwardEntryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateForwardEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateForwardEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardEntryResponse) GoString() string {
	return s.String()
}

func (s *CreateForwardEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateForwardEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateForwardEntryResponse) GetBody() *CreateForwardEntryResponseBody {
	return s.Body
}

func (s *CreateForwardEntryResponse) SetHeaders(v map[string]*string) *CreateForwardEntryResponse {
	s.Headers = v
	return s
}

func (s *CreateForwardEntryResponse) SetStatusCode(v int32) *CreateForwardEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateForwardEntryResponse) SetBody(v *CreateForwardEntryResponseBody) *CreateForwardEntryResponse {
	s.Body = v
	return s
}

func (s *CreateForwardEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
