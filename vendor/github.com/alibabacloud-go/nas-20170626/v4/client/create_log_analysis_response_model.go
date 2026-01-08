// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLogAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLogAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *CreateLogAnalysisResponseBody) *CreateLogAnalysisResponse
	GetBody() *CreateLogAnalysisResponseBody
}

type CreateLogAnalysisResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLogAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLogAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLogAnalysisResponse) GoString() string {
	return s.String()
}

func (s *CreateLogAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLogAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLogAnalysisResponse) GetBody() *CreateLogAnalysisResponseBody {
	return s.Body
}

func (s *CreateLogAnalysisResponse) SetHeaders(v map[string]*string) *CreateLogAnalysisResponse {
	s.Headers = v
	return s
}

func (s *CreateLogAnalysisResponse) SetStatusCode(v int32) *CreateLogAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLogAnalysisResponse) SetBody(v *CreateLogAnalysisResponseBody) *CreateLogAnalysisResponse {
	s.Body = v
	return s
}

func (s *CreateLogAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
