// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody
	GetImages() []*ListImagesResponseBodyImages
	SetNextToken(v string) *ListImagesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListImagesResponseBody
	GetRequestId() *string
}

type ListImagesResponseBody struct {
	// The image details.
	Images []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 3a6b93229825ac667104463b56790c91
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0FC4A1C7-421C-5EAB-9361-4C0338EFA287
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) GetImages() []*ListImagesResponseBodyImages {
	return s.Images
}

func (s *ListImagesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImagesResponseBody) SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListImagesResponseBody) SetNextToken(v string) *ListImagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImagesResponseBody) Validate() error {
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListImagesResponseBodyImages struct {
	// The architecture.
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The description.
	//
	// example:
	//
	// alibaba cloud linux 3 full for H800
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image ID.
	//
	// example:
	//
	// i190951671671438639388
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// CentOS_7.9_x86_64_FULL_20221110
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The image version.
	//
	// example:
	//
	// 7.9
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 20
	NodeCount *int64 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The platform.
	//
	// example:
	//
	// ALinux3
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The MD5 hash value of the file.
	//
	// example:
	//
	// 40741292480fc6d778138adcf8c
	ReleaseFileMd5 *string `json:"ReleaseFileMd5,omitempty" xml:"ReleaseFileMd5,omitempty"`
	// The image size.
	//
	// example:
	//
	// 5.8G
	ReleaseFileSize *string `json:"ReleaseFileSize,omitempty" xml:"ReleaseFileSize,omitempty"`
	// The image type.
	//
	// example:
	//
	// Public
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListImagesResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImages) GetArchitecture() *string {
	return s.Architecture
}

func (s *ListImagesResponseBodyImages) GetDescription() *string {
	return s.Description
}

func (s *ListImagesResponseBodyImages) GetImageId() *string {
	return s.ImageId
}

func (s *ListImagesResponseBodyImages) GetImageName() *string {
	return s.ImageName
}

func (s *ListImagesResponseBodyImages) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *ListImagesResponseBodyImages) GetNodeCount() *int64 {
	return s.NodeCount
}

func (s *ListImagesResponseBodyImages) GetPlatform() *string {
	return s.Platform
}

func (s *ListImagesResponseBodyImages) GetReleaseFileMd5() *string {
	return s.ReleaseFileMd5
}

func (s *ListImagesResponseBodyImages) GetReleaseFileSize() *string {
	return s.ReleaseFileSize
}

func (s *ListImagesResponseBodyImages) GetType() *string {
	return s.Type
}

func (s *ListImagesResponseBodyImages) SetArchitecture(v string) *ListImagesResponseBodyImages {
	s.Architecture = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetDescription(v string) *ListImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageId(v string) *ListImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageName(v string) *ListImagesResponseBodyImages {
	s.ImageName = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageVersion(v string) *ListImagesResponseBodyImages {
	s.ImageVersion = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetNodeCount(v int64) *ListImagesResponseBodyImages {
	s.NodeCount = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetPlatform(v string) *ListImagesResponseBodyImages {
	s.Platform = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetReleaseFileMd5(v string) *ListImagesResponseBodyImages {
	s.ReleaseFileMd5 = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetReleaseFileSize(v string) *ListImagesResponseBodyImages {
	s.ReleaseFileSize = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetType(v string) *ListImagesResponseBodyImages {
	s.Type = &v
	return s
}

func (s *ListImagesResponseBodyImages) Validate() error {
	return dara.Validate(s)
}
