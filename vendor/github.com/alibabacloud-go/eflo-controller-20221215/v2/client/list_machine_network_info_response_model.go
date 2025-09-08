// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMachineNetworkInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMachineNetworkInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMachineNetworkInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListMachineNetworkInfoResponseBody) *ListMachineNetworkInfoResponse
	GetBody() *ListMachineNetworkInfoResponseBody
}

type ListMachineNetworkInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMachineNetworkInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMachineNetworkInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMachineNetworkInfoResponse) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMachineNetworkInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMachineNetworkInfoResponse) GetBody() *ListMachineNetworkInfoResponseBody {
	return s.Body
}

func (s *ListMachineNetworkInfoResponse) SetHeaders(v map[string]*string) *ListMachineNetworkInfoResponse {
	s.Headers = v
	return s
}

func (s *ListMachineNetworkInfoResponse) SetStatusCode(v int32) *ListMachineNetworkInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMachineNetworkInfoResponse) SetBody(v *ListMachineNetworkInfoResponseBody) *ListMachineNetworkInfoResponse {
	s.Body = v
	return s
}

func (s *ListMachineNetworkInfoResponse) Validate() error {
	return dara.Validate(s)
}
