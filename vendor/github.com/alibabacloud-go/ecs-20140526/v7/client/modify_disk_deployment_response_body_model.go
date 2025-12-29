// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskDeploymentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDiskDeploymentResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ModifyDiskDeploymentResponseBody
	GetTaskId() *string
}

type ModifyDiskDeploymentResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D69846D9-F17F-51C0-8AC6-B4B71777****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the disk migration task.
	//
	// example:
	//
	// t-bp67acfmxazb4p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyDiskDeploymentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskDeploymentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDiskDeploymentResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyDiskDeploymentResponseBody) SetRequestId(v string) *ModifyDiskDeploymentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDiskDeploymentResponseBody) SetTaskId(v string) *ModifyDiskDeploymentResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyDiskDeploymentResponseBody) Validate() error {
	return dara.Validate(s)
}
