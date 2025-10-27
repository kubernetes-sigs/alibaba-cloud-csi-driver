// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelLifecycleRetrieveJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelLifecycleRetrieveJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelLifecycleRetrieveJobResponse
	GetStatusCode() *int32
	SetBody(v *CancelLifecycleRetrieveJobResponseBody) *CancelLifecycleRetrieveJobResponse
	GetBody() *CancelLifecycleRetrieveJobResponseBody
}

type CancelLifecycleRetrieveJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelLifecycleRetrieveJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelLifecycleRetrieveJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelLifecycleRetrieveJobResponse) GoString() string {
	return s.String()
}

func (s *CancelLifecycleRetrieveJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelLifecycleRetrieveJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelLifecycleRetrieveJobResponse) GetBody() *CancelLifecycleRetrieveJobResponseBody {
	return s.Body
}

func (s *CancelLifecycleRetrieveJobResponse) SetHeaders(v map[string]*string) *CancelLifecycleRetrieveJobResponse {
	s.Headers = v
	return s
}

func (s *CancelLifecycleRetrieveJobResponse) SetStatusCode(v int32) *CancelLifecycleRetrieveJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelLifecycleRetrieveJobResponse) SetBody(v *CancelLifecycleRetrieveJobResponseBody) *CancelLifecycleRetrieveJobResponse {
	s.Body = v
	return s
}

func (s *CancelLifecycleRetrieveJobResponse) Validate() error {
	return dara.Validate(s)
}
