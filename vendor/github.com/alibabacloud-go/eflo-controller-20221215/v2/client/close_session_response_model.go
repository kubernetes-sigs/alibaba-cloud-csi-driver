// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseSessionResponse
	GetStatusCode() *int32
	SetBody(v *CloseSessionResponseBody) *CloseSessionResponse
	GetBody() *CloseSessionResponseBody
}

type CloseSessionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseSessionResponse) GoString() string {
	return s.String()
}

func (s *CloseSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseSessionResponse) GetBody() *CloseSessionResponseBody {
	return s.Body
}

func (s *CloseSessionResponse) SetHeaders(v map[string]*string) *CloseSessionResponse {
	s.Headers = v
	return s
}

func (s *CloseSessionResponse) SetStatusCode(v int32) *CloseSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseSessionResponse) SetBody(v *CloseSessionResponseBody) *CloseSessionResponse {
	s.Body = v
	return s
}

func (s *CloseSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
