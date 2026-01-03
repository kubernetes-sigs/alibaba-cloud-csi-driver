// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveTagsResponse
	GetStatusCode() *int32
	SetBody(v *RemoveTagsResponseBody) *RemoveTagsResponse
	GetBody() *RemoveTagsResponseBody
}

type RemoveTagsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveTagsResponse) GoString() string {
	return s.String()
}

func (s *RemoveTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveTagsResponse) GetBody() *RemoveTagsResponseBody {
	return s.Body
}

func (s *RemoveTagsResponse) SetHeaders(v map[string]*string) *RemoveTagsResponse {
	s.Headers = v
	return s
}

func (s *RemoveTagsResponse) SetStatusCode(v int32) *RemoveTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveTagsResponse) SetBody(v *RemoveTagsResponseBody) *RemoveTagsResponse {
	s.Body = v
	return s
}

func (s *RemoveTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
