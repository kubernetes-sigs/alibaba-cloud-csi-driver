// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVSwitchResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVSwitchResponseBody) *DeleteVSwitchResponse
	GetBody() *DeleteVSwitchResponseBody
}

type DeleteVSwitchResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVSwitchResponse) GoString() string {
	return s.String()
}

func (s *DeleteVSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVSwitchResponse) GetBody() *DeleteVSwitchResponseBody {
	return s.Body
}

func (s *DeleteVSwitchResponse) SetHeaders(v map[string]*string) *DeleteVSwitchResponse {
	s.Headers = v
	return s
}

func (s *DeleteVSwitchResponse) SetStatusCode(v int32) *DeleteVSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVSwitchResponse) SetBody(v *DeleteVSwitchResponseBody) *DeleteVSwitchResponse {
	s.Body = v
	return s
}

func (s *DeleteVSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
