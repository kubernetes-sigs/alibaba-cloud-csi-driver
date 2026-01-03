// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrefixListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePrefixListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePrefixListResponse
	GetStatusCode() *int32
	SetBody(v *CreatePrefixListResponseBody) *CreatePrefixListResponse
	GetBody() *CreatePrefixListResponseBody
}

type CreatePrefixListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrefixListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrefixListResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePrefixListResponse) GoString() string {
	return s.String()
}

func (s *CreatePrefixListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePrefixListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePrefixListResponse) GetBody() *CreatePrefixListResponseBody {
	return s.Body
}

func (s *CreatePrefixListResponse) SetHeaders(v map[string]*string) *CreatePrefixListResponse {
	s.Headers = v
	return s
}

func (s *CreatePrefixListResponse) SetStatusCode(v int32) *CreatePrefixListResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrefixListResponse) SetBody(v *CreatePrefixListResponseBody) *CreatePrefixListResponse {
	s.Body = v
	return s
}

func (s *CreatePrefixListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
