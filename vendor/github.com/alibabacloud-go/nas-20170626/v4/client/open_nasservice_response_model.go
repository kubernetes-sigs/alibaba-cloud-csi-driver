// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenNASServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenNASServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenNASServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenNASServiceResponseBody) *OpenNASServiceResponse
	GetBody() *OpenNASServiceResponseBody
}

type OpenNASServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenNASServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenNASServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenNASServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenNASServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenNASServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenNASServiceResponse) GetBody() *OpenNASServiceResponseBody {
	return s.Body
}

func (s *OpenNASServiceResponse) SetHeaders(v map[string]*string) *OpenNASServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenNASServiceResponse) SetStatusCode(v int32) *OpenNASServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenNASServiceResponse) SetBody(v *OpenNASServiceResponseBody) *OpenNASServiceResponse {
	s.Body = v
	return s
}

func (s *OpenNASServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
