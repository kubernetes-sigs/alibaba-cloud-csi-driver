// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecycleBinRestoreJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRecycleBinRestoreJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRecycleBinRestoreJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateRecycleBinRestoreJobResponseBody) *CreateRecycleBinRestoreJobResponse
	GetBody() *CreateRecycleBinRestoreJobResponseBody
}

type CreateRecycleBinRestoreJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRecycleBinRestoreJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRecycleBinRestoreJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRecycleBinRestoreJobResponse) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinRestoreJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRecycleBinRestoreJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRecycleBinRestoreJobResponse) GetBody() *CreateRecycleBinRestoreJobResponseBody {
	return s.Body
}

func (s *CreateRecycleBinRestoreJobResponse) SetHeaders(v map[string]*string) *CreateRecycleBinRestoreJobResponse {
	s.Headers = v
	return s
}

func (s *CreateRecycleBinRestoreJobResponse) SetStatusCode(v int32) *CreateRecycleBinRestoreJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRecycleBinRestoreJobResponse) SetBody(v *CreateRecycleBinRestoreJobResponseBody) *CreateRecycleBinRestoreJobResponse {
	s.Body = v
	return s
}

func (s *CreateRecycleBinRestoreJobResponse) Validate() error {
	return dara.Validate(s)
}
