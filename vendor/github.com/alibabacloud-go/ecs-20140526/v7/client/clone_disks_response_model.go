// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneDisksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloneDisksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloneDisksResponse
	GetStatusCode() *int32
	SetBody(v *CloneDisksResponseBody) *CloneDisksResponse
	GetBody() *CloneDisksResponseBody
}

type CloneDisksResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneDisksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneDisksResponse) String() string {
	return dara.Prettify(s)
}

func (s CloneDisksResponse) GoString() string {
	return s.String()
}

func (s *CloneDisksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloneDisksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloneDisksResponse) GetBody() *CloneDisksResponseBody {
	return s.Body
}

func (s *CloneDisksResponse) SetHeaders(v map[string]*string) *CloneDisksResponse {
	s.Headers = v
	return s
}

func (s *CloneDisksResponse) SetStatusCode(v int32) *CloneDisksResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneDisksResponse) SetBody(v *CloneDisksResponseBody) *CloneDisksResponse {
	s.Body = v
	return s
}

func (s *CloneDisksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
