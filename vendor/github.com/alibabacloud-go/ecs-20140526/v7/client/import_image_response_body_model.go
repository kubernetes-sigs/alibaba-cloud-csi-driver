// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *ImportImageResponseBody
	GetImageId() *string
	SetRegionId(v string) *ImportImageResponseBody
	GetRegionId() *string
	SetRequestId(v string) *ImportImageResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ImportImageResponseBody
	GetTaskId() *string
}

type ImportImageResponseBody struct {
	// The image ID.
	//
	// example:
	//
	// m-bp67acfmxazb4p****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The image import task ID.
	//
	// example:
	//
	// t-bp67acfmxazb4p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImportImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportImageResponseBody) GoString() string {
	return s.String()
}

func (s *ImportImageResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *ImportImageResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *ImportImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportImageResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ImportImageResponseBody) SetImageId(v string) *ImportImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *ImportImageResponseBody) SetRegionId(v string) *ImportImageResponseBody {
	s.RegionId = &v
	return s
}

func (s *ImportImageResponseBody) SetRequestId(v string) *ImportImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportImageResponseBody) SetTaskId(v string) *ImportImageResponseBody {
	s.TaskId = &v
	return s
}

func (s *ImportImageResponseBody) Validate() error {
	return dara.Validate(s)
}
