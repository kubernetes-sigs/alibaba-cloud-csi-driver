// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLifecycleRetrieveJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLifecycleRetrieveJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLifecycleRetrieveJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateLifecycleRetrieveJobResponseBody) *CreateLifecycleRetrieveJobResponse
	GetBody() *CreateLifecycleRetrieveJobResponseBody
}

type CreateLifecycleRetrieveJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLifecycleRetrieveJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLifecycleRetrieveJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLifecycleRetrieveJobResponse) GoString() string {
	return s.String()
}

func (s *CreateLifecycleRetrieveJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLifecycleRetrieveJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLifecycleRetrieveJobResponse) GetBody() *CreateLifecycleRetrieveJobResponseBody {
	return s.Body
}

func (s *CreateLifecycleRetrieveJobResponse) SetHeaders(v map[string]*string) *CreateLifecycleRetrieveJobResponse {
	s.Headers = v
	return s
}

func (s *CreateLifecycleRetrieveJobResponse) SetStatusCode(v int32) *CreateLifecycleRetrieveJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLifecycleRetrieveJobResponse) SetBody(v *CreateLifecycleRetrieveJobResponseBody) *CreateLifecycleRetrieveJobResponse {
	s.Body = v
	return s
}

func (s *CreateLifecycleRetrieveJobResponse) Validate() error {
	return dara.Validate(s)
}
