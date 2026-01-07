// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoProvisioningGroupHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutoProvisioningGroupHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutoProvisioningGroupHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutoProvisioningGroupHistoryResponseBody) *DescribeAutoProvisioningGroupHistoryResponse
	GetBody() *DescribeAutoProvisioningGroupHistoryResponseBody
}

type DescribeAutoProvisioningGroupHistoryResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoProvisioningGroupHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutoProvisioningGroupHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutoProvisioningGroupHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutoProvisioningGroupHistoryResponse) GetBody() *DescribeAutoProvisioningGroupHistoryResponseBody {
	return s.Body
}

func (s *DescribeAutoProvisioningGroupHistoryResponse) SetHeaders(v map[string]*string) *DescribeAutoProvisioningGroupHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponse) SetStatusCode(v int32) *DescribeAutoProvisioningGroupHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponse) SetBody(v *DescribeAutoProvisioningGroupHistoryResponseBody) *DescribeAutoProvisioningGroupHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
