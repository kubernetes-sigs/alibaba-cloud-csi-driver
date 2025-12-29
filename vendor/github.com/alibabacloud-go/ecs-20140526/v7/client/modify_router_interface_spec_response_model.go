// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRouterInterfaceSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRouterInterfaceSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRouterInterfaceSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRouterInterfaceSpecResponseBody) *ModifyRouterInterfaceSpecResponse
	GetBody() *ModifyRouterInterfaceSpecResponseBody
}

type ModifyRouterInterfaceSpecResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRouterInterfaceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRouterInterfaceSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRouterInterfaceSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyRouterInterfaceSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRouterInterfaceSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRouterInterfaceSpecResponse) GetBody() *ModifyRouterInterfaceSpecResponseBody {
	return s.Body
}

func (s *ModifyRouterInterfaceSpecResponse) SetHeaders(v map[string]*string) *ModifyRouterInterfaceSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyRouterInterfaceSpecResponse) SetStatusCode(v int32) *ModifyRouterInterfaceSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRouterInterfaceSpecResponse) SetBody(v *ModifyRouterInterfaceSpecResponseBody) *ModifyRouterInterfaceSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyRouterInterfaceSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
