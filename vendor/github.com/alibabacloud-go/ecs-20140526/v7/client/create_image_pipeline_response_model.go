// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImagePipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateImagePipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateImagePipelineResponse
	GetStatusCode() *int32
	SetBody(v *CreateImagePipelineResponseBody) *CreateImagePipelineResponse
	GetBody() *CreateImagePipelineResponseBody
}

type CreateImagePipelineResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateImagePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateImagePipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateImagePipelineResponse) GoString() string {
	return s.String()
}

func (s *CreateImagePipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateImagePipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateImagePipelineResponse) GetBody() *CreateImagePipelineResponseBody {
	return s.Body
}

func (s *CreateImagePipelineResponse) SetHeaders(v map[string]*string) *CreateImagePipelineResponse {
	s.Headers = v
	return s
}

func (s *CreateImagePipelineResponse) SetStatusCode(v int32) *CreateImagePipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImagePipelineResponse) SetBody(v *CreateImagePipelineResponseBody) *CreateImagePipelineResponse {
	s.Body = v
	return s
}

func (s *CreateImagePipelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
