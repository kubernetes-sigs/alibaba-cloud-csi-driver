// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMachineTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMachineTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMachineTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListMachineTypesResponseBody) *ListMachineTypesResponse
	GetBody() *ListMachineTypesResponseBody
}

type ListMachineTypesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMachineTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMachineTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMachineTypesResponse) GoString() string {
	return s.String()
}

func (s *ListMachineTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMachineTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMachineTypesResponse) GetBody() *ListMachineTypesResponseBody {
	return s.Body
}

func (s *ListMachineTypesResponse) SetHeaders(v map[string]*string) *ListMachineTypesResponse {
	s.Headers = v
	return s
}

func (s *ListMachineTypesResponse) SetStatusCode(v int32) *ListMachineTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMachineTypesResponse) SetBody(v *ListMachineTypesResponseBody) *ListMachineTypesResponse {
	s.Body = v
	return s
}

func (s *ListMachineTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
