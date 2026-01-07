// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisksFullStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDisksFullStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDisksFullStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDisksFullStatusResponseBody) *DescribeDisksFullStatusResponse
	GetBody() *DescribeDisksFullStatusResponseBody
}

type DescribeDisksFullStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDisksFullStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDisksFullStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksFullStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDisksFullStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDisksFullStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDisksFullStatusResponse) GetBody() *DescribeDisksFullStatusResponseBody {
	return s.Body
}

func (s *DescribeDisksFullStatusResponse) SetHeaders(v map[string]*string) *DescribeDisksFullStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDisksFullStatusResponse) SetStatusCode(v int32) *DescribeDisksFullStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDisksFullStatusResponse) SetBody(v *DescribeDisksFullStatusResponseBody) *DescribeDisksFullStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeDisksFullStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
