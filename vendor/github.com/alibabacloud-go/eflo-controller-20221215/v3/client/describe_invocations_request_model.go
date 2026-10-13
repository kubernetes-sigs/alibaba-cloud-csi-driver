// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvocationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentEncoding(v string) *DescribeInvocationsRequest
	GetContentEncoding() *string
	SetIncludeOutput(v bool) *DescribeInvocationsRequest
	GetIncludeOutput() *bool
	SetInvokeId(v string) *DescribeInvocationsRequest
	GetInvokeId() *string
	SetNodeId(v string) *DescribeInvocationsRequest
	GetNodeId() *string
}

type DescribeInvocationsRequest struct {
	// The encoding mode of the `CommandContent` and `Output` fields in the response. Valid values:
	//
	// - PlainText: Returns the original command content and output.
	//
	// - Base64 (default): Returns the Base64-encoded command content and output.
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// Specifies whether to include the command output in the response.
	//
	// - true: Returns the output. You must specify the `InvokeId` or `NodeId` parameter.
	//
	// - false (default): Does not return the output.
	//
	// example:
	//
	// true
	IncludeOutput *bool `json:"IncludeOutput,omitempty" xml:"IncludeOutput,omitempty"`
	// The ID of the command execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-cd03crwys0lrls0
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// e01-cn-zvp2tgykr08
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeInvocationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *DescribeInvocationsRequest) GetIncludeOutput() *bool {
	return s.IncludeOutput
}

func (s *DescribeInvocationsRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeInvocationsRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeInvocationsRequest) SetContentEncoding(v string) *DescribeInvocationsRequest {
	s.ContentEncoding = &v
	return s
}

func (s *DescribeInvocationsRequest) SetIncludeOutput(v bool) *DescribeInvocationsRequest {
	s.IncludeOutput = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInvokeId(v string) *DescribeInvocationsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetNodeId(v string) *DescribeInvocationsRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeInvocationsRequest) Validate() error {
	return dara.Validate(s)
}
