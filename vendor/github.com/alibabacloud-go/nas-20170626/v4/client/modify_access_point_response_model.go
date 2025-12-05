// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccessPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAccessPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAccessPointResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAccessPointResponseBody) *ModifyAccessPointResponse
	GetBody() *ModifyAccessPointResponseBody
}

type ModifyAccessPointResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAccessPointResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccessPointResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccessPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAccessPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAccessPointResponse) GetBody() *ModifyAccessPointResponseBody {
	return s.Body
}

func (s *ModifyAccessPointResponse) SetHeaders(v map[string]*string) *ModifyAccessPointResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccessPointResponse) SetStatusCode(v int32) *ModifyAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAccessPointResponse) SetBody(v *ModifyAccessPointResponseBody) *ModifyAccessPointResponse {
	s.Body = v
	return s
}

func (s *ModifyAccessPointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
