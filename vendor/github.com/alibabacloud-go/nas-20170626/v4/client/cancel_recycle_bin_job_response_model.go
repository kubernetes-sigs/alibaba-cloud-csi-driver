// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRecycleBinJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelRecycleBinJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelRecycleBinJobResponse
	GetStatusCode() *int32
	SetBody(v *CancelRecycleBinJobResponseBody) *CancelRecycleBinJobResponse
	GetBody() *CancelRecycleBinJobResponseBody
}

type CancelRecycleBinJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelRecycleBinJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelRecycleBinJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelRecycleBinJobResponse) GoString() string {
	return s.String()
}

func (s *CancelRecycleBinJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelRecycleBinJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelRecycleBinJobResponse) GetBody() *CancelRecycleBinJobResponseBody {
	return s.Body
}

func (s *CancelRecycleBinJobResponse) SetHeaders(v map[string]*string) *CancelRecycleBinJobResponse {
	s.Headers = v
	return s
}

func (s *CancelRecycleBinJobResponse) SetStatusCode(v int32) *CancelRecycleBinJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelRecycleBinJobResponse) SetBody(v *CancelRecycleBinJobResponseBody) *CancelRecycleBinJobResponse {
	s.Body = v
	return s
}

func (s *CancelRecycleBinJobResponse) Validate() error {
	return dara.Validate(s)
}
