// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTagsResponse
	GetStatusCode() *int32
	SetBody(v *AddTagsResponseBody) *AddTagsResponse
	GetBody() *AddTagsResponseBody
}

type AddTagsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTagsResponse) GoString() string {
	return s.String()
}

func (s *AddTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTagsResponse) GetBody() *AddTagsResponseBody {
	return s.Body
}

func (s *AddTagsResponse) SetHeaders(v map[string]*string) *AddTagsResponse {
	s.Headers = v
	return s
}

func (s *AddTagsResponse) SetStatusCode(v int32) *AddTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTagsResponse) SetBody(v *AddTagsResponseBody) *AddTagsResponse {
	s.Body = v
	return s
}

func (s *AddTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
