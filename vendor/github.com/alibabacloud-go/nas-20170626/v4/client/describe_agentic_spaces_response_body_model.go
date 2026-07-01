// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticSpacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgenticSpaces(v *DescribeAgenticSpacesResponseBodyAgenticSpaces) *DescribeAgenticSpacesResponseBody
	GetAgenticSpaces() *DescribeAgenticSpacesResponseBodyAgenticSpaces
	SetNextToken(v string) *DescribeAgenticSpacesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeAgenticSpacesResponseBody
	GetRequestId() *string
}

type DescribeAgenticSpacesResponseBody struct {
	AgenticSpaces *DescribeAgenticSpacesResponseBodyAgenticSpaces `json:"AgenticSpaces,omitempty" xml:"AgenticSpaces,omitempty" type:"Struct"`
	// example:
	//
	// MTc2NTg1MTUyMzA1OTczNTc1OCM0NjQxMzQ****=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// BC7C825C-5F65-4B56-BEF6-98C56C7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAgenticSpacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticSpacesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAgenticSpacesResponseBody) GetAgenticSpaces() *DescribeAgenticSpacesResponseBodyAgenticSpaces {
	return s.AgenticSpaces
}

func (s *DescribeAgenticSpacesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAgenticSpacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAgenticSpacesResponseBody) SetAgenticSpaces(v *DescribeAgenticSpacesResponseBodyAgenticSpaces) *DescribeAgenticSpacesResponseBody {
	s.AgenticSpaces = v
	return s
}

func (s *DescribeAgenticSpacesResponseBody) SetNextToken(v string) *DescribeAgenticSpacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAgenticSpacesResponseBody) SetRequestId(v string) *DescribeAgenticSpacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAgenticSpacesResponseBody) Validate() error {
	if s.AgenticSpaces != nil {
		if err := s.AgenticSpaces.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAgenticSpacesResponseBodyAgenticSpaces struct {
	AgenticSpace []*DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace `json:"AgenticSpace,omitempty" xml:"AgenticSpace,omitempty" type:"Repeated"`
}

func (s DescribeAgenticSpacesResponseBodyAgenticSpaces) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticSpacesResponseBodyAgenticSpaces) GoString() string {
	return s.String()
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpaces) GetAgenticSpace() []*DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace {
	return s.AgenticSpace
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpaces) SetAgenticSpace(v []*DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) *DescribeAgenticSpacesResponseBodyAgenticSpaces {
	s.AgenticSpace = v
	return s
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpaces) Validate() error {
	if s.AgenticSpace != nil {
		for _, item := range s.AgenticSpace {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace struct {
	AgenticSpaceId *string                                                          `json:"AgenticSpaceId,omitempty" xml:"AgenticSpaceId,omitempty"`
	Azone          *string                                                          `json:"Azone,omitempty" xml:"Azone,omitempty"`
	CreateTimeUtc  *string                                                          `json:"CreateTimeUtc,omitempty" xml:"CreateTimeUtc,omitempty"`
	Description    *string                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	FileCountUsage *int64                                                           `json:"FileCountUsage,omitempty" xml:"FileCountUsage,omitempty"`
	FileSystemId   *string                                                          `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FileSystemPath *string                                                          `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	Quota          *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpaceQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
	SpaceUsage     *int64                                                           `json:"SpaceUsage,omitempty" xml:"SpaceUsage,omitempty"`
	Status         *string                                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTimeUtc  *string                                                          `json:"UpdateTimeUtc,omitempty" xml:"UpdateTimeUtc,omitempty"`
}

func (s DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) GoString() string {
	return s.String()
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) GetAgenticSpaceId() *string {
	return s.AgenticSpaceId
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) GetAzone() *string {
	return s.Azone
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) GetCreateTimeUtc() *string {
	return s.CreateTimeUtc
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) GetDescription() *string {
	return s.Description
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) GetFileCountUsage() *int64 {
	return s.FileCountUsage
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) GetFileSystemPath() *string {
	return s.FileSystemPath
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) GetQuota() *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpaceQuota {
	return s.Quota
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) GetSpaceUsage() *int64 {
	return s.SpaceUsage
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) GetStatus() *string {
	return s.Status
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) GetUpdateTimeUtc() *string {
	return s.UpdateTimeUtc
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) SetAgenticSpaceId(v string) *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace {
	s.AgenticSpaceId = &v
	return s
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) SetAzone(v string) *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace {
	s.Azone = &v
	return s
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) SetCreateTimeUtc(v string) *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace {
	s.CreateTimeUtc = &v
	return s
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) SetDescription(v string) *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace {
	s.Description = &v
	return s
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) SetFileCountUsage(v int64) *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace {
	s.FileCountUsage = &v
	return s
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) SetFileSystemId(v string) *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace {
	s.FileSystemId = &v
	return s
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) SetFileSystemPath(v string) *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace {
	s.FileSystemPath = &v
	return s
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) SetQuota(v *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpaceQuota) *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace {
	s.Quota = v
	return s
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) SetSpaceUsage(v int64) *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace {
	s.SpaceUsage = &v
	return s
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) SetStatus(v string) *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace {
	s.Status = &v
	return s
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) SetUpdateTimeUtc(v string) *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace {
	s.UpdateTimeUtc = &v
	return s
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpace) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpaceQuota struct {
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	SizeLimit      *int64 `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpaceQuota) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpaceQuota) GoString() string {
	return s.String()
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpaceQuota) GetFileCountLimit() *int64 {
	return s.FileCountLimit
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpaceQuota) GetSizeLimit() *int64 {
	return s.SizeLimit
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpaceQuota) SetFileCountLimit(v int64) *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpaceQuota {
	s.FileCountLimit = &v
	return s
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpaceQuota) SetSizeLimit(v int64) *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpaceQuota {
	s.SizeLimit = &v
	return s
}

func (s *DescribeAgenticSpacesResponseBodyAgenticSpacesAgenticSpaceQuota) Validate() error {
	return dara.Validate(s)
}
