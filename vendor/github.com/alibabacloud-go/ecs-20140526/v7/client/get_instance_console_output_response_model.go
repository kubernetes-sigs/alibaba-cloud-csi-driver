// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceConsoleOutputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceConsoleOutputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceConsoleOutputResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceConsoleOutputResponseBody) *GetInstanceConsoleOutputResponse
	GetBody() *GetInstanceConsoleOutputResponseBody
}

type GetInstanceConsoleOutputResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceConsoleOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceConsoleOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceConsoleOutputResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceConsoleOutputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceConsoleOutputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceConsoleOutputResponse) GetBody() *GetInstanceConsoleOutputResponseBody {
	return s.Body
}

func (s *GetInstanceConsoleOutputResponse) SetHeaders(v map[string]*string) *GetInstanceConsoleOutputResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceConsoleOutputResponse) SetStatusCode(v int32) *GetInstanceConsoleOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceConsoleOutputResponse) SetBody(v *GetInstanceConsoleOutputResponseBody) *GetInstanceConsoleOutputResponse {
	s.Body = v
	return s
}

func (s *GetInstanceConsoleOutputResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
