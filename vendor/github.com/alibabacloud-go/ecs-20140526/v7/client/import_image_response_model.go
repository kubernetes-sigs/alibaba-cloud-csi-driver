// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportImageResponse
	GetStatusCode() *int32
	SetBody(v *ImportImageResponseBody) *ImportImageResponse
	GetBody() *ImportImageResponseBody
}

type ImportImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportImageResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportImageResponse) GoString() string {
	return s.String()
}

func (s *ImportImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportImageResponse) GetBody() *ImportImageResponseBody {
	return s.Body
}

func (s *ImportImageResponse) SetHeaders(v map[string]*string) *ImportImageResponse {
	s.Headers = v
	return s
}

func (s *ImportImageResponse) SetStatusCode(v int32) *ImportImageResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportImageResponse) SetBody(v *ImportImageResponseBody) *ImportImageResponse {
	s.Body = v
	return s
}

func (s *ImportImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
