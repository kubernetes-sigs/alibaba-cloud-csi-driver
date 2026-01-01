// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataFlowAutoRefreshResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDataFlowAutoRefreshResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDataFlowAutoRefreshResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDataFlowAutoRefreshResponseBody) *ModifyDataFlowAutoRefreshResponse
	GetBody() *ModifyDataFlowAutoRefreshResponseBody
}

type ModifyDataFlowAutoRefreshResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDataFlowAutoRefreshResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDataFlowAutoRefreshResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataFlowAutoRefreshResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataFlowAutoRefreshResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDataFlowAutoRefreshResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDataFlowAutoRefreshResponse) GetBody() *ModifyDataFlowAutoRefreshResponseBody {
	return s.Body
}

func (s *ModifyDataFlowAutoRefreshResponse) SetHeaders(v map[string]*string) *ModifyDataFlowAutoRefreshResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataFlowAutoRefreshResponse) SetStatusCode(v int32) *ModifyDataFlowAutoRefreshResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshResponse) SetBody(v *ModifyDataFlowAutoRefreshResponseBody) *ModifyDataFlowAutoRefreshResponse {
	s.Body = v
	return s
}

func (s *ModifyDataFlowAutoRefreshResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
