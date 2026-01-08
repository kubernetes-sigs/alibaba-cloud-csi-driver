// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecycleBinDeleteJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRecycleBinDeleteJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRecycleBinDeleteJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateRecycleBinDeleteJobResponseBody) *CreateRecycleBinDeleteJobResponse
	GetBody() *CreateRecycleBinDeleteJobResponseBody
}

type CreateRecycleBinDeleteJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRecycleBinDeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRecycleBinDeleteJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRecycleBinDeleteJobResponse) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinDeleteJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRecycleBinDeleteJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRecycleBinDeleteJobResponse) GetBody() *CreateRecycleBinDeleteJobResponseBody {
	return s.Body
}

func (s *CreateRecycleBinDeleteJobResponse) SetHeaders(v map[string]*string) *CreateRecycleBinDeleteJobResponse {
	s.Headers = v
	return s
}

func (s *CreateRecycleBinDeleteJobResponse) SetStatusCode(v int32) *CreateRecycleBinDeleteJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecycleBinDeleteJobResponse) SetBody(v *CreateRecycleBinDeleteJobResponseBody) *CreateRecycleBinDeleteJobResponse {
	s.Body = v
	return s
}

func (s *CreateRecycleBinDeleteJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
