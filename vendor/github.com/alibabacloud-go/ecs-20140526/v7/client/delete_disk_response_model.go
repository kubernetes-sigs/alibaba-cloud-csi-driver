// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDiskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDiskResponseBody) *DeleteDiskResponse
	GetBody() *DeleteDiskResponseBody
}

type DeleteDiskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDiskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiskResponse) GoString() string {
	return s.String()
}

func (s *DeleteDiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDiskResponse) GetBody() *DeleteDiskResponseBody {
	return s.Body
}

func (s *DeleteDiskResponse) SetHeaders(v map[string]*string) *DeleteDiskResponse {
	s.Headers = v
	return s
}

func (s *DeleteDiskResponse) SetStatusCode(v int32) *DeleteDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDiskResponse) SetBody(v *DeleteDiskResponseBody) *DeleteDiskResponse {
	s.Body = v
	return s
}

func (s *DeleteDiskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
