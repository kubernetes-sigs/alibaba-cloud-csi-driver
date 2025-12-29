// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartImagePipelineExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartImagePipelineExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartImagePipelineExecutionResponse
	GetStatusCode() *int32
	SetBody(v *StartImagePipelineExecutionResponseBody) *StartImagePipelineExecutionResponse
	GetBody() *StartImagePipelineExecutionResponseBody
}

type StartImagePipelineExecutionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartImagePipelineExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartImagePipelineExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s StartImagePipelineExecutionResponse) GoString() string {
	return s.String()
}

func (s *StartImagePipelineExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartImagePipelineExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartImagePipelineExecutionResponse) GetBody() *StartImagePipelineExecutionResponseBody {
	return s.Body
}

func (s *StartImagePipelineExecutionResponse) SetHeaders(v map[string]*string) *StartImagePipelineExecutionResponse {
	s.Headers = v
	return s
}

func (s *StartImagePipelineExecutionResponse) SetStatusCode(v int32) *StartImagePipelineExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *StartImagePipelineExecutionResponse) SetBody(v *StartImagePipelineExecutionResponseBody) *StartImagePipelineExecutionResponse {
	s.Body = v
	return s
}

func (s *StartImagePipelineExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
