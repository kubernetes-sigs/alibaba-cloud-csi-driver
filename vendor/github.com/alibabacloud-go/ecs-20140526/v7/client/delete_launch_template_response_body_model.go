// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLaunchTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLaunchTemplateId(v string) *DeleteLaunchTemplateResponseBody
	GetLaunchTemplateId() *string
	SetLaunchTemplateVersionNumbers(v *DeleteLaunchTemplateResponseBodyLaunchTemplateVersionNumbers) *DeleteLaunchTemplateResponseBody
	GetLaunchTemplateVersionNumbers() *DeleteLaunchTemplateResponseBodyLaunchTemplateVersionNumbers
	SetRequestId(v string) *DeleteLaunchTemplateResponseBody
	GetRequestId() *string
}

type DeleteLaunchTemplateResponseBody struct {
	// The ID of the launch template. For more information, see [DescribeLaunchTemplates](https://help.aliyun.com/document_detail/73759.html).
	//
	// You must specify `LaunchTemplateId` or `LaunchTemplateName` to specify a launch template.
	//
	// example:
	//
	// lt-bp1apo0bbbkuy0rj****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The versions of the deleted launch template.
	LaunchTemplateVersionNumbers *DeleteLaunchTemplateResponseBodyLaunchTemplateVersionNumbers `json:"LaunchTemplateVersionNumbers,omitempty" xml:"LaunchTemplateVersionNumbers,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLaunchTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLaunchTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLaunchTemplateResponseBody) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *DeleteLaunchTemplateResponseBody) GetLaunchTemplateVersionNumbers() *DeleteLaunchTemplateResponseBodyLaunchTemplateVersionNumbers {
	return s.LaunchTemplateVersionNumbers
}

func (s *DeleteLaunchTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLaunchTemplateResponseBody) SetLaunchTemplateId(v string) *DeleteLaunchTemplateResponseBody {
	s.LaunchTemplateId = &v
	return s
}

func (s *DeleteLaunchTemplateResponseBody) SetLaunchTemplateVersionNumbers(v *DeleteLaunchTemplateResponseBodyLaunchTemplateVersionNumbers) *DeleteLaunchTemplateResponseBody {
	s.LaunchTemplateVersionNumbers = v
	return s
}

func (s *DeleteLaunchTemplateResponseBody) SetRequestId(v string) *DeleteLaunchTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLaunchTemplateResponseBody) Validate() error {
	if s.LaunchTemplateVersionNumbers != nil {
		if err := s.LaunchTemplateVersionNumbers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteLaunchTemplateResponseBodyLaunchTemplateVersionNumbers struct {
	VersionNumbers []*int64 `json:"versionNumbers,omitempty" xml:"versionNumbers,omitempty" type:"Repeated"`
}

func (s DeleteLaunchTemplateResponseBodyLaunchTemplateVersionNumbers) String() string {
	return dara.Prettify(s)
}

func (s DeleteLaunchTemplateResponseBodyLaunchTemplateVersionNumbers) GoString() string {
	return s.String()
}

func (s *DeleteLaunchTemplateResponseBodyLaunchTemplateVersionNumbers) GetVersionNumbers() []*int64 {
	return s.VersionNumbers
}

func (s *DeleteLaunchTemplateResponseBodyLaunchTemplateVersionNumbers) SetVersionNumbers(v []*int64) *DeleteLaunchTemplateResponseBodyLaunchTemplateVersionNumbers {
	s.VersionNumbers = v
	return s
}

func (s *DeleteLaunchTemplateResponseBodyLaunchTemplateVersionNumbers) Validate() error {
	return dara.Validate(s)
}
