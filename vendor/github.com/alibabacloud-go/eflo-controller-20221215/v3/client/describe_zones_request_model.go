// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeZonesRequest
	GetAcceptLanguage() *string
}

type DescribeZonesRequest struct {
	// The natural language that is used to filter responses. For more information, see RFC 7231. Valid values:
	//
	// zh-CN en-US Default value: zh-CN.
	//
	// Valid values:
	//
	// 	- en-US
	//
	// 	- zh-CN
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeZonesRequest) SetAcceptLanguage(v string) *DescribeZonesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeZonesRequest) Validate() error {
	return dara.Validate(s)
}
