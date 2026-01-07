// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLaunchTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLaunchTemplateId(v string) *CreateLaunchTemplateResponseBody
	GetLaunchTemplateId() *string
	SetLaunchTemplateVersionNumber(v int64) *CreateLaunchTemplateResponseBody
	GetLaunchTemplateVersionNumber() *int64
	SetRequestId(v string) *CreateLaunchTemplateResponseBody
	GetRequestId() *string
}

type CreateLaunchTemplateResponseBody struct {
	// The ID of the launch template.
	//
	// example:
	//
	// lt-m5eiaupmvm2op9d****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The version number of the launch template.
	//
	// example:
	//
	// 2
	LaunchTemplateVersionNumber *int64 `json:"LaunchTemplateVersionNumber,omitempty" xml:"LaunchTemplateVersionNumber,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLaunchTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateResponseBody) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *CreateLaunchTemplateResponseBody) GetLaunchTemplateVersionNumber() *int64 {
	return s.LaunchTemplateVersionNumber
}

func (s *CreateLaunchTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLaunchTemplateResponseBody) SetLaunchTemplateId(v string) *CreateLaunchTemplateResponseBody {
	s.LaunchTemplateId = &v
	return s
}

func (s *CreateLaunchTemplateResponseBody) SetLaunchTemplateVersionNumber(v int64) *CreateLaunchTemplateResponseBody {
	s.LaunchTemplateVersionNumber = &v
	return s
}

func (s *CreateLaunchTemplateResponseBody) SetRequestId(v string) *CreateLaunchTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLaunchTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
