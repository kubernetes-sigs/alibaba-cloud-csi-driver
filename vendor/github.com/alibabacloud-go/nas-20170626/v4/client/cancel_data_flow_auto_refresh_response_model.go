// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDataFlowAutoRefreshResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelDataFlowAutoRefreshResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelDataFlowAutoRefreshResponse
	GetStatusCode() *int32
	SetBody(v *CancelDataFlowAutoRefreshResponseBody) *CancelDataFlowAutoRefreshResponse
	GetBody() *CancelDataFlowAutoRefreshResponseBody
}

type CancelDataFlowAutoRefreshResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelDataFlowAutoRefreshResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelDataFlowAutoRefreshResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelDataFlowAutoRefreshResponse) GoString() string {
	return s.String()
}

func (s *CancelDataFlowAutoRefreshResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelDataFlowAutoRefreshResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelDataFlowAutoRefreshResponse) GetBody() *CancelDataFlowAutoRefreshResponseBody {
	return s.Body
}

func (s *CancelDataFlowAutoRefreshResponse) SetHeaders(v map[string]*string) *CancelDataFlowAutoRefreshResponse {
	s.Headers = v
	return s
}

func (s *CancelDataFlowAutoRefreshResponse) SetStatusCode(v int32) *CancelDataFlowAutoRefreshResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelDataFlowAutoRefreshResponse) SetBody(v *CancelDataFlowAutoRefreshResponseBody) *CancelDataFlowAutoRefreshResponse {
	s.Body = v
	return s
}

func (s *CancelDataFlowAutoRefreshResponse) Validate() error {
	return dara.Validate(s)
}
