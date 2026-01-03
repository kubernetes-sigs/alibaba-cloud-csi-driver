// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLaunchTemplateVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLaunchTemplateId(v string) *CreateLaunchTemplateVersionResponseBody
	GetLaunchTemplateId() *string
	SetLaunchTemplateVersionNumber(v int64) *CreateLaunchTemplateVersionResponseBody
	GetLaunchTemplateVersionNumber() *int64
	SetRequestId(v string) *CreateLaunchTemplateVersionResponseBody
	GetRequestId() *string
}

type CreateLaunchTemplateVersionResponseBody struct {
	// The ID of the launch template. For more information, see [DescribeLaunchTemplates](https://help.aliyun.com/document_detail/73759.html).
	//
	// You must specify `LaunchTemplateId` or `LaunchTemplateName` to specify a launch template.
	//
	// example:
	//
	// lt-bp1apo0bbbkuy0rj****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The number of the created version of the launch template.
	//
	// example:
	//
	// 2
	LaunchTemplateVersionNumber *int64 `json:"LaunchTemplateVersionNumber,omitempty" xml:"LaunchTemplateVersionNumber,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DExxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLaunchTemplateVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateVersionResponseBody) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *CreateLaunchTemplateVersionResponseBody) GetLaunchTemplateVersionNumber() *int64 {
	return s.LaunchTemplateVersionNumber
}

func (s *CreateLaunchTemplateVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLaunchTemplateVersionResponseBody) SetLaunchTemplateId(v string) *CreateLaunchTemplateVersionResponseBody {
	s.LaunchTemplateId = &v
	return s
}

func (s *CreateLaunchTemplateVersionResponseBody) SetLaunchTemplateVersionNumber(v int64) *CreateLaunchTemplateVersionResponseBody {
	s.LaunchTemplateVersionNumber = &v
	return s
}

func (s *CreateLaunchTemplateVersionResponseBody) SetRequestId(v string) *CreateLaunchTemplateVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLaunchTemplateVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
