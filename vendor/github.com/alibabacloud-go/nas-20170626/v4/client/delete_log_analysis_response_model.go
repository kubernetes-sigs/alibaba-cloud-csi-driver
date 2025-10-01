// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLogAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLogAnalysisResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLogAnalysisResponseBody) *DeleteLogAnalysisResponse
	GetBody() *DeleteLogAnalysisResponseBody
}

type DeleteLogAnalysisResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLogAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLogAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogAnalysisResponse) GoString() string {
	return s.String()
}

func (s *DeleteLogAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLogAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLogAnalysisResponse) GetBody() *DeleteLogAnalysisResponseBody {
	return s.Body
}

func (s *DeleteLogAnalysisResponse) SetHeaders(v map[string]*string) *DeleteLogAnalysisResponse {
	s.Headers = v
	return s
}

func (s *DeleteLogAnalysisResponse) SetStatusCode(v int32) *DeleteLogAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLogAnalysisResponse) SetBody(v *DeleteLogAnalysisResponseBody) *DeleteLogAnalysisResponse {
	s.Body = v
	return s
}

func (s *DeleteLogAnalysisResponse) Validate() error {
	return dara.Validate(s)
}
