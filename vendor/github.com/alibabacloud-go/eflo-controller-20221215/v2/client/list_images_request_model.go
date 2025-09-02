// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitecture(v string) *ListImagesRequest
	GetArchitecture() *string
	SetImageVersion(v string) *ListImagesRequest
	GetImageVersion() *string
	SetPlatform(v string) *ListImagesRequest
	GetPlatform() *string
}

type ListImagesRequest struct {
	// The architecture.
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The image version.
	//
	// example:
	//
	// 7.9
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// The platform.
	//
	// example:
	//
	// ALinux3
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s ListImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) GetArchitecture() *string {
	return s.Architecture
}

func (s *ListImagesRequest) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *ListImagesRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ListImagesRequest) SetArchitecture(v string) *ListImagesRequest {
	s.Architecture = &v
	return s
}

func (s *ListImagesRequest) SetImageVersion(v string) *ListImagesRequest {
	s.ImageVersion = &v
	return s
}

func (s *ListImagesRequest) SetPlatform(v string) *ListImagesRequest {
	s.Platform = &v
	return s
}

func (s *ListImagesRequest) Validate() error {
	return dara.Validate(s)
}
