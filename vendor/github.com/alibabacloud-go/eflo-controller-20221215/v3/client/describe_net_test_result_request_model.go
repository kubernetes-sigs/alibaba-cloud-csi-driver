// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetTestResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTestId(v string) *DescribeNetTestResultRequest
	GetTestId() *string
}

type DescribeNetTestResultRequest struct {
	// The ID of the test task. The unique identifier of a network test task.
	//
	// example:
	//
	// dr-uf6i0tv2refv8wz*****
	TestId *string `json:"TestId,omitempty" xml:"TestId,omitempty"`
}

func (s DescribeNetTestResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetTestResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultRequest) GetTestId() *string {
	return s.TestId
}

func (s *DescribeNetTestResultRequest) SetTestId(v string) *DescribeNetTestResultRequest {
	s.TestId = &v
	return s
}

func (s *DescribeNetTestResultRequest) Validate() error {
	return dara.Validate(s)
}
