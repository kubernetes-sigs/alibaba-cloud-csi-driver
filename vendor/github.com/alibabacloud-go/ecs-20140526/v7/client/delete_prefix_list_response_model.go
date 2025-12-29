// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrefixListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePrefixListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePrefixListResponse
	GetStatusCode() *int32
	SetBody(v *DeletePrefixListResponseBody) *DeletePrefixListResponse
	GetBody() *DeletePrefixListResponseBody
}

type DeletePrefixListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrefixListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrefixListResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePrefixListResponse) GoString() string {
	return s.String()
}

func (s *DeletePrefixListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePrefixListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePrefixListResponse) GetBody() *DeletePrefixListResponseBody {
	return s.Body
}

func (s *DeletePrefixListResponse) SetHeaders(v map[string]*string) *DeletePrefixListResponse {
	s.Headers = v
	return s
}

func (s *DeletePrefixListResponse) SetStatusCode(v int32) *DeletePrefixListResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrefixListResponse) SetBody(v *DeletePrefixListResponseBody) *DeletePrefixListResponse {
	s.Body = v
	return s
}

func (s *DeletePrefixListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
