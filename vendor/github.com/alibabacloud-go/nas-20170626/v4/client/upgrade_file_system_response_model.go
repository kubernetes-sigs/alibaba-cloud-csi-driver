// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeFileSystemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeFileSystemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeFileSystemResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeFileSystemResponseBody) *UpgradeFileSystemResponse
	GetBody() *UpgradeFileSystemResponseBody
}

type UpgradeFileSystemResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeFileSystemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeFileSystemResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeFileSystemResponse) GoString() string {
	return s.String()
}

func (s *UpgradeFileSystemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeFileSystemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeFileSystemResponse) GetBody() *UpgradeFileSystemResponseBody {
	return s.Body
}

func (s *UpgradeFileSystemResponse) SetHeaders(v map[string]*string) *UpgradeFileSystemResponse {
	s.Headers = v
	return s
}

func (s *UpgradeFileSystemResponse) SetStatusCode(v int32) *UpgradeFileSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeFileSystemResponse) SetBody(v *UpgradeFileSystemResponseBody) *UpgradeFileSystemResponse {
	s.Body = v
	return s
}

func (s *UpgradeFileSystemResponse) Validate() error {
	return dara.Validate(s)
}
