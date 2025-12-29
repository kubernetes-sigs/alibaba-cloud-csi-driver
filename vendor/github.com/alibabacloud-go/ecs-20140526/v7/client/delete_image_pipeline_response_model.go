// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImagePipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteImagePipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteImagePipelineResponse
	GetStatusCode() *int32
	SetBody(v *DeleteImagePipelineResponseBody) *DeleteImagePipelineResponse
	GetBody() *DeleteImagePipelineResponseBody
}

type DeleteImagePipelineResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteImagePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteImagePipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteImagePipelineResponse) GoString() string {
	return s.String()
}

func (s *DeleteImagePipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteImagePipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteImagePipelineResponse) GetBody() *DeleteImagePipelineResponseBody {
	return s.Body
}

func (s *DeleteImagePipelineResponse) SetHeaders(v map[string]*string) *DeleteImagePipelineResponse {
	s.Headers = v
	return s
}

func (s *DeleteImagePipelineResponse) SetStatusCode(v int32) *DeleteImagePipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImagePipelineResponse) SetBody(v *DeleteImagePipelineResponseBody) *DeleteImagePipelineResponse {
	s.Body = v
	return s
}

func (s *DeleteImagePipelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
