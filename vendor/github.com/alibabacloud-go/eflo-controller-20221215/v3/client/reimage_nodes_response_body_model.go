// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReimageNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReimageNodesResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ReimageNodesResponseBody
	GetTaskId() *string
}

type ReimageNodesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 15FBCD9B-C93F-54E8-A168-AADE7E66DAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// i158782151663841517926
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ReimageNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReimageNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ReimageNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReimageNodesResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ReimageNodesResponseBody) SetRequestId(v string) *ReimageNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReimageNodesResponseBody) SetTaskId(v string) *ReimageNodesResponseBody {
	s.TaskId = &v
	return s
}

func (s *ReimageNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
