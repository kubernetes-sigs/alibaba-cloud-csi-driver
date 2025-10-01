// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyDataFlowAutoRefreshResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyDataFlowAutoRefreshResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyDataFlowAutoRefreshResponse
	GetStatusCode() *int32
	SetBody(v *ApplyDataFlowAutoRefreshResponseBody) *ApplyDataFlowAutoRefreshResponse
	GetBody() *ApplyDataFlowAutoRefreshResponseBody
}

type ApplyDataFlowAutoRefreshResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyDataFlowAutoRefreshResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyDataFlowAutoRefreshResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyDataFlowAutoRefreshResponse) GoString() string {
	return s.String()
}

func (s *ApplyDataFlowAutoRefreshResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyDataFlowAutoRefreshResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyDataFlowAutoRefreshResponse) GetBody() *ApplyDataFlowAutoRefreshResponseBody {
	return s.Body
}

func (s *ApplyDataFlowAutoRefreshResponse) SetHeaders(v map[string]*string) *ApplyDataFlowAutoRefreshResponse {
	s.Headers = v
	return s
}

func (s *ApplyDataFlowAutoRefreshResponse) SetStatusCode(v int32) *ApplyDataFlowAutoRefreshResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshResponse) SetBody(v *ApplyDataFlowAutoRefreshResponseBody) *ApplyDataFlowAutoRefreshResponse {
	s.Body = v
	return s
}

func (s *ApplyDataFlowAutoRefreshResponse) Validate() error {
	return dara.Validate(s)
}
