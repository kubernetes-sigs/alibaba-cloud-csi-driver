// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFilesystemsAssociatedHpnZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFilesystemsAssociatedHpnZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFilesystemsAssociatedHpnZonesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFilesystemsAssociatedHpnZonesResponseBody) *DescribeFilesystemsAssociatedHpnZonesResponse
	GetBody() *DescribeFilesystemsAssociatedHpnZonesResponseBody
}

type DescribeFilesystemsAssociatedHpnZonesResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFilesystemsAssociatedHpnZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFilesystemsAssociatedHpnZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesystemsAssociatedHpnZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeFilesystemsAssociatedHpnZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFilesystemsAssociatedHpnZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFilesystemsAssociatedHpnZonesResponse) GetBody() *DescribeFilesystemsAssociatedHpnZonesResponseBody {
	return s.Body
}

func (s *DescribeFilesystemsAssociatedHpnZonesResponse) SetHeaders(v map[string]*string) *DescribeFilesystemsAssociatedHpnZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeFilesystemsAssociatedHpnZonesResponse) SetStatusCode(v int32) *DescribeFilesystemsAssociatedHpnZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFilesystemsAssociatedHpnZonesResponse) SetBody(v *DescribeFilesystemsAssociatedHpnZonesResponseBody) *DescribeFilesystemsAssociatedHpnZonesResponse {
	s.Body = v
	return s
}

func (s *DescribeFilesystemsAssociatedHpnZonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
