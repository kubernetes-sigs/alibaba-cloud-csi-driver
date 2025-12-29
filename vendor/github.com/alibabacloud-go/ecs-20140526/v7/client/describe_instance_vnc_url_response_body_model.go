// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceVncUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeInstanceVncUrlResponseBody
	GetRequestId() *string
	SetVncUrl(v string) *DescribeInstanceVncUrlResponseBody
	GetVncUrl() *string
}

type DescribeInstanceVncUrlResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The VNC logon address.
	//
	// >  The VNC logon address returned is valid only for 15 seconds. If a connection is not established within 15 seconds after a successful call, the VNC logon address expires and you must call the DescribeInstanceVncUrl operation to obtain a new logon address.
	//
	// example:
	//
	// wss%3A%2F%2Fhz01-vncproxy.aliyun.com%2Fwebsockify%2F%3Fs%3DDvh%252FIA%252BYc73gWO48cBx2gBxUDVzaAnSKr74pq30mzqUYgeUMcB%252FbkNixDxdEA996
	VncUrl *string `json:"VncUrl,omitempty" xml:"VncUrl,omitempty"`
}

func (s DescribeInstanceVncUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceVncUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceVncUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceVncUrlResponseBody) GetVncUrl() *string {
	return s.VncUrl
}

func (s *DescribeInstanceVncUrlResponseBody) SetRequestId(v string) *DescribeInstanceVncUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceVncUrlResponseBody) SetVncUrl(v string) *DescribeInstanceVncUrlResponseBody {
	s.VncUrl = &v
	return s
}

func (s *DescribeInstanceVncUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
