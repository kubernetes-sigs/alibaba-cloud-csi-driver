// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskMonitorDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiskMonitorDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiskMonitorDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiskMonitorDataResponseBody) *DescribeDiskMonitorDataResponse
	GetBody() *DescribeDiskMonitorDataResponseBody
}

type DescribeDiskMonitorDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiskMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiskMonitorDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiskMonitorDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiskMonitorDataResponse) GetBody() *DescribeDiskMonitorDataResponseBody {
	return s.Body
}

func (s *DescribeDiskMonitorDataResponse) SetHeaders(v map[string]*string) *DescribeDiskMonitorDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskMonitorDataResponse) SetStatusCode(v int32) *DescribeDiskMonitorDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskMonitorDataResponse) SetBody(v *DescribeDiskMonitorDataResponseBody) *DescribeDiskMonitorDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDiskMonitorDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
