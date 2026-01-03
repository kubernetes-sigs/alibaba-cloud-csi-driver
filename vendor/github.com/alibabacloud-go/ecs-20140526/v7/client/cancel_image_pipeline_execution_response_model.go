// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelImagePipelineExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelImagePipelineExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelImagePipelineExecutionResponse
	GetStatusCode() *int32
	SetBody(v *CancelImagePipelineExecutionResponseBody) *CancelImagePipelineExecutionResponse
	GetBody() *CancelImagePipelineExecutionResponseBody
}

type CancelImagePipelineExecutionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelImagePipelineExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelImagePipelineExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelImagePipelineExecutionResponse) GoString() string {
	return s.String()
}

func (s *CancelImagePipelineExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelImagePipelineExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelImagePipelineExecutionResponse) GetBody() *CancelImagePipelineExecutionResponseBody {
	return s.Body
}

func (s *CancelImagePipelineExecutionResponse) SetHeaders(v map[string]*string) *CancelImagePipelineExecutionResponse {
	s.Headers = v
	return s
}

func (s *CancelImagePipelineExecutionResponse) SetStatusCode(v int32) *CancelImagePipelineExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelImagePipelineExecutionResponse) SetBody(v *CancelImagePipelineExecutionResponseBody) *CancelImagePipelineExecutionResponse {
	s.Body = v
	return s
}

func (s *CancelImagePipelineExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
