// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLaunchTemplateVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLaunchTemplateVersions(v *DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersions) *DeleteLaunchTemplateVersionResponseBody
	GetLaunchTemplateVersions() *DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersions
	SetRequestId(v string) *DeleteLaunchTemplateVersionResponseBody
	GetRequestId() *string
}

type DeleteLaunchTemplateVersionResponseBody struct {
	// The deleted launch template versions.
	LaunchTemplateVersions *DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersions `json:"LaunchTemplateVersions,omitempty" xml:"LaunchTemplateVersions,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLaunchTemplateVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLaunchTemplateVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLaunchTemplateVersionResponseBody) GetLaunchTemplateVersions() *DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersions {
	return s.LaunchTemplateVersions
}

func (s *DeleteLaunchTemplateVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLaunchTemplateVersionResponseBody) SetLaunchTemplateVersions(v *DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersions) *DeleteLaunchTemplateVersionResponseBody {
	s.LaunchTemplateVersions = v
	return s
}

func (s *DeleteLaunchTemplateVersionResponseBody) SetRequestId(v string) *DeleteLaunchTemplateVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLaunchTemplateVersionResponseBody) Validate() error {
	if s.LaunchTemplateVersions != nil {
		if err := s.LaunchTemplateVersions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersions struct {
	LaunchTemplateVersion []*DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersionsLaunchTemplateVersion `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty" type:"Repeated"`
}

func (s DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersions) String() string {
	return dara.Prettify(s)
}

func (s DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersions) GoString() string {
	return s.String()
}

func (s *DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersions) GetLaunchTemplateVersion() []*DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersionsLaunchTemplateVersion {
	return s.LaunchTemplateVersion
}

func (s *DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersions) SetLaunchTemplateVersion(v []*DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersionsLaunchTemplateVersion) *DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersions {
	s.LaunchTemplateVersion = v
	return s
}

func (s *DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersions) Validate() error {
	if s.LaunchTemplateVersion != nil {
		for _, item := range s.LaunchTemplateVersion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersionsLaunchTemplateVersion struct {
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
}

func (s DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersionsLaunchTemplateVersion) String() string {
	return dara.Prettify(s)
}

func (s DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersionsLaunchTemplateVersion) GoString() string {
	return s.String()
}

func (s *DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersionsLaunchTemplateVersion) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersionsLaunchTemplateVersion) GetLaunchTemplateVersionNumber() *int64 {
	return s.LaunchTemplateVersionNumber
}

func (s *DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersionsLaunchTemplateVersion) SetLaunchTemplateId(v string) *DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersionsLaunchTemplateVersion {
	s.LaunchTemplateId = &v
	return s
}

func (s *DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersionsLaunchTemplateVersion) SetLaunchTemplateVersionNumber(v int64) *DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersionsLaunchTemplateVersion {
	s.LaunchTemplateVersionNumber = &v
	return s
}

func (s *DeleteLaunchTemplateVersionResponseBodyLaunchTemplateVersionsLaunchTemplateVersion) Validate() error {
	return dara.Validate(s)
}
